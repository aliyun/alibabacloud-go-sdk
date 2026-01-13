// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcInfoByInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcInfoByInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcInfoByInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcInfoByInstanceResponseBody) *ListVpcInfoByInstanceResponse
	GetBody() *ListVpcInfoByInstanceResponseBody
}

type ListVpcInfoByInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcInfoByInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcInfoByInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcInfoByInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListVpcInfoByInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcInfoByInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcInfoByInstanceResponse) GetBody() *ListVpcInfoByInstanceResponseBody {
	return s.Body
}

func (s *ListVpcInfoByInstanceResponse) SetHeaders(v map[string]*string) *ListVpcInfoByInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListVpcInfoByInstanceResponse) SetStatusCode(v int32) *ListVpcInfoByInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcInfoByInstanceResponse) SetBody(v *ListVpcInfoByInstanceResponseBody) *ListVpcInfoByInstanceResponse {
	s.Body = v
	return s
}

func (s *ListVpcInfoByInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
