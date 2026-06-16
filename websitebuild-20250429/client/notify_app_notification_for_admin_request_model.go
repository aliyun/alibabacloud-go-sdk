// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAppNotificationForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *NotifyAppNotificationForAdminRequest
	GetBizId() *string
	SetEnv(v string) *NotifyAppNotificationForAdminRequest
	GetEnv() *string
	SetSceneId(v string) *NotifyAppNotificationForAdminRequest
	GetSceneId() *string
}

type NotifyAppNotificationForAdminRequest struct {
	// example:
	//
	// WS20260206134746000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// staging
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 8c909373-6c33-41a7-aa38-3650e288a63e
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s NotifyAppNotificationForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s NotifyAppNotificationForAdminRequest) GoString() string {
	return s.String()
}

func (s *NotifyAppNotificationForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *NotifyAppNotificationForAdminRequest) GetEnv() *string {
	return s.Env
}

func (s *NotifyAppNotificationForAdminRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *NotifyAppNotificationForAdminRequest) SetBizId(v string) *NotifyAppNotificationForAdminRequest {
	s.BizId = &v
	return s
}

func (s *NotifyAppNotificationForAdminRequest) SetEnv(v string) *NotifyAppNotificationForAdminRequest {
	s.Env = &v
	return s
}

func (s *NotifyAppNotificationForAdminRequest) SetSceneId(v string) *NotifyAppNotificationForAdminRequest {
	s.SceneId = &v
	return s
}

func (s *NotifyAppNotificationForAdminRequest) Validate() error {
	return dara.Validate(s)
}
