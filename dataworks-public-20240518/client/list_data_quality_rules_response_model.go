// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityRulesResponseBody) *ListDataQualityRulesResponse
	GetBody() *ListDataQualityRulesResponseBody
}

type ListDataQualityRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityRulesResponse) GetBody() *ListDataQualityRulesResponseBody {
	return s.Body
}

func (s *ListDataQualityRulesResponse) SetHeaders(v map[string]*string) *ListDataQualityRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityRulesResponse) SetStatusCode(v int32) *ListDataQualityRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityRulesResponse) SetBody(v *ListDataQualityRulesResponseBody) *ListDataQualityRulesResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
