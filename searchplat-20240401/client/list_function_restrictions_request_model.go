// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionRestrictionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelType(v string) *ListFunctionRestrictionsRequest
	GetModelType() *string
	SetRegion(v string) *ListFunctionRestrictionsRequest
	GetRegion() *string
	SetSource(v string) *ListFunctionRestrictionsRequest
	GetSource() *string
}

type ListFunctionRestrictionsRequest struct {
	// The model type.
	//
	// example:
	//
	// native
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The source.
	//
	// example:
	//
	// user
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListFunctionRestrictionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionRestrictionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionRestrictionsRequest) GetModelType() *string {
	return s.ModelType
}

func (s *ListFunctionRestrictionsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListFunctionRestrictionsRequest) GetSource() *string {
	return s.Source
}

func (s *ListFunctionRestrictionsRequest) SetModelType(v string) *ListFunctionRestrictionsRequest {
	s.ModelType = &v
	return s
}

func (s *ListFunctionRestrictionsRequest) SetRegion(v string) *ListFunctionRestrictionsRequest {
	s.Region = &v
	return s
}

func (s *ListFunctionRestrictionsRequest) SetSource(v string) *ListFunctionRestrictionsRequest {
	s.Source = &v
	return s
}

func (s *ListFunctionRestrictionsRequest) Validate() error {
	return dara.Validate(s)
}
