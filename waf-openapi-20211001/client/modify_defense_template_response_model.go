// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseTemplateResponseBody) *ModifyDefenseTemplateResponse
	GetBody() *ModifyDefenseTemplateResponseBody
}

type ModifyDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseTemplateResponse) GetBody() *ModifyDefenseTemplateResponseBody {
	return s.Body
}

func (s *ModifyDefenseTemplateResponse) SetHeaders(v map[string]*string) *ModifyDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseTemplateResponse) SetStatusCode(v int32) *ModifyDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseTemplateResponse) SetBody(v *ModifyDefenseTemplateResponseBody) *ModifyDefenseTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
