// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSerialNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindSerialNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindSerialNumberResponse
	GetStatusCode() *int32
	SetBody(v *UnbindSerialNumberResponseBody) *UnbindSerialNumberResponse
	GetBody() *UnbindSerialNumberResponseBody
}

type UnbindSerialNumberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindSerialNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindSerialNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindSerialNumberResponse) GoString() string {
	return s.String()
}

func (s *UnbindSerialNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindSerialNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindSerialNumberResponse) GetBody() *UnbindSerialNumberResponseBody {
	return s.Body
}

func (s *UnbindSerialNumberResponse) SetHeaders(v map[string]*string) *UnbindSerialNumberResponse {
	s.Headers = v
	return s
}

func (s *UnbindSerialNumberResponse) SetStatusCode(v int32) *UnbindSerialNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindSerialNumberResponse) SetBody(v *UnbindSerialNumberResponseBody) *UnbindSerialNumberResponse {
	s.Body = v
	return s
}

func (s *UnbindSerialNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
