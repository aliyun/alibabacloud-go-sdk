// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidVsStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ForbidVsStreamRequest
	GetAppName() *string
	SetControlStreamAction(v string) *ForbidVsStreamRequest
	GetControlStreamAction() *string
	SetDomainName(v string) *ForbidVsStreamRequest
	GetDomainName() *string
	SetLiveStreamType(v string) *ForbidVsStreamRequest
	GetLiveStreamType() *string
	SetOneshot(v string) *ForbidVsStreamRequest
	GetOneshot() *string
	SetOwnerId(v int64) *ForbidVsStreamRequest
	GetOwnerId() *int64
	SetResumeTime(v string) *ForbidVsStreamRequest
	GetResumeTime() *string
	SetStreamName(v string) *ForbidVsStreamRequest
	GetStreamName() *string
}

type ForbidVsStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxApp
	AppName             *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ControlStreamAction *string `json:"ControlStreamAction,omitempty" xml:"ControlStreamAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// publisher
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	// example:
	//
	// yes
	Oneshot *string `json:"Oneshot,omitempty" xml:"Oneshot,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 2015-12-01T17:37:00Z
	ResumeTime *string `json:"ResumeTime,omitempty" xml:"ResumeTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ForbidVsStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ForbidVsStreamRequest) GoString() string {
	return s.String()
}

func (s *ForbidVsStreamRequest) GetAppName() *string {
	return s.AppName
}

func (s *ForbidVsStreamRequest) GetControlStreamAction() *string {
	return s.ControlStreamAction
}

func (s *ForbidVsStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ForbidVsStreamRequest) GetLiveStreamType() *string {
	return s.LiveStreamType
}

func (s *ForbidVsStreamRequest) GetOneshot() *string {
	return s.Oneshot
}

func (s *ForbidVsStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ForbidVsStreamRequest) GetResumeTime() *string {
	return s.ResumeTime
}

func (s *ForbidVsStreamRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *ForbidVsStreamRequest) SetAppName(v string) *ForbidVsStreamRequest {
	s.AppName = &v
	return s
}

func (s *ForbidVsStreamRequest) SetControlStreamAction(v string) *ForbidVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *ForbidVsStreamRequest) SetDomainName(v string) *ForbidVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ForbidVsStreamRequest) SetLiveStreamType(v string) *ForbidVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ForbidVsStreamRequest) SetOneshot(v string) *ForbidVsStreamRequest {
	s.Oneshot = &v
	return s
}

func (s *ForbidVsStreamRequest) SetOwnerId(v int64) *ForbidVsStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *ForbidVsStreamRequest) SetResumeTime(v string) *ForbidVsStreamRequest {
	s.ResumeTime = &v
	return s
}

func (s *ForbidVsStreamRequest) SetStreamName(v string) *ForbidVsStreamRequest {
	s.StreamName = &v
	return s
}

func (s *ForbidVsStreamRequest) Validate() error {
	return dara.Validate(s)
}
