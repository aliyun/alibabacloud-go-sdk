// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserVpcResponse
	GetStatusCode() *int32
	SetBody(v *ListUserVpcResponseBody) *ListUserVpcResponse
	GetBody() *ListUserVpcResponseBody
}

type ListUserVpcResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserVpcResponse) GoString() string {
	return s.String()
}

func (s *ListUserVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserVpcResponse) GetBody() *ListUserVpcResponseBody {
	return s.Body
}

func (s *ListUserVpcResponse) SetHeaders(v map[string]*string) *ListUserVpcResponse {
	s.Headers = v
	return s
}

func (s *ListUserVpcResponse) SetStatusCode(v int32) *ListUserVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserVpcResponse) SetBody(v *ListUserVpcResponseBody) *ListUserVpcResponse {
	s.Body = v
	return s
}

func (s *ListUserVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
