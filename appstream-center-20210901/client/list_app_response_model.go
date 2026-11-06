// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppResponse
	GetStatusCode() *int32
	SetBody(v *ListAppResponseBody) *ListAppResponse
	GetBody() *ListAppResponseBody
}

type ListAppResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppResponse) GoString() string {
	return s.String()
}

func (s *ListAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppResponse) GetBody() *ListAppResponseBody {
	return s.Body
}

func (s *ListAppResponse) SetHeaders(v map[string]*string) *ListAppResponse {
	s.Headers = v
	return s
}

func (s *ListAppResponse) SetStatusCode(v int32) *ListAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppResponse) SetBody(v *ListAppResponseBody) *ListAppResponse {
	s.Body = v
	return s
}

func (s *ListAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
