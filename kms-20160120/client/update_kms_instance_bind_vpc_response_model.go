// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKmsInstanceBindVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKmsInstanceBindVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKmsInstanceBindVpcResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKmsInstanceBindVpcResponseBody) *UpdateKmsInstanceBindVpcResponse
	GetBody() *UpdateKmsInstanceBindVpcResponseBody
}

type UpdateKmsInstanceBindVpcResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKmsInstanceBindVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKmsInstanceBindVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKmsInstanceBindVpcResponse) GoString() string {
	return s.String()
}

func (s *UpdateKmsInstanceBindVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKmsInstanceBindVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKmsInstanceBindVpcResponse) GetBody() *UpdateKmsInstanceBindVpcResponseBody {
	return s.Body
}

func (s *UpdateKmsInstanceBindVpcResponse) SetHeaders(v map[string]*string) *UpdateKmsInstanceBindVpcResponse {
	s.Headers = v
	return s
}

func (s *UpdateKmsInstanceBindVpcResponse) SetStatusCode(v int32) *UpdateKmsInstanceBindVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKmsInstanceBindVpcResponse) SetBody(v *UpdateKmsInstanceBindVpcResponseBody) *UpdateKmsInstanceBindVpcResponse {
	s.Body = v
	return s
}

func (s *UpdateKmsInstanceBindVpcResponse) Validate() error {
	return dara.Validate(s)
}
