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
	// example:
	//
	// 126000030
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 52775239-1575-5C07-A4AE-1835D120E4A6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// tcm9xac9dsfbfgm8hf5k94l3cqybwh9o3mn0iuyytdgd9qoejxf1crxsdvuvr8fu0zudk5px4vsa3e3fgcclplkiuo7kyy3sqgscvhejmooblaiv64ww8cvlxvin2urzyhooqj33y7gvodef0sxn22n9q58o7xlupabiknxsv46qe7kof8nuc4be8kyhi01
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
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
