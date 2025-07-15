// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetTopicByIdRequest
	GetAgentKey() *string
	SetId(v string) *GetTopicByIdRequest
	GetId() *string
}

type GetTopicByIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 数据ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTopicByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTopicByIdRequest) GoString() string {
	return s.String()
}

func (s *GetTopicByIdRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetTopicByIdRequest) GetId() *string {
	return s.Id
}

func (s *GetTopicByIdRequest) SetAgentKey(v string) *GetTopicByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *GetTopicByIdRequest) SetId(v string) *GetTopicByIdRequest {
	s.Id = &v
	return s
}

func (s *GetTopicByIdRequest) Validate() error {
	return dara.Validate(s)
}
