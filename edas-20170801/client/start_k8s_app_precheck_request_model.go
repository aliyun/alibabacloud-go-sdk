// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartK8sAppPrecheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *StartK8sAppPrecheckRequest
	GetAnnotations() *string
	SetAppId(v string) *StartK8sAppPrecheckRequest
	GetAppId() *string
	SetAppName(v string) *StartK8sAppPrecheckRequest
	GetAppName() *string
	SetClusterId(v string) *StartK8sAppPrecheckRequest
	GetClusterId() *string
	SetComponentIds(v string) *StartK8sAppPrecheckRequest
	GetComponentIds() *string
	SetConfigMountDescs(v string) *StartK8sAppPrecheckRequest
	GetConfigMountDescs() *string
	SetEmptyDirs(v string) *StartK8sAppPrecheckRequest
	GetEmptyDirs() *string
	SetEnvFroms(v string) *StartK8sAppPrecheckRequest
	GetEnvFroms() *string
	SetEnvs(v string) *StartK8sAppPrecheckRequest
	GetEnvs() *string
	SetImageUrl(v string) *StartK8sAppPrecheckRequest
	GetImageUrl() *string
	SetJavaStartUpConfig(v string) *StartK8sAppPrecheckRequest
	GetJavaStartUpConfig() *string
	SetLabels(v string) *StartK8sAppPrecheckRequest
	GetLabels() *string
	SetLimitEphemeralStorage(v int32) *StartK8sAppPrecheckRequest
	GetLimitEphemeralStorage() *int32
	SetLimitMem(v int32) *StartK8sAppPrecheckRequest
	GetLimitMem() *int32
	SetLimitmCpu(v int32) *StartK8sAppPrecheckRequest
	GetLimitmCpu() *int32
	SetLocalVolume(v string) *StartK8sAppPrecheckRequest
	GetLocalVolume() *string
	SetNamespace(v string) *StartK8sAppPrecheckRequest
	GetNamespace() *string
	SetPackageUrl(v string) *StartK8sAppPrecheckRequest
	GetPackageUrl() *string
	SetPvcMountDescs(v string) *StartK8sAppPrecheckRequest
	GetPvcMountDescs() *string
	SetRegionId(v string) *StartK8sAppPrecheckRequest
	GetRegionId() *string
	SetReplicas(v int32) *StartK8sAppPrecheckRequest
	GetReplicas() *int32
	SetRequestsEphemeralStorage(v int32) *StartK8sAppPrecheckRequest
	GetRequestsEphemeralStorage() *int32
	SetRequestsMem(v int32) *StartK8sAppPrecheckRequest
	GetRequestsMem() *int32
	SetRequestsmCpu(v int32) *StartK8sAppPrecheckRequest
	GetRequestsmCpu() *int32
}

type StartK8sAppPrecheckRequest struct {
	// The annotation of an application pod.
	//
	// example:
	//
	// {"annotation-name-1":"annotation-value-1","annotation-name-2":"annotation-value-2"}
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// af58edcf-f7eb-****-****-db4e425f****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application. The name must start with a letter, and can contain digits, letters, and hyphens (-). It can be up to 36 characters in length.
	//
	// example:
	//
	// testapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c37aec2a-bcca-4ec1-****-****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the application component. You can call the ListComponents operation to query application components. This parameter must be specified when the application runs in Apache Tomcat or in a standard Java application runtime environment. The Apache Tomcat application runtime environment is applicable to Dubbo applications that are deployed by using WAR packages. A standard Java application runtime environment is applicable to Spring Boot or Spring Cloud applications that are deployed by using JAR packages.
	//
	// Valid values for regular application component IDs:
	//
	// 	- 4: Apache Tomcat 7.0.91
	//
	// 	- 5: OpenJDK 1.8.x
	//
	// 	- 6: OpenJDK 1.7.x
	//
	// 	- 7: Apache Tomcat 8.5.42
	//
	// This parameter is available only for Java SDK 2.57.3 or later, or Python SDK 2.57.3 or later. Assume that you use an SDK that is not provided by Enterprise Distributed Application Service (EDAS), such as aliyun-python-sdk-core, aliyun-java-sdk-core, and Alibaba Cloud CLI. In this case, you can directly specify this parameter.
	//
	// example:
	//
	// 7
	ComponentIds *string `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty"`
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
	// The configuration for mounting a Kubernetes emptyDir volume to a directory in an elastic container instance. The following parameters are included in the configuration:
	//
	// 	- mountPath: The mount path in the container. This parameter is required.
	//
	// 	- readOnly: (Optional) The mount mode. The value true indicates the read-only mode. The value false indicates the read and write mode. Default value: false.
	//
	// 	- subPathExpr: (Optional) The regular expression that is used to match the subdirectory.
	//
	// example:
	//
	// [{"mountPath":"/app-log","subPathExpr":"$(POD_IP)"},{"readOnly":true,"mountPath":"/etc/nginx"}]
	EmptyDirs *string `json:"EmptyDirs,omitempty" xml:"EmptyDirs,omitempty"`
	// The Kubernetes environment variables that are configured in EnvFrom mode. A ConfigMap or Secret is mounted to a directory. Each key corresponds to a file in the directory, and the content of the file is the value of the key.
	//
	// The following parameters are included in the configuration of the EnvFroms parameter:
	//
	// 	- configMapRef: the ConfigMap that is referenced. The following parameter is included:
	//
	//     name: the name of the ConfigMap.
	//
	// 	- secretRef: the Secret that is referenced. The following parameter is included:
	//
	//     name: the name of the Secret.
	//
	// example:
	//
	// [{"name":"appname","valueFrom":{"configMapKeyRef":{"name":"appconf","key":"name"}}}]
	EnvFroms *string `json:"EnvFroms,omitempty" xml:"EnvFroms,omitempty"`
	// The environment variables that are used to deploy the application. The value must be a JSON array. Valid values: regular environment variables, Kubernetes ConfigMap environment variables, and Kubernetes Secret environment variables. Specify regular environment variables in the following format:
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
	// > If you want to cancel this configuration, set this parameter to an empty JSON array, which is in the format of "[]".
	//
	// example:
	//
	// [{"name":"x1","value":"y1"},{"name":"x2","valueFrom":{"configMapKeyRef":{"name":"my-config","key":"y2"}}},{"name":"x3","valueFrom":{"secretKeyRef":{"name":"my-secret","key":"y3"}}}]
	Envs *string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The URL of the image.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/mw/testapp:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The configuration of Java startup parameters for a Java application. These startup parameters involve the memory, application, garbage collection (GC) policy, tools, service registration and discovery, and custom configurations. Proper parameter settings help reduce the GC overheads, shorten the server response time, and improve the throughput. Set this parameter to a JSON string. In the example, original indicates the configuration value, and startup indicates a startup parameter. The system automatically concatenates all startup values as the settings of Java startup parameters for the application. To delete this configuration, leave the parameter value empty by entering `""` or `"{}"`. The following parameters are included in the configuration:
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
	// {"InitialHeapSize":{"original":512,"startup":"-Xms512m"},"MaxHeapSize":{"original":1024,"startup":"-Xmx1024m"}}
	JavaStartUpConfig *string `json:"JavaStartUpConfig,omitempty" xml:"JavaStartUpConfig,omitempty"`
	// The label of an application pod.
	//
	// example:
	//
	// {"label-name-1":"label-value-1","label-name-2":"label-value-2"}
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum size of space required by ephemeral storage. Unit: GB. The value 0 indicates that no limit is set on the ephemeral storage space.
	//
	// example:
	//
	// 4
	LimitEphemeralStorage *int32 `json:"LimitEphemeralStorage,omitempty" xml:"LimitEphemeralStorage,omitempty"`
	// The maximum size of memory allowed for each application instance when the application is running. Unit: MB. The value of LimitMem must be greater than or equal to that of RequestsMem.
	//
	// example:
	//
	// 4096
	LimitMem *int32 `json:"LimitMem,omitempty" xml:"LimitMem,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is running. Unit: millicores. The value 0 indicates that no limit is set on CPU cores.
	//
	// example:
	//
	// 1000
	LimitmCpu *int32 `json:"LimitmCpu,omitempty" xml:"LimitmCpu,omitempty"`
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
	// The namespace of the Kubernetes cluster. This parameter specifies the Kubernetes namespace in which your application is deployed. By default, the default namespace is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The URL of the deployment package.
	//
	// example:
	//
	// https://e***.oss-cn-beijing.aliyuncs.com/s***-1.0-SNAPSHOT-spring-boot.jar
	PackageUrl *string `json:"PackageUrl,omitempty" xml:"PackageUrl,omitempty"`
	// The configuration for mounting a Kubernetes PersistentVolumeClaim (PVC) to a directory in an elastic container instance. The following parameters are included in the configuration:
	//
	// 	- pvcName: the name of the PVC. Make sure that the volume exists and is in the Bound state.
	//
	// 	- mountPaths: the directory to which you want to mount the PVC. You can configure multiple directories. You can set the following two parameters for each mount directory:
	//
	//     	- mountPath: the mount path. The mount path must be an absolute path that starts with a forward slash (/).
	//
	//     	- readOnly: the mount mode. The value true indicates the read-only mode. The value false indicates the read and write mode. Default value: false.
	//
	// example:
	//
	// [{"pvcName":"nas-pvc-1","mountPaths":[{"mountPath":"/usr/share/nginx/data"},{"mountPath":"/usr/share/nginx/html","readOnly":true}]}]
	PvcMountDescs *string `json:"PvcMountDescs,omitempty" xml:"PvcMountDescs,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// 2
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// The minimum size of space required by ephemeral storage. Unit: GB. The value 0 indicates that no limit is set on the ephemeral storage space.
	//
	// example:
	//
	// 2
	RequestsEphemeralStorage *int32 `json:"RequestsEphemeralStorage,omitempty" xml:"RequestsEphemeralStorage,omitempty"`
	// The maximum size of memory allowed for each application instance when the application is created. Unit: MB. The value 0 indicates that no limit is set on the memory size. The value of RequestsMem cannot be greater than that of LimitMem.
	//
	// example:
	//
	// 1024
	RequestsMem *int32 `json:"RequestsMem,omitempty" xml:"RequestsMem,omitempty"`
	// The maximum number of CPU cores allowed for each application instance when the application is created. Unit: millicores.
	//
	// example:
	//
	// 500
	RequestsmCpu *int32 `json:"RequestsmCpu,omitempty" xml:"RequestsmCpu,omitempty"`
}

func (s StartK8sAppPrecheckRequest) String() string {
	return dara.Prettify(s)
}

func (s StartK8sAppPrecheckRequest) GoString() string {
	return s.String()
}

func (s *StartK8sAppPrecheckRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *StartK8sAppPrecheckRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartK8sAppPrecheckRequest) GetAppName() *string {
	return s.AppName
}

func (s *StartK8sAppPrecheckRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *StartK8sAppPrecheckRequest) GetComponentIds() *string {
	return s.ComponentIds
}

func (s *StartK8sAppPrecheckRequest) GetConfigMountDescs() *string {
	return s.ConfigMountDescs
}

func (s *StartK8sAppPrecheckRequest) GetEmptyDirs() *string {
	return s.EmptyDirs
}

func (s *StartK8sAppPrecheckRequest) GetEnvFroms() *string {
	return s.EnvFroms
}

func (s *StartK8sAppPrecheckRequest) GetEnvs() *string {
	return s.Envs
}

func (s *StartK8sAppPrecheckRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *StartK8sAppPrecheckRequest) GetJavaStartUpConfig() *string {
	return s.JavaStartUpConfig
}

func (s *StartK8sAppPrecheckRequest) GetLabels() *string {
	return s.Labels
}

func (s *StartK8sAppPrecheckRequest) GetLimitEphemeralStorage() *int32 {
	return s.LimitEphemeralStorage
}

func (s *StartK8sAppPrecheckRequest) GetLimitMem() *int32 {
	return s.LimitMem
}

func (s *StartK8sAppPrecheckRequest) GetLimitmCpu() *int32 {
	return s.LimitmCpu
}

func (s *StartK8sAppPrecheckRequest) GetLocalVolume() *string {
	return s.LocalVolume
}

func (s *StartK8sAppPrecheckRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *StartK8sAppPrecheckRequest) GetPackageUrl() *string {
	return s.PackageUrl
}

func (s *StartK8sAppPrecheckRequest) GetPvcMountDescs() *string {
	return s.PvcMountDescs
}

func (s *StartK8sAppPrecheckRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartK8sAppPrecheckRequest) GetReplicas() *int32 {
	return s.Replicas
}

func (s *StartK8sAppPrecheckRequest) GetRequestsEphemeralStorage() *int32 {
	return s.RequestsEphemeralStorage
}

func (s *StartK8sAppPrecheckRequest) GetRequestsMem() *int32 {
	return s.RequestsMem
}

func (s *StartK8sAppPrecheckRequest) GetRequestsmCpu() *int32 {
	return s.RequestsmCpu
}

func (s *StartK8sAppPrecheckRequest) SetAnnotations(v string) *StartK8sAppPrecheckRequest {
	s.Annotations = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetAppId(v string) *StartK8sAppPrecheckRequest {
	s.AppId = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetAppName(v string) *StartK8sAppPrecheckRequest {
	s.AppName = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetClusterId(v string) *StartK8sAppPrecheckRequest {
	s.ClusterId = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetComponentIds(v string) *StartK8sAppPrecheckRequest {
	s.ComponentIds = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetConfigMountDescs(v string) *StartK8sAppPrecheckRequest {
	s.ConfigMountDescs = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetEmptyDirs(v string) *StartK8sAppPrecheckRequest {
	s.EmptyDirs = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetEnvFroms(v string) *StartK8sAppPrecheckRequest {
	s.EnvFroms = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetEnvs(v string) *StartK8sAppPrecheckRequest {
	s.Envs = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetImageUrl(v string) *StartK8sAppPrecheckRequest {
	s.ImageUrl = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetJavaStartUpConfig(v string) *StartK8sAppPrecheckRequest {
	s.JavaStartUpConfig = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetLabels(v string) *StartK8sAppPrecheckRequest {
	s.Labels = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetLimitEphemeralStorage(v int32) *StartK8sAppPrecheckRequest {
	s.LimitEphemeralStorage = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetLimitMem(v int32) *StartK8sAppPrecheckRequest {
	s.LimitMem = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetLimitmCpu(v int32) *StartK8sAppPrecheckRequest {
	s.LimitmCpu = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetLocalVolume(v string) *StartK8sAppPrecheckRequest {
	s.LocalVolume = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetNamespace(v string) *StartK8sAppPrecheckRequest {
	s.Namespace = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetPackageUrl(v string) *StartK8sAppPrecheckRequest {
	s.PackageUrl = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetPvcMountDescs(v string) *StartK8sAppPrecheckRequest {
	s.PvcMountDescs = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetRegionId(v string) *StartK8sAppPrecheckRequest {
	s.RegionId = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetReplicas(v int32) *StartK8sAppPrecheckRequest {
	s.Replicas = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetRequestsEphemeralStorage(v int32) *StartK8sAppPrecheckRequest {
	s.RequestsEphemeralStorage = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetRequestsMem(v int32) *StartK8sAppPrecheckRequest {
	s.RequestsMem = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) SetRequestsmCpu(v int32) *StartK8sAppPrecheckRequest {
	s.RequestsmCpu = &v
	return s
}

func (s *StartK8sAppPrecheckRequest) Validate() error {
	return dara.Validate(s)
}
