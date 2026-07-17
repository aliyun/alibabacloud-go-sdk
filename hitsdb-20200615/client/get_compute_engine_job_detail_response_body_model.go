// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeEngineJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetComputeEngineJobDetailResponseBody
	GetAccessDeniedDetail() *string
	SetConfigs(v map[string]interface{}) *GetComputeEngineJobDetailResponseBody
	GetConfigs() map[string]interface{}
	SetCreateTime(v string) *GetComputeEngineJobDetailResponseBody
	GetCreateTime() *string
	SetEndpoint(v string) *GetComputeEngineJobDetailResponseBody
	GetEndpoint() *string
	SetExtraInfo(v map[string]interface{}) *GetComputeEngineJobDetailResponseBody
	GetExtraInfo() map[string]interface{}
	SetFinishTime(v string) *GetComputeEngineJobDetailResponseBody
	GetFinishTime() *string
	SetJobId(v string) *GetComputeEngineJobDetailResponseBody
	GetJobId() *string
	SetJobName(v string) *GetComputeEngineJobDetailResponseBody
	GetJobName() *string
	SetJobType(v string) *GetComputeEngineJobDetailResponseBody
	GetJobType() *string
	SetLastErrorCode(v string) *GetComputeEngineJobDetailResponseBody
	GetLastErrorCode() *string
	SetLastErrorInfo(v string) *GetComputeEngineJobDetailResponseBody
	GetLastErrorInfo() *string
	SetRequestId(v string) *GetComputeEngineJobDetailResponseBody
	GetRequestId() *string
	SetState(v string) *GetComputeEngineJobDetailResponseBody
	GetState() *string
}

type GetComputeEngineJobDetailResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Configs            map[string]interface{} `json:"Configs,omitempty" xml:"Configs,omitempty"`
	CreateTime         *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Endpoint           *string                `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ExtraInfo          map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	FinishTime         *string                `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	JobId              *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName            *string                `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobType            *string                `json:"JobType,omitempty" xml:"JobType,omitempty"`
	LastErrorCode      *string                `json:"LastErrorCode,omitempty" xml:"LastErrorCode,omitempty"`
	LastErrorInfo      *string                `json:"LastErrorInfo,omitempty" xml:"LastErrorInfo,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State              *string                `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetComputeEngineJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeEngineJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeEngineJobDetailResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetComputeEngineJobDetailResponseBody) GetConfigs() map[string]interface{} {
	return s.Configs
}

func (s *GetComputeEngineJobDetailResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetComputeEngineJobDetailResponseBody) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetComputeEngineJobDetailResponseBody) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *GetComputeEngineJobDetailResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetComputeEngineJobDetailResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetComputeEngineJobDetailResponseBody) GetJobName() *string {
	return s.JobName
}

func (s *GetComputeEngineJobDetailResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *GetComputeEngineJobDetailResponseBody) GetLastErrorCode() *string {
	return s.LastErrorCode
}

func (s *GetComputeEngineJobDetailResponseBody) GetLastErrorInfo() *string {
	return s.LastErrorInfo
}

func (s *GetComputeEngineJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeEngineJobDetailResponseBody) GetState() *string {
	return s.State
}

func (s *GetComputeEngineJobDetailResponseBody) SetAccessDeniedDetail(v string) *GetComputeEngineJobDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetConfigs(v map[string]interface{}) *GetComputeEngineJobDetailResponseBody {
	s.Configs = v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetCreateTime(v string) *GetComputeEngineJobDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetEndpoint(v string) *GetComputeEngineJobDetailResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetExtraInfo(v map[string]interface{}) *GetComputeEngineJobDetailResponseBody {
	s.ExtraInfo = v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetFinishTime(v string) *GetComputeEngineJobDetailResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetJobId(v string) *GetComputeEngineJobDetailResponseBody {
	s.JobId = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetJobName(v string) *GetComputeEngineJobDetailResponseBody {
	s.JobName = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetJobType(v string) *GetComputeEngineJobDetailResponseBody {
	s.JobType = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetLastErrorCode(v string) *GetComputeEngineJobDetailResponseBody {
	s.LastErrorCode = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetLastErrorInfo(v string) *GetComputeEngineJobDetailResponseBody {
	s.LastErrorInfo = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetRequestId(v string) *GetComputeEngineJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) SetState(v string) *GetComputeEngineJobDetailResponseBody {
	s.State = &v
	return s
}

func (s *GetComputeEngineJobDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
