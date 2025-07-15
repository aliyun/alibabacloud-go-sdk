// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveEdgeTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveEdgeTransferRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveEdgeTransferRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveEdgeTransferRequest
	GetRegionId() *string
}

type DeleteLiveEdgeTransferRequest struct {
	// The ingest domain. You can set only one stream relay configuration for an ingest domain.
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

func (s DeleteLiveEdgeTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveEdgeTransferRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveEdgeTransferRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveEdgeTransferRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveEdgeTransferRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveEdgeTransferRequest) SetDomainName(v string) *DeleteLiveEdgeTransferRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveEdgeTransferRequest) SetOwnerId(v int64) *DeleteLiveEdgeTransferRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveEdgeTransferRequest) SetRegionId(v string) *DeleteLiveEdgeTransferRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveEdgeTransferRequest) Validate() error {
	return dara.Validate(s)
}
