// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRegionAndIspRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnRegionAndIspRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnRegionAndIspRequest
	GetSecurityToken() *string
}

type DescribeDcdnRegionAndIspRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnRegionAndIspRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRegionAndIspRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnRegionAndIspRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnRegionAndIspRequest) SetOwnerId(v int64) *DescribeDcdnRegionAndIspRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnRegionAndIspRequest) SetSecurityToken(v string) *DescribeDcdnRegionAndIspRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnRegionAndIspRequest) Validate() error {
	return dara.Validate(s)
}
