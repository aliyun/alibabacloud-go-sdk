// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserInfoInChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsChannelExist(v bool) *DescribeUserInfoInChannelResponseBody
	GetIsChannelExist() *bool
	SetIsInChannel(v bool) *DescribeUserInfoInChannelResponseBody
	GetIsInChannel() *bool
	SetProperty(v []*DescribeUserInfoInChannelResponseBodyProperty) *DescribeUserInfoInChannelResponseBody
	GetProperty() []*DescribeUserInfoInChannelResponseBodyProperty
	SetRequestId(v string) *DescribeUserInfoInChannelResponseBody
	GetRequestId() *string
	SetTimestamp(v int32) *DescribeUserInfoInChannelResponseBody
	GetTimestamp() *int32
}

type DescribeUserInfoInChannelResponseBody struct {
	// Indicates whether the channel exists. Valid values:
	//
	// - **true**: The channel exists.
	//
	// - **false**: The channel does not exist.
	//
	// example:
	//
	// true
	IsChannelExist *bool `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	// Indicates whether the user is currently in the channel. Valid values:
	//
	// - **true**: The user is in the channel.
	//
	// - **false**: The user is not in the channel.
	//
	// example:
	//
	// true
	IsInChannel *bool `json:"IsInChannel,omitempty" xml:"IsInChannel,omitempty"`
	// The information of the user in the channel.
	Property []*DescribeUserInfoInChannelResponseBodyProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp of the current moment (UTC time).
	//
	// example:
	//
	// 1557909133
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeUserInfoInChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoInChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponseBody) GetIsChannelExist() *bool {
	return s.IsChannelExist
}

func (s *DescribeUserInfoInChannelResponseBody) GetIsInChannel() *bool {
	return s.IsInChannel
}

func (s *DescribeUserInfoInChannelResponseBody) GetProperty() []*DescribeUserInfoInChannelResponseBodyProperty {
	return s.Property
}

func (s *DescribeUserInfoInChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserInfoInChannelResponseBody) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *DescribeUserInfoInChannelResponseBody) SetIsChannelExist(v bool) *DescribeUserInfoInChannelResponseBody {
	s.IsChannelExist = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetIsInChannel(v bool) *DescribeUserInfoInChannelResponseBody {
	s.IsInChannel = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetProperty(v []*DescribeUserInfoInChannelResponseBodyProperty) *DescribeUserInfoInChannelResponseBody {
	s.Property = v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetRequestId(v string) *DescribeUserInfoInChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetTimestamp(v int32) *DescribeUserInfoInChannelResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) Validate() error {
	if s.Property != nil {
		for _, item := range s.Property {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserInfoInChannelResponseBodyProperty struct {
	// The time when the user joined the meeting (UTC time).
	//
	// example:
	//
	// 1557909133
	Join *int32 `json:"Join,omitempty" xml:"Join,omitempty"`
	// The role type in the meeting. Valid values:
	//
	// - **0**: The user role in communication mode.
	//
	//
	//
	// - **1**: The interactive role.
	//
	//
	//
	// - **2**: The live (view only) role.
	//
	// example:
	//
	// 1
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
	// The session ID.
	//
	// example:
	//
	// xa744sxx8rtobgj****
	Session *string `json:"Session,omitempty" xml:"Session,omitempty"`
}

func (s DescribeUserInfoInChannelResponseBodyProperty) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoInChannelResponseBodyProperty) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) GetJoin() *int32 {
	return s.Join
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) GetRole() *int32 {
	return s.Role
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) GetSession() *string {
	return s.Session
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetJoin(v int32) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Join = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetRole(v int32) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Role = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetSession(v string) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Session = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) Validate() error {
	return dara.Validate(s)
}
