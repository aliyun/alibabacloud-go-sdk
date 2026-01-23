// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableLineagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableLineagesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetTableLineagesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTableLineagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableLineagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableLineagesResponseBody
	GetSuccess() *bool
	SetTableLineageList(v []*GetTableLineagesResponseBodyTableLineageList) *GetTableLineagesResponseBody
	GetTableLineageList() []*GetTableLineagesResponseBodyTableLineageList
}

type GetTableLineagesResponseBody struct {
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
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TableLineageList []*GetTableLineagesResponseBodyTableLineageList `json:"TableLineageList,omitempty" xml:"TableLineageList,omitempty" type:"Repeated"`
}

func (s GetTableLineagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineagesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableLineagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableLineagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTableLineagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableLineagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableLineagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableLineagesResponseBody) GetTableLineageList() []*GetTableLineagesResponseBodyTableLineageList {
	return s.TableLineageList
}

func (s *GetTableLineagesResponseBody) SetCode(v string) *GetTableLineagesResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableLineagesResponseBody) SetHttpStatusCode(v int32) *GetTableLineagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTableLineagesResponseBody) SetMessage(v string) *GetTableLineagesResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableLineagesResponseBody) SetRequestId(v string) *GetTableLineagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableLineagesResponseBody) SetSuccess(v bool) *GetTableLineagesResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableLineagesResponseBody) SetTableLineageList(v []*GetTableLineagesResponseBodyTableLineageList) *GetTableLineagesResponseBody {
	s.TableLineageList = v
	return s
}

func (s *GetTableLineagesResponseBody) Validate() error {
	if s.TableLineageList != nil {
		for _, item := range s.TableLineageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableLineagesResponseBodyTableLineageList struct {
	// example:
	//
	// 1
	InputBizUnitId *int64 `json:"InputBizUnitId,omitempty" xml:"InputBizUnitId,omitempty"`
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

func (s GetTableLineagesResponseBodyTableLineageList) String() string {
	return dara.Prettify(s)
}

func (s GetTableLineagesResponseBodyTableLineageList) GoString() string {
	return s.String()
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputBizUnitId() *int64 {
	return s.InputBizUnitId
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputDataSourceId() *int64 {
	return s.InputDataSourceId
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputDataSourceType() *string {
	return s.InputDataSourceType
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputProjectId() *int64 {
	return s.InputProjectId
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputTableDeleted() *bool {
	return s.InputTableDeleted
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputTableEnv() *string {
	return s.InputTableEnv
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputTableGuid() *string {
	return s.InputTableGuid
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetInputTableName() *string {
	return s.InputTableName
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetNodeEnv() *string {
	return s.NodeEnv
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetNodeId() *string {
	return s.NodeId
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputBizUnitId() *int64 {
	return s.OutputBizUnitId
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputDataSourceId() *int64 {
	return s.OutputDataSourceId
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputDataSourceType() *string {
	return s.OutputDataSourceType
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputProjectId() *int64 {
	return s.OutputProjectId
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputTableDeleted() *bool {
	return s.OutputTableDeleted
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputTableEnv() *string {
	return s.OutputTableEnv
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputTableGuid() *string {
	return s.OutputTableGuid
}

func (s *GetTableLineagesResponseBodyTableLineageList) GetOutputTableName() *string {
	return s.OutputTableName
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputBizUnitId(v int64) *GetTableLineagesResponseBodyTableLineageList {
	s.InputBizUnitId = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputDataSourceId(v int64) *GetTableLineagesResponseBodyTableLineageList {
	s.InputDataSourceId = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputDataSourceType(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.InputDataSourceType = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputProjectId(v int64) *GetTableLineagesResponseBodyTableLineageList {
	s.InputProjectId = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputTableDeleted(v bool) *GetTableLineagesResponseBodyTableLineageList {
	s.InputTableDeleted = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputTableEnv(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.InputTableEnv = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputTableGuid(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.InputTableGuid = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetInputTableName(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.InputTableName = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetNodeEnv(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.NodeEnv = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetNodeId(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.NodeId = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputBizUnitId(v int64) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputBizUnitId = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputDataSourceId(v int64) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputDataSourceId = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputDataSourceType(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputDataSourceType = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputProjectId(v int64) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputProjectId = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputTableDeleted(v bool) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputTableDeleted = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputTableEnv(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputTableEnv = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputTableGuid(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputTableGuid = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) SetOutputTableName(v string) *GetTableLineagesResponseBodyTableLineageList {
	s.OutputTableName = &v
	return s
}

func (s *GetTableLineagesResponseBodyTableLineageList) Validate() error {
	return dara.Validate(s)
}
