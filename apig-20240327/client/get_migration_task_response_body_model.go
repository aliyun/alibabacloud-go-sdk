// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMigrationTaskResponseBody
	GetCode() *string
	SetData(v *GetMigrationTaskResponseBodyData) *GetMigrationTaskResponseBody
	GetData() *GetMigrationTaskResponseBodyData
	SetMessage(v string) *GetMigrationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMigrationTaskResponseBody
	GetRequestId() *string
}

type GetMigrationTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetMigrationTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 42EAF9DB-9082-5F11-8EE1-C2357906DA0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMigrationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetMigrationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMigrationTaskResponseBody) GetData() *GetMigrationTaskResponseBodyData {
	return s.Data
}

func (s *GetMigrationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMigrationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMigrationTaskResponseBody) SetCode(v string) *GetMigrationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetMigrationTaskResponseBody) SetData(v *GetMigrationTaskResponseBodyData) *GetMigrationTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetMigrationTaskResponseBody) SetMessage(v string) *GetMigrationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetMigrationTaskResponseBody) SetRequestId(v string) *GetMigrationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMigrationTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMigrationTaskResponseBodyData struct {
	// The ID of the associated HTTP API.
	//
	// example:
	//
	// api-xxxx
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// The API name. This value may be empty.
	//
	// example:
	//
	// ingress-api
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// The ID of the source ACK cluster.
	//
	// example:
	//
	// c-xxxxxx
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The name of the source cluster. The name is retrieved by looking up the service source based on gatewayId and clusterId. If the lookup fails, a hyphen (-) is returned as a fallback.
	//
	// example:
	//
	// my-cluster
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The namespace of the source cluster. This value is selected during SLB traffic switching.
	//
	// example:
	//
	// default
	ClusterNamespace *string `json:"clusterNamespace,omitempty" xml:"clusterNamespace,omitempty"`
	// The creation time. The value is a UNIX timestamp in seconds. Multiply the value by 1000 on the frontend before formatting.
	//
	// example:
	//
	// 1756262400
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The task description.
	//
	// example:
	//
	// workspace api monitor test
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
	// The name of the destination gateway. If the query fails, a hyphen (-) is returned as a fallback without blocking the return of the main task information.
	//
	// example:
	//
	// my-gateway
	GatewayName *string `json:"gatewayName,omitempty" xml:"gatewayName,omitempty"`
	// The Ingress configuration.
	IngressConfig *GetMigrationTaskResponseBodyDataIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	// The migration type.
	//
	// [_single.resp.200.props.data.migrationType.enum.Nginx Ingress]Nginx Ingress migration.
	//
	// example:
	//
	// Nginx Ingress
	MigrationType *string `json:"migrationType,omitempty" xml:"migrationType,omitempty"`
	// The name of the selected SLB service.
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
	// The traffic switching method.
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
	// The list of SLB virtual server groups.
	VirtualServices []*GetMigrationTaskResponseBodyDataVirtualServices `json:"virtualServices,omitempty" xml:"virtualServices,omitempty" type:"Repeated"`
	// The weight of the APIG-side service in the legacy flow (0-100).
	//
	// example:
	//
	// 50
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s GetMigrationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMigrationTaskResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *GetMigrationTaskResponseBodyData) GetApiName() *string {
	return s.ApiName
}

func (s *GetMigrationTaskResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMigrationTaskResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetMigrationTaskResponseBodyData) GetClusterNamespace() *string {
	return s.ClusterNamespace
}

func (s *GetMigrationTaskResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMigrationTaskResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMigrationTaskResponseBodyData) GetEnvId() *string {
	return s.EnvId
}

func (s *GetMigrationTaskResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetMigrationTaskResponseBodyData) GetGatewayName() *string {
	return s.GatewayName
}

func (s *GetMigrationTaskResponseBodyData) GetIngressConfig() *GetMigrationTaskResponseBodyDataIngressConfig {
	return s.IngressConfig
}

func (s *GetMigrationTaskResponseBodyData) GetMigrationType() *string {
	return s.MigrationType
}

func (s *GetMigrationTaskResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetMigrationTaskResponseBodyData) GetSlbId() *string {
	return s.SlbId
}

func (s *GetMigrationTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMigrationTaskResponseBodyData) GetSwitchType() *string {
	return s.SwitchType
}

func (s *GetMigrationTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMigrationTaskResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetMigrationTaskResponseBodyData) GetVirtualServices() []*GetMigrationTaskResponseBodyDataVirtualServices {
	return s.VirtualServices
}

func (s *GetMigrationTaskResponseBodyData) GetWeight() *int32 {
	return s.Weight
}

func (s *GetMigrationTaskResponseBodyData) SetApiId(v string) *GetMigrationTaskResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetApiName(v string) *GetMigrationTaskResponseBodyData {
	s.ApiName = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetClusterId(v string) *GetMigrationTaskResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetClusterName(v string) *GetMigrationTaskResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetClusterNamespace(v string) *GetMigrationTaskResponseBodyData {
	s.ClusterNamespace = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetCreateTime(v int64) *GetMigrationTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetDescription(v string) *GetMigrationTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetEnvId(v string) *GetMigrationTaskResponseBodyData {
	s.EnvId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetGatewayId(v string) *GetMigrationTaskResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetGatewayName(v string) *GetMigrationTaskResponseBodyData {
	s.GatewayName = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetIngressConfig(v *GetMigrationTaskResponseBodyDataIngressConfig) *GetMigrationTaskResponseBodyData {
	s.IngressConfig = v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetMigrationType(v string) *GetMigrationTaskResponseBodyData {
	s.MigrationType = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetServiceName(v string) *GetMigrationTaskResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetSlbId(v string) *GetMigrationTaskResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetStatus(v string) *GetMigrationTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetSwitchType(v string) *GetMigrationTaskResponseBodyData {
	s.SwitchType = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetTaskId(v string) *GetMigrationTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetUserId(v string) *GetMigrationTaskResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetVirtualServices(v []*GetMigrationTaskResponseBodyDataVirtualServices) *GetMigrationTaskResponseBodyData {
	s.VirtualServices = v
	return s
}

func (s *GetMigrationTaskResponseBodyData) SetWeight(v int32) *GetMigrationTaskResponseBodyData {
	s.Weight = &v
	return s
}

func (s *GetMigrationTaskResponseBodyData) Validate() error {
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

type GetMigrationTaskResponseBodyDataIngressConfig struct {
	// The Ingress class to listen on.
	//
	// example:
	//
	// nginx
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// The namespace in which to listen for Ingress resources.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s GetMigrationTaskResponseBodyDataIngressConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationTaskResponseBodyDataIngressConfig) GoString() string {
	return s.String()
}

func (s *GetMigrationTaskResponseBodyDataIngressConfig) GetIngressClass() *string {
	return s.IngressClass
}

func (s *GetMigrationTaskResponseBodyDataIngressConfig) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *GetMigrationTaskResponseBodyDataIngressConfig) SetIngressClass(v string) *GetMigrationTaskResponseBodyDataIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *GetMigrationTaskResponseBodyDataIngressConfig) SetWatchNamespace(v string) *GetMigrationTaskResponseBodyDataIngressConfig {
	s.WatchNamespace = &v
	return s
}

func (s *GetMigrationTaskResponseBodyDataIngressConfig) Validate() error {
	return dara.Validate(s)
}

type GetMigrationTaskResponseBodyDataVirtualServices struct {
	// The port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol. This parameter is required for CLB.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The backend server group ID.
	//
	// example:
	//
	// rsp-xxxx
	VirtualServiceGroupId *string `json:"virtualServiceGroupId,omitempty" xml:"virtualServiceGroupId,omitempty"`
	// The name.
	//
	// example:
	//
	// 80-tcp
	VirtualServiceGroupName *string `json:"virtualServiceGroupName,omitempty" xml:"virtualServiceGroupName,omitempty"`
}

func (s GetMigrationTaskResponseBodyDataVirtualServices) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationTaskResponseBodyDataVirtualServices) GoString() string {
	return s.String()
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) GetPort() *int32 {
	return s.Port
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) GetProtocol() *string {
	return s.Protocol
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) GetVirtualServiceGroupId() *string {
	return s.VirtualServiceGroupId
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) GetVirtualServiceGroupName() *string {
	return s.VirtualServiceGroupName
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) SetPort(v int32) *GetMigrationTaskResponseBodyDataVirtualServices {
	s.Port = &v
	return s
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) SetProtocol(v string) *GetMigrationTaskResponseBodyDataVirtualServices {
	s.Protocol = &v
	return s
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) SetVirtualServiceGroupId(v string) *GetMigrationTaskResponseBodyDataVirtualServices {
	s.VirtualServiceGroupId = &v
	return s
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) SetVirtualServiceGroupName(v string) *GetMigrationTaskResponseBodyDataVirtualServices {
	s.VirtualServiceGroupName = &v
	return s
}

func (s *GetMigrationTaskResponseBodyDataVirtualServices) Validate() error {
	return dara.Validate(s)
}
