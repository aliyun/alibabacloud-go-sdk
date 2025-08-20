// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSparkTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSparkTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSparkTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSparkTemplateResponseBody) *DeleteSparkTemplateResponse
	GetBody() *DeleteSparkTemplateResponseBody
}

type DeleteSparkTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSparkTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSparkTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSparkTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSparkTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSparkTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSparkTemplateResponse) GetBody() *DeleteSparkTemplateResponseBody {
	return s.Body
}

func (s *DeleteSparkTemplateResponse) SetHeaders(v map[string]*string) *DeleteSparkTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSparkTemplateResponse) SetStatusCode(v int32) *DeleteSparkTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSparkTemplateResponse) SetBody(v *DeleteSparkTemplateResponseBody) *DeleteSparkTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteSparkTemplateResponse) Validate() error {
	return dara.Validate(s)
}
