// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafTemplateRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWafTemplateRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWafTemplateRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListWafTemplateRulesResponseBody) *ListWafTemplateRulesResponse
	GetBody() *ListWafTemplateRulesResponseBody
}

type ListWafTemplateRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafTemplateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafTemplateRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWafTemplateRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWafTemplateRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWafTemplateRulesResponse) GetBody() *ListWafTemplateRulesResponseBody {
	return s.Body
}

func (s *ListWafTemplateRulesResponse) SetHeaders(v map[string]*string) *ListWafTemplateRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWafTemplateRulesResponse) SetStatusCode(v int32) *ListWafTemplateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafTemplateRulesResponse) SetBody(v *ListWafTemplateRulesResponseBody) *ListWafTemplateRulesResponse {
	s.Body = v
	return s
}

func (s *ListWafTemplateRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
