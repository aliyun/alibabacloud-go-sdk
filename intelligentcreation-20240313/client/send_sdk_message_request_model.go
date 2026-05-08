// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSdkMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SendSdkMessageRequest
	GetData() *string
	SetHeader(v string) *SendSdkMessageRequest
	GetHeader() *string
	SetModuleName(v string) *SendSdkMessageRequest
	GetModuleName() *string
	SetOperationName(v string) *SendSdkMessageRequest
	GetOperationName() *string
	SetUserId(v string) *SendSdkMessageRequest
	GetUserId() *string
}

type SendSdkMessageRequest struct {
	// example:
	//
	// {}
	Data   *string `json:"data,omitempty" xml:"data,omitempty"`
	Header *string `json:"header,omitempty" xml:"header,omitempty"`
	// example:
	//
	// avatar
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// getProject
	OperationName *string `json:"operationName,omitempty" xml:"operationName,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendSdkMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendSdkMessageRequest) GoString() string {
	return s.String()
}

func (s *SendSdkMessageRequest) GetData() *string {
	return s.Data
}

func (s *SendSdkMessageRequest) GetHeader() *string {
	return s.Header
}

func (s *SendSdkMessageRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *SendSdkMessageRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *SendSdkMessageRequest) GetUserId() *string {
	return s.UserId
}

func (s *SendSdkMessageRequest) SetData(v string) *SendSdkMessageRequest {
	s.Data = &v
	return s
}

func (s *SendSdkMessageRequest) SetHeader(v string) *SendSdkMessageRequest {
	s.Header = &v
	return s
}

func (s *SendSdkMessageRequest) SetModuleName(v string) *SendSdkMessageRequest {
	s.ModuleName = &v
	return s
}

func (s *SendSdkMessageRequest) SetOperationName(v string) *SendSdkMessageRequest {
	s.OperationName = &v
	return s
}

func (s *SendSdkMessageRequest) SetUserId(v string) *SendSdkMessageRequest {
	s.UserId = &v
	return s
}

func (s *SendSdkMessageRequest) Validate() error {
	return dara.Validate(s)
}
