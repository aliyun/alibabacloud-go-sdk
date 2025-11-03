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
	// The annotation of an application pod.
	//
	// example:
	//
	// {"annotation-name-1":"annotation-value-1","annotation-name-2":"annotation-value-2"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The application configuration when the application template is used. Set this parameter to a JSON array.
	//
	// example:
	//
	// {}
	AppConfig *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The name of the application. The name must start with a letter, and can contain digits, letters, and hyphens (-). It can be up to 36 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// doc-test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the template used to create the application. If you specify an application template when you create an application, the application template and the AppConfig parameter are used to configure the application. Other configurations are ignored.
	//
	// example:
	//
	// app-template001
	AppTemplateName *string `json:"AppTemplateName,omitempty" xml:"AppTemplateName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Application in the production environment
	ApplicationDescription *string `json:"ApplicationDescription,omitempty" xml:"ApplicationDescription,omitempty"`
	// The version of `EDAS Container`. The value of this parameter conflicts with that of the `EdasContainerVersion` parameter. We recommend that you use the `EdasContainerVersion` parameter.
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
	// The application startup command. If you specify this parameter, the value of this parameter will replace the startup command in the image.
	//
	// example:
	//
	// ls
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The arguments in the command. The parameter value is a JSON array of strings. An example is `[{"argument":"-c"},{"argument":"test"}]`, where `-c` and `test` are two arguments that can be set.
	//
	// example:
	//
	// [{"argument":"-lh"}]
	CommandArgs *string `json:"CommandArgs,omitempty" xml:"CommandArgs,omitempty"`
	// The configuration for mounting a Kubernetes ConfigMap or Secret to a directory in an elastic container instance. The following parameters are included in the configuration:
	//
	// 	- name: the name of the Kubernetes ConfigMap or Secret.
	//
	// 	- type: the type of the API object that you want to mount. You can mount a Kubernetes ConfigMap or Secret.
	//
	// 	- mountPath: the mount path. The mount path must be an absolute path that starts with a forward slash (/).
	//
	// example:
	//
	// [{"name":"nginx-config","type":"ConfigMap","mountPath":"/etc/nginx"},{"name":"tls-secret","type":"secret","mountPath":"/etc/ssh"}]
	ConfigMountDescs *string `json:"ConfigMountDescs,omitempty" xml:"ConfigMountDescs,omitempty"`
	// The ID of the repository used to build the image repository. If this parameter is left empty, the default repository provided by EDAS is used. Only the default repository provided by EDAS is supported.
	//
	// example:
	//
	// Leave empty
	ContainerRegistryId *string `json:"ContainerRegistryId,omitempty" xml:"ContainerRegistryId,omitempty"`
	// The ID of the cluster. This parameter is required only when you create the application in a cluster that has not been imported.
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
	CustomAffinity     *string `json:"CustomAffinity,omitempty" xml:"CustomAffinity,omitempty"`
	CustomAgentVersion *string `json:"CustomAgentVersion,omitempty" xml:"CustomAgentVersion,omitempty"`
	// The custom tolerances.
	//
	// example:
	//
	// demo
	CustomTolerations *string `json:"CustomTolerations,omitempty" xml:"CustomTolerations,omitempty"`
	// Specifies whether to distribute application instances across nodes. Value `true` indicates that application instances are distributed across nodes. Other values indicate that application instances are not distributed across nodes.
	//
	// example:
	//
	// true
	DeployAcrossNodes *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
	// Specifies whether to distribute application instances across zones. Value `true` indicates that application instances are distributed across zones. Other values indicate that application instances are not distributed across zones.
	//
	// example:
	//
	// true
	DeployAcrossZones *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	// The version of `EDAS Container` on which the deployment package of the application depends.
	//
	// > This parameter is unavailable if you deploy applications by using images.
	//
	// example:
	//
	// 3.5.9
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
	// The configuration for mounting a Kubernetes emptyDir volume to a directory in an elastic container instance. The following parameters are included in the configuration:
	//
	// 	- mountPath: The mount path in the container. This parameter is required.
	//
	// 	- readOnly: (Optional) The mount mode. Value true indicates the read-only mode. Value false indicates the read and write mode. Default value: false.
	//
	// 	- subPathExpr: (Optional) The regular expression that is used to match the subdirectory.
	//
	// example:
	//
	// [{"mountPath":"/app-log","subPathExpr":"$(POD_IP)"},{"readOnly":true,"mountPath":"/etc/nginx"}]
	EmptyDirs *string `json:"EmptyDirs,omitempty" xml:"EmptyDirs,omitempty"`
	// Specifies whether to enable access to Application High Availability Service (AHAS). Valid values:
	//
	// 	- true: enables access to AHAS.
	//
	// 	- false: does not enable access to AHAS.
	//
	// example:
	//
	// true
	EnableAhas *bool `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
	// Specifies whether to activate Alibaba Cloud Service Mesh (ASM). Set this parameter to true only when you create the application in a cluster that has not been imported and you need to use ASM.
	//
	// example:
	//
	// false
	EnableAsm *bool `json:"EnableAsm,omitempty" xml:"EnableAsm,omitempty"`
	// Specifies whether to enable the empty list protection feature. Valid values:
	//
	// 	- true: enables the empty list protection feature.
	//
	// 	- false: disables the empty list protection feature.
	//
	// example:
	//
	// false
	EnableEmptyPushReject *bool `json:"EnableEmptyPushReject,omitempty" xml:"EnableEmptyPushReject,omitempty"`
	// Specifies whether to enable graceful start rules. Valid values:
	//
	// 	- true: enables graceful start rules.
	//
	// 	- false: disables graceful start rules.
	//
	// example:
	//
	// true
	EnableLosslessRule *bool `json:"EnableLosslessRule,omitempty" xml:"EnableLosslessRule,omitempty"`
	// The Kubernetes environment variables that are configured in EnvFrom mode. A ConfigMap or Secret is mounted to a directory. Each key corresponds to a file in the directory, and the content of the file is the value of the key.
	//
	// The following parameters are included in the configuration:
	//
	// 	- configMapRef: the ConfigMap that is referenced. The following parameter is contained:
	//
	//     	- name: the name of the ConfigMap.
	//
	// 	- secretRef: the Secret that is referenced. The following parameter is contained:
	//
	//     	- name: the name of the Secret.
	//
	// example:
	//
	// [{"name":"appname","valueFrom":{"configMapKeyRef":{"name":"appconf","key":"name"}}}]
	EnvFroms *string `json:"EnvFroms,omitempty" xml:"EnvFroms,omitempty"`
	// The environment variables that are used to deploy the application. The value must be a JSON array. Valid values: regular environment variables, Kubernetes ConfigMap environment variables, or Kubernetes Secret environment variables. Specify regular environment variables in the following format:
	//
	// `{"name":"x", "value": "y"}`
	//
	// Specify Kubernetes ConfigMap environment variables in the following format to reference values from ConfigMaps:
	//
	// `{ "name": "x2", "valueFrom": { "configMapKeyRef": { "name": "my-config", "key": "y2" } } }`
	//
	// Specify Kubernetes Secret environment variables in the following format to reference values from Secrets:
	//
	// `{ "name": "x3", "valueFrom": { "secretKeyRef": { "name": "my-secret", "key": "y3" } } }`
	//
	// >  If you want to cancel this configuration, set this parameter to an empty JSON array in the format of "[]".
	//
	// example:
	//
	// [{"name":"x1","value":"y1"},{"name":"x2","valueFrom":{"configMapKeyRef":{"name":"my-config","key":"y2"}}},{"name":"x3","valueFrom":{"secretKeyRef":{"name":"my-secret","key":"y3"}}}]
	Envs          *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	FeatureConfig *string `json:"FeatureConfig,omitempty" xml:"FeatureConfig,omitempty"`
	// Mirror the target platform architecture, which is effective when deployed using war or jar. Enter an example:
	//
	// - Specify x86 64 architecture: Linux/amd64
	//
	// - Specify ARM 64 architecture: Linux/arm64
	//
	// - Specify the construction of dual architecture images: Linux/amd64, Linux/arm64
	//
	// - Do not input: default schema
	ImagePlatforms *string `json:"ImagePlatforms,omitempty" xml:"ImagePlatforms,omitempty"`
	// The URL of the image. This parameter is required if you set the `PackageType` parameter to `Image`.
	//
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/****_test/****-cons****:1.0
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Set the initialization container for the application Pod. Support setting the format YAML for container configuration, which is the value of Init container YAML configured with base64 encoding.
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
	// The ID of the Internet-facing SLB instance. If you do not specify this parameter, EDAS automatically purchases a new SLB instance for you.
	//
	// example:
	//
	// a3d4********
	InternetSlbId *string `json:"InternetSlbId,omitempty" xml:"InternetSlbId,omitempty"`
	// The frontend port of the Internet-facing SLB instance. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	InternetSlbPort *int32 `json:"InternetSlbPort,omitempty" xml:"InternetSlbPort,omitempty"`
	// The protocol used by the Internet-facing SLB instance. Valid values: TCP, HTTP, and HTTPS.
	//
	// example:
	//
	// TCP
	InternetSlbProtocol *string `json:"InternetSlbProtocol,omitempty" xml:"InternetSlbProtocol,omitempty"`
	// The backend port of the internal-facing SLB instance. This port also serves as the service port of the application. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8080
	InternetTargetPort *int32 `json:"InternetTargetPort,omitempty" xml:"InternetTargetPort,omitempty"`
	// The ID of the internal-facing SLB instance. If you do not specify this parameter, Enterprise Distributed Application Service (EDAS) automatically purchases a new SLB instance for you.
	//
	// example:
	//
	// ae93********
	IntranetSlbId *string `json:"IntranetSlbId,omitempty" xml:"IntranetSlbId,omitempty"`
	// The frontend port of the internal-facing SLB instance. Valid values: 1 to 65535.
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
	// The backend port of the internal-facing Server Load Balancer (SLB) instance. This port also serves as the service port of the application. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	IntranetTargetPort *int32 `json:"IntranetTargetPort,omitempty" xml:"IntranetTargetPort,omitempty"`
	// Specifies whether the application is a multi-language application.
	//
	// example:
	//
	// true
	IsMultilingualApp *bool `json:"IsMultilingualApp,omitempty" xml:"IsMultilingualApp,omitempty"`
	// The version of the Java Development Kit (JDK) on which the deployment package of the application depends. Valid values: Open JDK 7 and Open JDK 8. This parameter is unavailable if you deploy applications by using images.
	//
	// example:
	//
	// Open JDK 8
	JDK *string `json:"JDK,omitempty" xml:"JDK,omitempty"`
	// The configuration of Java startup parameters for a Java application. These startup parameters involve the memory, application, garbage collection (GC) policy, tools, service registration and discovery, and custom configurations. Appropriate parameter settings help reduce the GC overheads, shorten the server response time, and improve the throughput. Set this parameter to a JSON string. In the example, original indicates the configuration value, and startup indicates a startup parameter. The system automatically concatenates all startup values as the settings of Java startup parameters for the application. To delete this configuration, leave the parameter value empty by entering `""` or `"{}"`. The following parameters are included in the configuration:
	//
	// 	- InitialHeapSize: the initial size of the heap memory.
	//
	// 	- MaxHeapSize: the maximum size of the heap memory.
	//
	// 	- CustomParams: the custom parameters, such as JVM -D parameters.
	//
	// 	- Other parameters: You can view the JSON structure submitted by the frontend.
	//
	// example:
	//
	// {"InitialHeapSize":{"original":512,"startup":"-Xms512m"},"MaxHeapSize":{"original":1024,"startup":"-Xmx1024m"},"CustomParams":{"original":"-Dcustom.property.sample=false","startup":"-Dcustom.property.sample=false"}}
	JavaStartUpConfig *string `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	// The label of an application pod.
	//
	// example:
	//
	// {"label-name-1":"label-value-1","label-name-2":"label-value-2"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is running. Unit: cores. If the LimitmCpu parameter is specified, you can ignore this parameter.
	//
	// example:
	//
	// 4
	LimitCpu *int32 `json:"LimitCpu,omitempty" xml:"LimitCpu,omitempty"`
	// The maximum size of space required by ephemeral storage. Unit: GB. Value 0 indicates that no limit is set on the space size.
	//
	// example:
	//
	// 4
	LimitEphemeralStorage *int32 `json:"LimitEphemeralStorage,omitempty" xml:"LimitEphemeralStorage,omitempty"`
	// The maximum size of memory allowed for each application instance when the application is running. Unit: MB. The value of LimitMem must be greater than that of RequestsMem.
	//
	// example:
	//
	// 2
	LimitMem *int32 `json:"LimitMem,omitempty" xml:"LimitMem,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is running. Unit: millicores. Value 0 indicates that no limit is set on CPU cores.
	//
	// example:
	//
	// 1000
	LimitmCpu *int32 `json:"LimitmCpu,omitempty" xml:"LimitmCpu,omitempty"`
	// The configuration for the liveness check on the container. Example: `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}`.
	//
	// If you want to cancel this configuration, leave the parameter value empty by entering `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The configurations that are used when the host files are mounted to the container on which the application is running. Example: `[{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}\\]`. Description:
	//
	// 	- `nodePath`: the host path.
	//
	// 	- `mountPath`: the path in the container.
	//
	// 	- `type`: the mounting type.
	//
	// example:
	//
	// [{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}]
	LocalVolume *string `json:"LocalVolume,omitempty" xml:"LocalVolume,omitempty"`
	// The ID of the EDAS namespace. This parameter is required for a non-default namespace.
	//
	// example:
	//
	// cn-shenzhen:beta****
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// Specifies whether to enable the graceful rolling deployment mode and ensure that the service is registered before the readiness check. Valid values:
	//
	// 	- true: provides port 55199 and the /health path for the health check in a non-intrusive manner. When the service is registered, the system returns HTTP 200 status code. Otherwise, the system returns HTTP 500 status code.
	//
	//     **
	//
	//     **Note**If you set both the `LosslessRuleRelated` parameter and this parameter to `true`, the operation checks whether the service prefetching is complete.
	//
	// 	- false: does not check whether the service is registered.
	//
	// example:
	//
	// false
	LosslessRuleAligned *bool `json:"LosslessRuleAligned,omitempty" xml:"LosslessRuleAligned,omitempty"`
	// The delay of service registration. Valid values: 0 to 86400. Unit: seconds.
	//
	// example:
	//
	// 0
	LosslessRuleDelayTime *int32 `json:"LosslessRuleDelayTime,omitempty" xml:"LosslessRuleDelayTime,omitempty"`
	// The number of prefetching curves. Valid values: 0 to 20. The default value is 2, which is suitable for common prefetching scenarios. This value indicates that the received traffic of the provider during prefetching is displayed as a quadratic curve.
	//
	// example:
	//
	// 2
	LosslessRuleFuncType *int32 `json:"LosslessRuleFuncType,omitempty" xml:"LosslessRuleFuncType,omitempty"`
	// Specifies whether to enable the graceful rolling deployment mode and ensure that the service prefetching is complete before the readiness check. Valid values:
	//
	// 	- true: provides port 55199 and the /health path for the health check in a non-intrusive manner. When the service prefetching is complete, the system returns HTTP 200 status code. Otherwise, the system returns HTTP 500 status code.
	//
	// 	- false: does not check whether the service prefetching is complete.
	//
	// example:
	//
	// false
	LosslessRuleRelated *bool `json:"LosslessRuleRelated,omitempty" xml:"LosslessRuleRelated,omitempty"`
	// The service prefetching duration. Valid values: 0 to 86400. Unit: seconds.
	//
	// example:
	//
	// 120
	LosslessRuleWarmupTime *int32 `json:"LosslessRuleWarmupTime,omitempty" xml:"LosslessRuleWarmupTime,omitempty"`
	// The description of the NAS mounting configuration. Set this parameter to a serialized JSON string. Example: `[{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}\\]`. The `nasPath` parameter specifies the file storage path, and the `mountPath` parameter specifies the path to mount the file system to the container where the application is running.
	//
	// example:
	//
	// [{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}]
	MountDescs *string `json:"MountDescs,omitempty" xml:"MountDescs,omitempty"`
	// The namespace of the Kubernetes cluster. This parameter specifies the Kubernetes namespace in which your application is deployed. By default, the default namespace is used.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the Network Attached Storage (NAS) file system that you want to mount to the application. If you do not specify this parameter but specify the MountDescs parameter, a NAS file system is automatically purchased and mounted to the vSwitch in the VPC.
	//
	// example:
	//
	// dfs23****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
	// The type of the deployment package. Valid values: FatJar, WAR, and Image.
	//
	// example:
	//
	// WAR
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The URL of the deployment package. This parameter is required if you use a FatJar or WAR package to deploy the application.
	//
	// > The version of EDAS SDK for Java or Python must be V2.44.0 or later.
	//
	// example:
	//
	// https://e***.oss-cn-beijing.aliyuncs.com/s***-1.0-SNAPSHOT-spring-boot.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The version of the deployment package. This parameter is required if you use a FatJar or WAR package to deploy the application. You must specify a version.
	//
	// > The version of EDAS SDK for Java or Python must be V2.44.0 or later.
	//
	// example:
	//
	// 20200720
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The post-start script. Example: `{"exec":{"command":["cat","/etc/group"\\]}}`.
	//
	// If you want to cancel this configuration, leave this parameter empty by setting it to `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The pre-stop script. Example: `{"tcpSocket":{"host":"", "port":8080}}`.
	//
	// If you want to cancel this configuration, leave this parameter empty by setting it to `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The configuration for mounting a Kubernetes PersistentVolumeClaim (PVC) volume to a directory in an elastic container instance. The following parameters are included in the configuration:
	//
	// 	- pvcName: the name of the PVC volume. Make sure that the PVC volume is an existing volume and is in the Bound state.
	//
	// 	- mountPaths: the directory to which you want to mount the PVC volume. You can configure multiple directories. You can set the following two parameters for each mount directory:
	//
	//     	- mountPath: the mount path. The mount path must be an absolute path that starts with a forward slash (/).
	//
	//     	- readOnly: the mount mode. Value true indicates the read-only mode. Value false indicates the read and write mode. Default value: false.
	//
	// example:
	//
	// [{"pvcName":"nas-pvc-1","mountPaths":[{"mountPath":"/usr/share/nginx/data"},{"mountPath":"/usr/share/nginx/html","readOnly":true}]}]
	PvcMountDescs *string `json:"PvcMountDescs,omitempty" xml:"PvcMountDescs,omitempty"`
	// The configuration for the readiness check on the container. If the check fails, the traffic that passes through the Kubernetes Service is not transmitted to the container. Example: \\`{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": \[{"name": "test","value": "testvalue"}\\\\]}}\\`.``
	//
	// If you want to cancel this configuration, leave the parameter value empty by entering `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
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
	// The maximum number of CPU cores allowed for each application instance when the application is created. Unit: cores. Value 0 indicates that no limit is set on CPU cores. If the RequestsmCpu parameter is specified, the value of the RequestsmCpu parameter is used. You can ignore this parameter.
	//
	// example:
	//
	// 0
	RequestsCpu *int32 `json:"RequestsCpu,omitempty" xml:"RequestsCpu,omitempty"`
	// The minimum size of space required by ephemeral storage. Unit: GB. Value 0 indicates that no limit is set on the space size.
	//
	// example:
	//
	// 2
	RequestsEphemeralStorage *int32 `json:"RequestsEphemeralStorage,omitempty" xml:"RequestsEphemeralStorage,omitempty"`
	// The maximum size of memory allowed for each application instance when the application is created. Unit: MB. Value 0 indicates that no limit is set on the memory size. The value of RequestsMem cannot be greater than that of LimitMem.
	//
	// example:
	//
	// 0
	RequestsMem *int32 `json:"RequestsMem,omitempty" xml:"RequestsMem,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is created. Unit: millicores.
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
	// The name of the credential that is used to pull the images specified by the user. You must configure the Secret.
	//
	// example:
	//
	// edas-app-01-image-secret
	SecretName      *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	SecurityContext *string `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty"`
	// The configurations of services in a Kubernetes cluster.
	//
	// example:
	//
	// [{"name": "test-svc-create","serviceType":"ClusterIP","portMappings":[{"servicePort": {"targetPort":8080,"port":80,"protocol":"TCP"}}]}]
	ServiceConfigs *string `json:"ServiceConfigs,omitempty" xml:"ServiceConfigs,omitempty"`
	// Set up a Sidecar container for the application Pod. Support setting the format YAML for container configuration, which is the value of Sidecar container YAML configured with base64 encoding.
	//
	// example:
	//
	// [{"yamlEncoded":"Y29tbWFuZDoKICAtIHRhaWwKICAtICctZicKICAtIC9kZXYvbnVsbAppbWFnZTogJ2J1c3lib3g6bGF0ZXN0JwpuYW1lOiBidXN5Ym94Cg=="}]
	Sidecars *string `json:"Sidecars,omitempty" xml:"Sidecars,omitempty"`
	// The Logstore configuration. To delete this configuration, leave the parameter value empty by entering `""` or `"{}"`.
	//
	// 	- The following parameters are included in the configuration:
	//
	//     	- type: the collection type. Set this parameter to file to specify the file type. Set this parameter to stdout to specify the standard output type.
	//
	//     	- logstore: the name of the Logstore. Make sure that the name of the Logstore is unique in the cluster. The name must comply with the following rules:
	//
	//         	- The name can contain only lowercase letters, digits, hyphens (-), and underscores (_).
	//
	//         	- The name must start and end with a lowercase letter or a digit.
	//
	//         	- The name must be 3 to 63 characters in length. If you leave this parameter empty, the system automatically generates a name.
	//
	//     	- LogDir: If the standard output type is used, the collection path is stdout.log. If the file type is used, the collection path is the path of the collected file. Wildcards (\\*) are supported. The collection path must match the following regular expression: `^/(.+)/(.*)^/$`.
	//
	// example:
	//
	// [{"logstore":"thisisanotherfilelog","type":"file","logDir":"/var/log/*"},{"logstore":"","type":"stdout","logDir":"stdout.log"},{"logstore":"thisisafilelog","type":"file","logDir":"/tmp/log/*"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// The startup probe can be used to detect the viability of slow start containers, avoiding them from being killed before startup. The format is as follows: {"FailureThreshold": 3, "initialDelaySeconds": 5, "SuccessThreshold": 1, "timeoutSeconds": 1, "https Get": {"path": "/consumer", "port": 8080, "scheme": "HTTP", "https Headers": [{"name": "test", "value": "testvalue"}]}.
	//
	// If set to "" or {}, it means delete, and if not set, it means ignore.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Startup *string `json:"Startup,omitempty" xml:"Startup,omitempty"`
	// The storage type of the NAS file system.
	//
	// 	- Valid values for General-purpose NAS file systems: Capacity and Performance.
	//
	// 	- Valid values for Extreme NAS file systems: Standard and Advance.
	//
	// You can set this parameter only to Performance.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Set the grace stop timeout for the application. Unit: seconds.
	//
	// example:
	//
	// 120
	TerminateGracePeriod *int32 `json:"TerminateGracePeriod,omitempty" xml:"TerminateGracePeriod,omitempty"`
	// The timeout period of the change process. Valid values: 1 to 1800. Unit: seconds. If you do not specify this Unidentifiedparameter, the default value 1800 is used.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The URI encoding scheme. Valid values: ISO-8859-1, GBK, GB2312, and UTF-8.
	//
	// > If you do not specify this parameter in the application configurations, the default URI encoding scheme in the Tomcat container is applied.
	//
	// example:
	//
	// GBK
	UriEncoding *string `json:"UriEncoding,omitempty" xml:"UriEncoding,omitempty"`
	// Specifies whether to use the encoding scheme specified in the request body for URI query parameters.
	//
	// > If this parameter is not specified in application configuration, the default value false is applied.
	//
	// example:
	//
	// false
	UseBodyEncoding *bool `json:"UseBodyEncoding,omitempty" xml:"UseBodyEncoding,omitempty"`
	// When using custom JDK runtime, it is necessary to configure the basic image address. The address needs to be publicly accessible, and the EDAS server will pull the image to build the application image.
	//
	// example:
	//
	// openjdk:8u302
	UserBaseImageUrl *string `json:"UserBaseImageUrl,omitempty" xml:"UserBaseImageUrl,omitempty"`
	// The version of the Tomcat container on which the deployment package of the application depends. This parameter is applicable to Spring Cloud and Dubbo applications that you deploy by using WAR packages. This parameter is unavailable if you deploy applications by using images.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	// The configuration of the Tomcat container. If you want to cancel this configuration, set this parameter to "" or "{}". The following parameters are included in the configuration:
	//
	// 	- useDefaultConfig: specifies whether to use the default configuration. Value true indicates that the default configuration is used. Value false indicates that the custom configuration is used. If the default configuration is used, the following parameters do not take effect:
	//
	// 	- contextInputType: the type of the access path for the application. Valid values:
	//
	//     	- war: The access path is the name of the WAR package. You do not need to specify a custom path.
	//
	//     	- root: The access path for the application is `/`. You do not need to specify a custom path.
	//
	//     	- custom: If you select this option, you must specify a custom path for the contextPath parameter.
	//
	// 	- contextPath: the custom access path for the application. This parameter is required only when you set the contextInputType parameter to custom.
	//
	// 	- httpPort: the port number. The port number ranges from 1024 to 65535. Though the admin permissions are configured for the container, the root permissions are required to perform operations on ports whose number is less than 1024. Enter a value that ranges from 1024 to 65535 because the container has only the admin permissions. If you do not configure this parameter, the default port number 8080 is used.
	//
	// 	- maxThreads: the maximum number of connections in the connection pool. Default value: 400.
	//
	//     **
	//
	//     **Note**This parameter significantly affects application performance. We recommend that you consult with technical support before you set this parameter.
	//
	// 	- uriEncoding: the URI encoding scheme in the Tomcat container. Valid values: UTF-8, ISO-8859-1, GBK, and GB2312. If you do not specify this parameter, the default value ISO-8859-1 is used.
	//
	// 	- useBodyEncoding: specifies whether to use the encoding scheme specified in the request body for URI query parameters.
	//
	// 	- useAdvancedServerXml: specifies whether to use advanced configurations to customize the server.xml file. If the preceding parameter types and specific parameters cannot meet your requirements, you can use advanced configurations to customize the server.xml file of Tomcat.
	//
	// 	- serverXml: the content of the server.xml file customized by using advanced configurations. This parameter takes effect only when you set the useAdvancedServerXml parameter to true.
	//
	// example:
	//
	// {"useDefaultConfig":false,"contextInputType":"custom","contextPath":"hello","httpPort":8088,"maxThreads":400,"uriEncoding":"UTF-8","useBodyEncoding":true,"useAdvancedServerXml":false}
	WebContainerConfig *string `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty"`
	// The type of Workload when creating an application is currently only supported for the Deployment type.
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
