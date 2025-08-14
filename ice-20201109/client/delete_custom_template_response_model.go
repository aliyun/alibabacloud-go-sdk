// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomTemplateResponseBody) *DeleteCustomTemplateResponse
	GetBody() *DeleteCustomTemplateResponseBody
}

type DeleteCustomTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomTemplateResponse) GetBody() *DeleteCustomTemplateResponseBody {
	return s.Body
}

func (s *DeleteCustomTemplateResponse) SetHeaders(v map[string]*string) *DeleteCustomTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTemplateResponse) SetStatusCode(v int32) *DeleteCustomTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTemplateResponse) SetBody(v *DeleteCustomTemplateResponseBody) *DeleteCustomTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomTemplateResponse) Validate() error {
	return dara.Validate(s)
}
