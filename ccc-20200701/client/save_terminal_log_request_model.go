// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTerminalLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *SaveTerminalLogRequest
	GetAppName() *string
	SetCallId(v string) *SaveTerminalLogRequest
	GetCallId() *string
	SetContent(v string) *SaveTerminalLogRequest
	GetContent() *string
	SetDataType(v int32) *SaveTerminalLogRequest
	GetDataType() *int32
	SetInstanceId(v string) *SaveTerminalLogRequest
	GetInstanceId() *string
	SetJobId(v string) *SaveTerminalLogRequest
	GetJobId() *string
	SetMethodName(v string) *SaveTerminalLogRequest
	GetMethodName() *string
	SetStatus(v string) *SaveTerminalLogRequest
	GetStatus() *string
	SetUniqueRequestId(v string) *SaveTerminalLogRequest
	GetUniqueRequestId() *string
}

type SaveTerminalLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// CCCClient
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e13c9740-1e37-123b-21b6-00163e352f9
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// none
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-b8b0ca63-330c-4e65-8ae3-9de2c7ce7683
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// TurnServerTest
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// none
	UniqueRequestId *string `json:"UniqueRequestId,omitempty" xml:"UniqueRequestId,omitempty"`
}

func (s SaveTerminalLogRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTerminalLogRequest) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogRequest) GetAppName() *string {
	return s.AppName
}

func (s *SaveTerminalLogRequest) GetCallId() *string {
	return s.CallId
}

func (s *SaveTerminalLogRequest) GetContent() *string {
	return s.Content
}

func (s *SaveTerminalLogRequest) GetDataType() *int32 {
	return s.DataType
}

func (s *SaveTerminalLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveTerminalLogRequest) GetJobId() *string {
	return s.JobId
}

func (s *SaveTerminalLogRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *SaveTerminalLogRequest) GetStatus() *string {
	return s.Status
}

func (s *SaveTerminalLogRequest) GetUniqueRequestId() *string {
	return s.UniqueRequestId
}

func (s *SaveTerminalLogRequest) SetAppName(v string) *SaveTerminalLogRequest {
	s.AppName = &v
	return s
}

func (s *SaveTerminalLogRequest) SetCallId(v string) *SaveTerminalLogRequest {
	s.CallId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetContent(v string) *SaveTerminalLogRequest {
	s.Content = &v
	return s
}

func (s *SaveTerminalLogRequest) SetDataType(v int32) *SaveTerminalLogRequest {
	s.DataType = &v
	return s
}

func (s *SaveTerminalLogRequest) SetInstanceId(v string) *SaveTerminalLogRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetJobId(v string) *SaveTerminalLogRequest {
	s.JobId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetMethodName(v string) *SaveTerminalLogRequest {
	s.MethodName = &v
	return s
}

func (s *SaveTerminalLogRequest) SetStatus(v string) *SaveTerminalLogRequest {
	s.Status = &v
	return s
}

func (s *SaveTerminalLogRequest) SetUniqueRequestId(v string) *SaveTerminalLogRequest {
	s.UniqueRequestId = &v
	return s
}

func (s *SaveTerminalLogRequest) Validate() error {
	return dara.Validate(s)
}
