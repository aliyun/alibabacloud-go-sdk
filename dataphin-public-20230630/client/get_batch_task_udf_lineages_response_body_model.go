// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskUdfLineagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchTaskUdfLineagesResponseBody
	GetCode() *string
	SetData(v *GetBatchTaskUdfLineagesResponseBodyData) *GetBatchTaskUdfLineagesResponseBody
	GetData() *GetBatchTaskUdfLineagesResponseBodyData
	SetHttpStatusCode(v int32) *GetBatchTaskUdfLineagesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBatchTaskUdfLineagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTaskUdfLineagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBatchTaskUdfLineagesResponseBody
	GetSuccess() *bool
}

type GetBatchTaskUdfLineagesResponseBody struct {
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBatchTaskUdfLineagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBatchTaskUdfLineagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchTaskUdfLineagesResponseBody) GetData() *GetBatchTaskUdfLineagesResponseBodyData {
	return s.Data
}

func (s *GetBatchTaskUdfLineagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBatchTaskUdfLineagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTaskUdfLineagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTaskUdfLineagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBatchTaskUdfLineagesResponseBody) SetCode(v string) *GetBatchTaskUdfLineagesResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBody) SetData(v *GetBatchTaskUdfLineagesResponseBodyData) *GetBatchTaskUdfLineagesResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBody) SetHttpStatusCode(v int32) *GetBatchTaskUdfLineagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBody) SetMessage(v string) *GetBatchTaskUdfLineagesResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBody) SetRequestId(v string) *GetBatchTaskUdfLineagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBody) SetSuccess(v bool) *GetBatchTaskUdfLineagesResponseBody {
	s.Success = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchTaskUdfLineagesResponseBodyData struct {
	LineageGroupList []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList `json:"LineageGroupList,omitempty" xml:"LineageGroupList,omitempty" type:"Repeated"`
}

func (s GetBatchTaskUdfLineagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponseBodyData) GetLineageGroupList() []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList {
	return s.LineageGroupList
}

func (s *GetBatchTaskUdfLineagesResponseBodyData) SetLineageGroupList(v []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) *GetBatchTaskUdfLineagesResponseBodyData {
	s.LineageGroupList = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyData) Validate() error {
	if s.LineageGroupList != nil {
		for _, item := range s.LineageGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList struct {
	InputLineageList  []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList  `json:"InputLineageList,omitempty" xml:"InputLineageList,omitempty" type:"Repeated"`
	OutputLineageList []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList `json:"OutputLineageList,omitempty" xml:"OutputLineageList,omitempty" type:"Repeated"`
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) GetInputLineageList() []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	return s.InputLineageList
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) GetOutputLineageList() []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	return s.OutputLineageList
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) SetInputLineageList(v []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList {
	s.InputLineageList = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) SetOutputLineageList(v []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList {
	s.OutputLineageList = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupList) Validate() error {
	if s.InputLineageList != nil {
		for _, item := range s.InputLineageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OutputLineageList != nil {
		for _, item := range s.OutputLineageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList struct {
	// example:
	//
	// 103111231
	BizUnitId *string `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// xx测试
	BizUnitName *string                                                                              `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	ColumnList  []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// example:
	//
	// test xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// t_input
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// dev
	Env       *string `json:"Env,omitempty" xml:"Env,omitempty"`
	FullTable *bool   `json:"FullTable,omitempty" xml:"FullTable,omitempty"`
	// example:
	//
	// Guid_101121
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// t_input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 20112101
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// 131211211
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// prj_test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// dim
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetBizUnitId() *string {
	return s.BizUnitId
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetColumnList() []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList {
	return s.ColumnList
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetDescription() *string {
	return s.Description
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetEnv() *string {
	return s.Env
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetFullTable() *bool {
	return s.FullTable
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetGuid() *string {
	return s.Guid
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetName() *string {
	return s.Name
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) GetSubType() *string {
	return s.SubType
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetBizUnitId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.BizUnitId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetBizUnitName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.BizUnitName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetColumnList(v []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.ColumnList = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetDescription(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.Description = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetDisplayName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.DisplayName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetEnv(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.Env = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetFullTable(v bool) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.FullTable = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetGuid(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.Guid = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.Name = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetOwnerName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.OwnerName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetOwnerUserId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.OwnerUserId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetProjectId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetProjectName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.ProjectName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) SetSubType(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList {
	s.SubType = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageList) Validate() error {
	if s.ColumnList != nil {
		for _, item := range s.ColumnList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList struct {
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// c011
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// c011
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PartitionKey *bool   `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	PrimaryKey   *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) GetDataType() *string {
	return s.DataType
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) GetDescription() *string {
	return s.Description
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) GetId() *string {
	return s.Id
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) GetName() *string {
	return s.Name
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) GetPartitionKey() *bool {
	return s.PartitionKey
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) SetDataType(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList {
	s.DataType = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) SetDescription(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList {
	s.Description = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) SetId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList {
	s.Id = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) SetName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList {
	s.Name = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) SetPartitionKey(v bool) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList {
	s.PartitionKey = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) SetPrimaryKey(v bool) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList {
	s.PrimaryKey = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListInputLineageListColumnList) Validate() error {
	return dara.Validate(s)
}

type GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList struct {
	// example:
	//
	// 103111231
	BizUnitId *string `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// xx测试
	BizUnitName *string                                                                               `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	ColumnList  []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// example:
	//
	// test xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// t_input
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// dev
	Env       *string `json:"Env,omitempty" xml:"Env,omitempty"`
	FullTable *bool   `json:"FullTable,omitempty" xml:"FullTable,omitempty"`
	// example:
	//
	// Guid_101121
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// example:
	//
	// t_input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 20112101
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// 131211211
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// prj_test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// dim
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetBizUnitId() *string {
	return s.BizUnitId
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetColumnList() []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList {
	return s.ColumnList
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetDescription() *string {
	return s.Description
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetEnv() *string {
	return s.Env
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetFullTable() *bool {
	return s.FullTable
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetGuid() *string {
	return s.Guid
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetName() *string {
	return s.Name
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) GetSubType() *string {
	return s.SubType
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetBizUnitId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.BizUnitId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetBizUnitName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.BizUnitName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetColumnList(v []*GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.ColumnList = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetDescription(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.Description = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetDisplayName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.DisplayName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetEnv(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.Env = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetFullTable(v bool) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.FullTable = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetGuid(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.Guid = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.Name = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetOwnerName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.OwnerName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetOwnerUserId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.OwnerUserId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetProjectId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetProjectName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.ProjectName = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) SetSubType(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList {
	s.SubType = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageList) Validate() error {
	if s.ColumnList != nil {
		for _, item := range s.ColumnList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList struct {
	// example:
	//
	// varchar
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// c011
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// c011
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PartitionKey *bool   `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	PrimaryKey   *bool   `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) GetDataType() *string {
	return s.DataType
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) GetDescription() *string {
	return s.Description
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) GetId() *string {
	return s.Id
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) GetName() *string {
	return s.Name
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) GetPartitionKey() *bool {
	return s.PartitionKey
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) SetDataType(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList {
	s.DataType = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) SetDescription(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList {
	s.Description = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) SetId(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList {
	s.Id = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) SetName(v string) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList {
	s.Name = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) SetPartitionKey(v bool) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList {
	s.PartitionKey = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) SetPrimaryKey(v bool) *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList {
	s.PrimaryKey = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponseBodyDataLineageGroupListOutputLineageListColumnList) Validate() error {
	return dara.Validate(s)
}
