// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConversionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConversionDataResponse
	GetStatusCode() *int32
	SetBody(v *ConversionDataResponseBody) *ConversionDataResponse
	GetBody() *ConversionDataResponseBody
}

type ConversionDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConversionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConversionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s ConversionDataResponse) GoString() string {
	return s.String()
}

func (s *ConversionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConversionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConversionDataResponse) GetBody() *ConversionDataResponseBody {
	return s.Body
}

func (s *ConversionDataResponse) SetHeaders(v map[string]*string) *ConversionDataResponse {
	s.Headers = v
	return s
}

func (s *ConversionDataResponse) SetStatusCode(v int32) *ConversionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ConversionDataResponse) SetBody(v *ConversionDataResponseBody) *ConversionDataResponse {
	s.Body = v
	return s
}

func (s *ConversionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
