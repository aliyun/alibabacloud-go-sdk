// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebFilterRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebFilterRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebFilterRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListWebFilterRulesResponseBody) *ListWebFilterRulesResponse
	GetBody() *ListWebFilterRulesResponseBody
}

type ListWebFilterRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebFilterRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebFilterRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebFilterRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWebFilterRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebFilterRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebFilterRulesResponse) GetBody() *ListWebFilterRulesResponseBody {
	return s.Body
}

func (s *ListWebFilterRulesResponse) SetHeaders(v map[string]*string) *ListWebFilterRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWebFilterRulesResponse) SetStatusCode(v int32) *ListWebFilterRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebFilterRulesResponse) SetBody(v *ListWebFilterRulesResponseBody) *ListWebFilterRulesResponse {
	s.Body = v
	return s
}

func (s *ListWebFilterRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
