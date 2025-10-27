// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *GetJobInstanceListResponseBody) *GetJobInstanceListResponse
	GetBody() *GetJobInstanceListResponseBody
}

type GetJobInstanceListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobInstanceListResponse) GetBody() *GetJobInstanceListResponseBody {
	return s.Body
}

func (s *GetJobInstanceListResponse) SetHeaders(v map[string]*string) *GetJobInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetJobInstanceListResponse) SetStatusCode(v int32) *GetJobInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobInstanceListResponse) SetBody(v *GetJobInstanceListResponseBody) *GetJobInstanceListResponse {
	s.Body = v
	return s
}

func (s *GetJobInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
