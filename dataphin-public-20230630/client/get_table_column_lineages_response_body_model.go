// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableColumnLineagesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetTableColumnLineagesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTableColumnLineagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableColumnLineagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableColumnLineagesResponseBody
	GetSuccess() *bool
	SetTableColumnLineageList(v []*GetTableColumnLineagesResponseBodyTableColumnLineageList) *GetTableColumnLineagesResponseBody
	GetTableColumnLineageList() []*GetTableColumnLineagesResponseBodyTableColumnLineageList
}

type GetTableColumnLineagesResponseBody struct {
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
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId              *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TableColumnLineageList []*GetTableColumnLineagesResponseBodyTableColumnLineageList `json:"TableColumnLineageList,omitempty" xml:"TableColumnLineageList,omitempty" type:"Repeated"`
}

func (s GetTableColumnLineagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineagesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableColumnLineagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTableColumnLineagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableColumnLineagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableColumnLineagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableColumnLineagesResponseBody) GetTableColumnLineageList() []*GetTableColumnLineagesResponseBodyTableColumnLineageList {
	return s.TableColumnLineageList
}

func (s *GetTableColumnLineagesResponseBody) SetCode(v string) *GetTableColumnLineagesResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableColumnLineagesResponseBody) SetHttpStatusCode(v int32) *GetTableColumnLineagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTableColumnLineagesResponseBody) SetMessage(v string) *GetTableColumnLineagesResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableColumnLineagesResponseBody) SetRequestId(v string) *GetTableColumnLineagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBody) SetSuccess(v bool) *GetTableColumnLineagesResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableColumnLineagesResponseBody) SetTableColumnLineageList(v []*GetTableColumnLineagesResponseBodyTableColumnLineageList) *GetTableColumnLineagesResponseBody {
	s.TableColumnLineageList = v
	return s
}

func (s *GetTableColumnLineagesResponseBody) Validate() error {
	if s.TableColumnLineageList != nil {
		for _, item := range s.TableColumnLineageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableColumnLineagesResponseBodyTableColumnLineageList struct {
	// example:
	//
	// 1
	InputBizUnitId *int64 `json:"InputBizUnitId,omitempty" xml:"InputBizUnitId,omitempty"`
	// example:
	//
	// 1121.col1
	InputColumnGuid *string `json:"InputColumnGuid,omitempty" xml:"InputColumnGuid,omitempty"`
	// example:
	//
	// col1
	InputColumnName *string `json:"InputColumnName,omitempty" xml:"InputColumnName,omitempty"`
	// example:
	//
	// 1
	InputDataSourceId *int64 `json:"InputDataSourceId,omitempty" xml:"InputDataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	InputDataSourceType *string `json:"InputDataSourceType,omitempty" xml:"InputDataSourceType,omitempty"`
	// example:
	//
	// 1233
	InputProjectId    *int64 `json:"InputProjectId,omitempty" xml:"InputProjectId,omitempty"`
	InputTableDeleted *bool  `json:"InputTableDeleted,omitempty" xml:"InputTableDeleted,omitempty"`
	// example:
	//
	// dev
	InputTableEnv *string `json:"InputTableEnv,omitempty" xml:"InputTableEnv,omitempty"`
	// example:
	//
	// 123211
	InputTableGuid *string `json:"InputTableGuid,omitempty" xml:"InputTableGuid,omitempty"`
	// example:
	//
	// t_input
	InputTableName *string `json:"InputTableName,omitempty" xml:"InputTableName,omitempty"`
	// example:
	//
	// dev
	NodeEnv *string `json:"NodeEnv,omitempty" xml:"NodeEnv,omitempty"`
	// example:
	//
	// 110021
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1
	OutputBizUnitId *int64 `json:"OutputBizUnitId,omitempty" xml:"OutputBizUnitId,omitempty"`
	// example:
	//
	// 2231.col2
	OutputColumnGuid *string `json:"OutputColumnGuid,omitempty" xml:"OutputColumnGuid,omitempty"`
	// example:
	//
	// col2
	OutputColumnName *string `json:"OutputColumnName,omitempty" xml:"OutputColumnName,omitempty"`
	// example:
	//
	// 1
	OutputDataSourceId *int64 `json:"OutputDataSourceId,omitempty" xml:"OutputDataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	OutputDataSourceType *string `json:"OutputDataSourceType,omitempty" xml:"OutputDataSourceType,omitempty"`
	// example:
	//
	// 1233
	OutputProjectId    *int64 `json:"OutputProjectId,omitempty" xml:"OutputProjectId,omitempty"`
	OutputTableDeleted *bool  `json:"OutputTableDeleted,omitempty" xml:"OutputTableDeleted,omitempty"`
	// example:
	//
	// dev
	OutputTableEnv *string `json:"OutputTableEnv,omitempty" xml:"OutputTableEnv,omitempty"`
	// example:
	//
	// 2231
	OutputTableGuid *string `json:"OutputTableGuid,omitempty" xml:"OutputTableGuid,omitempty"`
	// example:
	//
	// t_output
	OutputTableName *string `json:"OutputTableName,omitempty" xml:"OutputTableName,omitempty"`
}

func (s GetTableColumnLineagesResponseBodyTableColumnLineageList) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineagesResponseBodyTableColumnLineageList) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputBizUnitId() *int64 {
	return s.InputBizUnitId
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputColumnGuid() *string {
	return s.InputColumnGuid
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputColumnName() *string {
	return s.InputColumnName
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputDataSourceId() *int64 {
	return s.InputDataSourceId
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputDataSourceType() *string {
	return s.InputDataSourceType
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputProjectId() *int64 {
	return s.InputProjectId
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputTableDeleted() *bool {
	return s.InputTableDeleted
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputTableEnv() *string {
	return s.InputTableEnv
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputTableGuid() *string {
	return s.InputTableGuid
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetInputTableName() *string {
	return s.InputTableName
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetNodeEnv() *string {
	return s.NodeEnv
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputBizUnitId() *int64 {
	return s.OutputBizUnitId
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputColumnGuid() *string {
	return s.OutputColumnGuid
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputColumnName() *string {
	return s.OutputColumnName
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputDataSourceId() *int64 {
	return s.OutputDataSourceId
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputDataSourceType() *string {
	return s.OutputDataSourceType
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputProjectId() *int64 {
	return s.OutputProjectId
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputTableDeleted() *bool {
	return s.OutputTableDeleted
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputTableEnv() *string {
	return s.OutputTableEnv
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputTableGuid() *string {
	return s.OutputTableGuid
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) GetOutputTableName() *string {
	return s.OutputTableName
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputBizUnitId(v int64) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputBizUnitId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputColumnGuid(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputColumnGuid = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputColumnName(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputColumnName = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputDataSourceId(v int64) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputDataSourceId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputDataSourceType(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputDataSourceType = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputProjectId(v int64) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputProjectId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputTableDeleted(v bool) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputTableDeleted = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputTableEnv(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputTableEnv = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputTableGuid(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputTableGuid = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetInputTableName(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.InputTableName = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetNodeEnv(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.NodeEnv = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetNodeId(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.NodeId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputBizUnitId(v int64) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputBizUnitId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputColumnGuid(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputColumnGuid = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputColumnName(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputColumnName = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputDataSourceId(v int64) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputDataSourceId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputDataSourceType(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputDataSourceType = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputProjectId(v int64) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputProjectId = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputTableDeleted(v bool) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputTableDeleted = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputTableEnv(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputTableEnv = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputTableGuid(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputTableGuid = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) SetOutputTableName(v string) *GetTableColumnLineagesResponseBodyTableColumnLineageList {
	s.OutputTableName = &v
	return s
}

func (s *GetTableColumnLineagesResponseBodyTableColumnLineageList) Validate() error {
	return dara.Validate(s)
}
