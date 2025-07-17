// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCronClearOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataCronClearOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataCronClearOrderRequest
	GetComment() *string
	SetParam(v *CreateDataCronClearOrderRequestParam) *CreateDataCronClearOrderRequest
	GetParam() *CreateDataCronClearOrderRequestParam
	SetRelatedUserList(v []*int64) *CreateDataCronClearOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateDataCronClearOrderRequest
	GetTid() *int64
}

type CreateDataCronClearOrderRequest struct {
	// The key of the attachment for the ticket. The attachment provides more instructions for this operation.
	//
	// You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to query the key of the attachment.
	//
	// example:
	//
	// order_attachement.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the data change. This reduces unnecessary communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	Param *CreateDataCronClearOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The stakeholders of this operation. All stakeholders can view the ticket details and assist in the approval process. Irrelevant users other than Data Management (DMS) administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// >  The ID of the tenant is displayed when you move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the Manage DMS tenants topic.
	//
	// example:
	//
	// 123454324
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataCronClearOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCronClearOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataCronClearOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataCronClearOrderRequest) GetParam() *CreateDataCronClearOrderRequestParam {
	return s.Param
}

func (s *CreateDataCronClearOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateDataCronClearOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataCronClearOrderRequest) SetAttachmentKey(v string) *CreateDataCronClearOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetComment(v string) *CreateDataCronClearOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetParam(v *CreateDataCronClearOrderRequestParam) *CreateDataCronClearOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetRelatedUserList(v []*int64) *CreateDataCronClearOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataCronClearOrderRequest) SetTid(v int64) *CreateDataCronClearOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataCronClearOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataCronClearOrderRequestParam struct {
	// The reason for the data change.
	//
	// example:
	//
	// test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The tables for which you want to clear historical data.
	//
	// This parameter is required.
	CronClearItemList []*CreateDataCronClearOrderRequestParamCronClearItemList `json:"CronClearItemList,omitempty" xml:"CronClearItemList,omitempty" type:"Repeated"`
	// The crontab expression that you can use to run the task at a specified time. For more information, see [Crontab expression](https://help.aliyun.com/document_detail/206581.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	CronFormat *string `json:"CronFormat,omitempty" xml:"CronFormat,omitempty"`
	// The databases for which you want to clear historical data.
	//
	// This parameter is required.
	DbItemList []*CreateDataCronClearOrderRequestParamDbItemList `json:"DbItemList,omitempty" xml:"DbItemList,omitempty" type:"Repeated"`
	// The amount of time taken to run the task. Unit: hours.
	//
	// >  If the **specifyDuration*	- parameter is set to **true**, this parameter is required.
	//
	// example:
	//
	// 4
	DurationHour *int64 `json:"DurationHour,omitempty" xml:"DurationHour,omitempty"`
	// Specifies whether to specify an end time for the task. Valid values:
	//
	// 	- **true**: specifies an end time for the task. The task is automatically suspended after this end time.
	//
	// 	- **false**: does not specify an end time for the task. The task is stopped after the historical data is cleared.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SpecifyDuration *bool `json:"specifyDuration,omitempty" xml:"specifyDuration,omitempty"`
}

func (s CreateDataCronClearOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCronClearOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequestParam) GetClassify() *string {
	return s.Classify
}

func (s *CreateDataCronClearOrderRequestParam) GetCronClearItemList() []*CreateDataCronClearOrderRequestParamCronClearItemList {
	return s.CronClearItemList
}

func (s *CreateDataCronClearOrderRequestParam) GetCronFormat() *string {
	return s.CronFormat
}

func (s *CreateDataCronClearOrderRequestParam) GetDbItemList() []*CreateDataCronClearOrderRequestParamDbItemList {
	return s.DbItemList
}

func (s *CreateDataCronClearOrderRequestParam) GetDurationHour() *int64 {
	return s.DurationHour
}

func (s *CreateDataCronClearOrderRequestParam) GetSpecifyDuration() *bool {
	return s.SpecifyDuration
}

func (s *CreateDataCronClearOrderRequestParam) SetClassify(v string) *CreateDataCronClearOrderRequestParam {
	s.Classify = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetCronClearItemList(v []*CreateDataCronClearOrderRequestParamCronClearItemList) *CreateDataCronClearOrderRequestParam {
	s.CronClearItemList = v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetCronFormat(v string) *CreateDataCronClearOrderRequestParam {
	s.CronFormat = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetDbItemList(v []*CreateDataCronClearOrderRequestParamDbItemList) *CreateDataCronClearOrderRequestParam {
	s.DbItemList = v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetDurationHour(v int64) *CreateDataCronClearOrderRequestParam {
	s.DurationHour = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) SetSpecifyDuration(v bool) *CreateDataCronClearOrderRequestParam {
	s.SpecifyDuration = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParam) Validate() error {
	return dara.Validate(s)
}

type CreateDataCronClearOrderRequestParamCronClearItemList struct {
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// gmt_create
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The filter conditions.
	//
	// example:
	//
	// where 1 = 1
	FilterSQL *string `json:"FilterSQL,omitempty" xml:"FilterSQL,omitempty"`
	// The retention period of the historical data. Unit: days. For example, if you set the parameter to 7, DMS deletes the data that is retained for more than seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	RemainDays *int64 `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
	// The name of the table. You can call the [ListTables](https://help.aliyun.com/document_detail/141878.html) operation to query the name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// t1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of time granularity. If the ColumnName parameter specifies a field of a time type, this parameter is required. Valid values:
	//
	// 	- **MILLISECONDS**: milliseconds
	//
	// 	- **SECONDS**: seconds
	//
	// example:
	//
	// MILLISECONDS
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
}

func (s CreateDataCronClearOrderRequestParamCronClearItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCronClearOrderRequestParamCronClearItemList) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) GetColumnName() *string {
	return s.ColumnName
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) GetFilterSQL() *string {
	return s.FilterSQL
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) GetRemainDays() *int64 {
	return s.RemainDays
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) GetTableName() *string {
	return s.TableName
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetColumnName(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.ColumnName = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetFilterSQL(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.FilterSQL = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetRemainDays(v int64) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.RemainDays = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetTableName(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.TableName = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) SetTimeUnit(v string) *CreateDataCronClearOrderRequestParamCronClearItemList {
	s.TimeUnit = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamCronClearItemList) Validate() error {
	return dara.Validate(s)
}

type CreateDataCronClearOrderRequestParamDbItemList struct {
	// The ID of the database. You can call the [SearchDatabases](https://help.aliyun.com/document_detail/141876.html) operation to query the ID of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is not a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
}

func (s CreateDataCronClearOrderRequestParamDbItemList) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCronClearOrderRequestParamDbItemList) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderRequestParamDbItemList) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateDataCronClearOrderRequestParamDbItemList) GetLogic() *bool {
	return s.Logic
}

func (s *CreateDataCronClearOrderRequestParamDbItemList) SetDbId(v int64) *CreateDataCronClearOrderRequestParamDbItemList {
	s.DbId = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamDbItemList) SetLogic(v bool) *CreateDataCronClearOrderRequestParamDbItemList {
	s.Logic = &v
	return s
}

func (s *CreateDataCronClearOrderRequestParamDbItemList) Validate() error {
	return dara.Validate(s)
}
