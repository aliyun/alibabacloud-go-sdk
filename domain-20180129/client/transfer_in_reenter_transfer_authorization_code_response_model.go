// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInReenterTransferAuthorizationCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferInReenterTransferAuthorizationCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferInReenterTransferAuthorizationCodeResponse
	GetStatusCode() *int32
	SetBody(v *TransferInReenterTransferAuthorizationCodeResponseBody) *TransferInReenterTransferAuthorizationCodeResponse
	GetBody() *TransferInReenterTransferAuthorizationCodeResponseBody
}

type TransferInReenterTransferAuthorizationCodeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransferInReenterTransferAuthorizationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferInReenterTransferAuthorizationCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferInReenterTransferAuthorizationCodeResponse) GoString() string {
	return s.String()
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) GetBody() *TransferInReenterTransferAuthorizationCodeResponseBody {
	return s.Body
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) SetHeaders(v map[string]*string) *TransferInReenterTransferAuthorizationCodeResponse {
	s.Headers = v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) SetStatusCode(v int32) *TransferInReenterTransferAuthorizationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) SetBody(v *TransferInReenterTransferAuthorizationCodeResponseBody) *TransferInReenterTransferAuthorizationCodeResponse {
	s.Body = v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
