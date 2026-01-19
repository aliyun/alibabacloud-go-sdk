// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWebRtcInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *SaveWebRtcInfoRequest
	GetCallId() *string
	SetContent(v string) *SaveWebRtcInfoRequest
	GetContent() *string
	SetContentType(v string) *SaveWebRtcInfoRequest
	GetContentType() *string
	SetInstanceId(v string) *SaveWebRtcInfoRequest
	GetInstanceId() *string
	SetJobId(v string) *SaveWebRtcInfoRequest
	GetJobId() *string
}

type SaveWebRtcInfoRequest struct {
	// This parameter is required.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SaveWebRtcInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveWebRtcInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoRequest) GetCallId() *string {
	return s.CallId
}

func (s *SaveWebRtcInfoRequest) GetContent() *string {
	return s.Content
}

func (s *SaveWebRtcInfoRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SaveWebRtcInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveWebRtcInfoRequest) GetJobId() *string {
	return s.JobId
}

func (s *SaveWebRtcInfoRequest) SetCallId(v string) *SaveWebRtcInfoRequest {
	s.CallId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetContent(v string) *SaveWebRtcInfoRequest {
	s.Content = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetContentType(v string) *SaveWebRtcInfoRequest {
	s.ContentType = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetInstanceId(v string) *SaveWebRtcInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetJobId(v string) *SaveWebRtcInfoRequest {
	s.JobId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) Validate() error {
	return dara.Validate(s)
}
