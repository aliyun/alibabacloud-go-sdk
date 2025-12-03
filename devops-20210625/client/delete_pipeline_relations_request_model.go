// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRelObjectId(v string) *DeletePipelineRelationsRequest
	GetRelObjectId() *string
	SetRelObjectType(v string) *DeletePipelineRelationsRequest
	GetRelObjectType() *string
}

type DeletePipelineRelationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11
	RelObjectId *string `json:"relObjectId,omitempty" xml:"relObjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VARIABLE_GROUP
	RelObjectType *string `json:"relObjectType,omitempty" xml:"relObjectType,omitempty"`
}

func (s DeletePipelineRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineRelationsRequest) GoString() string {
	return s.String()
}

func (s *DeletePipelineRelationsRequest) GetRelObjectId() *string {
	return s.RelObjectId
}

func (s *DeletePipelineRelationsRequest) GetRelObjectType() *string {
	return s.RelObjectType
}

func (s *DeletePipelineRelationsRequest) SetRelObjectId(v string) *DeletePipelineRelationsRequest {
	s.RelObjectId = &v
	return s
}

func (s *DeletePipelineRelationsRequest) SetRelObjectType(v string) *DeletePipelineRelationsRequest {
	s.RelObjectType = &v
	return s
}

func (s *DeletePipelineRelationsRequest) Validate() error {
	return dara.Validate(s)
}
