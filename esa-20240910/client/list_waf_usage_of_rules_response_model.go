// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafUsageOfRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWafUsageOfRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWafUsageOfRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListWafUsageOfRulesResponseBody) *ListWafUsageOfRulesResponse
	GetBody() *ListWafUsageOfRulesResponseBody
}

type ListWafUsageOfRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafUsageOfRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafUsageOfRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWafUsageOfRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWafUsageOfRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWafUsageOfRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWafUsageOfRulesResponse) GetBody() *ListWafUsageOfRulesResponseBody {
	return s.Body
}

func (s *ListWafUsageOfRulesResponse) SetHeaders(v map[string]*string) *ListWafUsageOfRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWafUsageOfRulesResponse) SetStatusCode(v int32) *ListWafUsageOfRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafUsageOfRulesResponse) SetBody(v *ListWafUsageOfRulesResponseBody) *ListWafUsageOfRulesResponse {
	s.Body = v
	return s
}

func (s *ListWafUsageOfRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
