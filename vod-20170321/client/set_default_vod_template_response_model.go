// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultVodTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultVodTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultVodTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultVodTemplateResponseBody) *SetDefaultVodTemplateResponse
	GetBody() *SetDefaultVodTemplateResponseBody
}

type SetDefaultVodTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultVodTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultVodTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultVodTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultVodTemplateResponse) GetBody() *SetDefaultVodTemplateResponseBody {
	return s.Body
}

func (s *SetDefaultVodTemplateResponse) SetHeaders(v map[string]*string) *SetDefaultVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultVodTemplateResponse) SetStatusCode(v int32) *SetDefaultVodTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultVodTemplateResponse) SetBody(v *SetDefaultVodTemplateResponseBody) *SetDefaultVodTemplateResponse {
	s.Body = v
	return s
}

func (s *SetDefaultVodTemplateResponse) Validate() error {
	return dara.Validate(s)
}
