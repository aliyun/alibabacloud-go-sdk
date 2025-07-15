// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveSnapshotNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveSnapshotNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveSnapshotNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveSnapshotNotifyConfigRequest
	GetRegionId() *string
}

type DescribeLiveSnapshotNotifyConfigRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.yourdomain***.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveSnapshotNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveSnapshotNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveSnapshotNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveSnapshotNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveSnapshotNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveSnapshotNotifyConfigRequest) SetDomainName(v string) *DescribeLiveSnapshotNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigRequest) SetOwnerId(v int64) *DescribeLiveSnapshotNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigRequest) SetRegionId(v string) *DescribeLiveSnapshotNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveSnapshotNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
