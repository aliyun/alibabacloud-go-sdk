// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveSnapshotNotifyConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveSnapshotNotifyConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveSnapshotNotifyConfigRequest
	GetRegionId() *string
}

type DeleteLiveSnapshotNotifyConfigRequest struct {
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

func (s DeleteLiveSnapshotNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotNotifyConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveSnapshotNotifyConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveSnapshotNotifyConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveSnapshotNotifyConfigRequest) SetDomainName(v string) *DeleteLiveSnapshotNotifyConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveSnapshotNotifyConfigRequest) SetOwnerId(v int64) *DeleteLiveSnapshotNotifyConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveSnapshotNotifyConfigRequest) SetRegionId(v string) *DeleteLiveSnapshotNotifyConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveSnapshotNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}
