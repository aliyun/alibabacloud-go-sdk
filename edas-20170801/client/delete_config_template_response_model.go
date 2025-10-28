// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConfigTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConfigTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConfigTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConfigTemplateResponseBody) *DeleteConfigTemplateResponse
	GetBody() *DeleteConfigTemplateResponseBody
}

type DeleteConfigTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConfigTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConfigTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConfigTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConfigTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConfigTemplateResponse) GetBody() *DeleteConfigTemplateResponseBody {
	return s.Body
}

func (s *DeleteConfigTemplateResponse) SetHeaders(v map[string]*string) *DeleteConfigTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigTemplateResponse) SetStatusCode(v int32) *DeleteConfigTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigTemplateResponse) SetBody(v *DeleteConfigTemplateResponseBody) *DeleteConfigTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteConfigTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
