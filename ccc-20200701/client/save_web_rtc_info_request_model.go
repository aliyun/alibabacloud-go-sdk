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
	//
	// example:
	//
	// e13c9740-1e37-123b-21b6-00163e352f9
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"media_source":{},"remote_inbound_rtp":{},"outbound_rtp":{},"inbound_rtp":{},"remote_outbound_rtp":{},"candidate":{},"basic":{"callId":"e13c9740-1e37-123b-21b6-00163e352f9","timestamp":"1647262108395","callStartTime":"1647262108393","uid":"user-test","access_point":"shanghai","browser":"90","ip":"127.0.0.1"}}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-b8b0ca63-330c-4e65-8ae3-9de2c7ce7683
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
