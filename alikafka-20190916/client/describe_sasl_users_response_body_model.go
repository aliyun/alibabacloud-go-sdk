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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Simple Authentication and Security Layer (SASL) users.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeSaslUsersResponseBodySaslUserListSaslUserVO struct {
	// The encryption method.
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// SCRAM-SHA-256
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	// The password.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The type of the SASL user. Valid values:
	//
	// 	- **plain**: a simple mechanism that uses usernames and passwords to verify user identities. ApsaraMQ for Kafka provides an improved PLAIN mechanism that allows you to dynamically add SASL users without the need to restart an instance.
	//
	// 	- **SCRAM**: a mechanism that uses usernames and passwords to verify user identities. Compared with the PLAIN mechanism, this mechanism provides better security protection. ApsaraMQ for Kafka uses the SCRAM-SHA-256 algorithm.
	//
	// 	- **LDAP**: This value is available only for the SASL users of ApsaraMQ for Confluent instances.
	//
	// Default value: **plain**.
	//
	// example:
	//
	// scram
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username.
	//
	// example:
	//
	// test12***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
