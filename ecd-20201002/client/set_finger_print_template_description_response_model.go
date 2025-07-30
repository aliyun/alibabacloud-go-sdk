// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFingerPrintTemplateDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetFingerPrintTemplateDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetFingerPrintTemplateDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *SetFingerPrintTemplateDescriptionResponseBody) *SetFingerPrintTemplateDescriptionResponse
	GetBody() *SetFingerPrintTemplateDescriptionResponseBody
}

type SetFingerPrintTemplateDescriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFingerPrintTemplateDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionResponse) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetFingerPrintTemplateDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetFingerPrintTemplateDescriptionResponse) GetBody() *SetFingerPrintTemplateDescriptionResponseBody {
	return s.Body
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetHeaders(v map[string]*string) *SetFingerPrintTemplateDescriptionResponse {
	s.Headers = v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetStatusCode(v int32) *SetFingerPrintTemplateDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetBody(v *SetFingerPrintTemplateDescriptionResponseBody) *SetFingerPrintTemplateDescriptionResponse {
	s.Body = v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponse) Validate() error {
	return dara.Validate(s)
}
