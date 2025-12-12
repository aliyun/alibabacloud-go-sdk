// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWhitelistTemplateResponseBody) *DeleteWhitelistTemplateResponse
	GetBody() *DeleteWhitelistTemplateResponseBody
}

type DeleteWhitelistTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWhitelistTemplateResponse) GetBody() *DeleteWhitelistTemplateResponseBody {
	return s.Body
}

func (s *DeleteWhitelistTemplateResponse) SetHeaders(v map[string]*string) *DeleteWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhitelistTemplateResponse) SetStatusCode(v int32) *DeleteWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWhitelistTemplateResponse) SetBody(v *DeleteWhitelistTemplateResponseBody) *DeleteWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
