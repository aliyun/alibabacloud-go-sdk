// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStreamJobsResponseBody
	GetCode() *string
	SetData(v []*GetStreamJobsResponseBodyData) *GetStreamJobsResponseBody
	GetData() []*GetStreamJobsResponseBodyData
	SetHttpStatusCode(v int32) *GetStreamJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStreamJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStreamJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStreamJobsResponseBody
	GetSuccess() *bool
}

type GetStreamJobsResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetStreamJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStreamJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStreamJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetStreamJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStreamJobsResponseBody) GetData() []*GetStreamJobsResponseBodyData {
	return s.Data
}

func (s *GetStreamJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStreamJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStreamJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStreamJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStreamJobsResponseBody) SetCode(v string) *GetStreamJobsResponseBody {
	s.Code = &v
	return s
}

func (s *GetStreamJobsResponseBody) SetData(v []*GetStreamJobsResponseBodyData) *GetStreamJobsResponseBody {
	s.Data = v
	return s
}

func (s *GetStreamJobsResponseBody) SetHttpStatusCode(v int32) *GetStreamJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStreamJobsResponseBody) SetMessage(v string) *GetStreamJobsResponseBody {
	s.Message = &v
	return s
}

func (s *GetStreamJobsResponseBody) SetRequestId(v string) *GetStreamJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStreamJobsResponseBody) SetSuccess(v bool) *GetStreamJobsResponseBody {
	s.Success = &v
	return s
}

func (s *GetStreamJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStreamJobsResponseBodyData struct {
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 716555
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// flink_order_detail
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 300006788
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// 300006788
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 7162269257990111
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// FLINK_SQL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStreamJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStreamJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStreamJobsResponseBodyData) GetEnv() *string {
	return s.Env
}

func (s *GetStreamJobsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetStreamJobsResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *GetStreamJobsResponseBodyData) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetStreamJobsResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetStreamJobsResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetStreamJobsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetStreamJobsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetStreamJobsResponseBodyData) SetEnv(v string) *GetStreamJobsResponseBodyData {
	s.Env = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) SetId(v int64) *GetStreamJobsResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) SetJobName(v string) *GetStreamJobsResponseBodyData {
	s.JobName = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) SetModifierId(v string) *GetStreamJobsResponseBodyData {
	s.ModifierId = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) SetOwnerId(v string) *GetStreamJobsResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) SetProjectId(v string) *GetStreamJobsResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) SetStatus(v string) *GetStreamJobsResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) SetType(v string) *GetStreamJobsResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetStreamJobsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
