// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckIntlFixPriceDomainStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckIntlFixPriceDomainStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckIntlFixPriceDomainStatusResponse
	GetStatusCode() *int32
	SetBody(v *CheckIntlFixPriceDomainStatusResponseBody) *CheckIntlFixPriceDomainStatusResponse
	GetBody() *CheckIntlFixPriceDomainStatusResponseBody
}

type CheckIntlFixPriceDomainStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckIntlFixPriceDomainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckIntlFixPriceDomainStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckIntlFixPriceDomainStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckIntlFixPriceDomainStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckIntlFixPriceDomainStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckIntlFixPriceDomainStatusResponse) GetBody() *CheckIntlFixPriceDomainStatusResponseBody {
	return s.Body
}

func (s *CheckIntlFixPriceDomainStatusResponse) SetHeaders(v map[string]*string) *CheckIntlFixPriceDomainStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponse) SetStatusCode(v int32) *CheckIntlFixPriceDomainStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponse) SetBody(v *CheckIntlFixPriceDomainStatusResponseBody) *CheckIntlFixPriceDomainStatusResponse {
	s.Body = v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
