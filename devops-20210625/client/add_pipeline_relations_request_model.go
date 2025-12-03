// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPipelineRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelObjectIds(v string) *AddPipelineRelationsRequest
	GetRelObjectIds() *string
	SetRelObjectType(v string) *AddPipelineRelationsRequest
	GetRelObjectType() *string
}

type AddPipelineRelationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11,22
	RelObjectIds *string `json:"relObjectIds,omitempty" xml:"relObjectIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VARIABLE_GROUP
	RelObjectType *string `json:"relObjectType,omitempty" xml:"relObjectType,omitempty"`
}

func (s AddPipelineRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPipelineRelationsRequest) GoString() string {
	return s.String()
}

func (s *AddPipelineRelationsRequest) GetRelObjectIds() *string {
	return s.RelObjectIds
}

func (s *AddPipelineRelationsRequest) GetRelObjectType() *string {
	return s.RelObjectType
}

func (s *AddPipelineRelationsRequest) SetRelObjectIds(v string) *AddPipelineRelationsRequest {
	s.RelObjectIds = &v
	return s
}

func (s *AddPipelineRelationsRequest) SetRelObjectType(v string) *AddPipelineRelationsRequest {
	s.RelObjectType = &v
	return s
}

func (s *AddPipelineRelationsRequest) Validate() error {
	return dara.Validate(s)
}
