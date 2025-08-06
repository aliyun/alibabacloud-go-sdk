// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVodTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVodTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVodTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVodTemplateResponseBody) *UpdateVodTemplateResponse
	GetBody() *UpdateVodTemplateResponseBody
}

type UpdateVodTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVodTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateVodTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVodTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVodTemplateResponse) GetBody() *UpdateVodTemplateResponseBody {
	return s.Body
}

func (s *UpdateVodTemplateResponse) SetHeaders(v map[string]*string) *UpdateVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateVodTemplateResponse) SetStatusCode(v int32) *UpdateVodTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVodTemplateResponse) SetBody(v *UpdateVodTemplateResponseBody) *UpdateVodTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateVodTemplateResponse) Validate() error {
	return dara.Validate(s)
}
