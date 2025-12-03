// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelObjectType(v string) *ListPipelineRelationsRequest
	GetRelObjectType() *string
}

type ListPipelineRelationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// VARIABLE_GROUP
	RelObjectType *string `json:"relObjectType,omitempty" xml:"relObjectType,omitempty"`
}

func (s ListPipelineRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRelationsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineRelationsRequest) GetRelObjectType() *string {
	return s.RelObjectType
}

func (s *ListPipelineRelationsRequest) SetRelObjectType(v string) *ListPipelineRelationsRequest {
	s.RelObjectType = &v
	return s
}

func (s *ListPipelineRelationsRequest) Validate() error {
	return dara.Validate(s)
}
