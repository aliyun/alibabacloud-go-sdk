// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappHealthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWhatsappHealthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWhatsappHealthStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetWhatsappHealthStatusResponseBody) *GetWhatsappHealthStatusResponse
	GetBody() *GetWhatsappHealthStatusResponseBody
}

type GetWhatsappHealthStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWhatsappHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWhatsappHealthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *GetWhatsappHealthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWhatsappHealthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWhatsappHealthStatusResponse) GetBody() *GetWhatsappHealthStatusResponseBody {
	return s.Body
}

func (s *GetWhatsappHealthStatusResponse) SetHeaders(v map[string]*string) *GetWhatsappHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *GetWhatsappHealthStatusResponse) SetStatusCode(v int32) *GetWhatsappHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWhatsappHealthStatusResponse) SetBody(v *GetWhatsappHealthStatusResponseBody) *GetWhatsappHealthStatusResponse {
	s.Body = v
	return s
}

func (s *GetWhatsappHealthStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
