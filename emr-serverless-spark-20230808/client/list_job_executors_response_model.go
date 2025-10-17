// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobExecutorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobExecutorsResponse
	GetStatusCode() *int32
	SetBody(v *ListJobExecutorsResponseBody) *ListJobExecutorsResponse
	GetBody() *ListJobExecutorsResponseBody
}

type ListJobExecutorsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobExecutorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsResponse) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobExecutorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobExecutorsResponse) GetBody() *ListJobExecutorsResponseBody {
	return s.Body
}

func (s *ListJobExecutorsResponse) SetHeaders(v map[string]*string) *ListJobExecutorsResponse {
	s.Headers = v
	return s
}

func (s *ListJobExecutorsResponse) SetStatusCode(v int32) *ListJobExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobExecutorsResponse) SetBody(v *ListJobExecutorsResponseBody) *ListJobExecutorsResponse {
	s.Body = v
	return s
}

func (s *ListJobExecutorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
