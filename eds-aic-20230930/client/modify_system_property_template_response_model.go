// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySystemPropertyTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySystemPropertyTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySystemPropertyTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifySystemPropertyTemplateResponseBody) *ModifySystemPropertyTemplateResponse
	GetBody() *ModifySystemPropertyTemplateResponseBody
}

type ModifySystemPropertyTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySystemPropertyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySystemPropertyTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemPropertyTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifySystemPropertyTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySystemPropertyTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySystemPropertyTemplateResponse) GetBody() *ModifySystemPropertyTemplateResponseBody {
	return s.Body
}

func (s *ModifySystemPropertyTemplateResponse) SetHeaders(v map[string]*string) *ModifySystemPropertyTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifySystemPropertyTemplateResponse) SetStatusCode(v int32) *ModifySystemPropertyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySystemPropertyTemplateResponse) SetBody(v *ModifySystemPropertyTemplateResponseBody) *ModifySystemPropertyTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifySystemPropertyTemplateResponse) Validate() error {
	return dara.Validate(s)
}
