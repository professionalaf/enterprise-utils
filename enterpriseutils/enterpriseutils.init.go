package enterpriseutils

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"enterprise-utils.PythonLambdaWithPrivatePypi",
		reflect.TypeOf((*PythonLambdaWithPrivatePypi)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "codePath", GoGetter: "CodePath"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "indexUrl", GoGetter: "IndexUrl"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedHost", GoGetter: "TrustedHost"},
		},
		func() interface{} {
			j := jsiiProxy_PythonLambdaWithPrivatePypi{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"enterprise-utils.PythonLambdaWithPrivatePypiProps",
		reflect.TypeOf((*PythonLambdaWithPrivatePypiProps)(nil)).Elem(),
	)
}
