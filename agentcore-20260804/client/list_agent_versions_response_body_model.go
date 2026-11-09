// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentVersionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAgentVersionsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListAgentVersionsResponseBodyItems) *ListAgentVersionsResponseBody
	GetItems() []*ListAgentVersionsResponseBodyItems
	SetMaxResults(v int32) *ListAgentVersionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAgentVersionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAgentVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgentVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentVersionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgentVersionsResponseBody
	GetTotalCount() *int64
}

type ListAgentVersionsResponseBody struct {
	// The business status code. The value is SUCCESS if the request is successful, or an error code if the request fails.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of fixed versions on the current page, sorted in descending order of creation time. The configuration is not included.
	Items []*ListAgentVersionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries per page used in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. The value is success if the request is successful, or an error message if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. Pass this value as the nextToken request parameter to query the next page. An empty value indicates that no more data is available.
	//
	// example:
	//
	// YWdlbnQtdmVyc2lvbnM6MjA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// 5E8F3C2A-1B4D-4F6E-9A7C-0D2B8E6F4A1C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of undeleted fixed versions of the agent, excluding the system-maintained latest version.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAgentVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAgentVersionsResponseBody) GetItems() []*ListAgentVersionsResponseBodyItems {
	return s.Items
}

func (s *ListAgentVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgentVersionsResponseBody) SetCode(v string) *ListAgentVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetHttpStatusCode(v int32) *ListAgentVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetItems(v []*ListAgentVersionsResponseBodyItems) *ListAgentVersionsResponseBody {
	s.Items = v
	return s
}

func (s *ListAgentVersionsResponseBody) SetMaxResults(v int32) *ListAgentVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetMessage(v string) *ListAgentVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetNextToken(v string) *ListAgentVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetRequestId(v string) *ListAgentVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetSuccess(v bool) *ListAgentVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentVersionsResponseBody) SetTotalCount(v int64) *ListAgentVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentVersionsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentVersionsResponseBodyItems struct {
	// The ID of the managed or custom agent to which the version belongs.
	//
	// example:
	//
	// agent-0123456789abcdef0123456789abcdef
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The version number generated by the system in ascending order for the agent, such as v1 and v2. It is used when querying, deleting, and referencing the version in service endpoints.
	//
	// example:
	//
	// v1
	AgentVersion *string `json:"agentVersion,omitempty" xml:"agentVersion,omitempty"`
	// Indicates only whether the version template and artifact archiving are complete, and does not represent the deployment or running status. READY indicates complete, and INVALID indicates missing or inconsistent archiving.
	//
	// example:
	//
	// READY
	ArtifactStatus *string `json:"artifactStatus,omitempty" xml:"artifactStatus,omitempty"`
	// The time when the version was saved, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-12T01:02:03Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The UID of the Alibaba Cloud account of the caller who saved the version. If a RAM user makes the call, this is the UID of the Alibaba Cloud account to which the RAM user belongs.
	//
	// example:
	//
	// 1234567890123456
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The version description provided when saving the version, up to 1,024 characters. It is empty if no description was provided when saving.
	//
	// example:
	//
	// Stable production version
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The effective release serial number of the agent when the version was saved, which is the effectiveSpecVersion at that time. This value does not change after the version is saved.
	//
	// example:
	//
	// 2
	TargetSpecVersion *int64 `json:"targetSpecVersion,omitempty" xml:"targetSpecVersion,omitempty"`
	// The display name specified by the user when saving the version. For earlier saved versions without a recorded name, the agentVersion is returned.
	//
	// example:
	//
	// Production stable version V1.0
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty"`
	// The ID of the AgentCore workspace where the agent that owns the version resides.
	//
	// example:
	//
	// ws-0123456789abcdef01234
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListAgentVersionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAgentVersionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAgentVersionsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentVersionsResponseBodyItems) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *ListAgentVersionsResponseBodyItems) GetArtifactStatus() *string {
	return s.ArtifactStatus
}

func (s *ListAgentVersionsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAgentVersionsResponseBodyItems) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListAgentVersionsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListAgentVersionsResponseBodyItems) GetTargetSpecVersion() *int64 {
	return s.TargetSpecVersion
}

func (s *ListAgentVersionsResponseBodyItems) GetVersionName() *string {
	return s.VersionName
}

func (s *ListAgentVersionsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAgentVersionsResponseBodyItems) SetAgentId(v string) *ListAgentVersionsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetAgentVersion(v string) *ListAgentVersionsResponseBodyItems {
	s.AgentVersion = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetArtifactStatus(v string) *ListAgentVersionsResponseBodyItems {
	s.ArtifactStatus = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetCreatedAt(v string) *ListAgentVersionsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetCreatedBy(v string) *ListAgentVersionsResponseBodyItems {
	s.CreatedBy = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetDescription(v string) *ListAgentVersionsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetTargetSpecVersion(v int64) *ListAgentVersionsResponseBodyItems {
	s.TargetSpecVersion = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetVersionName(v string) *ListAgentVersionsResponseBodyItems {
	s.VersionName = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) SetWorkspaceId(v string) *ListAgentVersionsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListAgentVersionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
