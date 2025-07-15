// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveSnapshotNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddLiveSnapshotNotifyConfigRequest
	GetDomainName() *string
	SetNotifyAuthKey(v string) *AddLiveSnapshotNotifyConfigRequest
	GetNotifyAuthKey() *string
	SetNotifyReqAuth(v string) *AddLiveSnapshotNotifyConfigRequest
	GetNotifyReqAuth() *string
	SetNotifyUrl(v string) *AddLiveSnapshotNotifyConfigRequest
	GetNotifyUrl() *string
	SetOwnerId(v int64) *AddLiveSnapshotNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddLiveSnapshotNotifyConfigRequest
	GetRegionId() *string
}

type AddLiveSnapshotNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.yourdomain***.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The callback authentication key. The key must be 16 to 32 characters in length and can contain only letters and digits.
	//
	// > This parameter is required if you set the NotifyReqAuth parameter to **yes**.
	//
	// example:
	//
	// yourkey
	NotifyAuthKey *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// Specifies whether to enable callback authentication. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no*	- (default)
	//
	// > This parameter is required if you set the NotifyAuthKey parameter to yes.
	//
	// example:
	//
	// yes
	NotifyReqAuth *string `json:"NotifyReqAuth,omitempty" xml:"NotifyReqAuth,omitempty"`
	// The callback URL. Specify a valid URL that is up to 500 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://callback.yourdomain***.com
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddLiveSnapshotNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveSnapshotNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *AddLiveSnapshotNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddLiveSnapshotNotifyConfigRequest) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *AddLiveSnapshotNotifyConfigRequest) GetNotifyReqAuth() *string {
	return s.NotifyReqAuth
}

func (s *AddLiveSnapshotNotifyConfigRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *AddLiveSnapshotNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveSnapshotNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveSnapshotNotifyConfigRequest) SetDomainName(v string) *AddLiveSnapshotNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigRequest) SetNotifyAuthKey(v string) *AddLiveSnapshotNotifyConfigRequest {
	s.NotifyAuthKey = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigRequest) SetNotifyReqAuth(v string) *AddLiveSnapshotNotifyConfigRequest {
	s.NotifyReqAuth = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigRequest) SetNotifyUrl(v string) *AddLiveSnapshotNotifyConfigRequest {
	s.NotifyUrl = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigRequest) SetOwnerId(v int64) *AddLiveSnapshotNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigRequest) SetRegionId(v string) *AddLiveSnapshotNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveSnapshotNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
