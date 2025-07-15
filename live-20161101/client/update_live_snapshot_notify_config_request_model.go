// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateLiveSnapshotNotifyConfigRequest
	GetDomainName() *string
	SetNotifyAuthKey(v string) *UpdateLiveSnapshotNotifyConfigRequest
	GetNotifyAuthKey() *string
	SetNotifyReqAuth(v string) *UpdateLiveSnapshotNotifyConfigRequest
	GetNotifyReqAuth() *string
	SetNotifyUrl(v string) *UpdateLiveSnapshotNotifyConfigRequest
	GetNotifyUrl() *string
	SetOwnerId(v int64) *UpdateLiveSnapshotNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateLiveSnapshotNotifyConfigRequest
	GetRegionId() *string
}

type UpdateLiveSnapshotNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// ww.yourdomain***.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback authentication key. The key can be 16 to 32 characters in length and can contain only letters and digits.
	//
	// >  This parameter is required if you set the NotifyReqAuth parameter to **yes**.
	//
	// example:
	//
	// yourkey
	NotifyAuthKey *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// Specifies whether to enable callback authentication. Valid values:
	//
	// 	- **yes**: enables callback authentication
	//
	// 	- **no**: disables callback authentication
	//
	// >  Default value: **no**. If you set this parameter to **yes**, the NotifyAuthKey parameter is required.
	//
	// example:
	//
	// yes
	NotifyReqAuth *string `json:"NotifyReqAuth,omitempty" xml:"NotifyReqAuth,omitempty"`
	// The callback URL. Specify a valid URL that is up to 500 characters in length.
	//
	// example:
	//
	// http://callback.yourdomain***.com
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateLiveSnapshotNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) GetNotifyReqAuth() *string {
	return s.NotifyReqAuth
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) SetDomainName(v string) *UpdateLiveSnapshotNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) SetNotifyAuthKey(v string) *UpdateLiveSnapshotNotifyConfigRequest {
	s.NotifyAuthKey = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) SetNotifyReqAuth(v string) *UpdateLiveSnapshotNotifyConfigRequest {
	s.NotifyReqAuth = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) SetNotifyUrl(v string) *UpdateLiveSnapshotNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) SetOwnerId(v int64) *UpdateLiveSnapshotNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) SetRegionId(v string) *UpdateLiveSnapshotNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
