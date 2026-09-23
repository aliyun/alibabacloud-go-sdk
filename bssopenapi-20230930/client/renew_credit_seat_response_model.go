// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewCreditSeatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewCreditSeatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewCreditSeatResponse
	GetStatusCode() *int32
	SetBody(v *RenewCreditSeatResponseBody) *RenewCreditSeatResponse
	GetBody() *RenewCreditSeatResponseBody
}

type RenewCreditSeatResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewCreditSeatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewCreditSeatResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewCreditSeatResponse) GoString() string {
	return s.String()
}

func (s *RenewCreditSeatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewCreditSeatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewCreditSeatResponse) GetBody() *RenewCreditSeatResponseBody {
	return s.Body
}

func (s *RenewCreditSeatResponse) SetHeaders(v map[string]*string) *RenewCreditSeatResponse {
	s.Headers = v
	return s
}

func (s *RenewCreditSeatResponse) SetStatusCode(v int32) *RenewCreditSeatResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewCreditSeatResponse) SetBody(v *RenewCreditSeatResponseBody) *RenewCreditSeatResponse {
	s.Body = v
	return s
}

func (s *RenewCreditSeatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
