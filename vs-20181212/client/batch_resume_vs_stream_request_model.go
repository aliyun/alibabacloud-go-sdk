// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchResumeVsStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *BatchResumeVsStreamRequest
	GetChannel() *string
	SetControlStreamAction(v string) *BatchResumeVsStreamRequest
	GetControlStreamAction() *string
	SetDomainName(v string) *BatchResumeVsStreamRequest
	GetDomainName() *string
	SetLiveStreamType(v string) *BatchResumeVsStreamRequest
	GetLiveStreamType() *string
	SetOwnerId(v int64) *BatchResumeVsStreamRequest
	GetOwnerId() *int64
}

type BatchResumeVsStreamRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// live/stream1
	Channel             *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ControlStreamAction *string `json:"ControlStreamAction,omitempty" xml:"ControlStreamAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// publisher
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchResumeVsStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchResumeVsStreamRequest) GoString() string {
	return s.String()
}

func (s *BatchResumeVsStreamRequest) GetChannel() *string {
	return s.Channel
}

func (s *BatchResumeVsStreamRequest) GetControlStreamAction() *string {
	return s.ControlStreamAction
}

func (s *BatchResumeVsStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchResumeVsStreamRequest) GetLiveStreamType() *string {
	return s.LiveStreamType
}

func (s *BatchResumeVsStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchResumeVsStreamRequest) SetChannel(v string) *BatchResumeVsStreamRequest {
	s.Channel = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetControlStreamAction(v string) *BatchResumeVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetDomainName(v string) *BatchResumeVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetLiveStreamType(v string) *BatchResumeVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *BatchResumeVsStreamRequest) SetOwnerId(v int64) *BatchResumeVsStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchResumeVsStreamRequest) Validate() error {
	return dara.Validate(s)
}
