// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse
	GetBody() *DeleteTemplateResponseBody
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTemplateResponse) GetBody() *DeleteTemplateResponseBody {
	return s.Body
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteTemplateResponse) Validate() error {
	return dara.Validate(s)
}
