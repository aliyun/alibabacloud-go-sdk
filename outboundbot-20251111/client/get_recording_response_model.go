// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecordingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecordingResponse
	GetStatusCode() *int32
	SetBody(v *GetRecordingResponseBody) *GetRecordingResponse
	GetBody() *GetRecordingResponseBody
}

type GetRecordingResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecordingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecordingResponse) GoString() string {
	return s.String()
}

func (s *GetRecordingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecordingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecordingResponse) GetBody() *GetRecordingResponseBody {
	return s.Body
}

func (s *GetRecordingResponse) SetHeaders(v map[string]*string) *GetRecordingResponse {
	s.Headers = v
	return s
}

func (s *GetRecordingResponse) SetStatusCode(v int32) *GetRecordingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordingResponse) SetBody(v *GetRecordingResponseBody) *GetRecordingResponse {
	s.Body = v
	return s
}

func (s *GetRecordingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
