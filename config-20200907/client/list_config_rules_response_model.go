// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigRulesResponseBody) *ListConfigRulesResponse
	GetBody() *ListConfigRulesResponseBody
}

type ListConfigRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigRulesResponse) GetBody() *ListConfigRulesResponseBody {
	return s.Body
}

func (s *ListConfigRulesResponse) SetHeaders(v map[string]*string) *ListConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRulesResponse) SetStatusCode(v int32) *ListConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRulesResponse) SetBody(v *ListConfigRulesResponseBody) *ListConfigRulesResponse {
	s.Body = v
	return s
}

func (s *ListConfigRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
