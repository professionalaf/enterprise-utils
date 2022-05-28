import path from 'path';
import { Code, Function, Runtime } from 'aws-cdk-lib/aws-lambda';
import { Construct } from 'constructs';

export interface PythonLambdaWithPrivatePypiProps {
  /**
   * --index-url for private pypi repo
   */
  readonly indexUrl: string;

  /**
   * --trusted-host for private repo
   */
  readonly trustedHost: string;
  /**
   * Relative path to function code
   */
  readonly codePath: string;
  /**
   * Handler location
   *
   * @default - 'index.handler'
   */
  readonly handler?: string;
}

export class PythonLambdaWithPrivatePypi extends Construct {
  public readonly indexUrl: string;
  public readonly trustedHost: string;
  public readonly codePath: string;
  public readonly handler?: string;

  constructor(scope: Construct, id: string, props: PythonLambdaWithPrivatePypiProps) {
    super(scope, id);

    this.indexUrl = props.indexUrl;
    this.trustedHost = props.trustedHost;
    this.codePath = props.codePath;
    this.handler = props.handler ? props.handler: 'index.handler';

    new Function(this, `${id}Lambda`, {
      runtime: Runtime.PYTHON_3_9,
      handler: this.handler,
      code: Code.fromAsset(path.join(__dirname, this.codePath), {
        bundling: {
          image: Runtime.PYTHON_3_9.bundlingImage,
          command: [
            'bash',
            '-c',
            'pip install ' +
            `--index-url ${this.indexUrl} ` +
            `--trusted-host ${this.trustedHost} ` +
            '-r requirements.txt ' +
            '-t /asset-output/python ' +
            '&& cp -au . /asset/output/python',
          ],
        },
      },
      ),
    });
  }
}