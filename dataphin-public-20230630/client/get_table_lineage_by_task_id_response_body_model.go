// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineageByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableLineageByTaskIdResponseBody
	GetCode() *string
	SetData(v []*GetTableLineageByTaskIdResponseBodyData) *GetTableLineageByTaskIdResponseBody
	GetData() []*GetTableLineageByTaskIdResponseBodyData
	SetHttpStatusCode(v int32) *GetTableLineageByTaskIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTableLineageByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableLineageByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableLineageByTaskIdResponseBody
	GetSuccess() *bool
}

type GetTableLineageByTaskIdResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetTableLineageByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableLineageByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineageByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableLineageByTaskIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableLineageByTaskIdResponseBody) GetData() []*GetTableLineageByTaskIdResponseBodyData {
	return s.Data
}

func (s *GetTableLineageByTaskIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTableLineageByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableLineageByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableLineageByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableLineageByTaskIdResponseBody) SetCode(v string) *GetTableLineageByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBody) SetData(v []*GetTableLineageByTaskIdResponseBodyData) *GetTableLineageByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTableLineageByTaskIdResponseBody) SetHttpStatusCode(v int32) *GetTableLineageByTaskIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBody) SetMessage(v string) *GetTableLineageByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBody) SetRequestId(v string) *GetTableLineageByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBody) SetSuccess(v bool) *GetTableLineageByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBody) Validate() error {
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

type GetTableLineageByTaskIdResponseBodyData struct {
	// example:
	//
	// 123
	InputBizUnitId *int64 `json:"InputBizUnitId,omitempty" xml:"InputBizUnitId,omitempty"`
	// example:
	//
	// 123
	InputDataSourceId *int64 `json:"InputDataSourceId,omitempty" xml:"InputDataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	InputDataSourceType *string `json:"InputDataSourceType,omitempty" xml:"InputDataSourceType,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	InputDbType *string `json:"InputDbType,omitempty" xml:"InputDbType,omitempty"`
	// example:
	//
	// DEV
	InputEnv *string `json:"InputEnv,omitempty" xml:"InputEnv,omitempty"`
	// example:
	//
	// 123
	InputProjectId    *int64 `json:"InputProjectId,omitempty" xml:"InputProjectId,omitempty"`
	InputTableDeleted *bool  `json:"InputTableDeleted,omitempty" xml:"InputTableDeleted,omitempty"`
	// example:
	//
	// odps.123.test_project.order
	InputTableId *string `json:"InputTableId,omitempty" xml:"InputTableId,omitempty"`
	// example:
	//
	// order
	InputTableName *string `json:"InputTableName,omitempty" xml:"InputTableName,omitempty"`
	// example:
	//
	// PHYSICAL_TABLE
	InputTableType *string `json:"InputTableType,omitempty" xml:"InputTableType,omitempty"`
	// example:
	//
	// 123
	OutputBizUnitId *int64 `json:"OutputBizUnitId,omitempty" xml:"OutputBizUnitId,omitempty"`
	// example:
	//
	// 123
	OutputDataSourceId *int64 `json:"OutputDataSourceId,omitempty" xml:"OutputDataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	OutputDataSourceType *string `json:"OutputDataSourceType,omitempty" xml:"OutputDataSourceType,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	OutputDbType *string `json:"OutputDbType,omitempty" xml:"OutputDbType,omitempty"`
	// example:
	//
	// DEV/PROD
	OutputEnv *string `json:"OutputEnv,omitempty" xml:"OutputEnv,omitempty"`
	// example:
	//
	// 123
	OutputProjectId    *int64 `json:"OutputProjectId,omitempty" xml:"OutputProjectId,omitempty"`
	OutputTableDeleted *bool  `json:"OutputTableDeleted,omitempty" xml:"OutputTableDeleted,omitempty"`
	// example:
	//
	// odps.123.test_project.order
	OutputTableId *string `json:"OutputTableId,omitempty" xml:"OutputTableId,omitempty"`
	// example:
	//
	// order
	OutputTableName *string `json:"OutputTableName,omitempty" xml:"OutputTableName,omitempty"`
	// example:
	//
	// PHYSICAL_TABLE
	OutputTableType *string `json:"OutputTableType,omitempty" xml:"OutputTableType,omitempty"`
	// example:
	//
	// DEV
	TaskEnv *string `json:"TaskEnv,omitempty" xml:"TaskEnv,omitempty"`
	// example:
	//
	// n_123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 12345
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetTableLineageByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineageByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputBizUnitId() *int64 {
	return s.InputBizUnitId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputDataSourceId() *int64 {
	return s.InputDataSourceId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputDataSourceType() *string {
	return s.InputDataSourceType
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputDbType() *string {
	return s.InputDbType
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputEnv() *string {
	return s.InputEnv
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputProjectId() *int64 {
	return s.InputProjectId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputTableDeleted() *bool {
	return s.InputTableDeleted
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputTableId() *string {
	return s.InputTableId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputTableName() *string {
	return s.InputTableName
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetInputTableType() *string {
	return s.InputTableType
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputBizUnitId() *int64 {
	return s.OutputBizUnitId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputDataSourceId() *int64 {
	return s.OutputDataSourceId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputDataSourceType() *string {
	return s.OutputDataSourceType
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputDbType() *string {
	return s.OutputDbType
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputEnv() *string {
	return s.OutputEnv
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputProjectId() *int64 {
	return s.OutputProjectId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputTableDeleted() *bool {
	return s.OutputTableDeleted
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputTableId() *string {
	return s.OutputTableId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputTableName() *string {
	return s.OutputTableName
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetOutputTableType() *string {
	return s.OutputTableType
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetTaskEnv() *string {
	return s.TaskEnv
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTableLineageByTaskIdResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputBizUnitId(v int64) *GetTableLineageByTaskIdResponseBodyData {
	s.InputBizUnitId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputDataSourceId(v int64) *GetTableLineageByTaskIdResponseBodyData {
	s.InputDataSourceId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputDataSourceType(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.InputDataSourceType = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputDbType(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.InputDbType = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputEnv(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.InputEnv = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputProjectId(v int64) *GetTableLineageByTaskIdResponseBodyData {
	s.InputProjectId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputTableDeleted(v bool) *GetTableLineageByTaskIdResponseBodyData {
	s.InputTableDeleted = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputTableId(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.InputTableId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputTableName(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.InputTableName = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetInputTableType(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.InputTableType = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputBizUnitId(v int64) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputBizUnitId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputDataSourceId(v int64) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputDataSourceId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputDataSourceType(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputDataSourceType = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputDbType(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputDbType = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputEnv(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputEnv = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputProjectId(v int64) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputProjectId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputTableDeleted(v bool) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputTableDeleted = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputTableId(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputTableId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputTableName(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputTableName = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetOutputTableType(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.OutputTableType = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetTaskEnv(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.TaskEnv = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetTaskId(v string) *GetTableLineageByTaskIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) SetTenantId(v int64) *GetTableLineageByTaskIdResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetTableLineageByTaskIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
