// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnUserResourcePackageRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnUserResourcePackageRequest
	GetSecurityToken() *string
	SetStatus(v string) *DescribeDcdnUserResourcePackageRequest
	GetStatus() *string
}

type DescribeDcdnUserResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The status of the resource plan. Valid values:
	//
	// 	- **valid**: valid
	//
	// 	- **closed**: expired
	//
	// 	- If you leave this parameter empty, all resource plans are queried.
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnUserResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserResourcePackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnUserResourcePackageRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnUserResourcePackageRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnUserResourcePackageRequest) SetOwnerId(v int64) *DescribeDcdnUserResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageRequest) SetSecurityToken(v string) *DescribeDcdnUserResourcePackageRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageRequest) SetStatus(v string) *DescribeDcdnUserResourcePackageRequest {
	s.Status = &v
	return s
}

func (s *DescribeDcdnUserResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
