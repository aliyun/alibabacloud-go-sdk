// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateAccountResponseBody
	GetCode() *int32
	SetData(v *CreateAccountResponseBodyData) *CreateAccountResponseBody
	GetData() *CreateAccountResponseBodyData
	SetMessage(v string) *CreateAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAccountResponseBody
	GetSuccess() *bool
}

type CreateAccountResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92385FD2-624A-48C9-8FB5-753F2AFA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateAccountResponseBody) GetData() *CreateAccountResponseBodyData {
	return s.Data
}

func (s *CreateAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAccountResponseBody) SetCode(v int32) *CreateAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAccountResponseBody) SetData(v *CreateAccountResponseBodyData) *CreateAccountResponseBody {
	s.Data = v
	return s
}

func (s *CreateAccountResponseBody) SetMessage(v string) *CreateAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) SetSuccess(v bool) *CreateAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAccountResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAccountResponseBodyData struct {
	// The AccessKey ID that is used to create the password.
	//
	// example:
	//
	// LTAI****************
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The timestamp that indicates when the password was created. Unit: milliseconds.
	//
	// example:
	//
	// 1671175303522
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Alibaba Cloud account ID or RAM user to which the AccessKey pair that is used to create the static username and password belongs.
	//
	// example:
	//
	// 15657*********01
	MasterUId *int64 `json:"MasterUId,omitempty" xml:"MasterUId,omitempty"`
	// The created static password.
	//
	// example:
	//
	// NEMxQTYzNjdDRTVDNDI1NUU5NjE3**************1MzNGODoxNjcxMTc1MzEzNTIy
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// ***环境
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The created static username.
	//
	// example:
	//
	// MjphbXFwLWNuLXVxbTJ6cjc2djAwMzpMVEFJNX*******ZNMWVSWnRFSjZ2Zm8=
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBodyData) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateAccountResponseBodyData) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *CreateAccountResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAccountResponseBodyData) GetMasterUId() *int64 {
	return s.MasterUId
}

func (s *CreateAccountResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *CreateAccountResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *CreateAccountResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *CreateAccountResponseBodyData) SetAccessKey(v string) *CreateAccountResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetCreateTimeStamp(v int64) *CreateAccountResponseBodyData {
	s.CreateTimeStamp = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetInstanceId(v string) *CreateAccountResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetMasterUId(v int64) *CreateAccountResponseBodyData {
	s.MasterUId = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetPassword(v string) *CreateAccountResponseBodyData {
	s.Password = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetRemark(v string) *CreateAccountResponseBodyData {
	s.Remark = &v
	return s
}

func (s *CreateAccountResponseBodyData) SetUserName(v string) *CreateAccountResponseBodyData {
	s.UserName = &v
	return s
}

func (s *CreateAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
