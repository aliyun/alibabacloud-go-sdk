// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemConfigRulesResponseBody) *ListSystemConfigRulesResponse
	GetBody() *ListSystemConfigRulesResponseBody
}

type ListSystemConfigRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemConfigRulesResponse) GetBody() *ListSystemConfigRulesResponseBody {
	return s.Body
}

func (s *ListSystemConfigRulesResponse) SetHeaders(v map[string]*string) *ListSystemConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemConfigRulesResponse) SetStatusCode(v int32) *ListSystemConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemConfigRulesResponse) SetBody(v *ListSystemConfigRulesResponseBody) *ListSystemConfigRulesResponse {
	s.Body = v
	return s
}

func (s *ListSystemConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
