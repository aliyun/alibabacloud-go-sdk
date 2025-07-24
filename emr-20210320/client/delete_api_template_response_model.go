// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApiTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApiTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApiTemplateResponseBody) *DeleteApiTemplateResponse
	GetBody() *DeleteApiTemplateResponseBody
}

type DeleteApiTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApiTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApiTemplateResponse) GetBody() *DeleteApiTemplateResponseBody {
	return s.Body
}

func (s *DeleteApiTemplateResponse) SetHeaders(v map[string]*string) *DeleteApiTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiTemplateResponse) SetStatusCode(v int32) *DeleteApiTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiTemplateResponse) SetBody(v *DeleteApiTemplateResponseBody) *DeleteApiTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteApiTemplateResponse) Validate() error {
	return dara.Validate(s)
}
