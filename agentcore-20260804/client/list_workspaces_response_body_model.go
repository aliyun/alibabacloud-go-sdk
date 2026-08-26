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
	SetItems(v []*ListWorkspacesResponseBodyItems) *ListWorkspacesResponseBody
	GetItems() []*ListWorkspacesResponseBodyItems
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
}

type ListWorkspacesResponseBody struct {
	// The business status code.
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
	// The list of workspaces.
	Items []*ListWorkspacesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page used for this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. This value is empty if no more pages exist.
	//
	// example:
	//
	// d29ya3NwYWNlLW9mZnNldDo0MA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of workspaces that match the query conditions.
	//
	// example:
	//
	// 42
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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

func (s *ListWorkspacesResponseBody) GetItems() []*ListWorkspacesResponseBodyItems {
	return s.Items
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

func (s *ListWorkspacesResponseBody) SetCode(v string) *ListWorkspacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetHttpStatusCode(v int32) *ListWorkspacesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetItems(v []*ListWorkspacesResponseBodyItems) *ListWorkspacesResponseBody {
	s.Items = v
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

func (s *ListWorkspacesResponseBody) Validate() error {
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

type ListWorkspacesResponseBodyItems struct {
	// The creation time of the workspace.
	//
	// example:
	//
	// 2026-08-06T03:56:56Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// production-agents
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration of the workspace.
	NetworkConfiguration *ListWorkspacesResponseBodyItemsNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The region ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The workspace status. Valid values:
	//
	// - Initializing
	//
	// - Initialized
	//
	// - Deleting
	//
	// - Deleted.
	//
	// example:
	//
	// Initialized
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the tenant to which the workspace belongs.
	//
	// example:
	//
	// tenant-123456
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListWorkspacesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWorkspacesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListWorkspacesResponseBodyItems) GetNetworkConfiguration() *ListWorkspacesResponseBodyItemsNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *ListWorkspacesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWorkspacesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListWorkspacesResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkspacesResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListWorkspacesResponseBodyItems) SetCreateTime(v string) *ListWorkspacesResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetName(v string) *ListWorkspacesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetNetworkConfiguration(v *ListWorkspacesResponseBodyItemsNetworkConfiguration) *ListWorkspacesResponseBodyItems {
	s.NetworkConfiguration = v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetRegionId(v string) *ListWorkspacesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetStatus(v string) *ListWorkspacesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetTenantId(v string) *ListWorkspacesResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) SetWorkspaceId(v string) *ListWorkspacesResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItems) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspacesResponseBodyItemsNetworkConfiguration struct {
	// The VPC network configuration of the user.
	Vpc *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s ListWorkspacesResponseBodyItemsNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyItemsNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfiguration) GetVpc() *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc {
	return s.Vpc
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfiguration) SetVpc(v *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) *ListWorkspacesResponseBodyItemsNetworkConfiguration {
	s.Vpc = v
	return s
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfiguration) Validate() error {
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkspacesResponseBodyItemsNetworkConfigurationVpc struct {
	// Indicates whether the VPC network is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The ID of the user VPC.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) String() string {
	return dara.Prettify(s)
}

func (s ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) SetEnabled(v bool) *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc {
	s.Enabled = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) SetVSwitchIds(v []*string) *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc {
	s.VSwitchIds = v
	return s
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) SetVpcId(v string) *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc {
	s.VpcId = &v
	return s
}

func (s *ListWorkspacesResponseBodyItemsNetworkConfigurationVpc) Validate() error {
	return dara.Validate(s)
}
