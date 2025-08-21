// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVsStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ResumeVsStreamRequest
	GetAppName() *string
	SetControlStreamAction(v string) *ResumeVsStreamRequest
	GetControlStreamAction() *string
	SetDomainName(v string) *ResumeVsStreamRequest
	GetDomainName() *string
	SetLiveStreamType(v string) *ResumeVsStreamRequest
	GetLiveStreamType() *string
	SetOwnerId(v int64) *ResumeVsStreamRequest
	GetOwnerId() *int64
	SetStreamName(v string) *ResumeVsStreamRequest
	GetStreamName() *string
}

type ResumeVsStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxApp
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
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ResumeVsStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeVsStreamRequest) GoString() string {
	return s.String()
}

func (s *ResumeVsStreamRequest) GetAppName() *string {
	return s.AppName
}

func (s *ResumeVsStreamRequest) GetControlStreamAction() *string {
	return s.ControlStreamAction
}

func (s *ResumeVsStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ResumeVsStreamRequest) GetLiveStreamType() *string {
	return s.LiveStreamType
}

func (s *ResumeVsStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResumeVsStreamRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *ResumeVsStreamRequest) SetAppName(v string) *ResumeVsStreamRequest {
	s.AppName = &v
	return s
}

func (s *ResumeVsStreamRequest) SetControlStreamAction(v string) *ResumeVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *ResumeVsStreamRequest) SetDomainName(v string) *ResumeVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ResumeVsStreamRequest) SetLiveStreamType(v string) *ResumeVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ResumeVsStreamRequest) SetOwnerId(v int64) *ResumeVsStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeVsStreamRequest) SetStreamName(v string) *ResumeVsStreamRequest {
	s.StreamName = &v
	return s
}

func (s *ResumeVsStreamRequest) Validate() error {
	return dara.Validate(s)
}
