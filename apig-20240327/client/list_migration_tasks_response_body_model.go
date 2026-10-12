// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMigrationTasksResponseBody
	GetCode() *string
	SetData(v *ListMigrationTasksResponseBodyData) *ListMigrationTasksResponseBody
	GetData() *ListMigrationTasksResponseBodyData
	SetMessage(v string) *ListMigrationTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMigrationTasksResponseBody
	GetRequestId() *string
}

type ListMigrationTasksResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListMigrationTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 019FB5FB-615B-52AB-A92F-D40A3193DA96
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMigrationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMigrationTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMigrationTasksResponseBody) GetData() *ListMigrationTasksResponseBodyData {
	return s.Data
}

func (s *ListMigrationTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMigrationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMigrationTasksResponseBody) SetCode(v string) *ListMigrationTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListMigrationTasksResponseBody) SetData(v *ListMigrationTasksResponseBodyData) *ListMigrationTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListMigrationTasksResponseBody) SetMessage(v string) *ListMigrationTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListMigrationTasksResponseBody) SetRequestId(v string) *ListMigrationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMigrationTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMigrationTasksResponseBodyData struct {
	// The list of migration tasks.
	Items []*ListMigrationTasksResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 25
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListMigrationTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMigrationTasksResponseBodyData) GetItems() []*ListMigrationTasksResponseBodyDataItems {
	return s.Items
}

func (s *ListMigrationTasksResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMigrationTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMigrationTasksResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListMigrationTasksResponseBodyData) SetItems(v []*ListMigrationTasksResponseBodyDataItems) *ListMigrationTasksResponseBodyData {
	s.Items = v
	return s
}

func (s *ListMigrationTasksResponseBodyData) SetPageNumber(v int32) *ListMigrationTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationTasksResponseBodyData) SetPageSize(v int32) *ListMigrationTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMigrationTasksResponseBodyData) SetTotalSize(v int32) *ListMigrationTasksResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListMigrationTasksResponseBodyData) Validate() error {
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

type ListMigrationTasksResponseBodyDataItems struct {
	// The associated HTTP API.
	//
	// example:
	//
	// api-xxxx
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// The associated HTTP API.
	//
	// example:
	//
	// api-name
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// The source ACK cluster. A hyphen (-) is displayed if the cluster name fails to be queried.
	//
	// example:
	//
	// c-xxxxxx
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// vpc_hz_domain_1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The namespace for SLB traffic switching.
	//
	// example:
	//
	// default
	ClusterNamespace *string `json:"clusterNamespace,omitempty" xml:"clusterNamespace,omitempty"`
	// The creation time, in Unix seconds.
	//
	// example:
	//
	// 1756262400
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The task description.
	//
	// example:
	//
	// Migration test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the destination environment.
	//
	// example:
	//
	// env-xxxx
	EnvId *string `json:"envId,omitempty" xml:"envId,omitempty"`
	// The ID of the destination gateway.
	//
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The name of the gateway to which the rule belongs.
	//
	// example:
	//
	// prod-gateway
	GatewayName *string `json:"gatewayName,omitempty" xml:"gatewayName,omitempty"`
	// The monitored Ingress Class configuration.
	IngressConfig *ListMigrationTasksResponseBodyDataItemsIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	// The migration type. Currently, only Nginx Ingress is supported.
	//
	// example:
	//
	// Nginx Ingress
	MigrationType *string `json:"migrationType,omitempty" xml:"migrationType,omitempty"`
	// The name of the SLB service.
	//
	// example:
	//
	// nginx-ingress-lb
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The ID of the SLB instance used for traffic switching.
	//
	// example:
	//
	// lb-bp1xxxx
	SlbId *string `json:"slbId,omitempty" xml:"slbId,omitempty"`
	// The task status.
	//
	// example:
	//
	// FlowSwitch
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The traffic switching method. Valid values: DNS and SLB.
	//
	// example:
	//
	// SLB
	SwitchType *string `json:"switchType,omitempty" xml:"switchType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// mt-xxxxxxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The ID of the user to whom the task belongs.
	//
	// example:
	//
	// 1234567890
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The list of virtual server groups.
	VirtualServices []*ListMigrationTasksResponseBodyDataItemsVirtualServices `json:"virtualServices,omitempty" xml:"virtualServices,omitempty" type:"Repeated"`
	// The weight of the service on the API gateway side in the legacy workflow. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ListMigrationTasksResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTasksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListMigrationTasksResponseBodyDataItems) GetApiId() *string {
	return s.ApiId
}

func (s *ListMigrationTasksResponseBodyDataItems) GetApiName() *string {
	return s.ApiName
}

func (s *ListMigrationTasksResponseBodyDataItems) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListMigrationTasksResponseBodyDataItems) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListMigrationTasksResponseBodyDataItems) GetClusterNamespace() *string {
	return s.ClusterNamespace
}

func (s *ListMigrationTasksResponseBodyDataItems) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMigrationTasksResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListMigrationTasksResponseBodyDataItems) GetEnvId() *string {
	return s.EnvId
}

func (s *ListMigrationTasksResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListMigrationTasksResponseBodyDataItems) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListMigrationTasksResponseBodyDataItems) GetIngressConfig() *ListMigrationTasksResponseBodyDataItemsIngressConfig {
	return s.IngressConfig
}

func (s *ListMigrationTasksResponseBodyDataItems) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListMigrationTasksResponseBodyDataItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListMigrationTasksResponseBodyDataItems) GetSlbId() *string {
	return s.SlbId
}

func (s *ListMigrationTasksResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListMigrationTasksResponseBodyDataItems) GetSwitchType() *string {
	return s.SwitchType
}

func (s *ListMigrationTasksResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListMigrationTasksResponseBodyDataItems) GetUserId() *string {
	return s.UserId
}

func (s *ListMigrationTasksResponseBodyDataItems) GetVirtualServices() []*ListMigrationTasksResponseBodyDataItemsVirtualServices {
	return s.VirtualServices
}

func (s *ListMigrationTasksResponseBodyDataItems) GetWeight() *int32 {
	return s.Weight
}

func (s *ListMigrationTasksResponseBodyDataItems) SetApiId(v string) *ListMigrationTasksResponseBodyDataItems {
	s.ApiId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetApiName(v string) *ListMigrationTasksResponseBodyDataItems {
	s.ApiName = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetClusterId(v string) *ListMigrationTasksResponseBodyDataItems {
	s.ClusterId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetClusterName(v string) *ListMigrationTasksResponseBodyDataItems {
	s.ClusterName = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetClusterNamespace(v string) *ListMigrationTasksResponseBodyDataItems {
	s.ClusterNamespace = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetCreateTime(v int64) *ListMigrationTasksResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetDescription(v string) *ListMigrationTasksResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetEnvId(v string) *ListMigrationTasksResponseBodyDataItems {
	s.EnvId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetGatewayId(v string) *ListMigrationTasksResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetGatewayName(v string) *ListMigrationTasksResponseBodyDataItems {
	s.GatewayName = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetIngressConfig(v *ListMigrationTasksResponseBodyDataItemsIngressConfig) *ListMigrationTasksResponseBodyDataItems {
	s.IngressConfig = v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetMigrationType(v string) *ListMigrationTasksResponseBodyDataItems {
	s.MigrationType = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetServiceName(v string) *ListMigrationTasksResponseBodyDataItems {
	s.ServiceName = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetSlbId(v string) *ListMigrationTasksResponseBodyDataItems {
	s.SlbId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetStatus(v string) *ListMigrationTasksResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetSwitchType(v string) *ListMigrationTasksResponseBodyDataItems {
	s.SwitchType = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetTaskId(v string) *ListMigrationTasksResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetUserId(v string) *ListMigrationTasksResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetVirtualServices(v []*ListMigrationTasksResponseBodyDataItemsVirtualServices) *ListMigrationTasksResponseBodyDataItems {
	s.VirtualServices = v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) SetWeight(v int32) *ListMigrationTasksResponseBodyDataItems {
	s.Weight = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItems) Validate() error {
	if s.IngressConfig != nil {
		if err := s.IngressConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VirtualServices != nil {
		for _, item := range s.VirtualServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMigrationTasksResponseBodyDataItemsIngressConfig struct {
	// The monitored Ingress Class.
	//
	// example:
	//
	// nginx
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// The monitored namespace.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s ListMigrationTasksResponseBodyDataItemsIngressConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTasksResponseBodyDataItemsIngressConfig) GoString() string {
	return s.String()
}

func (s *ListMigrationTasksResponseBodyDataItemsIngressConfig) GetIngressClass() *string {
	return s.IngressClass
}

func (s *ListMigrationTasksResponseBodyDataItemsIngressConfig) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *ListMigrationTasksResponseBodyDataItemsIngressConfig) SetIngressClass(v string) *ListMigrationTasksResponseBodyDataItemsIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItemsIngressConfig) SetWatchNamespace(v string) *ListMigrationTasksResponseBodyDataItemsIngressConfig {
	s.WatchNamespace = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItemsIngressConfig) Validate() error {
	return dara.Validate(s)
}

type ListMigrationTasksResponseBodyDataItemsVirtualServices struct {
	// The listener port. Valid values: 1 to 65535. This parameter is required in Step 3 of the frontend form.
	//
	// example:
	//
	// 443
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The listener protocol. This parameter is required for CLB migration. During MSE migration validation, CLB requests without a protocol are directly rejected. See migrationtask.go#L95-L97. This parameter can be left empty for NLB scenarios.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The virtual server group ID. For CLB, this is the VServerGroupId. For NLB, this is the ServerGroupId. This parameter is required in Step 3 of the frontend form.
	//
	// example:
	//
	// rsp-xxxx
	VirtualServiceGroupId *string `json:"virtualServiceGroupId,omitempty" xml:"virtualServiceGroupId,omitempty"`
	// The name of the virtual server group.
	//
	// example:
	//
	// 80-tcp
	VirtualServiceGroupName *string `json:"virtualServiceGroupName,omitempty" xml:"virtualServiceGroupName,omitempty"`
}

func (s ListMigrationTasksResponseBodyDataItemsVirtualServices) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationTasksResponseBodyDataItemsVirtualServices) GoString() string {
	return s.String()
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) GetPort() *int32 {
	return s.Port
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) GetProtocol() *string {
	return s.Protocol
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) GetVirtualServiceGroupId() *string {
	return s.VirtualServiceGroupId
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) GetVirtualServiceGroupName() *string {
	return s.VirtualServiceGroupName
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) SetPort(v int32) *ListMigrationTasksResponseBodyDataItemsVirtualServices {
	s.Port = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) SetProtocol(v string) *ListMigrationTasksResponseBodyDataItemsVirtualServices {
	s.Protocol = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) SetVirtualServiceGroupId(v string) *ListMigrationTasksResponseBodyDataItemsVirtualServices {
	s.VirtualServiceGroupId = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) SetVirtualServiceGroupName(v string) *ListMigrationTasksResponseBodyDataItemsVirtualServices {
	s.VirtualServiceGroupName = &v
	return s
}

func (s *ListMigrationTasksResponseBodyDataItemsVirtualServices) Validate() error {
	return dara.Validate(s)
}
