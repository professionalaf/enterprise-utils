// enterprise-utils
package enterpriseutils

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/professionalaf/enterprise-utils/enterpriseutils/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/professionalaf/enterprise-utils/enterpriseutils/internal"
)

type PythonLambdaWithPrivatePypi interface {
	constructs.Construct
	CodePath() *string
	Handler() *string
	IndexUrl() *string
	// The tree node.
	Node() constructs.Node
	TrustedHost() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for PythonLambdaWithPrivatePypi
type jsiiProxy_PythonLambdaWithPrivatePypi struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PythonLambdaWithPrivatePypi) CodePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLambdaWithPrivatePypi) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLambdaWithPrivatePypi) IndexUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLambdaWithPrivatePypi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLambdaWithPrivatePypi) TrustedHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedHost",
		&returns,
	)
	return returns
}


func NewPythonLambdaWithPrivatePypi(scope constructs.Construct, id *string, props *PythonLambdaWithPrivatePypiProps) PythonLambdaWithPrivatePypi {
	_init_.Initialize()

	j := jsiiProxy_PythonLambdaWithPrivatePypi{}

	_jsii_.Create(
		"enterprise-utils.PythonLambdaWithPrivatePypi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPythonLambdaWithPrivatePypi_Override(p PythonLambdaWithPrivatePypi, scope constructs.Construct, id *string, props *PythonLambdaWithPrivatePypiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"enterprise-utils.PythonLambdaWithPrivatePypi",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PythonLambdaWithPrivatePypi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"enterprise-utils.PythonLambdaWithPrivatePypi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonLambdaWithPrivatePypi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type PythonLambdaWithPrivatePypiProps struct {
	// Relative path to function code.
	CodePath *string `field:"required" json:"codePath" yaml:"codePath"`
	// --index-url for private pypi repo.
	IndexUrl *string `field:"required" json:"indexUrl" yaml:"indexUrl"`
	// --trusted-host for private repo.
	TrustedHost *string `field:"required" json:"trustedHost" yaml:"trustedHost"`
	// Handler location.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
}

