// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllInstancesResponseBody
	GetCode() *string
	SetData(v []*ListAllInstancesResponseBodyData) *ListAllInstancesResponseBody
	GetData() []*ListAllInstancesResponseBodyData
	SetMaxResults(v int32) *ListAllInstancesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAllInstancesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAllInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAllInstancesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAllInstancesResponseBody
	GetTotal() *int64
}

type ListAllInstancesResponseBody struct {
	// Status code
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data
	Data []*ListAllInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Maximum number of returned items
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Error message:
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the error message of the request.
	//
	// example:
	//
	// instance not exists
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Paging cursor.
	//
	// > If not empty, more data is available.
	//
	// example:
	//
	// c2f78a783f49457caba6bace6f6f79e4
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Total number of query results.
	//
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAllInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllInstancesResponseBody) GetData() []*ListAllInstancesResponseBodyData {
	return s.Data
}

func (s *ListAllInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAllInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAllInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllInstancesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAllInstancesResponseBody) SetCode(v string) *ListAllInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllInstancesResponseBody) SetData(v []*ListAllInstancesResponseBodyData) *ListAllInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListAllInstancesResponseBody) SetMaxResults(v int32) *ListAllInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAllInstancesResponseBody) SetMessage(v string) *ListAllInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllInstancesResponseBody) SetNextToken(v string) *ListAllInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAllInstancesResponseBody) SetRequestId(v string) *ListAllInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllInstancesResponseBody) SetTotal(v int64) *ListAllInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *ListAllInstancesResponseBody) Validate() error {
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

type ListAllInstancesResponseBodyData struct {
	// Agent configuration ID
	//
	// example:
	//
	// test_id
	AgentConfigId *string `json:"agentConfigId,omitempty" xml:"agentConfigId,omitempty"`
	// Agent configuration name
	//
	// example:
	//
	// test_name
	AgentConfigName *string `json:"agentConfigName,omitempty" xml:"agentConfigName,omitempty"`
	// Extension information.
	Attributes []*ListAllInstancesResponseBodyDataAttributes `json:"attributes,omitempty" xml:"attributes,omitempty" type:"Repeated"`
	// Cluster ID
	//
	// example:
	//
	// 3b24a621-acb3-11ef-8c90-00163e1029af
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// zjk_vpc_domain_1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// Image ID.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20250117.vhd
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// Installation level
	//
	// example:
	//
	// Cluster
	InstallLevel *string `json:"installLevel,omitempty" xml:"installLevel,omitempty"`
	// Installation type
	//
	// example:
	//
	// console
	InstallType *string `json:"installType,omitempty" xml:"installType,omitempty"`
	// Instance ID
	//
	// example:
	//
	// i-bp17uabeke9v7n30abm2
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// Instance type
	//
	// example:
	//
	// ecs
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// Milvus version
	//
	// example:
	//
	// 5.10.134-18.al8.x86_64
	KernelVersion *string `json:"kernelVersion,omitempty" xml:"kernelVersion,omitempty"`
	// Management level
	//
	// example:
	//
	// cluster
	ManageLevel *string `json:"manageLevel,omitempty" xml:"manageLevel,omitempty"`
	// Management type
	//
	// example:
	//
	// managed
	ManageType *string `json:"manageType,omitempty" xml:"manageType,omitempty"`
	// Operating system architecture
	//
	// example:
	//
	// x86_64
	OsArch *string `json:"osArch,omitempty" xml:"osArch,omitempty"`
	// Operating system health score
	//
	// example:
	//
	// 100
	OsHealthScore *int32 `json:"osHealthScore,omitempty" xml:"osHealthScore,omitempty"`
	// Operating system name
	//
	// example:
	//
	// alios
	OsName *string `json:"osName,omitempty" xml:"osName,omitempty"`
	// Private IP address.
	//
	// example:
	//
	// 172.21.172.7
	PrivateIp *string `json:"privateIp,omitempty" xml:"privateIp,omitempty"`
	// Public IP address.
	//
	// example:
	//
	// 47.98.215.58
	PublicIp *string `json:"publicIp,omitempty" xml:"publicIp,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// 3b24a621-acb3-11ef-8c90-00163e1029af
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Resource group name
	//
	// example:
	//
	// xxxx
	ResourceGroupName *string `json:"resourceGroupName,omitempty" xml:"resourceGroupName,omitempty"`
	// The running status of the instance. Valid values:
	//
	// - **Running**: The instance is running.
	//
	// - **Offline**: The instance is offline.
	//
	// > An instance in the Offline state indicates that the heartbeat from the edge zone to the SysOM Server has been lost. It does not mean that the corresponding ECS instance is not running.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAllInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllInstancesResponseBodyData) GetAgentConfigId() *string {
	return s.AgentConfigId
}

func (s *ListAllInstancesResponseBodyData) GetAgentConfigName() *string {
	return s.AgentConfigName
}

func (s *ListAllInstancesResponseBodyData) GetAttributes() []*ListAllInstancesResponseBodyDataAttributes {
	return s.Attributes
}

func (s *ListAllInstancesResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAllInstancesResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListAllInstancesResponseBodyData) GetImageId() *string {
	return s.ImageId
}

func (s *ListAllInstancesResponseBodyData) GetInstallLevel() *string {
	return s.InstallLevel
}

func (s *ListAllInstancesResponseBodyData) GetInstallType() *string {
	return s.InstallType
}

func (s *ListAllInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAllInstancesResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListAllInstancesResponseBodyData) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListAllInstancesResponseBodyData) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *ListAllInstancesResponseBodyData) GetManageLevel() *string {
	return s.ManageLevel
}

func (s *ListAllInstancesResponseBodyData) GetManageType() *string {
	return s.ManageType
}

func (s *ListAllInstancesResponseBodyData) GetOsArch() *string {
	return s.OsArch
}

func (s *ListAllInstancesResponseBodyData) GetOsHealthScore() *int32 {
	return s.OsHealthScore
}

func (s *ListAllInstancesResponseBodyData) GetOsName() *string {
	return s.OsName
}

func (s *ListAllInstancesResponseBodyData) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListAllInstancesResponseBodyData) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListAllInstancesResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAllInstancesResponseBodyData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListAllInstancesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAllInstancesResponseBodyData) SetAgentConfigId(v string) *ListAllInstancesResponseBodyData {
	s.AgentConfigId = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetAgentConfigName(v string) *ListAllInstancesResponseBodyData {
	s.AgentConfigName = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetAttributes(v []*ListAllInstancesResponseBodyDataAttributes) *ListAllInstancesResponseBodyData {
	s.Attributes = v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetClusterId(v string) *ListAllInstancesResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetClusterName(v string) *ListAllInstancesResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetImageId(v string) *ListAllInstancesResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetInstallLevel(v string) *ListAllInstancesResponseBodyData {
	s.InstallLevel = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetInstallType(v string) *ListAllInstancesResponseBodyData {
	s.InstallType = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetInstanceId(v string) *ListAllInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetInstanceName(v string) *ListAllInstancesResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetInstanceType(v string) *ListAllInstancesResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetKernelVersion(v string) *ListAllInstancesResponseBodyData {
	s.KernelVersion = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetManageLevel(v string) *ListAllInstancesResponseBodyData {
	s.ManageLevel = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetManageType(v string) *ListAllInstancesResponseBodyData {
	s.ManageType = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetOsArch(v string) *ListAllInstancesResponseBodyData {
	s.OsArch = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetOsHealthScore(v int32) *ListAllInstancesResponseBodyData {
	s.OsHealthScore = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetOsName(v string) *ListAllInstancesResponseBodyData {
	s.OsName = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetPrivateIp(v string) *ListAllInstancesResponseBodyData {
	s.PrivateIp = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetPublicIp(v string) *ListAllInstancesResponseBodyData {
	s.PublicIp = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetResourceGroupId(v string) *ListAllInstancesResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetResourceGroupName(v string) *ListAllInstancesResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) SetStatus(v string) *ListAllInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAllInstancesResponseBodyData) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllInstancesResponseBodyDataAttributes struct {
	// information key
	//
	// example:
	//
	// sysom
	InfoKey *string `json:"infoKey,omitempty" xml:"infoKey,omitempty"`
	// information type
	//
	// example:
	//
	// instance_tag
	InfoType *string `json:"infoType,omitempty" xml:"infoType,omitempty"`
	// information value
	//
	// example:
	//
	// diagnosis
	InfoValue *string `json:"infoValue,omitempty" xml:"infoValue,omitempty"`
}

func (s ListAllInstancesResponseBodyDataAttributes) String() string {
	return dara.Prettify(s)
}

func (s ListAllInstancesResponseBodyDataAttributes) GoString() string {
	return s.String()
}

func (s *ListAllInstancesResponseBodyDataAttributes) GetInfoKey() *string {
	return s.InfoKey
}

func (s *ListAllInstancesResponseBodyDataAttributes) GetInfoType() *string {
	return s.InfoType
}

func (s *ListAllInstancesResponseBodyDataAttributes) GetInfoValue() *string {
	return s.InfoValue
}

func (s *ListAllInstancesResponseBodyDataAttributes) SetInfoKey(v string) *ListAllInstancesResponseBodyDataAttributes {
	s.InfoKey = &v
	return s
}

func (s *ListAllInstancesResponseBodyDataAttributes) SetInfoType(v string) *ListAllInstancesResponseBodyDataAttributes {
	s.InfoType = &v
	return s
}

func (s *ListAllInstancesResponseBodyDataAttributes) SetInfoValue(v string) *ListAllInstancesResponseBodyDataAttributes {
	s.InfoValue = &v
	return s
}

func (s *ListAllInstancesResponseBodyDataAttributes) Validate() error {
	return dara.Validate(s)
}
