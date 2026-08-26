// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeLiveStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ResumeLiveStreamRequest
	GetAppName() *string
	SetDomainName(v string) *ResumeLiveStreamRequest
	GetDomainName() *string
	SetLiveStreamType(v string) *ResumeLiveStreamRequest
	GetLiveStreamType() *string
	SetOwnerId(v int64) *ResumeLiveStreamRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *ResumeLiveStreamRequest
	GetSecurityToken() *string
	SetStreamName(v string) *ResumeLiveStreamRequest
	GetStreamName() *string
}

type ResumeLiveStreamRequest struct {
	// The name of the application to which the ingest stream belongs. You can set this parameter to 	- to indicate all AppName values (no restriction on AppName). You can view the AppName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
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
	// Specifies whether the stream is a streamer ingest stream or a client playback stream. Currently supported value: **publisher*	- (streamer ingest).
	//
	// This parameter is required.
	//
	// example:
	//
	// publisher
	LiveStreamType *string `json:"LiveStreamType,omitempty" xml:"LiveStreamType,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the ingest stream. You can view the StreamName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ResumeLiveStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeLiveStreamRequest) GoString() string {
	return s.String()
}

func (s *ResumeLiveStreamRequest) GetAppName() *string {
	return s.AppName
}

func (s *ResumeLiveStreamRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ResumeLiveStreamRequest) GetLiveStreamType() *string {
	return s.LiveStreamType
}

func (s *ResumeLiveStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResumeLiveStreamRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ResumeLiveStreamRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *ResumeLiveStreamRequest) SetAppName(v string) *ResumeLiveStreamRequest {
	s.AppName = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetDomainName(v string) *ResumeLiveStreamRequest {
	s.DomainName = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetLiveStreamType(v string) *ResumeLiveStreamRequest {
	s.LiveStreamType = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetOwnerId(v int64) *ResumeLiveStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetSecurityToken(v string) *ResumeLiveStreamRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResumeLiveStreamRequest) SetStreamName(v string) *ResumeLiveStreamRequest {
	s.StreamName = &v
	return s
}

func (s *ResumeLiveStreamRequest) Validate() error {
	return dara.Validate(s)
}
