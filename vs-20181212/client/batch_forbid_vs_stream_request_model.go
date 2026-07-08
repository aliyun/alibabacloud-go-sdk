// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchForbidVsStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *BatchForbidVsStreamRequest
	GetChannel() *string
	SetControlStreamAction(v string) *BatchForbidVsStreamRequest
	GetControlStreamAction() *string
	SetDomainName(v string) *BatchForbidVsStreamRequest
	GetDomainName() *string
	SetLiveStreamType(v string) *BatchForbidVsStreamRequest
	GetLiveStreamType() *string
	SetOneshot(v string) *BatchForbidVsStreamRequest
	GetOneshot() *string
	SetOwnerId(v int64) *BatchForbidVsStreamRequest
	GetOwnerId() *int64
	SetResumeTime(v string) *BatchForbidVsStreamRequest
	GetResumeTime() *string
}

type BatchForbidVsStreamRequest struct {
	// The stream name.
	//
	// > - Format: AppName/StreamName.
	//
	// >
	//
	// > - Specify multiple names, separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// live/stream1,live/stream2
	Channel             *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	ControlStreamAction *string `json:"ControlStreamAction,omitempty" xml:"ControlStreamAction,omitempty"`
	// Your accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether the operation applies to stream ingest by a streamer or stream pulling by a client. Valid values:
	//
	// - publisher (streamer ingest)
	//
	// > Only publisher is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// publisher
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	// Specifies whether to stop ingest without adding the stream to the blacklist. Valid values:
	//
	// - yes
	//
	// - no
	//
	// example:
	//
	// yes
	Oneshot *string `json:"Oneshot,omitempty" xml:"Oneshot,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time to resume ingest.
	//
	// > Use UTC format. Example: 2015-12-01T17:37:00Z
	//
	// example:
	//
	// 2015-12-01T17:37:00Z
	ResumeTime *string `json:"ResumeTime,omitempty" xml:"ResumeTime,omitempty"`
}

func (s BatchForbidVsStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchForbidVsStreamRequest) GoString() string {
	return s.String()
}

func (s *BatchForbidVsStreamRequest) GetChannel() *string {
	return s.Channel
}

func (s *BatchForbidVsStreamRequest) GetControlStreamAction() *string {
	return s.ControlStreamAction
}

func (s *BatchForbidVsStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *BatchForbidVsStreamRequest) GetLiveStreamType() *string {
	return s.LiveStreamType
}

func (s *BatchForbidVsStreamRequest) GetOneshot() *string {
	return s.Oneshot
}

func (s *BatchForbidVsStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchForbidVsStreamRequest) GetResumeTime() *string {
	return s.ResumeTime
}

func (s *BatchForbidVsStreamRequest) SetChannel(v string) *BatchForbidVsStreamRequest {
	s.Channel = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetControlStreamAction(v string) *BatchForbidVsStreamRequest {
	s.ControlStreamAction = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetDomainName(v string) *BatchForbidVsStreamRequest {
	s.DomainName = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetLiveStreamType(v string) *BatchForbidVsStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetOneshot(v string) *BatchForbidVsStreamRequest {
	s.Oneshot = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetOwnerId(v int64) *BatchForbidVsStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchForbidVsStreamRequest) SetResumeTime(v string) *BatchForbidVsStreamRequest {
	s.ResumeTime = &v
	return s
}

func (s *BatchForbidVsStreamRequest) Validate() error {
	return dara.Validate(s)
}
