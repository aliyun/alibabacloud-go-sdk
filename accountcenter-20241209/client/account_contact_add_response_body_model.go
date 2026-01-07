// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AccountContactAddResponseBody
	GetCode() *string
	SetData(v *AccountContactAddResponseBodyData) *AccountContactAddResponseBody
	GetData() *AccountContactAddResponseBodyData
	SetMessage(v string) *AccountContactAddResponseBody
	GetMessage() *string
	SetRequestId(v string) *AccountContactAddResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AccountContactAddResponseBody
	GetSuccess() *bool
}

type AccountContactAddResponseBody struct {
	// example:
	//
	// 200
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AccountContactAddResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
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

func (s AccountContactAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccountContactAddResponseBody) GoString() string {
	return s.String()
}

func (s *AccountContactAddResponseBody) GetCode() *string {
	return s.Code
}

func (s *AccountContactAddResponseBody) GetData() *AccountContactAddResponseBodyData {
	return s.Data
}

func (s *AccountContactAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AccountContactAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccountContactAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccountContactAddResponseBody) SetCode(v string) *AccountContactAddResponseBody {
	s.Code = &v
	return s
}

func (s *AccountContactAddResponseBody) SetData(v *AccountContactAddResponseBodyData) *AccountContactAddResponseBody {
	s.Data = v
	return s
}

func (s *AccountContactAddResponseBody) SetMessage(v string) *AccountContactAddResponseBody {
	s.Message = &v
	return s
}

func (s *AccountContactAddResponseBody) SetRequestId(v string) *AccountContactAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountContactAddResponseBody) SetSuccess(v bool) *AccountContactAddResponseBody {
	s.Success = &v
	return s
}

func (s *AccountContactAddResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AccountContactAddResponseBodyData struct {
	// example:
	//
	// xxx
	ContactId *int64                                        `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ErrorList []*AccountContactAddResponseBodyDataErrorList `json:"ErrorList,omitempty" xml:"ErrorList,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AccountContactAddResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AccountContactAddResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccountContactAddResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *AccountContactAddResponseBodyData) GetErrorList() []*AccountContactAddResponseBodyDataErrorList {
	return s.ErrorList
}

func (s *AccountContactAddResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *AccountContactAddResponseBodyData) SetContactId(v int64) *AccountContactAddResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *AccountContactAddResponseBodyData) SetErrorList(v []*AccountContactAddResponseBodyDataErrorList) *AccountContactAddResponseBodyData {
	s.ErrorList = v
	return s
}

func (s *AccountContactAddResponseBodyData) SetResult(v bool) *AccountContactAddResponseBodyData {
	s.Result = &v
	return s
}

func (s *AccountContactAddResponseBodyData) Validate() error {
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

type AccountContactAddResponseBodyDataErrorList struct {
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

func (s AccountContactAddResponseBodyDataErrorList) String() string {
	return dara.Prettify(s)
}

func (s AccountContactAddResponseBodyDataErrorList) GoString() string {
	return s.String()
}

func (s *AccountContactAddResponseBodyDataErrorList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AccountContactAddResponseBodyDataErrorList) GetErrorDesc() *string {
	return s.ErrorDesc
}

func (s *AccountContactAddResponseBodyDataErrorList) GetItem() *string {
	return s.Item
}

func (s *AccountContactAddResponseBodyDataErrorList) SetErrorCode(v string) *AccountContactAddResponseBodyDataErrorList {
	s.ErrorCode = &v
	return s
}

func (s *AccountContactAddResponseBodyDataErrorList) SetErrorDesc(v string) *AccountContactAddResponseBodyDataErrorList {
	s.ErrorDesc = &v
	return s
}

func (s *AccountContactAddResponseBodyDataErrorList) SetItem(v string) *AccountContactAddResponseBodyDataErrorList {
	s.Item = &v
	return s
}

func (s *AccountContactAddResponseBodyDataErrorList) Validate() error {
	return dara.Validate(s)
}
