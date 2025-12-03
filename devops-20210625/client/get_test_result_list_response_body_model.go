// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTestResultListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTestResultListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetTestResultListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetTestResultListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTestResultListResponseBody
	GetSuccess() *bool
	SetTestResult(v []*GetTestResultListResponseBodyTestResult) *GetTestResultListResponseBody
	GetTestResult() []*GetTestResultListResponseBodyTestResult
}

type GetTestResultListResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success    *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
	TestResult []*GetTestResultListResponseBodyTestResult `json:"testResult,omitempty" xml:"testResult,omitempty" type:"Repeated"`
}

func (s GetTestResultListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTestResultListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTestResultListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTestResultListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetTestResultListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTestResultListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTestResultListResponseBody) GetTestResult() []*GetTestResultListResponseBodyTestResult {
	return s.TestResult
}

func (s *GetTestResultListResponseBody) SetErrorCode(v string) *GetTestResultListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTestResultListResponseBody) SetErrorMsg(v string) *GetTestResultListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTestResultListResponseBody) SetRequestId(v string) *GetTestResultListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTestResultListResponseBody) SetSuccess(v bool) *GetTestResultListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTestResultListResponseBody) SetTestResult(v []*GetTestResultListResponseBodyTestResult) *GetTestResultListResponseBody {
	s.TestResult = v
	return s
}

func (s *GetTestResultListResponseBody) Validate() error {
	if s.TestResult != nil {
		for _, item := range s.TestResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTestResultListResponseBodyTestResult struct {
	AssignedTo *GetTestResultListResponseBodyTestResultAssignedTo `json:"assignedTo,omitempty" xml:"assignedTo,omitempty" type:"Struct"`
	// example:
	//
	// 8
	BugCount *int64 `json:"bugCount,omitempty" xml:"bugCount,omitempty"`
	// example:
	//
	// Req
	CategoryIdentifier *string                                                `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	CustomFields       []*GetTestResultListResponseBodyTestResultCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	GmtCreate          *int64                                                 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	SpaceIdentifier    *string                                                `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// 测试工作项
	Subject             *string                                                    `json:"subject,omitempty" xml:"subject,omitempty"`
	TestResultExecutor  *GetTestResultListResponseBodyTestResultTestResultExecutor `json:"testResultExecutor,omitempty" xml:"testResultExecutor,omitempty" type:"Struct"`
	TestResultGmtCreate *int64                                                     `json:"testResultGmtCreate,omitempty" xml:"testResultGmtCreate,omitempty"`
	// example:
	//
	// a8bxxxxxxxxxxxxxxxx54
	TestResultIdentifier *string `json:"testResultIdentifier,omitempty" xml:"testResultIdentifier,omitempty"`
	// example:
	//
	// TO DO
	TestResultStatus   *string `json:"testResultStatus,omitempty" xml:"testResultStatus,omitempty"`
	TestcaseIdentifier *string `json:"testcaseIdentifier,omitempty" xml:"testcaseIdentifier,omitempty"`
}

func (s GetTestResultListResponseBodyTestResult) String() string {
	return dara.Prettify(s)
}

func (s GetTestResultListResponseBodyTestResult) GoString() string {
	return s.String()
}

func (s *GetTestResultListResponseBodyTestResult) GetAssignedTo() *GetTestResultListResponseBodyTestResultAssignedTo {
	return s.AssignedTo
}

func (s *GetTestResultListResponseBodyTestResult) GetBugCount() *int64 {
	return s.BugCount
}

func (s *GetTestResultListResponseBodyTestResult) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *GetTestResultListResponseBodyTestResult) GetCustomFields() []*GetTestResultListResponseBodyTestResultCustomFields {
	return s.CustomFields
}

func (s *GetTestResultListResponseBodyTestResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetTestResultListResponseBodyTestResult) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *GetTestResultListResponseBodyTestResult) GetSubject() *string {
	return s.Subject
}

func (s *GetTestResultListResponseBodyTestResult) GetTestResultExecutor() *GetTestResultListResponseBodyTestResultTestResultExecutor {
	return s.TestResultExecutor
}

func (s *GetTestResultListResponseBodyTestResult) GetTestResultGmtCreate() *int64 {
	return s.TestResultGmtCreate
}

func (s *GetTestResultListResponseBodyTestResult) GetTestResultIdentifier() *string {
	return s.TestResultIdentifier
}

func (s *GetTestResultListResponseBodyTestResult) GetTestResultStatus() *string {
	return s.TestResultStatus
}

func (s *GetTestResultListResponseBodyTestResult) GetTestcaseIdentifier() *string {
	return s.TestcaseIdentifier
}

func (s *GetTestResultListResponseBodyTestResult) SetAssignedTo(v *GetTestResultListResponseBodyTestResultAssignedTo) *GetTestResultListResponseBodyTestResult {
	s.AssignedTo = v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetBugCount(v int64) *GetTestResultListResponseBodyTestResult {
	s.BugCount = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetCategoryIdentifier(v string) *GetTestResultListResponseBodyTestResult {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetCustomFields(v []*GetTestResultListResponseBodyTestResultCustomFields) *GetTestResultListResponseBodyTestResult {
	s.CustomFields = v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetGmtCreate(v int64) *GetTestResultListResponseBodyTestResult {
	s.GmtCreate = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetSpaceIdentifier(v string) *GetTestResultListResponseBodyTestResult {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetSubject(v string) *GetTestResultListResponseBodyTestResult {
	s.Subject = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetTestResultExecutor(v *GetTestResultListResponseBodyTestResultTestResultExecutor) *GetTestResultListResponseBodyTestResult {
	s.TestResultExecutor = v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetTestResultGmtCreate(v int64) *GetTestResultListResponseBodyTestResult {
	s.TestResultGmtCreate = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetTestResultIdentifier(v string) *GetTestResultListResponseBodyTestResult {
	s.TestResultIdentifier = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetTestResultStatus(v string) *GetTestResultListResponseBodyTestResult {
	s.TestResultStatus = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) SetTestcaseIdentifier(v string) *GetTestResultListResponseBodyTestResult {
	s.TestcaseIdentifier = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResult) Validate() error {
	if s.AssignedTo != nil {
		if err := s.AssignedTo.Validate(); err != nil {
			return err
		}
	}
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TestResultExecutor != nil {
		if err := s.TestResultExecutor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTestResultListResponseBodyTestResultAssignedTo struct {
	AssignedToIdenttifier *string `json:"assignedToIdenttifier,omitempty" xml:"assignedToIdenttifier,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetTestResultListResponseBodyTestResultAssignedTo) String() string {
	return dara.Prettify(s)
}

func (s GetTestResultListResponseBodyTestResultAssignedTo) GoString() string {
	return s.String()
}

func (s *GetTestResultListResponseBodyTestResultAssignedTo) GetAssignedToIdenttifier() *string {
	return s.AssignedToIdenttifier
}

func (s *GetTestResultListResponseBodyTestResultAssignedTo) GetName() *string {
	return s.Name
}

func (s *GetTestResultListResponseBodyTestResultAssignedTo) SetAssignedToIdenttifier(v string) *GetTestResultListResponseBodyTestResultAssignedTo {
	s.AssignedToIdenttifier = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultAssignedTo) SetName(v string) *GetTestResultListResponseBodyTestResultAssignedTo {
	s.Name = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultAssignedTo) Validate() error {
	return dara.Validate(s)
}

type GetTestResultListResponseBodyTestResultCustomFields struct {
	FieldClassName  *string `json:"fieldClassName,omitempty" xml:"fieldClassName,omitempty"`
	FieldFormat     *string `json:"fieldFormat,omitempty" xml:"fieldFormat,omitempty"`
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	Value           *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetTestResultListResponseBodyTestResultCustomFields) String() string {
	return dara.Prettify(s)
}

func (s GetTestResultListResponseBodyTestResultCustomFields) GoString() string {
	return s.String()
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) GetFieldClassName() *string {
	return s.FieldClassName
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) GetFieldFormat() *string {
	return s.FieldFormat
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) GetValue() *string {
	return s.Value
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) SetFieldClassName(v string) *GetTestResultListResponseBodyTestResultCustomFields {
	s.FieldClassName = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) SetFieldFormat(v string) *GetTestResultListResponseBodyTestResultCustomFields {
	s.FieldFormat = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) SetFieldIdentifier(v string) *GetTestResultListResponseBodyTestResultCustomFields {
	s.FieldIdentifier = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) SetValue(v string) *GetTestResultListResponseBodyTestResultCustomFields {
	s.Value = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultCustomFields) Validate() error {
	return dara.Validate(s)
}

type GetTestResultListResponseBodyTestResultTestResultExecutor struct {
	ExecutorIdentifier *string `json:"executorIdentifier,omitempty" xml:"executorIdentifier,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetTestResultListResponseBodyTestResultTestResultExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetTestResultListResponseBodyTestResultTestResultExecutor) GoString() string {
	return s.String()
}

func (s *GetTestResultListResponseBodyTestResultTestResultExecutor) GetExecutorIdentifier() *string {
	return s.ExecutorIdentifier
}

func (s *GetTestResultListResponseBodyTestResultTestResultExecutor) GetName() *string {
	return s.Name
}

func (s *GetTestResultListResponseBodyTestResultTestResultExecutor) SetExecutorIdentifier(v string) *GetTestResultListResponseBodyTestResultTestResultExecutor {
	s.ExecutorIdentifier = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultTestResultExecutor) SetName(v string) *GetTestResultListResponseBodyTestResultTestResultExecutor {
	s.Name = &v
	return s
}

func (s *GetTestResultListResponseBodyTestResultTestResultExecutor) Validate() error {
	return dara.Validate(s)
}
