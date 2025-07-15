// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteModelRequest
	GetGroupId() *string
	SetModelName(v string) *DeleteModelRequest
	GetModelName() *string
}

type DeleteModelRequest struct {
	// The ID of the API group to which the model belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30e792398d6c4569b04c0e53a3494381
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the model.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *DeleteModelRequest) SetGroupId(v string) *DeleteModelRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteModelRequest) SetModelName(v string) *DeleteModelRequest {
	s.ModelName = &v
	return s
}

func (s *DeleteModelRequest) Validate() error {
	return dara.Validate(s)
}
