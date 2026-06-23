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
	// The remarks on the static user.
	//
	// example:
	//
	// **	- environment
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The AccessKey ID of your Alibaba Cloud account or RAM user. For more information about how to obtain an AccessKey ID, see [Create an AccessKey](https://help.aliyun.com/document_detail/116401.html).
	//
	// > If you use the AccessKey of a RAM user to create a static username and password to access ApsaraMQ for RabbitMQ and to send and receive messages, make sure that the RAM user is granted the required permissions. For more information, see [RAM access policies](https://help.aliyun.com/document_detail/146559.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAccessKeyID
	AccountAccessKey *string `json:"accountAccessKey,omitempty" xml:"accountAccessKey,omitempty"`
	// The timestamp that indicates when the username and password are created. Unit: milliseconds.
	//
	// > This timestamp is used to calculate the static password. You can customize this value. This is not the timestamp that the system generates when the username and password are created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1671175303522
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance. This specifies the instance for which you want to create a static username and password.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The signature of the AccessKey secret. The system calculates the static password based on the signature, the AccessKey secret signature, and the username.
	//
	// The AccessKey secret signature is calculated using the HmacSHA1 algorithm on the creation timestamp of the specified username and the AccessKey ID. For more information about how to calculate the signature, see the **Signature algorithm sample code*	- section in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c1a6367ce4c4255e9617326f9133ac635******
	SecretSign *string `json:"secretSign,omitempty" xml:"secretSign,omitempty"`
	// The signature. The system calculates the static password based on the signature, the AccessKey secret signature, and the username.
	//
	// The signature is calculated using the HmacSHA1 algorithm on the creation timestamp of the specified username and the AccessKey ID. For more information about how to calculate the signature, see the **Signature algorithm sample code*	- section in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22c2d7d1769cb53c5a6d9213248e2de524******
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The static username that you want to create.
	//
	// The value of this parameter is a Base64-encoded string that is constructed from the instance ID and the AccessKey ID. For more information about how to calculate the value, see the **Username calculation sample code*	- section in this topic.
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
