// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitQualityRuleTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitQualityRuleTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitQualityRuleTasksResponse
	GetStatusCode() *int32
	SetBody(v *SubmitQualityRuleTasksResponseBody) *SubmitQualityRuleTasksResponse
	GetBody() *SubmitQualityRuleTasksResponseBody
}

type SubmitQualityRuleTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitQualityRuleTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitQualityRuleTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitQualityRuleTasksResponse) GoString() string {
	return s.String()
}

func (s *SubmitQualityRuleTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitQualityRuleTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitQualityRuleTasksResponse) GetBody() *SubmitQualityRuleTasksResponseBody {
	return s.Body
}

func (s *SubmitQualityRuleTasksResponse) SetHeaders(v map[string]*string) *SubmitQualityRuleTasksResponse {
	s.Headers = v
	return s
}

func (s *SubmitQualityRuleTasksResponse) SetStatusCode(v int32) *SubmitQualityRuleTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQualityRuleTasksResponse) SetBody(v *SubmitQualityRuleTasksResponseBody) *SubmitQualityRuleTasksResponse {
	s.Body = v
	return s
}

func (s *SubmitQualityRuleTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
