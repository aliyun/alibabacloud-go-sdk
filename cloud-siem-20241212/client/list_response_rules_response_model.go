// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResponseRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResponseRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResponseRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListResponseRulesResponseBody) *ListResponseRulesResponse
	GetBody() *ListResponseRulesResponseBody
}

type ListResponseRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResponseRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResponseRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRulesResponse) GoString() string {
	return s.String()
}

func (s *ListResponseRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResponseRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResponseRulesResponse) GetBody() *ListResponseRulesResponseBody {
	return s.Body
}

func (s *ListResponseRulesResponse) SetHeaders(v map[string]*string) *ListResponseRulesResponse {
	s.Headers = v
	return s
}

func (s *ListResponseRulesResponse) SetStatusCode(v int32) *ListResponseRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResponseRulesResponse) SetBody(v *ListResponseRulesResponseBody) *ListResponseRulesResponse {
	s.Body = v
	return s
}

func (s *ListResponseRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
