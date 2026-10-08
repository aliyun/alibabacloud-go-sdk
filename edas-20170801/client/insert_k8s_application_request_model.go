// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *InsertK8sApplicationRequest
	GetAnnotations() *string
	SetAppConfig(v string) *InsertK8sApplicationRequest
	GetAppConfig() *string
	SetAppName(v string) *InsertK8sApplicationRequest
	GetAppName() *string
	SetAppTemplateName(v string) *InsertK8sApplicationRequest
	GetAppTemplateName() *string
	SetApplicationDescription(v string) *InsertK8sApplicationRequest
	GetApplicationDescription() *string
	SetBuildPackId(v string) *InsertK8sApplicationRequest
	GetBuildPackId() *string
	SetClusterId(v string) *InsertK8sApplicationRequest
	GetClusterId() *string
	SetCommand(v string) *InsertK8sApplicationRequest
	GetCommand() *string
	SetCommandArgs(v string) *InsertK8sApplicationRequest
	GetCommandArgs() *string
	SetConfigMountDescs(v string) *InsertK8sApplicationRequest
	GetConfigMountDescs() *string
	SetContainerRegistryId(v string) *InsertK8sApplicationRequest
	GetContainerRegistryId() *string
	SetCsClusterId(v string) *InsertK8sApplicationRequest
	GetCsClusterId() *string
	SetCustomAffinity(v string) *InsertK8sApplicationRequest
	GetCustomAffinity() *string
	SetCustomAgentVersion(v string) *InsertK8sApplicationRequest
	GetCustomAgentVersion() *string
	SetCustomTolerations(v string) *InsertK8sApplicationRequest
	GetCustomTolerations() *string
	SetDeployAcrossNodes(v string) *InsertK8sApplicationRequest
	GetDeployAcrossNodes() *string
	SetDeployAcrossZones(v string) *InsertK8sApplicationRequest
	GetDeployAcrossZones() *string
	SetEdasContainerVersion(v string) *InsertK8sApplicationRequest
	GetEdasContainerVersion() *string
	SetEmptyDirs(v string) *InsertK8sApplicationRequest
	GetEmptyDirs() *string
	SetEnableAhas(v bool) *InsertK8sApplicationRequest
	GetEnableAhas() *bool
	SetEnableAsm(v bool) *InsertK8sApplicationRequest
	GetEnableAsm() *bool
	SetEnableEmptyPushReject(v bool) *InsertK8sApplicationRequest
	GetEnableEmptyPushReject() *bool
	SetEnableLosslessRule(v bool) *InsertK8sApplicationRequest
	GetEnableLosslessRule() *bool
	SetEnvFroms(v string) *InsertK8sApplicationRequest
	GetEnvFroms() *string
	SetEnvs(v string) *InsertK8sApplicationRequest
	GetEnvs() *string
	SetFeatureConfig(v string) *InsertK8sApplicationRequest
	GetFeatureConfig() *string
	SetImagePlatforms(v string) *InsertK8sApplicationRequest
	GetImagePlatforms() *string
	SetImageUrl(v string) *InsertK8sApplicationRequest
	GetImageUrl() *string
	SetInitContainers(v string) *InsertK8sApplicationRequest
	GetInitContainers() *string
	SetInternetSlbId(v string) *InsertK8sApplicationRequest
	GetInternetSlbId() *string
	SetInternetSlbPort(v int32) *InsertK8sApplicationRequest
	GetInternetSlbPort() *int32
	SetInternetSlbProtocol(v string) *InsertK8sApplicationRequest
	GetInternetSlbProtocol() *string
	SetInternetTargetPort(v int32) *InsertK8sApplicationRequest
	GetInternetTargetPort() *int32
	SetIntranetSlbId(v string) *InsertK8sApplicationRequest
	GetIntranetSlbId() *string
	SetIntranetSlbPort(v int32) *InsertK8sApplicationRequest
	GetIntranetSlbPort() *int32
	SetIntranetSlbProtocol(v string) *InsertK8sApplicationRequest
	GetIntranetSlbProtocol() *string
	SetIntranetTargetPort(v int32) *InsertK8sApplicationRequest
	GetIntranetTargetPort() *int32
	SetIsMultilingualApp(v bool) *InsertK8sApplicationRequest
	GetIsMultilingualApp() *bool
	SetJDK(v string) *InsertK8sApplicationRequest
	GetJDK() *string
	SetJavaStartUpConfig(v string) *InsertK8sApplicationRequest
	GetJavaStartUpConfig() *string
	SetLabels(v string) *InsertK8sApplicationRequest
	GetLabels() *string
	SetLimitCpu(v int32) *InsertK8sApplicationRequest
	GetLimitCpu() *int32
	SetLimitEphemeralStorage(v int32) *InsertK8sApplicationRequest
	GetLimitEphemeralStorage() *int32
	SetLimitMem(v int32) *InsertK8sApplicationRequest
	GetLimitMem() *int32
	SetLimitmCpu(v int32) *InsertK8sApplicationRequest
	GetLimitmCpu() *int32
	SetLiveness(v string) *InsertK8sApplicationRequest
	GetLiveness() *string
	SetLocalVolume(v string) *InsertK8sApplicationRequest
	GetLocalVolume() *string
	SetLogicalRegionId(v string) *InsertK8sApplicationRequest
	GetLogicalRegionId() *string
	SetLosslessRuleAligned(v bool) *InsertK8sApplicationRequest
	GetLosslessRuleAligned() *bool
	SetLosslessRuleDelayTime(v int32) *InsertK8sApplicationRequest
	GetLosslessRuleDelayTime() *int32
	SetLosslessRuleFuncType(v int32) *InsertK8sApplicationRequest
	GetLosslessRuleFuncType() *int32
	SetLosslessRuleRelated(v bool) *InsertK8sApplicationRequest
	GetLosslessRuleRelated() *bool
	SetLosslessRuleWarmupTime(v int32) *InsertK8sApplicationRequest
	GetLosslessRuleWarmupTime() *int32
	SetMountDescs(v string) *InsertK8sApplicationRequest
	GetMountDescs() *string
	SetNamespace(v string) *InsertK8sApplicationRequest
	GetNamespace() *string
	SetNasId(v string) *InsertK8sApplicationRequest
	GetNasId() *string
	SetPackageType(v string) *InsertK8sApplicationRequest
	GetPackageType() *string
	SetPackageUrl(v string) *InsertK8sApplicationRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *InsertK8sApplicationRequest
	GetPackageVersion() *string
	SetPostStart(v string) *InsertK8sApplicationRequest
	GetPostStart() *string
	SetPreStop(v string) *InsertK8sApplicationRequest
	GetPreStop() *string
	SetPvcMountDescs(v string) *InsertK8sApplicationRequest
	GetPvcMountDescs() *string
	SetReadiness(v string) *InsertK8sApplicationRequest
	GetReadiness() *string
	SetReplicas(v int32) *InsertK8sApplicationRequest
	GetReplicas() *int32
	SetRepoId(v string) *InsertK8sApplicationRequest
	GetRepoId() *string
	SetRequestsCpu(v int32) *InsertK8sApplicationRequest
	GetRequestsCpu() *int32
	SetRequestsEphemeralStorage(v int32) *InsertK8sApplicationRequest
	GetRequestsEphemeralStorage() *int32
	SetRequestsMem(v int32) *InsertK8sApplicationRequest
	GetRequestsMem() *int32
	SetRequestsmCpu(v int32) *InsertK8sApplicationRequest
	GetRequestsmCpu() *int32
	SetResourceGroupId(v string) *InsertK8sApplicationRequest
	GetResourceGroupId() *string
	SetRuntimeClassName(v string) *InsertK8sApplicationRequest
	GetRuntimeClassName() *string
	SetSecretName(v string) *InsertK8sApplicationRequest
	GetSecretName() *string
	SetSecurityContext(v string) *InsertK8sApplicationRequest
	GetSecurityContext() *string
	SetServiceConfigs(v string) *InsertK8sApplicationRequest
	GetServiceConfigs() *string
	SetSidecars(v string) *InsertK8sApplicationRequest
	GetSidecars() *string
	SetSlsConfigs(v string) *InsertK8sApplicationRequest
	GetSlsConfigs() *string
	SetStartup(v string) *InsertK8sApplicationRequest
	GetStartup() *string
	SetStorageType(v string) *InsertK8sApplicationRequest
	GetStorageType() *string
	SetTerminateGracePeriod(v int32) *InsertK8sApplicationRequest
	GetTerminateGracePeriod() *int32
	SetTimeout(v int32) *InsertK8sApplicationRequest
	GetTimeout() *int32
	SetUriEncoding(v string) *InsertK8sApplicationRequest
	GetUriEncoding() *string
	SetUseBodyEncoding(v bool) *InsertK8sApplicationRequest
	GetUseBodyEncoding() *bool
	SetUserBaseImageUrl(v string) *InsertK8sApplicationRequest
	GetUserBaseImageUrl() *string
	SetWebContainer(v string) *InsertK8sApplicationRequest
	GetWebContainer() *string
	SetWebContainerConfig(v string) *InsertK8sApplicationRequest
	GetWebContainerConfig() *string
	SetWorkloadType(v string) *InsertK8sApplicationRequest
	GetWorkloadType() *string
}

type InsertK8sApplicationRequest struct {
	// The annotations of the application pod.
	//
	// example:
	//
	// {"annotation-name-1":"annotation-value-1","annotation-name-2":"annotation-value-2"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The application configuration when an application template is used. The value is a JSON string.
	//
	// example:
	//
	// {}
	AppConfig *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The name of the application. The name must start with a letter and can contain digits, letters, and hyphens (-). The name can be up to 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc-test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the application template that is used to create the application. If you specify an application template when you create the application, the application template and the AppConfig parameter are preferentially used to determine the application configuration. Other configurations are ignored.
	//
	// example:
	//
	// app-template001
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Production Environment
	ApplicationDescription *string `json:"ApplicationDescription,omitempty" xml:"ApplicationDescription,omitempty"`
	// The version of EDAS Container. This parameter conflicts with `EdasContainerVersion`. Use the `EdasContainerVersion` parameter instead.
	//
	// example:
	//
	// -1
	BuildPackId *string `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
	// The ID of the cluster. You can call the ListCluster operation to query the cluster ID. For more information, see [ListCluster](https://help.aliyun.com/document_detail/154995.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// c9cd****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The startup command of the application. If you set this parameter, the original startup command of the image is overridden.
	//
	// example:
	//
	// ls
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments for the startup command. The arguments are a JSON array of strings. Example: `[{"argument":"-c"},{"argument":"test"}]`. In this example, `-c` and `test` are two arguments.
	//
	// example:
	//
	// [{"argument":"-lh"}]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The configuration for mounting Kubernetes ConfigMaps and Secrets. You can mount ConfigMaps and Secrets to specified directories in a container. The following parameters are included in ConfigMountDescs:
	//
	// - name: The name of the ConfigMap or Secret.
	//
	// - type: The configuration type. Valid values: ConfigMap and Secret.
	//
	// - mountPath: The mount path. The path must be an absolute path that starts with a forward slash (/).
	//
	// example:
	//
	// [{"name":"nginx-config","type":"ConfigMap","mountPath":"/etc/nginx"},{"name":"tls-secret","type":"secret","mountPath":"/etc/ssh"}]
	ConfigMountDescs *string `json:"ConfigMountDescs,omitempty" xml:"ConfigMountDescs,omitempty"`
	// The ID of the repository that is used to build the image repository. If you leave this parameter empty, the default repository provided by EDAS is used. Currently, only the default repository provided by EDAS is supported.
	//
	// example:
	//
	// leave empty
	ContainerRegistryId *string `json:"ContainerRegistryId,omitempty" xml:"ContainerRegistryId,omitempty"`
	// You must specify CsClusterId only when you create an application in a cluster that has never been imported.
	//
	// example:
	//
	// abcdefg
	CsClusterId *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty"`
	// The custom affinity.
	//
	// example:
	//
	// demo
	CustomAffinity *string `json:"CustomAffinity,omitempty" xml:"CustomAffinity,omitempty"`
	// The version of the agent.
	//
	// example:
	//
	// 2.8.3,3.2.10,4.3.1
	CustomAgentVersion *string `json:"CustomAgentVersion,omitempty" xml:"CustomAgentVersion,omitempty"`
	// The custom tolerations.
	//
	// example:
	//
	// demo
	CustomTolerations *string `json:"CustomTolerations,omitempty" xml:"CustomTolerations,omitempty"`
	// Specifies whether to distribute application instances to multiple nodes. A value of `true` means yes. Other values mean no.
	//
	// example:
	//
	// true
	DeployAcrossNodes *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
	// Specifies whether to distribute application instances to multiple zones. A value of `true` means yes. Other values mean no.
	//
	// example:
	//
	// true
	DeployAcrossZones *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	// The version of the `EDAS-Container` on which the deployment package depends.
	//
	// > This parameter is not supported for image-based deployments.
	//
	// example:
	//
	// 3.5.9
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The configuration for mounting a Kubernetes emptyDir volume. You can mount an emptyDir volume to a specified directory in a container. The following parameters are included in EmptyDirs:
	//
	// - mountPath: The mount path in the container. This parameter is required.
	//
	// - readOnly: Specifies whether the volume is read-only. This parameter is optional. true specifies read-only. false specifies read and write. Default value: false.
	//
	// - subPathExpr: The subdirectory expression. This parameter is optional.
	//
	// example:
	//
	// [{"mountPath":"/app-log","subPathExpr":"$(POD_IP)"},{"readOnly":true,"mountPath":"/etc/nginx"}]
	EmptyDirs *string `json:"EmptyDirs,omitempty" xml:"EmptyDirs,omitempty"`
	// Specifies whether to enable Application High Availability Service (AHAS):
	//
	// - true: Enable AHAS.
	//
	// - false: Do not enable AHAS.
	//
	// example:
	//
	// true
	EnableAhas *bool `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// You must set this parameter to true only when you create an application in a cluster that has never been imported and enable Service Mesh (ASM).
	//
	// example:
	//
	// false
	EnableAsm *bool `json:"EnableAsm,omitempty" xml:"EnableAsm,omitempty"`
	// Specifies whether to enable protection against empty pushes:
	//
	// - true: Enable protection against empty pushes.
	//
	// - false: Do not enable protection against empty pushes.
	//
	// example:
	//
	// false
	EnableEmptyPushReject *bool `json:"EnableEmptyPushReject,omitempty" xml:"EnableEmptyPushReject,omitempty"`
	// Specifies whether to enable the graceful start rule:
	//
	// - true: Enable the graceful start rule.
	//
	// - false: Do not enable the graceful start rule.
	//
	// example:
	//
	// true
	EnableLosslessRule *bool `json:"EnableLosslessRule,omitempty" xml:"EnableLosslessRule,omitempty"`
	// The configuration for environment variables of the Kubernetes EnvFrom type. You can mount a specified ConfigMap or Secret to a specified directory. Each key corresponds to a file in the directory. The content of the file is the value of the key.
	//
	// The following parameters are included in EnvFroms:
	//
	// - configMapRef: The reference to the ConfigMap. This field includes the following parameter:
	//
	//   - name: The name of the ConfigMap.
	//
	// - secretRef: The reference to the Secret. This field includes the following parameter:
	//
	//   - name: The name of the Secret.
	//
	// example:
	//
	// [{"name":"appname","valueFrom":{"configMapKeyRef":{"name":"appconf","key":"name"}}}]
	EnvFroms *string `json:"EnvFroms,omitempty" xml:"EnvFroms,omitempty"`
	// The environment variables for the deployment. The value must be a JSON array of objects. Three types of environment variables are supported: regular environment variables, Kubernetes ConfigMap environment variables, and Kubernetes Secret environment variables. The format of a regular environment variable is as follows:
	//
	// `{"name":"x", "value": "y"}`
	//
	// You can use a ConfigMap to inject the value of a specific key into a container\\"s environment variable. The format is as follows:
	//
	// `{ "name": "x2", "valueFrom": { "configMapKeyRef": { "name": "my-config", "key": "y2" } } }`
	//
	// You can use a Secret to inject the value of a specific key into a container\\"s environment variable. The format is as follows:
	//
	// `{ "name": "x3", "valueFrom": { "secretKeyRef": { "name": "my-secret", "key": "y3" } } }`
	//
	// > To clear this configuration, set the value to an empty JSON array ([]).
	//
	// example:
	//
	// [{"name":"x1","value":"y1"},{"name":"x2","valueFrom":{"configMapKeyRef":{"name":"my-config","key":"y2"}}},{"name":"x3","valueFrom":{"secretKeyRef":{"name":"my-secret","key":"y3"}}}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The configuration of the custom monitoring and administration solution.
	//
	// example:
	//
	// {"features":[{"name":"base.combination.arms","enable":true},{"name":"base.combination.mse","enable":true}]}
	FeatureConfig *string `json:"FeatureConfig,omitempty" xml:"FeatureConfig,omitempty"`
	// The architecture of the image platform. This parameter is valid when you use a WAR or JAR package for deployment. Examples:
	//
	// - To specify the x86-64 architecture, enter linux/amd64.
	//
	// - To specify the ARM64 architecture, enter linux/arm64.
	//
	// - To build a dual-architecture image, enter linux/amd64,linux/arm64.
	//
	// - If you do not enter a value, the default architecture is used.
	//
	// example:
	//
	// linux/arm64,linux/amd64
	ImagePlatforms *string `json:"ImagePlatforms,omitempty" xml:"ImagePlatforms,omitempty"`
	// The address of the image. This parameter is required when you set `PackageType` to `Image`.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/****_test/****-cons****:1.0
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The init containers for the application pod. You can set the container configuration in the YAML format. The value is the Base64-encoded YAML configuration of the init container.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "yamlEncoded": "Y29tbWFuZDoKICAtIHNsZWVwCiAgLSAnNjAnCmltYWdlOiAnYnVzeWJveDpsYXRlc3QnCm5hbWU6IGluaXQtYnVzeWJveAo="
	//
	//       }
	//
	// ]
	InitContainers *string `json:"InitContainers,omitempty" xml:"InitContainers,omitempty"`
	// The ID of the internet-facing SLB instance. If you do not specify this parameter, EDAS automatically purchases a new SLB instance for you.
	//
	// example:
	//
	// a3d4********
	InternetSlbId *string `json:"InternetSlbId,omitempty" xml:"InternetSlbId,omitempty"`
	// The frontend port of the internet-facing SLB instance. The value must be in the range of 1 to 65535.
	//
	// example:
	//
	// 80
	InternetSlbPort *int32 `json:"InternetSlbPort,omitempty" xml:"InternetSlbPort,omitempty"`
	// The protocol used by the internet-facing SLB instance. Valid values: TCP, HTTP, and HTTPS.
	//
	// example:
	//
	// TCP
	InternetSlbProtocol *string `json:"InternetSlbProtocol,omitempty" xml:"InternetSlbProtocol,omitempty"`
	// The backend port of the internal SLB instance, which also serves as the service port for the application. The port number must be an integer from 1 to 65535.
	//
	// example:
	//
	// 8080
	InternetTargetPort *int32 `json:"InternetTargetPort,omitempty" xml:"InternetTargetPort,omitempty"`
	// The ID of the internal-facing SLB instance. If you do not specify this parameter, EDAS automatically purchases a new SLB instance for you.
	//
	// example:
	//
	// ae93********
	IntranetSlbId *string `json:"IntranetSlbId,omitempty" xml:"IntranetSlbId,omitempty"`
	// The frontend port of the internal-facing SLB instance. The value must be in the range of 1 to 65535.
	//
	// example:
	//
	// 80
	IntranetSlbPort *int32 `json:"IntranetSlbPort,omitempty" xml:"IntranetSlbPort,omitempty"`
	// The protocol used by the internal-facing SLB instance. Valid values: TCP, HTTP, and HTTPS.
	//
	// example:
	//
	// TCP
	IntranetSlbProtocol *string `json:"IntranetSlbProtocol,omitempty" xml:"IntranetSlbProtocol,omitempty"`
	// The backend port of the internal-facing SLB instance. This is also the service port of the application. The value must be in the range of 1 to 65535.
	//
	// example:
	//
	// 80
	IntranetTargetPort *int32 `json:"IntranetTargetPort,omitempty" xml:"IntranetTargetPort,omitempty"`
	// Specifies whether the application is a multilingual application.
	//
	// example:
	//
	// true
	IsMultilingualApp *bool `json:"IsMultilingualApp,omitempty" xml:"IsMultilingualApp,omitempty"`
	// The version of the Java Development Kit (JDK) on which the deployment package depends. Valid values: Open JDK 7, Open JDK 8, and Custom OpenJDK. This parameter is not supported for image-based deployments. If you select Custom OpenJDK, you must also specify the UserBaseImageUrl parameter.
	//
	// example:
	//
	// Open JDK 8
	JDK *string `json:"JDK,omitempty" xml:"JDK,omitempty"`
	// The Java startup parameters. You can configure startup parameters for a Java application. You can configure memory, application, garbage collection (GC) policy, tools, service registration and discovery, and custom parameters. Proper parameter configuration helps reduce GC overhead, shorten server response time, and improve throughput. The value is a JSON string. original specifies the configuration value, and startup specifies the startup parameter. The system automatically concatenates all startup values as the Java startup parameters for the application. To clear the configuration, set the value to `""` or `"{}"`. The keys in the JSON string are described as follows:
	//
	// - InitialHeapSize: the initial heap size.
	//
	// - MaxHeapSize: the maximum heap size.
	//
	// - CustomParams: custom content, such as JVM -D parameters.
	//
	// - Other keys: You can view the JSON structure submitted by the frontend.
	//
	// example:
	//
	// {"InitialHeapSize":{"original":512,"startup":"-Xms512m"},"MaxHeapSize":{"original":1024,"startup":"-Xmx1024m"},"CustomParams":{"original":"-Dcustom.property.sample=false","startup":"-Dcustom.property.sample=false"}}
	JavaStartUpConfig *string `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	// The labels of the application pod.
	//
	// example:
	//
	// {"label-name-1":"label-value-1","label-name-2":"label-value-2"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of CPU cores that can be used by an application instance. If you specify LimitmCpu, this parameter is ignored.
	//
	// example:
	//
	// 4
	LimitCpu *int32 `json:"LimitCpu,omitempty" xml:"LimitCpu,omitempty"`
	// The maximum ephemeral storage. Unit: GB. A value of 0 means no limit.
	//
	// example:
	//
	// 4
	LimitEphemeralStorage *int32 `json:"LimitEphemeralStorage,omitempty" xml:"LimitEphemeralStorage,omitempty"`
	// The maximum amount of memory that can be used by an application instance. Unit: MB. The value of LimitMem must be greater than or equal to the value of RequestsMem.
	//
	// example:
	//
	// 2
	LimitMem *int32 `json:"LimitMem,omitempty" xml:"LimitMem,omitempty"`
	// The maximum number of CPU cores that can be used by an application instance. Unit: millicores. A value of 0 means no limit.
	//
	// example:
	//
	// 1000
	LimitmCpu *int32 `json:"LimitmCpu,omitempty" xml:"LimitmCpu,omitempty"`
	// The liveness probe of the container. Example: `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}`.
	//
	// To clear this configuration, set the value to `""` or `{}`. If you do not set this parameter, it is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The configuration for mounting a host file to a container. Example: `[{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}]`. The following parameters are included:
	//
	// - `nodePath`: the path on the host.
	//
	// - `mountPath`: the path in the container.
	//
	// - `type`: the mount type.
	//
	// example:
	//
	// [{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}]
	LocalVolume *string `json:"LocalVolume,omitempty" xml:"LocalVolume,omitempty"`
	// The ID of the EDAS namespace. This parameter is required if you want to use a non-default namespace.
	//
	// example:
	//
	// cn-shenzhen:beta****
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// Specifies whether to enable the graceful rolling deployment mode in which service registration is complete before the readiness probe is passed:
	//
	// - true: A health check URL is provided for the application on port 55199. The path is /health. The URL returns 200 after the service is registered. Otherwise, the URL returns 500.
	//
	//   > If you also set `LosslessRuleRelated` to `true`, this URL is used to check whether the service warm-up is complete.
	//
	// - false: A URL is not provided for the application to check whether the service is registered.
	//
	// example:
	//
	// false
	LosslessRuleAligned *bool `json:"LosslessRuleAligned,omitempty" xml:"LosslessRuleAligned,omitempty"`
	// The delay of service registration. Unit: seconds. The value must be in the range of 0 to 86400.
	//
	// example:
	//
	// 0
	LosslessRuleDelayTime *int32 `json:"LosslessRuleDelayTime,omitempty" xml:"LosslessRuleDelayTime,omitempty"`
	// The warm-up curve of the service. The value must be in the range of 0 to 20. Default value: 2. This value is suitable for normal warm-up scenarios and indicates that the traffic that the service provider receives follows a quadratic curve during the warm-up period.
	//
	// example:
	//
	// 2
	LosslessRuleFuncType *int32 `json:"LosslessRuleFuncType,omitempty" xml:"LosslessRuleFuncType,omitempty"`
	// Specifies whether to enable the graceful rolling deployment mode in which service warm-up is complete before the readiness probe is passed:
	//
	// - true: A health check URL is provided for the application on port 55199. The path is /health. The URL returns 200 after the service warm-up is complete. Otherwise, the URL returns 500.
	//
	// - false: A URL is not provided for the application to check whether the service warm-up is complete.
	//
	// example:
	//
	// false
	LosslessRuleRelated *bool `json:"LosslessRuleRelated,omitempty" xml:"LosslessRuleRelated,omitempty"`
	// The warm-up duration of the service. Unit: seconds. The value must be in the range of 0 to 86400.
	//
	// example:
	//
	// 120
	LosslessRuleWarmupTime *int32 `json:"LosslessRuleWarmupTime,omitempty" xml:"LosslessRuleWarmupTime,omitempty"`
	// The description of the mount configuration. The value is a serialized JSON string. Example: `[{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}]`. `nasPath` specifies the file storage path. `mountPath` specifies the path to which the file system is mounted in the container.
	//
	// example:
	//
	// [{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}]
	MountDescs *string `json:"MountDescs,omitempty" xml:"MountDescs,omitempty"`
	// The namespace of the Kubernetes cluster. This parameter determines the Kubernetes namespace in which your application is deployed. The default value is default.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the NAS file system that you want to mount. If you do not specify this parameter but mountDescs is specified, a new NAS file system is automatically purchased and mounted to a vSwitch in the VPC.
	//
	// example:
	//
	// dfs23****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The type of the application package. Valid values: FatJar, WAR, and Image.
	//
	// example:
	//
	// WAR
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. This parameter is required for applications that are deployed using a FatJar or WAR package.
	//
	// > The version of the EDAS POP API SDK for Java or Python must be 2.44.0 or later.
	//
	// example:
	//
	// https://e***.oss-cn-beijing.aliyuncs.com/s***-1.0-SNAPSHOT-spring-boot.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version number of the deployment package. This parameter is required for WAR and FatJar packages. You can define the meaning of the version number.
	//
	// > The version of the EDAS POP API SDK for Java or Python must be 2.44.0 or later.
	//
	// example:
	//
	// 20200720
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The script that is run after the container is started. Example: `{"exec":{"command":["cat","/etc/group"]}}`.
	//
	// To clear this configuration, set the value to `""` or `{}`. If you do not set this parameter, it is ignored.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The script that is run before the container is stopped. Example: `{"tcpSocket":{"host":"", "port":8080}}`.
	//
	// To clear this configuration, set the value to `""` or `{}`. If you do not set this parameter, it is ignored.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The configuration for mounting a Kubernetes PersistentVolumeClaim (PVC). You can mount a Kubernetes PVC volume to a specified directory in a container. The following parameters are included in PvcMountDescs:
	//
	// - pvcName: The name of the PVC volume. The PVC volume must exist and be in the Bound state.
	//
	// - mountPaths: The list of mount directories. You can configure multiple mount directories. Each mount directory supports two parameters.
	//
	//   - mountPath: The mount path. The path must be an absolute path that starts with a forward slash (/).
	//
	//   - readOnly: The mount mode. true specifies the read-only mode. false specifies the read and write mode. Default value: false.
	//
	// example:
	//
	// [{"pvcName":"nas-pvc-1","mountPaths":[{"mountPath":"/usr/share/nginx/data"},{"mountPath":"/usr/share/nginx/html","readOnly":true}]}]
	PvcMountDescs *string `json:"PvcMountDescs,omitempty" xml:"PvcMountDescs,omitempty"`
	// The readiness probe of the container. If the check fails, traffic is not routed to the container through the Kubernetes Service. Example: `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}]}}`.
	//
	// To clear this configuration, set the value to `""` or `{}`. If you do not set this parameter, it is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}]}}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 4
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// ced********
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The number of CPU cores requested for an application instance upon creation. Unit: cores. A value of 0 means no limit. If you specify RequestsmCpu, this parameter is ignored.
	//
	// example:
	//
	// 0
	RequestsCpu *int32 `json:"RequestsCpu,omitempty" xml:"RequestsCpu,omitempty"`
	// The minimum ephemeral storage. Unit: GB. A value of 0 means no limit.
	//
	// example:
	//
	// 2
	RequestsEphemeralStorage *int32 `json:"RequestsEphemeralStorage,omitempty" xml:"RequestsEphemeralStorage,omitempty"`
	// The amount of memory requested for an application instance upon creation. Unit: MB. A value of 0 means no limit. The value of RequestsMem cannot be greater than the value of LimitMem.
	//
	// example:
	//
	// 0
	RequestsMem *int32 `json:"RequestsMem,omitempty" xml:"RequestsMem,omitempty"`
	// The number of CPU cores requested for an application instance upon creation. Unit: millicores.
	//
	// example:
	//
	// 500
	RequestsmCpu *int32 `json:"RequestsmCpu,omitempty" xml:"RequestsmCpu,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 461
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the container runtime. This parameter is applicable only to clusters that use sandboxed containers.
	//
	// example:
	//
	// runc
	RuntimeClassName *string `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty"`
	// The name of the image pull secret. You must create the secret.
	//
	// example:
	//
	// edas-app-01-image-secret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The SecurityContext attribute for the application pod container. The value is the Base64-encoded YAML configuration of the SecurityContext.
	//
	// example:
	//
	// {"yamlEncoded":"cnVuQXNVc2VyOiAwCnJ1bkFzR3JvdXA6IDA="}
	SecurityContext *string `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty"`
	// The configuration of the Kubernetes Service.
	//
	// example:
	//
	// [{"name": "test-svc-create","serviceType":"ClusterIP","portMappings":[{"servicePort": {"targetPort":8080,"port":80,"protocol":"TCP"}}]}]
	ServiceConfigs *string `json:"ServiceConfigs,omitempty" xml:"ServiceConfigs,omitempty"`
	// The sidecar containers for the application pod. You can set the container configuration in the YAML format. The value is the Base64-encoded YAML configuration of the sidecar container.
	//
	// example:
	//
	// [{"yamlEncoded":"Y29tbWFuZDoKICAtIHRhaWwKICAtICctZicKICAtIC9kZXYvbnVsbAppbWFnZTogJ2J1c3lib3g6bGF0ZXN0JwpuYW1lOiBidXN5Ym94Cg=="}]
	Sidecars *string `json:"Sidecars,omitempty" xml:"Sidecars,omitempty"`
	// The Logstore configuration. To clear the configuration, set the value to `""` or `"{}"`:
	//
	// - Configs:
	//
	//   - type: The collection type. file indicates the file type. stdout indicates the standard output type.
	//
	//   - Logstore: The name of the Logstore. Make sure that the Logstore name is unique in the same cluster and meets the following naming conventions:
	//
	//     - The name can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	//     - The name must start and end with a lowercase letter or a digit.
	//
	//     - The name must be 3 to 63 characters in length. If you leave this parameter empty, the system automatically generates a name.
	//
	//   - LogDir: If the collection type is standard output, the collection path is stdout.log. If the collection type is file, the collection path is the path of the file to be collected. Wildcards are supported. The collection path must match the following regular expression: `^/(.+)/(.*)^/$`.
	//
	// example:
	//
	// [{"logstore":"thisisanotherfilelog","type":"file","logDir":"/var/log/*"},{"logstore":"","type":"stdout","logDir":"stdout.log"},{"logstore":"thisisafilelog","type":"file","logDir":"/tmp/log/*"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The startup probe. You can use a startup probe to check the liveness of a slow-start container and prevent the container from being killed before it is started. Example: {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}]}}.
	//
	// To clear this configuration, set the value to "" or {}. If you do not set this parameter, it is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Startup *string `json:"Startup,omitempty" xml:"Startup,omitempty"`
	// The storage type of the NAS file system. Valid values:
	//
	// - General-purpose NAS file systems: Capacity and Performance
	//
	// - Extreme NAS file systems: Standard and Advance
	//
	// Currently, only the Performance type is supported.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The timeout period for a graceful stop. Unit: seconds.
	//
	// example:
	//
	// 120
	TerminateGracePeriod *int32 `json:"TerminateGracePeriod,omitempty" xml:"TerminateGracePeriod,omitempty"`
	// The timeout period for the change process. Unit: seconds. The value must be in the range of 1 to 1800. If you do not specify this parameter, the default value 1800 is used.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The URI encoding scheme. Valid values: ISO-8859-1, GBK, GB2312, and UTF-8.
	//
	// > If you do not set this parameter for the application, the default value of Tomcat is used.
	//
	// example:
	//
	// GBK
	UriEncoding *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	// Specifies whether to enable useBodyEncodingForURI.
	//
	// > If you do not set this parameter for the application, the default value false is used.
	//
	// example:
	//
	// false
	UseBodyEncoding *bool `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	// If you use a custom JDK runtime, you must configure the address of the base image. The address must be accessible over the Internet. The EDAS server pulls the image to build an application image.
	//
	// example:
	//
	// openjdk:8u302
	UserBaseImageUrl *string `json:"UserBaseImageUrl,omitempty" xml:"UserBaseImageUrl,omitempty"`
	// The version of the Tomcat container on which the deployment package depends. This parameter is applicable to Spring Cloud and Dubbo applications that are deployed using a WAR package. This parameter is not supported for image-based deployments.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	// The configuration of the Tomcat container. To clear the configuration, set the value to "" or "{}":
	//
	// - useDefaultConfig: Specifies whether to use the default configuration. If you set this parameter to true, the custom configuration is not used. If you set this parameter to false, the custom configuration is used. If you do not use the custom configuration, the following parameter settings do not take effect.
	//
	// - contextInputType: The access path of the application.
	//
	//   - war: You do not need to specify a custom path. The access path is the name of the WAR package.
	//
	//   - root: You do not need to specify a custom path. The access path is `/`.
	//
	//   - custom: You must specify a custom path in the contextPath parameter.
	//
	// - contextPath: The custom path. This parameter is required only when you set contextInputType to custom.
	//
	// - httpPort: The port number. The value must be in the range of 1024 to 65535. Ports smaller than 1024 require root permissions. Because the container is configured with administrator permissions, specify a port number greater than 1024. If you do not specify this parameter, the default port 8080 is used.
	//
	// - maxThreads: The maximum number of connections in the connection pool. Default value: 400.
	//
	//   > This parameter greatly affects application performance. Configure this parameter with the help of a professional.
	//
	// - uriEncoding: The encoding format for Tomcat. Valid values: UTF-8, ISO-8859-1, GBK, and GB2312. If you do not specify this parameter, the default value ISO-8859-1 is used.
	//
	// - useBodyEncoding: Specifies whether to use BodyEncoding for URLs.
	//
	// - useAdvancedServerXml: Specifies whether to use advanced settings to customize the server.xml file. If the preceding parameter types and specific parameters cannot meet your requirements, you can use advanced settings to directly edit the server.xml file of Tomcat.
	//
	// - serverXml: The content of the server.xml file that is customized in the advanced settings. This parameter takes effect only when useAdvancedServerXml is set to true.
	//
	// example:
	//
	// {"useDefaultConfig":false,"contextInputType":"custom","contextPath":"hello","httpPort":8088,"maxThreads":400,"uriEncoding":"UTF-8","useBodyEncoding":true,"useAdvancedServerXml":false}
	WebContainerConfig *string `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty"`
	// The type of the workload. Currently, only deployments are supported.
	//
	// example:
	//
	// Deployment
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s InsertK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *InsertK8sApplicationRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *InsertK8sApplicationRequest) GetAppConfig() *string {
	return s.AppConfig
}

func (s *InsertK8sApplicationRequest) GetAppName() *string {
	return s.AppName
}

func (s *InsertK8sApplicationRequest) GetAppTemplateName() *string {
	return s.AppTemplateName
}

func (s *InsertK8sApplicationRequest) GetApplicationDescription() *string {
	return s.ApplicationDescription
}

func (s *InsertK8sApplicationRequest) GetBuildPackId() *string {
	return s.BuildPackId
}

func (s *InsertK8sApplicationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InsertK8sApplicationRequest) GetCommand() *string {
	return s.Command
}

func (s *InsertK8sApplicationRequest) GetCommandArgs() *string {
	return s.CommandArgs
}

func (s *InsertK8sApplicationRequest) GetConfigMountDescs() *string {
	return s.ConfigMountDescs
}

func (s *InsertK8sApplicationRequest) GetContainerRegistryId() *string {
	return s.ContainerRegistryId
}

func (s *InsertK8sApplicationRequest) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *InsertK8sApplicationRequest) GetCustomAffinity() *string {
	return s.CustomAffinity
}

func (s *InsertK8sApplicationRequest) GetCustomAgentVersion() *string {
	return s.CustomAgentVersion
}

func (s *InsertK8sApplicationRequest) GetCustomTolerations() *string {
	return s.CustomTolerations
}

func (s *InsertK8sApplicationRequest) GetDeployAcrossNodes() *string {
	return s.DeployAcrossNodes
}

func (s *InsertK8sApplicationRequest) GetDeployAcrossZones() *string {
	return s.DeployAcrossZones
}

func (s *InsertK8sApplicationRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *InsertK8sApplicationRequest) GetEmptyDirs() *string {
	return s.EmptyDirs
}

func (s *InsertK8sApplicationRequest) GetEnableAhas() *bool {
	return s.EnableAhas
}

func (s *InsertK8sApplicationRequest) GetEnableAsm() *bool {
	return s.EnableAsm
}

func (s *InsertK8sApplicationRequest) GetEnableEmptyPushReject() *bool {
	return s.EnableEmptyPushReject
}

func (s *InsertK8sApplicationRequest) GetEnableLosslessRule() *bool {
	return s.EnableLosslessRule
}

func (s *InsertK8sApplicationRequest) GetEnvFroms() *string {
	return s.EnvFroms
}

func (s *InsertK8sApplicationRequest) GetEnvs() *string {
	return s.Envs
}

func (s *InsertK8sApplicationRequest) GetFeatureConfig() *string {
	return s.FeatureConfig
}

func (s *InsertK8sApplicationRequest) GetImagePlatforms() *string {
	return s.ImagePlatforms
}

func (s *InsertK8sApplicationRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *InsertK8sApplicationRequest) GetInitContainers() *string {
	return s.InitContainers
}

func (s *InsertK8sApplicationRequest) GetInternetSlbId() *string {
	return s.InternetSlbId
}

func (s *InsertK8sApplicationRequest) GetInternetSlbPort() *int32 {
	return s.InternetSlbPort
}

func (s *InsertK8sApplicationRequest) GetInternetSlbProtocol() *string {
	return s.InternetSlbProtocol
}

func (s *InsertK8sApplicationRequest) GetInternetTargetPort() *int32 {
	return s.InternetTargetPort
}

func (s *InsertK8sApplicationRequest) GetIntranetSlbId() *string {
	return s.IntranetSlbId
}

func (s *InsertK8sApplicationRequest) GetIntranetSlbPort() *int32 {
	return s.IntranetSlbPort
}

func (s *InsertK8sApplicationRequest) GetIntranetSlbProtocol() *string {
	return s.IntranetSlbProtocol
}

func (s *InsertK8sApplicationRequest) GetIntranetTargetPort() *int32 {
	return s.IntranetTargetPort
}

func (s *InsertK8sApplicationRequest) GetIsMultilingualApp() *bool {
	return s.IsMultilingualApp
}

func (s *InsertK8sApplicationRequest) GetJDK() *string {
	return s.JDK
}

func (s *InsertK8sApplicationRequest) GetJavaStartUpConfig() *string {
	return s.JavaStartUpConfig
}

func (s *InsertK8sApplicationRequest) GetLabels() *string {
	return s.Labels
}

func (s *InsertK8sApplicationRequest) GetLimitCpu() *int32 {
	return s.LimitCpu
}

func (s *InsertK8sApplicationRequest) GetLimitEphemeralStorage() *int32 {
	return s.LimitEphemeralStorage
}

func (s *InsertK8sApplicationRequest) GetLimitMem() *int32 {
	return s.LimitMem
}

func (s *InsertK8sApplicationRequest) GetLimitmCpu() *int32 {
	return s.LimitmCpu
}

func (s *InsertK8sApplicationRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *InsertK8sApplicationRequest) GetLocalVolume() *string {
	return s.LocalVolume
}

func (s *InsertK8sApplicationRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *InsertK8sApplicationRequest) GetLosslessRuleAligned() *bool {
	return s.LosslessRuleAligned
}

func (s *InsertK8sApplicationRequest) GetLosslessRuleDelayTime() *int32 {
	return s.LosslessRuleDelayTime
}

func (s *InsertK8sApplicationRequest) GetLosslessRuleFuncType() *int32 {
	return s.LosslessRuleFuncType
}

func (s *InsertK8sApplicationRequest) GetLosslessRuleRelated() *bool {
	return s.LosslessRuleRelated
}

func (s *InsertK8sApplicationRequest) GetLosslessRuleWarmupTime() *int32 {
	return s.LosslessRuleWarmupTime
}

func (s *InsertK8sApplicationRequest) GetMountDescs() *string {
	return s.MountDescs
}

func (s *InsertK8sApplicationRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *InsertK8sApplicationRequest) GetNasId() *string {
	return s.NasId
}

func (s *InsertK8sApplicationRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *InsertK8sApplicationRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *InsertK8sApplicationRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *InsertK8sApplicationRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *InsertK8sApplicationRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *InsertK8sApplicationRequest) GetPvcMountDescs() *string {
	return s.PvcMountDescs
}

func (s *InsertK8sApplicationRequest) GetReadiness() *string {
	return s.Readiness
}

func (s *InsertK8sApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *InsertK8sApplicationRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *InsertK8sApplicationRequest) GetRequestsCpu() *int32 {
	return s.RequestsCpu
}

func (s *InsertK8sApplicationRequest) GetRequestsEphemeralStorage() *int32 {
	return s.RequestsEphemeralStorage
}

func (s *InsertK8sApplicationRequest) GetRequestsMem() *int32 {
	return s.RequestsMem
}

func (s *InsertK8sApplicationRequest) GetRequestsmCpu() *int32 {
	return s.RequestsmCpu
}

func (s *InsertK8sApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InsertK8sApplicationRequest) GetRuntimeClassName() *string {
	return s.RuntimeClassName
}

func (s *InsertK8sApplicationRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *InsertK8sApplicationRequest) GetSecurityContext() *string {
	return s.SecurityContext
}

func (s *InsertK8sApplicationRequest) GetServiceConfigs() *string {
	return s.ServiceConfigs
}

func (s *InsertK8sApplicationRequest) GetSidecars() *string {
	return s.Sidecars
}

func (s *InsertK8sApplicationRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *InsertK8sApplicationRequest) GetStartup() *string {
	return s.Startup
}

func (s *InsertK8sApplicationRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *InsertK8sApplicationRequest) GetTerminateGracePeriod() *int32 {
	return s.TerminateGracePeriod
}

func (s *InsertK8sApplicationRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *InsertK8sApplicationRequest) GetUriEncoding() *string {
	return s.UriEncoding
}

func (s *InsertK8sApplicationRequest) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *InsertK8sApplicationRequest) GetUserBaseImageUrl() *string {
	return s.UserBaseImageUrl
}

func (s *InsertK8sApplicationRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *InsertK8sApplicationRequest) GetWebContainerConfig() *string {
	return s.WebContainerConfig
}

func (s *InsertK8sApplicationRequest) GetWorkloadType() *string {
	return s.WorkloadType
}

func (s *InsertK8sApplicationRequest) SetAnnotations(v string) *InsertK8sApplicationRequest {
	s.Annotations = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetAppConfig(v string) *InsertK8sApplicationRequest {
	s.AppConfig = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetAppName(v string) *InsertK8sApplicationRequest {
	s.AppName = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetAppTemplateName(v string) *InsertK8sApplicationRequest {
	s.AppTemplateName = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetApplicationDescription(v string) *InsertK8sApplicationRequest {
	s.ApplicationDescription = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetBuildPackId(v string) *InsertK8sApplicationRequest {
	s.BuildPackId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetClusterId(v string) *InsertK8sApplicationRequest {
	s.ClusterId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetCommand(v string) *InsertK8sApplicationRequest {
	s.Command = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetCommandArgs(v string) *InsertK8sApplicationRequest {
	s.CommandArgs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetConfigMountDescs(v string) *InsertK8sApplicationRequest {
	s.ConfigMountDescs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetContainerRegistryId(v string) *InsertK8sApplicationRequest {
	s.ContainerRegistryId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetCsClusterId(v string) *InsertK8sApplicationRequest {
	s.CsClusterId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetCustomAffinity(v string) *InsertK8sApplicationRequest {
	s.CustomAffinity = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetCustomAgentVersion(v string) *InsertK8sApplicationRequest {
	s.CustomAgentVersion = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetCustomTolerations(v string) *InsertK8sApplicationRequest {
	s.CustomTolerations = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetDeployAcrossNodes(v string) *InsertK8sApplicationRequest {
	s.DeployAcrossNodes = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetDeployAcrossZones(v string) *InsertK8sApplicationRequest {
	s.DeployAcrossZones = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEdasContainerVersion(v string) *InsertK8sApplicationRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEmptyDirs(v string) *InsertK8sApplicationRequest {
	s.EmptyDirs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEnableAhas(v bool) *InsertK8sApplicationRequest {
	s.EnableAhas = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEnableAsm(v bool) *InsertK8sApplicationRequest {
	s.EnableAsm = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEnableEmptyPushReject(v bool) *InsertK8sApplicationRequest {
	s.EnableEmptyPushReject = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEnableLosslessRule(v bool) *InsertK8sApplicationRequest {
	s.EnableLosslessRule = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEnvFroms(v string) *InsertK8sApplicationRequest {
	s.EnvFroms = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetEnvs(v string) *InsertK8sApplicationRequest {
	s.Envs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetFeatureConfig(v string) *InsertK8sApplicationRequest {
	s.FeatureConfig = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetImagePlatforms(v string) *InsertK8sApplicationRequest {
	s.ImagePlatforms = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetImageUrl(v string) *InsertK8sApplicationRequest {
	s.ImageUrl = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetInitContainers(v string) *InsertK8sApplicationRequest {
	s.InitContainers = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetInternetSlbId(v string) *InsertK8sApplicationRequest {
	s.InternetSlbId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetInternetSlbPort(v int32) *InsertK8sApplicationRequest {
	s.InternetSlbPort = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetInternetSlbProtocol(v string) *InsertK8sApplicationRequest {
	s.InternetSlbProtocol = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetInternetTargetPort(v int32) *InsertK8sApplicationRequest {
	s.InternetTargetPort = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetIntranetSlbId(v string) *InsertK8sApplicationRequest {
	s.IntranetSlbId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetIntranetSlbPort(v int32) *InsertK8sApplicationRequest {
	s.IntranetSlbPort = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetIntranetSlbProtocol(v string) *InsertK8sApplicationRequest {
	s.IntranetSlbProtocol = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetIntranetTargetPort(v int32) *InsertK8sApplicationRequest {
	s.IntranetTargetPort = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetIsMultilingualApp(v bool) *InsertK8sApplicationRequest {
	s.IsMultilingualApp = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetJDK(v string) *InsertK8sApplicationRequest {
	s.JDK = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetJavaStartUpConfig(v string) *InsertK8sApplicationRequest {
	s.JavaStartUpConfig = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLabels(v string) *InsertK8sApplicationRequest {
	s.Labels = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLimitCpu(v int32) *InsertK8sApplicationRequest {
	s.LimitCpu = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLimitEphemeralStorage(v int32) *InsertK8sApplicationRequest {
	s.LimitEphemeralStorage = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLimitMem(v int32) *InsertK8sApplicationRequest {
	s.LimitMem = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLimitmCpu(v int32) *InsertK8sApplicationRequest {
	s.LimitmCpu = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLiveness(v string) *InsertK8sApplicationRequest {
	s.Liveness = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLocalVolume(v string) *InsertK8sApplicationRequest {
	s.LocalVolume = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLogicalRegionId(v string) *InsertK8sApplicationRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLosslessRuleAligned(v bool) *InsertK8sApplicationRequest {
	s.LosslessRuleAligned = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLosslessRuleDelayTime(v int32) *InsertK8sApplicationRequest {
	s.LosslessRuleDelayTime = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLosslessRuleFuncType(v int32) *InsertK8sApplicationRequest {
	s.LosslessRuleFuncType = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLosslessRuleRelated(v bool) *InsertK8sApplicationRequest {
	s.LosslessRuleRelated = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetLosslessRuleWarmupTime(v int32) *InsertK8sApplicationRequest {
	s.LosslessRuleWarmupTime = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetMountDescs(v string) *InsertK8sApplicationRequest {
	s.MountDescs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetNamespace(v string) *InsertK8sApplicationRequest {
	s.Namespace = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetNasId(v string) *InsertK8sApplicationRequest {
	s.NasId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetPackageType(v string) *InsertK8sApplicationRequest {
	s.PackageType = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetPackageUrl(v string) *InsertK8sApplicationRequest {
	s.PackageUrl = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetPackageVersion(v string) *InsertK8sApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetPostStart(v string) *InsertK8sApplicationRequest {
	s.PostStart = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetPreStop(v string) *InsertK8sApplicationRequest {
	s.PreStop = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetPvcMountDescs(v string) *InsertK8sApplicationRequest {
	s.PvcMountDescs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetReadiness(v string) *InsertK8sApplicationRequest {
	s.Readiness = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetReplicas(v int32) *InsertK8sApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetRepoId(v string) *InsertK8sApplicationRequest {
	s.RepoId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetRequestsCpu(v int32) *InsertK8sApplicationRequest {
	s.RequestsCpu = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetRequestsEphemeralStorage(v int32) *InsertK8sApplicationRequest {
	s.RequestsEphemeralStorage = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetRequestsMem(v int32) *InsertK8sApplicationRequest {
	s.RequestsMem = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetRequestsmCpu(v int32) *InsertK8sApplicationRequest {
	s.RequestsmCpu = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetResourceGroupId(v string) *InsertK8sApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetRuntimeClassName(v string) *InsertK8sApplicationRequest {
	s.RuntimeClassName = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetSecretName(v string) *InsertK8sApplicationRequest {
	s.SecretName = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetSecurityContext(v string) *InsertK8sApplicationRequest {
	s.SecurityContext = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetServiceConfigs(v string) *InsertK8sApplicationRequest {
	s.ServiceConfigs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetSidecars(v string) *InsertK8sApplicationRequest {
	s.Sidecars = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetSlsConfigs(v string) *InsertK8sApplicationRequest {
	s.SlsConfigs = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetStartup(v string) *InsertK8sApplicationRequest {
	s.Startup = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetStorageType(v string) *InsertK8sApplicationRequest {
	s.StorageType = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetTerminateGracePeriod(v int32) *InsertK8sApplicationRequest {
	s.TerminateGracePeriod = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetTimeout(v int32) *InsertK8sApplicationRequest {
	s.Timeout = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetUriEncoding(v string) *InsertK8sApplicationRequest {
	s.UriEncoding = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetUseBodyEncoding(v bool) *InsertK8sApplicationRequest {
	s.UseBodyEncoding = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetUserBaseImageUrl(v string) *InsertK8sApplicationRequest {
	s.UserBaseImageUrl = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetWebContainer(v string) *InsertK8sApplicationRequest {
	s.WebContainer = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetWebContainerConfig(v string) *InsertK8sApplicationRequest {
	s.WebContainerConfig = &v
	return s
}

func (s *InsertK8sApplicationRequest) SetWorkloadType(v string) *InsertK8sApplicationRequest {
	s.WorkloadType = &v
	return s
}

func (s *InsertK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
