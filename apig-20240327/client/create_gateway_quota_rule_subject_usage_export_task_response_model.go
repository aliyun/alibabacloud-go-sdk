// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayQuotaRuleSubjectUsageExportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse
	GetBody() *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody
}

type CreateGatewayQuotaRuleSubjectUsageExportTaskResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) GetBody() *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	return s.Body
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) SetHeaders(v map[string]*string) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) SetStatusCode(v int32) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) SetBody(v *CreateGatewayQuotaRuleSubjectUsageExportTaskResponseBody) *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
