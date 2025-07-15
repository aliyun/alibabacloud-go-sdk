// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDefaultVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDefaultVpcResponse
	GetStatusCode() *int32
	SetBody(v *CreateDefaultVpcResponseBody) *CreateDefaultVpcResponse
	GetBody() *CreateDefaultVpcResponseBody
}

type CreateDefaultVpcResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefaultVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefaultVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultVpcResponse) GoString() string {
	return s.String()
}

func (s *CreateDefaultVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDefaultVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDefaultVpcResponse) GetBody() *CreateDefaultVpcResponseBody {
	return s.Body
}

func (s *CreateDefaultVpcResponse) SetHeaders(v map[string]*string) *CreateDefaultVpcResponse {
	s.Headers = v
	return s
}

func (s *CreateDefaultVpcResponse) SetStatusCode(v int32) *CreateDefaultVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefaultVpcResponse) SetBody(v *CreateDefaultVpcResponseBody) *CreateDefaultVpcResponse {
	s.Body = v
	return s
}

func (s *CreateDefaultVpcResponse) Validate() error {
	return dara.Validate(s)
}
