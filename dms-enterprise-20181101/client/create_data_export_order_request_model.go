// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataExportOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentKey(v string) *CreateDataExportOrderRequest
	GetAttachmentKey() *string
	SetComment(v string) *CreateDataExportOrderRequest
	GetComment() *string
	SetParentId(v int64) *CreateDataExportOrderRequest
	GetParentId() *int64
	SetPluginParam(v *CreateDataExportOrderRequestPluginParam) *CreateDataExportOrderRequest
	GetPluginParam() *CreateDataExportOrderRequestPluginParam
	SetRealLoginUserUid(v string) *CreateDataExportOrderRequest
	GetRealLoginUserUid() *string
	SetRelatedUserList(v []*int64) *CreateDataExportOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateDataExportOrderRequest
	GetTid() *int64
}

type CreateDataExportOrderRequest struct {
	// The key of the attachment that provides more instructions for the ticket. You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to obtain the attachment key.
	//
	// example:
	//
	// order_attachment.txt
	AttachmentKey *string `json:"AttachmentKey,omitempty" xml:"AttachmentKey,omitempty"`
	// The purpose or objective of the ticket. This parameter helps reduce unnecessary communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// business_test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the parent ticket.
	//
	// example:
	//
	// 877****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	PluginParam *CreateDataExportOrderRequestPluginParam `json:"PluginParam,omitempty" xml:"PluginParam,omitempty" type:"Struct"`
	// The UID of the Alibaba Cloud account that actually calls the API.
	//
	// example:
	//
	// 21400447956867****
	RealLoginUserUid *string `json:"RealLoginUserUid,omitempty" xml:"RealLoginUserUid,omitempty"`
	// The stakeholders involved in this operation.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the DMS console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataExportOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataExportOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataExportOrderRequest) GetAttachmentKey() *string {
	return s.AttachmentKey
}

func (s *CreateDataExportOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataExportOrderRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDataExportOrderRequest) GetPluginParam() *CreateDataExportOrderRequestPluginParam {
	return s.PluginParam
}

func (s *CreateDataExportOrderRequest) GetRealLoginUserUid() *string {
	return s.RealLoginUserUid
}

func (s *CreateDataExportOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateDataExportOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataExportOrderRequest) SetAttachmentKey(v string) *CreateDataExportOrderRequest {
	s.AttachmentKey = &v
	return s
}

func (s *CreateDataExportOrderRequest) SetComment(v string) *CreateDataExportOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataExportOrderRequest) SetParentId(v int64) *CreateDataExportOrderRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataExportOrderRequest) SetPluginParam(v *CreateDataExportOrderRequestPluginParam) *CreateDataExportOrderRequest {
	s.PluginParam = v
	return s
}

func (s *CreateDataExportOrderRequest) SetRealLoginUserUid(v string) *CreateDataExportOrderRequest {
	s.RealLoginUserUid = &v
	return s
}

func (s *CreateDataExportOrderRequest) SetRelatedUserList(v []*int64) *CreateDataExportOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataExportOrderRequest) SetTid(v int64) *CreateDataExportOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataExportOrderRequest) Validate() error {
	if s.PluginParam != nil {
		if err := s.PluginParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataExportOrderRequestPluginParam struct {
	// The estimated number of data rows to be affected.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The reason for the export ticket.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The database ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The SQL statements that can be executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM DMS_test
	//
	//  LIMIT 20;
	ExeSQL *string `json:"ExeSQL,omitempty" xml:"ExeSQL,omitempty"`
	// Specifies whether to skip verification. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IgnoreAffectRows *bool `json:"IgnoreAffectRows,omitempty" xml:"IgnoreAffectRows,omitempty"`
	// The reason for skipping verification. This parameter is required if you set IgnoreAffectRows to true.
	//
	// example:
	//
	// Test only, does not affect the business, and does not require verification.
	IgnoreAffectRowsReason *string `json:"IgnoreAffectRowsReason,omitempty" xml:"IgnoreAffectRowsReason,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 132****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the database is a logical database. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > If you set this parameter to **true**, the database that you specify must be a logical database.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The information about the watermarks.
	Watermark *CreateDataExportOrderRequestPluginParamWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
}

func (s CreateDataExportOrderRequestPluginParam) String() string {
	return dara.Prettify(s)
}

func (s CreateDataExportOrderRequestPluginParam) GoString() string {
	return s.String()
}

func (s *CreateDataExportOrderRequestPluginParam) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *CreateDataExportOrderRequestPluginParam) GetClassify() *string {
	return s.Classify
}

func (s *CreateDataExportOrderRequestPluginParam) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateDataExportOrderRequestPluginParam) GetExeSQL() *string {
	return s.ExeSQL
}

func (s *CreateDataExportOrderRequestPluginParam) GetIgnoreAffectRows() *bool {
	return s.IgnoreAffectRows
}

func (s *CreateDataExportOrderRequestPluginParam) GetIgnoreAffectRowsReason() *string {
	return s.IgnoreAffectRowsReason
}

func (s *CreateDataExportOrderRequestPluginParam) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *CreateDataExportOrderRequestPluginParam) GetLogic() *bool {
	return s.Logic
}

func (s *CreateDataExportOrderRequestPluginParam) GetWatermark() *CreateDataExportOrderRequestPluginParamWatermark {
	return s.Watermark
}

func (s *CreateDataExportOrderRequestPluginParam) SetAffectRows(v int64) *CreateDataExportOrderRequestPluginParam {
	s.AffectRows = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetClassify(v string) *CreateDataExportOrderRequestPluginParam {
	s.Classify = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetDbId(v int64) *CreateDataExportOrderRequestPluginParam {
	s.DbId = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetExeSQL(v string) *CreateDataExportOrderRequestPluginParam {
	s.ExeSQL = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetIgnoreAffectRows(v bool) *CreateDataExportOrderRequestPluginParam {
	s.IgnoreAffectRows = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetIgnoreAffectRowsReason(v string) *CreateDataExportOrderRequestPluginParam {
	s.IgnoreAffectRowsReason = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetInstanceId(v int64) *CreateDataExportOrderRequestPluginParam {
	s.InstanceId = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetLogic(v bool) *CreateDataExportOrderRequestPluginParam {
	s.Logic = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) SetWatermark(v *CreateDataExportOrderRequestPluginParamWatermark) *CreateDataExportOrderRequestPluginParam {
	s.Watermark = v
	return s
}

func (s *CreateDataExportOrderRequestPluginParam) Validate() error {
	if s.Watermark != nil {
		if err := s.Watermark.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataExportOrderRequestPluginParamWatermark struct {
	// The field into which the watermark is to be embedded.
	//
	// example:
	//
	// user_number
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The information to be embedded as a watermark into data.
	//
	// example:
	//
	// test
	DataWatermark *string `json:"DataWatermark,omitempty" xml:"DataWatermark,omitempty"`
	// The information to be embedded as a watermark into files.
	//
	// example:
	//
	// test
	FileWatermark *string `json:"FileWatermark,omitempty" xml:"FileWatermark,omitempty"`
	// One or more primary keys or unique keys.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The methods in which the watermark is embedded.
	WatermarkTypes []*string `json:"WatermarkTypes,omitempty" xml:"WatermarkTypes,omitempty" type:"Repeated"`
}

func (s CreateDataExportOrderRequestPluginParamWatermark) String() string {
	return dara.Prettify(s)
}

func (s CreateDataExportOrderRequestPluginParamWatermark) GoString() string {
	return s.String()
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) GetColumnName() *string {
	return s.ColumnName
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) GetDataWatermark() *string {
	return s.DataWatermark
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) GetFileWatermark() *string {
	return s.FileWatermark
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) GetKeys() []*string {
	return s.Keys
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) GetWatermarkTypes() []*string {
	return s.WatermarkTypes
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) SetColumnName(v string) *CreateDataExportOrderRequestPluginParamWatermark {
	s.ColumnName = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) SetDataWatermark(v string) *CreateDataExportOrderRequestPluginParamWatermark {
	s.DataWatermark = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) SetFileWatermark(v string) *CreateDataExportOrderRequestPluginParamWatermark {
	s.FileWatermark = &v
	return s
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) SetKeys(v []*string) *CreateDataExportOrderRequestPluginParamWatermark {
	s.Keys = v
	return s
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) SetWatermarkTypes(v []*string) *CreateDataExportOrderRequestPluginParamWatermark {
	s.WatermarkTypes = v
	return s
}

func (s *CreateDataExportOrderRequestPluginParamWatermark) Validate() error {
	return dara.Validate(s)
}
