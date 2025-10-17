// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivyComputeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLivyComputeResponseBody
	GetCode() *string
	SetData(v *GetLivyComputeResponseBodyData) *GetLivyComputeResponseBody
	GetData() *GetLivyComputeResponseBodyData
	SetMessage(v string) *GetLivyComputeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLivyComputeResponseBody
	GetRequestId() *string
}

type GetLivyComputeResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetLivyComputeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLivyComputeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLivyComputeResponseBody) GetData() *GetLivyComputeResponseBodyData {
	return s.Data
}

func (s *GetLivyComputeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLivyComputeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLivyComputeResponseBody) SetCode(v string) *GetLivyComputeResponseBody {
	s.Code = &v
	return s
}

func (s *GetLivyComputeResponseBody) SetData(v *GetLivyComputeResponseBodyData) *GetLivyComputeResponseBody {
	s.Data = v
	return s
}

func (s *GetLivyComputeResponseBody) SetMessage(v string) *GetLivyComputeResponseBody {
	s.Message = &v
	return s
}

func (s *GetLivyComputeResponseBody) SetRequestId(v string) *GetLivyComputeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLivyComputeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLivyComputeResponseBodyData struct {
	// example:
	//
	// Token
	AuthType              *string                                              `json:"authType,omitempty" xml:"authType,omitempty"`
	AutoStopConfiguration *GetLivyComputeResponseBodyDataAutoStopConfiguration `json:"autoStopConfiguration,omitempty" xml:"autoStopConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// lc-xxxxxxxxxxxxx
	ComputeId *string `json:"computeId,omitempty" xml:"computeId,omitempty"`
	// example:
	//
	// 1
	CpuLimit *string `json:"cpuLimit,omitempty" xml:"cpuLimit,omitempty"`
	// example:
	//
	// alice
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// example:
	//
	// true
	EnablePublic *bool `json:"enablePublic,omitempty" xml:"enablePublic,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou.data.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// emr-spark-livy-gateway-cn-hangzhou-internal.aliyun.com/api/v1/workspace/w-xxxxxxxxx/livycompute/lc-xxxxxxxxxxx
	EndpointInner *string `json:"endpointInner,omitempty" xml:"endpointInner,omitempty"`
	// example:
	//
	// ev-cq31c7tlhtgm9nrrlj4g
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// example:
	//
	// 1749456094000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// {
	//
	//   "sparkDefaultsConf": "spark.driver.cores     1\\nspark.driver.memory    4g\\nspark.executor.cores   1\\nspark.executor.memory  4g\\n",
	//
	//   "sparkBlackListConf": "spark.driver.cores\\nspark.driver.memory",
	//
	//   "livyConf": "livy.server.session.timeout  1h\\n",
	//
	//   "livyClientConf": "livy.rsc.sql.num-rows  1000\\n"
	//
	// }
	LivyServerConf *string `json:"livyServerConf,omitempty" xml:"livyServerConf,omitempty"`
	// example:
	//
	// 0.8.0
	LivyVersion *string `json:"livyVersion,omitempty" xml:"livyVersion,omitempty"`
	// example:
	//
	// 4Gi
	MemoryLimit *string `json:"memoryLimit,omitempty" xml:"memoryLimit,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// test
	NetworkName *string `json:"networkName,omitempty" xml:"networkName,omitempty"`
	// example:
	//
	// root_queue
	QueueName *string `json:"queueName,omitempty" xml:"queueName,omitempty"`
	// example:
	//
	// 10000001
	RamUserId *string `json:"ramUserId,omitempty" xml:"ramUserId,omitempty"`
	// example:
	//
	// esr-4.3.0 (Spark 3.5.2, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// example:
	//
	// 1749456094000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetLivyComputeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponseBodyData) GetAuthType() *string {
	return s.AuthType
}

func (s *GetLivyComputeResponseBodyData) GetAutoStopConfiguration() *GetLivyComputeResponseBodyDataAutoStopConfiguration {
	return s.AutoStopConfiguration
}

func (s *GetLivyComputeResponseBodyData) GetComputeId() *string {
	return s.ComputeId
}

func (s *GetLivyComputeResponseBodyData) GetCpuLimit() *string {
	return s.CpuLimit
}

func (s *GetLivyComputeResponseBodyData) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetLivyComputeResponseBodyData) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *GetLivyComputeResponseBodyData) GetEnablePublic() *bool {
	return s.EnablePublic
}

func (s *GetLivyComputeResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetLivyComputeResponseBodyData) GetEndpointInner() *string {
	return s.EndpointInner
}

func (s *GetLivyComputeResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetLivyComputeResponseBodyData) GetFusion() *bool {
	return s.Fusion
}

func (s *GetLivyComputeResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetLivyComputeResponseBodyData) GetLivyServerConf() *string {
	return s.LivyServerConf
}

func (s *GetLivyComputeResponseBodyData) GetLivyVersion() *string {
	return s.LivyVersion
}

func (s *GetLivyComputeResponseBodyData) GetMemoryLimit() *string {
	return s.MemoryLimit
}

func (s *GetLivyComputeResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetLivyComputeResponseBodyData) GetNetworkName() *string {
	return s.NetworkName
}

func (s *GetLivyComputeResponseBodyData) GetQueueName() *string {
	return s.QueueName
}

func (s *GetLivyComputeResponseBodyData) GetRamUserId() *string {
	return s.RamUserId
}

func (s *GetLivyComputeResponseBodyData) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *GetLivyComputeResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetLivyComputeResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetLivyComputeResponseBodyData) SetAuthType(v string) *GetLivyComputeResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetAutoStopConfiguration(v *GetLivyComputeResponseBodyDataAutoStopConfiguration) *GetLivyComputeResponseBodyData {
	s.AutoStopConfiguration = v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetComputeId(v string) *GetLivyComputeResponseBodyData {
	s.ComputeId = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetCpuLimit(v string) *GetLivyComputeResponseBodyData {
	s.CpuLimit = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetCreatedBy(v string) *GetLivyComputeResponseBodyData {
	s.CreatedBy = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetDisplayReleaseVersion(v string) *GetLivyComputeResponseBodyData {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEnablePublic(v bool) *GetLivyComputeResponseBodyData {
	s.EnablePublic = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEndpoint(v string) *GetLivyComputeResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEndpointInner(v string) *GetLivyComputeResponseBodyData {
	s.EndpointInner = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetEnvironmentId(v string) *GetLivyComputeResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetFusion(v bool) *GetLivyComputeResponseBodyData {
	s.Fusion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetGmtCreate(v int64) *GetLivyComputeResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetLivyServerConf(v string) *GetLivyComputeResponseBodyData {
	s.LivyServerConf = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetLivyVersion(v string) *GetLivyComputeResponseBodyData {
	s.LivyVersion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetMemoryLimit(v string) *GetLivyComputeResponseBodyData {
	s.MemoryLimit = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetName(v string) *GetLivyComputeResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetNetworkName(v string) *GetLivyComputeResponseBodyData {
	s.NetworkName = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetQueueName(v string) *GetLivyComputeResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetRamUserId(v string) *GetLivyComputeResponseBodyData {
	s.RamUserId = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetReleaseVersion(v string) *GetLivyComputeResponseBodyData {
	s.ReleaseVersion = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetStartTime(v int64) *GetLivyComputeResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) SetStatus(v string) *GetLivyComputeResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetLivyComputeResponseBodyData) Validate() error {
	if s.AutoStopConfiguration != nil {
		if err := s.AutoStopConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLivyComputeResponseBodyDataAutoStopConfiguration struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 300
	IdleTimeoutMinutes *int64 `json:"idleTimeoutMinutes,omitempty" xml:"idleTimeoutMinutes,omitempty"`
}

func (s GetLivyComputeResponseBodyDataAutoStopConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetLivyComputeResponseBodyDataAutoStopConfiguration) GoString() string {
	return s.String()
}

func (s *GetLivyComputeResponseBodyDataAutoStopConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *GetLivyComputeResponseBodyDataAutoStopConfiguration) GetIdleTimeoutMinutes() *int64 {
	return s.IdleTimeoutMinutes
}

func (s *GetLivyComputeResponseBodyDataAutoStopConfiguration) SetEnable(v bool) *GetLivyComputeResponseBodyDataAutoStopConfiguration {
	s.Enable = &v
	return s
}

func (s *GetLivyComputeResponseBodyDataAutoStopConfiguration) SetIdleTimeoutMinutes(v int64) *GetLivyComputeResponseBodyDataAutoStopConfiguration {
	s.IdleTimeoutMinutes = &v
	return s
}

func (s *GetLivyComputeResponseBodyDataAutoStopConfiguration) Validate() error {
	return dara.Validate(s)
}
