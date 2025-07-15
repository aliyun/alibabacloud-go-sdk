// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySnapshotCallbackAuthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QuerySnapshotCallbackAuthRequest
	GetDomainName() *string
	SetOwnerId(v int64) *QuerySnapshotCallbackAuthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QuerySnapshotCallbackAuthRequest
	GetRegionId() *string
}

type QuerySnapshotCallbackAuthRequest struct {
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

func (s QuerySnapshotCallbackAuthRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotCallbackAuthRequest) GoString() string {
	return s.String()
}

func (s *QuerySnapshotCallbackAuthRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QuerySnapshotCallbackAuthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySnapshotCallbackAuthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QuerySnapshotCallbackAuthRequest) SetDomainName(v string) *QuerySnapshotCallbackAuthRequest {
	s.DomainName = &v
	return s
}

func (s *QuerySnapshotCallbackAuthRequest) SetOwnerId(v int64) *QuerySnapshotCallbackAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySnapshotCallbackAuthRequest) SetRegionId(v string) *QuerySnapshotCallbackAuthRequest {
	s.RegionId = &v
	return s
}

func (s *QuerySnapshotCallbackAuthRequest) Validate() error {
	return dara.Validate(s)
}
