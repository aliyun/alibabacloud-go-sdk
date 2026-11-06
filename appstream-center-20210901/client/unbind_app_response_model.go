// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAppResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAppResponseBody) *UnbindAppResponse
	GetBody() *UnbindAppResponseBody
}

type UnbindAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAppResponse) GoString() string {
	return s.String()
}

func (s *UnbindAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAppResponse) GetBody() *UnbindAppResponseBody {
	return s.Body
}

func (s *UnbindAppResponse) SetHeaders(v map[string]*string) *UnbindAppResponse {
	s.Headers = v
	return s
}

func (s *UnbindAppResponse) SetStatusCode(v int32) *UnbindAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAppResponse) SetBody(v *UnbindAppResponseBody) *UnbindAppResponse {
	s.Body = v
	return s
}

func (s *UnbindAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
