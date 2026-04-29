// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePolarClawMCPServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RemovePolarClawMCPServerResponseBody
	GetApplicationId() *string
	SetCode(v int32) *RemovePolarClawMCPServerResponseBody
	GetCode() *int32
	SetMessage(v string) *RemovePolarClawMCPServerResponseBody
	GetMessage() *string
	SetOk(v bool) *RemovePolarClawMCPServerResponseBody
	GetOk() *bool
	SetRequestId(v string) *RemovePolarClawMCPServerResponseBody
	GetRequestId() *string
	SetServerName(v string) *RemovePolarClawMCPServerResponseBody
	GetServerName() *string
}

type RemovePolarClawMCPServerResponseBody struct {
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
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test-v1
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s RemovePolarClawMCPServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePolarClawMCPServerResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePolarClawMCPServerResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RemovePolarClawMCPServerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemovePolarClawMCPServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemovePolarClawMCPServerResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *RemovePolarClawMCPServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePolarClawMCPServerResponseBody) GetServerName() *string {
	return s.ServerName
}

func (s *RemovePolarClawMCPServerResponseBody) SetApplicationId(v string) *RemovePolarClawMCPServerResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *RemovePolarClawMCPServerResponseBody) SetCode(v int32) *RemovePolarClawMCPServerResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePolarClawMCPServerResponseBody) SetMessage(v string) *RemovePolarClawMCPServerResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePolarClawMCPServerResponseBody) SetOk(v bool) *RemovePolarClawMCPServerResponseBody {
	s.Ok = &v
	return s
}

func (s *RemovePolarClawMCPServerResponseBody) SetRequestId(v string) *RemovePolarClawMCPServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePolarClawMCPServerResponseBody) SetServerName(v string) *RemovePolarClawMCPServerResponseBody {
	s.ServerName = &v
	return s
}

func (s *RemovePolarClawMCPServerResponseBody) Validate() error {
	return dara.Validate(s)
}
