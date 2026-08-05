// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFunctionInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *GetFunctionInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *GetFunctionInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *GetFunctionInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFunctionInstanceResponseBody
	GetRequestId() *string
	SetResult(v *GetFunctionInstanceResponseBodyResult) *GetFunctionInstanceResponseBody
	GetResult() *GetFunctionInstanceResponseBodyResult
	SetStatus(v string) *GetFunctionInstanceResponseBody
	GetStatus() *string
}

type GetFunctionInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// not found
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The time consumed.
	//
	// example:
	//
	// 11.627
	Latency *int64 `json:"latency,omitempty" xml:"latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// "xx not found"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C56462F4-CCB3-10BF-A3D8-FEE53C72B65C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *GetFunctionInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetFunctionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFunctionInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *GetFunctionInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *GetFunctionInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFunctionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFunctionInstanceResponseBody) GetResult() *GetFunctionInstanceResponseBodyResult {
	return s.Result
}

func (s *GetFunctionInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionInstanceResponseBody) SetCode(v string) *GetFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetHttpCode(v int64) *GetFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetLatency(v int64) *GetFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetMessage(v string) *GetFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetRequestId(v string) *GetFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetResult(v *GetFunctionInstanceResponseBodyResult) *GetFunctionInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetStatus(v string) *GetFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionInstanceResponseBodyResult struct {
	// The ownership information.
	Belongs *GetFunctionInstanceResponseBodyResultBelongs `json:"belongs,omitempty" xml:"belongs,omitempty" type:"Struct"`
	// The specific configuration items.
	CreateParameters []*GetFunctionInstanceResponseBodyResultCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 1724998630466
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The cron expression for the timed scheduling task.
	//
	// example:
	//
	// ""
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The extended information.
	//
	// example:
	//
	// ""
	ExtendInfo *string `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	// The configuration type. Valid values:
	//
	// - nl2sql
	//
	// - embedding-tuning
	//
	// - deployment
	//
	// - notebook.
	//
	// example:
	//
	// nl2sql
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// The configuration type. PAAS (default): requires training before use.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The configuration name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The model type. The valid values vary based on the configuration type (functionName):
	//
	// - ops-query-analyze-nl2sql-001 (nl2sql)
	//
	// - ops-embedding-dim-reduction-001 (embedding-tuning)
	//
	// - native (deployment)
	//
	// - dsw (notebook).
	//
	// example:
	//
	// dsw
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The source.
	//
	// example:
	//
	// user
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The status. Valid values:
	//
	// - available
	//
	// - unavailable.
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task information.
	Task *GetFunctionInstanceResponseBodyResultTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
	// The training version ID.
	//
	// example:
	//
	// 21
	VersionId *int64 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResult) GetBelongs() *GetFunctionInstanceResponseBodyResultBelongs {
	return s.Belongs
}

func (s *GetFunctionInstanceResponseBodyResult) GetCreateParameters() []*GetFunctionInstanceResponseBodyResultCreateParameters {
	return s.CreateParameters
}

func (s *GetFunctionInstanceResponseBodyResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetFunctionInstanceResponseBodyResult) GetCron() *string {
	return s.Cron
}

func (s *GetFunctionInstanceResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetFunctionInstanceResponseBodyResult) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetFunctionInstanceResponseBodyResult) GetFunctionName() *string {
	return s.FunctionName
}

func (s *GetFunctionInstanceResponseBodyResult) GetFunctionType() *string {
	return s.FunctionType
}

func (s *GetFunctionInstanceResponseBodyResult) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetFunctionInstanceResponseBodyResult) GetModelType() *string {
	return s.ModelType
}

func (s *GetFunctionInstanceResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *GetFunctionInstanceResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetFunctionInstanceResponseBodyResult) GetTask() *GetFunctionInstanceResponseBodyResultTask {
	return s.Task
}

func (s *GetFunctionInstanceResponseBodyResult) GetVersionId() *int64 {
	return s.VersionId
}

func (s *GetFunctionInstanceResponseBodyResult) SetBelongs(v *GetFunctionInstanceResponseBodyResultBelongs) *GetFunctionInstanceResponseBodyResult {
	s.Belongs = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateParameters(v []*GetFunctionInstanceResponseBodyResultCreateParameters) *GetFunctionInstanceResponseBodyResult {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateTime(v int64) *GetFunctionInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCron(v string) *GetFunctionInstanceResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetDescription(v string) *GetFunctionInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetExtendInfo(v string) *GetFunctionInstanceResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionName(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionType(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetInstanceName(v string) *GetFunctionInstanceResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetModelType(v string) *GetFunctionInstanceResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetSource(v string) *GetFunctionInstanceResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetStatus(v string) *GetFunctionInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetTask(v *GetFunctionInstanceResponseBodyResultTask) *GetFunctionInstanceResponseBodyResult {
	s.Task = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetVersionId(v int64) *GetFunctionInstanceResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) Validate() error {
	if s.Belongs != nil {
		if err := s.Belongs.Validate(); err != nil {
			return err
		}
	}
	if s.CreateParameters != nil {
		for _, item := range s.CreateParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFunctionInstanceResponseBodyResultBelongs struct {
	// The category.
	//
	// example:
	//
	// ""
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The industry type.
	//
	// example:
	//
	// ""
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The language.
	//
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultBelongs) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultBelongs) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) GetCategory() *string {
	return s.Category
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) GetDomain() *string {
	return s.Domain
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) GetLanguage() *string {
	return s.Language
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetCategory(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Category = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetDomain(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Domain = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetLanguage(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Language = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) Validate() error {
	return dara.Validate(s)
}

type GetFunctionInstanceResponseBodyResultCreateParameters struct {
	// The parameter name.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) GetName() *string {
	return s.Name
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) GetValue() *string {
	return s.Value
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetName(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetValue(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Value = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) Validate() error {
	return dara.Validate(s)
}

type GetFunctionInstanceResponseBodyResultTask struct {
	// The task status. Valid values:
	//
	// - success: Succeeded.
	//
	// - failed: Failed.
	//
	// - untrained: Pending training.
	//
	// - pending: Scheduling.
	//
	// - running: Training in progress.
	//
	// example:
	//
	// success
	DagStatus *string `json:"dagStatus,omitempty" xml:"dagStatus,omitempty"`
	// The last training time.
	//
	// example:
	//
	// 1724998630466
	LastRunTime *int64 `json:"lastRunTime,omitempty" xml:"lastRunTime,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultTask) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultTask) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultTask) GetDagStatus() *string {
	return s.DagStatus
}

func (s *GetFunctionInstanceResponseBodyResultTask) GetLastRunTime() *int64 {
	return s.LastRunTime
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetDagStatus(v string) *GetFunctionInstanceResponseBodyResultTask {
	s.DagStatus = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetLastRunTime(v int64) *GetFunctionInstanceResponseBodyResultTask {
	s.LastRunTime = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultTask) Validate() error {
	return dara.Validate(s)
}
