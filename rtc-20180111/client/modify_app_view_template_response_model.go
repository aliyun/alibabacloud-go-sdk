// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppViewTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppViewTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppViewTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppViewTemplateResponseBody) *ModifyAppViewTemplateResponse
	GetBody() *ModifyAppViewTemplateResponseBody
}

type ModifyAppViewTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppViewTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppViewTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppViewTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppViewTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppViewTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppViewTemplateResponse) GetBody() *ModifyAppViewTemplateResponseBody {
	return s.Body
}

func (s *ModifyAppViewTemplateResponse) SetHeaders(v map[string]*string) *ModifyAppViewTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppViewTemplateResponse) SetStatusCode(v int32) *ModifyAppViewTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppViewTemplateResponse) SetBody(v *ModifyAppViewTemplateResponseBody) *ModifyAppViewTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyAppViewTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
