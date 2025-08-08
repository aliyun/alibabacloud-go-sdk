// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenderServicesByTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenderServicesByTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenderServicesByTemplateResponse
	GetStatusCode() *int32
	SetBody(v *RenderServicesByTemplateResponseBody) *RenderServicesByTemplateResponse
	GetBody() *RenderServicesByTemplateResponseBody
}

type RenderServicesByTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenderServicesByTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderServicesByTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s RenderServicesByTemplateResponse) GoString() string {
	return s.String()
}

func (s *RenderServicesByTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenderServicesByTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenderServicesByTemplateResponse) GetBody() *RenderServicesByTemplateResponseBody {
	return s.Body
}

func (s *RenderServicesByTemplateResponse) SetHeaders(v map[string]*string) *RenderServicesByTemplateResponse {
	s.Headers = v
	return s
}

func (s *RenderServicesByTemplateResponse) SetStatusCode(v int32) *RenderServicesByTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *RenderServicesByTemplateResponse) SetBody(v *RenderServicesByTemplateResponseBody) *RenderServicesByTemplateResponse {
	s.Body = v
	return s
}

func (s *RenderServicesByTemplateResponse) Validate() error {
	return dara.Validate(s)
}
