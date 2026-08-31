// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAssetsGovernObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataAssetsGovernObjectResponseBody
	GetCode() *string
	SetGovernObjectInfo(v *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) *GetDataAssetsGovernObjectResponseBody
	GetGovernObjectInfo() *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo
	SetHttpStatusCode(v int32) *GetDataAssetsGovernObjectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataAssetsGovernObjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataAssetsGovernObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataAssetsGovernObjectResponseBody
	GetSuccess() *bool
}

type GetDataAssetsGovernObjectResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The governance object details.
	GovernObjectInfo *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo `json:"GovernObjectInfo,omitempty" xml:"GovernObjectInfo,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend response exception details.
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
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataAssetsGovernObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataAssetsGovernObjectResponseBody) GetGovernObjectInfo() *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	return s.GovernObjectInfo
}

func (s *GetDataAssetsGovernObjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataAssetsGovernObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataAssetsGovernObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataAssetsGovernObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataAssetsGovernObjectResponseBody) SetCode(v string) *GetDataAssetsGovernObjectResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBody) SetGovernObjectInfo(v *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) *GetDataAssetsGovernObjectResponseBody {
	s.GovernObjectInfo = v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBody) SetHttpStatusCode(v int32) *GetDataAssetsGovernObjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBody) SetMessage(v string) *GetDataAssetsGovernObjectResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBody) SetRequestId(v string) *GetDataAssetsGovernObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBody) SetSuccess(v bool) *GetDataAssetsGovernObjectResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBody) Validate() error {
	if s.GovernObjectInfo != nil {
		if err := s.GovernObjectInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataAssetsGovernObjectResponseBodyGovernObjectInfo struct {
	// The time when the governance object was reported.
	//
	// example:
	//
	// 2026-08-31 10:06:01
	CommitTime *string `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The governance object ID.
	//
	// example:
	//
	// 96928483120
	GovernItemId *int64 `json:"GovernItemId,omitempty" xml:"GovernItemId,omitempty"`
	// The governance object ID.
	//
	// example:
	//
	// 54295947412
	GovernObjectId *int64 `json:"GovernObjectId,omitempty" xml:"GovernObjectId,omitempty"`
	// Indicates whether rectification is in progress.
	IsRectify *bool `json:"IsRectify,omitempty" xml:"IsRectify,omitempty"`
	// The list of owners.
	Owners []*GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
	// The governance issue object.
	Problem *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem `json:"Problem,omitempty" xml:"Problem,omitempty" type:"Struct"`
	// The properties.
	//
	// example:
	//
	// "properties": {
	//
	//                 "gmt_create": "2026-08-24 06:00:19.649",
	//
	//                 "index_compute_type": "",
	//
	//                 "table_env": "PROD",
	//
	//                 "table_datasource_id": "1",
	//
	//                 "index_catalog": "",
	//
	//                 "qd_feature_owner": "",
	//
	//                 "rule_task_start_time": "2026-08-24 06:00:01",
	//
	//                 "table_id": "odps.300023201.fashion_ads.api2mysql_demo",
	//
	//                 "rule_strength": "WEAK",
	//
	//                 "table_biz_unit_name": "LD_Fashion",
	//
	//                 "table_name": "fashion_ads.api2mysql_demo",
	//
	//                 "index_type": "",
	//
	//                 "table_datasource_from": "META_DATA",
	//
	//                 "datasource_type": "MAX_COMPUTE",
	//
	//                 "datasource_scope": "OFFLINE",
	//
	//                 "template_zh_tw_name": "欄位空值校正",
	//
	//                 "problem_submit_type": "SYSTEM",
	//
	//                 "template_type": "FIELD_NULL_VALUE_VALIDATE",
	//
	//                 "index_name_cn": "",
	//
	//                 "datasource_name": "Dataphin",
	//
	//                 "id": "909586",
	//
	//                 "index_id": "",
	//
	//                 "validate_partition": "ds=\\"20260824\\"",
	//
	//                 "index_owner_id": "",
	//
	//                 "datasource_owner": "300006218",
	//
	//                 "rule_name": "date_odps_test_2025-09-22 20:32:07",
	//
	//                 "watch_type": "TABLE",
	//
	//                 "validate_status": "NOT_PASS",
	//
	//                 "qd_feature_id": "",
	//
	//                 "table_desc": "api2mysql_demo",
	//
	//                 "is_ignore": "false",
	//
	//                 "rule_desc": "",
	//
	//                 "table_partitioned": "true",
	//
	//                 "template_owner": "300006218",
	//
	//                 "index_biz_unit_id": "",
	//
	//                 "table_biz_unit_id": "6865277495315392",
	//
	//                 "index_biz_unit_name": "",
	//
	//                 "watch_env": "PROD",
	//
	//                 "problem_contact_other": "",
	//
	//                 "status": "NEW",
	//
	//                 "datasource_from": "META_DATA",
	//
	//                 "table_project_name": "fashion_ads",
	//
	//                 "tenant_id": "300023201",
	//
	//                 "datasource_env": "PROD",
	//
	//                 "template_en_name": "Verify Field Null Values",
	//
	//                 "commit_time": "2026-08-24 06:00:18.73",
	//
	//                 "gmt_modified": "2026-08-24 06:00:19.649",
	//
	//                 "qd_feature_name": "",
	//
	//                 "table_catalog": "fashion_ads",
	//
	//                 "rule_status": "ENABLE",
	//
	//                 "problem_contact_mail": "",
	//
	//                 "rule_task_status": "SUCCESS",
	//
	//                 "rule_validate_object_type": "COLUMN",
	//
	//                 "watch_status": "ENABLE",
	//
	//                 "index_guid": "",
	//
	//                 "system_template": "true",
	//
	//                 "quality_owners": "300006218",
	//
	//                 "index_name": "",
	//
	//                 "problem_contact_phone": "",
	//
	//                 "watch_task_id": "8199222",
	//
	//                 "rule_task_id": "8199227",
	//
	//                 "index_desc": "",
	//
	//                 "table_type": "PHYSICAL_TABLE",
	//
	//                 "table_project_id": "6865331520706176",
	//
	//                 "trace_id": "time:4326023",
	//
	//                 "datasource_id": "1",
	//
	//                 "qd_feature_code": "",
	//
	//                 "rule_validate_object_name": "date_odps",
	//
	//                 "problem_desc": "date_odps_test_2025-09-22 20:32:07",
	//
	//                 "table_owner": "300006218",
	//
	//                 "quality_owner_groups": "",
	//
	//                 "govern_item_id": "100",
	//
	//                 "rule_catalogs": "COMPLETENESS",
	//
	//                 "table_datasource_type": "MAX_COMPUTE",
	//
	//                 "template_zh_cn_name": "字段空值校验",
	//
	//                 "rule_task_biz_date": "20260824 06:00:00",
	//
	//                 "problem_types": "[{\\"value\\":\\"COMPLETENESS\\"}]",
	//
	//                 "rule_id": "4322944",
	//
	//                 "problem_attachment_file_ids": "null",
	//
	//                 "watch_name": "",
	//
	//                 "template_name": "",
	//
	//                 "schedule_params": "",
	//
	//                 "is_rectify": "false",
	//
	//                 "rule_task_end_time": "",
	//
	//                 "watch_id": "3841908",
	//
	//                 "validate_result": "false",
	//
	//                 "qd_feature_guid": "",
	//
	//                 "govern_object_id": "910181",
	//
	//                 "template_id": "100"
	//
	//             }
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The ID of the rectification.
	//
	// example:
	//
	// 49169072991
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
	// The name of the rectification.
	//
	// example:
	//
	// Rectification process 1
	RectifyName *string `json:"RectifyName,omitempty" xml:"RectifyName,omitempty"`
	// The rectification status.
	//
	// example:
	//
	// NEW
	RectifyStatus *string `json:"RectifyStatus,omitempty" xml:"RectifyStatus,omitempty"`
	// The ID of the user who performs the rectification.
	//
	// example:
	//
	// 566777
	RectifyUser *string `json:"RectifyUser,omitempty" xml:"RectifyUser,omitempty"`
	// The display name of the rectification user.
	//
	// example:
	//
	// John
	RectifyUserName *string `json:"RectifyUserName,omitempty" xml:"RectifyUserName,omitempty"`
	// The related knowledge base.
	RelatedKnowledge []*GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge `json:"RelatedKnowledge,omitempty" xml:"RelatedKnowledge,omitempty" type:"Repeated"`
	// The status of the governance object.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The submission method.
	//
	// example:
	//
	// SYSTEM
	SubmitType *string `json:"SubmitType,omitempty" xml:"SubmitType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// -17163770809
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetCommitTime() *string {
	return s.CommitTime
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetGovernItemId() *int64 {
	return s.GovernItemId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetGovernObjectId() *int64 {
	return s.GovernObjectId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetIsRectify() *bool {
	return s.IsRectify
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetOwners() []*GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners {
	return s.Owners
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetProblem() *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	return s.Problem
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetRectifyName() *string {
	return s.RectifyName
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetRectifyStatus() *string {
	return s.RectifyStatus
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetRectifyUser() *string {
	return s.RectifyUser
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetRectifyUserName() *string {
	return s.RectifyUserName
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetRelatedKnowledge() []*GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	return s.RelatedKnowledge
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetStatus() *string {
	return s.Status
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetSubmitType() *string {
	return s.SubmitType
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetCommitTime(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.CommitTime = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetGovernItemId(v int64) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.GovernItemId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetGovernObjectId(v int64) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.GovernObjectId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetIsRectify(v bool) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.IsRectify = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetOwners(v []*GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.Owners = v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetProblem(v *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.Problem = v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetProperties(v map[string]interface{}) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.Properties = v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetRectifyId(v int64) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.RectifyId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetRectifyName(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.RectifyName = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetRectifyStatus(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.RectifyStatus = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetRectifyUser(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.RectifyUser = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetRectifyUserName(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.RectifyUserName = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetRelatedKnowledge(v []*GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.RelatedKnowledge = v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetStatus(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.Status = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetSubmitType(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.SubmitType = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) SetTenantId(v int64) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo {
	s.TenantId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfo) Validate() error {
	if s.Owners != nil {
		for _, item := range s.Owners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Problem != nil {
		if err := s.Problem.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedKnowledge != nil {
		for _, item := range s.RelatedKnowledge {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners struct {
	// The display name of the user.
	//
	// example:
	//
	// 龚恒菊2088822037866701
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 300006218
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) GetUserId() *string {
	return s.UserId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) SetDisplayName(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners {
	s.DisplayName = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) SetUserId(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners {
	s.UserId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoOwners) Validate() error {
	return dara.Validate(s)
}

type GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem struct {
	// The object ID.
	//
	// example:
	//
	// 36ea160807b14216b62a939327941e8b
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The ID of the parent object.
	//
	// example:
	//
	// 913836
	ParentObjectId *string `json:"ParentObjectId,omitempty" xml:"ParentObjectId,omitempty"`
	// The contact email for the governance issue.
	//
	// example:
	//
	// 178986769@gmail.com
	ProblemContactMail *string `json:"ProblemContactMail,omitempty" xml:"ProblemContactMail,omitempty"`
	// The other contact information for the governance issue.
	//
	// example:
	//
	// Jane
	ProblemContactOther *string `json:"ProblemContactOther,omitempty" xml:"ProblemContactOther,omitempty"`
	// The contact phone number for the governance issue.
	//
	// example:
	//
	// 12596178752
	ProblemContactPhone *string `json:"ProblemContactPhone,omitempty" xml:"ProblemContactPhone,omitempty"`
	// The description of the governance issue.
	//
	// example:
	//
	// Test issue
	ProblemDesc *string `json:"ProblemDesc,omitempty" xml:"ProblemDesc,omitempty"`
	// The submission method of the issue.
	//
	// example:
	//
	// SYSTEM
	ProblemSubmitType *string `json:"ProblemSubmitType,omitempty" xml:"ProblemSubmitType,omitempty"`
	// The user who submitted the issue.
	//
	// example:
	//
	// 300006218
	ProblemSubmitter *string `json:"ProblemSubmitter,omitempty" xml:"ProblemSubmitter,omitempty"`
	// The username of the user who submitted the issue.
	//
	// example:
	//
	// John
	ProblemSubmitterUserName *string `json:"ProblemSubmitterUserName,omitempty" xml:"ProblemSubmitterUserName,omitempty"`
	// The types of the governance issue.
	ProblemTypes []*string `json:"ProblemTypes,omitempty" xml:"ProblemTypes,omitempty" type:"Repeated"`
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetParentObjectId() *string {
	return s.ParentObjectId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemContactMail() *string {
	return s.ProblemContactMail
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemContactOther() *string {
	return s.ProblemContactOther
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemContactPhone() *string {
	return s.ProblemContactPhone
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemDesc() *string {
	return s.ProblemDesc
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemSubmitType() *string {
	return s.ProblemSubmitType
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemSubmitter() *string {
	return s.ProblemSubmitter
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemSubmitterUserName() *string {
	return s.ProblemSubmitterUserName
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) GetProblemTypes() []*string {
	return s.ProblemTypes
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetObjectId(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ObjectId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetParentObjectId(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ParentObjectId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemContactMail(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemContactMail = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemContactOther(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemContactOther = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemContactPhone(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemContactPhone = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemDesc(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemDesc = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemSubmitType(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemSubmitType = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemSubmitter(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemSubmitter = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemSubmitterUserName(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemSubmitterUserName = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) SetProblemTypes(v []*string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem {
	s.ProblemTypes = v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoProblem) Validate() error {
	return dara.Validate(s)
}

type GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge struct {
	// The cause of the issue.
	//
	// example:
	//
	// NC tag test
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The description.
	//
	// example:
	//
	// Yangchun Maternal and Child Health Cloud Service Space
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the knowledge entry.
	//
	// example:
	//
	// 522072057231
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The owner.
	//
	// example:
	//
	// leisatc
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// buc_166994
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The Solutions.
	//
	// example:
	//
	// RPBioOnly
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The title.
	//
	// example:
	//
	// Makassar International Eight Festival & Forum (F8 Makassar)
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GetCause() *string {
	return s.Cause
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GetDesc() *string {
	return s.Desc
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GetOwner() *string {
	return s.Owner
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GetSolution() *string {
	return s.Solution
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) GetTitle() *string {
	return s.Title
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) SetCause(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	s.Cause = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) SetDesc(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	s.Desc = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) SetKnowledgeId(v int64) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	s.KnowledgeId = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) SetOwner(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	s.Owner = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) SetOwnerName(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	s.OwnerName = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) SetSolution(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	s.Solution = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) SetTitle(v string) *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge {
	s.Title = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponseBodyGovernObjectInfoRelatedKnowledge) Validate() error {
	return dara.Validate(s)
}
