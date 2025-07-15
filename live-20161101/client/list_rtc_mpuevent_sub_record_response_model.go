// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRtcMPUEventSubRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRtcMPUEventSubRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRtcMPUEventSubRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListRtcMPUEventSubRecordResponseBody) *ListRtcMPUEventSubRecordResponse
	GetBody() *ListRtcMPUEventSubRecordResponseBody
}

type ListRtcMPUEventSubRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRtcMPUEventSubRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRtcMPUEventSubRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUEventSubRecordResponse) GoString() string {
	return s.String()
}

func (s *ListRtcMPUEventSubRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRtcMPUEventSubRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRtcMPUEventSubRecordResponse) GetBody() *ListRtcMPUEventSubRecordResponseBody {
	return s.Body
}

func (s *ListRtcMPUEventSubRecordResponse) SetHeaders(v map[string]*string) *ListRtcMPUEventSubRecordResponse {
	s.Headers = v
	return s
}

func (s *ListRtcMPUEventSubRecordResponse) SetStatusCode(v int32) *ListRtcMPUEventSubRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRtcMPUEventSubRecordResponse) SetBody(v *ListRtcMPUEventSubRecordResponseBody) *ListRtcMPUEventSubRecordResponse {
	s.Body = v
	return s
}

func (s *ListRtcMPUEventSubRecordResponse) Validate() error {
	return dara.Validate(s)
}
