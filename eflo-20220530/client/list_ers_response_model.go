// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListErsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListErsResponse
	GetStatusCode() *int32
	SetBody(v *ListErsResponseBody) *ListErsResponse
	GetBody() *ListErsResponseBody
}

type ListErsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListErsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListErsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListErsResponse) GoString() string {
	return s.String()
}

func (s *ListErsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListErsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListErsResponse) GetBody() *ListErsResponseBody {
	return s.Body
}

func (s *ListErsResponse) SetHeaders(v map[string]*string) *ListErsResponse {
	s.Headers = v
	return s
}

func (s *ListErsResponse) SetStatusCode(v int32) *ListErsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListErsResponse) SetBody(v *ListErsResponseBody) *ListErsResponse {
	s.Body = v
	return s
}

func (s *ListErsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
