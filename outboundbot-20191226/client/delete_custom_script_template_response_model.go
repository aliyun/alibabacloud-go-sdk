// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomScriptTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomScriptTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomScriptTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomScriptTemplateResponseBody) *DeleteCustomScriptTemplateResponse
	GetBody() *DeleteCustomScriptTemplateResponseBody
}

type DeleteCustomScriptTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomScriptTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomScriptTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomScriptTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomScriptTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomScriptTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomScriptTemplateResponse) GetBody() *DeleteCustomScriptTemplateResponseBody {
	return s.Body
}

func (s *DeleteCustomScriptTemplateResponse) SetHeaders(v map[string]*string) *DeleteCustomScriptTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomScriptTemplateResponse) SetStatusCode(v int32) *DeleteCustomScriptTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomScriptTemplateResponse) SetBody(v *DeleteCustomScriptTemplateResponseBody) *DeleteCustomScriptTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomScriptTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
