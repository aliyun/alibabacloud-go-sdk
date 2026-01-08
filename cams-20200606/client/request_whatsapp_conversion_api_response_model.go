// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestWhatsappConversionApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RequestWhatsappConversionApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RequestWhatsappConversionApiResponse
	GetStatusCode() *int32
	SetBody(v *RequestWhatsappConversionApiResponseBody) *RequestWhatsappConversionApiResponse
	GetBody() *RequestWhatsappConversionApiResponseBody
}

type RequestWhatsappConversionApiResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestWhatsappConversionApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestWhatsappConversionApiResponse) String() string {
	return dara.Prettify(s)
}

func (s RequestWhatsappConversionApiResponse) GoString() string {
	return s.String()
}

func (s *RequestWhatsappConversionApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RequestWhatsappConversionApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RequestWhatsappConversionApiResponse) GetBody() *RequestWhatsappConversionApiResponseBody {
	return s.Body
}

func (s *RequestWhatsappConversionApiResponse) SetHeaders(v map[string]*string) *RequestWhatsappConversionApiResponse {
	s.Headers = v
	return s
}

func (s *RequestWhatsappConversionApiResponse) SetStatusCode(v int32) *RequestWhatsappConversionApiResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestWhatsappConversionApiResponse) SetBody(v *RequestWhatsappConversionApiResponseBody) *RequestWhatsappConversionApiResponse {
	s.Body = v
	return s
}

func (s *RequestWhatsappConversionApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
