// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizeRuleTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomizeRuleTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomizeRuleTestResultResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomizeRuleTestResultResponseBody) *ListCustomizeRuleTestResultResponse
	GetBody() *ListCustomizeRuleTestResultResponseBody
}

type ListCustomizeRuleTestResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomizeRuleTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomizeRuleTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponse) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomizeRuleTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomizeRuleTestResultResponse) GetBody() *ListCustomizeRuleTestResultResponseBody {
	return s.Body
}

func (s *ListCustomizeRuleTestResultResponse) SetHeaders(v map[string]*string) *ListCustomizeRuleTestResultResponse {
	s.Headers = v
	return s
}

func (s *ListCustomizeRuleTestResultResponse) SetStatusCode(v int32) *ListCustomizeRuleTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponse) SetBody(v *ListCustomizeRuleTestResultResponseBody) *ListCustomizeRuleTestResultResponse {
	s.Body = v
	return s
}

func (s *ListCustomizeRuleTestResultResponse) Validate() error {
	return dara.Validate(s)
}
