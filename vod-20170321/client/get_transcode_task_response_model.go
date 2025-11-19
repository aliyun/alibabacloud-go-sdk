// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranscodeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranscodeTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetTranscodeTaskResponseBody) *GetTranscodeTaskResponse
	GetBody() *GetTranscodeTaskResponseBody
}

type GetTranscodeTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscodeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscodeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranscodeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranscodeTaskResponse) GetBody() *GetTranscodeTaskResponseBody {
	return s.Body
}

func (s *GetTranscodeTaskResponse) SetHeaders(v map[string]*string) *GetTranscodeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeTaskResponse) SetStatusCode(v int32) *GetTranscodeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscodeTaskResponse) SetBody(v *GetTranscodeTaskResponseBody) *GetTranscodeTaskResponse {
	s.Body = v
	return s
}

func (s *GetTranscodeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
