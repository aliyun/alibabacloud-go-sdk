// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataQualityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataQualityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataQualityTemplateResponseBody) *UpdateDataQualityTemplateResponse
	GetBody() *UpdateDataQualityTemplateResponseBody
}

type UpdateDataQualityTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataQualityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataQualityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataQualityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataQualityTemplateResponse) GetBody() *UpdateDataQualityTemplateResponseBody {
	return s.Body
}

func (s *UpdateDataQualityTemplateResponse) SetHeaders(v map[string]*string) *UpdateDataQualityTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataQualityTemplateResponse) SetStatusCode(v int32) *UpdateDataQualityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataQualityTemplateResponse) SetBody(v *UpdateDataQualityTemplateResponseBody) *UpdateDataQualityTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateDataQualityTemplateResponse) Validate() error {
	return dara.Validate(s)
}
