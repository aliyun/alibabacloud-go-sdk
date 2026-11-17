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
	// The optional token used for request idempotency. It can be up to 128 bytes in length after leading and trailing whitespaces are removed. We recommend that you use ASCII characters.
	//
	// When the same user retries with the same token and identical request parameters in the current region, the existing task is reused and its current status is returned. If the same token corresponds to different parameters, a parameter error is returned. The compared parameters include gatewayId, ruleId, subjectId, filterFailedRequests, format, and whether startTime and endTime are passed and their values. Omitting the time is not equivalent to explicitly passing 0.
	//
	// Reusing a task does not recalculate the time range or regenerate the file. When you need to re-export data (including when the original task fails or the file expires), use a new token.
	//
	// example:
	//
	// usage-export-20260729-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The end time of the export time range. Use a non-negative Unix timestamp in seconds. This time is exclusive.
	//
	// If only endTime is specified, one rule period is deduced forward from this time as the startTime. If neither is specified, the end time is the task creation time. If both are specified, endTime must be greater than or equal to startTime, and the span cannot exceed one rule period.
	//
	// example:
	//
	// 1783935200
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Specifies whether to export only consumption details with valid usage. Default value: false.
	//
	// - false: Does not filter records by usage.
	//
	// - true: Filters out records with zero usage based on the token or credit metering dimension of the rule, and retains only records with usage.
	//
	// example:
	//
	// true
	FilterFailedRequests *bool `json:"filterFailedRequests,omitempty" xml:"filterFailedRequests,omitempty"`
	// The format of the exported file. Currently, only CSV_GZIP is supported. If omitted or an empty string is passed, this format is used by default. The file is a GZIP-compressed CSV file named `{taskId}.csv.gz` and must be decompressed before reading.
	//
	// example:
	//
	// CSV_GZIP
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The start time of the export time range. Use a non-negative Unix timestamp in seconds. This time is inclusive.
	//
	// - If neither startTime nor endTime is specified: The start point of the current rule epoch to the node creation time is used.
	//
	// - If only startTime is specified: One rule epoch is deduced backward from this time as the endTime.
	//
	// - If both are specified: endTime must be greater than or equal to startTime, and the span cannot exceed one rule epoch.
	//
	// The actual export start point is also limited by the rule creation time, rule reset, or the effective start point of purge history records. You can confirm this by querying the effectiveWindowStart returned by the node.
	//
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
