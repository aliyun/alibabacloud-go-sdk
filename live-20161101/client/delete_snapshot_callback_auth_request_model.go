// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotCallbackAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteSnapshotCallbackAuthRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteSnapshotCallbackAuthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSnapshotCallbackAuthRequest
	GetRegionId() *string
}

type DeleteSnapshotCallbackAuthRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSnapshotCallbackAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotCallbackAuthRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotCallbackAuthRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteSnapshotCallbackAuthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSnapshotCallbackAuthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnapshotCallbackAuthRequest) SetDomainName(v string) *DeleteSnapshotCallbackAuthRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSnapshotCallbackAuthRequest) SetOwnerId(v int64) *DeleteSnapshotCallbackAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSnapshotCallbackAuthRequest) SetRegionId(v string) *DeleteSnapshotCallbackAuthRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotCallbackAuthRequest) Validate() error {
	return dara.Validate(s)
}
