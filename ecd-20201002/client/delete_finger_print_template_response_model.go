// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFingerPrintTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFingerPrintTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFingerPrintTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFingerPrintTemplateResponseBody) *DeleteFingerPrintTemplateResponse
	GetBody() *DeleteFingerPrintTemplateResponseBody
}

type DeleteFingerPrintTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFingerPrintTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFingerPrintTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFingerPrintTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFingerPrintTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFingerPrintTemplateResponse) GetBody() *DeleteFingerPrintTemplateResponseBody {
	return s.Body
}

func (s *DeleteFingerPrintTemplateResponse) SetHeaders(v map[string]*string) *DeleteFingerPrintTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFingerPrintTemplateResponse) SetStatusCode(v int32) *DeleteFingerPrintTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFingerPrintTemplateResponse) SetBody(v *DeleteFingerPrintTemplateResponseBody) *DeleteFingerPrintTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteFingerPrintTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
