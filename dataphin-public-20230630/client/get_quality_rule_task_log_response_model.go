// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTaskLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityRuleTaskLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityRuleTaskLogResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityRuleTaskLogResponseBody) *GetQualityRuleTaskLogResponse
	GetBody() *GetQualityRuleTaskLogResponseBody
}

type GetQualityRuleTaskLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityRuleTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityRuleTaskLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTaskLogResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTaskLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityRuleTaskLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityRuleTaskLogResponse) GetBody() *GetQualityRuleTaskLogResponseBody {
	return s.Body
}

func (s *GetQualityRuleTaskLogResponse) SetHeaders(v map[string]*string) *GetQualityRuleTaskLogResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleTaskLogResponse) SetStatusCode(v int32) *GetQualityRuleTaskLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityRuleTaskLogResponse) SetBody(v *GetQualityRuleTaskLogResponseBody) *GetQualityRuleTaskLogResponse {
	s.Body = v
	return s
}

func (s *GetQualityRuleTaskLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
