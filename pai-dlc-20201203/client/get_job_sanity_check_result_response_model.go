// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobSanityCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobSanityCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobSanityCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *GetJobSanityCheckResultResponseBody) *GetJobSanityCheckResultResponse
	GetBody() *GetJobSanityCheckResultResponseBody
}

type GetJobSanityCheckResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobSanityCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobSanityCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobSanityCheckResultResponse) GoString() string {
	return s.String()
}

func (s *GetJobSanityCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobSanityCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobSanityCheckResultResponse) GetBody() *GetJobSanityCheckResultResponseBody {
	return s.Body
}

func (s *GetJobSanityCheckResultResponse) SetHeaders(v map[string]*string) *GetJobSanityCheckResultResponse {
	s.Headers = v
	return s
}

func (s *GetJobSanityCheckResultResponse) SetStatusCode(v int32) *GetJobSanityCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobSanityCheckResultResponse) SetBody(v *GetJobSanityCheckResultResponseBody) *GetJobSanityCheckResultResponse {
	s.Body = v
	return s
}

func (s *GetJobSanityCheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
