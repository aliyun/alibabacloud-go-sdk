// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafRulesetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWafRulesetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWafRulesetsResponse
	GetStatusCode() *int32
	SetBody(v *ListWafRulesetsResponseBody) *ListWafRulesetsResponse
	GetBody() *ListWafRulesetsResponseBody
}

type ListWafRulesetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWafRulesetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWafRulesetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWafRulesetsResponse) GoString() string {
	return s.String()
}

func (s *ListWafRulesetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWafRulesetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWafRulesetsResponse) GetBody() *ListWafRulesetsResponseBody {
	return s.Body
}

func (s *ListWafRulesetsResponse) SetHeaders(v map[string]*string) *ListWafRulesetsResponse {
	s.Headers = v
	return s
}

func (s *ListWafRulesetsResponse) SetStatusCode(v int32) *ListWafRulesetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWafRulesetsResponse) SetBody(v *ListWafRulesetsResponseBody) *ListWafRulesetsResponse {
	s.Body = v
	return s
}

func (s *ListWafRulesetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
