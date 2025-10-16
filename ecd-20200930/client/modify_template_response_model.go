// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTemplateResponseBody) *ModifyTemplateResponse
	GetBody() *ModifyTemplateResponseBody
}

type ModifyTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTemplateResponse) GetBody() *ModifyTemplateResponseBody {
	return s.Body
}

func (s *ModifyTemplateResponse) SetHeaders(v map[string]*string) *ModifyTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyTemplateResponse) SetStatusCode(v int32) *ModifyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTemplateResponse) SetBody(v *ModifyTemplateResponseBody) *ModifyTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
