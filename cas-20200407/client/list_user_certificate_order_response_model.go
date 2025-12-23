// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserCertificateOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserCertificateOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserCertificateOrderResponse
	GetStatusCode() *int32
	SetBody(v *ListUserCertificateOrderResponseBody) *ListUserCertificateOrderResponse
	GetBody() *ListUserCertificateOrderResponseBody
}

type ListUserCertificateOrderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserCertificateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserCertificateOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserCertificateOrderResponse) GoString() string {
	return s.String()
}

func (s *ListUserCertificateOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserCertificateOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserCertificateOrderResponse) GetBody() *ListUserCertificateOrderResponseBody {
	return s.Body
}

func (s *ListUserCertificateOrderResponse) SetHeaders(v map[string]*string) *ListUserCertificateOrderResponse {
	s.Headers = v
	return s
}

func (s *ListUserCertificateOrderResponse) SetStatusCode(v int32) *ListUserCertificateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserCertificateOrderResponse) SetBody(v *ListUserCertificateOrderResponseBody) *ListUserCertificateOrderResponse {
	s.Body = v
	return s
}

func (s *ListUserCertificateOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
