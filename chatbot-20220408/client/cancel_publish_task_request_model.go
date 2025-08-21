// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPublishTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CancelPublishTaskRequest
	GetAgentKey() *string
	SetId(v int64) *CancelPublishTaskRequest
	GetId() *int64
}

type CancelPublishTaskRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 8521
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CancelPublishTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPublishTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelPublishTaskRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CancelPublishTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *CancelPublishTaskRequest) SetAgentKey(v string) *CancelPublishTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelPublishTaskRequest) SetId(v int64) *CancelPublishTaskRequest {
	s.Id = &v
	return s
}

func (s *CancelPublishTaskRequest) Validate() error {
	return dara.Validate(s)
}
