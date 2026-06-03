// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcknowledgeTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *AcknowledgeTaskResultRequest
	GetLang() *string
	SetTaskDetailNo(v []*string) *AcknowledgeTaskResultRequest
	GetTaskDetailNo() []*string
	SetUserClientIp(v string) *AcknowledgeTaskResultRequest
	GetUserClientIp() *string
}

type AcknowledgeTaskResultRequest struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	TaskDetailNo []*string `json:"TaskDetailNo,omitempty" xml:"TaskDetailNo,omitempty" type:"Repeated"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s AcknowledgeTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s AcknowledgeTaskResultRequest) GoString() string {
	return s.String()
}

func (s *AcknowledgeTaskResultRequest) GetLang() *string {
	return s.Lang
}

func (s *AcknowledgeTaskResultRequest) GetTaskDetailNo() []*string {
	return s.TaskDetailNo
}

func (s *AcknowledgeTaskResultRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *AcknowledgeTaskResultRequest) SetLang(v string) *AcknowledgeTaskResultRequest {
	s.Lang = &v
	return s
}

func (s *AcknowledgeTaskResultRequest) SetTaskDetailNo(v []*string) *AcknowledgeTaskResultRequest {
	s.TaskDetailNo = v
	return s
}

func (s *AcknowledgeTaskResultRequest) SetUserClientIp(v string) *AcknowledgeTaskResultRequest {
	s.UserClientIp = &v
	return s
}

func (s *AcknowledgeTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
