// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRtcMPUTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRtcMPUTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRtcMPUTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *ListRtcMPUTaskDetailResponseBody) *ListRtcMPUTaskDetailResponse
	GetBody() *ListRtcMPUTaskDetailResponseBody
}

type ListRtcMPUTaskDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRtcMPUTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRtcMPUTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRtcMPUTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRtcMPUTaskDetailResponse) GetBody() *ListRtcMPUTaskDetailResponseBody {
	return s.Body
}

func (s *ListRtcMPUTaskDetailResponse) SetHeaders(v map[string]*string) *ListRtcMPUTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *ListRtcMPUTaskDetailResponse) SetStatusCode(v int32) *ListRtcMPUTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponse) SetBody(v *ListRtcMPUTaskDetailResponseBody) *ListRtcMPUTaskDetailResponse {
	s.Body = v
	return s
}

func (s *ListRtcMPUTaskDetailResponse) Validate() error {
	return dara.Validate(s)
}
