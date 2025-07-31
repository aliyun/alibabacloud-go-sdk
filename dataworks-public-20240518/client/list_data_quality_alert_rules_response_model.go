// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityAlertRulesResponseBody) *ListDataQualityAlertRulesResponse
	GetBody() *ListDataQualityAlertRulesResponseBody
}

type ListDataQualityAlertRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityAlertRulesResponse) GetBody() *ListDataQualityAlertRulesResponseBody {
	return s.Body
}

func (s *ListDataQualityAlertRulesResponse) SetHeaders(v map[string]*string) *ListDataQualityAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityAlertRulesResponse) SetStatusCode(v int32) *ListDataQualityAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityAlertRulesResponse) SetBody(v *ListDataQualityAlertRulesResponseBody) *ListDataQualityAlertRulesResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityAlertRulesResponse) Validate() error {
	return dara.Validate(s)
}
