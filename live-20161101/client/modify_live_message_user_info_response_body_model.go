// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailList(v []*ModifyLiveMessageUserInfoResponseBodyFailList) *ModifyLiveMessageUserInfoResponseBody
	GetFailList() []*ModifyLiveMessageUserInfoResponseBodyFailList
	SetRequestId(v string) *ModifyLiveMessageUserInfoResponseBody
	GetRequestId() *string
	SetSuccessList(v []*ModifyLiveMessageUserInfoResponseBodySuccessList) *ModifyLiveMessageUserInfoResponseBody
	GetSuccessList() []*ModifyLiveMessageUserInfoResponseBodySuccessList
}

type ModifyLiveMessageUserInfoResponseBody struct {
	// The users whose information failed to be modified.
	FailList []*ModifyLiveMessageUserInfoResponseBodyFailList `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3271ACD2-F143-1204-AFDB-9A87C131****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The users whose information was modified.
	SuccessList []*ModifyLiveMessageUserInfoResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s ModifyLiveMessageUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageUserInfoResponseBody) GetFailList() []*ModifyLiveMessageUserInfoResponseBodyFailList {
	return s.FailList
}

func (s *ModifyLiveMessageUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveMessageUserInfoResponseBody) GetSuccessList() []*ModifyLiveMessageUserInfoResponseBodySuccessList {
	return s.SuccessList
}

func (s *ModifyLiveMessageUserInfoResponseBody) SetFailList(v []*ModifyLiveMessageUserInfoResponseBodyFailList) *ModifyLiveMessageUserInfoResponseBody {
	s.FailList = v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBody) SetRequestId(v string) *ModifyLiveMessageUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBody) SetSuccessList(v []*ModifyLiveMessageUserInfoResponseBodySuccessList) *ModifyLiveMessageUserInfoResponseBody {
	s.SuccessList = v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBody) Validate() error {
	if s.FailList != nil {
		for _, item := range s.FailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessList != nil {
		for _, item := range s.SuccessList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyLiveMessageUserInfoResponseBodyFailList struct {
	// The error code.
	//
	// example:
	//
	// 440
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the group to which the user belongs. For failed modification, the information of the user is not updated when you query the users in the group. You can try again after you check the failure reason and fix the issue.
	//
	// example:
	//
	// grouptest2
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The reason why the information of the user failed to be modified.
	//
	// example:
	//
	// group not exists or already deleted
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Indicates whether the group to which the user belongs is modified. In this case, the group is not modified.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyLiveMessageUserInfoResponseBodyFailList) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageUserInfoResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) GetCode() *int32 {
	return s.Code
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) GetReason() *string {
	return s.Reason
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) SetCode(v int32) *ModifyLiveMessageUserInfoResponseBodyFailList {
	s.Code = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) SetGroupId(v string) *ModifyLiveMessageUserInfoResponseBodyFailList {
	s.GroupId = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) SetReason(v string) *ModifyLiveMessageUserInfoResponseBodyFailList {
	s.Reason = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) SetSuccess(v bool) *ModifyLiveMessageUserInfoResponseBodyFailList {
	s.Success = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBodyFailList) Validate() error {
	return dara.Validate(s)
}

type ModifyLiveMessageUserInfoResponseBodySuccessList struct {
	// The ID of the group to which the user belongs. For successful modification, the information of the user is updated when you query the users in the group.
	//
	// example:
	//
	// grouptest1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the group to which the user belongs is modified. In this case, the group is modified.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyLiveMessageUserInfoResponseBodySuccessList) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageUserInfoResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageUserInfoResponseBodySuccessList) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyLiveMessageUserInfoResponseBodySuccessList) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyLiveMessageUserInfoResponseBodySuccessList) SetGroupId(v string) *ModifyLiveMessageUserInfoResponseBodySuccessList {
	s.GroupId = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBodySuccessList) SetSuccess(v bool) *ModifyLiveMessageUserInfoResponseBodySuccessList {
	s.Success = &v
	return s
}

func (s *ModifyLiveMessageUserInfoResponseBodySuccessList) Validate() error {
	return dara.Validate(s)
}
