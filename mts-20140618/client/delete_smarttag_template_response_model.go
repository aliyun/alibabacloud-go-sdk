// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmarttagTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmarttagTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmarttagTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmarttagTemplateResponseBody) *DeleteSmarttagTemplateResponse
	GetBody() *DeleteSmarttagTemplateResponseBody
}

type DeleteSmarttagTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmarttagTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmarttagTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmarttagTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmarttagTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmarttagTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmarttagTemplateResponse) GetBody() *DeleteSmarttagTemplateResponseBody {
	return s.Body
}

func (s *DeleteSmarttagTemplateResponse) SetHeaders(v map[string]*string) *DeleteSmarttagTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmarttagTemplateResponse) SetStatusCode(v int32) *DeleteSmarttagTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmarttagTemplateResponse) SetBody(v *DeleteSmarttagTemplateResponseBody) *DeleteSmarttagTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteSmarttagTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
