// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteMaterialByIdRequest
	GetAgentKey() *string
	SetId(v int64) *DeleteMaterialByIdRequest
	GetId() *int64
}

type DeleteMaterialByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteMaterialByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteMaterialByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteMaterialByIdRequest) SetAgentKey(v string) *DeleteMaterialByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteMaterialByIdRequest) SetId(v int64) *DeleteMaterialByIdRequest {
	s.Id = &v
	return s
}

func (s *DeleteMaterialByIdRequest) Validate() error {
	return dara.Validate(s)
}
