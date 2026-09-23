// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCreditSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCreditSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCreditSeatsResponse
	GetStatusCode() *int32
	SetBody(v *AddCreditSeatsResponseBody) *AddCreditSeatsResponse
	GetBody() *AddCreditSeatsResponseBody
}

type AddCreditSeatsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCreditSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCreditSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCreditSeatsResponse) GoString() string {
	return s.String()
}

func (s *AddCreditSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCreditSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCreditSeatsResponse) GetBody() *AddCreditSeatsResponseBody {
	return s.Body
}

func (s *AddCreditSeatsResponse) SetHeaders(v map[string]*string) *AddCreditSeatsResponse {
	s.Headers = v
	return s
}

func (s *AddCreditSeatsResponse) SetStatusCode(v int32) *AddCreditSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCreditSeatsResponse) SetBody(v *AddCreditSeatsResponseBody) *AddCreditSeatsResponse {
	s.Body = v
	return s
}

func (s *AddCreditSeatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
