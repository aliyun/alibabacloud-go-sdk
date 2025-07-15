// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLivePackageConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLivePackageConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLivePackageConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLivePackageConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLivePackageConfigRequest
	GetStreamName() *string
}

type DeleteLivePackageConfigRequest struct {
	// App name, `*` matches all names.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Live streaming domain (primary playback domain).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Live stream name, `*` matches all streams under AppName.
	//
	// This parameter is required.
	//
	// example:
	//
	// StreamName
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLivePackageConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLivePackageConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLivePackageConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLivePackageConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLivePackageConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLivePackageConfigRequest) SetAppName(v string) *DeleteLivePackageConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLivePackageConfigRequest) SetDomainName(v string) *DeleteLivePackageConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLivePackageConfigRequest) SetOwnerId(v int64) *DeleteLivePackageConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLivePackageConfigRequest) SetRegionId(v string) *DeleteLivePackageConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLivePackageConfigRequest) SetStreamName(v string) *DeleteLivePackageConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLivePackageConfigRequest) Validate() error {
	return dara.Validate(s)
}
