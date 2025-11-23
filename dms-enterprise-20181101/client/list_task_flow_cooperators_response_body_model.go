// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowCooperatorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCooperatorList(v *ListTaskFlowCooperatorsResponseBodyCooperatorList) *ListTaskFlowCooperatorsResponseBody
	GetCooperatorList() *ListTaskFlowCooperatorsResponseBodyCooperatorList
	SetErrorCode(v string) *ListTaskFlowCooperatorsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTaskFlowCooperatorsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTaskFlowCooperatorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskFlowCooperatorsResponseBody
	GetSuccess() *bool
}

type ListTaskFlowCooperatorsResponseBody struct {
	// The users that are involved in the task flow.
	CooperatorList *ListTaskFlowCooperatorsResponseBodyCooperatorList `json:"CooperatorList,omitempty" xml:"CooperatorList,omitempty" type:"Struct"`
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
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 54C99C70-2DFF-5A8C-A252-EBAA1EB668EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskFlowCooperatorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowCooperatorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskFlowCooperatorsResponseBody) GetCooperatorList() *ListTaskFlowCooperatorsResponseBodyCooperatorList {
	return s.CooperatorList
}

func (s *ListTaskFlowCooperatorsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTaskFlowCooperatorsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTaskFlowCooperatorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskFlowCooperatorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskFlowCooperatorsResponseBody) SetCooperatorList(v *ListTaskFlowCooperatorsResponseBodyCooperatorList) *ListTaskFlowCooperatorsResponseBody {
	s.CooperatorList = v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBody) SetErrorCode(v string) *ListTaskFlowCooperatorsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBody) SetErrorMessage(v string) *ListTaskFlowCooperatorsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBody) SetRequestId(v string) *ListTaskFlowCooperatorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBody) SetSuccess(v bool) *ListTaskFlowCooperatorsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBody) Validate() error {
	if s.CooperatorList != nil {
		if err := s.CooperatorList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskFlowCooperatorsResponseBodyCooperatorList struct {
	Cooperator []*ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator `json:"Cooperator,omitempty" xml:"Cooperator,omitempty" type:"Repeated"`
}

func (s ListTaskFlowCooperatorsResponseBodyCooperatorList) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowCooperatorsResponseBodyCooperatorList) GoString() string {
	return s.String()
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorList) GetCooperator() []*ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator {
	return s.Cooperator
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorList) SetCooperator(v []*ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) *ListTaskFlowCooperatorsResponseBodyCooperatorList {
	s.Cooperator = v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorList) Validate() error {
	if s.Cooperator != nil {
		for _, item := range s.Cooperator {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator struct {
	// The email address of the user.
	//
	// example:
	//
	// test@XX.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username.
	//
	// example:
	//
	// name
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The alias of the user.
	//
	// example:
	//
	// name
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// userId.
	//
	// example:
	//
	// 123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) GoString() string {
	return s.String()
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) GetEmail() *string {
	return s.Email
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) GetLoginName() *string {
	return s.LoginName
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) GetNickName() *string {
	return s.NickName
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) GetUserId() *string {
	return s.UserId
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) SetEmail(v string) *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator {
	s.Email = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) SetLoginName(v string) *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator {
	s.LoginName = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) SetNickName(v string) *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator {
	s.NickName = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) SetUserId(v string) *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator {
	s.UserId = &v
	return s
}

func (s *ListTaskFlowCooperatorsResponseBodyCooperatorListCooperator) Validate() error {
	return dara.Validate(s)
}
