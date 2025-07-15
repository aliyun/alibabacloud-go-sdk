// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveStreamBlockRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveStreamBlockRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveStreamBlockRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveStreamBlockRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLiveStreamBlockRequest
	GetStreamName() *string
}

type DeleteLiveStreamBlockRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
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
	// The name of the live stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveStreamBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamBlockRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamBlockRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveStreamBlockRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveStreamBlockRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveStreamBlockRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLiveStreamBlockRequest) SetAppName(v string) *DeleteLiveStreamBlockRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveStreamBlockRequest) SetDomainName(v string) *DeleteLiveStreamBlockRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveStreamBlockRequest) SetOwnerId(v int64) *DeleteLiveStreamBlockRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamBlockRequest) SetRegionId(v string) *DeleteLiveStreamBlockRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveStreamBlockRequest) SetStreamName(v string) *DeleteLiveStreamBlockRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLiveStreamBlockRequest) Validate() error {
	return dara.Validate(s)
}
