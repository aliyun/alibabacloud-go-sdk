// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTestcaseListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTestcaseListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetTestcaseListResponseBody
	GetErrorMsg() *string
	SetMaxResults(v int64) *GetTestcaseListResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *GetTestcaseListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetTestcaseListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTestcaseListResponseBody
	GetSuccess() *bool
	SetTestcase(v []*GetTestcaseListResponseBodyTestcase) *GetTestcaseListResponseBody
	GetTestcase() []*GetTestcaseListResponseBodyTestcase
	SetTotalCount(v int64) *GetTestcaseListResponseBody
	GetTotalCount() *int64
}

type GetTestcaseListResponseBody struct {
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
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success  *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
	Testcase []*GetTestcaseListResponseBodyTestcase `json:"testcase,omitempty" xml:"testcase,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetTestcaseListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTestcaseListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTestcaseListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetTestcaseListResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetTestcaseListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetTestcaseListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTestcaseListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTestcaseListResponseBody) GetTestcase() []*GetTestcaseListResponseBodyTestcase {
	return s.Testcase
}

func (s *GetTestcaseListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTestcaseListResponseBody) SetErrorCode(v string) *GetTestcaseListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTestcaseListResponseBody) SetErrorMsg(v string) *GetTestcaseListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTestcaseListResponseBody) SetMaxResults(v int64) *GetTestcaseListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetTestcaseListResponseBody) SetNextToken(v string) *GetTestcaseListResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetTestcaseListResponseBody) SetRequestId(v string) *GetTestcaseListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTestcaseListResponseBody) SetSuccess(v bool) *GetTestcaseListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTestcaseListResponseBody) SetTestcase(v []*GetTestcaseListResponseBodyTestcase) *GetTestcaseListResponseBody {
	s.Testcase = v
	return s
}

func (s *GetTestcaseListResponseBody) SetTotalCount(v int64) *GetTestcaseListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTestcaseListResponseBody) Validate() error {
	if s.Testcase != nil {
		for _, item := range s.Testcase {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTestcaseListResponseBodyTestcase struct {
	AssignedTo *GetTestcaseListResponseBodyTestcaseAssignedTo `json:"assignedTo,omitempty" xml:"assignedTo,omitempty" type:"Struct"`
	// example:
	//
	// Req
	CategoryIdentifier *string                                            `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	CustomFields       []*GetTestcaseListResponseBodyTestcaseCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string                                   `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Space      *GetTestcaseListResponseBodyTestcaseSpace `json:"space,omitempty" xml:"space,omitempty" type:"Struct"`
	// example:
	//
	// 测试工作项
	Subject *string                                    `json:"subject,omitempty" xml:"subject,omitempty"`
	Tags    []*GetTestcaseListResponseBodyTestcaseTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s GetTestcaseListResponseBodyTestcase) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListResponseBodyTestcase) GoString() string {
	return s.String()
}

func (s *GetTestcaseListResponseBodyTestcase) GetAssignedTo() *GetTestcaseListResponseBodyTestcaseAssignedTo {
	return s.AssignedTo
}

func (s *GetTestcaseListResponseBodyTestcase) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *GetTestcaseListResponseBodyTestcase) GetCustomFields() []*GetTestcaseListResponseBodyTestcaseCustomFields {
	return s.CustomFields
}

func (s *GetTestcaseListResponseBodyTestcase) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetTestcaseListResponseBodyTestcase) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetTestcaseListResponseBodyTestcase) GetSpace() *GetTestcaseListResponseBodyTestcaseSpace {
	return s.Space
}

func (s *GetTestcaseListResponseBodyTestcase) GetSubject() *string {
	return s.Subject
}

func (s *GetTestcaseListResponseBodyTestcase) GetTags() []*GetTestcaseListResponseBodyTestcaseTags {
	return s.Tags
}

func (s *GetTestcaseListResponseBodyTestcase) SetAssignedTo(v *GetTestcaseListResponseBodyTestcaseAssignedTo) *GetTestcaseListResponseBodyTestcase {
	s.AssignedTo = v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) SetCategoryIdentifier(v string) *GetTestcaseListResponseBodyTestcase {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) SetCustomFields(v []*GetTestcaseListResponseBodyTestcaseCustomFields) *GetTestcaseListResponseBodyTestcase {
	s.CustomFields = v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) SetGmtCreate(v int64) *GetTestcaseListResponseBodyTestcase {
	s.GmtCreate = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) SetIdentifier(v string) *GetTestcaseListResponseBodyTestcase {
	s.Identifier = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) SetSpace(v *GetTestcaseListResponseBodyTestcaseSpace) *GetTestcaseListResponseBodyTestcase {
	s.Space = v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) SetSubject(v string) *GetTestcaseListResponseBodyTestcase {
	s.Subject = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) SetTags(v []*GetTestcaseListResponseBodyTestcaseTags) *GetTestcaseListResponseBodyTestcase {
	s.Tags = v
	return s
}

func (s *GetTestcaseListResponseBodyTestcase) Validate() error {
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
	if s.Space != nil {
		if err := s.Space.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTestcaseListResponseBodyTestcaseAssignedTo struct {
	// example:
	//
	// 12xxxxx456
	AssignedToIdenttifier *string `json:"assignedToIdenttifier,omitempty" xml:"assignedToIdenttifier,omitempty"`
	Name                  *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetTestcaseListResponseBodyTestcaseAssignedTo) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListResponseBodyTestcaseAssignedTo) GoString() string {
	return s.String()
}

func (s *GetTestcaseListResponseBodyTestcaseAssignedTo) GetAssignedToIdenttifier() *string {
	return s.AssignedToIdenttifier
}

func (s *GetTestcaseListResponseBodyTestcaseAssignedTo) GetName() *string {
	return s.Name
}

func (s *GetTestcaseListResponseBodyTestcaseAssignedTo) SetAssignedToIdenttifier(v string) *GetTestcaseListResponseBodyTestcaseAssignedTo {
	s.AssignedToIdenttifier = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseAssignedTo) SetName(v string) *GetTestcaseListResponseBodyTestcaseAssignedTo {
	s.Name = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseAssignedTo) Validate() error {
	return dara.Validate(s)
}

type GetTestcaseListResponseBodyTestcaseCustomFields struct {
	// example:
	//
	// User
	FieldClassName *string `json:"fieldClassName,omitempty" xml:"fieldClassName,omitempty"`
	// example:
	//
	// Input
	FieldFormat *string `json:"fieldFormat,omitempty" xml:"fieldFormat,omitempty"`
	// example:
	//
	// 85702b33f14bfa82cb458173ba
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// example:
	//
	// d7f112f9d023e2108fa1b0d8
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetTestcaseListResponseBodyTestcaseCustomFields) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListResponseBodyTestcaseCustomFields) GoString() string {
	return s.String()
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) GetFieldClassName() *string {
	return s.FieldClassName
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) GetFieldFormat() *string {
	return s.FieldFormat
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) GetValue() *string {
	return s.Value
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) SetFieldClassName(v string) *GetTestcaseListResponseBodyTestcaseCustomFields {
	s.FieldClassName = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) SetFieldFormat(v string) *GetTestcaseListResponseBodyTestcaseCustomFields {
	s.FieldFormat = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) SetFieldIdentifier(v string) *GetTestcaseListResponseBodyTestcaseCustomFields {
	s.FieldIdentifier = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) SetValue(v string) *GetTestcaseListResponseBodyTestcaseCustomFields {
	s.Value = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseCustomFields) Validate() error {
	return dara.Validate(s)
}

type GetTestcaseListResponseBodyTestcaseSpace struct {
	// example:
	//
	// 22c32972b853cd703dbf0efe4c
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// TestRepo
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTestcaseListResponseBodyTestcaseSpace) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListResponseBodyTestcaseSpace) GoString() string {
	return s.String()
}

func (s *GetTestcaseListResponseBodyTestcaseSpace) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *GetTestcaseListResponseBodyTestcaseSpace) GetType() *string {
	return s.Type
}

func (s *GetTestcaseListResponseBodyTestcaseSpace) SetSpaceIdentifier(v string) *GetTestcaseListResponseBodyTestcaseSpace {
	s.SpaceIdentifier = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseSpace) SetType(v string) *GetTestcaseListResponseBodyTestcaseSpace {
	s.Type = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseSpace) Validate() error {
	return dara.Validate(s)
}

type GetTestcaseListResponseBodyTestcaseTags struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 85702b33f14bxxxxxx58173ba
	TagIdentifier *string `json:"tagIdentifier,omitempty" xml:"tagIdentifier,omitempty"`
}

func (s GetTestcaseListResponseBodyTestcaseTags) String() string {
	return dara.Prettify(s)
}

func (s GetTestcaseListResponseBodyTestcaseTags) GoString() string {
	return s.String()
}

func (s *GetTestcaseListResponseBodyTestcaseTags) GetName() *string {
	return s.Name
}

func (s *GetTestcaseListResponseBodyTestcaseTags) GetTagIdentifier() *string {
	return s.TagIdentifier
}

func (s *GetTestcaseListResponseBodyTestcaseTags) SetName(v string) *GetTestcaseListResponseBodyTestcaseTags {
	s.Name = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseTags) SetTagIdentifier(v string) *GetTestcaseListResponseBodyTestcaseTags {
	s.TagIdentifier = &v
	return s
}

func (s *GetTestcaseListResponseBodyTestcaseTags) Validate() error {
	return dara.Validate(s)
}
