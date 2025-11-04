// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranscodeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranscodeJobResponse
	GetStatusCode() *int32
	SetBody(v *GetTranscodeJobResponseBody) *GetTranscodeJobResponse
	GetBody() *GetTranscodeJobResponseBody
}

type GetTranscodeJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscodeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscodeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeJobResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranscodeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranscodeJobResponse) GetBody() *GetTranscodeJobResponseBody {
	return s.Body
}

func (s *GetTranscodeJobResponse) SetHeaders(v map[string]*string) *GetTranscodeJobResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeJobResponse) SetStatusCode(v int32) *GetTranscodeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscodeJobResponse) SetBody(v *GetTranscodeJobResponseBody) *GetTranscodeJobResponse {
	s.Body = v
	return s
}

func (s *GetTranscodeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
