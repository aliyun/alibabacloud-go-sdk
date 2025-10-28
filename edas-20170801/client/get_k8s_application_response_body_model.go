// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplcation(v *GetK8sApplicationResponseBodyApplcation) *GetK8sApplicationResponseBody
	GetApplcation() *GetK8sApplicationResponseBodyApplcation
	SetCode(v int32) *GetK8sApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *GetK8sApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetK8sApplicationResponseBody
	GetRequestId() *string
}

type GetK8sApplicationResponseBody struct {
	// The information about the application.
	Applcation *GetK8sApplicationResponseBodyApplcation `json:"Applcation,omitempty" xml:"Applcation,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1053-08e4-47a5-b2ab-5c0323de7b5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetK8sApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBody) GetApplcation() *GetK8sApplicationResponseBodyApplcation {
	return s.Applcation
}

func (s *GetK8sApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetK8sApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetK8sApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetK8sApplicationResponseBody) SetApplcation(v *GetK8sApplicationResponseBodyApplcation) *GetK8sApplicationResponseBody {
	s.Applcation = v
	return s
}

func (s *GetK8sApplicationResponseBody) SetCode(v int32) *GetK8sApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetK8sApplicationResponseBody) SetMessage(v string) *GetK8sApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetK8sApplicationResponseBody) SetRequestId(v string) *GetK8sApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetK8sApplicationResponseBody) Validate() error {
	if s.Applcation != nil {
		if err := s.Applcation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetK8sApplicationResponseBodyApplcation struct {
	// The basic information about the application.
	App *GetK8sApplicationResponseBodyApplcationApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// example:
	//
	// a5281053-****-47a5-b2ab-5c0323de****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The configurations.
	Conf *GetK8sApplicationResponseBodyApplcationConf `json:"Conf,omitempty" xml:"Conf,omitempty" type:"Struct"`
	// The information about the instance group in which the application is deployed.
	DeployGroups *GetK8sApplicationResponseBodyApplcationDeployGroups `json:"DeployGroups,omitempty" xml:"DeployGroups,omitempty" type:"Struct"`
	// The information about the image.
	ImageInfo *GetK8sApplicationResponseBodyApplcationImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	// The information about the latest version.
	LatestVersion *GetK8sApplicationResponseBodyApplcationLatestVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty" type:"Struct"`
}

func (s GetK8sApplicationResponseBodyApplcation) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcation) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcation) GetApp() *GetK8sApplicationResponseBodyApplcationApp {
	return s.App
}

func (s *GetK8sApplicationResponseBodyApplcation) GetAppId() *string {
	return s.AppId
}

func (s *GetK8sApplicationResponseBodyApplcation) GetConf() *GetK8sApplicationResponseBodyApplcationConf {
	return s.Conf
}

func (s *GetK8sApplicationResponseBodyApplcation) GetDeployGroups() *GetK8sApplicationResponseBodyApplcationDeployGroups {
	return s.DeployGroups
}

func (s *GetK8sApplicationResponseBodyApplcation) GetImageInfo() *GetK8sApplicationResponseBodyApplcationImageInfo {
	return s.ImageInfo
}

func (s *GetK8sApplicationResponseBodyApplcation) GetLatestVersion() *GetK8sApplicationResponseBodyApplcationLatestVersion {
	return s.LatestVersion
}

func (s *GetK8sApplicationResponseBodyApplcation) SetApp(v *GetK8sApplicationResponseBodyApplcationApp) *GetK8sApplicationResponseBodyApplcation {
	s.App = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetAppId(v string) *GetK8sApplicationResponseBodyApplcation {
	s.AppId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetConf(v *GetK8sApplicationResponseBodyApplcationConf) *GetK8sApplicationResponseBodyApplcation {
	s.Conf = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetDeployGroups(v *GetK8sApplicationResponseBodyApplcationDeployGroups) *GetK8sApplicationResponseBodyApplcation {
	s.DeployGroups = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetImageInfo(v *GetK8sApplicationResponseBodyApplcationImageInfo) *GetK8sApplicationResponseBodyApplcation {
	s.ImageInfo = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) SetLatestVersion(v *GetK8sApplicationResponseBodyApplcationLatestVersion) *GetK8sApplicationResponseBodyApplcation {
	s.LatestVersion = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcation) Validate() error {
	if s.App != nil {
		if err := s.App.Validate(); err != nil {
			return err
		}
	}
	if s.Conf != nil {
		if err := s.Conf.Validate(); err != nil {
			return err
		}
	}
	if s.DeployGroups != nil {
		if err := s.DeployGroups.Validate(); err != nil {
			return err
		}
	}
	if s.ImageInfo != nil {
		if err := s.ImageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.LatestVersion != nil {
		if err := s.LatestVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetK8sApplicationResponseBodyApplcationApp struct {
	// The annotation of an application pod.
	//
	// example:
	//
	// {"test-annokey":"test-annovalue"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// example:
	//
	// 00ee517d-dd7d-4d4e-****-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The type of the application.
	//
	// example:
	//
	// War
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The build package number of Enterprise Distributed Application Service (EDAS) Container.
	//
	// example:
	//
	// 57
	BuildpackId *int32 `json:"BuildpackId,omitempty" xml:"BuildpackId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c37aec2a-bcca-4ec1-****-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The startup command.
	//
	// example:
	//
	// ls
	Cmd *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	// The list of commands.
	CmdArgs *GetK8sApplicationResponseBodyApplcationAppCmdArgs `json:"CmdArgs,omitempty" xml:"CmdArgs,omitempty" type:"Struct"`
	// The ID of the cluster to which the container belongs.
	//
	// example:
	//
	// c383bc813c1974e****451b50c0c8****
	CsClusterId *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty"`
	// The deployment type of the application. Example: Image.
	//
	// example:
	//
	// Image
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The application type. Valid values:
	//
	// 	- General: native Java application
	//
	// 	- Pandora: Pandora application
	//
	// 	- Multilingual: multilingual application
	//
	// example:
	//
	// General
	DevelopType *string `json:"DevelopType,omitempty" xml:"DevelopType,omitempty"`
	// The version of EDAS Container.
	//
	// example:
	//
	// 3.60.0
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// Indicates whether the Empty List Protection feature is enabled for the application.
	//
	// example:
	//
	// true
	EnableEmptyPushReject *bool `json:"EnableEmptyPushReject,omitempty" xml:"EnableEmptyPushReject,omitempty"`
	// Indicates whether the Graceful Release feature is enabled for the application.
	//
	// example:
	//
	// true
	EnableLosslessRule *bool `json:"EnableLosslessRule,omitempty" xml:"EnableLosslessRule,omitempty"`
	// The list of environment variables.
	EnvList *GetK8sApplicationResponseBodyApplcationAppEnvList `json:"EnvList,omitempty" xml:"EnvList,omitempty" type:"Struct"`
	// The feature annotations. Possible values:
	//
	// 	- base.combination.edas: enables EDAS integrated management solution.
	//
	// 	- base.combination.arms: enables ARMS monitoring.
	//
	// 	- base.combination.mse: enables MSE microservices governance.
	//
	// 	- base.combination.none: enables lifecycle management.
	//
	// example:
	//
	// base.combination.edas
	FeatureAnnotations *string `json:"FeatureAnnotations,omitempty" xml:"FeatureAnnotations,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 4
	Instances *int32 `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The number of application instances before the last auto scaling operation.
	//
	// example:
	//
	// 10
	InstancesBeforeScaling *int32 `json:"InstancesBeforeScaling,omitempty" xml:"InstancesBeforeScaling,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The label of an application pod.
	//
	// example:
	//
	// {"test-labelkey":"test-labelvalue"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of CPU cores allowed. Unit: millicores. 1,000 millicores equal one CPU core.
	//
	// example:
	//
	// 1000
	LimitCpuM *int32 `json:"LimitCpuM,omitempty" xml:"LimitCpuM,omitempty"`
	// The maximum size of space required by ephemeral storage. Unit: GB. Value 0 indicates that no limit is set on the space size.
	//
	// example:
	//
	// 4
	LimitEphemeralStorage *string `json:"LimitEphemeralStorage,omitempty" xml:"LimitEphemeralStorage,omitempty"`
	// The maximum size of the memory allowed. Unit: MiB.
	//
	// example:
	//
	// 1024
	LimitMem *int32 `json:"LimitMem,omitempty" xml:"LimitMem,omitempty"`
	// Indicates whether the Graceful Rolling Release and Configure Complete Service Registration before Readiness Probing feature is enabled for the application.
	//
	// example:
	//
	// true
	LosslessRuleAligned *bool `json:"LosslessRuleAligned,omitempty" xml:"LosslessRuleAligned,omitempty"`
	// The delay of service registration. Unit: seconds.
	//
	// example:
	//
	// 120
	LosslessRuleDelayTime *int32 `json:"LosslessRuleDelayTime,omitempty" xml:"LosslessRuleDelayTime,omitempty"`
	// The number of prefetching curves.
	//
	// example:
	//
	// 2
	LosslessRuleFuncType *int32 `json:"LosslessRuleFuncType,omitempty" xml:"LosslessRuleFuncType,omitempty"`
	// Indicates whether the Graceful Rolling Release and Configure Complete Service Prefetching before Readiness Probing feature is enabled for the application.
	//
	// example:
	//
	// true
	LosslessRuleRelated *bool `json:"LosslessRuleRelated,omitempty" xml:"LosslessRuleRelated,omitempty"`
	// The service prefetching duration. Unit: seconds.
	//
	// example:
	//
	// 120
	LosslessRuleWarmupTime *int32 `json:"LosslessRuleWarmupTime,omitempty" xml:"LosslessRuleWarmupTime,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of requested CPU cores. Unit: millicores. 1,000 millicores equal one CPU core.
	//
	// example:
	//
	// 1000
	RequestCpuM *int32 `json:"RequestCpuM,omitempty" xml:"RequestCpuM,omitempty"`
	// The size of space reserved for ephemeral storage resources. Unit: GB. Value 0 indicates that no limit is set on the space size.
	//
	// example:
	//
	// 2
	RequestEphemeralStorage *string `json:"RequestEphemeralStorage,omitempty" xml:"RequestEphemeralStorage,omitempty"`
	// The size of the reserved memory. Unit: MiB.
	//
	// example:
	//
	// 1024
	RequestMem *int32 `json:"RequestMem,omitempty" xml:"RequestMem,omitempty"`
	// The configuration information about the Server Load Balancer (SLB).
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "addressType": "intranet",
	//
	//     "externalTrafficPolicy": "Local",
	//
	//     "ip": "192.168.254.***",
	//
	//     "name": "intranet-testapp",
	//
	//     "portMappings": [
	//
	//       {
	//
	//         "loadBalancerProtocol": "TCP",
	//
	//         "servicePort": {
	//
	//           "port": 8080,
	//
	//           "protocol": "TCP",
	//
	//           "targetPort": 18081,
	//
	//           "vServerGroupName": "k8s/31414/intranet-testapp/default/cc90e0c9508a44667bdae2e83d3******"
	//
	//         }
	//
	//       }
	//
	//     ],
	//
	//     "scheduler": "rr",
	//
	//     "serviceType": "LoadBalancer",
	//
	//     "slbId": "lb-bp1ikoh3nrpgqsm******",
	//
	//     "source": "create",
	//
	//     "specification": "slb.s3.large"
	//
	//   }
	//
	// ]
	SlbInfo *string `json:"SlbInfo,omitempty" xml:"SlbInfo,omitempty"`
	// The version of Apache Tomcat.
	//
	// example:
	//
	// 8.5.55
	TomcatVersion *string `json:"TomcatVersion,omitempty" xml:"TomcatVersion,omitempty"`
	// The workload type. Valid values: Deployment and StatefulSet. If you do not specify this parameter, Deployment is used.
	//
	// example:
	//
	// Deployment
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s GetK8sApplicationResponseBodyApplcationApp) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationApp) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetAnnotations() *string {
	return s.Annotations
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetAppId() *string {
	return s.AppId
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetBuildpackId() *int32 {
	return s.BuildpackId
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetCmd() *string {
	return s.Cmd
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetCmdArgs() *GetK8sApplicationResponseBodyApplcationAppCmdArgs {
	return s.CmdArgs
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetDeployType() *string {
	return s.DeployType
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetDevelopType() *string {
	return s.DevelopType
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetEnableEmptyPushReject() *bool {
	return s.EnableEmptyPushReject
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetEnableLosslessRule() *bool {
	return s.EnableLosslessRule
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetEnvList() *GetK8sApplicationResponseBodyApplcationAppEnvList {
	return s.EnvList
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetFeatureAnnotations() *string {
	return s.FeatureAnnotations
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetInstances() *int32 {
	return s.Instances
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetInstancesBeforeScaling() *int32 {
	return s.InstancesBeforeScaling
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLabels() *string {
	return s.Labels
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLimitCpuM() *int32 {
	return s.LimitCpuM
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLimitEphemeralStorage() *string {
	return s.LimitEphemeralStorage
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLimitMem() *int32 {
	return s.LimitMem
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLosslessRuleAligned() *bool {
	return s.LosslessRuleAligned
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLosslessRuleDelayTime() *int32 {
	return s.LosslessRuleDelayTime
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLosslessRuleFuncType() *int32 {
	return s.LosslessRuleFuncType
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLosslessRuleRelated() *bool {
	return s.LosslessRuleRelated
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetLosslessRuleWarmupTime() *int32 {
	return s.LosslessRuleWarmupTime
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetRegionId() *string {
	return s.RegionId
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetRequestCpuM() *int32 {
	return s.RequestCpuM
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetRequestEphemeralStorage() *string {
	return s.RequestEphemeralStorage
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetRequestMem() *int32 {
	return s.RequestMem
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetSlbInfo() *string {
	return s.SlbInfo
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetTomcatVersion() *string {
	return s.TomcatVersion
}

func (s *GetK8sApplicationResponseBodyApplcationApp) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetAnnotations(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.Annotations = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetAppId(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.AppId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetApplicationName(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.ApplicationName = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetApplicationType(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.ApplicationType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetBuildpackId(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.BuildpackId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetClusterId(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.ClusterId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetCmd(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.Cmd = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetCmdArgs(v *GetK8sApplicationResponseBodyApplcationAppCmdArgs) *GetK8sApplicationResponseBodyApplcationApp {
	s.CmdArgs = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetCsClusterId(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.CsClusterId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetDeployType(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.DeployType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetDevelopType(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.DevelopType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetEdasContainerVersion(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.EdasContainerVersion = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetEnableEmptyPushReject(v bool) *GetK8sApplicationResponseBodyApplcationApp {
	s.EnableEmptyPushReject = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetEnableLosslessRule(v bool) *GetK8sApplicationResponseBodyApplcationApp {
	s.EnableLosslessRule = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetEnvList(v *GetK8sApplicationResponseBodyApplcationAppEnvList) *GetK8sApplicationResponseBodyApplcationApp {
	s.EnvList = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetFeatureAnnotations(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.FeatureAnnotations = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetInstances(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.Instances = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetInstancesBeforeScaling(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.InstancesBeforeScaling = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetK8sNamespace(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.K8sNamespace = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLabels(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.Labels = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLimitCpuM(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.LimitCpuM = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLimitEphemeralStorage(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.LimitEphemeralStorage = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLimitMem(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.LimitMem = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLosslessRuleAligned(v bool) *GetK8sApplicationResponseBodyApplcationApp {
	s.LosslessRuleAligned = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLosslessRuleDelayTime(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.LosslessRuleDelayTime = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLosslessRuleFuncType(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.LosslessRuleFuncType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLosslessRuleRelated(v bool) *GetK8sApplicationResponseBodyApplcationApp {
	s.LosslessRuleRelated = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetLosslessRuleWarmupTime(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.LosslessRuleWarmupTime = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetRegionId(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.RegionId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetRequestCpuM(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.RequestCpuM = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetRequestEphemeralStorage(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.RequestEphemeralStorage = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetRequestMem(v int32) *GetK8sApplicationResponseBodyApplcationApp {
	s.RequestMem = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetSlbInfo(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.SlbInfo = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetTomcatVersion(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.TomcatVersion = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) SetWorkloadType(v string) *GetK8sApplicationResponseBodyApplcationApp {
	s.WorkloadType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationApp) Validate() error {
	if s.CmdArgs != nil {
		if err := s.CmdArgs.Validate(); err != nil {
			return err
		}
	}
	if s.EnvList != nil {
		if err := s.EnvList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetK8sApplicationResponseBodyApplcationAppCmdArgs struct {
	CmdArg []*string `json:"CmdArg,omitempty" xml:"CmdArg,omitempty" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationAppCmdArgs) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationAppCmdArgs) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationAppCmdArgs) GetCmdArg() []*string {
	return s.CmdArg
}

func (s *GetK8sApplicationResponseBodyApplcationAppCmdArgs) SetCmdArg(v []*string) *GetK8sApplicationResponseBodyApplcationAppCmdArgs {
	s.CmdArg = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationAppCmdArgs) Validate() error {
	return dara.Validate(s)
}

type GetK8sApplicationResponseBodyApplcationAppEnvList struct {
	Env []*GetK8sApplicationResponseBodyApplcationAppEnvListEnv `json:"Env,omitempty" xml:"Env,omitempty" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvList) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvList) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvList) GetEnv() []*GetK8sApplicationResponseBodyApplcationAppEnvListEnv {
	return s.Env
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvList) SetEnv(v []*GetK8sApplicationResponseBodyApplcationAppEnvListEnv) *GetK8sApplicationResponseBodyApplcationAppEnvList {
	s.Env = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvList) Validate() error {
	if s.Env != nil {
		for _, item := range s.Env {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetK8sApplicationResponseBodyApplcationAppEnvListEnv struct {
	// The name of the environment variable.
	//
	// example:
	//
	// CATALINA_OPTS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// -Xmx 1024m -Dhsf.default.tid=false $(EDAS_CATALINA_OPTS)
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvListEnv) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationAppEnvListEnv) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvListEnv) GetName() *string {
	return s.Name
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvListEnv) GetValue() *string {
	return s.Value
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvListEnv) SetName(v string) *GetK8sApplicationResponseBodyApplcationAppEnvListEnv {
	s.Name = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvListEnv) SetValue(v string) *GetK8sApplicationResponseBodyApplcationAppEnvListEnv {
	s.Value = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationAppEnvListEnv) Validate() error {
	return dara.Validate(s)
}

type GetK8sApplicationResponseBodyApplcationConf struct {
	// The affinity configuration of the pod.
	//
	// example:
	//
	// "{\\"nodeAffinity\\":{\\"requiredDuringSchedulingIgnoredDuringExecution\\":{\\"nodeSelectorTerms\\":[{\\"matchExpressions\\":[{\\"key\\":\\"beta.kubernetes.io/arch\\",\\"operator\\":\\"NotIn\\",\\"values\\":[\\"arm64\\",\\"arm32\\"]}]}]},\\"preferredDuringSchedulingIgnoredDuringExecution\\":[{\\"weight\\":5,\\"preference\\":{\\"matchExpressions\\":[{\\"key\\":\\"kubernetes.io/os\\",\\"operator\\":\\"In\\",\\"values\\":[\\"linux\\"]}]}}]},\\"podAffinity\\":{\\"requiredDuringSchedulingIgnoredDuringExecution\\":[{\\"labelSelector\\":{\\"matchExpressions\\":[{\\"key\\":\\"edas.oam.acname\\",\\"operator\\":\\"NotIn\\",\\"values\\":[\\"edas-test-app\\"]}]},\\"namespaces\\":[\\"default\\"],\\"topologyKey\\":\\"kubernetes.io/hostname\\"}]},\\"podAntiAffinity\\":{\\"preferredDuringSchedulingIgnoredDuringExecution\\":[{\\"weight\\":15,\\"podAffinityTerm\\":{\\"labelSelector\\":{\\"matchExpressions\\":[{\\"key\\":\\"edas.oam.acname\\",\\"operator\\":\\"In\\",\\"values\\":[\\"edas-test-app-2\\"]}]},\\"namespaces\\":[\\"default\\"],\\"topologyKey\\":\\"failure-domain.beta.kubernetes.io/zone\\"}}]}}"
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// Indicates whether the application is connected to Application High Availability Service (AHAS).
	//
	// example:
	//
	// true
	AhasEnabled *bool `json:"AhasEnabled,omitempty" xml:"AhasEnabled,omitempty"`
	// Indicates whether the application instances are deployed across nodes.
	//
	// 	- Value `true` indicates that the application instances are deployed across nodes.
	//
	// 	- Other values indicate that the application instances are not deployed across nodes.
	//
	// example:
	//
	// true
	DeployAcrossNodes *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
	// Indicates whether the application instances are deployed across zones.
	//
	// 	- Value `true` indicates that the application instances are deployed across zones.
	//
	// 	- Other values indicate that the application instances are not deployed across zones.
	//
	// example:
	//
	// true
	DeployAcrossZones *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	// The startup parameters for a JAR application. This parameter is deprecated.
	//
	// example:
	//
	// -lh
	JarStartArgs *string `json:"JarStartArgs,omitempty" xml:"JarStartArgs,omitempty"`
	// The startup options for a JAR application. This parameter is deprecated.
	//
	// example:
	//
	// -h
	JarStartOptions *string `json:"JarStartOptions,omitempty" xml:"JarStartOptions,omitempty"`
	// The startup command.
	//
	// example:
	//
	// ls
	K8sCmd *string `json:"K8sCmd,omitempty" xml:"K8sCmd,omitempty"`
	// The parameters of the startup command.
	//
	// example:
	//
	// -lh
	K8sCmdArgs *string `json:"K8sCmdArgs,omitempty" xml:"K8sCmdArgs,omitempty"`
	// The information about the local storage.
	//
	// example:
	//
	// [{"type":"","nodePath":"/mnt/","mountPath":"/mnt/"}]
	K8sLocalvolumeInfo *string `json:"K8sLocalvolumeInfo,omitempty" xml:"K8sLocalvolumeInfo,omitempty"`
	// The information about the File Storage NAS (NAS) storage.
	//
	// example:
	//
	// [{"nasPath":"/mnt/","mountPath":"/mnt/"}]
	K8sNasInfo *string `json:"K8sNasInfo,omitempty" xml:"K8sNasInfo,omitempty"`
	// The information about the storage.
	//
	// example:
	//
	// "{\\"hostPaths\\":\\"[]\\",\\"emptyDirs\\":\\"[]\\"}"
	K8sVolumeInfo *string `json:"K8sVolumeInfo,omitempty" xml:"K8sVolumeInfo,omitempty"`
	// The information about the liveness check on the container.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The script executed after the container is started.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script executed before the container is stopped.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The information about the readiness check on the container.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}\\]}}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The type of the container runtime. This parameter is applicable only to clusters that use sandboxed containers.
	//
	// example:
	//
	// runc
	RuntimeClassName *string `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty"`
	// The scheduling tolerance configuration of the pod.
	//
	// example:
	//
	// "[{\\"key\\":\\"edas-taint-key2\\",\\"operator\\":\\"Exists\\",\\"effect\\":\\"NoExecute\\",\\"tolerationSeconds\\":50},{\\"key\\":\\"edas-taint-key\\",\\"operator\\":\\"Equal\\",\\"value\\":\\"edas-taint-value\\",\\"effect\\":\\"PreferNoSchedule\\"}]"
	Tolerations *string `json:"Tolerations,omitempty" xml:"Tolerations,omitempty"`
	// The URL of the base image. If you use a custom Java Development Kit (JDK) runtime, you must specify this parameter.
	//
	// example:
	//
	// openjdk:8u302
	UserBaseImageUrl *string `json:"UserBaseImageUrl,omitempty" xml:"UserBaseImageUrl,omitempty"`
}

func (s GetK8sApplicationResponseBodyApplcationConf) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationConf) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetAffinity() *string {
	return s.Affinity
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetAhasEnabled() *bool {
	return s.AhasEnabled
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetDeployAcrossNodes() *string {
	return s.DeployAcrossNodes
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetDeployAcrossZones() *string {
	return s.DeployAcrossZones
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetJarStartArgs() *string {
	return s.JarStartArgs
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetJarStartOptions() *string {
	return s.JarStartOptions
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetK8sCmd() *string {
	return s.K8sCmd
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetK8sCmdArgs() *string {
	return s.K8sCmdArgs
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetK8sLocalvolumeInfo() *string {
	return s.K8sLocalvolumeInfo
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetK8sNasInfo() *string {
	return s.K8sNasInfo
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetK8sVolumeInfo() *string {
	return s.K8sVolumeInfo
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetLiveness() *string {
	return s.Liveness
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetPostStart() *string {
	return s.PostStart
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetPreStop() *string {
	return s.PreStop
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetReadiness() *string {
	return s.Readiness
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetRuntimeClassName() *string {
	return s.RuntimeClassName
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetTolerations() *string {
	return s.Tolerations
}

func (s *GetK8sApplicationResponseBodyApplcationConf) GetUserBaseImageUrl() *string {
	return s.UserBaseImageUrl
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetAffinity(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.Affinity = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetAhasEnabled(v bool) *GetK8sApplicationResponseBodyApplcationConf {
	s.AhasEnabled = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetDeployAcrossNodes(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.DeployAcrossNodes = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetDeployAcrossZones(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.DeployAcrossZones = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetJarStartArgs(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.JarStartArgs = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetJarStartOptions(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.JarStartOptions = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sCmd(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sCmd = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sCmdArgs(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sCmdArgs = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sLocalvolumeInfo(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sLocalvolumeInfo = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sNasInfo(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sNasInfo = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetK8sVolumeInfo(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.K8sVolumeInfo = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetLiveness(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.Liveness = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetPostStart(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.PostStart = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetPreStop(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.PreStop = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetReadiness(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.Readiness = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetRuntimeClassName(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.RuntimeClassName = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetTolerations(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.Tolerations = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) SetUserBaseImageUrl(v string) *GetK8sApplicationResponseBodyApplcationConf {
	s.UserBaseImageUrl = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationConf) Validate() error {
	return dara.Validate(s)
}

type GetK8sApplicationResponseBodyApplcationDeployGroups struct {
	DeployGroup []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup `json:"DeployGroup,omitempty" xml:"DeployGroup,omitempty" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroups) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroups) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroups) GetDeployGroup() []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup {
	return s.DeployGroup
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroups) SetDeployGroup(v []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) *GetK8sApplicationResponseBodyApplcationDeployGroups {
	s.DeployGroup = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroups) Validate() error {
	if s.DeployGroup != nil {
		for _, item := range s.DeployGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup struct {
	// The information about the component.
	Components *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Struct"`
	// The environment variable. This parameter is different from the EnvList parameter. This parameter specifies the referenced configuration of the ConfigMap or Secret.
	//
	// example:
	//
	// "["{\\"name\\":\\"test1\\",\\"valueFrom\\":{\\"configMapKeyRef\\":{\\"name\\":\\"edas-demo-configmap\\",\\"key\\":\\"key1\\"}}}","{\\"name\\":\\"k2\\",\\"value\\":\\"v2\\"}","{\\"name\\":\\"s1\\",\\"valueFrom\\":{\\"secretKeyRef\\":{\\"name\\":\\"edas-demo-secret\\",\\"key\\":\\"k1\\"}}}"]"
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The source of the environment variable.
	//
	// example:
	//
	// [{"configMapRef":{"name":"test-cm"}}]
	EnvFrom *string `json:"EnvFrom,omitempty" xml:"EnvFrom,omitempty"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) GetComponents() *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents {
	return s.Components
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) GetEnv() *string {
	return s.Env
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) GetEnvFrom() *string {
	return s.EnvFrom
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) SetComponents(v *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup {
	s.Components = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) SetEnv(v string) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup {
	s.Env = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) SetEnvFrom(v string) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup {
	s.EnvFrom = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroup) Validate() error {
	if s.Components != nil {
		if err := s.Components.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents struct {
	Components []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) GetComponents() []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents {
	return s.Components
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) SetComponents(v []*GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents {
	s.Components = v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponents) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents struct {
	// The component ID.
	//
	// example:
	//
	// 5
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The keyword that is included in the component name.
	//
	// example:
	//
	// Open JDK 8
	ComponentKey *string `json:"ComponentKey,omitempty" xml:"ComponentKey,omitempty"`
	// The component type. Valid values:
	//
	// example:
	//
	// JDK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) GetComponentId() *string {
	return s.ComponentId
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) GetComponentKey() *string {
	return s.ComponentKey
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) GetType() *string {
	return s.Type
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) SetComponentId(v string) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents {
	s.ComponentId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) SetComponentKey(v string) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents {
	s.ComponentKey = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) SetType(v string) *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents {
	s.Type = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationDeployGroupsDeployGroupComponentsComponents) Validate() error {
	return dara.Validate(s)
}

type GetK8sApplicationResponseBodyApplcationImageInfo struct {
	// The URL of the image.
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The region ID of the image repository.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// cn-hangzhou
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// 131****067006888_shared_repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the image repository belongs.
	//
	// example:
	//
	// edas-server****-user
	RepoNamespace *string `json:"RepoNamespace,omitempty" xml:"RepoNamespace,omitempty"`
	// The source type of the image repository.
	//
	// example:
	//
	// ALI_HUB
	RepoOriginType *string `json:"RepoOriginType,omitempty" xml:"RepoOriginType,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// 5a166fbd-9d76-4f98-****-781659d9f54c_1572485443282
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetK8sApplicationResponseBodyApplcationImageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationImageInfo) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) GetRepoId() *string {
	return s.RepoId
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) GetRepoName() *string {
	return s.RepoName
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) GetRepoNamespace() *string {
	return s.RepoNamespace
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) GetRepoOriginType() *string {
	return s.RepoOriginType
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) GetTag() *string {
	return s.Tag
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetImageUrl(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.ImageUrl = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRegionId(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RegionId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoId(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoId = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoName(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoName = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoNamespace(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoNamespace = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetRepoOriginType(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.RepoOriginType = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) SetTag(v string) *GetK8sApplicationResponseBodyApplcationImageInfo {
	s.Tag = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationImageInfo) Validate() error {
	return dara.Validate(s)
}

type GetK8sApplicationResponseBodyApplcationLatestVersion struct {
	// The version of the deployment package.
	//
	// example:
	//
	// 20200720
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The URL of the deployment package. This parameter is required if you use a FatJar or WAR package to deploy the application.
	//
	// example:
	//
	// https://e***.oss-cn-beijing.aliyuncs.com/s***-1.0-SNAPSHOT-spring-boot.jar
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The URL of the deployment package. This parameter is required if you use a FatJar or WAR package to deploy the application.
	//
	// example:
	//
	// https://e***.oss-cn-beijing.aliyuncs.com/s***-1.0-SNAPSHOT-spring-boot.jar
	WarUrl *string `json:"WarUrl,omitempty" xml:"WarUrl,omitempty"`
}

func (s GetK8sApplicationResponseBodyApplcationLatestVersion) String() string {
	return dara.Prettify(s)
}

func (s GetK8sApplicationResponseBodyApplcationLatestVersion) GoString() string {
	return s.String()
}

func (s *GetK8sApplicationResponseBodyApplcationLatestVersion) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *GetK8sApplicationResponseBodyApplcationLatestVersion) GetUrl() *string {
	return s.Url
}

func (s *GetK8sApplicationResponseBodyApplcationLatestVersion) GetWarUrl() *string {
	return s.WarUrl
}

func (s *GetK8sApplicationResponseBodyApplcationLatestVersion) SetPackageVersion(v string) *GetK8sApplicationResponseBodyApplcationLatestVersion {
	s.PackageVersion = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationLatestVersion) SetUrl(v string) *GetK8sApplicationResponseBodyApplcationLatestVersion {
	s.Url = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationLatestVersion) SetWarUrl(v string) *GetK8sApplicationResponseBodyApplcationLatestVersion {
	s.WarUrl = &v
	return s
}

func (s *GetK8sApplicationResponseBodyApplcationLatestVersion) Validate() error {
	return dara.Validate(s)
}
