// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhatsappConversionApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWhatsappConversionApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWhatsappConversionApiResponse
	GetStatusCode() *int32
	SetBody(v *CreateWhatsappConversionApiResponseBody) *CreateWhatsappConversionApiResponse
	GetBody() *CreateWhatsappConversionApiResponseBody
}

type CreateWhatsappConversionApiResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWhatsappConversionApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWhatsappConversionApiResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWhatsappConversionApiResponse) GoString() string {
	return s.String()
}

func (s *CreateWhatsappConversionApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWhatsappConversionApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWhatsappConversionApiResponse) GetBody() *CreateWhatsappConversionApiResponseBody {
	return s.Body
}

func (s *CreateWhatsappConversionApiResponse) SetHeaders(v map[string]*string) *CreateWhatsappConversionApiResponse {
	s.Headers = v
	return s
}

func (s *CreateWhatsappConversionApiResponse) SetStatusCode(v int32) *CreateWhatsappConversionApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWhatsappConversionApiResponse) SetBody(v *CreateWhatsappConversionApiResponseBody) *CreateWhatsappConversionApiResponse {
	s.Body = v
	return s
}

func (s *CreateWhatsappConversionApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
