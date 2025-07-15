// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaterialByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetMaterialByIdRequest
	GetAgentKey() *string
	SetId(v int64) *GetMaterialByIdRequest
	GetId() *int64
}

type GetMaterialByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 67c520d1fa43455ea44fb69fa402d54d_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMaterialByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMaterialByIdRequest) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetMaterialByIdRequest) GetId() *int64 {
	return s.Id
}

func (s *GetMaterialByIdRequest) SetAgentKey(v string) *GetMaterialByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *GetMaterialByIdRequest) SetId(v int64) *GetMaterialByIdRequest {
	s.Id = &v
	return s
}

func (s *GetMaterialByIdRequest) Validate() error {
	return dara.Validate(s)
}
