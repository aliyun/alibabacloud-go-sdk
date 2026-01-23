// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityRuleTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQualityRuleTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQualityRuleTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListQualityRuleTasksResponseBody) *ListQualityRuleTasksResponse
	GetBody() *ListQualityRuleTasksResponseBody
}

type ListQualityRuleTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQualityRuleTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQualityRuleTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQualityRuleTasksResponse) GoString() string {
	return s.String()
}

func (s *ListQualityRuleTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQualityRuleTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQualityRuleTasksResponse) GetBody() *ListQualityRuleTasksResponseBody {
	return s.Body
}

func (s *ListQualityRuleTasksResponse) SetHeaders(v map[string]*string) *ListQualityRuleTasksResponse {
	s.Headers = v
	return s
}

func (s *ListQualityRuleTasksResponse) SetStatusCode(v int32) *ListQualityRuleTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQualityRuleTasksResponse) SetBody(v *ListQualityRuleTasksResponseBody) *ListQualityRuleTasksResponse {
	s.Body = v
	return s
}

func (s *ListQualityRuleTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
