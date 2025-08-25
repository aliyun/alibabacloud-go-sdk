// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePushMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RevokePushMessageRequest
	GetAppId() *string
	SetMessageId(v string) *RevokePushMessageRequest
	GetMessageId() *string
	SetTargetId(v string) *RevokePushMessageRequest
	GetTargetId() *string
	SetTenantId(v string) *RevokePushMessageRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *RevokePushMessageRequest
	GetWorkspaceId() *string
}

type RevokePushMessageRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// This parameter is required.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RevokePushMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokePushMessageRequest) GoString() string {
	return s.String()
}

func (s *RevokePushMessageRequest) GetAppId() *string {
	return s.AppId
}

func (s *RevokePushMessageRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *RevokePushMessageRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *RevokePushMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevokePushMessageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RevokePushMessageRequest) SetAppId(v string) *RevokePushMessageRequest {
	s.AppId = &v
	return s
}

func (s *RevokePushMessageRequest) SetMessageId(v string) *RevokePushMessageRequest {
	s.MessageId = &v
	return s
}

func (s *RevokePushMessageRequest) SetTargetId(v string) *RevokePushMessageRequest {
	s.TargetId = &v
	return s
}

func (s *RevokePushMessageRequest) SetTenantId(v string) *RevokePushMessageRequest {
	s.TenantId = &v
	return s
}

func (s *RevokePushMessageRequest) SetWorkspaceId(v string) *RevokePushMessageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RevokePushMessageRequest) Validate() error {
	return dara.Validate(s)
}
