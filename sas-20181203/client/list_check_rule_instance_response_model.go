// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckRuleInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckRuleInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckRuleInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckRuleInstanceResponseBody) *ListCheckRuleInstanceResponse
	GetBody() *ListCheckRuleInstanceResponseBody
}

type ListCheckRuleInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckRuleInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckRuleInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRuleInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListCheckRuleInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckRuleInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckRuleInstanceResponse) GetBody() *ListCheckRuleInstanceResponseBody {
	return s.Body
}

func (s *ListCheckRuleInstanceResponse) SetHeaders(v map[string]*string) *ListCheckRuleInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListCheckRuleInstanceResponse) SetStatusCode(v int32) *ListCheckRuleInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckRuleInstanceResponse) SetBody(v *ListCheckRuleInstanceResponseBody) *ListCheckRuleInstanceResponse {
	s.Body = v
	return s
}

func (s *ListCheckRuleInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
