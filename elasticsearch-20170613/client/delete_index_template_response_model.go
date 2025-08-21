// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIndexTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIndexTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIndexTemplateResponseBody) *DeleteIndexTemplateResponse
	GetBody() *DeleteIndexTemplateResponseBody
}

type DeleteIndexTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIndexTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIndexTemplateResponse) GetBody() *DeleteIndexTemplateResponseBody {
	return s.Body
}

func (s *DeleteIndexTemplateResponse) SetHeaders(v map[string]*string) *DeleteIndexTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexTemplateResponse) SetStatusCode(v int32) *DeleteIndexTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexTemplateResponse) SetBody(v *DeleteIndexTemplateResponseBody) *DeleteIndexTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteIndexTemplateResponse) Validate() error {
	return dara.Validate(s)
}
