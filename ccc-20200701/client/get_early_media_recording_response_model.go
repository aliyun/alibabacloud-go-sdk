// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEarlyMediaRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEarlyMediaRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEarlyMediaRecordingResponse
	GetStatusCode() *int32
	SetBody(v *GetEarlyMediaRecordingResponseBody) *GetEarlyMediaRecordingResponse
	GetBody() *GetEarlyMediaRecordingResponseBody
}

type GetEarlyMediaRecordingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEarlyMediaRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEarlyMediaRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEarlyMediaRecordingResponse) GoString() string {
	return s.String()
}

func (s *GetEarlyMediaRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEarlyMediaRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEarlyMediaRecordingResponse) GetBody() *GetEarlyMediaRecordingResponseBody {
	return s.Body
}

func (s *GetEarlyMediaRecordingResponse) SetHeaders(v map[string]*string) *GetEarlyMediaRecordingResponse {
	s.Headers = v
	return s
}

func (s *GetEarlyMediaRecordingResponse) SetStatusCode(v int32) *GetEarlyMediaRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEarlyMediaRecordingResponse) SetBody(v *GetEarlyMediaRecordingResponseBody) *GetEarlyMediaRecordingResponse {
	s.Body = v
	return s
}

func (s *GetEarlyMediaRecordingResponse) Validate() error {
	return dara.Validate(s)
}
