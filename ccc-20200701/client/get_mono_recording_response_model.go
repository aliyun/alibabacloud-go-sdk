// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonoRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonoRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonoRecordingResponse
	GetStatusCode() *int32
	SetBody(v *GetMonoRecordingResponseBody) *GetMonoRecordingResponse
	GetBody() *GetMonoRecordingResponseBody
}

type GetMonoRecordingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMonoRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonoRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonoRecordingResponse) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonoRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonoRecordingResponse) GetBody() *GetMonoRecordingResponseBody {
	return s.Body
}

func (s *GetMonoRecordingResponse) SetHeaders(v map[string]*string) *GetMonoRecordingResponse {
	s.Headers = v
	return s
}

func (s *GetMonoRecordingResponse) SetStatusCode(v int32) *GetMonoRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonoRecordingResponse) SetBody(v *GetMonoRecordingResponseBody) *GetMonoRecordingResponse {
	s.Body = v
	return s
}

func (s *GetMonoRecordingResponse) Validate() error {
	return dara.Validate(s)
}
