// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientRuleFileTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileProtectClientRuleFileTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileProtectClientRuleFileTypeResponse
	GetStatusCode() *int32
	SetBody(v *ListFileProtectClientRuleFileTypeResponseBody) *ListFileProtectClientRuleFileTypeResponse
	GetBody() *ListFileProtectClientRuleFileTypeResponseBody
}

type ListFileProtectClientRuleFileTypeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileProtectClientRuleFileTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileProtectClientRuleFileTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientRuleFileTypeResponse) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientRuleFileTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileProtectClientRuleFileTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileProtectClientRuleFileTypeResponse) GetBody() *ListFileProtectClientRuleFileTypeResponseBody {
	return s.Body
}

func (s *ListFileProtectClientRuleFileTypeResponse) SetHeaders(v map[string]*string) *ListFileProtectClientRuleFileTypeResponse {
	s.Headers = v
	return s
}

func (s *ListFileProtectClientRuleFileTypeResponse) SetStatusCode(v int32) *ListFileProtectClientRuleFileTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileProtectClientRuleFileTypeResponse) SetBody(v *ListFileProtectClientRuleFileTypeResponseBody) *ListFileProtectClientRuleFileTypeResponse {
	s.Body = v
	return s
}

func (s *ListFileProtectClientRuleFileTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
