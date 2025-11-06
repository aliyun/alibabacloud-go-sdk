// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchStaticAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAccessKey(v string) *FetchStaticAccountRequest
	GetAccountAccessKey() *string
	SetConsoleSessionId(v string) *FetchStaticAccountRequest
	GetConsoleSessionId() *string
	SetCreateTimeStamp(v int64) *FetchStaticAccountRequest
	GetCreateTimeStamp() *int64
	SetInstanceId(v string) *FetchStaticAccountRequest
	GetInstanceId() *string
	SetRemark(v string) *FetchStaticAccountRequest
	GetRemark() *string
	SetSKey(v string) *FetchStaticAccountRequest
	GetSKey() *string
	SetSecretSign(v string) *FetchStaticAccountRequest
	GetSecretSign() *string
	SetUserName(v string) *FetchStaticAccountRequest
	GetUserName() *string
}

type FetchStaticAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAccessKeyID
	AccountAccessKey *string `json:"AccountAccessKey,omitempty" xml:"AccountAccessKey,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1671175303522
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// amqp-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 备注示例
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22c2d7d1769cb53c5a6d9213248e2de524******
	SKey *string `json:"SKey,omitempty" xml:"SKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4c1a6367ce4c4255e9617326f9133ac635******
	SecretSign *string `json:"SecretSign,omitempty" xml:"SecretSign,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Mjo****************
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s FetchStaticAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchStaticAccountRequest) GoString() string {
	return s.String()
}

func (s *FetchStaticAccountRequest) GetAccountAccessKey() *string {
	return s.AccountAccessKey
}

func (s *FetchStaticAccountRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *FetchStaticAccountRequest) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *FetchStaticAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FetchStaticAccountRequest) GetRemark() *string {
	return s.Remark
}

func (s *FetchStaticAccountRequest) GetSKey() *string {
	return s.SKey
}

func (s *FetchStaticAccountRequest) GetSecretSign() *string {
	return s.SecretSign
}

func (s *FetchStaticAccountRequest) GetUserName() *string {
	return s.UserName
}

func (s *FetchStaticAccountRequest) SetAccountAccessKey(v string) *FetchStaticAccountRequest {
	s.AccountAccessKey = &v
	return s
}

func (s *FetchStaticAccountRequest) SetConsoleSessionId(v string) *FetchStaticAccountRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *FetchStaticAccountRequest) SetCreateTimeStamp(v int64) *FetchStaticAccountRequest {
	s.CreateTimeStamp = &v
	return s
}

func (s *FetchStaticAccountRequest) SetInstanceId(v string) *FetchStaticAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *FetchStaticAccountRequest) SetRemark(v string) *FetchStaticAccountRequest {
	s.Remark = &v
	return s
}

func (s *FetchStaticAccountRequest) SetSKey(v string) *FetchStaticAccountRequest {
	s.SKey = &v
	return s
}

func (s *FetchStaticAccountRequest) SetSecretSign(v string) *FetchStaticAccountRequest {
	s.SecretSign = &v
	return s
}

func (s *FetchStaticAccountRequest) SetUserName(v string) *FetchStaticAccountRequest {
	s.UserName = &v
	return s
}

func (s *FetchStaticAccountRequest) Validate() error {
	return dara.Validate(s)
}
