// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertTemplateResponseBody) *DeleteAlertTemplateResponse
	GetBody() *DeleteAlertTemplateResponseBody
}

type DeleteAlertTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertTemplateResponse) GetBody() *DeleteAlertTemplateResponseBody {
	return s.Body
}

func (s *DeleteAlertTemplateResponse) SetHeaders(v map[string]*string) *DeleteAlertTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertTemplateResponse) SetStatusCode(v int32) *DeleteAlertTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertTemplateResponse) SetBody(v *DeleteAlertTemplateResponseBody) *DeleteAlertTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
