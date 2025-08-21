// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishTaskStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetPublishTaskStateRequest
	GetAgentKey() *string
	SetId(v int64) *GetPublishTaskStateRequest
	GetId() *int64
}

type GetPublishTaskStateRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetPublishTaskStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPublishTaskStateRequest) GoString() string {
	return s.String()
}

func (s *GetPublishTaskStateRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetPublishTaskStateRequest) GetId() *int64 {
	return s.Id
}

func (s *GetPublishTaskStateRequest) SetAgentKey(v string) *GetPublishTaskStateRequest {
	s.AgentKey = &v
	return s
}

func (s *GetPublishTaskStateRequest) SetId(v int64) *GetPublishTaskStateRequest {
	s.Id = &v
	return s
}

func (s *GetPublishTaskStateRequest) Validate() error {
	return dara.Validate(s)
}
