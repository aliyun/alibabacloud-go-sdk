// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManagedRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManagedRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListManagedRulesResponseBody) *ListManagedRulesResponse
	GetBody() *ListManagedRulesResponseBody
}

type ListManagedRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManagedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManagedRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManagedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManagedRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManagedRulesResponse) GetBody() *ListManagedRulesResponseBody {
	return s.Body
}

func (s *ListManagedRulesResponse) SetHeaders(v map[string]*string) *ListManagedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListManagedRulesResponse) SetStatusCode(v int32) *ListManagedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManagedRulesResponse) SetBody(v *ListManagedRulesResponseBody) *ListManagedRulesResponse {
	s.Body = v
	return s
}

func (s *ListManagedRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
