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
	// The name of the application to which the ingest stream belongs. You can view the AppName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
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
	// Specifies whether to disable stream ingest or streaming. Currently, only disabling stream ingest is supported: **publisher**.
	//
	// This parameter is required.
	//
	// example:
	//
	// publisher
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	// Specifies whether to only interrupt the stream without adding it to the blacklist. Valid values:
	//
	// - **yes**: Only interrupts the stream without adding it to the blacklist (supports upstream ingest or upstream streaming).
	//
	// - **no**: Interrupts the stream and adds it to the blacklist.
	//
	// > Default value: no.
	//
	// example:
	//
	// yes
	Oneshot *string `json:"Oneshot,omitempty" xml:"Oneshot,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time to resume the stream. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// > - If the **Oneshot*	- parameter is set to **no*	- and ResumeTime is not specified, the live stream is disabled for 6 months by default.
	//
	// > - If a value is specified, the restriction is lifted at the time specified by ResumeTime and the live stream is resumed.
	//
	// example:
	//
	// 2015-12-01T10:37:00Z
	ResumeTime *string `json:"ResumeTime,omitempty" xml:"ResumeTime,omitempty"`
	// The name of the ingest stream. You can view the StreamName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
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
