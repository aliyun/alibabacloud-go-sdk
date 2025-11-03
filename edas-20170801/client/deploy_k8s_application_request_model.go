// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployK8sApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *DeployK8sApplicationRequest
	GetAnnotations() *string
	SetAppId(v string) *DeployK8sApplicationRequest
	GetAppId() *string
	SetArgs(v string) *DeployK8sApplicationRequest
	GetArgs() *string
	SetBatchTimeout(v int32) *DeployK8sApplicationRequest
	GetBatchTimeout() *int32
	SetBatchWaitTime(v int32) *DeployK8sApplicationRequest
	GetBatchWaitTime() *int32
	SetBuildPackId(v string) *DeployK8sApplicationRequest
	GetBuildPackId() *string
	SetCanaryRuleId(v string) *DeployK8sApplicationRequest
	GetCanaryRuleId() *string
	SetChangeOrderDesc(v string) *DeployK8sApplicationRequest
	GetChangeOrderDesc() *string
	SetCommand(v string) *DeployK8sApplicationRequest
	GetCommand() *string
	SetConfigMountDescs(v string) *DeployK8sApplicationRequest
	GetConfigMountDescs() *string
	SetCpuLimit(v int32) *DeployK8sApplicationRequest
	GetCpuLimit() *int32
	SetCpuRequest(v int32) *DeployK8sApplicationRequest
	GetCpuRequest() *int32
	SetCustomAffinity(v string) *DeployK8sApplicationRequest
	GetCustomAffinity() *string
	SetCustomAgentVersion(v string) *DeployK8sApplicationRequest
	GetCustomAgentVersion() *string
	SetCustomTolerations(v string) *DeployK8sApplicationRequest
	GetCustomTolerations() *string
	SetDeployAcrossNodes(v string) *DeployK8sApplicationRequest
	GetDeployAcrossNodes() *string
	SetDeployAcrossZones(v string) *DeployK8sApplicationRequest
	GetDeployAcrossZones() *string
	SetEdasContainerVersion(v string) *DeployK8sApplicationRequest
	GetEdasContainerVersion() *string
	SetEmptyDirs(v string) *DeployK8sApplicationRequest
	GetEmptyDirs() *string
	SetEnableAhas(v bool) *DeployK8sApplicationRequest
	GetEnableAhas() *bool
	SetEnableEmptyPushReject(v bool) *DeployK8sApplicationRequest
	GetEnableEmptyPushReject() *bool
	SetEnableLosslessRule(v bool) *DeployK8sApplicationRequest
	GetEnableLosslessRule() *bool
	SetEnvFroms(v string) *DeployK8sApplicationRequest
	GetEnvFroms() *string
	SetEnvs(v string) *DeployK8sApplicationRequest
	GetEnvs() *string
	SetImage(v string) *DeployK8sApplicationRequest
	GetImage() *string
	SetImagePlatforms(v string) *DeployK8sApplicationRequest
	GetImagePlatforms() *string
	SetImageTag(v string) *DeployK8sApplicationRequest
	GetImageTag() *string
	SetInitContainers(v string) *DeployK8sApplicationRequest
	GetInitContainers() *string
	SetJDK(v string) *DeployK8sApplicationRequest
	GetJDK() *string
	SetJavaStartUpConfig(v string) *DeployK8sApplicationRequest
	GetJavaStartUpConfig() *string
	SetLabels(v string) *DeployK8sApplicationRequest
	GetLabels() *string
	SetLimitEphemeralStorage(v int32) *DeployK8sApplicationRequest
	GetLimitEphemeralStorage() *int32
	SetLiveness(v string) *DeployK8sApplicationRequest
	GetLiveness() *string
	SetLocalVolume(v string) *DeployK8sApplicationRequest
	GetLocalVolume() *string
	SetLosslessRuleAligned(v bool) *DeployK8sApplicationRequest
	GetLosslessRuleAligned() *bool
	SetLosslessRuleDelayTime(v int32) *DeployK8sApplicationRequest
	GetLosslessRuleDelayTime() *int32
	SetLosslessRuleFuncType(v int32) *DeployK8sApplicationRequest
	GetLosslessRuleFuncType() *int32
	SetLosslessRuleRelated(v bool) *DeployK8sApplicationRequest
	GetLosslessRuleRelated() *bool
	SetLosslessRuleWarmupTime(v int32) *DeployK8sApplicationRequest
	GetLosslessRuleWarmupTime() *int32
	SetMcpuLimit(v int32) *DeployK8sApplicationRequest
	GetMcpuLimit() *int32
	SetMcpuRequest(v int32) *DeployK8sApplicationRequest
	GetMcpuRequest() *int32
	SetMemoryLimit(v int32) *DeployK8sApplicationRequest
	GetMemoryLimit() *int32
	SetMemoryRequest(v int32) *DeployK8sApplicationRequest
	GetMemoryRequest() *int32
	SetMountDescs(v string) *DeployK8sApplicationRequest
	GetMountDescs() *string
	SetNasId(v string) *DeployK8sApplicationRequest
	GetNasId() *string
	SetPackageUrl(v string) *DeployK8sApplicationRequest
	GetPackageUrl() *string
	SetPackageVersion(v string) *DeployK8sApplicationRequest
	GetPackageVersion() *string
	SetPackageVersionId(v string) *DeployK8sApplicationRequest
	GetPackageVersionId() *string
	SetPostStart(v string) *DeployK8sApplicationRequest
	GetPostStart() *string
	SetPreStop(v string) *DeployK8sApplicationRequest
	GetPreStop() *string
	SetPvcMountDescs(v string) *DeployK8sApplicationRequest
	GetPvcMountDescs() *string
	SetReadiness(v string) *DeployK8sApplicationRequest
	GetReadiness() *string
	SetReplicas(v int32) *DeployK8sApplicationRequest
	GetReplicas() *int32
	SetRequestsEphemeralStorage(v int32) *DeployK8sApplicationRequest
	GetRequestsEphemeralStorage() *int32
	SetRuntimeClassName(v string) *DeployK8sApplicationRequest
	GetRuntimeClassName() *string
	SetSecurityContext(v string) *DeployK8sApplicationRequest
	GetSecurityContext() *string
	SetSidecars(v string) *DeployK8sApplicationRequest
	GetSidecars() *string
	SetSlsConfigs(v string) *DeployK8sApplicationRequest
	GetSlsConfigs() *string
	SetStartup(v string) *DeployK8sApplicationRequest
	GetStartup() *string
	SetStorageType(v string) *DeployK8sApplicationRequest
	GetStorageType() *string
	SetTerminateGracePeriod(v int32) *DeployK8sApplicationRequest
	GetTerminateGracePeriod() *int32
	SetTrafficControlStrategy(v string) *DeployK8sApplicationRequest
	GetTrafficControlStrategy() *string
	SetUpdateStrategy(v string) *DeployK8sApplicationRequest
	GetUpdateStrategy() *string
	SetUriEncoding(v string) *DeployK8sApplicationRequest
	GetUriEncoding() *string
	SetUseBodyEncoding(v bool) *DeployK8sApplicationRequest
	GetUseBodyEncoding() *bool
	SetUserBaseImageUrl(v string) *DeployK8sApplicationRequest
	GetUserBaseImageUrl() *string
	SetVolumesStr(v string) *DeployK8sApplicationRequest
	GetVolumesStr() *string
	SetWebContainer(v string) *DeployK8sApplicationRequest
	GetWebContainer() *string
	SetWebContainerConfig(v string) *DeployK8sApplicationRequest
	GetWebContainerConfig() *string
}

type DeployK8sApplicationRequest struct {
	// The annotation of an application pod.
	//
	// example:
	//
	// {"annotation-name-1":"annotation-value-1","annotation-name-2":"annotation-value-2"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// e83acea6-****-47e1-96ae-c0e953772cdc
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The argument array in the container start-up command. Set this parameter to a JSON array in the format of `["args1","args2"\\]`, where each key is set to a string. If you want to cancel this configuration, set this parameter to an empty JSON array in the format of `"[\\]"`.
	//
	// example:
	//
	// ["args1","args2"]
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The timeout period for an at-a-time release. Unit: seconds.
	//
	// example:
	//
	// 60
	BatchTimeout *int32 `json:"BatchTimeout,omitempty" xml:"BatchTimeout,omitempty"`
	// The minimum time interval for the phased release of pods. For more information, see [minReadySeconds](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#min-ready-seconds).
	//
	// example:
	//
	// 0
	BatchWaitTime *int32 `json:"BatchWaitTime,omitempty" xml:"BatchWaitTime,omitempty"`
	// The build package number of EDAS Container.
	//
	// 	- You do not need to set the parameter if you do not need to change the EDAS Container version during the deployment.
	//
	// 	- Set the parameter if you need to update the EDAS Container version of the application during the deployment.
	//
	// You can query the build package number by using one of the following methods:
	//
	// 	- Call the ListBuildPack operation. For more information, see [ListBuildPack](https://help.aliyun.com/document_detail/423222.html).
	//
	// 	- Obtain the value in the **Build package number*	- column of the [Release notes for EDAS Container](https://help.aliyun.com/document_detail/92614.html) topic. For example, `59` indicates `EDAS Container 3.5.8`.
	//
	// example:
	//
	// 59
	BuildPackId *string `json:"BuildPackId,omitempty" xml:"BuildPackId,omitempty"`
	// example:
	//
	// a8daf22e-****-968c7ff2ea34
	CanaryRuleId *string `json:"CanaryRuleId,omitempty" xml:"CanaryRuleId,omitempty"`
	// The description of the change process.
	//
	// example:
	//
	// Upgrade
	ChangeOrderDesc *string `json:"ChangeOrderDesc,omitempty" xml:"ChangeOrderDesc,omitempty"`
	// The commands that you run to start the container.
	//
	// > If you want to cancel this configuration, set this parameter to an empty string in the format of `""`.
	//
	// example:
	//
	// ls
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
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
	// [
	//
	//       {
	//
	//             "name": "nginx-config",
	//
	//             "type": "ConfigMap",
	//
	//             "mountPath": "/etc/nginx"
	//
	//       },
	//
	//       {
	//
	//             "name": "tls-secret",
	//
	//             "type": "Secret",
	//
	//             "mountPath": "/etc/ssh"
	//
	//       }
	//
	// ]
	ConfigMountDescs *string `json:"ConfigMountDescs,omitempty" xml:"ConfigMountDescs,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is running. Unit: cores. Value 0 indicates that no limit is set on CPU cores.
	//
	// example:
	//
	// 1
	CpuLimit *int32 `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// The number of CPU cores requested for each application instance when the application is running. Unit: cores. We recommend that you set this parameter. Value 0 indicates that no limit is set on CPU cores.
	//
	// > You must set this parameter together with the CpuLimit parameter. Make sure that the value of this parameter does not exceed that of the CpuLimit parameter.
	//
	// example:
	//
	// 0
	CpuRequest *int32 `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	// The affinity configuration of the pod. This parameter takes effect only if both the DeployAcrossNodes and DeployAcrossZones parameters are set to false.
	//
	// example:
	//
	// {"nodeAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":{"nodeSelectorTerms":[{"matchExpressions":[{"key":"beta.kubernetes.io/arch","operator":"NotIn","values":["arm64","arm32"]}]}]},"preferredDuringSchedulingIgnoredDuringExecution":[{"weight":5,"preference":{"matchExpressions":[{"key":"kubernetes.io/os","operator":"In","values":["linux"]}]}}]},"podAffinity":{"requiredDuringSchedulingIgnoredDuringExecution":[{"namespaces":["default"],"topologyKey":"kubernetes.io/hostname","labelSelector":{"matchExpressions":[{"key":"edas.oam.acname","operator":"NotIn","values":["edas-test-app"]}]}}]},"podAntiAffinity":{"preferredDuringSchedulingIgnoredDuringExecution":[{"podAffinityTerm":{"namespaces":["default"],"topologyKey":"failure-domain.beta.kubernetes.io/zone","labelSelector":{"matchExpressions":[{"key":"edas.oam.acname","operator":"In","values":["edas-test-app-2"]}]}},"weight":15}]}}
	CustomAffinity *string `json:"CustomAffinity,omitempty" xml:"CustomAffinity,omitempty"`
	// example:
	//
	// 3.1.4
	CustomAgentVersion *string `json:"CustomAgentVersion,omitempty" xml:"CustomAgentVersion,omitempty"`
	// The scheduling tolerance configuration of the pod. This parameter takes effect only if both the DeployAcrossNodes and DeployAcrossZones parameters are set to false.
	//
	// example:
	//
	// [{"key":"edas-taint-key2","operator":"Exists","effect":"NoExecute","tolerationSeconds":50},{"key":"edas-taint-key","operator":"Equal","value":"edas-taint-value","effect":"PreferNoSchedule"}]
	CustomTolerations *string `json:"CustomTolerations,omitempty" xml:"CustomTolerations,omitempty"`
	// Specifies whether to distribute application instances to multiple nodes. Value true indicates that application instances are distrubuted across zones. Other values indicate that application instances are not distributed across zones.
	//
	// example:
	//
	// true
	DeployAcrossNodes *string `json:"DeployAcrossNodes,omitempty" xml:"DeployAcrossNodes,omitempty"`
	// Specifies whether to distribute application instances across zones. Value true indicates that application instances are distrubuted across zones. Other values indicate that application instances are not distributed across zones.
	//
	// example:
	//
	// true
	DeployAcrossZones *string `json:"DeployAcrossZones,omitempty" xml:"DeployAcrossZones,omitempty"`
	// The version of EDAS Container on which the deployment package of the application depends. This parameter is applicable to High-Speed Service Framework (HSF) applications that you deploy by using WAR packages. This parameter is unavailable if you deploy applications by using images.
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
	// Specifies whether to enable access to Application High Availability Service (AHAS).
	//
	// example:
	//
	// true
	EnableAhas *bool `json:"EnableAhas,omitempty" xml:"EnableAhas,omitempty"`
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
	// This parameter contains the following parameters:
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
	// The environment variables that are used to deploy the application. Set this parameter to a JSON array. Valid values: regular environment variables, Kubernetes ConfigMap environment variables, and Kubernetes Secret environment variables. Specify regular environment variables in the following format:
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
	// >  If you want to cancel this configuration, set this parameter to an empty JSON array, which is in the format of "[]".
	//
	// example:
	//
	// [{"name":"x1","value":"y1"},{"name":"x2","valueFrom":{"configMapKeyRef":{"name":"my-config","key":"y2"}}},{"name":"x3","valueFrom":{"secretKeyRef":{"name":"my-secret","key":"y3"}}}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The absolute URL of the image. This parameter setting overwrites the setting of the ImageTag parameter.
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The destination image platform. This parameter takes effect only when you deploy applications by using .war or .jar packages.
	//
	// 	- If you want to specify x86_64, set the value to linux/amd64.
	//
	// 	- If you want to specify ARM64, set the value to linux/arm64.
	//
	// 	- If you want to specify both x86_64 and ARM64, set the value to linux/amd64,linux/arm64.
	//
	// 	- If you leave this parameter empty, the default architecture is used.
	ImagePlatforms *string `json:"ImagePlatforms,omitempty" xml:"ImagePlatforms,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// latest
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
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
	// The version of the Java Development Kit (JDK) on which the deployment package of the application depends. Open JDK 7 and Open JDK 8 are supported. This parameter is unavailable if you deploy applications by using images.
	//
	// example:
	//
	// Open JDK 8
	JDK *string `json:"JDK,omitempty" xml:"JDK,omitempty"`
	// The configuration of Java startup parameters for a Java application. These startup parameters involve the memory, application, garbage collection (GC) policy, tools, service registration and discovery, and custom configurations. Proper parameter settings help reduce the GC overheads, shorten the server response time, and improve the throughput. Set this parameter to a JSON string. In the example, original indicates the configuration value, and startup indicates a startup parameter. The system automatically concatenates all startup values as the settings of Java startup parameters for the application. To delete this configuration, leave the parameter value empty by entering `""` or `"{}"`.
	//
	// example:
	//
	// {"InitialHeapSize":{"original":512,"startup":"-Xms512m"},"MaxHeapSize":{"original":1024,"startup":"-Xmx1024m"}}
	JavaStartUpConfig *string `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	// The label of an application pod.
	//
	// example:
	//
	// {"label-name-1":"label-value-1","label-name-2":"label-value-2"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum size of space required by ephemeral storage. Unit: GB. Value 0 indicates that no limit is set on the ephemeral storage space.
	//
	// example:
	//
	// 4
	LimitEphemeralStorage *int32 `json:"LimitEphemeralStorage,omitempty" xml:"LimitEphemeralStorage,omitempty"`
	// The configuration for the liveness check on the container. Example: `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}`. If you want to cancel this configuration, set this parameter to `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The configurations that are used when the host files are mounted to the container on which the application is running. Example: `[{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}\\]`. The nodePath parameter specifies the host path, the mountPath parameter specifies the path within the container, and the type parameter specifies the mounting type.
	//
	// example:
	//
	// [{"type":"","nodePath":"/localfiles","mountPath":"/app/files"},{"type":"Directory","nodePath":"/mnt","mountPath":"/app/storage"}]
	LocalVolume *string `json:"LocalVolume,omitempty" xml:"LocalVolume,omitempty"`
	// Specifies whether to enable Graceful Rolling Release and configure Complete Service Registration before Readiness Probing. Valid values:
	//
	// 	- true: If you turn on the switch, the system uses the /health path and provides port 55199 for the health check. The system does not intrude into the application. When the service is registered, the system returns HTTP 200 status code. Otherwise, the system returns HTTP 500 status code.
	//
	// > If you set both the LosslessRuleRelated parameter and this parameter to true, the operation checks whether the service prefetching is complete.
	//
	// 	- false: If you turn off the switch, the system does not provide a port to check whether the service is registered.
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
	// The number of prefetching curves. Valid values: 0 to 20. The default value is 2, which is suitable for common prefetching scenarios. This value indicates that the received traffic amount of the provider during prefetching is displayed as a quadratic curve.
	//
	// example:
	//
	// 2
	LosslessRuleFuncType *int32 `json:"LosslessRuleFuncType,omitempty" xml:"LosslessRuleFuncType,omitempty"`
	// Specifies whether to enable Graceful Rolling Release and configure Complete Service Prefetching before Readiness Probing. Valid values:
	//
	// 	- true: If you turn on the switch, the system uses the /health path and provides port 55199 for the health check. The system does not intrude into the application. When service prefetching is complete, the system returns HTTP 200 status code. Otherwise, the system returns HTTP 500 status code.
	//
	// 	- false: If you turn off the switch, the system does not provide a port to check whether service prefetching is complete.
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
	// The maximum number of CPU cores allowed. Unit: cores. Value 0 indicates that no limit is set on CPU cores.
	//
	// example:
	//
	// 0
	McpuLimit *int32 `json:"McpuLimit,omitempty" xml:"McpuLimit,omitempty"`
	// The minimum number of CPU cores required. Unit: cores. Value 0 indicates that no limit is set on CPU cores.
	//
	// > You must set this parameter together with the CpuLimit parameter. Make sure that the value of this parameter does not exceed that of the CpuLimit parameter.
	//
	// example:
	//
	// 4
	McpuRequest *int32 `json:"McpuRequest,omitempty" xml:"McpuRequest,omitempty"`
	// The maximum size of memory allowed for each application instance when the application is running. Unit: MB. Value 0 indicates that no limit is set on the memory size.
	//
	// example:
	//
	// 0
	MemoryLimit *int32 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// The size of memory requested for each application instance when the application is running. Unit: MB. We recommend that you set this parameter. If you do not want to apply for a memory quota, set this parameter to 0.
	//
	// > You must set this parameter together with the MemoryLimit parameter. Make sure that the value of this parameter does not exceed that of the MemoryLimit parameter.
	//
	// example:
	//
	// 0
	MemoryRequest *int32 `json:"MemoryRequest,omitempty" xml:"MemoryRequest,omitempty"`
	// The description of the NAS mounting configuration. Set this parameter to a serialized JSON string. Example: `[{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}\\]`. The nasPath parameter specifies the file storage path, and the mountPath parameter specifies the path to mount the file system to the container in which the application is running.
	//
	// example:
	//
	// [{"nasPath": "/k8s","mountPath": "/mnt"},{"nasPath": "/files","mountPath": "/app/files"}]
	MountDescs *string `json:"MountDescs,omitempty" xml:"MountDescs,omitempty"`
	// The ID of the File Storage NAS (NAS) file system mounted to the container in which the application is running. The NAS file system must be in the same region as the cluster. The NAS file system must have an available mount target, or have a mount target on the vSwitch in the virtual private cloud (VPC) in which the application resides. If you do not specify this parameter but specify the MountDescs parameter, a NAS file system is automatically purchased and mounted to the vSwitch in the VPC.
	//
	// example:
	//
	// dfs23****
	NasId *string `json:"NasId,omitempty" xml:"NasId,omitempty"`
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
	// The version ID of the deployment package.
	//
	// example:
	//
	// 2bcc********
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	// The post-start script. Example: `{"exec":{"command":["cat","/etc/group"\\]}}`. If you want to cancel this configuration, set this parameter to `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PostStart *string `json:"PostStart,omitempty" xml:"PostStart,omitempty"`
	// The pre-stop script. Example: `{"tcpSocket":{"host":"", "port":8080}}`. If you want to cancel this configuration, set this parameter to `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {\\"exec\\":{\\"command\\":[\\"ls\\",\\"/\\"]}}"
	PreStop *string `json:"PreStop,omitempty" xml:"PreStop,omitempty"`
	// The configuration for mounting a Kubernetes PersistentVolumeClaim (PVC) to a directory in an elastic container instance. The following parameters are included in the configuration:
	//
	// 	- pvcName: the name of the PVC. Make sure that the volume exists and is in the Bound state.
	//
	// 	- mountPaths: the directory to which you want to mount the PVC. You can configure multiple directories. You can set the following two parameters for each mount directory:
	//
	//     	- mountPath: the mount path. The mount path must be an absolute path that starts with a forward slash (/).
	//
	//     	- readOnly: the mount mode. Value true indicates the read-only mode. Value false indicates the read and write mode. Default value: false.
	//
	// example:
	//
	// [{"pvcName":"nas-pvc-1","mountPaths":[{"mountPath":"/usr/share/nginx/data"},{"mountPath":"/usr/share/nginx/html","readOnly":true}]}]
	PvcMountDescs *string `json:"PvcMountDescs,omitempty" xml:"PvcMountDescs,omitempty"`
	// The configuration for the readiness check on the container. If the check fails, the traffic that passes through the Kubernetes service is not transmitted to the container. Example: `{"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}\\]}}`. If you want to cancel this configuration, set this parameter to `""` or `{}`. If you do not specify this parameter, this configuration is ignored.
	//
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"httpGet": {"path": "/consumer","port": 8080,"scheme": "HTTP","httpHeaders": [{"name": "test","value": "testvalue"}]}}
	Readiness *string `json:"Readiness,omitempty" xml:"Readiness,omitempty"`
	// The number of application instances. The minimum value is 0.
	//
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The minimum size of space required by ephemeral storage. Unit: GB. Value 0 indicates that no limit is set on the ephemeral storage space.
	//
	// example:
	//
	// 2
	RequestsEphemeralStorage *int32 `json:"RequestsEphemeralStorage,omitempty" xml:"RequestsEphemeralStorage,omitempty"`
	// The type of the container runtime. Valid values:
	//
	// 	- runc: standard container runtime
	//
	// 	- runv: sandboxed container runtime
	//
	// This parameter is applicable only to clusters that use sandboxed containers.
	//
	// example:
	//
	// runc
	RuntimeClassName *string `json:"RuntimeClassName,omitempty" xml:"RuntimeClassName,omitempty"`
	SecurityContext  *string `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty"`
	// example:
	//
	// [
	//
	//       {
	//
	//             "yamlEncoded": "Y29tbWFuZDoKICAtIHRhaWwKICAtICctZicKICAtIC9kZXYvbnVsbAppbWFnZTogJ2J1c3lib3g6bGF0ZXN0JwpuYW1lOiBidXN5Ym94Cg=="
	//
	//       }
	//
	// ]
	Sidecars *string `json:"Sidecars,omitempty" xml:"Sidecars,omitempty"`
	// The Logstore configuration. If you want to cancel this configuration, leave the parameter value empty by entering `""` or `"{}"`.
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
	//     	- logDir: If the standard output type is used, the collection path is stdout.log. If the file type is used, the collection path is the path of the collected file. Wildcards (\\*) are supported. The collection path must match the following regular expression: `^/(.+)/(.*)^/$`.
	//
	// example:
	//
	// [{"logstore":"thisisanotherfilelog","type":"file","logDir":"/var/log/*"},{"logstore":"","type":"stdout","logDir":"stdout.log"},{"logstore":"thisisafilelog","type":"file","logDir":"/tmp/log/*"}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
	// example:
	//
	// {"failureThreshold": 3,"initialDelaySeconds": 5,"successThreshold": 1,"timeoutSeconds": 1,"tcpSocket":{"host":"", "port":8080}}
	Startup *string `json:"Startup,omitempty" xml:"Startup,omitempty"`
	// The storage type of the NAS file system.
	//
	// 	- Valid values for General-purpose NAS file systems: Capacity and Performance.
	//
	// 	- Valid values for Extreme NAS file systems: standard and advance.
	//
	// You can set this parameter only to Performance.
	//
	// example:
	//
	// Performance
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// 120
	TerminateGracePeriod *int32 `json:"TerminateGracePeriod,omitempty" xml:"TerminateGracePeriod,omitempty"`
	// The traffic adjustment policy for a canary release.
	//
	// example:
	//
	// {"http":{"rules":[{"conditionType":"percent","percent":10}]}}
	TrafficControlStrategy *string `json:"TrafficControlStrategy,omitempty" xml:"TrafficControlStrategy,omitempty"`
	// The phased release policy.
	//
	// 	- Example 1: One instance for a canary release + Two subsequent batches + Automatic batching + 1-minute batch interval.
	//
	// `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}`
	//
	// 	- Example 2: One instance for a canary release + Two subsequent batches + Manual batching.
	//
	// `{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"manual"},"grayUpdate":{"gray":1}}`
	//
	// 	- Example 3: Two batches + Automatic batching + 0-minute batch interval.
	//
	// `{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":0}}`
	//
	// example:
	//
	// {"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}
	UpdateStrategy *string `json:"UpdateStrategy,omitempty" xml:"UpdateStrategy,omitempty"`
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
	// example:
	//
	// openjdk:8u302
	UserBaseImageUrl *string `json:"UserBaseImageUrl,omitempty" xml:"UserBaseImageUrl,omitempty"`
	// The data volume.
	//
	// example:
	//
	// test
	VolumesStr *string `json:"VolumesStr,omitempty" xml:"VolumesStr,omitempty"`
	// The version of the Tomcat container on which the deployment package of the application depends. This parameter is applicable to Spring Cloud and Dubbo applications that you deploy by using WAR packages. This parameter is unavailable if you deploy applications by using images.
	//
	// example:
	//
	// apache-tomcat-7.0.91
	WebContainer *string `json:"WebContainer,omitempty" xml:"WebContainer,omitempty"`
	// The Tomcat container configuration. If you want to cancel this configuration, set this parameter to `""` or `"{}"`. The following parameters are included in the configuration:
	//
	// 	- useDefaultConfig: specifies whether to use the default configuration. Value true indicates to use the default configuration. Value false indicates to use the custom configuration. If the default configuration is used, the following parameters do not take effect.
	//
	// 	- contextInputType: the type of the access path for the application. Valid values:
	//
	//     	- war: The access path for the application is the name of the WAR package. You do not need to specify a custom path.
	//
	//     	- root: The access path for the application is /. You do not need to specify a custom path.
	//
	//     	- custom: If you select this option, you must specify a custom path for the contextPath parameter.
	//
	// 	- contextPath: the custom access path for the application. This parameter is required only when you set the contextInputType parameter to custom.
	//
	// 	- httpPort: the port number. The port number ranges from 1024 to 65535. Though the admin permissions are configured for the container, the root permissions are required to perform operations on ports whose number is less than 1024. Enter a value that ranges from 1025 to 65535 because the container has only the admin permissions. If you do not configure this parameter, the default port number 8080 is used.
	//
	// 	- maxThreads: the maximum number of connections in the connection pool. Default value: 400.
	//
	//     **
	//
	//     **Note**This parameter greatly affects the application performance. We recommend that you set this parameter under professional guidance.
	//
	// 	- uriEncoding: the URI encoding scheme in the Tomcat container. Valid values: UTF-8, ISO-8859-1, GBK, and GB2312. If you do not specify this parameter, the default value ISO-8859-1 is used.
	//
	// 	- useBodyEncoding: specifies whether to use the encoding scheme specified in the request body for URI query parameters.
	//
	// 	- useAdvancedServerXml: specifies whether to use advanced configurations to customize the `server.xml` file. If the preceding parameter types and specific parameters cannot meet your requirements, you can use advanced configurations to customize the `server.xml` file of Tomcat.
	//
	// 	- serverXml: the content of the `server.xml` file customized by using advanced configurations. This parameter takes effect only when you set the useAdvancedServerXml parameter to true.
	//
	// example:
	//
	// {"useDefaultConfig":false,"contextInputType":"custom","contextPath":"hello","httpPort":8088,"maxThreads":400,"uriEncoding":"UTF-8","useBodyEncoding":true,"useAdvancedServerXml":false}
	WebContainerConfig *string `json:"WebContainerConfig,omitempty" xml:"WebContainerConfig,omitempty"`
}

func (s DeployK8sApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployK8sApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *DeployK8sApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeployK8sApplicationRequest) GetArgs() *string {
	return s.Args
}

func (s *DeployK8sApplicationRequest) GetBatchTimeout() *int32 {
	return s.BatchTimeout
}

func (s *DeployK8sApplicationRequest) GetBatchWaitTime() *int32 {
	return s.BatchWaitTime
}

func (s *DeployK8sApplicationRequest) GetBuildPackId() *string {
	return s.BuildPackId
}

func (s *DeployK8sApplicationRequest) GetCanaryRuleId() *string {
	return s.CanaryRuleId
}

func (s *DeployK8sApplicationRequest) GetChangeOrderDesc() *string {
	return s.ChangeOrderDesc
}

func (s *DeployK8sApplicationRequest) GetCommand() *string {
	return s.Command
}

func (s *DeployK8sApplicationRequest) GetConfigMountDescs() *string {
	return s.ConfigMountDescs
}

func (s *DeployK8sApplicationRequest) GetCpuLimit() *int32 {
	return s.CpuLimit
}

func (s *DeployK8sApplicationRequest) GetCpuRequest() *int32 {
	return s.CpuRequest
}

func (s *DeployK8sApplicationRequest) GetCustomAffinity() *string {
	return s.CustomAffinity
}

func (s *DeployK8sApplicationRequest) GetCustomAgentVersion() *string {
	return s.CustomAgentVersion
}

func (s *DeployK8sApplicationRequest) GetCustomTolerations() *string {
	return s.CustomTolerations
}

func (s *DeployK8sApplicationRequest) GetDeployAcrossNodes() *string {
	return s.DeployAcrossNodes
}

func (s *DeployK8sApplicationRequest) GetDeployAcrossZones() *string {
	return s.DeployAcrossZones
}

func (s *DeployK8sApplicationRequest) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DeployK8sApplicationRequest) GetEmptyDirs() *string {
	return s.EmptyDirs
}

func (s *DeployK8sApplicationRequest) GetEnableAhas() *bool {
	return s.EnableAhas
}

func (s *DeployK8sApplicationRequest) GetEnableEmptyPushReject() *bool {
	return s.EnableEmptyPushReject
}

func (s *DeployK8sApplicationRequest) GetEnableLosslessRule() *bool {
	return s.EnableLosslessRule
}

func (s *DeployK8sApplicationRequest) GetEnvFroms() *string {
	return s.EnvFroms
}

func (s *DeployK8sApplicationRequest) GetEnvs() *string {
	return s.Envs
}

func (s *DeployK8sApplicationRequest) GetImage() *string {
	return s.Image
}

func (s *DeployK8sApplicationRequest) GetImagePlatforms() *string {
	return s.ImagePlatforms
}

func (s *DeployK8sApplicationRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DeployK8sApplicationRequest) GetInitContainers() *string {
	return s.InitContainers
}

func (s *DeployK8sApplicationRequest) GetJDK() *string {
	return s.JDK
}

func (s *DeployK8sApplicationRequest) GetJavaStartUpConfig() *string {
	return s.JavaStartUpConfig
}

func (s *DeployK8sApplicationRequest) GetLabels() *string {
	return s.Labels
}

func (s *DeployK8sApplicationRequest) GetLimitEphemeralStorage() *int32 {
	return s.LimitEphemeralStorage
}

func (s *DeployK8sApplicationRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *DeployK8sApplicationRequest) GetLocalVolume() *string {
	return s.LocalVolume
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleAligned() *bool {
	return s.LosslessRuleAligned
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleDelayTime() *int32 {
	return s.LosslessRuleDelayTime
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleFuncType() *int32 {
	return s.LosslessRuleFuncType
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleRelated() *bool {
	return s.LosslessRuleRelated
}

func (s *DeployK8sApplicationRequest) GetLosslessRuleWarmupTime() *int32 {
	return s.LosslessRuleWarmupTime
}

func (s *DeployK8sApplicationRequest) GetMcpuLimit() *int32 {
	return s.McpuLimit
}

func (s *DeployK8sApplicationRequest) GetMcpuRequest() *int32 {
	return s.McpuRequest
}

func (s *DeployK8sApplicationRequest) GetMemoryLimit() *int32 {
	return s.MemoryLimit
}

func (s *DeployK8sApplicationRequest) GetMemoryRequest() *int32 {
	return s.MemoryRequest
}

func (s *DeployK8sApplicationRequest) GetMountDescs() *string {
	return s.MountDescs
}

func (s *DeployK8sApplicationRequest) GetNasId() *string {
	return s.NasId
}

func (s *DeployK8sApplicationRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *DeployK8sApplicationRequest) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *DeployK8sApplicationRequest) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *DeployK8sApplicationRequest) GetPostStart() *string {
	return s.PostStart
}

func (s *DeployK8sApplicationRequest) GetPreStop() *string {
	return s.PreStop
}

func (s *DeployK8sApplicationRequest) GetPvcMountDescs() *string {
	return s.PvcMountDescs
}

func (s *DeployK8sApplicationRequest) GetReadiness() *string {
	return s.Readiness
}

func (s *DeployK8sApplicationRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *DeployK8sApplicationRequest) GetRequestsEphemeralStorage() *int32 {
	return s.RequestsEphemeralStorage
}

func (s *DeployK8sApplicationRequest) GetRuntimeClassName() *string {
	return s.RuntimeClassName
}

func (s *DeployK8sApplicationRequest) GetSecurityContext() *string {
	return s.SecurityContext
}

func (s *DeployK8sApplicationRequest) GetSidecars() *string {
	return s.Sidecars
}

func (s *DeployK8sApplicationRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *DeployK8sApplicationRequest) GetStartup() *string {
	return s.Startup
}

func (s *DeployK8sApplicationRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DeployK8sApplicationRequest) GetTerminateGracePeriod() *int32 {
	return s.TerminateGracePeriod
}

func (s *DeployK8sApplicationRequest) GetTrafficControlStrategy() *string {
	return s.TrafficControlStrategy
}

func (s *DeployK8sApplicationRequest) GetUpdateStrategy() *string {
	return s.UpdateStrategy
}

func (s *DeployK8sApplicationRequest) GetUriEncoding() *string {
	return s.UriEncoding
}

func (s *DeployK8sApplicationRequest) GetUseBodyEncoding() *bool {
	return s.UseBodyEncoding
}

func (s *DeployK8sApplicationRequest) GetUserBaseImageUrl() *string {
	return s.UserBaseImageUrl
}

func (s *DeployK8sApplicationRequest) GetVolumesStr() *string {
	return s.VolumesStr
}

func (s *DeployK8sApplicationRequest) GetWebContainer() *string {
	return s.WebContainer
}

func (s *DeployK8sApplicationRequest) GetWebContainerConfig() *string {
	return s.WebContainerConfig
}

func (s *DeployK8sApplicationRequest) SetAnnotations(v string) *DeployK8sApplicationRequest {
	s.Annotations = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetAppId(v string) *DeployK8sApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetArgs(v string) *DeployK8sApplicationRequest {
	s.Args = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetBatchTimeout(v int32) *DeployK8sApplicationRequest {
	s.BatchTimeout = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetBatchWaitTime(v int32) *DeployK8sApplicationRequest {
	s.BatchWaitTime = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetBuildPackId(v string) *DeployK8sApplicationRequest {
	s.BuildPackId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCanaryRuleId(v string) *DeployK8sApplicationRequest {
	s.CanaryRuleId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetChangeOrderDesc(v string) *DeployK8sApplicationRequest {
	s.ChangeOrderDesc = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCommand(v string) *DeployK8sApplicationRequest {
	s.Command = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetConfigMountDescs(v string) *DeployK8sApplicationRequest {
	s.ConfigMountDescs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCpuLimit(v int32) *DeployK8sApplicationRequest {
	s.CpuLimit = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCpuRequest(v int32) *DeployK8sApplicationRequest {
	s.CpuRequest = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCustomAffinity(v string) *DeployK8sApplicationRequest {
	s.CustomAffinity = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCustomAgentVersion(v string) *DeployK8sApplicationRequest {
	s.CustomAgentVersion = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetCustomTolerations(v string) *DeployK8sApplicationRequest {
	s.CustomTolerations = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetDeployAcrossNodes(v string) *DeployK8sApplicationRequest {
	s.DeployAcrossNodes = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetDeployAcrossZones(v string) *DeployK8sApplicationRequest {
	s.DeployAcrossZones = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEdasContainerVersion(v string) *DeployK8sApplicationRequest {
	s.EdasContainerVersion = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEmptyDirs(v string) *DeployK8sApplicationRequest {
	s.EmptyDirs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnableAhas(v bool) *DeployK8sApplicationRequest {
	s.EnableAhas = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnableEmptyPushReject(v bool) *DeployK8sApplicationRequest {
	s.EnableEmptyPushReject = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnableLosslessRule(v bool) *DeployK8sApplicationRequest {
	s.EnableLosslessRule = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnvFroms(v string) *DeployK8sApplicationRequest {
	s.EnvFroms = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetEnvs(v string) *DeployK8sApplicationRequest {
	s.Envs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetImage(v string) *DeployK8sApplicationRequest {
	s.Image = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetImagePlatforms(v string) *DeployK8sApplicationRequest {
	s.ImagePlatforms = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetImageTag(v string) *DeployK8sApplicationRequest {
	s.ImageTag = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetInitContainers(v string) *DeployK8sApplicationRequest {
	s.InitContainers = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetJDK(v string) *DeployK8sApplicationRequest {
	s.JDK = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetJavaStartUpConfig(v string) *DeployK8sApplicationRequest {
	s.JavaStartUpConfig = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLabels(v string) *DeployK8sApplicationRequest {
	s.Labels = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLimitEphemeralStorage(v int32) *DeployK8sApplicationRequest {
	s.LimitEphemeralStorage = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLiveness(v string) *DeployK8sApplicationRequest {
	s.Liveness = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLocalVolume(v string) *DeployK8sApplicationRequest {
	s.LocalVolume = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleAligned(v bool) *DeployK8sApplicationRequest {
	s.LosslessRuleAligned = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleDelayTime(v int32) *DeployK8sApplicationRequest {
	s.LosslessRuleDelayTime = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleFuncType(v int32) *DeployK8sApplicationRequest {
	s.LosslessRuleFuncType = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleRelated(v bool) *DeployK8sApplicationRequest {
	s.LosslessRuleRelated = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetLosslessRuleWarmupTime(v int32) *DeployK8sApplicationRequest {
	s.LosslessRuleWarmupTime = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMcpuLimit(v int32) *DeployK8sApplicationRequest {
	s.McpuLimit = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMcpuRequest(v int32) *DeployK8sApplicationRequest {
	s.McpuRequest = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMemoryLimit(v int32) *DeployK8sApplicationRequest {
	s.MemoryLimit = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMemoryRequest(v int32) *DeployK8sApplicationRequest {
	s.MemoryRequest = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetMountDescs(v string) *DeployK8sApplicationRequest {
	s.MountDescs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetNasId(v string) *DeployK8sApplicationRequest {
	s.NasId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPackageUrl(v string) *DeployK8sApplicationRequest {
	s.PackageUrl = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPackageVersion(v string) *DeployK8sApplicationRequest {
	s.PackageVersion = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPackageVersionId(v string) *DeployK8sApplicationRequest {
	s.PackageVersionId = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPostStart(v string) *DeployK8sApplicationRequest {
	s.PostStart = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPreStop(v string) *DeployK8sApplicationRequest {
	s.PreStop = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetPvcMountDescs(v string) *DeployK8sApplicationRequest {
	s.PvcMountDescs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetReadiness(v string) *DeployK8sApplicationRequest {
	s.Readiness = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetReplicas(v int32) *DeployK8sApplicationRequest {
	s.Replicas = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetRequestsEphemeralStorage(v int32) *DeployK8sApplicationRequest {
	s.RequestsEphemeralStorage = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetRuntimeClassName(v string) *DeployK8sApplicationRequest {
	s.RuntimeClassName = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetSecurityContext(v string) *DeployK8sApplicationRequest {
	s.SecurityContext = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetSidecars(v string) *DeployK8sApplicationRequest {
	s.Sidecars = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetSlsConfigs(v string) *DeployK8sApplicationRequest {
	s.SlsConfigs = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetStartup(v string) *DeployK8sApplicationRequest {
	s.Startup = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetStorageType(v string) *DeployK8sApplicationRequest {
	s.StorageType = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetTerminateGracePeriod(v int32) *DeployK8sApplicationRequest {
	s.TerminateGracePeriod = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetTrafficControlStrategy(v string) *DeployK8sApplicationRequest {
	s.TrafficControlStrategy = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUpdateStrategy(v string) *DeployK8sApplicationRequest {
	s.UpdateStrategy = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUriEncoding(v string) *DeployK8sApplicationRequest {
	s.UriEncoding = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUseBodyEncoding(v bool) *DeployK8sApplicationRequest {
	s.UseBodyEncoding = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetUserBaseImageUrl(v string) *DeployK8sApplicationRequest {
	s.UserBaseImageUrl = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetVolumesStr(v string) *DeployK8sApplicationRequest {
	s.VolumesStr = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetWebContainer(v string) *DeployK8sApplicationRequest {
	s.WebContainer = &v
	return s
}

func (s *DeployK8sApplicationRequest) SetWebContainerConfig(v string) *DeployK8sApplicationRequest {
	s.WebContainerConfig = &v
	return s
}

func (s *DeployK8sApplicationRequest) Validate() error {
	return dara.Validate(s)
}
