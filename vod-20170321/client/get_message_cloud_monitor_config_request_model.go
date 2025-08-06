// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCloudMonitorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMessageCloudMonitorConfigRequest
	GetAppId() *string
	SetOwnerAccount(v string) *GetMessageCloudMonitorConfigRequest
	GetOwnerAccount() *string
}

type GetMessageCloudMonitorConfigRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s GetMessageCloudMonitorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCloudMonitorConfigRequest) GoString() string {
	return s.String()
}

func (s *GetMessageCloudMonitorConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMessageCloudMonitorConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetMessageCloudMonitorConfigRequest) SetAppId(v string) *GetMessageCloudMonitorConfigRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageCloudMonitorConfigRequest) SetOwnerAccount(v string) *GetMessageCloudMonitorConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMessageCloudMonitorConfigRequest) Validate() error {
	return dara.Validate(s)
}
