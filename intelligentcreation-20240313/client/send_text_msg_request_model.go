// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTextMsgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *SendTextMsgRequest
	GetProjectId() *string
	SetRequestId(v string) *SendTextMsgRequest
	GetRequestId() *string
	SetSessionId(v string) *SendTextMsgRequest
	GetSessionId() *string
	SetText(v string) *SendTextMsgRequest
	GetText() *string
	SetType(v int32) *SendTextMsgRequest
	GetType() *int32
}

type SendTextMsgRequest struct {
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	Type      *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendTextMsgRequest) String() string {
	return dara.Prettify(s)
}

func (s SendTextMsgRequest) GoString() string {
	return s.String()
}

func (s *SendTextMsgRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SendTextMsgRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *SendTextMsgRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendTextMsgRequest) GetText() *string {
	return s.Text
}

func (s *SendTextMsgRequest) GetType() *int32 {
	return s.Type
}

func (s *SendTextMsgRequest) SetProjectId(v string) *SendTextMsgRequest {
	s.ProjectId = &v
	return s
}

func (s *SendTextMsgRequest) SetRequestId(v string) *SendTextMsgRequest {
	s.RequestId = &v
	return s
}

func (s *SendTextMsgRequest) SetSessionId(v string) *SendTextMsgRequest {
	s.SessionId = &v
	return s
}

func (s *SendTextMsgRequest) SetText(v string) *SendTextMsgRequest {
	s.Text = &v
	return s
}

func (s *SendTextMsgRequest) SetType(v int32) *SendTextMsgRequest {
	s.Type = &v
	return s
}

func (s *SendTextMsgRequest) Validate() error {
	return dara.Validate(s)
}
