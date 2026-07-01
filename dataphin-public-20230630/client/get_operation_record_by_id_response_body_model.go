// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOperationRecordByIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetOperationRecordByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOperationRecordByIdResponseBody
	GetMessage() *string
	SetOperationLogDTO(v *GetOperationRecordByIdResponseBodyOperationLogDTO) *GetOperationRecordByIdResponseBody
	GetOperationLogDTO() *GetOperationRecordByIdResponseBodyOperationLogDTO
	SetRequestId(v string) *GetOperationRecordByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOperationRecordByIdResponseBody
	GetSuccess() *bool
}

type GetOperationRecordByIdResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message         *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	OperationLogDTO *GetOperationRecordByIdResponseBodyOperationLogDTO `json:"OperationLogDTO,omitempty" xml:"OperationLogDTO,omitempty" type:"Struct"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOperationRecordByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationRecordByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOperationRecordByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOperationRecordByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOperationRecordByIdResponseBody) GetOperationLogDTO() *GetOperationRecordByIdResponseBodyOperationLogDTO {
	return s.OperationLogDTO
}

func (s *GetOperationRecordByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationRecordByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOperationRecordByIdResponseBody) SetCode(v string) *GetOperationRecordByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetOperationRecordByIdResponseBody) SetHttpStatusCode(v int32) *GetOperationRecordByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOperationRecordByIdResponseBody) SetMessage(v string) *GetOperationRecordByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetOperationRecordByIdResponseBody) SetOperationLogDTO(v *GetOperationRecordByIdResponseBodyOperationLogDTO) *GetOperationRecordByIdResponseBody {
	s.OperationLogDTO = v
	return s
}

func (s *GetOperationRecordByIdResponseBody) SetRequestId(v string) *GetOperationRecordByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationRecordByIdResponseBody) SetSuccess(v bool) *GetOperationRecordByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetOperationRecordByIdResponseBody) Validate() error {
	if s.OperationLogDTO != nil {
		if err := s.OperationLogDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationRecordByIdResponseBodyOperationLogDTO struct {
	// example:
	//
	// 2025-01-15 10:30:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 0
	CodeType *int32 `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// example:
	//
	// 120
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 测试任务
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// onedata-ide
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 987654321
	OperationId *int64 `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// example:
	//
	// 131211211
	ProjectId      *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RelationTables []*string `json:"RelationTables,omitempty" xml:"RelationTables,omitempty" type:"Repeated"`
	// example:
	//
	// 30231123
	Runner *string `json:"Runner,omitempty" xml:"Runner,omitempty"`
	// example:
	//
	// 张三
	RunnerName *string `json:"RunnerName,omitempty" xml:"RunnerName,omitempty"`
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetOperationRecordByIdResponseBodyOperationLogDTO) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordByIdResponseBodyOperationLogDTO) GoString() string {
	return s.String()
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetCodeType() *int32 {
	return s.CodeType
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetDuration() *int64 {
	return s.Duration
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetId() *int64 {
	return s.Id
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetName() *string {
	return s.Name
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetOperationId() *int64 {
	return s.OperationId
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetRelationTables() []*string {
	return s.RelationTables
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetRunner() *string {
	return s.Runner
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetRunnerName() *string {
	return s.RunnerName
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetStatus() *int32 {
	return s.Status
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetBeginTime(v string) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.BeginTime = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetCodeType(v int32) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.CodeType = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetDuration(v int64) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.Duration = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetId(v int64) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.Id = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetName(v string) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.Name = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetObjectType(v string) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.ObjectType = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetOperationId(v int64) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.OperationId = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetProjectId(v int64) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.ProjectId = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetRelationTables(v []*string) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.RelationTables = v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetRunner(v string) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.Runner = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetRunnerName(v string) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.RunnerName = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetStatus(v int32) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.Status = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) SetTenantId(v int64) *GetOperationRecordByIdResponseBodyOperationLogDTO {
	s.TenantId = &v
	return s
}

func (s *GetOperationRecordByIdResponseBodyOperationLogDTO) Validate() error {
	return dara.Validate(s)
}
