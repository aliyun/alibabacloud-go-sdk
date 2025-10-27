// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRuleTargetAllResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRuleTargetAllResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRuleTargetAllResponse
	GetStatusCode() *int32
	SetBody(v *ListRuleTargetAllResponseBody) *ListRuleTargetAllResponse
	GetBody() *ListRuleTargetAllResponseBody
}

type ListRuleTargetAllResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRuleTargetAllResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRuleTargetAllResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRuleTargetAllResponse) GoString() string {
	return s.String()
}

func (s *ListRuleTargetAllResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRuleTargetAllResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRuleTargetAllResponse) GetBody() *ListRuleTargetAllResponseBody {
	return s.Body
}

func (s *ListRuleTargetAllResponse) SetHeaders(v map[string]*string) *ListRuleTargetAllResponse {
	s.Headers = v
	return s
}

func (s *ListRuleTargetAllResponse) SetStatusCode(v int32) *ListRuleTargetAllResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRuleTargetAllResponse) SetBody(v *ListRuleTargetAllResponseBody) *ListRuleTargetAllResponse {
	s.Body = v
	return s
}

func (s *ListRuleTargetAllResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
