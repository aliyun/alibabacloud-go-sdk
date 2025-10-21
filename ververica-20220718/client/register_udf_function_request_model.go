// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUdfFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *RegisterUdfFunctionRequest
	GetClassName() *string
	SetFunctionName(v string) *RegisterUdfFunctionRequest
	GetFunctionName() *string
	SetUdfArtifactName(v string) *RegisterUdfFunctionRequest
	GetUdfArtifactName() *string
}

type RegisterUdfFunctionRequest struct {
	// The name of the class that corresponds to the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// orderRank
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The name of the UDF. In most cases, the name of the UDF is the same as the class name. You can specify a name for the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// orderRank
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The name of the JAR or Python file that corresponds to the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-udf
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s RegisterUdfFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterUdfFunctionRequest) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionRequest) GetClassName() *string {
	return s.ClassName
}

func (s *RegisterUdfFunctionRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *RegisterUdfFunctionRequest) GetUdfArtifactName() *string {
	return s.UdfArtifactName
}

func (s *RegisterUdfFunctionRequest) SetClassName(v string) *RegisterUdfFunctionRequest {
	s.ClassName = &v
	return s
}

func (s *RegisterUdfFunctionRequest) SetFunctionName(v string) *RegisterUdfFunctionRequest {
	s.FunctionName = &v
	return s
}

func (s *RegisterUdfFunctionRequest) SetUdfArtifactName(v string) *RegisterUdfFunctionRequest {
	s.UdfArtifactName = &v
	return s
}

func (s *RegisterUdfFunctionRequest) Validate() error {
	return dara.Validate(s)
}
