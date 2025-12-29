// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDConversionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UAIDConversionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UAIDConversionResponse
	GetStatusCode() *int32
	SetBody(v *UAIDConversionResponseBody) *UAIDConversionResponse
	GetBody() *UAIDConversionResponseBody
}

type UAIDConversionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UAIDConversionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UAIDConversionResponse) String() string {
	return dara.Prettify(s)
}

func (s UAIDConversionResponse) GoString() string {
	return s.String()
}

func (s *UAIDConversionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UAIDConversionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UAIDConversionResponse) GetBody() *UAIDConversionResponseBody {
	return s.Body
}

func (s *UAIDConversionResponse) SetHeaders(v map[string]*string) *UAIDConversionResponse {
	s.Headers = v
	return s
}

func (s *UAIDConversionResponse) SetStatusCode(v int32) *UAIDConversionResponse {
	s.StatusCode = &v
	return s
}

func (s *UAIDConversionResponse) SetBody(v *UAIDConversionResponseBody) *UAIDConversionResponse {
	s.Body = v
	return s
}

func (s *UAIDConversionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
