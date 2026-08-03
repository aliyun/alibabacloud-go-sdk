// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidatePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidatePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidatePriceResponse
	GetStatusCode() *int32
	SetBody(v *ValidatePriceResponseBody) *ValidatePriceResponse
	GetBody() *ValidatePriceResponseBody
}

type ValidatePriceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidatePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidatePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidatePriceResponse) GoString() string {
	return s.String()
}

func (s *ValidatePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidatePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidatePriceResponse) GetBody() *ValidatePriceResponseBody {
	return s.Body
}

func (s *ValidatePriceResponse) SetHeaders(v map[string]*string) *ValidatePriceResponse {
	s.Headers = v
	return s
}

func (s *ValidatePriceResponse) SetStatusCode(v int32) *ValidatePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidatePriceResponse) SetBody(v *ValidatePriceResponseBody) *ValidatePriceResponse {
	s.Body = v
	return s
}

func (s *ValidatePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
