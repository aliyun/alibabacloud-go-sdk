// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlows(v []*ListFlowsResponseBodyFlows) *ListFlowsResponseBody
	GetFlows() []*ListFlowsResponseBodyFlows
	SetMaxResults(v int32) *ListFlowsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListFlowsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFlowsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListFlowsResponseBody
	GetTotalCount() *int32
}

type ListFlowsResponseBody struct {
	// The list of connector flows.
	Flows []*ListFlowsResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// The page size.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// AAAAAVY3rYiv9VoUJQSiCitgjgRCGYBJsWLezzDdq2+aDNB4j8bgHiGAwZWnCMJPepC+WYZ4J1BLMwE7gJm++1ku/AI=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 02FB9782-5390-5DF9-A1DA-9F2DE2FA1E3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBody) GetFlows() []*ListFlowsResponseBodyFlows {
	return s.Flows
}

func (s *ListFlowsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFlowsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFlowsResponseBody) SetFlows(v []*ListFlowsResponseBodyFlows) *ListFlowsResponseBody {
	s.Flows = v
	return s
}

func (s *ListFlowsResponseBody) SetMaxResults(v int32) *ListFlowsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFlowsResponseBody) SetNextToken(v string) *ListFlowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFlowsResponseBody) SetRequestId(v string) *ListFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResponseBody) SetTotalCount(v int32) *ListFlowsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFlowsResponseBody) Validate() error {
	if s.Flows != nil {
		for _, item := range s.Flows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowsResponseBodyFlows struct {
	// Indicates whether the connector flow is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The connector flow description.
	//
	// example:
	//
	// 在未认证的公众号中更实用Qwen3开源版本自动回复
	FlowDesc *string `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	// The connector flow ID.
	//
	// example:
	//
	// flow-xxxxxxxx
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The connector flow name.
	//
	// example:
	//
	// 钉钉群聊 - 流式调用百炼应用-小助手
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The connector flow template content.
	//
	// example:
	//
	// {}
	FlowTemplate *string `json:"FlowTemplate,omitempty" xml:"FlowTemplate,omitempty"`
	// The connector flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The connector flow version status.
	//
	// example:
	//
	// 0 - 草稿
	//
	// 1- 发布
	//
	// 2- 下线
	FlowVersionStatus *string `json:"FlowVersionStatus,omitempty" xml:"FlowVersionStatus,omitempty"`
	// The time when the connector flow was created.
	//
	// example:
	//
	// 2025-12-30T02:29:51Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the connector flow was last modified.
	//
	// example:
	//
	// 2025-12-30T02:29:48Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The released version number.
	//
	// example:
	//
	// 1
	ReleasedVersion *int32 `json:"ReleasedVersion,omitempty" xml:"ReleasedVersion,omitempty"`
	// The tag values.
	Tags []*ListFlowsResponseBodyFlowsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The webhook URL.
	//
	// example:
	//
	// https://{uid}.appflow.aliyunnest.com/webhook/xxxxxx
	WebhookUrl *string `json:"WebhookUrl,omitempty" xml:"WebhookUrl,omitempty"`
}

func (s ListFlowsResponseBodyFlows) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlows) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListFlowsResponseBodyFlows) GetFlowDesc() *string {
	return s.FlowDesc
}

func (s *ListFlowsResponseBodyFlows) GetFlowId() *string {
	return s.FlowId
}

func (s *ListFlowsResponseBodyFlows) GetFlowName() *string {
	return s.FlowName
}

func (s *ListFlowsResponseBodyFlows) GetFlowTemplate() *string {
	return s.FlowTemplate
}

func (s *ListFlowsResponseBodyFlows) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *ListFlowsResponseBodyFlows) GetFlowVersionStatus() *string {
	return s.FlowVersionStatus
}

func (s *ListFlowsResponseBodyFlows) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListFlowsResponseBodyFlows) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListFlowsResponseBodyFlows) GetReleasedVersion() *int32 {
	return s.ReleasedVersion
}

func (s *ListFlowsResponseBodyFlows) GetTags() []*ListFlowsResponseBodyFlowsTags {
	return s.Tags
}

func (s *ListFlowsResponseBodyFlows) GetWebhookUrl() *string {
	return s.WebhookUrl
}

func (s *ListFlowsResponseBodyFlows) SetEnabled(v bool) *ListFlowsResponseBodyFlows {
	s.Enabled = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowDesc(v string) *ListFlowsResponseBodyFlows {
	s.FlowDesc = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowId(v string) *ListFlowsResponseBodyFlows {
	s.FlowId = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowName(v string) *ListFlowsResponseBodyFlows {
	s.FlowName = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowTemplate(v string) *ListFlowsResponseBodyFlows {
	s.FlowTemplate = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowVersion(v string) *ListFlowsResponseBodyFlows {
	s.FlowVersion = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowVersionStatus(v string) *ListFlowsResponseBodyFlows {
	s.FlowVersionStatus = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetGmtCreate(v string) *ListFlowsResponseBodyFlows {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetGmtModified(v string) *ListFlowsResponseBodyFlows {
	s.GmtModified = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetReleasedVersion(v int32) *ListFlowsResponseBodyFlows {
	s.ReleasedVersion = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetTags(v []*ListFlowsResponseBodyFlowsTags) *ListFlowsResponseBodyFlows {
	s.Tags = v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetWebhookUrl(v string) *ListFlowsResponseBodyFlows {
	s.WebhookUrl = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlowsResponseBodyFlowsTags struct {
	// The tag key. The tag key can be up to 64 characters in length.
	//
	// example:
	//
	// tuoluo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// example:
	//
	// {\\"connectorId\\": \\"connector-5e43ef26b53e4a158529\\", \\"authConfig\\": \\"{\\\\\\"apiKey\\\\\\": \\\\\\"sk-77e7c7ed775f4fd388d505d37b2b1b33\\\\\\"}\\"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFlowsResponseBodyFlowsTags) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResponseBodyFlowsTags) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlowsTags) GetKey() *string {
	return s.Key
}

func (s *ListFlowsResponseBodyFlowsTags) GetValue() *string {
	return s.Value
}

func (s *ListFlowsResponseBodyFlowsTags) SetKey(v string) *ListFlowsResponseBodyFlowsTags {
	s.Key = &v
	return s
}

func (s *ListFlowsResponseBodyFlowsTags) SetValue(v string) *ListFlowsResponseBodyFlowsTags {
	s.Value = &v
	return s
}

func (s *ListFlowsResponseBodyFlowsTags) Validate() error {
	return dara.Validate(s)
}
