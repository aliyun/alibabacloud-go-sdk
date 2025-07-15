// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordVodConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteLiveRecordVodConfigRequest
	GetAppName() *string
	SetDomainName(v string) *DeleteLiveRecordVodConfigRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DeleteLiveRecordVodConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveRecordVodConfigRequest
	GetRegionId() *string
	SetStreamName(v string) *DeleteLiveRecordVodConfigRequest
	GetStreamName() *string
}

type DeleteLiveRecordVodConfigRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The main streaming domain.
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
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DeleteLiveRecordVodConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordVodConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordVodConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteLiveRecordVodConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveRecordVodConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveRecordVodConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveRecordVodConfigRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DeleteLiveRecordVodConfigRequest) SetAppName(v string) *DeleteLiveRecordVodConfigRequest {
	s.AppName = &v
	return s
}

func (s *DeleteLiveRecordVodConfigRequest) SetDomainName(v string) *DeleteLiveRecordVodConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveRecordVodConfigRequest) SetOwnerId(v int64) *DeleteLiveRecordVodConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveRecordVodConfigRequest) SetRegionId(v string) *DeleteLiveRecordVodConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveRecordVodConfigRequest) SetStreamName(v string) *DeleteLiveRecordVodConfigRequest {
	s.StreamName = &v
	return s
}

func (s *DeleteLiveRecordVodConfigRequest) Validate() error {
	return dara.Validate(s)
}
