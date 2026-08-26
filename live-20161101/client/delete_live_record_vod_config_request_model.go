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
	// The AppName of the live stream. View AppNames on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
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
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stream. View StreamNames on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
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
