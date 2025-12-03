// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTestCaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateTestCaseResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateTestCaseResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateTestCaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTestCaseResponseBody
	GetSuccess() *bool
	SetTestcase(v *CreateTestCaseResponseBodyTestcase) *CreateTestCaseResponseBody
	GetTestcase() *CreateTestCaseResponseBodyTestcase
}

type CreateTestCaseResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 18E50717-93A4-53BC-A30D-963F742A1CE6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success  *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	Testcase *CreateTestCaseResponseBodyTestcase `json:"testcase,omitempty" xml:"testcase,omitempty" type:"Struct"`
}

func (s CreateTestCaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTestCaseResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateTestCaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTestCaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTestCaseResponseBody) GetTestcase() *CreateTestCaseResponseBodyTestcase {
	return s.Testcase
}

func (s *CreateTestCaseResponseBody) SetErrorCode(v string) *CreateTestCaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTestCaseResponseBody) SetErrorMsg(v string) *CreateTestCaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateTestCaseResponseBody) SetRequestId(v string) *CreateTestCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTestCaseResponseBody) SetSuccess(v bool) *CreateTestCaseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTestCaseResponseBody) SetTestcase(v *CreateTestCaseResponseBodyTestcase) *CreateTestCaseResponseBody {
	s.Testcase = v
	return s
}

func (s *CreateTestCaseResponseBody) Validate() error {
	if s.Testcase != nil {
		if err := s.Testcase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTestCaseResponseBodyTestcase struct {
	AssignedTo *CreateTestCaseResponseBodyTestcaseAssignedTo `json:"assignedTo,omitempty" xml:"assignedTo,omitempty" type:"Struct"`
	// example:
	//
	// TestCase
	CategoryIdentifier *string                                       `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator            *CreateTestCaseResponseBodyTestcaseCreator    `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	DetailInfo         *CreateTestCaseResponseBodyTestcaseDetailInfo `json:"detailInfo,omitempty" xml:"detailInfo,omitempty" type:"Struct"`
	Directory          *CreateTestCaseResponseBodyTestcaseDirectory  `json:"directory,omitempty" xml:"directory,omitempty" type:"Struct"`
	// example:
	//
	// 5a73f81c834d013361d92bdcce
	Identifier *string                                     `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier   *CreateTestCaseResponseBodyTestcaseModifier `json:"modifier,omitempty" xml:"modifier,omitempty" type:"Struct"`
	// example:
	//
	// a18571eba8fe9267cd8375fc06
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// TestRepo
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Subject   *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcase) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcase) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcase) GetAssignedTo() *CreateTestCaseResponseBodyTestcaseAssignedTo {
	return s.AssignedTo
}

func (s *CreateTestCaseResponseBodyTestcase) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *CreateTestCaseResponseBodyTestcase) GetCreator() *CreateTestCaseResponseBodyTestcaseCreator {
	return s.Creator
}

func (s *CreateTestCaseResponseBodyTestcase) GetDetailInfo() *CreateTestCaseResponseBodyTestcaseDetailInfo {
	return s.DetailInfo
}

func (s *CreateTestCaseResponseBodyTestcase) GetDirectory() *CreateTestCaseResponseBodyTestcaseDirectory {
	return s.Directory
}

func (s *CreateTestCaseResponseBodyTestcase) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateTestCaseResponseBodyTestcase) GetModifier() *CreateTestCaseResponseBodyTestcaseModifier {
	return s.Modifier
}

func (s *CreateTestCaseResponseBodyTestcase) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *CreateTestCaseResponseBodyTestcase) GetSpaceType() *string {
	return s.SpaceType
}

func (s *CreateTestCaseResponseBodyTestcase) GetSubject() *string {
	return s.Subject
}

func (s *CreateTestCaseResponseBodyTestcase) SetAssignedTo(v *CreateTestCaseResponseBodyTestcaseAssignedTo) *CreateTestCaseResponseBodyTestcase {
	s.AssignedTo = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetCategoryIdentifier(v string) *CreateTestCaseResponseBodyTestcase {
	s.CategoryIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetCreator(v *CreateTestCaseResponseBodyTestcaseCreator) *CreateTestCaseResponseBodyTestcase {
	s.Creator = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetDetailInfo(v *CreateTestCaseResponseBodyTestcaseDetailInfo) *CreateTestCaseResponseBodyTestcase {
	s.DetailInfo = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetDirectory(v *CreateTestCaseResponseBodyTestcaseDirectory) *CreateTestCaseResponseBodyTestcase {
	s.Directory = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetIdentifier(v string) *CreateTestCaseResponseBodyTestcase {
	s.Identifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetModifier(v *CreateTestCaseResponseBodyTestcaseModifier) *CreateTestCaseResponseBodyTestcase {
	s.Modifier = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetSpaceIdentifier(v string) *CreateTestCaseResponseBodyTestcase {
	s.SpaceIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetSpaceType(v string) *CreateTestCaseResponseBodyTestcase {
	s.SpaceType = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) SetSubject(v string) *CreateTestCaseResponseBodyTestcase {
	s.Subject = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcase) Validate() error {
	if s.AssignedTo != nil {
		if err := s.AssignedTo.Validate(); err != nil {
			return err
		}
	}
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.DetailInfo != nil {
		if err := s.DetailInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTestCaseResponseBodyTestcaseAssignedTo struct {
	// example:
	//
	// 134xxx343xxxxx
	AssignIdentifier *string `json:"assignIdentifier,omitempty" xml:"assignIdentifier,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 3c2253c22xxxxxxxx53a
	TbRoleId *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcaseAssignedTo) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseAssignedTo) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseAssignedTo) GetAssignIdentifier() *string {
	return s.AssignIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseAssignedTo) GetName() *string {
	return s.Name
}

func (s *CreateTestCaseResponseBodyTestcaseAssignedTo) GetTbRoleId() *string {
	return s.TbRoleId
}

func (s *CreateTestCaseResponseBodyTestcaseAssignedTo) SetAssignIdentifier(v string) *CreateTestCaseResponseBodyTestcaseAssignedTo {
	s.AssignIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseAssignedTo) SetName(v string) *CreateTestCaseResponseBodyTestcaseAssignedTo {
	s.Name = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseAssignedTo) SetTbRoleId(v string) *CreateTestCaseResponseBodyTestcaseAssignedTo {
	s.TbRoleId = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseAssignedTo) Validate() error {
	return dara.Validate(s)
}

type CreateTestCaseResponseBodyTestcaseCreator struct {
	// example:
	//
	// 134xxx343xxxxx
	CreateIdentifier *string `json:"createIdentifier,omitempty" xml:"createIdentifier,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcaseCreator) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseCreator) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseCreator) GetCreateIdentifier() *string {
	return s.CreateIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseCreator) GetName() *string {
	return s.Name
}

func (s *CreateTestCaseResponseBodyTestcaseCreator) SetCreateIdentifier(v string) *CreateTestCaseResponseBodyTestcaseCreator {
	s.CreateIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseCreator) SetName(v string) *CreateTestCaseResponseBodyTestcaseCreator {
	s.Name = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseCreator) Validate() error {
	return dara.Validate(s)
}

type CreateTestCaseResponseBodyTestcaseDetailInfo struct {
	ExpectedResult *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult `json:"expectedResult,omitempty" xml:"expectedResult,omitempty" type:"Struct"`
	Precondition   *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition   `json:"precondition,omitempty" xml:"precondition,omitempty" type:"Struct"`
	StepContent    *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent    `json:"stepContent,omitempty" xml:"stepContent,omitempty" type:"Struct"`
	// example:
	//
	// TEXT/TABLE
	StepType *string `json:"stepType,omitempty" xml:"stepType,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfo) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) GetExpectedResult() *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	return s.ExpectedResult
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) GetPrecondition() *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	return s.Precondition
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) GetStepContent() *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	return s.StepContent
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) GetStepType() *string {
	return s.StepType
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) SetExpectedResult(v *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) *CreateTestCaseResponseBodyTestcaseDetailInfo {
	s.ExpectedResult = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) SetPrecondition(v *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) *CreateTestCaseResponseBodyTestcaseDetailInfo {
	s.Precondition = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) SetStepContent(v *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) *CreateTestCaseResponseBodyTestcaseDetailInfo {
	s.StepContent = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) SetStepType(v string) *CreateTestCaseResponseBodyTestcaseDetailInfo {
	s.StepType = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfo) Validate() error {
	if s.ExpectedResult != nil {
		if err := s.ExpectedResult.Validate(); err != nil {
			return err
		}
	}
	if s.Precondition != nil {
		if err := s.Precondition.Validate(); err != nil {
			return err
		}
	}
	if s.StepContent != nil {
		if err := s.StepContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult struct {
	ExpectContent *string `json:"expectContent,omitempty" xml:"expectContent,omitempty"`
	// example:
	//
	// RICHTEXT
	ExpectContentType *string `json:"expectContentType,omitempty" xml:"expectContentType,omitempty"`
	// example:
	//
	// 59253164xxxxxxf2e98dbc7e27
	ExpectIdentifier *string `json:"expectIdentifier,omitempty" xml:"expectIdentifier,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GetExpectContent() *string {
	return s.ExpectContent
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GetExpectContentType() *string {
	return s.ExpectContentType
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GetExpectIdentifier() *string {
	return s.ExpectIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) SetExpectContent(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	s.ExpectContent = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) SetExpectContentType(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	s.ExpectContentType = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) SetExpectIdentifier(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	s.ExpectIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) Validate() error {
	return dara.Validate(s)
}

type CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition struct {
	PreContent *string `json:"preContent,omitempty" xml:"preContent,omitempty"`
	// example:
	//
	// RICHTEXT
	PreContentType *string `json:"preContentType,omitempty" xml:"preContentType,omitempty"`
	// example:
	//
	// 59253164xxxxxxf2e98dbc7e27
	PreIdentifier *string `json:"preIdentifier,omitempty" xml:"preIdentifier,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GetPreContent() *string {
	return s.PreContent
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GetPreContentType() *string {
	return s.PreContentType
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GetPreIdentifier() *string {
	return s.PreIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) SetPreContent(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	s.PreContent = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) SetPreContentType(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	s.PreContentType = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) SetPreIdentifier(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	s.PreIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoPrecondition) Validate() error {
	return dara.Validate(s)
}

type CreateTestCaseResponseBodyTestcaseDetailInfoStepContent struct {
	StepContent *string `json:"stepContent,omitempty" xml:"stepContent,omitempty"`
	// example:
	//
	// RICHTEXT
	StepContentType *string `json:"stepContentType,omitempty" xml:"stepContentType,omitempty"`
	// example:
	//
	// 59253164xxxxxxf2e98dbc7e27
	StepIdentifier *string `json:"stepIdentifier,omitempty" xml:"stepIdentifier,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) GetStepContent() *string {
	return s.StepContent
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) GetStepContentType() *string {
	return s.StepContentType
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) GetStepIdentifier() *string {
	return s.StepIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) SetStepContent(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	s.StepContent = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) SetStepContentType(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	s.StepContentType = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) SetStepIdentifier(v string) *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	s.StepIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDetailInfoStepContent) Validate() error {
	return dara.Validate(s)
}

type CreateTestCaseResponseBodyTestcaseDirectory struct {
	// example:
	//
	// 0bc1150dcxxxxxxxx04c
	ChildIdentifier *string `json:"childIdentifier,omitempty" xml:"childIdentifier,omitempty"`
	// example:
	//
	// 0bc1150dcxxxxxxxx04c
	DirectoryIdentifier *string   `json:"directoryIdentifier,omitempty" xml:"directoryIdentifier,omitempty"`
	Name                *string   `json:"name,omitempty" xml:"name,omitempty"`
	PathName            []*string `json:"pathName,omitempty" xml:"pathName,omitempty" type:"Repeated"`
}

func (s CreateTestCaseResponseBodyTestcaseDirectory) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseDirectory) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) GetChildIdentifier() *string {
	return s.ChildIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) GetDirectoryIdentifier() *string {
	return s.DirectoryIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) GetName() *string {
	return s.Name
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) GetPathName() []*string {
	return s.PathName
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) SetChildIdentifier(v string) *CreateTestCaseResponseBodyTestcaseDirectory {
	s.ChildIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) SetDirectoryIdentifier(v string) *CreateTestCaseResponseBodyTestcaseDirectory {
	s.DirectoryIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) SetName(v string) *CreateTestCaseResponseBodyTestcaseDirectory {
	s.Name = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) SetPathName(v []*string) *CreateTestCaseResponseBodyTestcaseDirectory {
	s.PathName = v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseDirectory) Validate() error {
	return dara.Validate(s)
}

type CreateTestCaseResponseBodyTestcaseModifier struct {
	// example:
	//
	// 134xxx343xxxxx
	ModifyIdentifier *string `json:"modifyIdentifier,omitempty" xml:"modifyIdentifier,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateTestCaseResponseBodyTestcaseModifier) String() string {
	return dara.Prettify(s)
}

func (s CreateTestCaseResponseBodyTestcaseModifier) GoString() string {
	return s.String()
}

func (s *CreateTestCaseResponseBodyTestcaseModifier) GetModifyIdentifier() *string {
	return s.ModifyIdentifier
}

func (s *CreateTestCaseResponseBodyTestcaseModifier) GetName() *string {
	return s.Name
}

func (s *CreateTestCaseResponseBodyTestcaseModifier) SetModifyIdentifier(v string) *CreateTestCaseResponseBodyTestcaseModifier {
	s.ModifyIdentifier = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseModifier) SetName(v string) *CreateTestCaseResponseBodyTestcaseModifier {
	s.Name = &v
	return s
}

func (s *CreateTestCaseResponseBodyTestcaseModifier) Validate() error {
	return dara.Validate(s)
}
