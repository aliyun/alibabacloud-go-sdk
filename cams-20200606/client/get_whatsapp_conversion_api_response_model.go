// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappConversionApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWhatsappConversionApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWhatsappConversionApiResponse
	GetStatusCode() *int32
	SetBody(v *GetWhatsappConversionApiResponseBody) *GetWhatsappConversionApiResponse
	GetBody() *GetWhatsappConversionApiResponseBody
}

type GetWhatsappConversionApiResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWhatsappConversionApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWhatsappConversionApiResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappConversionApiResponse) GoString() string {
	return s.String()
}

func (s *GetWhatsappConversionApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWhatsappConversionApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWhatsappConversionApiResponse) GetBody() *GetWhatsappConversionApiResponseBody {
	return s.Body
}

func (s *GetWhatsappConversionApiResponse) SetHeaders(v map[string]*string) *GetWhatsappConversionApiResponse {
	s.Headers = v
	return s
}

func (s *GetWhatsappConversionApiResponse) SetStatusCode(v int32) *GetWhatsappConversionApiResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWhatsappConversionApiResponse) SetBody(v *GetWhatsappConversionApiResponseBody) *GetWhatsappConversionApiResponse {
	s.Body = v
	return s
}

func (s *GetWhatsappConversionApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
