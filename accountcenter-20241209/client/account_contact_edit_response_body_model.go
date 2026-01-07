// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactEditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AccountContactEditResponseBody
	GetCode() *string
	SetData(v *AccountContactEditResponseBodyData) *AccountContactEditResponseBody
	GetData() *AccountContactEditResponseBodyData
	SetMessage(v string) *AccountContactEditResponseBody
	GetMessage() *string
	SetRequestId(v string) *AccountContactEditResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AccountContactEditResponseBody
	GetSuccess() *bool
}

type AccountContactEditResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AccountContactEditResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E4E192DF-798B-58AE-B8BF-EBC15E2E85F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AccountContactEditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccountContactEditResponseBody) GoString() string {
	return s.String()
}

func (s *AccountContactEditResponseBody) GetCode() *string {
	return s.Code
}

func (s *AccountContactEditResponseBody) GetData() *AccountContactEditResponseBodyData {
	return s.Data
}

func (s *AccountContactEditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AccountContactEditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccountContactEditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccountContactEditResponseBody) SetCode(v string) *AccountContactEditResponseBody {
	s.Code = &v
	return s
}

func (s *AccountContactEditResponseBody) SetData(v *AccountContactEditResponseBodyData) *AccountContactEditResponseBody {
	s.Data = v
	return s
}

func (s *AccountContactEditResponseBody) SetMessage(v string) *AccountContactEditResponseBody {
	s.Message = &v
	return s
}

func (s *AccountContactEditResponseBody) SetRequestId(v string) *AccountContactEditResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountContactEditResponseBody) SetSuccess(v bool) *AccountContactEditResponseBody {
	s.Success = &v
	return s
}

func (s *AccountContactEditResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AccountContactEditResponseBodyData struct {
	// example:
	//
	// xxx
	ContactId *int64                                         `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ErrorList []*AccountContactEditResponseBodyDataErrorList `json:"ErrorList,omitempty" xml:"ErrorList,omitempty" type:"Repeated"`
	// example:
	//
	// {\\"TeamDesktops\\": [], \\"PrivateDesktops\\": [], \\"UnallocatedTeamDesktops\\": []}
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AccountContactEditResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AccountContactEditResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccountContactEditResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *AccountContactEditResponseBodyData) GetErrorList() []*AccountContactEditResponseBodyDataErrorList {
	return s.ErrorList
}

func (s *AccountContactEditResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *AccountContactEditResponseBodyData) SetContactId(v int64) *AccountContactEditResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *AccountContactEditResponseBodyData) SetErrorList(v []*AccountContactEditResponseBodyDataErrorList) *AccountContactEditResponseBodyData {
	s.ErrorList = v
	return s
}

func (s *AccountContactEditResponseBodyData) SetResult(v bool) *AccountContactEditResponseBodyData {
	s.Result = &v
	return s
}

func (s *AccountContactEditResponseBodyData) Validate() error {
	if s.ErrorList != nil {
		for _, item := range s.ErrorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AccountContactEditResponseBodyDataErrorList struct {
	// example:
	//
	// MOBILE_CODE_ILLEGAL
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// MOBILE_CODE_ILLEGAL
	ErrorDesc *string `json:"ErrorDesc,omitempty" xml:"ErrorDesc,omitempty"`
	// example:
	//
	// MOBILE_VERIFY_CODE
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
}

func (s AccountContactEditResponseBodyDataErrorList) String() string {
	return dara.Prettify(s)
}

func (s AccountContactEditResponseBodyDataErrorList) GoString() string {
	return s.String()
}

func (s *AccountContactEditResponseBodyDataErrorList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AccountContactEditResponseBodyDataErrorList) GetErrorDesc() *string {
	return s.ErrorDesc
}

func (s *AccountContactEditResponseBodyDataErrorList) GetItem() *string {
	return s.Item
}

func (s *AccountContactEditResponseBodyDataErrorList) SetErrorCode(v string) *AccountContactEditResponseBodyDataErrorList {
	s.ErrorCode = &v
	return s
}

func (s *AccountContactEditResponseBodyDataErrorList) SetErrorDesc(v string) *AccountContactEditResponseBodyDataErrorList {
	s.ErrorDesc = &v
	return s
}

func (s *AccountContactEditResponseBodyDataErrorList) SetItem(v string) *AccountContactEditResponseBodyDataErrorList {
	s.Item = &v
	return s
}

func (s *AccountContactEditResponseBodyDataErrorList) Validate() error {
	return dara.Validate(s)
}
