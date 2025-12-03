// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectWorkitemTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProjectWorkitemTypesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListProjectWorkitemTypesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListProjectWorkitemTypesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectWorkitemTypesResponseBody
	GetSuccess() *bool
	SetWorkitemTypes(v []*ListProjectWorkitemTypesResponseBodyWorkitemTypes) *ListProjectWorkitemTypesResponseBody
	GetWorkitemTypes() []*ListProjectWorkitemTypesResponseBodyWorkitemTypes
}

type ListProjectWorkitemTypesResponseBody struct {
	// example:
	//
	// 例：Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// errormessage
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true或者false
	Success       *bool                                                `json:"success,omitempty" xml:"success,omitempty"`
	WorkitemTypes []*ListProjectWorkitemTypesResponseBodyWorkitemTypes `json:"workitemTypes,omitempty" xml:"workitemTypes,omitempty" type:"Repeated"`
}

func (s ListProjectWorkitemTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectWorkitemTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProjectWorkitemTypesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListProjectWorkitemTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectWorkitemTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectWorkitemTypesResponseBody) GetWorkitemTypes() []*ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	return s.WorkitemTypes
}

func (s *ListProjectWorkitemTypesResponseBody) SetErrorCode(v string) *ListProjectWorkitemTypesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetErrorMessage(v string) *ListProjectWorkitemTypesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetRequestId(v string) *ListProjectWorkitemTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetSuccess(v bool) *ListProjectWorkitemTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) SetWorkitemTypes(v []*ListProjectWorkitemTypesResponseBodyWorkitemTypes) *ListProjectWorkitemTypesResponseBody {
	s.WorkitemTypes = v
	return s
}

func (s *ListProjectWorkitemTypesResponseBody) Validate() error {
	if s.WorkitemTypes != nil {
		for _, item := range s.WorkitemTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectWorkitemTypesResponseBodyWorkitemTypes struct {
	// example:
	//
	// 用户阿里云pk，例如19xxxx31947xxxx
	AddUser *string `json:"addUser,omitempty" xml:"addUser,omitempty"`
	// example:
	//
	// Req
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// example:
	//
	// 用户阿里云pk，例如19xxxx31947xxxx
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// true或者false
	DefaultType *bool `json:"defaultType,omitempty" xml:"defaultType,omitempty"`
	// example:
	//
	// 该类型的具体信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true或者false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 1641870287000
	GmtAdd *int64 `json:"gmtAdd,omitempty" xml:"gmtAdd,omitempty"`
	// example:
	//
	// 1620455467000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// dfexxxxxf4fee18xxxxx36
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 例：业务类需求
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 例：Business Requirement
	NameEn *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	// example:
	//
	// true
	SystemDefault *bool `json:"systemDefault,omitempty" xml:"systemDefault,omitempty"`
}

func (s ListProjectWorkitemTypesResponseBodyWorkitemTypes) String() string {
	return dara.Prettify(s)
}

func (s ListProjectWorkitemTypesResponseBodyWorkitemTypes) GoString() string {
	return s.String()
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetAddUser() *string {
	return s.AddUser
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetCreator() *string {
	return s.Creator
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetDefaultType() *bool {
	return s.DefaultType
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetDescription() *string {
	return s.Description
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetEnable() *bool {
	return s.Enable
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetGmtAdd() *int64 {
	return s.GmtAdd
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetName() *string {
	return s.Name
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetNameEn() *string {
	return s.NameEn
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) GetSystemDefault() *bool {
	return s.SystemDefault
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetAddUser(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.AddUser = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetCategoryIdentifier(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.CategoryIdentifier = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetCreator(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Creator = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetDefaultType(v bool) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.DefaultType = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetDescription(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Description = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetEnable(v bool) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Enable = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetGmtAdd(v int64) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.GmtAdd = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetGmtCreate(v int64) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetIdentifier(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Identifier = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetName(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.Name = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetNameEn(v string) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.NameEn = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) SetSystemDefault(v bool) *ListProjectWorkitemTypesResponseBodyWorkitemTypes {
	s.SystemDefault = &v
	return s
}

func (s *ListProjectWorkitemTypesResponseBodyWorkitemTypes) Validate() error {
	return dara.Validate(s)
}
