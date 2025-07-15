// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SetLiveStreamBlockRequest
	GetAppName() *string
	SetBlockType(v string) *SetLiveStreamBlockRequest
	GetBlockType() *string
	SetDomainName(v string) *SetLiveStreamBlockRequest
	GetDomainName() *string
	SetLocationList(v string) *SetLiveStreamBlockRequest
	GetLocationList() *string
	SetOwnerId(v int64) *SetLiveStreamBlockRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLiveStreamBlockRequest
	GetRegionId() *string
	SetReleaseTime(v string) *SetLiveStreamBlockRequest
	GetReleaseTime() *string
	SetStreamName(v string) *SetLiveStreamBlockRequest
	GetStreamName() *string
}

type SetLiveStreamBlockRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The blocking type. Valid values: blacklist and whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// blacklist
	BlockType *string `json:"BlockType,omitempty" xml:"BlockType,omitempty"`
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The blocked region. If you specify multiple regions, such as CN and AS, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// CN
	LocationList *string `json:"LocationList,omitempty" xml:"LocationList,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the blocking ends. The time must be in UTC. If you do not specify this parameter, the blocking is valid for 7 days by default.
	//
	// example:
	//
	// 2016-06-29T19:00:00Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s SetLiveStreamBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamBlockRequest) GoString() string {
	return s.String()
}

func (s *SetLiveStreamBlockRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetLiveStreamBlockRequest) GetBlockType() *string {
	return s.BlockType
}

func (s *SetLiveStreamBlockRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetLiveStreamBlockRequest) GetLocationList() *string {
	return s.LocationList
}

func (s *SetLiveStreamBlockRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLiveStreamBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLiveStreamBlockRequest) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *SetLiveStreamBlockRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *SetLiveStreamBlockRequest) SetAppName(v string) *SetLiveStreamBlockRequest {
	s.AppName = &v
	return s
}

func (s *SetLiveStreamBlockRequest) SetBlockType(v string) *SetLiveStreamBlockRequest {
	s.BlockType = &v
	return s
}

func (s *SetLiveStreamBlockRequest) SetDomainName(v string) *SetLiveStreamBlockRequest {
	s.DomainName = &v
	return s
}

func (s *SetLiveStreamBlockRequest) SetLocationList(v string) *SetLiveStreamBlockRequest {
	s.LocationList = &v
	return s
}

func (s *SetLiveStreamBlockRequest) SetOwnerId(v int64) *SetLiveStreamBlockRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLiveStreamBlockRequest) SetRegionId(v string) *SetLiveStreamBlockRequest {
	s.RegionId = &v
	return s
}

func (s *SetLiveStreamBlockRequest) SetReleaseTime(v string) *SetLiveStreamBlockRequest {
	s.ReleaseTime = &v
	return s
}

func (s *SetLiveStreamBlockRequest) SetStreamName(v string) *SetLiveStreamBlockRequest {
	s.StreamName = &v
	return s
}

func (s *SetLiveStreamBlockRequest) Validate() error {
	return dara.Validate(s)
}
