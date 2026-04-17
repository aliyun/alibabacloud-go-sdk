// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSaslUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeSaslUsersResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeSaslUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSaslUsersResponseBody
	GetRequestId() *string
	SetSaslUserList(v *DescribeSaslUsersResponseBodySaslUserList) *DescribeSaslUsersResponseBody
	GetSaslUserList() *DescribeSaslUsersResponseBodySaslUserList
	SetSuccess(v bool) *DescribeSaslUsersResponseBody
	GetSuccess() *bool
}

type DescribeSaslUsersResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9E3B3592-5994-4F65-A61E-E62A77A****
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SaslUserList *DescribeSaslUsersResponseBodySaslUserList `json:"SaslUserList,omitempty" xml:"SaslUserList,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSaslUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSaslUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeSaslUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSaslUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSaslUsersResponseBody) GetSaslUserList() *DescribeSaslUsersResponseBodySaslUserList {
	return s.SaslUserList
}

func (s *DescribeSaslUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSaslUsersResponseBody) SetCode(v int32) *DescribeSaslUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetMessage(v string) *DescribeSaslUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetRequestId(v string) *DescribeSaslUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetSaslUserList(v *DescribeSaslUsersResponseBodySaslUserList) *DescribeSaslUsersResponseBody {
	s.SaslUserList = v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetSuccess(v bool) *DescribeSaslUsersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSaslUsersResponseBody) Validate() error {
	if s.SaslUserList != nil {
		if err := s.SaslUserList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSaslUsersResponseBodySaslUserList struct {
	SaslUserVO []*DescribeSaslUsersResponseBodySaslUserListSaslUserVO `json:"SaslUserVO,omitempty" xml:"SaslUserVO,omitempty" type:"Repeated"`
}

func (s DescribeSaslUsersResponseBodySaslUserList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSaslUsersResponseBodySaslUserList) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponseBodySaslUserList) GetSaslUserVO() []*DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	return s.SaslUserVO
}

func (s *DescribeSaslUsersResponseBodySaslUserList) SetSaslUserVO(v []*DescribeSaslUsersResponseBodySaslUserListSaslUserVO) *DescribeSaslUsersResponseBodySaslUserList {
	s.SaslUserVO = v
	return s
}

func (s *DescribeSaslUsersResponseBodySaslUserList) Validate() error {
	if s.SaslUserVO != nil {
		for _, item := range s.SaslUserVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSaslUsersResponseBodySaslUserListSaslUserVO struct {
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	Password  *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Username  *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeSaslUsersResponseBodySaslUserListSaslUserVO) String() string {
	return dara.Prettify(s)
}

func (s DescribeSaslUsersResponseBodySaslUserListSaslUserVO) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) GetMechanism() *string {
	return s.Mechanism
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) GetPassword() *string {
	return s.Password
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) GetType() *string {
	return s.Type
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) GetUsername() *string {
	return s.Username
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetMechanism(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Mechanism = &v
	return s
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetPassword(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Password = &v
	return s
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetType(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Type = &v
	return s
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetUsername(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Username = &v
	return s
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) Validate() error {
	return dara.Validate(s)
}
