// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayQuotaRuleSubjectUsageExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest
	GetClientToken() *string
	SetEndTime(v int64) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest
	GetEndTime() *int64
	SetFilterFailedRequests(v bool) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest
	GetFilterFailedRequests() *bool
	SetFormat(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest
	GetFormat() *string
	SetStartTime(v int64) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest
	GetStartTime() *int64
}

type CreateGatewayQuotaRuleSubjectUsageExportTaskRequest struct {
	// example:
	//
	// usage-export-20260729-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// 1783935200
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// true
	FilterFailedRequests *bool `json:"filterFailedRequests,omitempty" xml:"filterFailedRequests,omitempty"`
	// example:
	//
	// CSV_GZIP
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// 1783931600
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) GetFilterFailedRequests() *bool {
	return s.FilterFailedRequests
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) GetFormat() *string {
	return s.Format
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) SetClientToken(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) SetEndTime(v int64) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest {
	s.EndTime = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) SetFilterFailedRequests(v bool) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest {
	s.FilterFailedRequests = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) SetFormat(v string) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest {
	s.Format = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) SetStartTime(v int64) *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateGatewayQuotaRuleSubjectUsageExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
