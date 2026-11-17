// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationCustomRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationCustomRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationCustomRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationCustomRulesResponseBody) *ListApplicationCustomRulesResponse
	GetBody() *ListApplicationCustomRulesResponseBody
}

type ListApplicationCustomRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationCustomRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationCustomRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCustomRulesResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationCustomRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationCustomRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationCustomRulesResponse) GetBody() *ListApplicationCustomRulesResponseBody {
	return s.Body
}

func (s *ListApplicationCustomRulesResponse) SetHeaders(v map[string]*string) *ListApplicationCustomRulesResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationCustomRulesResponse) SetStatusCode(v int32) *ListApplicationCustomRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationCustomRulesResponse) SetBody(v *ListApplicationCustomRulesResponseBody) *ListApplicationCustomRulesResponse {
	s.Body = v
	return s
}

func (s *ListApplicationCustomRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
