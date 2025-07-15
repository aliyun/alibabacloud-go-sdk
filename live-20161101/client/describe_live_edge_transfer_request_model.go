// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveEdgeTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveEdgeTransferRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveEdgeTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveEdgeTransferRequest
	GetRegionId() *string
}

type DescribeLiveEdgeTransferRequest struct {
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveEdgeTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveEdgeTransferRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveEdgeTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveEdgeTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveEdgeTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveEdgeTransferRequest) SetDomainName(v string) *DescribeLiveEdgeTransferRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveEdgeTransferRequest) SetOwnerId(v int64) *DescribeLiveEdgeTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveEdgeTransferRequest) SetRegionId(v string) *DescribeLiveEdgeTransferRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveEdgeTransferRequest) Validate() error {
	return dara.Validate(s)
}
