// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsWithTagRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsWithTagRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsWithTagRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsWithTagRulesResponseBody) *ListApplicationsWithTagRulesResponse
	GetBody() *ListApplicationsWithTagRulesResponseBody
}

type ListApplicationsWithTagRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsWithTagRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsWithTagRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsWithTagRulesResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsWithTagRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsWithTagRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsWithTagRulesResponse) GetBody() *ListApplicationsWithTagRulesResponseBody {
	return s.Body
}

func (s *ListApplicationsWithTagRulesResponse) SetHeaders(v map[string]*string) *ListApplicationsWithTagRulesResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsWithTagRulesResponse) SetStatusCode(v int32) *ListApplicationsWithTagRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsWithTagRulesResponse) SetBody(v *ListApplicationsWithTagRulesResponseBody) *ListApplicationsWithTagRulesResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsWithTagRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
