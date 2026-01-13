// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindInstance2VpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindInstance2VpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindInstance2VpcResponse
	GetStatusCode() *int32
	SetBody(v *UnbindInstance2VpcResponseBody) *UnbindInstance2VpcResponse
	GetBody() *UnbindInstance2VpcResponseBody
}

type UnbindInstance2VpcResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindInstance2VpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindInstance2VpcResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindInstance2VpcResponse) GoString() string {
	return s.String()
}

func (s *UnbindInstance2VpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindInstance2VpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindInstance2VpcResponse) GetBody() *UnbindInstance2VpcResponseBody {
	return s.Body
}

func (s *UnbindInstance2VpcResponse) SetHeaders(v map[string]*string) *UnbindInstance2VpcResponse {
	s.Headers = v
	return s
}

func (s *UnbindInstance2VpcResponse) SetStatusCode(v int32) *UnbindInstance2VpcResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindInstance2VpcResponse) SetBody(v *UnbindInstance2VpcResponseBody) *UnbindInstance2VpcResponse {
	s.Body = v
	return s
}

func (s *UnbindInstance2VpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
