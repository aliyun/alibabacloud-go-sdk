// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTestCaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTestcase(v *UpdateTestCaseResponseBodyTestcase) *UpdateTestCaseResponseBody
	GetTestcase() *UpdateTestCaseResponseBodyTestcase
	SetErrorCode(v string) *UpdateTestCaseResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *UpdateTestCaseResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *UpdateTestCaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTestCaseResponseBody
	GetSuccess() *bool
}

type UpdateTestCaseResponseBody struct {
	Testcase *UpdateTestCaseResponseBodyTestcase `json:"Testcase,omitempty" xml:"Testcase,omitempty" type:"Struct"`
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
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTestCaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBody) GetTestcase() *UpdateTestCaseResponseBodyTestcase {
	return s.Testcase
}

func (s *UpdateTestCaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateTestCaseResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateTestCaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTestCaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTestCaseResponseBody) SetTestcase(v *UpdateTestCaseResponseBodyTestcase) *UpdateTestCaseResponseBody {
	s.Testcase = v
	return s
}

func (s *UpdateTestCaseResponseBody) SetErrorCode(v string) *UpdateTestCaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTestCaseResponseBody) SetErrorMsg(v string) *UpdateTestCaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTestCaseResponseBody) SetRequestId(v string) *UpdateTestCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTestCaseResponseBody) SetSuccess(v bool) *UpdateTestCaseResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTestCaseResponseBody) Validate() error {
	if s.Testcase != nil {
		if err := s.Testcase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTestCaseResponseBodyTestcase struct {
	AssignedTo *UpdateTestCaseResponseBodyTestcaseAssignedTo `json:"assignedTo,omitempty" xml:"assignedTo,omitempty" type:"Struct"`
	// example:
	//
	// TestCase
	CategoryIdentifier *string                                       `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	Creator            *UpdateTestCaseResponseBodyTestcaseCreator    `json:"creator,omitempty" xml:"creator,omitempty" type:"Struct"`
	DetailInfo         *UpdateTestCaseResponseBodyTestcaseDetailInfo `json:"detailInfo,omitempty" xml:"detailInfo,omitempty" type:"Struct"`
	Directory          *UpdateTestCaseResponseBodyTestcaseDirectory  `json:"directory,omitempty" xml:"directory,omitempty" type:"Struct"`
	// example:
	//
	// c7f7033b021ead52cc42721382
	Identifier *string                                     `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Modifier   *UpdateTestCaseResponseBodyTestcaseModifier `json:"modifier,omitempty" xml:"modifier,omitempty" type:"Struct"`
	// example:
	//
	// 1e7d7a412b91a2144ec4aa8411
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// TestRepo
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	Subject   *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcase) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcase) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcase) GetAssignedTo() *UpdateTestCaseResponseBodyTestcaseAssignedTo {
	return s.AssignedTo
}

func (s *UpdateTestCaseResponseBodyTestcase) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcase) GetCreator() *UpdateTestCaseResponseBodyTestcaseCreator {
	return s.Creator
}

func (s *UpdateTestCaseResponseBodyTestcase) GetDetailInfo() *UpdateTestCaseResponseBodyTestcaseDetailInfo {
	return s.DetailInfo
}

func (s *UpdateTestCaseResponseBodyTestcase) GetDirectory() *UpdateTestCaseResponseBodyTestcaseDirectory {
	return s.Directory
}

func (s *UpdateTestCaseResponseBodyTestcase) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateTestCaseResponseBodyTestcase) GetModifier() *UpdateTestCaseResponseBodyTestcaseModifier {
	return s.Modifier
}

func (s *UpdateTestCaseResponseBodyTestcase) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcase) GetSpaceType() *string {
	return s.SpaceType
}

func (s *UpdateTestCaseResponseBodyTestcase) GetSubject() *string {
	return s.Subject
}

func (s *UpdateTestCaseResponseBodyTestcase) SetAssignedTo(v *UpdateTestCaseResponseBodyTestcaseAssignedTo) *UpdateTestCaseResponseBodyTestcase {
	s.AssignedTo = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetCategoryIdentifier(v string) *UpdateTestCaseResponseBodyTestcase {
	s.CategoryIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetCreator(v *UpdateTestCaseResponseBodyTestcaseCreator) *UpdateTestCaseResponseBodyTestcase {
	s.Creator = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetDetailInfo(v *UpdateTestCaseResponseBodyTestcaseDetailInfo) *UpdateTestCaseResponseBodyTestcase {
	s.DetailInfo = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetDirectory(v *UpdateTestCaseResponseBodyTestcaseDirectory) *UpdateTestCaseResponseBodyTestcase {
	s.Directory = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetIdentifier(v string) *UpdateTestCaseResponseBodyTestcase {
	s.Identifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetModifier(v *UpdateTestCaseResponseBodyTestcaseModifier) *UpdateTestCaseResponseBodyTestcase {
	s.Modifier = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetSpaceIdentifier(v string) *UpdateTestCaseResponseBodyTestcase {
	s.SpaceIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetSpaceType(v string) *UpdateTestCaseResponseBodyTestcase {
	s.SpaceType = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) SetSubject(v string) *UpdateTestCaseResponseBodyTestcase {
	s.Subject = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcase) Validate() error {
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

type UpdateTestCaseResponseBodyTestcaseAssignedTo struct {
	// example:
	//
	// 1316xxxxxx8624xxx
	AssignIdentifier *string `json:"assignIdentifier,omitempty" xml:"assignIdentifier,omitempty"`
	Name             *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// xxxxxxxewdds12xxx
	TbRoleId *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcaseAssignedTo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseAssignedTo) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseAssignedTo) GetAssignIdentifier() *string {
	return s.AssignIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseAssignedTo) GetName() *string {
	return s.Name
}

func (s *UpdateTestCaseResponseBodyTestcaseAssignedTo) GetTbRoleId() *string {
	return s.TbRoleId
}

func (s *UpdateTestCaseResponseBodyTestcaseAssignedTo) SetAssignIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseAssignedTo {
	s.AssignIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseAssignedTo) SetName(v string) *UpdateTestCaseResponseBodyTestcaseAssignedTo {
	s.Name = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseAssignedTo) SetTbRoleId(v string) *UpdateTestCaseResponseBodyTestcaseAssignedTo {
	s.TbRoleId = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseAssignedTo) Validate() error {
	return dara.Validate(s)
}

type UpdateTestCaseResponseBodyTestcaseCreator struct {
	// example:
	//
	// 1316xxxxxx8624xxx
	CreateIdentifier *string `json:"createIdentifier,omitempty" xml:"createIdentifier,omitempty"`
	// example:
	//
	// xxxxxxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcaseCreator) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseCreator) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseCreator) GetCreateIdentifier() *string {
	return s.CreateIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseCreator) GetName() *string {
	return s.Name
}

func (s *UpdateTestCaseResponseBodyTestcaseCreator) SetCreateIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseCreator {
	s.CreateIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseCreator) SetName(v string) *UpdateTestCaseResponseBodyTestcaseCreator {
	s.Name = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseCreator) Validate() error {
	return dara.Validate(s)
}

type UpdateTestCaseResponseBodyTestcaseDetailInfo struct {
	ExpectedResult *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult `json:"expectedResult,omitempty" xml:"expectedResult,omitempty" type:"Struct"`
	Precondition   *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition   `json:"precondition,omitempty" xml:"precondition,omitempty" type:"Struct"`
	StepContent    *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent    `json:"stepContent,omitempty" xml:"stepContent,omitempty" type:"Struct"`
	// example:
	//
	// TEXT/TABLE
	StepType *string `json:"stepType,omitempty" xml:"stepType,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfo) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) GetExpectedResult() *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	return s.ExpectedResult
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) GetPrecondition() *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	return s.Precondition
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) GetStepContent() *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	return s.StepContent
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) GetStepType() *string {
	return s.StepType
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) SetExpectedResult(v *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) *UpdateTestCaseResponseBodyTestcaseDetailInfo {
	s.ExpectedResult = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) SetPrecondition(v *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) *UpdateTestCaseResponseBodyTestcaseDetailInfo {
	s.Precondition = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) SetStepContent(v *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) *UpdateTestCaseResponseBodyTestcaseDetailInfo {
	s.StepContent = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) SetStepType(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfo {
	s.StepType = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfo) Validate() error {
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

type UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult struct {
	ExpectContent *string `json:"expectContent,omitempty" xml:"expectContent,omitempty"`
	// example:
	//
	// MARKDOWN
	ExpectContentType *string `json:"expectContentType,omitempty" xml:"expectContentType,omitempty"`
	// example:
	//
	// 685340d13127b01185335bd360
	ExpectIdentifier *string `json:"expectIdentifier,omitempty" xml:"expectIdentifier,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GetExpectContent() *string {
	return s.ExpectContent
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GetExpectContentType() *string {
	return s.ExpectContentType
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) GetExpectIdentifier() *string {
	return s.ExpectIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) SetExpectContent(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	s.ExpectContent = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) SetExpectContentType(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	s.ExpectContentType = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) SetExpectIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult {
	s.ExpectIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoExpectedResult) Validate() error {
	return dara.Validate(s)
}

type UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition struct {
	PreContent *string `json:"preContent,omitempty" xml:"preContent,omitempty"`
	// example:
	//
	// MARKDOWN
	PreContentType *string `json:"preContentType,omitempty" xml:"preContentType,omitempty"`
	// example:
	//
	// 3354596c7b3004480b635acf95
	PreIdentifier *string `json:"preIdentifier,omitempty" xml:"preIdentifier,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GetPreContent() *string {
	return s.PreContent
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GetPreContentType() *string {
	return s.PreContentType
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) GetPreIdentifier() *string {
	return s.PreIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) SetPreContent(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	s.PreContent = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) SetPreContentType(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	s.PreContentType = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) SetPreIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition {
	s.PreIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoPrecondition) Validate() error {
	return dara.Validate(s)
}

type UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent struct {
	StepContent *string `json:"stepContent,omitempty" xml:"stepContent,omitempty"`
	// example:
	//
	// RICHTEXT
	StepContentType *string `json:"stepContentType,omitempty" xml:"stepContentType,omitempty"`
	// example:
	//
	// ad504e6cdcd2165b28eb1e1b9f
	StepIdentifier *string `json:"stepIdentifier,omitempty" xml:"stepIdentifier,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) GetStepContent() *string {
	return s.StepContent
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) GetStepContentType() *string {
	return s.StepContentType
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) GetStepIdentifier() *string {
	return s.StepIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) SetStepContent(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	s.StepContent = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) SetStepContentType(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	s.StepContentType = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) SetStepIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent {
	s.StepIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDetailInfoStepContent) Validate() error {
	return dara.Validate(s)
}

type UpdateTestCaseResponseBodyTestcaseDirectory struct {
	// example:
	//
	// 2973f597c14c6f533fffdcd05c
	ChildIdentifier *string `json:"childIdentifier,omitempty" xml:"childIdentifier,omitempty"`
	// example:
	//
	// e27b8eace6501ce51cf5d56784
	DirectoryIdentifier *string   `json:"directoryIdentifier,omitempty" xml:"directoryIdentifier,omitempty"`
	Name                *string   `json:"name,omitempty" xml:"name,omitempty"`
	PathName            []*string `json:"pathName,omitempty" xml:"pathName,omitempty" type:"Repeated"`
}

func (s UpdateTestCaseResponseBodyTestcaseDirectory) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseDirectory) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) GetChildIdentifier() *string {
	return s.ChildIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) GetDirectoryIdentifier() *string {
	return s.DirectoryIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) GetName() *string {
	return s.Name
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) GetPathName() []*string {
	return s.PathName
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) SetChildIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseDirectory {
	s.ChildIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) SetDirectoryIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseDirectory {
	s.DirectoryIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) SetName(v string) *UpdateTestCaseResponseBodyTestcaseDirectory {
	s.Name = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) SetPathName(v []*string) *UpdateTestCaseResponseBodyTestcaseDirectory {
	s.PathName = v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseDirectory) Validate() error {
	return dara.Validate(s)
}

type UpdateTestCaseResponseBodyTestcaseModifier struct {
	// example:
	//
	// 1316xxxxxx8624xxx
	ModifyIdentifier *string `json:"modifyIdentifier,omitempty" xml:"modifyIdentifier,omitempty"`
	// example:
	//
	// xxxxxxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateTestCaseResponseBodyTestcaseModifier) String() string {
	return dara.Prettify(s)
}

func (s UpdateTestCaseResponseBodyTestcaseModifier) GoString() string {
	return s.String()
}

func (s *UpdateTestCaseResponseBodyTestcaseModifier) GetModifyIdentifier() *string {
	return s.ModifyIdentifier
}

func (s *UpdateTestCaseResponseBodyTestcaseModifier) GetName() *string {
	return s.Name
}

func (s *UpdateTestCaseResponseBodyTestcaseModifier) SetModifyIdentifier(v string) *UpdateTestCaseResponseBodyTestcaseModifier {
	s.ModifyIdentifier = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseModifier) SetName(v string) *UpdateTestCaseResponseBodyTestcaseModifier {
	s.Name = &v
	return s
}

func (s *UpdateTestCaseResponseBodyTestcaseModifier) Validate() error {
	return dara.Validate(s)
}
