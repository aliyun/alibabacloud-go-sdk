// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivacyRuleSharedMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivacyRuleSharedMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivacyRuleSharedMemberResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivacyRuleSharedMemberResponseBody) *ListPrivacyRuleSharedMemberResponse
	GetBody() *ListPrivacyRuleSharedMemberResponseBody
}

type ListPrivacyRuleSharedMemberResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivacyRuleSharedMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivacyRuleSharedMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivacyRuleSharedMemberResponse) GoString() string {
	return s.String()
}

func (s *ListPrivacyRuleSharedMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivacyRuleSharedMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivacyRuleSharedMemberResponse) GetBody() *ListPrivacyRuleSharedMemberResponseBody {
	return s.Body
}

func (s *ListPrivacyRuleSharedMemberResponse) SetHeaders(v map[string]*string) *ListPrivacyRuleSharedMemberResponse {
	s.Headers = v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponse) SetStatusCode(v int32) *ListPrivacyRuleSharedMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponse) SetBody(v *ListPrivacyRuleSharedMemberResponseBody) *ListPrivacyRuleSharedMemberResponse {
	s.Body = v
	return s
}

func (s *ListPrivacyRuleSharedMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
