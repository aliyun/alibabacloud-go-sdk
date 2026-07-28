// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindResourceControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindResourceControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindResourceControlResponse
	GetStatusCode() *int32
	SetBody(v *UnbindResourceControlResponseBody) *UnbindResourceControlResponse
	GetBody() *UnbindResourceControlResponseBody
}

type UnbindResourceControlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindResourceControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindResourceControlResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindResourceControlResponse) GoString() string {
	return s.String()
}

func (s *UnbindResourceControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindResourceControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindResourceControlResponse) GetBody() *UnbindResourceControlResponseBody {
	return s.Body
}

func (s *UnbindResourceControlResponse) SetHeaders(v map[string]*string) *UnbindResourceControlResponse {
	s.Headers = v
	return s
}

func (s *UnbindResourceControlResponse) SetStatusCode(v int32) *UnbindResourceControlResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindResourceControlResponse) SetBody(v *UnbindResourceControlResponseBody) *UnbindResourceControlResponse {
	s.Body = v
	return s
}

func (s *UnbindResourceControlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
