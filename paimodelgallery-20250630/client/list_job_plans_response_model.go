// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListJobPlansResponseBody) *ListJobPlansResponse
	GetBody() *ListJobPlansResponseBody
}

type ListJobPlansResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobPlansResponse) GoString() string {
	return s.String()
}

func (s *ListJobPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobPlansResponse) GetBody() *ListJobPlansResponseBody {
	return s.Body
}

func (s *ListJobPlansResponse) SetHeaders(v map[string]*string) *ListJobPlansResponse {
	s.Headers = v
	return s
}

func (s *ListJobPlansResponse) SetStatusCode(v int32) *ListJobPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobPlansResponse) SetBody(v *ListJobPlansResponseBody) *ListJobPlansResponse {
	s.Body = v
	return s
}

func (s *ListJobPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
