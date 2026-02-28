// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonoRecordingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentChannelId(v string) *ListMonoRecordingsRequest
	GetAgentChannelId() *string
	SetAgentId(v string) *ListMonoRecordingsRequest
	GetAgentId() *string
	SetContactId(v string) *ListMonoRecordingsRequest
	GetContactId() *string
	SetInstanceId(v string) *ListMonoRecordingsRequest
	GetInstanceId() *string
}

type ListMonoRecordingsRequest struct {
	// example:
	//
	// ch-user-****-****-1772180844645-job-*****
	AgentChannelId *string `json:"AgentChannelId,omitempty" xml:"AgentChannelId,omitempty"`
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-25697383427137****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListMonoRecordingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMonoRecordingsRequest) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsRequest) GetAgentChannelId() *string {
	return s.AgentChannelId
}

func (s *ListMonoRecordingsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMonoRecordingsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *ListMonoRecordingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMonoRecordingsRequest) SetAgentChannelId(v string) *ListMonoRecordingsRequest {
	s.AgentChannelId = &v
	return s
}

func (s *ListMonoRecordingsRequest) SetAgentId(v string) *ListMonoRecordingsRequest {
	s.AgentId = &v
	return s
}

func (s *ListMonoRecordingsRequest) SetContactId(v string) *ListMonoRecordingsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMonoRecordingsRequest) SetInstanceId(v string) *ListMonoRecordingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMonoRecordingsRequest) Validate() error {
	return dara.Validate(s)
}
