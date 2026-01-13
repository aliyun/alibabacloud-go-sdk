// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindInstance2VpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindInstance2VpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindInstance2VpcResponse
	GetStatusCode() *int32
	SetBody(v *BindInstance2VpcResponseBody) *BindInstance2VpcResponse
	GetBody() *BindInstance2VpcResponseBody
}

type BindInstance2VpcResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindInstance2VpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindInstance2VpcResponse) String() string {
	return dara.Prettify(s)
}

func (s BindInstance2VpcResponse) GoString() string {
	return s.String()
}

func (s *BindInstance2VpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindInstance2VpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindInstance2VpcResponse) GetBody() *BindInstance2VpcResponseBody {
	return s.Body
}

func (s *BindInstance2VpcResponse) SetHeaders(v map[string]*string) *BindInstance2VpcResponse {
	s.Headers = v
	return s
}

func (s *BindInstance2VpcResponse) SetStatusCode(v int32) *BindInstance2VpcResponse {
	s.StatusCode = &v
	return s
}

func (s *BindInstance2VpcResponse) SetBody(v *BindInstance2VpcResponseBody) *BindInstance2VpcResponse {
	s.Body = v
	return s
}

func (s *BindInstance2VpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
