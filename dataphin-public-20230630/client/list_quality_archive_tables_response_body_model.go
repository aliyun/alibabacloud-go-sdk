// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityArchiveTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQualityArchiveTablesResponseBody
	GetCode() *string
	SetData(v *ListQualityArchiveTablesResponseBodyData) *ListQualityArchiveTablesResponseBody
	GetData() *ListQualityArchiveTablesResponseBodyData
	SetHttpStatusCode(v int32) *ListQualityArchiveTablesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListQualityArchiveTablesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListQualityArchiveTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityArchiveTablesResponseBody
	GetSuccess() *bool
}

type ListQualityArchiveTablesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of querying the anomaly archived table list.
	Data *ListQualityArchiveTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
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
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityArchiveTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityArchiveTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityArchiveTablesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQualityArchiveTablesResponseBody) GetData() *ListQualityArchiveTablesResponseBodyData {
	return s.Data
}

func (s *ListQualityArchiveTablesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityArchiveTablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQualityArchiveTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityArchiveTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityArchiveTablesResponseBody) SetCode(v string) *ListQualityArchiveTablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBody) SetData(v *ListQualityArchiveTablesResponseBodyData) *ListQualityArchiveTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListQualityArchiveTablesResponseBody) SetHttpStatusCode(v int32) *ListQualityArchiveTablesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBody) SetMessage(v string) *ListQualityArchiveTablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBody) SetRequestId(v string) *ListQualityArchiveTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBody) SetSuccess(v bool) *ListQualityArchiveTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityArchiveTablesResponseBodyData struct {
	// The list of anomaly archived tables.
	ArchiveTableList []*ListQualityArchiveTablesResponseBodyDataArchiveTableList `json:"ArchiveTableList,omitempty" xml:"ArchiveTableList,omitempty" type:"Repeated"`
	// The number of custom anomaly archived tables.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityArchiveTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListQualityArchiveTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQualityArchiveTablesResponseBodyData) GetArchiveTableList() []*ListQualityArchiveTablesResponseBodyDataArchiveTableList {
	return s.ArchiveTableList
}

func (s *ListQualityArchiveTablesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityArchiveTablesResponseBodyData) SetArchiveTableList(v []*ListQualityArchiveTablesResponseBodyDataArchiveTableList) *ListQualityArchiveTablesResponseBodyData {
	s.ArchiveTableList = v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyData) SetTotalCount(v int64) *ListQualityArchiveTablesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyData) Validate() error {
	if s.ArchiveTableList != nil {
		for _, item := range s.ArchiveTableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityArchiveTablesResponseBodyDataArchiveTableList struct {
	// The ID of the archived table. This ID is used when you update, switch to active, or delete the archived table.
	//
	// example:
	//
	// 7673533
	ArchiveTableId *int64 `json:"ArchiveTableId,omitempty" xml:"ArchiveTableId,omitempty"`
	// The full table name in the format of project_name.table_name.
	//
	// example:
	//
	// Train.a01_reanme_exception_data
	ArchiveTableName *string `json:"ArchiveTableName,omitempty" xml:"ArchiveTableName,omitempty"`
	// The DDL statement for creating the archived table, which includes dataphin_quality_	- system fields and the dataphin_quality_validate_date partition field definition.
	//
	// example:
	//
	// create table mfg_fin_cdm.a_shixin_b_exception_data\\n         (  \\n            dataphin_quality_tenant_id varchar(64) comment \\"Tenant ID\\"\\n         , \\n            dataphin_quality_rule_id varchar(64) comment \\"Quality rule ID\\"\\n         , \\n            dataphin_quality_rule_name string comment \\"Quality rule name\\"\\n         , \\n            dataphin_quality_column_name varchar(1024) comment \\"Validation field name\\"\\n         , \\n            dataphin_quality_watch_task_id varchar(128) comment \\"Monitored object task ID\\"\\n         , \\n            dataphin_quality_rule_task_id varchar(64) comment \\"Rule task ID\\"\\n         , \\n            dataphin_quality_validate_time varchar(64) comment \\"Quality validation time\\"\\n         , \\n            dataphin_quality_archive_mode varchar(32) comment \\"Anomaly archiving mode, ONLY_ERROR_FIELD/FULL_RECORD\\"\\n         , \\n            dataphin_quality_error_data string comment \\"Anomaly data\\"\\n         , \\n            执行依据文号 string comment \\"\\"\\n         , \\n            立案时间 string comment \\"\\"\\n         , \\n            案号 string comment \\"\\"\\n         , \\n            执行法院 string comment \\"\\"\\n         , \\n            性别 string comment \\"\\"\\n         , \\n            省份 string comment \\"\\"\\n         , \\n            被执行人的履行情况 string comment \\"\\"\\n         , \\n            发布时间 string comment \\"\\"\\n         , \\n            姓名 string comment \\"\\"\\n         , \\n            身份证号 string comment \\"\\"\\n         , \\n            失信被执行人行为具体情形 string comment \\"\\"\\n         ) \\n        partitioned by (dataphin_quality_validate_date string comment \\"Validation date (partition field)\\")
	Ddl *string `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
	// Indicates whether this is the active archived table. At least one active archived table must exist under the same monitored object.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The lifecycle in days. An empty value indicates no lifecycle is configured.
	//
	// example:
	//
	// 30
	Lifecycle *int32 `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The maximum number of records to archive per validation. A value of -1 indicates full archiving.
	//
	// example:
	//
	// 10000
	MaxArchiveCount *int64 `json:"MaxArchiveCount,omitempty" xml:"MaxArchiveCount,omitempty"`
}

func (s ListQualityArchiveTablesResponseBodyDataArchiveTableList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityArchiveTablesResponseBodyDataArchiveTableList) GoString() string {
	return s.String()
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) GetArchiveTableId() *int64 {
	return s.ArchiveTableId
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) GetArchiveTableName() *string {
	return s.ArchiveTableName
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) GetDdl() *string {
	return s.Ddl
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) GetLifecycle() *int32 {
	return s.Lifecycle
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) GetMaxArchiveCount() *int64 {
	return s.MaxArchiveCount
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) SetArchiveTableId(v int64) *ListQualityArchiveTablesResponseBodyDataArchiveTableList {
	s.ArchiveTableId = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) SetArchiveTableName(v string) *ListQualityArchiveTablesResponseBodyDataArchiveTableList {
	s.ArchiveTableName = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) SetDdl(v string) *ListQualityArchiveTablesResponseBodyDataArchiveTableList {
	s.Ddl = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) SetIsDefault(v bool) *ListQualityArchiveTablesResponseBodyDataArchiveTableList {
	s.IsDefault = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) SetLifecycle(v int32) *ListQualityArchiveTablesResponseBodyDataArchiveTableList {
	s.Lifecycle = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) SetMaxArchiveCount(v int64) *ListQualityArchiveTablesResponseBodyDataArchiveTableList {
	s.MaxArchiveCount = &v
	return s
}

func (s *ListQualityArchiveTablesResponseBodyDataArchiveTableList) Validate() error {
	return dara.Validate(s)
}
