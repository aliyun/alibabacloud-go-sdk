// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnRegionAndIspRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeCdnRegionAndIspRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnRegionAndIspRequest
	GetSecurityToken() *string
}

type DescribeCdnRegionAndIspRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnRegionAndIspRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnRegionAndIspRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnRegionAndIspRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnRegionAndIspRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnRegionAndIspRequest) SetOwnerId(v int64) *DescribeCdnRegionAndIspRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnRegionAndIspRequest) SetSecurityToken(v string) *DescribeCdnRegionAndIspRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnRegionAndIspRequest) Validate() error {
	return dara.Validate(s)
}
