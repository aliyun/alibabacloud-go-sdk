// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterQueueInfoByEnvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetClusterQueueInfoByEnvResponseBody
	GetCode() *string
	SetData(v []*GetClusterQueueInfoByEnvResponseBodyData) *GetClusterQueueInfoByEnvResponseBody
	GetData() []*GetClusterQueueInfoByEnvResponseBodyData
	SetHttpStatusCode(v int32) *GetClusterQueueInfoByEnvResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetClusterQueueInfoByEnvResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetClusterQueueInfoByEnvResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetClusterQueueInfoByEnvResponseBody
	GetSuccess() *bool
}

type GetClusterQueueInfoByEnvResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*GetClusterQueueInfoByEnvResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetClusterQueueInfoByEnvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterQueueInfoByEnvResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterQueueInfoByEnvResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetClusterQueueInfoByEnvResponseBody) GetData() []*GetClusterQueueInfoByEnvResponseBodyData {
	return s.Data
}

func (s *GetClusterQueueInfoByEnvResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetClusterQueueInfoByEnvResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClusterQueueInfoByEnvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterQueueInfoByEnvResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetClusterQueueInfoByEnvResponseBody) SetCode(v string) *GetClusterQueueInfoByEnvResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBody) SetData(v []*GetClusterQueueInfoByEnvResponseBodyData) *GetClusterQueueInfoByEnvResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBody) SetHttpStatusCode(v int32) *GetClusterQueueInfoByEnvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBody) SetMessage(v string) *GetClusterQueueInfoByEnvResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBody) SetRequestId(v string) *GetClusterQueueInfoByEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBody) SetSuccess(v bool) *GetClusterQueueInfoByEnvResponseBody {
	s.Success = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBody) Validate() error {
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

type GetClusterQueueInfoByEnvResponseBodyData struct {
	// The annotations.
	//
	// example:
	//
	// {  "creator": "new_datasource@test.aliyunid.com",  "modifier": "new_datasource@test.aliyunid.com" }
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 61187014-a3ba-4cdd-8609-1f0aa3df4a3d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-31 10:29:17
	CreateAt *string `json:"CreateAt,omitempty" xml:"CreateAt,omitempty"`
	// The Flink image registry address.
	//
	// example:
	//
	// xxxx-registry-vpc.cn-shanghai.cr.aliyuncs.com/xxxx/flink:1.15.4-scala_2.12
	FlinkImageRegistry *string `json:"FlinkImageRegistry,omitempty" xml:"FlinkImageRegistry,omitempty"`
	// The Flink image repository.
	//
	// example:
	//
	// xxxx-registry-vpc.cn-shanghai.cr.aliyuncs.com
	FlinkImageRepository *string `json:"FlinkImageRepository,omitempty" xml:"FlinkImageRepository,omitempty"`
	// The Flink image tag.
	//
	// example:
	//
	// 1.15.4
	FlinkImageTag *string `json:"FlinkImageTag,omitempty" xml:"FlinkImageTag,omitempty"`
	// The Flink version.
	//
	// example:
	//
	// 1.15
	FlinkVersion *string `json:"FlinkVersion,omitempty" xml:"FlinkVersion,omitempty"`
	// The labels.
	//
	// example:
	//
	// cdh
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The maximum number of vCores provided by each queue.
	//
	// example:
	//
	// 10
	MaxVcore *string `json:"MaxVcore,omitempty" xml:"MaxVcore,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-10-31 10:29:17
	ModifiedAt *string `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	// The namespace.
	//
	// example:
	//
	// dataphinv45prod
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The queue name, which serves as a unique identifier. For Kubernetes, this is the resource group name.
	//
	// example:
	//
	// default-queue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The resource version.
	//
	// example:
	//
	// 1
	ResourceVersion *string `json:"ResourceVersion,omitempty" xml:"ResourceVersion,omitempty"`
	// The resource specification.
	//
	// example:
	//
	// {\\n  \\"kind\\" : \\"kubernetes\\",\\n  \\"kubernetes\\{"namespace" : "n1730341728989z7",    "clusterName" : "a51578bdcce145"  },  "state" : "ONLINE"}
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The VVP cluster type.
	//
	// example:
	//
	// PREJOB
	VvpClusterType *string `json:"VvpClusterType,omitempty" xml:"VvpClusterType,omitempty"`
}

func (s GetClusterQueueInfoByEnvResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetClusterQueueInfoByEnvResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetAnnotations() *string {
	return s.Annotations
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetCreateAt() *string {
	return s.CreateAt
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetFlinkImageRegistry() *string {
	return s.FlinkImageRegistry
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetFlinkImageRepository() *string {
	return s.FlinkImageRepository
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetFlinkImageTag() *string {
	return s.FlinkImageTag
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetFlinkVersion() *string {
	return s.FlinkVersion
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetMaxVcore() *string {
	return s.MaxVcore
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetQueueName() *string {
	return s.QueueName
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetResourceVersion() *string {
	return s.ResourceVersion
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetSpec() *string {
	return s.Spec
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) GetVvpClusterType() *string {
	return s.VvpClusterType
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetAnnotations(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.Annotations = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetClusterId(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetCreateAt(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.CreateAt = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetFlinkImageRegistry(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.FlinkImageRegistry = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetFlinkImageRepository(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.FlinkImageRepository = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetFlinkImageTag(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.FlinkImageTag = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetFlinkVersion(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.FlinkVersion = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetLabels(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.Labels = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetMaxVcore(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.MaxVcore = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetModifiedAt(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.ModifiedAt = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetNamespace(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetQueueName(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetResourceVersion(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.ResourceVersion = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetSpec(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.Spec = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) SetVvpClusterType(v string) *GetClusterQueueInfoByEnvResponseBodyData {
	s.VvpClusterType = &v
	return s
}

func (s *GetClusterQueueInfoByEnvResponseBodyData) Validate() error {
	return dara.Validate(s)
}
