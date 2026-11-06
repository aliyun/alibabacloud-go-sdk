// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNumberBindingsOfInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNumberBindingsOfInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNumberBindingsOfInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListNumberBindingsOfInstanceResponseBody) *ListNumberBindingsOfInstanceResponse
	GetBody() *ListNumberBindingsOfInstanceResponseBody
}

type ListNumberBindingsOfInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNumberBindingsOfInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNumberBindingsOfInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNumberBindingsOfInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListNumberBindingsOfInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNumberBindingsOfInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNumberBindingsOfInstanceResponse) GetBody() *ListNumberBindingsOfInstanceResponseBody {
	return s.Body
}

func (s *ListNumberBindingsOfInstanceResponse) SetHeaders(v map[string]*string) *ListNumberBindingsOfInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListNumberBindingsOfInstanceResponse) SetStatusCode(v int32) *ListNumberBindingsOfInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNumberBindingsOfInstanceResponse) SetBody(v *ListNumberBindingsOfInstanceResponseBody) *ListNumberBindingsOfInstanceResponse {
	s.Body = v
	return s
}

func (s *ListNumberBindingsOfInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
