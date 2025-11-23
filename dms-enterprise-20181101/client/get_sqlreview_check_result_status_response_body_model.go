// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSQLReviewCheckResultStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResultStatus(v *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) *GetSQLReviewCheckResultStatusResponseBody
	GetCheckResultStatus() *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus
	SetErrorCode(v string) *GetSQLReviewCheckResultStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSQLReviewCheckResultStatusResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetSQLReviewCheckResultStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSQLReviewCheckResultStatusResponseBody
	GetSuccess() *bool
}

type GetSQLReviewCheckResultStatusResponseBody struct {
	// The result of the SQL review.
	CheckResultStatus *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus `json:"CheckResultStatus,omitempty" xml:"CheckResultStatus,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSQLReviewCheckResultStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBody) GetCheckResultStatus() *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	return s.CheckResultStatus
}

func (s *GetSQLReviewCheckResultStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSQLReviewCheckResultStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSQLReviewCheckResultStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSQLReviewCheckResultStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetCheckResultStatus(v *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) *GetSQLReviewCheckResultStatusResponseBody {
	s.CheckResultStatus = v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetErrorCode(v string) *GetSQLReviewCheckResultStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetErrorMessage(v string) *GetSQLReviewCheckResultStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetRequestId(v string) *GetSQLReviewCheckResultStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) SetSuccess(v bool) *GetSQLReviewCheckResultStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBody) Validate() error {
	if s.CheckResultStatus != nil {
		if err := s.CheckResultStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus struct {
	// The result of the SQL status check.
	CheckStatusResult *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult `json:"CheckStatusResult,omitempty" xml:"CheckStatusResult,omitempty" type:"Struct"`
	// The number of SQL statements that were reviewed.
	//
	// example:
	//
	// 10
	CheckedCount *int64 `json:"CheckedCount,omitempty" xml:"CheckedCount,omitempty"`
	// The optimization suggestion for SQL statements.
	SQLReviewResult *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult `json:"SQLReviewResult,omitempty" xml:"SQLReviewResult,omitempty" type:"Struct"`
	// The total number of SQL statements.
	//
	// example:
	//
	// 10
	TotalSQLCount *int64 `json:"TotalSQLCount,omitempty" xml:"TotalSQLCount,omitempty"`
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) GetCheckStatusResult() *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	return s.CheckStatusResult
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) GetCheckedCount() *int64 {
	return s.CheckedCount
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) GetSQLReviewResult() *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	return s.SQLReviewResult
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) GetTotalSQLCount() *int64 {
	return s.TotalSQLCount
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetCheckStatusResult(v *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.CheckStatusResult = v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetCheckedCount(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.CheckedCount = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetSQLReviewResult(v *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.SQLReviewResult = v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) SetTotalSQLCount(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus {
	s.TotalSQLCount = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatus) Validate() error {
	if s.CheckStatusResult != nil {
		if err := s.CheckStatusResult.Validate(); err != nil {
			return err
		}
	}
	if s.SQLReviewResult != nil {
		if err := s.SQLReviewResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult struct {
	// The number of SQL statements that failed to pass the review.
	//
	// example:
	//
	// 1
	CheckNotPass *int64 `json:"CheckNotPass,omitempty" xml:"CheckNotPass,omitempty"`
	// The number of SQL statements that passed the review.
	//
	// example:
	//
	// 8
	CheckPass *int64 `json:"CheckPass,omitempty" xml:"CheckPass,omitempty"`
	// The number of SQL statements that failed to pass the manual review.
	//
	// example:
	//
	// 0
	ForceNotPass *int64 `json:"ForceNotPass,omitempty" xml:"ForceNotPass,omitempty"`
	// The number of SQL statements that passed the manual review.
	//
	// example:
	//
	// 1
	ForcePass *int64 `json:"ForcePass,omitempty" xml:"ForcePass,omitempty"`
	// The number of SQL statements to be reviewed.
	//
	// example:
	//
	// 0
	New *int64 `json:"New,omitempty" xml:"New,omitempty"`
	// The number of abnormal SQL statements.
	//
	// example:
	//
	// 0
	Unknown *int64 `json:"Unknown,omitempty" xml:"Unknown,omitempty"`
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GetCheckNotPass() *int64 {
	return s.CheckNotPass
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GetCheckPass() *int64 {
	return s.CheckPass
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GetForceNotPass() *int64 {
	return s.ForceNotPass
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GetForcePass() *int64 {
	return s.ForcePass
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GetNew() *int64 {
	return s.New
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) GetUnknown() *int64 {
	return s.Unknown
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetCheckNotPass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.CheckNotPass = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetCheckPass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.CheckPass = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetForceNotPass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.ForceNotPass = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetForcePass(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.ForcePass = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetNew(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.New = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) SetUnknown(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult {
	s.Unknown = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusCheckStatusResult) Validate() error {
	return dara.Validate(s)
}

type GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult struct {
	// The number of SQL statements that must be modified.
	//
	// example:
	//
	// 1
	MustImprove *int64 `json:"MustImprove,omitempty" xml:"MustImprove,omitempty"`
	// The number of SQL statements that have potential issues.
	//
	// example:
	//
	// 0
	PotentialIssue *int64 `json:"PotentialIssue,omitempty" xml:"PotentialIssue,omitempty"`
	// The number of SQL statements that can be modified.
	//
	// example:
	//
	// 3
	SuggestImprove *int64 `json:"SuggestImprove,omitempty" xml:"SuggestImprove,omitempty"`
	// The number of SQL statements that can use indexes.
	//
	// example:
	//
	// 2
	TableIndexSuggest *int64 `json:"TableIndexSuggest,omitempty" xml:"TableIndexSuggest,omitempty"`
	// The number of SQL statements that can be used for lock-free data changes.
	//
	// example:
	//
	// 0
	UseDmsDmlUnlock *int64 `json:"UseDmsDmlUnlock,omitempty" xml:"UseDmsDmlUnlock,omitempty"`
	// The number of SQL statements that can be used for lock-free schema changes.
	//
	// example:
	//
	// 0
	UseDmsToolkit *int64 `json:"UseDmsToolkit,omitempty" xml:"UseDmsToolkit,omitempty"`
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GoString() string {
	return s.String()
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GetMustImprove() *int64 {
	return s.MustImprove
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GetPotentialIssue() *int64 {
	return s.PotentialIssue
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GetSuggestImprove() *int64 {
	return s.SuggestImprove
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GetTableIndexSuggest() *int64 {
	return s.TableIndexSuggest
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GetUseDmsDmlUnlock() *int64 {
	return s.UseDmsDmlUnlock
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) GetUseDmsToolkit() *int64 {
	return s.UseDmsToolkit
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetMustImprove(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.MustImprove = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetPotentialIssue(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.PotentialIssue = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetSuggestImprove(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.SuggestImprove = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetTableIndexSuggest(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.TableIndexSuggest = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetUseDmsDmlUnlock(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.UseDmsDmlUnlock = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) SetUseDmsToolkit(v int64) *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult {
	s.UseDmsToolkit = &v
	return s
}

func (s *GetSQLReviewCheckResultStatusResponseBodyCheckResultStatusSQLReviewResult) Validate() error {
	return dara.Validate(s)
}
