// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleForProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckServiceLinkedRoleForProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckServiceLinkedRoleForProductResponse
	GetStatusCode() *int32
	SetBody(v *CheckServiceLinkedRoleForProductResponseBody) *CheckServiceLinkedRoleForProductResponse
	GetBody() *CheckServiceLinkedRoleForProductResponseBody
}

type CheckServiceLinkedRoleForProductResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceLinkedRoleForProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceLinkedRoleForProductResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForProductResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckServiceLinkedRoleForProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckServiceLinkedRoleForProductResponse) GetBody() *CheckServiceLinkedRoleForProductResponseBody {
	return s.Body
}

func (s *CheckServiceLinkedRoleForProductResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleForProductResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleForProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponse) SetBody(v *CheckServiceLinkedRoleForProductResponseBody) *CheckServiceLinkedRoleForProductResponse {
	s.Body = v
	return s
}

func (s *CheckServiceLinkedRoleForProductResponse) Validate() error {
	return dara.Validate(s)
}
