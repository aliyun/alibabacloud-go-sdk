// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotificationPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNotificationPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNotificationPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListNotificationPoliciesResponseBody) *ListNotificationPoliciesResponse
	GetBody() *ListNotificationPoliciesResponseBody
}

type ListNotificationPoliciesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNotificationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNotificationPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNotificationPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNotificationPoliciesResponse) GetBody() *ListNotificationPoliciesResponseBody {
	return s.Body
}

func (s *ListNotificationPoliciesResponse) SetHeaders(v map[string]*string) *ListNotificationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListNotificationPoliciesResponse) SetStatusCode(v int32) *ListNotificationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNotificationPoliciesResponse) SetBody(v *ListNotificationPoliciesResponseBody) *ListNotificationPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListNotificationPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
