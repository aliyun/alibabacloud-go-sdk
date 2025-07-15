// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RejectChatRequest
	GetInstanceId() *string
	SetJobId(v string) *RejectChatRequest
	GetJobId() *string
}

type RejectChatRequest struct {
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

func (s RejectChatRequest) String() string {
	return dara.Prettify(s)
}

func (s RejectChatRequest) GoString() string {
	return s.String()
}

func (s *RejectChatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RejectChatRequest) GetJobId() *string {
	return s.JobId
}

func (s *RejectChatRequest) SetInstanceId(v string) *RejectChatRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectChatRequest) SetJobId(v string) *RejectChatRequest {
	s.JobId = &v
	return s
}

func (s *RejectChatRequest) Validate() error {
	return dara.Validate(s)
}
