// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemark(v string) *CreateAccountRequest
	GetRemark() *string
	SetAccountAccessKey(v string) *CreateAccountRequest
	GetAccountAccessKey() *string
	SetCreateTimestamp(v int64) *CreateAccountRequest
	GetCreateTimestamp() *int64
	SetInstanceId(v string) *CreateAccountRequest
	GetInstanceId() *string
	SetSecretSign(v string) *CreateAccountRequest
	GetSecretSign() *string
	SetSignature(v string) *CreateAccountRequest
	GetSignature() *string
	SetUserName(v string) *CreateAccountRequest
	GetUserName() *string
}

type CreateAccountRequest struct {
	// example:
	//
	// ***环境
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The AccessKey ID of your Alibaba Cloud account or RAM user. For information about how to obtain an AccessKey pair, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
	//
	// >  If you use the pair of static username and password that is created by using the Accesskey pair of a RAM user to access ApsaraMQ for RabbitMQ to send and receive messages, make sure that the RAM user is granted the required permissions. For more information, see [RAM policies](https://help.aliyun.com/document_detail/146559.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAccessKeyID
	AccountAccessKey *string `json:"accountAccessKey,omitempty" xml:"accountAccessKey,omitempty"`
	// The timestamp that indicates when the password is created. Unit: milliseconds.
	//
	// >  This timestamp is specified by you and is used to generate a static password. The timestamp is not the timestamp that indicates when the system generates the password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671175303522
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The ID of the instance for which you want to create a pair of static username and password.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The AccessKey secret signature. The system generates a static password based on the signature in the request, the AccessKey secret signature, and the username.
	//
	// The system uses the HMAC-SHA1 algorithm to generate the AccessKey secret signature based on the timestamp that indicates when the username is created and the AccessKey ID. For more information, see the **"Sample code on how to generate a signature"*	- section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c1a6367ce4c4255e9617326f9133ac635******
	SecretSign *string `json:"secretSign,omitempty" xml:"secretSign,omitempty"`
	// The signature. The system generates a static password based on the signature in the request, the AccessKey secret signature, and the username.
	//
	// The system uses the HMAC-SHA1 algorithm to generate the signature based on the timestamp that indicates when the username is created and the AccessKey ID. For more information, see the **"Sample code on how to generate a signature"*	- section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22c2d7d1769cb53c5a6d9213248e2de524******
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The static username that you want to create.
	//
	// The value of this parameter is a Base64-encoded string that is generated based on the instance ID and AccessKey ID. For more information, see the "**Sample code on how to generate a username**" section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// Mjo****************
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateAccountRequest) GetAccountAccessKey() *string {
	return s.AccountAccessKey
}

func (s *CreateAccountRequest) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *CreateAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAccountRequest) GetSecretSign() *string {
	return s.SecretSign
}

func (s *CreateAccountRequest) GetSignature() *string {
	return s.Signature
}

func (s *CreateAccountRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateAccountRequest) SetRemark(v string) *CreateAccountRequest {
	s.Remark = &v
	return s
}

func (s *CreateAccountRequest) SetAccountAccessKey(v string) *CreateAccountRequest {
	s.AccountAccessKey = &v
	return s
}

func (s *CreateAccountRequest) SetCreateTimestamp(v int64) *CreateAccountRequest {
	s.CreateTimestamp = &v
	return s
}

func (s *CreateAccountRequest) SetInstanceId(v string) *CreateAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetSecretSign(v string) *CreateAccountRequest {
	s.SecretSign = &v
	return s
}

func (s *CreateAccountRequest) SetSignature(v string) *CreateAccountRequest {
	s.Signature = &v
	return s
}

func (s *CreateAccountRequest) SetUserName(v string) *CreateAccountRequest {
	s.UserName = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}
