// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPathFieldTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiPathFieldTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiPathFieldTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetApiPathFieldTemplateResponseBody) *GetApiPathFieldTemplateResponse
	GetBody() *GetApiPathFieldTemplateResponseBody
}

type GetApiPathFieldTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiPathFieldTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiPathFieldTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiPathFieldTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiPathFieldTemplateResponse) GetBody() *GetApiPathFieldTemplateResponseBody {
	return s.Body
}

func (s *GetApiPathFieldTemplateResponse) SetHeaders(v map[string]*string) *GetApiPathFieldTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetApiPathFieldTemplateResponse) SetStatusCode(v int32) *GetApiPathFieldTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiPathFieldTemplateResponse) SetBody(v *GetApiPathFieldTemplateResponseBody) *GetApiPathFieldTemplateResponse {
	s.Body = v
	return s
}

func (s *GetApiPathFieldTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
