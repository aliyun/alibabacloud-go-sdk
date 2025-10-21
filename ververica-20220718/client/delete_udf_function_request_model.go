// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *DeleteUdfFunctionRequest
	GetClassName() *string
	SetUdfArtifactName(v string) *DeleteUdfFunctionRequest
	GetUdfArtifactName() *string
}

type DeleteUdfFunctionRequest struct {
	// The name of the class that corresponds to the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// Category
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// The name of the resource that corresponds to the UDF that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-udf
	UdfArtifactName *string `json:"udfArtifactName,omitempty" xml:"udfArtifactName,omitempty"`
}

func (s DeleteUdfFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteUdfFunctionRequest) GetClassName() *string {
	return s.ClassName
}

func (s *DeleteUdfFunctionRequest) GetUdfArtifactName() *string {
	return s.UdfArtifactName
}

func (s *DeleteUdfFunctionRequest) SetClassName(v string) *DeleteUdfFunctionRequest {
	s.ClassName = &v
	return s
}

func (s *DeleteUdfFunctionRequest) SetUdfArtifactName(v string) *DeleteUdfFunctionRequest {
	s.UdfArtifactName = &v
	return s
}

func (s *DeleteUdfFunctionRequest) Validate() error {
	return dara.Validate(s)
}
