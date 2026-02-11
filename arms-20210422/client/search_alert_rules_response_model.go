// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *SearchAlertRulesResponseBody) *SearchAlertRulesResponse
	GetBody() *SearchAlertRulesResponseBody
}

type SearchAlertRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAlertRulesResponse) GetBody() *SearchAlertRulesResponseBody {
	return s.Body
}

func (s *SearchAlertRulesResponse) SetHeaders(v map[string]*string) *SearchAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertRulesResponse) SetStatusCode(v int32) *SearchAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertRulesResponse) SetBody(v *SearchAlertRulesResponseBody) *SearchAlertRulesResponse {
	s.Body = v
	return s
}

func (s *SearchAlertRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
