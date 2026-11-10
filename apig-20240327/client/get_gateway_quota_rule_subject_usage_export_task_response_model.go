// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleSubjectUsageExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayQuotaRuleSubjectUsageExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) *GetGatewayQuotaRuleSubjectUsageExportTaskResponse
	GetBody() *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody
}

type GetGatewayQuotaRuleSubjectUsageExportTaskResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponse) GetBody() *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	return s.Body
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponse) SetHeaders(v map[string]*string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponse) SetStatusCode(v int32) *GetGatewayQuotaRuleSubjectUsageExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponse) SetBody(v *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) *GetGatewayQuotaRuleSubjectUsageExportTaskResponse {
	s.Body = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
