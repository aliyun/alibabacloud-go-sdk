// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEcsSlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *BindEcsSlbRequest
	GetAppId() *string
	SetDeployGroupId(v string) *BindEcsSlbRequest
	GetDeployGroupId() *string
	SetListenerHealthCheckUrl(v string) *BindEcsSlbRequest
	GetListenerHealthCheckUrl() *string
	SetListenerPort(v int32) *BindEcsSlbRequest
	GetListenerPort() *int32
	SetListenerProtocol(v string) *BindEcsSlbRequest
	GetListenerProtocol() *string
	SetSlbId(v string) *BindEcsSlbRequest
	GetSlbId() *string
	SetVForwardingUrlRule(v string) *BindEcsSlbRequest
	GetVForwardingUrlRule() *string
	SetVServerGroupId(v string) *BindEcsSlbRequest
	GetVServerGroupId() *string
	SetVServerGroupName(v string) *BindEcsSlbRequest
	GetVServerGroupName() *string
}

type BindEcsSlbRequest struct {
	// The ID of the application. You can query the application ID by calling the ListApplication operation. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 93fdd228-*****-ed2ae98de18d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance group whose application you want to bind. You can call the ListDeployGroup operation to query the group ID. For more information, see [ListDeployGroup](https://help.aliyun.com/document_detail/62077.html).
	//
	// example:
	//
	// 577f4c50-16ee-43d8-****-****
	DeployGroupId *string `json:"DeployGroupId,omitempty" xml:"DeployGroupId,omitempty"`
	// The health check URL.
	//
	// example:
	//
	// /_ehc.html
	ListenerHealthCheckUrl *string `json:"ListenerHealthCheckUrl,omitempty" xml:"ListenerHealthCheckUrl,omitempty"`
	// The listener port for the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol for the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-uf6j54m3a****cj01jx8
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The forwarding rule of the SLB listener.
	//
	// example:
	//
	// example.com/forwarding
	VForwardingUrlRule *string `json:"VForwardingUrlRule,omitempty" xml:"VForwardingUrlRule,omitempty"`
	// The ID of the vServer group for the SLB instance.
	//
	// example:
	//
	// rsp-2ze****6l*****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The name of the vServer group.
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s BindEcsSlbRequest) String() string {
	return dara.Prettify(s)
}

func (s BindEcsSlbRequest) GoString() string {
	return s.String()
}

func (s *BindEcsSlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *BindEcsSlbRequest) GetDeployGroupId() *string {
	return s.DeployGroupId
}

func (s *BindEcsSlbRequest) GetListenerHealthCheckUrl() *string {
	return s.ListenerHealthCheckUrl
}

func (s *BindEcsSlbRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *BindEcsSlbRequest) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *BindEcsSlbRequest) GetSlbId() *string {
	return s.SlbId
}

func (s *BindEcsSlbRequest) GetVForwardingUrlRule() *string {
	return s.VForwardingUrlRule
}

func (s *BindEcsSlbRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *BindEcsSlbRequest) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *BindEcsSlbRequest) SetAppId(v string) *BindEcsSlbRequest {
	s.AppId = &v
	return s
}

func (s *BindEcsSlbRequest) SetDeployGroupId(v string) *BindEcsSlbRequest {
	s.DeployGroupId = &v
	return s
}

func (s *BindEcsSlbRequest) SetListenerHealthCheckUrl(v string) *BindEcsSlbRequest {
	s.ListenerHealthCheckUrl = &v
	return s
}

func (s *BindEcsSlbRequest) SetListenerPort(v int32) *BindEcsSlbRequest {
	s.ListenerPort = &v
	return s
}

func (s *BindEcsSlbRequest) SetListenerProtocol(v string) *BindEcsSlbRequest {
	s.ListenerProtocol = &v
	return s
}

func (s *BindEcsSlbRequest) SetSlbId(v string) *BindEcsSlbRequest {
	s.SlbId = &v
	return s
}

func (s *BindEcsSlbRequest) SetVForwardingUrlRule(v string) *BindEcsSlbRequest {
	s.VForwardingUrlRule = &v
	return s
}

func (s *BindEcsSlbRequest) SetVServerGroupId(v string) *BindEcsSlbRequest {
	s.VServerGroupId = &v
	return s
}

func (s *BindEcsSlbRequest) SetVServerGroupName(v string) *BindEcsSlbRequest {
	s.VServerGroupName = &v
	return s
}

func (s *BindEcsSlbRequest) Validate() error {
	return dara.Validate(s)
}
