// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityRuleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityRuleTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityRuleTaskResponseBody) *GetQualityRuleTaskResponse
	GetBody() *GetQualityRuleTaskResponseBody
}

type GetQualityRuleTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityRuleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityRuleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTaskResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityRuleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityRuleTaskResponse) GetBody() *GetQualityRuleTaskResponseBody {
	return s.Body
}

func (s *GetQualityRuleTaskResponse) SetHeaders(v map[string]*string) *GetQualityRuleTaskResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleTaskResponse) SetStatusCode(v int32) *GetQualityRuleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityRuleTaskResponse) SetBody(v *GetQualityRuleTaskResponseBody) *GetQualityRuleTaskResponse {
	s.Body = v
	return s
}

func (s *GetQualityRuleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
