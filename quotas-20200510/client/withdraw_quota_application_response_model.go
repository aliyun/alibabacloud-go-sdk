// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawQuotaApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawQuotaApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawQuotaApplicationResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawQuotaApplicationResponseBody) *WithdrawQuotaApplicationResponse
	GetBody() *WithdrawQuotaApplicationResponseBody
}

type WithdrawQuotaApplicationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawQuotaApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawQuotaApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawQuotaApplicationResponse) GoString() string {
	return s.String()
}

func (s *WithdrawQuotaApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawQuotaApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawQuotaApplicationResponse) GetBody() *WithdrawQuotaApplicationResponseBody {
	return s.Body
}

func (s *WithdrawQuotaApplicationResponse) SetHeaders(v map[string]*string) *WithdrawQuotaApplicationResponse {
	s.Headers = v
	return s
}

func (s *WithdrawQuotaApplicationResponse) SetStatusCode(v int32) *WithdrawQuotaApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawQuotaApplicationResponse) SetBody(v *WithdrawQuotaApplicationResponseBody) *WithdrawQuotaApplicationResponse {
	s.Body = v
	return s
}

func (s *WithdrawQuotaApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
