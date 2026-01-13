// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcInfoByVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcInfoByVpcResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcInfoByVpcResponseBody) *ListVpcInfoByVpcResponse
	GetBody() *ListVpcInfoByVpcResponseBody
}

type ListVpcInfoByVpcResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcInfoByVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcInfoByVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByVpcResponse) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcInfoByVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcInfoByVpcResponse) GetBody() *ListVpcInfoByVpcResponseBody {
	return s.Body
}

func (s *ListVpcInfoByVpcResponse) SetHeaders(v map[string]*string) *ListVpcInfoByVpcResponse {
	s.Headers = v
	return s
}

func (s *ListVpcInfoByVpcResponse) SetStatusCode(v int32) *ListVpcInfoByVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcInfoByVpcResponse) SetBody(v *ListVpcInfoByVpcResponseBody) *ListVpcInfoByVpcResponse {
	s.Body = v
	return s
}

func (s *ListVpcInfoByVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
