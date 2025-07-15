// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveVerifyContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveVerifyContentRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveVerifyContentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveVerifyContentRequest
	GetRegionId() *string
}

type DescribeLiveVerifyContentRequest struct {
	// The domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveVerifyContentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveVerifyContentRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveVerifyContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveVerifyContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveVerifyContentRequest) SetDomainName(v string) *DescribeLiveVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveVerifyContentRequest) SetOwnerId(v int64) *DescribeLiveVerifyContentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveVerifyContentRequest) SetRegionId(v string) *DescribeLiveVerifyContentRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveVerifyContentRequest) Validate() error {
	return dara.Validate(s)
}
