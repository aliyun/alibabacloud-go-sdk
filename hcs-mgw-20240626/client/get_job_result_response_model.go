// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobResultResponse
	GetStatusCode() *int32
	SetBody(v *GetJobResultResponseBody) *GetJobResultResponse
	GetBody() *GetJobResultResponseBody
}

type GetJobResultResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetJobResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobResultResponse) GetBody() *GetJobResultResponseBody {
	return s.Body
}

func (s *GetJobResultResponse) SetHeaders(v map[string]*string) *GetJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetJobResultResponse) SetStatusCode(v int32) *GetJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResultResponse) SetBody(v *GetJobResultResponseBody) *GetJobResultResponse {
	s.Body = v
	return s
}

func (s *GetJobResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
