// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpconnFromVbrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpconnFromVbrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpconnFromVbrResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpconnFromVbrResponseBody) *CreateVpconnFromVbrResponse
	GetBody() *CreateVpconnFromVbrResponseBody
}

type CreateVpconnFromVbrResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpconnFromVbrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpconnFromVbrResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpconnFromVbrResponse) GoString() string {
	return s.String()
}

func (s *CreateVpconnFromVbrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpconnFromVbrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpconnFromVbrResponse) GetBody() *CreateVpconnFromVbrResponseBody {
	return s.Body
}

func (s *CreateVpconnFromVbrResponse) SetHeaders(v map[string]*string) *CreateVpconnFromVbrResponse {
	s.Headers = v
	return s
}

func (s *CreateVpconnFromVbrResponse) SetStatusCode(v int32) *CreateVpconnFromVbrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpconnFromVbrResponse) SetBody(v *CreateVpconnFromVbrResponseBody) *CreateVpconnFromVbrResponse {
	s.Body = v
	return s
}

func (s *CreateVpconnFromVbrResponse) Validate() error {
	return dara.Validate(s)
}
