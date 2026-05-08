// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSdkStreamMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SendSdkStreamMessageRequest
	GetData() *string
	SetHeader(v string) *SendSdkStreamMessageRequest
	GetHeader() *string
	SetModuleName(v string) *SendSdkStreamMessageRequest
	GetModuleName() *string
	SetOperationName(v string) *SendSdkStreamMessageRequest
	GetOperationName() *string
	SetUserId(v string) *SendSdkStreamMessageRequest
	GetUserId() *string
}

type SendSdkStreamMessageRequest struct {
	// example:
	//
	// {"test":""}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// header
	//
	// example:
	//
	// {}
	Header *string `json:"header,omitempty" xml:"header,omitempty"`
	// example:
	//
	// avatar
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// GetProject
	OperationName *string `json:"operationName,omitempty" xml:"operationName,omitempty"`
	// example:
	//
	// 123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendSdkStreamMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendSdkStreamMessageRequest) GoString() string {
	return s.String()
}

func (s *SendSdkStreamMessageRequest) GetData() *string {
	return s.Data
}

func (s *SendSdkStreamMessageRequest) GetHeader() *string {
	return s.Header
}

func (s *SendSdkStreamMessageRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *SendSdkStreamMessageRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *SendSdkStreamMessageRequest) GetUserId() *string {
	return s.UserId
}

func (s *SendSdkStreamMessageRequest) SetData(v string) *SendSdkStreamMessageRequest {
	s.Data = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetHeader(v string) *SendSdkStreamMessageRequest {
	s.Header = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetModuleName(v string) *SendSdkStreamMessageRequest {
	s.ModuleName = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetOperationName(v string) *SendSdkStreamMessageRequest {
	s.OperationName = &v
	return s
}

func (s *SendSdkStreamMessageRequest) SetUserId(v string) *SendSdkStreamMessageRequest {
	s.UserId = &v
	return s
}

func (s *SendSdkStreamMessageRequest) Validate() error {
	return dara.Validate(s)
}
