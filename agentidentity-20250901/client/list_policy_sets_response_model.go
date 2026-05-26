// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicySetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPolicySetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPolicySetsResponse
	GetStatusCode() *int32
	SetBody(v *ListPolicySetsResponseBody) *ListPolicySetsResponse
	GetBody() *ListPolicySetsResponseBody
}

type ListPolicySetsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicySetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicySetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPolicySetsResponse) GoString() string {
	return s.String()
}

func (s *ListPolicySetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPolicySetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPolicySetsResponse) GetBody() *ListPolicySetsResponseBody {
	return s.Body
}

func (s *ListPolicySetsResponse) SetHeaders(v map[string]*string) *ListPolicySetsResponse {
	s.Headers = v
	return s
}

func (s *ListPolicySetsResponse) SetStatusCode(v int32) *ListPolicySetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicySetsResponse) SetBody(v *ListPolicySetsResponseBody) *ListPolicySetsResponse {
	s.Body = v
	return s
}

func (s *ListPolicySetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
