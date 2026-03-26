// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkspacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkspacesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListWorkspacesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListWorkspacesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListWorkspacesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkspacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkspacesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkspacesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListWorkspacesResponseBody
	GetTotalCount() *int32
	SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody
	GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces
}

type ListWorkspacesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// otqqmWigyzM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 43ED6679-672D-5EB6-B3AC-E0EF90129DFF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 31
	TotalCount *int32                                  `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"workspaces,omitempty" xml:"workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkspacesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWorkspacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWorkspacesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkspacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkspacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkspacesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkspacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkspacesResponseBody) GetWorkspaces() []*ListWorkspacesResponseBodyWorkspaces {
	return s.Workspaces
}

func (s *ListWorkspacesResponseBody) SetCode(v string) *ListWorkspacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetHttpStatusCode(v int32) *ListWorkspacesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetMessage(v string) *ListWorkspacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetSuccess(v bool) *ListWorkspacesResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

func (s *ListWorkspacesResponseBody) Validate() error {
	if s.Workspaces != nil {
		for _, item := range s.Workspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkspacesResponseBodyWorkspaces struct {
	// API Host。
	//
	// example:
	//
	// llm-34o9ts1dai60z5sf.cn-beijing.maas.aliyuncs.com
	ApiHost *string `json:"apiHost,omitempty" xml:"apiHost,omitempty"`
	// example:
	//
	// 1774338222000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// cn-beijing
	Region     *string `json:"region,omitempty" xml:"region,omitempty"`
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// example:
	//
	// global
	ServiceSite *string `json:"serviceSite,omitempty" xml:"serviceSite,omitempty"`
	// example:
	//
	// llm-34o9ts1dai60z5sf
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// example:
	//
	// default
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetApiHost() *string {
	return s.ApiHost
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetRegion() *string {
	return s.Region
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetRegionName() *string {
	return s.RegionName
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetServiceSite() *string {
	return s.ServiceSite
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyWorkspaces) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetApiHost(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ApiHost = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetGmtCreate(v int64) *ListWorkspacesResponseBodyWorkspaces {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRegion(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Region = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRegionName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RegionName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetServiceSite(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.ServiceSite = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) Validate() error {
	return dara.Validate(s)
}
