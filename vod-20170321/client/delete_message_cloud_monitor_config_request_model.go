// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCloudMonitorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMessageCloudMonitorConfigRequest
	GetAppId() *string
	SetOwnerAccount(v string) *DeleteMessageCloudMonitorConfigRequest
	GetOwnerAccount() *string
}

type DeleteMessageCloudMonitorConfigRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteMessageCloudMonitorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCloudMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageCloudMonitorConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMessageCloudMonitorConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteMessageCloudMonitorConfigRequest) SetAppId(v string) *DeleteMessageCloudMonitorConfigRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMessageCloudMonitorConfigRequest) SetOwnerAccount(v string) *DeleteMessageCloudMonitorConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteMessageCloudMonitorConfigRequest) Validate() error {
	return dara.Validate(s)
}
