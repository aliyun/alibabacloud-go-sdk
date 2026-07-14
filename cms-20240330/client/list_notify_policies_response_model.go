// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotifyPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNotifyPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNotifyPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListNotifyPoliciesResponseBody) *ListNotifyPoliciesResponse
	GetBody() *ListNotifyPoliciesResponseBody
}

type ListNotifyPoliciesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNotifyPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNotifyPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNotifyPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListNotifyPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNotifyPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNotifyPoliciesResponse) GetBody() *ListNotifyPoliciesResponseBody {
	return s.Body
}

func (s *ListNotifyPoliciesResponse) SetHeaders(v map[string]*string) *ListNotifyPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListNotifyPoliciesResponse) SetStatusCode(v int32) *ListNotifyPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNotifyPoliciesResponse) SetBody(v *ListNotifyPoliciesResponseBody) *ListNotifyPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListNotifyPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
