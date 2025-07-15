// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveShiftConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveShiftConfigsRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveShiftConfigsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveShiftConfigsRequest
	GetRegionId() *string
}

type DescribeLiveShiftConfigsRequest struct {
	// The streaming domain.
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

func (s DescribeLiveShiftConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveShiftConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveShiftConfigsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveShiftConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveShiftConfigsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveShiftConfigsRequest) SetDomainName(v string) *DescribeLiveShiftConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveShiftConfigsRequest) SetOwnerId(v int64) *DescribeLiveShiftConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveShiftConfigsRequest) SetRegionId(v string) *DescribeLiveShiftConfigsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveShiftConfigsRequest) Validate() error {
	return dara.Validate(s)
}
