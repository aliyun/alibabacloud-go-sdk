// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGovernObjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGovernObjectsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListGovernObjectsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGovernObjectsResponseBody
	GetMessage() *string
	SetPageResult(v *ListGovernObjectsResponseBodyPageResult) *ListGovernObjectsResponseBody
	GetPageResult() *ListGovernObjectsResponseBodyPageResult
	SetRequestId(v string) *ListGovernObjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGovernObjectsResponseBody
	GetSuccess() *bool
}

type ListGovernObjectsResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// The paged query result.
	PageResult *ListGovernObjectsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
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

func (s ListGovernObjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGovernObjectsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGovernObjectsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGovernObjectsResponseBody) GetPageResult() *ListGovernObjectsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListGovernObjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGovernObjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGovernObjectsResponseBody) SetCode(v string) *ListGovernObjectsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGovernObjectsResponseBody) SetHttpStatusCode(v int32) *ListGovernObjectsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGovernObjectsResponseBody) SetMessage(v string) *ListGovernObjectsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGovernObjectsResponseBody) SetPageResult(v *ListGovernObjectsResponseBodyPageResult) *ListGovernObjectsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListGovernObjectsResponseBody) SetRequestId(v string) *ListGovernObjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGovernObjectsResponseBody) SetSuccess(v bool) *ListGovernObjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGovernObjectsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGovernObjectsResponseBodyPageResult struct {
	// The paged list of governance objects.
	Data []*ListGovernObjectsResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGovernObjectsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsResponseBodyPageResult) GetData() []*ListGovernObjectsResponseBodyPageResultData {
	return s.Data
}

func (s *ListGovernObjectsResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGovernObjectsResponseBodyPageResult) SetData(v []*ListGovernObjectsResponseBodyPageResultData) *ListGovernObjectsResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResult) SetTotalCount(v int64) *ListGovernObjectsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResult) Validate() error {
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

type ListGovernObjectsResponseBodyPageResultData struct {
	// The time when the record was reported.
	//
	// example:
	//
	// 2026-08-31 10:10:59
	CommitTime *string `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The governance object ID.
	//
	// example:
	//
	// 139487419630
	GovernItemId *int64 `json:"GovernItemId,omitempty" xml:"GovernItemId,omitempty"`
	// The governance object ID.
	//
	// example:
	//
	// -643545112181
	GovernObjectId *int64 `json:"GovernObjectId,omitempty" xml:"GovernObjectId,omitempty"`
	// Indicates whether rectification is in progress.
	IsRectify *bool `json:"IsRectify,omitempty" xml:"IsRectify,omitempty"`
	// The list of owners.
	Owners []*ListGovernObjectsResponseBodyPageResultDataOwners `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
	// The governance issue object.
	Problem *ListGovernObjectsResponseBodyPageResultDataProblem `json:"Problem,omitempty" xml:"Problem,omitempty" type:"Struct"`
	// The property values.
	//
	// example:
	//
	// {
	//
	//                     "gmt_create": "2026-08-31 06:00:22.296",
	//
	//                     "index_compute_type": "",
	//
	//                     "table_env": "PROD",
	//
	//                     "table_datasource_id": "1",
	//
	//                     "owner_id": "300006218",
	//
	//                     "index_catalog": "",
	//
	//                     "qd_feature_owner": "",
	//
	//                     "rule_task_start_time": "2026-08-31 06:00:00",
	//
	//                     "table_id": "odps.300023201.fashion_ads.api2mysql_demo",
	//
	//                     "rule_strength": "WEAK",
	//
	//                     "table_biz_unit_name": "LD_Fashion",
	//
	//                     "table_name": "fashion_ads.api2mysql_demo",
	//
	//                     "index_type": "",
	//
	//                     "table_datasource_from": "META_DATA",
	//
	//                     "datasource_type": "MaxCompute",
	//
	//                     "datasource_scope": "OFFLINE",
	//
	//                     "template_zh_tw_name": "欄位空值校正",
	//
	//                     "problem_submit_type": "SYSTEM",
	//
	//                     "template_type": "FIELD_NULL_VALUE_VALIDATE",
	//
	//                     "index_name_cn": "",
	//
	//                     "datasource_name": "Dataphin",
	//
	//                     "id": "918363",
	//
	//                     "index_id": "",
	//
	//                     "validate_partition": "ds=\\"20260831\\"",
	//
	//                     "index_owner_id": "",
	//
	//                     "datasource_owner": "300006218",
	//
	//                     "rule_name": "date_odps_test_2025-09-22 20:32:07",
	//
	//                     "watch_type": "TABLE",
	//
	//                     "validate_status": "NOT_PASS",
	//
	//                     "qd_feature_id": "",
	//
	//                     "table_desc": "api2mysql_demo",
	//
	//                     "is_ignore": "false",
	//
	//                     "rule_desc": "",
	//
	//                     "table_partitioned": "true",
	//
	//                     "template_owner": "300006218",
	//
	//                     "index_biz_unit_id": "",
	//
	//                     "table_biz_unit_id": "6865277495315392",
	//
	//                     "index_biz_unit_name": "",
	//
	//                     "watch_env": "PROD",
	//
	//                     "problem_contact_other": "",
	//
	//                     "status": "NEW",
	//
	//                     "datasource_from": "META_DATA",
	//
	//                     "table_project_name": "fashion_ads",
	//
	//                     "tenant_id": "300023201",
	//
	//                     "datasource_env": "PROD",
	//
	//                     "template_en_name": "Verify Field Null Values",
	//
	//                     "commit_time": "2026-08-31 06:00:21.135",
	//
	//                     "gmt_modified": "2026-08-31 06:00:22.296",
	//
	//                     "qd_feature_name": "",
	//
	//                     "table_catalog": "fashion_ads",
	//
	//                     "rule_status": "ENABLE",
	//
	//                     "problem_contact_mail": "",
	//
	//                     "rule_task_status": "SUCCESS",
	//
	//                     "rule_validate_object_type": "COLUMN",
	//
	//                     "watch_status": "ENABLE",
	//
	//                     "index_guid": "",
	//
	//                     "system_template": "true",
	//
	//                     "quality_owners": "300006218",
	//
	//                     "index_name": "",
	//
	//                     "problem_contact_phone": "",
	//
	//                     "watch_task_id": "8325922",
	//
	//                     "rule_task_id": "8328383",
	//
	//                     "index_desc": "",
	//
	//                     "table_type": "PHYSICAL_TABLE",
	//
	//                     "table_project_id": "6865331520706176",
	//
	//                     "trace_id": "time:4326023",
	//
	//                     "datasource_id": "1",
	//
	//                     "qd_feature_code": "",
	//
	//                     "rule_validate_object_name": "date_odps",
	//
	//                     "problem_desc": "date_odps_test_2025-09-22 20:32:07",
	//
	//                     "table_owner": "300006218",
	//
	//                     "quality_owner_groups": "",
	//
	//                     "govern_item_id": "100",
	//
	//                     "rule_catalogs": "COMPLETENESS",
	//
	//                     "table_datasource_type": "MAX_COMPUTE",
	//
	//                     "template_zh_cn_name": "字段空值校验",
	//
	//                     "rule_task_biz_date": "20260831 06:00:00",
	//
	//                     "problem_types": "[{\\"value\\":\\"COMPLETENESS\\"}]",
	//
	//                     "rule_id": "4322944",
	//
	//                     "problem_attachment_file_ids": "null",
	//
	//                     "watch_name": "",
	//
	//                     "template_name": "",
	//
	//                     "schedule_params": "",
	//
	//                     "is_rectify": "false",
	//
	//                     "rule_task_end_time": "",
	//
	//                     "watch_id": "3841908",
	//
	//                     "validate_result": "false",
	//
	//                     "qd_feature_guid": "",
	//
	//                     "govern_object_id": "913836",
	//
	//                     "template_id": "100"
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The rectification ID.
	//
	// example:
	//
	// -787032739353
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
	// The name of the rectification.
	//
	// example:
	//
	// Test rectification
	RectifyName *string `json:"RectifyName,omitempty" xml:"RectifyName,omitempty"`
	// The rectification status.
	//
	// example:
	//
	// NEW
	RectifyStatus *string `json:"RectifyStatus,omitempty" xml:"RectifyStatus,omitempty"`
	// The related knowledge base entries.
	RelatedKnowledge []*ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge `json:"RelatedKnowledge,omitempty" xml:"RelatedKnowledge,omitempty" type:"Repeated"`
	// The status of the governance object.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListGovernObjectsResponseBodyPageResultData) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetCommitTime() *string {
	return s.CommitTime
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetGovernItemId() *int64 {
	return s.GovernItemId
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetGovernObjectId() *int64 {
	return s.GovernObjectId
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetIsRectify() *bool {
	return s.IsRectify
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetOwners() []*ListGovernObjectsResponseBodyPageResultDataOwners {
	return s.Owners
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetProblem() *ListGovernObjectsResponseBodyPageResultDataProblem {
	return s.Problem
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetRectifyName() *string {
	return s.RectifyName
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetRectifyStatus() *string {
	return s.RectifyStatus
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetRelatedKnowledge() []*ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	return s.RelatedKnowledge
}

func (s *ListGovernObjectsResponseBodyPageResultData) GetStatus() *string {
	return s.Status
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetCommitTime(v string) *ListGovernObjectsResponseBodyPageResultData {
	s.CommitTime = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetGovernItemId(v int64) *ListGovernObjectsResponseBodyPageResultData {
	s.GovernItemId = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetGovernObjectId(v int64) *ListGovernObjectsResponseBodyPageResultData {
	s.GovernObjectId = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetIsRectify(v bool) *ListGovernObjectsResponseBodyPageResultData {
	s.IsRectify = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetOwners(v []*ListGovernObjectsResponseBodyPageResultDataOwners) *ListGovernObjectsResponseBodyPageResultData {
	s.Owners = v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetProblem(v *ListGovernObjectsResponseBodyPageResultDataProblem) *ListGovernObjectsResponseBodyPageResultData {
	s.Problem = v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetProperties(v map[string]interface{}) *ListGovernObjectsResponseBodyPageResultData {
	s.Properties = v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetRectifyId(v int64) *ListGovernObjectsResponseBodyPageResultData {
	s.RectifyId = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetRectifyName(v string) *ListGovernObjectsResponseBodyPageResultData {
	s.RectifyName = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetRectifyStatus(v string) *ListGovernObjectsResponseBodyPageResultData {
	s.RectifyStatus = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetRelatedKnowledge(v []*ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) *ListGovernObjectsResponseBodyPageResultData {
	s.RelatedKnowledge = v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) SetStatus(v string) *ListGovernObjectsResponseBodyPageResultData {
	s.Status = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultData) Validate() error {
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

type ListGovernObjectsResponseBodyPageResultDataOwners struct {
	// The display name of the user.
	//
	// example:
	//
	// Yang Jing 2088252351182803
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListGovernObjectsResponseBodyPageResultDataOwners) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsResponseBodyPageResultDataOwners) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsResponseBodyPageResultDataOwners) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListGovernObjectsResponseBodyPageResultDataOwners) GetUserId() *string {
	return s.UserId
}

func (s *ListGovernObjectsResponseBodyPageResultDataOwners) SetDisplayName(v string) *ListGovernObjectsResponseBodyPageResultDataOwners {
	s.DisplayName = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataOwners) SetUserId(v string) *ListGovernObjectsResponseBodyPageResultDataOwners {
	s.UserId = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataOwners) Validate() error {
	return dara.Validate(s)
}

type ListGovernObjectsResponseBodyPageResultDataProblem struct {
	// The object ID.
	//
	// example:
	//
	// 9223058119411358258
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
	// 126983612986391@gamail.com
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
	// 16278902467
	ProblemContactPhone *string `json:"ProblemContactPhone,omitempty" xml:"ProblemContactPhone,omitempty"`
	// The description of the governance issue.
	//
	// example:
	//
	// Governance issue description
	ProblemDesc *string `json:"ProblemDesc,omitempty" xml:"ProblemDesc,omitempty"`
	// The submission type of the issue.
	//
	// example:
	//
	// SYSTEM
	ProblemSubmitType *string `json:"ProblemSubmitType,omitempty" xml:"ProblemSubmitType,omitempty"`
	// The user who submitted the issue.
	//
	// example:
	//
	// 123456
	ProblemSubmitter *string `json:"ProblemSubmitter,omitempty" xml:"ProblemSubmitter,omitempty"`
	// The username of the issue submitter.
	//
	// example:
	//
	// John
	ProblemSubmitterUserName *string `json:"ProblemSubmitterUserName,omitempty" xml:"ProblemSubmitterUserName,omitempty"`
	// The types of the governance issue.
	ProblemTypes []*string `json:"ProblemTypes,omitempty" xml:"ProblemTypes,omitempty" type:"Repeated"`
}

func (s ListGovernObjectsResponseBodyPageResultDataProblem) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsResponseBodyPageResultDataProblem) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetParentObjectId() *string {
	return s.ParentObjectId
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemContactMail() *string {
	return s.ProblemContactMail
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemContactOther() *string {
	return s.ProblemContactOther
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemContactPhone() *string {
	return s.ProblemContactPhone
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemDesc() *string {
	return s.ProblemDesc
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemSubmitType() *string {
	return s.ProblemSubmitType
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemSubmitter() *string {
	return s.ProblemSubmitter
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemSubmitterUserName() *string {
	return s.ProblemSubmitterUserName
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) GetProblemTypes() []*string {
	return s.ProblemTypes
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetObjectId(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ObjectId = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetParentObjectId(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ParentObjectId = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemContactMail(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemContactMail = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemContactOther(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemContactOther = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemContactPhone(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemContactPhone = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemDesc(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemDesc = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemSubmitType(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemSubmitType = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemSubmitter(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemSubmitter = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemSubmitterUserName(v string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemSubmitterUserName = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) SetProblemTypes(v []*string) *ListGovernObjectsResponseBodyPageResultDataProblem {
	s.ProblemTypes = v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataProblem) Validate() error {
	return dara.Validate(s)
}

type ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge struct {
	// The cause of the issue.
	//
	// example:
	//
	// Host exception. The instance was migrated with data loss
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The description.
	//
	// example:
	//
	// Store scheduled power on/off business domain canary list\\n
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The knowledge entry ID.
	//
	// example:
	//
	// -341426256859
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The owner.
	//
	// example:
	//
	// -mnneiiwtemj-wjuggee
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// buc_459782
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The Solutions.
	//
	// example:
	//
	// {\\"Type\\":\\"text\\",\\"Value\\":\\"OSS public network access\\"}
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The title.
	//
	// example:
	//
	// Makassar International Eight Festival & Forum (F8 Makassar)
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GetCause() *string {
	return s.Cause
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GetDesc() *string {
	return s.Desc
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GetOwner() *string {
	return s.Owner
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GetSolution() *string {
	return s.Solution
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) GetTitle() *string {
	return s.Title
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) SetCause(v string) *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	s.Cause = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) SetDesc(v string) *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	s.Desc = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) SetKnowledgeId(v int64) *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	s.KnowledgeId = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) SetOwner(v string) *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	s.Owner = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) SetOwnerName(v string) *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	s.OwnerName = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) SetSolution(v string) *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	s.Solution = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) SetTitle(v string) *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge {
	s.Title = &v
	return s
}

func (s *ListGovernObjectsResponseBodyPageResultDataRelatedKnowledge) Validate() error {
	return dara.Validate(s)
}
