// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sIngressRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListK8sIngressRulesResponseBody
	GetCode() *int32
	SetData(v []*ListK8sIngressRulesResponseBodyData) *ListK8sIngressRulesResponseBody
	GetData() []*ListK8sIngressRulesResponseBodyData
	SetMessage(v string) *ListK8sIngressRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListK8sIngressRulesResponseBody
	GetRequestId() *string
}

type ListK8sIngressRulesResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data []*ListK8sIngressRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5C1C9DE7-88FF-4B56-A47B-3DBBCEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListK8sIngressRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListK8sIngressRulesResponseBody) GetData() []*ListK8sIngressRulesResponseBodyData {
	return s.Data
}

func (s *ListK8sIngressRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListK8sIngressRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8sIngressRulesResponseBody) SetCode(v int32) *ListK8sIngressRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListK8sIngressRulesResponseBody) SetData(v []*ListK8sIngressRulesResponseBodyData) *ListK8sIngressRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListK8sIngressRulesResponseBody) SetMessage(v string) *ListK8sIngressRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListK8sIngressRulesResponseBody) SetRequestId(v string) *ListK8sIngressRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8sIngressRulesResponseBody) Validate() error {
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

type ListK8sIngressRulesResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// 5b2b4ab4-efbc-4a81-9c45-a5942881****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// my-dev-cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The Ingresses.
	IngressConfs []*ListK8sIngressRulesResponseBodyDataIngressConfs `json:"IngressConfs,omitempty" xml:"IngressConfs,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListK8sIngressRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sIngressRulesResponseBodyData) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListK8sIngressRulesResponseBodyData) GetIngressConfs() []*ListK8sIngressRulesResponseBodyDataIngressConfs {
	return s.IngressConfs
}

func (s *ListK8sIngressRulesResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListK8sIngressRulesResponseBodyData) SetClusterId(v string) *ListK8sIngressRulesResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyData) SetClusterName(v string) *ListK8sIngressRulesResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyData) SetIngressConfs(v []*ListK8sIngressRulesResponseBodyDataIngressConfs) *ListK8sIngressRulesResponseBodyData {
	s.IngressConfs = v
	return s
}

func (s *ListK8sIngressRulesResponseBodyData) SetRegionId(v string) *ListK8sIngressRulesResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyData) Validate() error {
	if s.IngressConfs != nil {
		for _, item := range s.IngressConfs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sIngressRulesResponseBodyDataIngressConfs struct {
	// The ID of the ALB instance.
	//
	// example:
	//
	// alb-5sde0tq62r********
	AlbId *string `json:"AlbId,omitempty" xml:"AlbId,omitempty"`
	// The annotations.
	//
	// example:
	//
	// {"test-annotation":"test-annotation-value"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The time when the Ingress was created.
	//
	// example:
	//
	// 2021-04-27 20:16:52
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The monitoring URL of the Ingress.
	//
	// example:
	//
	// http://grafana.console.aliyun.com/d/10xxxx/ingress
	DashboardUrl *string `json:"DashboardUrl,omitempty" xml:"DashboardUrl,omitempty"`
	// The IP address of the Ingress.
	//
	// example:
	//
	// 47.11x.xx.xx
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The Ingress type. Valid values:
	//
	// 	- **NginxIngress**: NGINX Ingress controller
	//
	// 	- **AlbIngress**: ALB Ingress controller
	//
	// Default value: NginxIngress.
	//
	// example:
	//
	// NginxIngress
	IngressType *string `json:"IngressType,omitempty" xml:"IngressType,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"test-label": "test-labels"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The ID of the MSE gateway.
	//
	// example:
	//
	// gw-xxxxxxxx
	MseGatewayId *string `json:"MseGatewayId,omitempty" xml:"MseGatewayId,omitempty"`
	// The name of the MSE gateway.
	//
	// example:
	//
	// gw-test-name
	MseGatewayName *string `json:"MseGatewayName,omitempty" xml:"MseGatewayName,omitempty"`
	// The Ingress name.
	//
	// example:
	//
	// my-ingress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Kubernetes namespace to which the Ingress belongs.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The URL used for basic monitoring of the open source version.
	//
	// example:
	//
	// https://g.console.aliyun.com/d/xxxxxx/nginx-ingress-dashboard-official
	OfficalBasicUrl *string `json:"OfficalBasicUrl,omitempty" xml:"OfficalBasicUrl,omitempty"`
	// The URL used for request performance monitoring of the open source version.
	//
	// example:
	//
	// https://g.console.aliyun.com/d/xxxxxx/request-handling-performance-official
	OfficalRequestUrl *string `json:"OfficalRequestUrl,omitempty" xml:"OfficalRequestUrl,omitempty"`
	// The routing rules.
	Rules []*ListK8sIngressRulesResponseBodyDataIngressConfsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Indicates whether SSL redirection is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SslRedirect *bool `json:"SslRedirect,omitempty" xml:"SslRedirect,omitempty"`
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfs) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfs) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetAlbId() *string {
	return s.AlbId
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetAnnotations() *string {
	return s.Annotations
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetDashboardUrl() *string {
	return s.DashboardUrl
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetIngressType() *string {
	return s.IngressType
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetLabels() *string {
	return s.Labels
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetMseGatewayId() *string {
	return s.MseGatewayId
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetMseGatewayName() *string {
	return s.MseGatewayName
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetName() *string {
	return s.Name
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetOfficalBasicUrl() *string {
	return s.OfficalBasicUrl
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetOfficalRequestUrl() *string {
	return s.OfficalRequestUrl
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetRules() []*ListK8sIngressRulesResponseBodyDataIngressConfsRules {
	return s.Rules
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) GetSslRedirect() *bool {
	return s.SslRedirect
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetAlbId(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.AlbId = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetAnnotations(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.Annotations = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetCreationTime(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.CreationTime = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetDashboardUrl(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.DashboardUrl = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetEndpoint(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.Endpoint = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetIngressType(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.IngressType = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetLabels(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.Labels = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetMseGatewayId(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.MseGatewayId = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetMseGatewayName(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.MseGatewayName = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetName(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.Name = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetNamespace(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.Namespace = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetOfficalBasicUrl(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.OfficalBasicUrl = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetOfficalRequestUrl(v string) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.OfficalRequestUrl = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetRules(v []*ListK8sIngressRulesResponseBodyDataIngressConfsRules) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.Rules = v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) SetSslRedirect(v bool) *ListK8sIngressRulesResponseBodyDataIngressConfs {
	s.SslRedirect = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfs) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sIngressRulesResponseBodyDataIngressConfsRules struct {
	// Indicates whether TLS is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableTls *bool `json:"EnableTls,omitempty" xml:"EnableTls,omitempty"`
	// The domain name to be accessed.
	//
	// example:
	//
	// example.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The paths to be accessed.
	Paths []*ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The name of the Secret that stores the Transport Layer Security (TLS) certificate.
	//
	// example:
	//
	// my-secret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfsRules) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfsRules) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) GetEnableTls() *bool {
	return s.EnableTls
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) GetHost() *string {
	return s.Host
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) GetPaths() []*ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	return s.Paths
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) GetSecretName() *string {
	return s.SecretName
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) SetEnableTls(v bool) *ListK8sIngressRulesResponseBodyDataIngressConfsRules {
	s.EnableTls = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) SetHost(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRules {
	s.Host = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) SetPaths(v []*ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) *ListK8sIngressRulesResponseBodyDataIngressConfsRules {
	s.Paths = v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) SetSecretName(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRules {
	s.SecretName = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRules) Validate() error {
	if s.Paths != nil {
		for _, item := range s.Paths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths struct {
	// The ID of the EDAS application.
	//
	// example:
	//
	// 43d30ba5-c568-460c-8080-d447ed1a****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the EDAS application.
	//
	// example:
	//
	// my-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The configurations of the backend Service.
	Backend *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend `json:"Backend,omitempty" xml:"Backend,omitempty" type:"Struct"`
	// The collection rate that is set based on the trace query feature. You can add a trace ID to a gateway to use the trace query feature of EDAS.
	//
	// example:
	//
	// 100
	CollectRate *int32 `json:"CollectRate,omitempty" xml:"CollectRate,omitempty"`
	// The path to be accessed.
	//
	// example:
	//
	// /foo/bar
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path type that determines how a path is matched.
	//
	// 	- ImplementationSpecific (default)
	//
	// 	- Exact
	//
	// 	- Prefix
	//
	// example:
	//
	// ImplementationSpecific
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// The state of the Ingress. Valid values:
	//
	// 	- **Normal**: The Ingress works as expected.
	//
	// 	- **ServiceNotFound**: The backend Service does not exist.
	//
	// 	- **InvalidServicePort**: The Service port is invalid.
	//
	// 	- **NotManagedService**: The Service is not managed by EDAS.
	//
	// 	- **Unknown**: An unknown error occurred.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GetAppId() *string {
	return s.AppId
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GetAppName() *string {
	return s.AppName
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GetBackend() *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend {
	return s.Backend
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GetCollectRate() *int32 {
	return s.CollectRate
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GetPath() *string {
	return s.Path
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GetPathType() *string {
	return s.PathType
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) GetStatus() *string {
	return s.Status
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) SetAppId(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	s.AppId = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) SetAppName(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	s.AppName = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) SetBackend(v *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	s.Backend = v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) SetCollectRate(v int32) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	s.CollectRate = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) SetPath(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	s.Path = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) SetPathType(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	s.PathType = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) SetStatus(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths {
	s.Status = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPaths) Validate() error {
	if s.Backend != nil {
		if err := s.Backend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend struct {
	// The name of the backend Service.
	//
	// example:
	//
	// http-service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The port of the backend Service.
	//
	// example:
	//
	// 8080
	ServicePort *string `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) String() string {
	return dara.Prettify(s)
}

func (s ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) GoString() string {
	return s.String()
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) GetServicePort() *string {
	return s.ServicePort
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) SetServiceName(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend {
	s.ServiceName = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) SetServicePort(v string) *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend {
	s.ServicePort = &v
	return s
}

func (s *ListK8sIngressRulesResponseBodyDataIngressConfsRulesPathsBackend) Validate() error {
	return dara.Validate(s)
}
