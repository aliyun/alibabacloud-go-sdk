// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidLiveStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ForbidLiveStreamRequest
	GetAppName() *string
	SetDomainName(v string) *ForbidLiveStreamRequest
	GetDomainName() *string
	SetLiveStreamType(v string) *ForbidLiveStreamRequest
	GetLiveStreamType() *string
	SetOneshot(v string) *ForbidLiveStreamRequest
	GetOneshot() *string
	SetOwnerId(v int64) *ForbidLiveStreamRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ForbidLiveStreamRequest
	GetRegionId() *string
	SetResumeTime(v string) *ForbidLiveStreamRequest
	GetResumeTime() *string
	SetStreamName(v string) *ForbidLiveStreamRequest
	GetStreamName() *string
}

type ForbidLiveStreamRequest struct {
	// The name of the application to which the live stream belongs. You can view the application name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether the live stream is ingested by a streamer or played by a viewer. Set the value to **publisher**.
	//
	// This parameter is required.
	//
	// example:
	//
	// publisher
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	// Specifies whether to only interrupt the live stream without adding the ingest URL of the live stream to the blacklist. Valid values:
	//
	// 	- **yes**: interrupts the live stream but does not add the ingest URL of the live stream to the blacklist. This value is available only when the live stream is ingested or played in the upstream.
	//
	// 	- **no**: disables the live stream and adds the ingest URL of the live stream to the blacklist.
	//
	// >  If you do not specify this parameter, the default value no is used.
	//
	// example:
	//
	// yes
	Oneshot  *string `json:"Oneshot,omitempty" xml:"Oneshot,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the live stream is resumed. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >
	//
	// 	- If you set the **Oneshot*	- parameter to **no*	- and do not specify this parameter, the live stream is disabled for six months by default.
	//
	// 	- If you specify this parameter, the live stream is resumed at the specified point in time.
	//
	// example:
	//
	// 2015-12-01T10:37:00Z
	ResumeTime *string `json:"ResumeTime,omitempty" xml:"ResumeTime,omitempty"`
	// The name of the ingested stream. You can view the stream name on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page of the ApsaraVideo Live console.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ForbidLiveStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ForbidLiveStreamRequest) GoString() string {
	return s.String()
}

func (s *ForbidLiveStreamRequest) GetAppName() *string {
	return s.AppName
}

func (s *ForbidLiveStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ForbidLiveStreamRequest) GetLiveStreamType() *string {
	return s.LiveStreamType
}

func (s *ForbidLiveStreamRequest) GetOneshot() *string {
	return s.Oneshot
}

func (s *ForbidLiveStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ForbidLiveStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ForbidLiveStreamRequest) GetResumeTime() *string {
	return s.ResumeTime
}

func (s *ForbidLiveStreamRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *ForbidLiveStreamRequest) SetAppName(v string) *ForbidLiveStreamRequest {
	s.AppName = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetDomainName(v string) *ForbidLiveStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetLiveStreamType(v string) *ForbidLiveStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetOneshot(v string) *ForbidLiveStreamRequest {
	s.Oneshot = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetOwnerId(v int64) *ForbidLiveStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetRegionId(v string) *ForbidLiveStreamRequest {
	s.RegionId = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetResumeTime(v string) *ForbidLiveStreamRequest {
	s.ResumeTime = &v
	return s
}

func (s *ForbidLiveStreamRequest) SetStreamName(v string) *ForbidLiveStreamRequest {
	s.StreamName = &v
	return s
}

func (s *ForbidLiveStreamRequest) Validate() error {
	return dara.Validate(s)
}
