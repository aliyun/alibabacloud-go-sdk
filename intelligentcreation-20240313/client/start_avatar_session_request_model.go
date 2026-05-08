// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAvatarSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelToken(v string) *StartAvatarSessionRequest
	GetChannelToken() *string
	SetCustomPushUrl(v string) *StartAvatarSessionRequest
	GetCustomPushUrl() *string
	SetCustomUserId(v string) *StartAvatarSessionRequest
	GetCustomUserId() *string
	SetProjectId(v string) *StartAvatarSessionRequest
	GetProjectId() *string
	SetRequestId(v string) *StartAvatarSessionRequest
	GetRequestId() *string
}

type StartAvatarSessionRequest struct {
	ChannelToken  *string `json:"channelToken,omitempty" xml:"channelToken,omitempty"`
	CustomPushUrl *string `json:"customPushUrl,omitempty" xml:"customPushUrl,omitempty"`
	CustomUserId  *string `json:"customUserId,omitempty" xml:"customUserId,omitempty"`
	// example:
	//
	// 13534711288320
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 15ED6083-B0B8-5B2A-BEDB-94A5C687C812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartAvatarSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAvatarSessionRequest) GoString() string {
	return s.String()
}

func (s *StartAvatarSessionRequest) GetChannelToken() *string {
	return s.ChannelToken
}

func (s *StartAvatarSessionRequest) GetCustomPushUrl() *string {
	return s.CustomPushUrl
}

func (s *StartAvatarSessionRequest) GetCustomUserId() *string {
	return s.CustomUserId
}

func (s *StartAvatarSessionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *StartAvatarSessionRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAvatarSessionRequest) SetChannelToken(v string) *StartAvatarSessionRequest {
	s.ChannelToken = &v
	return s
}

func (s *StartAvatarSessionRequest) SetCustomPushUrl(v string) *StartAvatarSessionRequest {
	s.CustomPushUrl = &v
	return s
}

func (s *StartAvatarSessionRequest) SetCustomUserId(v string) *StartAvatarSessionRequest {
	s.CustomUserId = &v
	return s
}

func (s *StartAvatarSessionRequest) SetProjectId(v string) *StartAvatarSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *StartAvatarSessionRequest) SetRequestId(v string) *StartAvatarSessionRequest {
	s.RequestId = &v
	return s
}

func (s *StartAvatarSessionRequest) Validate() error {
	return dara.Validate(s)
}
