// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiChannelRecordingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentChannelId(v string) *ListMultiChannelRecordingsRequest
	GetAgentChannelId() *string
	SetAgentId(v string) *ListMultiChannelRecordingsRequest
	GetAgentId() *string
	SetContactId(v string) *ListMultiChannelRecordingsRequest
	GetContactId() *string
	SetInstanceId(v string) *ListMultiChannelRecordingsRequest
	GetInstanceId() *string
}

type ListMultiChannelRecordingsRequest struct {
	// example:
	//
	// ch-user-****-****-1772180844645-job-******
	AgentChannelId *string `json:"AgentChannelId,omitempty" xml:"AgentChannelId,omitempty"`
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-25920271311543****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListMultiChannelRecordingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiChannelRecordingsRequest) GoString() string {
	return s.String()
}

func (s *ListMultiChannelRecordingsRequest) GetAgentChannelId() *string {
	return s.AgentChannelId
}

func (s *ListMultiChannelRecordingsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMultiChannelRecordingsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListMultiChannelRecordingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMultiChannelRecordingsRequest) SetAgentChannelId(v string) *ListMultiChannelRecordingsRequest {
	s.AgentChannelId = &v
	return s
}

func (s *ListMultiChannelRecordingsRequest) SetAgentId(v string) *ListMultiChannelRecordingsRequest {
	s.AgentId = &v
	return s
}

func (s *ListMultiChannelRecordingsRequest) SetContactId(v string) *ListMultiChannelRecordingsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMultiChannelRecordingsRequest) SetInstanceId(v string) *ListMultiChannelRecordingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMultiChannelRecordingsRequest) Validate() error {
	return dara.Validate(s)
}
