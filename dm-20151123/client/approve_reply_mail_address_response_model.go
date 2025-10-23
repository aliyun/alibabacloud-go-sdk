// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveReplyMailAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveReplyMailAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveReplyMailAddressResponse
	GetStatusCode() *int32
	SetBody(v *ApproveReplyMailAddressResponseBody) *ApproveReplyMailAddressResponse
	GetBody() *ApproveReplyMailAddressResponseBody
}

type ApproveReplyMailAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveReplyMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveReplyMailAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveReplyMailAddressResponse) GoString() string {
	return s.String()
}

func (s *ApproveReplyMailAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveReplyMailAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveReplyMailAddressResponse) GetBody() *ApproveReplyMailAddressResponseBody {
	return s.Body
}

func (s *ApproveReplyMailAddressResponse) SetHeaders(v map[string]*string) *ApproveReplyMailAddressResponse {
	s.Headers = v
	return s
}

func (s *ApproveReplyMailAddressResponse) SetStatusCode(v int32) *ApproveReplyMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveReplyMailAddressResponse) SetBody(v *ApproveReplyMailAddressResponseBody) *ApproveReplyMailAddressResponse {
	s.Body = v
	return s
}

func (s *ApproveReplyMailAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
