// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListViberServiceMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListViberServiceMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListViberServiceMessageResponse
	GetStatusCode() *int32
	SetBody(v *ListViberServiceMessageResponseBody) *ListViberServiceMessageResponse
	GetBody() *ListViberServiceMessageResponseBody
}

type ListViberServiceMessageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListViberServiceMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListViberServiceMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListViberServiceMessageResponse) GoString() string {
	return s.String()
}

func (s *ListViberServiceMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListViberServiceMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListViberServiceMessageResponse) GetBody() *ListViberServiceMessageResponseBody {
	return s.Body
}

func (s *ListViberServiceMessageResponse) SetHeaders(v map[string]*string) *ListViberServiceMessageResponse {
	s.Headers = v
	return s
}

func (s *ListViberServiceMessageResponse) SetStatusCode(v int32) *ListViberServiceMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListViberServiceMessageResponse) SetBody(v *ListViberServiceMessageResponseBody) *ListViberServiceMessageResponse {
	s.Body = v
	return s
}

func (s *ListViberServiceMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
