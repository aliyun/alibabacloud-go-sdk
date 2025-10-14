// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransitVisaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransitVisaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransitVisaResponse
	GetStatusCode() *int32
	SetBody(v *TransitVisaResponseBody) *TransitVisaResponse
	GetBody() *TransitVisaResponseBody
}

type TransitVisaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TransitVisaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransitVisaResponse) String() string {
	return dara.Prettify(s)
}

func (s TransitVisaResponse) GoString() string {
	return s.String()
}

func (s *TransitVisaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransitVisaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransitVisaResponse) GetBody() *TransitVisaResponseBody {
	return s.Body
}

func (s *TransitVisaResponse) SetHeaders(v map[string]*string) *TransitVisaResponse {
	s.Headers = v
	return s
}

func (s *TransitVisaResponse) SetStatusCode(v int32) *TransitVisaResponse {
	s.StatusCode = &v
	return s
}

func (s *TransitVisaResponse) SetBody(v *TransitVisaResponseBody) *TransitVisaResponse {
	s.Body = v
	return s
}

func (s *TransitVisaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
