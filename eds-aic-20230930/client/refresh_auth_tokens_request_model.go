// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAuthTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireSeconds(v int64) *RefreshAuthTokensRequest
	GetExpireSeconds() *int64
	SetInstanceIds(v string) *RefreshAuthTokensRequest
	GetInstanceIds() *string
	SetLicenseKeys(v string) *RefreshAuthTokensRequest
	GetLicenseKeys() *string
}

type RefreshAuthTokensRequest struct {
	// The validity period in seconds.
	//
	// example:
	//
	// 600
	ExpireSeconds *int64 `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	// The list of instance IDs.
	//
	// example:
	//
	// acp-2zef0gov2nh2l3xxx,acp-2zef0gov2nh2l3yyy
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The list of license keys.
	//
	// example:
	//
	// lk-abcdef1234567890,lk-abcdef1234567891
	LicenseKeys *string `json:"LicenseKeys,omitempty" xml:"LicenseKeys,omitempty"`
}

func (s RefreshAuthTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAuthTokensRequest) GoString() string {
	return s.String()
}

func (s *RefreshAuthTokensRequest) GetExpireSeconds() *int64 {
	return s.ExpireSeconds
}

func (s *RefreshAuthTokensRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *RefreshAuthTokensRequest) GetLicenseKeys() *string {
	return s.LicenseKeys
}

func (s *RefreshAuthTokensRequest) SetExpireSeconds(v int64) *RefreshAuthTokensRequest {
	s.ExpireSeconds = &v
	return s
}

func (s *RefreshAuthTokensRequest) SetInstanceIds(v string) *RefreshAuthTokensRequest {
	s.InstanceIds = &v
	return s
}

func (s *RefreshAuthTokensRequest) SetLicenseKeys(v string) *RefreshAuthTokensRequest {
	s.LicenseKeys = &v
	return s
}

func (s *RefreshAuthTokensRequest) Validate() error {
	return dara.Validate(s)
}
