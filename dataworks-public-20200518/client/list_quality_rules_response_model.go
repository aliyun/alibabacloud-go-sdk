// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityRulesResponseBody) *ListQualityRulesResponse
	GetBody() *ListQualityRulesResponseBody
}

type ListQualityRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRulesResponse) GoString() string {
	return s.String()
}

func (s *ListQualityRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityRulesResponse) GetBody() *ListQualityRulesResponseBody {
	return s.Body
}

func (s *ListQualityRulesResponse) SetHeaders(v map[string]*string) *ListQualityRulesResponse {
	s.Headers = v
	return s
}

func (s *ListQualityRulesResponse) SetStatusCode(v int32) *ListQualityRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityRulesResponse) SetBody(v *ListQualityRulesResponseBody) *ListQualityRulesResponse {
	s.Body = v
	return s
}

func (s *ListQualityRulesResponse) Validate() error {
	return dara.Validate(s)
}
