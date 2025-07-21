// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckReplyToMailAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckReplyToMailAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckReplyToMailAddressResponse
	GetStatusCode() *int32
	SetBody(v *CheckReplyToMailAddressResponseBody) *CheckReplyToMailAddressResponse
	GetBody() *CheckReplyToMailAddressResponseBody
}

type CheckReplyToMailAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckReplyToMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckReplyToMailAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckReplyToMailAddressResponse) GoString() string {
	return s.String()
}

func (s *CheckReplyToMailAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckReplyToMailAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckReplyToMailAddressResponse) GetBody() *CheckReplyToMailAddressResponseBody {
	return s.Body
}

func (s *CheckReplyToMailAddressResponse) SetHeaders(v map[string]*string) *CheckReplyToMailAddressResponse {
	s.Headers = v
	return s
}

func (s *CheckReplyToMailAddressResponse) SetStatusCode(v int32) *CheckReplyToMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckReplyToMailAddressResponse) SetBody(v *CheckReplyToMailAddressResponseBody) *CheckReplyToMailAddressResponse {
	s.Body = v
	return s
}

func (s *CheckReplyToMailAddressResponse) Validate() error {
	return dara.Validate(s)
}
