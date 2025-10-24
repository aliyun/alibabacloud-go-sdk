// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseTemplateStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseTemplateStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseTemplateStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseTemplateStatusResponseBody) *ModifyDefenseTemplateStatusResponse
	GetBody() *ModifyDefenseTemplateStatusResponseBody
}

type ModifyDefenseTemplateStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseTemplateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseTemplateStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseTemplateStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseTemplateStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseTemplateStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseTemplateStatusResponse) GetBody() *ModifyDefenseTemplateStatusResponseBody {
	return s.Body
}

func (s *ModifyDefenseTemplateStatusResponse) SetHeaders(v map[string]*string) *ModifyDefenseTemplateStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseTemplateStatusResponse) SetStatusCode(v int32) *ModifyDefenseTemplateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseTemplateStatusResponse) SetBody(v *ModifyDefenseTemplateStatusResponseBody) *ModifyDefenseTemplateStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseTemplateStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
