// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListJobExecutionsResponseBody) *ListJobExecutionsResponse
	GetBody() *ListJobExecutionsResponseBody
}

type ListJobExecutionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobExecutionsResponse) GetBody() *ListJobExecutionsResponseBody {
	return s.Body
}

func (s *ListJobExecutionsResponse) SetHeaders(v map[string]*string) *ListJobExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListJobExecutionsResponse) SetStatusCode(v int32) *ListJobExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobExecutionsResponse) SetBody(v *ListJobExecutionsResponseBody) *ListJobExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListJobExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
