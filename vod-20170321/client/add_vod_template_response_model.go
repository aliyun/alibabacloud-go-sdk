// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVodTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddVodTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddVodTemplateResponse
	GetStatusCode() *int32
	SetBody(v *AddVodTemplateResponseBody) *AddVodTemplateResponse
	GetBody() *AddVodTemplateResponseBody
}

type AddVodTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddVodTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddVodTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddVodTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddVodTemplateResponse) GetBody() *AddVodTemplateResponseBody {
	return s.Body
}

func (s *AddVodTemplateResponse) SetHeaders(v map[string]*string) *AddVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddVodTemplateResponse) SetStatusCode(v int32) *AddVodTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVodTemplateResponse) SetBody(v *AddVodTemplateResponseBody) *AddVodTemplateResponse {
	s.Body = v
	return s
}

func (s *AddVodTemplateResponse) Validate() error {
	return dara.Validate(s)
}
