// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppAgentTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppAgentTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppAgentTemplateResponseBody) *ModifyAppAgentTemplateResponse
	GetBody() *ModifyAppAgentTemplateResponseBody
}

type ModifyAppAgentTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppAgentTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppAgentTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppAgentTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppAgentTemplateResponse) GetBody() *ModifyAppAgentTemplateResponseBody {
	return s.Body
}

func (s *ModifyAppAgentTemplateResponse) SetHeaders(v map[string]*string) *ModifyAppAgentTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppAgentTemplateResponse) SetStatusCode(v int32) *ModifyAppAgentTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppAgentTemplateResponse) SetBody(v *ModifyAppAgentTemplateResponseBody) *ModifyAppAgentTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyAppAgentTemplateResponse) Validate() error {
	return dara.Validate(s)
}
