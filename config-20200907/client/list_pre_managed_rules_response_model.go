// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPreManagedRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPreManagedRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPreManagedRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListPreManagedRulesResponseBody) *ListPreManagedRulesResponse
	GetBody() *ListPreManagedRulesResponseBody
}

type ListPreManagedRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPreManagedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPreManagedRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPreManagedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPreManagedRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPreManagedRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPreManagedRulesResponse) GetBody() *ListPreManagedRulesResponseBody {
	return s.Body
}

func (s *ListPreManagedRulesResponse) SetHeaders(v map[string]*string) *ListPreManagedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPreManagedRulesResponse) SetStatusCode(v int32) *ListPreManagedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPreManagedRulesResponse) SetBody(v *ListPreManagedRulesResponseBody) *ListPreManagedRulesResponse {
	s.Body = v
	return s
}

func (s *ListPreManagedRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
