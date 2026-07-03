// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRdUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListRdUsersResponseBodyData) *ListRdUsersResponseBody
	GetData() []*ListRdUsersResponseBodyData
	SetRequestId(v string) *ListRdUsersResponseBody
	GetRequestId() *string
}

type ListRdUsersResponseBody struct {
	// The data returned.
	Data []*ListRdUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRdUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRdUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRdUsersResponseBody) GetData() []*ListRdUsersResponseBodyData {
	return s.Data
}

func (s *ListRdUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRdUsersResponseBody) SetData(v []*ListRdUsersResponseBodyData) *ListRdUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListRdUsersResponseBody) SetRequestId(v string) *ListRdUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRdUsersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRdUsersResponseBodyData struct {
	// Indicates whether the account is delegated to view its own resources.
	//
	// example:
	//
	// true
	DelegatedOrNot *bool `json:"DelegatedOrNot,omitempty" xml:"DelegatedOrNot,omitempty"`
	// Indicates whether the account is managed by the multi-account control feature of Threat Analysis. Valid values:
	//
	// - true: The account is managed.
	//
	// - false: The account is not managed.
	//
	// example:
	//
	// true
	Joined *bool `json:"Joined,omitempty" xml:"Joined,omitempty"`
	// The time when the account was added.
	//
	// example:
	//
	// 2013-10-01 00:00:00
	JoinedTime *string `json:"JoinedTime,omitempty" xml:"JoinedTime,omitempty"`
	// The ID of the Alibaba Cloud account that purchased Threat Analysis.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The ID of the member Alibaba Cloud account.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The name of the member Alibaba Cloud account.
	//
	// example:
	//
	// sas_account_xxx
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
}

func (s ListRdUsersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRdUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRdUsersResponseBodyData) GetDelegatedOrNot() *bool {
	return s.DelegatedOrNot
}

func (s *ListRdUsersResponseBodyData) GetJoined() *bool {
	return s.Joined
}

func (s *ListRdUsersResponseBodyData) GetJoinedTime() *string {
	return s.JoinedTime
}

func (s *ListRdUsersResponseBodyData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *ListRdUsersResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListRdUsersResponseBodyData) GetSubUserName() *string {
	return s.SubUserName
}

func (s *ListRdUsersResponseBodyData) SetDelegatedOrNot(v bool) *ListRdUsersResponseBodyData {
	s.DelegatedOrNot = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetJoined(v bool) *ListRdUsersResponseBodyData {
	s.Joined = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetJoinedTime(v string) *ListRdUsersResponseBodyData {
	s.JoinedTime = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetMainUserId(v int64) *ListRdUsersResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetSubUserId(v int64) *ListRdUsersResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *ListRdUsersResponseBodyData) SetSubUserName(v string) *ListRdUsersResponseBodyData {
	s.SubUserName = &v
	return s
}

func (s *ListRdUsersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
