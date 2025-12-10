// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobResponse
	GetStatusCode() *int32
	SetBody(v *ListJobResponseBody) *ListJobResponse
	GetBody() *ListJobResponseBody
}

type ListJobResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobResponse) GoString() string {
	return s.String()
}

func (s *ListJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobResponse) GetBody() *ListJobResponseBody {
	return s.Body
}

func (s *ListJobResponse) SetHeaders(v map[string]*string) *ListJobResponse {
	s.Headers = v
	return s
}

func (s *ListJobResponse) SetStatusCode(v int32) *ListJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobResponse) SetBody(v *ListJobResponseBody) *ListJobResponse {
	s.Body = v
	return s
}

func (s *ListJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
