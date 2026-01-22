// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFundAccountTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFundAccountTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFundAccountTransferResponse
	GetStatusCode() *int32
	SetBody(v *CreateFundAccountTransferResponseBody) *CreateFundAccountTransferResponse
	GetBody() *CreateFundAccountTransferResponseBody
}

type CreateFundAccountTransferResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFundAccountTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFundAccountTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountTransferResponse) GoString() string {
	return s.String()
}

func (s *CreateFundAccountTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFundAccountTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFundAccountTransferResponse) GetBody() *CreateFundAccountTransferResponseBody {
	return s.Body
}

func (s *CreateFundAccountTransferResponse) SetHeaders(v map[string]*string) *CreateFundAccountTransferResponse {
	s.Headers = v
	return s
}

func (s *CreateFundAccountTransferResponse) SetStatusCode(v int32) *CreateFundAccountTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFundAccountTransferResponse) SetBody(v *CreateFundAccountTransferResponseBody) *CreateFundAccountTransferResponse {
	s.Body = v
	return s
}

func (s *CreateFundAccountTransferResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
