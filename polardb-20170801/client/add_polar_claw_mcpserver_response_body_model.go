// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarClawMCPServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AddPolarClawMCPServerResponseBody
	GetApplicationId() *string
	SetCode(v int32) *AddPolarClawMCPServerResponseBody
	GetCode() *int32
	SetMessage(v string) *AddPolarClawMCPServerResponseBody
	GetMessage() *string
	SetOk(v bool) *AddPolarClawMCPServerResponseBody
	GetOk() *bool
	SetRequestId(v string) *AddPolarClawMCPServerResponseBody
	GetRequestId() *string
	SetServerName(v string) *AddPolarClawMCPServerResponseBody
	GetServerName() *string
}

type AddPolarClawMCPServerResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// True
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// 2281C6C9-CBAB-1AFD-8400-670750CF6025_2212
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-v1
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s AddPolarClawMCPServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPolarClawMCPServerResponseBody) GoString() string {
	return s.String()
}

func (s *AddPolarClawMCPServerResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AddPolarClawMCPServerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddPolarClawMCPServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPolarClawMCPServerResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *AddPolarClawMCPServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPolarClawMCPServerResponseBody) GetServerName() *string {
	return s.ServerName
}

func (s *AddPolarClawMCPServerResponseBody) SetApplicationId(v string) *AddPolarClawMCPServerResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *AddPolarClawMCPServerResponseBody) SetCode(v int32) *AddPolarClawMCPServerResponseBody {
	s.Code = &v
	return s
}

func (s *AddPolarClawMCPServerResponseBody) SetMessage(v string) *AddPolarClawMCPServerResponseBody {
	s.Message = &v
	return s
}

func (s *AddPolarClawMCPServerResponseBody) SetOk(v bool) *AddPolarClawMCPServerResponseBody {
	s.Ok = &v
	return s
}

func (s *AddPolarClawMCPServerResponseBody) SetRequestId(v string) *AddPolarClawMCPServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPolarClawMCPServerResponseBody) SetServerName(v string) *AddPolarClawMCPServerResponseBody {
	s.ServerName = &v
	return s
}

func (s *AddPolarClawMCPServerResponseBody) Validate() error {
	return dara.Validate(s)
}
