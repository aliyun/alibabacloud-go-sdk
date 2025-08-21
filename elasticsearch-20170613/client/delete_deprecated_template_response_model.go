// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeprecatedTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeprecatedTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeprecatedTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeprecatedTemplateResponseBody) *DeleteDeprecatedTemplateResponse
	GetBody() *DeleteDeprecatedTemplateResponseBody
}

type DeleteDeprecatedTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeprecatedTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeprecatedTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeprecatedTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeprecatedTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeprecatedTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeprecatedTemplateResponse) GetBody() *DeleteDeprecatedTemplateResponseBody {
	return s.Body
}

func (s *DeleteDeprecatedTemplateResponse) SetHeaders(v map[string]*string) *DeleteDeprecatedTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeprecatedTemplateResponse) SetStatusCode(v int32) *DeleteDeprecatedTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeprecatedTemplateResponse) SetBody(v *DeleteDeprecatedTemplateResponseBody) *DeleteDeprecatedTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteDeprecatedTemplateResponse) Validate() error {
	return dara.Validate(s)
}
