// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhatAppTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWhatAppTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWhatAppTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListWhatAppTemplateResponseBody) *ListWhatAppTemplateResponse
	GetBody() *ListWhatAppTemplateResponseBody
}

type ListWhatAppTemplateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWhatAppTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWhatAppTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWhatAppTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListWhatAppTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWhatAppTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWhatAppTemplateResponse) GetBody() *ListWhatAppTemplateResponseBody {
	return s.Body
}

func (s *ListWhatAppTemplateResponse) SetHeaders(v map[string]*string) *ListWhatAppTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListWhatAppTemplateResponse) SetStatusCode(v int32) *ListWhatAppTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWhatAppTemplateResponse) SetBody(v *ListWhatAppTemplateResponseBody) *ListWhatAppTemplateResponse {
	s.Body = v
	return s
}

func (s *ListWhatAppTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
