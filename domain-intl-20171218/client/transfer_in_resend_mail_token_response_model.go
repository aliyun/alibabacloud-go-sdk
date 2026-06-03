// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInResendMailTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferInResendMailTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferInResendMailTokenResponse
	GetStatusCode() *int32
	SetBody(v *TransferInResendMailTokenResponseBody) *TransferInResendMailTokenResponse
	GetBody() *TransferInResendMailTokenResponseBody
}

type TransferInResendMailTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferInResendMailTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferInResendMailTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferInResendMailTokenResponse) GoString() string {
	return s.String()
}

func (s *TransferInResendMailTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferInResendMailTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferInResendMailTokenResponse) GetBody() *TransferInResendMailTokenResponseBody {
	return s.Body
}

func (s *TransferInResendMailTokenResponse) SetHeaders(v map[string]*string) *TransferInResendMailTokenResponse {
	s.Headers = v
	return s
}

func (s *TransferInResendMailTokenResponse) SetStatusCode(v int32) *TransferInResendMailTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInResendMailTokenResponse) SetBody(v *TransferInResendMailTokenResponseBody) *TransferInResendMailTokenResponse {
	s.Body = v
	return s
}

func (s *TransferInResendMailTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
