// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateDNSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivateDNSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivateDNSResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivateDNSResponseBody) *UpdatePrivateDNSResponse
	GetBody() *UpdatePrivateDNSResponseBody
}

type UpdatePrivateDNSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateDNSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateDNSResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateDNSResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateDNSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivateDNSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivateDNSResponse) GetBody() *UpdatePrivateDNSResponseBody {
	return s.Body
}

func (s *UpdatePrivateDNSResponse) SetHeaders(v map[string]*string) *UpdatePrivateDNSResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateDNSResponse) SetStatusCode(v int32) *UpdatePrivateDNSResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateDNSResponse) SetBody(v *UpdatePrivateDNSResponseBody) *UpdatePrivateDNSResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivateDNSResponse) Validate() error {
	return dara.Validate(s)
}
