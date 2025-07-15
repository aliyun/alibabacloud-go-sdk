// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AcceptChatRequest
	GetInstanceId() *string
	SetJobId(v string) *AcceptChatRequest
	GetJobId() *string
}

type AcceptChatRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chat-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s AcceptChatRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptChatRequest) GoString() string {
	return s.String()
}

func (s *AcceptChatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AcceptChatRequest) GetJobId() *string {
	return s.JobId
}

func (s *AcceptChatRequest) SetInstanceId(v string) *AcceptChatRequest {
	s.InstanceId = &v
	return s
}

func (s *AcceptChatRequest) SetJobId(v string) *AcceptChatRequest {
	s.JobId = &v
	return s
}

func (s *AcceptChatRequest) Validate() error {
	return dara.Validate(s)
}
