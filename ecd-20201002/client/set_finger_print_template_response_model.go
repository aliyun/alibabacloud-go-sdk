// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFingerPrintTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetFingerPrintTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetFingerPrintTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SetFingerPrintTemplateResponseBody) *SetFingerPrintTemplateResponse
	GetBody() *SetFingerPrintTemplateResponseBody
}

type SetFingerPrintTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFingerPrintTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFingerPrintTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetFingerPrintTemplateResponse) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetFingerPrintTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetFingerPrintTemplateResponse) GetBody() *SetFingerPrintTemplateResponseBody {
	return s.Body
}

func (s *SetFingerPrintTemplateResponse) SetHeaders(v map[string]*string) *SetFingerPrintTemplateResponse {
	s.Headers = v
	return s
}

func (s *SetFingerPrintTemplateResponse) SetStatusCode(v int32) *SetFingerPrintTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFingerPrintTemplateResponse) SetBody(v *SetFingerPrintTemplateResponseBody) *SetFingerPrintTemplateResponse {
	s.Body = v
	return s
}

func (s *SetFingerPrintTemplateResponse) Validate() error {
	return dara.Validate(s)
}
