// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRectifyFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRectifyFlowResponseBody
	GetCode() *string
	SetData(v *GetRectifyFlowResponseBodyData) *GetRectifyFlowResponseBody
	GetData() *GetRectifyFlowResponseBodyData
	SetHttpStatusCode(v int32) *GetRectifyFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRectifyFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRectifyFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRectifyFlowResponseBody
	GetSuccess() *bool
}

type GetRectifyFlowResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The rectification details.
	Data *GetRectifyFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRectifyFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRectifyFlowResponseBody) GetData() *GetRectifyFlowResponseBodyData {
	return s.Data
}

func (s *GetRectifyFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRectifyFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRectifyFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRectifyFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRectifyFlowResponseBody) SetCode(v string) *GetRectifyFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetRectifyFlowResponseBody) SetData(v *GetRectifyFlowResponseBodyData) *GetRectifyFlowResponseBody {
	s.Data = v
	return s
}

func (s *GetRectifyFlowResponseBody) SetHttpStatusCode(v int32) *GetRectifyFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRectifyFlowResponseBody) SetMessage(v string) *GetRectifyFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetRectifyFlowResponseBody) SetRequestId(v string) *GetRectifyFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRectifyFlowResponseBody) SetSuccess(v bool) *GetRectifyFlowResponseBody {
	s.Success = &v
	return s
}

func (s *GetRectifyFlowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRectifyFlowResponseBodyData struct {
	// The completion time as a UNIX timestamp in milliseconds. This parameter is empty if the rectification is not completed.
	//
	// example:
	//
	// 57753888419
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The pagination information for rectification issues. If the rectification has been accepted, the acceptance snapshot is used.
	Issues *GetRectifyFlowResponseBodyDataIssues `json:"Issues,omitempty" xml:"Issues,omitempty" type:"Struct"`
	// The total number of current issues in the process.
	//
	// example:
	//
	// 415492404179
	ProblemCount *int64 `json:"ProblemCount,omitempty" xml:"ProblemCount,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 14561234364
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
	// The process name.
	//
	// example:
	//
	// test
	RectifyName *string `json:"RectifyName,omitempty" xml:"RectifyName,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	RectifyPriority *string `json:"RectifyPriority,omitempty" xml:"RectifyPriority,omitempty"`
	// The rectifier ID.
	//
	// example:
	//
	// 3001
	RectifyUser *string `json:"RectifyUser,omitempty" xml:"RectifyUser,omitempty"`
	// The rectifier name.
	//
	// example:
	//
	// test
	RectifyUserName *string `json:"RectifyUserName,omitempty" xml:"RectifyUserName,omitempty"`
	// The start time, in UNIX milliseconds.
	//
	// example:
	//
	// 496002949493
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The process status.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The initiator ID.
	//
	// example:
	//
	// 3001
	SubmitUser *string `json:"SubmitUser,omitempty" xml:"SubmitUser,omitempty"`
	// The initiator name.
	//
	// example:
	//
	// test
	SubmitUserName *string `json:"SubmitUserName,omitempty" xml:"SubmitUserName,omitempty"`
	// The verifier ID.
	//
	// example:
	//
	// 3001
	VerifyUser *string `json:"VerifyUser,omitempty" xml:"VerifyUser,omitempty"`
	// The verifier name.
	//
	// example:
	//
	// test
	VerifyUserName *string `json:"VerifyUserName,omitempty" xml:"VerifyUserName,omitempty"`
}

func (s GetRectifyFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponseBodyData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetRectifyFlowResponseBodyData) GetIssues() *GetRectifyFlowResponseBodyDataIssues {
	return s.Issues
}

func (s *GetRectifyFlowResponseBodyData) GetProblemCount() *int64 {
	return s.ProblemCount
}

func (s *GetRectifyFlowResponseBodyData) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *GetRectifyFlowResponseBodyData) GetRectifyName() *string {
	return s.RectifyName
}

func (s *GetRectifyFlowResponseBodyData) GetRectifyPriority() *string {
	return s.RectifyPriority
}

func (s *GetRectifyFlowResponseBodyData) GetRectifyUser() *string {
	return s.RectifyUser
}

func (s *GetRectifyFlowResponseBodyData) GetRectifyUserName() *string {
	return s.RectifyUserName
}

func (s *GetRectifyFlowResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRectifyFlowResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetRectifyFlowResponseBodyData) GetSubmitUser() *string {
	return s.SubmitUser
}

func (s *GetRectifyFlowResponseBodyData) GetSubmitUserName() *string {
	return s.SubmitUserName
}

func (s *GetRectifyFlowResponseBodyData) GetVerifyUser() *string {
	return s.VerifyUser
}

func (s *GetRectifyFlowResponseBodyData) GetVerifyUserName() *string {
	return s.VerifyUserName
}

func (s *GetRectifyFlowResponseBodyData) SetFinishTime(v int64) *GetRectifyFlowResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetIssues(v *GetRectifyFlowResponseBodyDataIssues) *GetRectifyFlowResponseBodyData {
	s.Issues = v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetProblemCount(v int64) *GetRectifyFlowResponseBodyData {
	s.ProblemCount = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetRectifyId(v int64) *GetRectifyFlowResponseBodyData {
	s.RectifyId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetRectifyName(v string) *GetRectifyFlowResponseBodyData {
	s.RectifyName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetRectifyPriority(v string) *GetRectifyFlowResponseBodyData {
	s.RectifyPriority = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetRectifyUser(v string) *GetRectifyFlowResponseBodyData {
	s.RectifyUser = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetRectifyUserName(v string) *GetRectifyFlowResponseBodyData {
	s.RectifyUserName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetStartTime(v int64) *GetRectifyFlowResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetStatus(v string) *GetRectifyFlowResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetSubmitUser(v string) *GetRectifyFlowResponseBodyData {
	s.SubmitUser = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetSubmitUserName(v string) *GetRectifyFlowResponseBodyData {
	s.SubmitUserName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetVerifyUser(v string) *GetRectifyFlowResponseBodyData {
	s.VerifyUser = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) SetVerifyUserName(v string) *GetRectifyFlowResponseBodyData {
	s.VerifyUserName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyData) Validate() error {
	if s.Issues != nil {
		if err := s.Issues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRectifyFlowResponseBodyDataIssues struct {
	// The issues on the current page. The fields are consistent with the issue list.
	Data []*GetRectifyFlowResponseBodyDataIssuesData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of issues.
	//
	// example:
	//
	// 467439222172
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetRectifyFlowResponseBodyDataIssues) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponseBodyDataIssues) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponseBodyDataIssues) GetData() []*GetRectifyFlowResponseBodyDataIssuesData {
	return s.Data
}

func (s *GetRectifyFlowResponseBodyDataIssues) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetRectifyFlowResponseBodyDataIssues) SetData(v []*GetRectifyFlowResponseBodyDataIssuesData) *GetRectifyFlowResponseBodyDataIssues {
	s.Data = v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssues) SetTotalCount(v int64) *GetRectifyFlowResponseBodyDataIssues {
	s.TotalCount = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssues) Validate() error {
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

type GetRectifyFlowResponseBodyDataIssuesData struct {
	// The submission time.
	//
	// example:
	//
	// 2026-10-22
	CommitTime *string `json:"CommitTime,omitempty" xml:"CommitTime,omitempty"`
	// The ID of the governance object.
	//
	// example:
	//
	// 768244593451
	GovernItemId *int64 `json:"GovernItemId,omitempty" xml:"GovernItemId,omitempty"`
	// The ID of the governance object.
	//
	// example:
	//
	// 28783986302
	GovernObjectId *int64 `json:"GovernObjectId,omitempty" xml:"GovernObjectId,omitempty"`
	// Specifies whether the rectification is in progress.
	//
	// example:
	//
	// true
	IsRectify *bool `json:"IsRectify,omitempty" xml:"IsRectify,omitempty"`
	// The list of owners.
	Owners []*GetRectifyFlowResponseBodyDataIssuesDataOwners `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
	// The governance issue object.
	Problem *GetRectifyFlowResponseBodyDataIssuesDataProblem `json:"Problem,omitempty" xml:"Problem,omitempty" type:"Struct"`
	// The property map of the governance object.
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The ID of the rectification.
	//
	// example:
	//
	// 813969529153
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
	// The name of the rectification.
	//
	// example:
	//
	// test
	RectifyName *string `json:"RectifyName,omitempty" xml:"RectifyName,omitempty"`
	// The status of the rectification.
	//
	// example:
	//
	// SUCCESS
	RectifyStatus *string `json:"RectifyStatus,omitempty" xml:"RectifyStatus,omitempty"`
	// The ID of the user who performs the rectification.
	//
	// example:
	//
	// 3001
	RectifyUser *string `json:"RectifyUser,omitempty" xml:"RectifyUser,omitempty"`
	// The display name of the user who performs the rectification.
	//
	// example:
	//
	// test
	RectifyUserName *string `json:"RectifyUserName,omitempty" xml:"RectifyUserName,omitempty"`
	// The related knowledge base.
	RelatedKnowledge []*GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge `json:"RelatedKnowledge,omitempty" xml:"RelatedKnowledge,omitempty" type:"Repeated"`
	// The status of the governance object.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The submission method.
	//
	// example:
	//
	// OPEN API
	SubmitType *string `json:"SubmitType,omitempty" xml:"SubmitType,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 575244494646
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetRectifyFlowResponseBodyDataIssuesData) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponseBodyDataIssuesData) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetCommitTime() *string {
	return s.CommitTime
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetGovernItemId() *int64 {
	return s.GovernItemId
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetGovernObjectId() *int64 {
	return s.GovernObjectId
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetIsRectify() *bool {
	return s.IsRectify
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetOwners() []*GetRectifyFlowResponseBodyDataIssuesDataOwners {
	return s.Owners
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetProblem() *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	return s.Problem
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetProperties() map[string]*string {
	return s.Properties
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetRectifyName() *string {
	return s.RectifyName
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetRectifyStatus() *string {
	return s.RectifyStatus
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetRectifyUser() *string {
	return s.RectifyUser
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetRectifyUserName() *string {
	return s.RectifyUserName
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetRelatedKnowledge() []*GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	return s.RelatedKnowledge
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetStatus() *string {
	return s.Status
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetSubmitType() *string {
	return s.SubmitType
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetCommitTime(v string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.CommitTime = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetGovernItemId(v int64) *GetRectifyFlowResponseBodyDataIssuesData {
	s.GovernItemId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetGovernObjectId(v int64) *GetRectifyFlowResponseBodyDataIssuesData {
	s.GovernObjectId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetIsRectify(v bool) *GetRectifyFlowResponseBodyDataIssuesData {
	s.IsRectify = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetOwners(v []*GetRectifyFlowResponseBodyDataIssuesDataOwners) *GetRectifyFlowResponseBodyDataIssuesData {
	s.Owners = v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetProblem(v *GetRectifyFlowResponseBodyDataIssuesDataProblem) *GetRectifyFlowResponseBodyDataIssuesData {
	s.Problem = v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetProperties(v map[string]*string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.Properties = v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetRectifyId(v int64) *GetRectifyFlowResponseBodyDataIssuesData {
	s.RectifyId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetRectifyName(v string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.RectifyName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetRectifyStatus(v string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.RectifyStatus = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetRectifyUser(v string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.RectifyUser = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetRectifyUserName(v string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.RectifyUserName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetRelatedKnowledge(v []*GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) *GetRectifyFlowResponseBodyDataIssuesData {
	s.RelatedKnowledge = v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetStatus(v string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.Status = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetSubmitType(v string) *GetRectifyFlowResponseBodyDataIssuesData {
	s.SubmitType = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) SetTenantId(v int64) *GetRectifyFlowResponseBodyDataIssuesData {
	s.TenantId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesData) Validate() error {
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

type GetRectifyFlowResponseBodyDataIssuesDataOwners struct {
	// The display name of the user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 301
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetRectifyFlowResponseBodyDataIssuesDataOwners) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponseBodyDataIssuesDataOwners) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataOwners) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataOwners) GetUserId() *string {
	return s.UserId
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataOwners) SetDisplayName(v string) *GetRectifyFlowResponseBodyDataIssuesDataOwners {
	s.DisplayName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataOwners) SetUserId(v string) *GetRectifyFlowResponseBodyDataIssuesDataOwners {
	s.UserId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataOwners) Validate() error {
	return dara.Validate(s)
}

type GetRectifyFlowResponseBodyDataIssuesDataProblem struct {
	// The ID of the object.
	//
	// example:
	//
	// 123
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The ID of the parent object.
	//
	// example:
	//
	// 123
	ParentObjectId *string `json:"ParentObjectId,omitempty" xml:"ParentObjectId,omitempty"`
	// The contact email for the governance issue.
	//
	// example:
	//
	// 123@test.com
	ProblemContactMail *string `json:"ProblemContactMail,omitempty" xml:"ProblemContactMail,omitempty"`
	// Other contact information for the governance issue.
	//
	// example:
	//
	// 12345
	ProblemContactOther *string `json:"ProblemContactOther,omitempty" xml:"ProblemContactOther,omitempty"`
	// The contact phone number for the governance issue.
	//
	// example:
	//
	// 123
	ProblemContactPhone *string `json:"ProblemContactPhone,omitempty" xml:"ProblemContactPhone,omitempty"`
	// The description of the governance issue.
	//
	// example:
	//
	// test
	ProblemDesc *string `json:"ProblemDesc,omitempty" xml:"ProblemDesc,omitempty"`
	// The submission method of the issue.
	//
	// example:
	//
	// OPENAPI
	ProblemSubmitType *string `json:"ProblemSubmitType,omitempty" xml:"ProblemSubmitType,omitempty"`
	// The user who submitted the issue.
	//
	// example:
	//
	// 3001
	ProblemSubmitter *string `json:"ProblemSubmitter,omitempty" xml:"ProblemSubmitter,omitempty"`
	// The name of the user who submitted the issue.
	//
	// example:
	//
	// test
	ProblemSubmitterUserName *string `json:"ProblemSubmitterUserName,omitempty" xml:"ProblemSubmitterUserName,omitempty"`
	// The type of the governance issue.
	ProblemTypes []*string `json:"ProblemTypes,omitempty" xml:"ProblemTypes,omitempty" type:"Repeated"`
}

func (s GetRectifyFlowResponseBodyDataIssuesDataProblem) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponseBodyDataIssuesDataProblem) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetParentObjectId() *string {
	return s.ParentObjectId
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemContactMail() *string {
	return s.ProblemContactMail
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemContactOther() *string {
	return s.ProblemContactOther
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemContactPhone() *string {
	return s.ProblemContactPhone
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemDesc() *string {
	return s.ProblemDesc
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemSubmitType() *string {
	return s.ProblemSubmitType
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemSubmitter() *string {
	return s.ProblemSubmitter
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemSubmitterUserName() *string {
	return s.ProblemSubmitterUserName
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) GetProblemTypes() []*string {
	return s.ProblemTypes
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetObjectId(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ObjectId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetParentObjectId(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ParentObjectId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemContactMail(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemContactMail = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemContactOther(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemContactOther = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemContactPhone(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemContactPhone = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemDesc(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemDesc = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemSubmitType(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemSubmitType = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemSubmitter(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemSubmitter = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemSubmitterUserName(v string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemSubmitterUserName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) SetProblemTypes(v []*string) *GetRectifyFlowResponseBodyDataIssuesDataProblem {
	s.ProblemTypes = v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataProblem) Validate() error {
	return dara.Validate(s)
}

type GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge struct {
	// The cause of the issue.
	//
	// example:
	//
	// test
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The ID of the knowledge.
	//
	// example:
	//
	// 649144693069
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// The owner.
	//
	// example:
	//
	// 3001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The Solutions.
	//
	// example:
	//
	// teset
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The title.
	//
	// example:
	//
	// teset
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GetCause() *string {
	return s.Cause
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GetDesc() *string {
	return s.Desc
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GetOwner() *string {
	return s.Owner
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GetSolution() *string {
	return s.Solution
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) GetTitle() *string {
	return s.Title
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) SetCause(v string) *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	s.Cause = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) SetDesc(v string) *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	s.Desc = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) SetKnowledgeId(v int64) *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	s.KnowledgeId = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) SetOwner(v string) *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	s.Owner = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) SetOwnerName(v string) *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	s.OwnerName = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) SetSolution(v string) *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	s.Solution = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) SetTitle(v string) *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge {
	s.Title = &v
	return s
}

func (s *GetRectifyFlowResponseBodyDataIssuesDataRelatedKnowledge) Validate() error {
	return dara.Validate(s)
}
