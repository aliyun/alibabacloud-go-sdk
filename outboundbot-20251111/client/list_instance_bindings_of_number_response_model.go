// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceBindingsOfNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceBindingsOfNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceBindingsOfNumberResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceBindingsOfNumberResponseBody) *ListInstanceBindingsOfNumberResponse
	GetBody() *ListInstanceBindingsOfNumberResponseBody
}

type ListInstanceBindingsOfNumberResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceBindingsOfNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceBindingsOfNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceBindingsOfNumberResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceBindingsOfNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceBindingsOfNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceBindingsOfNumberResponse) GetBody() *ListInstanceBindingsOfNumberResponseBody {
	return s.Body
}

func (s *ListInstanceBindingsOfNumberResponse) SetHeaders(v map[string]*string) *ListInstanceBindingsOfNumberResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceBindingsOfNumberResponse) SetStatusCode(v int32) *ListInstanceBindingsOfNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceBindingsOfNumberResponse) SetBody(v *ListInstanceBindingsOfNumberResponseBody) *ListInstanceBindingsOfNumberResponse {
	s.Body = v
	return s
}

func (s *ListInstanceBindingsOfNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
