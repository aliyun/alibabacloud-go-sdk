// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListResourceGroupsResponseBodyData) *ListResourceGroupsResponseBody
	GetData() []*ListResourceGroupsResponseBodyData
	SetMaxResults(v int32) *ListResourceGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListResourceGroupsResponseBody
	GetTotalCount() *int32
}

type ListResourceGroupsResponseBody struct {
	Data []*ListResourceGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 10。
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 1。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 16
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) GetData() []*ListResourceGroupsResponseBodyData {
	return s.Data
}

func (s *ListResourceGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceGroupsResponseBody) SetData(v []*ListResourceGroupsResponseBodyData) *ListResourceGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetMaxResults(v int32) *ListResourceGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetNextToken(v string) *ListResourceGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetTotalCount(v int32) *ListResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyData struct {
	AssociatedClusterTemplates []*ListResourceGroupsResponseBodyDataAssociatedClusterTemplates `json:"associatedClusterTemplates,omitempty" xml:"associatedClusterTemplates,omitempty" type:"Repeated"`
	AssociatedClusters         []*ListResourceGroupsResponseBodyDataAssociatedClusters         `json:"associatedClusters,omitempty" xml:"associatedClusters,omitempty" type:"Repeated"`
	AssociatedWorkspaces       []*ListResourceGroupsResponseBodyDataAssociatedWorkspaces       `json:"associatedWorkspaces,omitempty" xml:"associatedWorkspaces,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-26T02:10:04Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2。
	NodeMaxCount *int32 `json:"nodeMaxCount,omitempty" xml:"nodeMaxCount,omitempty"`
	// example:
	//
	// 1。
	NodeMinCount *int32 `json:"nodeMinCount,omitempty" xml:"nodeMinCount,omitempty"`
	// example:
	//
	// EMR_W1_SMALL。
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// example:
	//
	// postpaid
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// example:
	//
	// rg-acfm4ewqrznxxxx
	ResourceGroupId   *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ResourceGroupName *string `json:"resourceGroupName,omitempty" xml:"resourceGroupName,omitempty"`
	// example:
	//
	// CLUSTER_RESOURCE_GROUP。
	ResourceGroupType *string `json:"resourceGroupType,omitempty" xml:"resourceGroupType,omitempty"`
}

func (s ListResourceGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyData) GetAssociatedClusterTemplates() []*ListResourceGroupsResponseBodyDataAssociatedClusterTemplates {
	return s.AssociatedClusterTemplates
}

func (s *ListResourceGroupsResponseBodyData) GetAssociatedClusters() []*ListResourceGroupsResponseBodyDataAssociatedClusters {
	return s.AssociatedClusters
}

func (s *ListResourceGroupsResponseBodyData) GetAssociatedWorkspaces() []*ListResourceGroupsResponseBodyDataAssociatedWorkspaces {
	return s.AssociatedWorkspaces
}

func (s *ListResourceGroupsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceGroupsResponseBodyData) GetNodeMaxCount() *int32 {
	return s.NodeMaxCount
}

func (s *ListResourceGroupsResponseBodyData) GetNodeMinCount() *int32 {
	return s.NodeMinCount
}

func (s *ListResourceGroupsResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *ListResourceGroupsResponseBodyData) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListResourceGroupsResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceGroupsResponseBodyData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListResourceGroupsResponseBodyData) GetResourceGroupType() *string {
	return s.ResourceGroupType
}

func (s *ListResourceGroupsResponseBodyData) SetAssociatedClusterTemplates(v []*ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) *ListResourceGroupsResponseBodyData {
	s.AssociatedClusterTemplates = v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetAssociatedClusters(v []*ListResourceGroupsResponseBodyDataAssociatedClusters) *ListResourceGroupsResponseBodyData {
	s.AssociatedClusters = v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetAssociatedWorkspaces(v []*ListResourceGroupsResponseBodyDataAssociatedWorkspaces) *ListResourceGroupsResponseBodyData {
	s.AssociatedWorkspaces = v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetCreateTime(v string) *ListResourceGroupsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetNodeMaxCount(v int32) *ListResourceGroupsResponseBodyData {
	s.NodeMaxCount = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetNodeMinCount(v int32) *ListResourceGroupsResponseBodyData {
	s.NodeMinCount = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetNodeType(v string) *ListResourceGroupsResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetPaymentType(v string) *ListResourceGroupsResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetResourceGroupId(v string) *ListResourceGroupsResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetResourceGroupName(v string) *ListResourceGroupsResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) SetResourceGroupType(v string) *ListResourceGroupsResponseBodyData {
	s.ResourceGroupType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyData) Validate() error {
	if s.AssociatedClusterTemplates != nil {
		for _, item := range s.AssociatedClusterTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AssociatedClusters != nil {
		for _, item := range s.AssociatedClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AssociatedWorkspaces != nil {
		for _, item := range s.AssociatedWorkspaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyDataAssociatedClusterTemplates struct {
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// example:
	//
	// 78723f0dc36。
	TemplateId   *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) SetClusterType(v string) *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates {
	s.ClusterType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) SetTemplateId(v string) *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) SetTemplateName(v string) *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusterTemplates) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsResponseBodyDataAssociatedClusters struct {
	// example:
	//
	// c-3cd8ba94b36cxxxx
	ClusterId   *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
}

func (s ListResourceGroupsResponseBodyDataAssociatedClusters) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyDataAssociatedClusters) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusters) SetClusterId(v string) *ListResourceGroupsResponseBodyDataAssociatedClusters {
	s.ClusterId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusters) SetClusterName(v string) *ListResourceGroupsResponseBodyDataAssociatedClusters {
	s.ClusterName = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusters) SetClusterType(v string) *ListResourceGroupsResponseBodyDataAssociatedClusters {
	s.ClusterType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedClusters) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsResponseBodyDataAssociatedWorkspaces struct {
	// example:
	//
	// 1200827。
	WorkspaceId   *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s ListResourceGroupsResponseBodyDataAssociatedWorkspaces) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyDataAssociatedWorkspaces) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyDataAssociatedWorkspaces) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListResourceGroupsResponseBodyDataAssociatedWorkspaces) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListResourceGroupsResponseBodyDataAssociatedWorkspaces) SetWorkspaceId(v string) *ListResourceGroupsResponseBodyDataAssociatedWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedWorkspaces) SetWorkspaceName(v string) *ListResourceGroupsResponseBodyDataAssociatedWorkspaces {
	s.WorkspaceName = &v
	return s
}

func (s *ListResourceGroupsResponseBodyDataAssociatedWorkspaces) Validate() error {
	return dara.Validate(s)
}
