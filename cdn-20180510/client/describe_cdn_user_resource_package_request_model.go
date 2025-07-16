// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeCdnUserResourcePackageRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnUserResourcePackageRequest
	GetSecurityToken() *string
	SetStatus(v string) *DescribeCdnUserResourcePackageRequest
	GetStatus() *string
}

type DescribeCdnUserResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The status of the resource plan that you want to query. Valid values:
	//
	// 	- **valid**: valid
	//
	// 	- **closed**: expired
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdnUserResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserResourcePackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnUserResourcePackageRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnUserResourcePackageRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnUserResourcePackageRequest) SetOwnerId(v int64) *DescribeCdnUserResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnUserResourcePackageRequest) SetSecurityToken(v string) *DescribeCdnUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnUserResourcePackageRequest) SetStatus(v string) *DescribeCdnUserResourcePackageRequest {
	s.Status = &v
	return s
}

func (s *DescribeCdnUserResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
