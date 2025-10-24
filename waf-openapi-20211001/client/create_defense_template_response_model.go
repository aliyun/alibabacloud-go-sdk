// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDefenseTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDefenseTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateDefenseTemplateResponseBody) *CreateDefenseTemplateResponse
	GetBody() *CreateDefenseTemplateResponseBody
}

type CreateDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefenseTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDefenseTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDefenseTemplateResponse) GetBody() *CreateDefenseTemplateResponseBody {
	return s.Body
}

func (s *CreateDefenseTemplateResponse) SetHeaders(v map[string]*string) *CreateDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseTemplateResponse) SetStatusCode(v int32) *CreateDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseTemplateResponse) SetBody(v *CreateDefenseTemplateResponseBody) *CreateDefenseTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateDefenseTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
