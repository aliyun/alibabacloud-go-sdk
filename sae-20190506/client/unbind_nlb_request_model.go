// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindNlbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UnbindNlbRequest
	GetAppId() *string
	SetNlbId(v string) *UnbindNlbRequest
	GetNlbId() *string
	SetPort(v int32) *UnbindNlbRequest
	GetPort() *int32
	SetProtocol(v string) *UnbindNlbRequest
	GetProtocol() *string
}

type UnbindNlbRequest struct {
	// A short description of struct
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of NLB instance.
	//
	// example:
	//
	// nlb-7z7jjbzz44d82c9***
	NlbId *string `json:"NlbId,omitempty" xml:"NlbId,omitempty"`
	// The listener port of the instance. Valid values: 0 to 65535.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **TCP**.
	//
	// 	- **UDP**.
	//
	// 	- **TCPSSL**.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s UnbindNlbRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindNlbRequest) GoString() string {
	return s.String()
}

func (s *UnbindNlbRequest) GetAppId() *string {
	return s.AppId
}

func (s *UnbindNlbRequest) GetNlbId() *string {
	return s.NlbId
}

func (s *UnbindNlbRequest) GetPort() *int32 {
	return s.Port
}

func (s *UnbindNlbRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UnbindNlbRequest) SetAppId(v string) *UnbindNlbRequest {
	s.AppId = &v
	return s
}

func (s *UnbindNlbRequest) SetNlbId(v string) *UnbindNlbRequest {
	s.NlbId = &v
	return s
}

func (s *UnbindNlbRequest) SetPort(v int32) *UnbindNlbRequest {
	s.Port = &v
	return s
}

func (s *UnbindNlbRequest) SetProtocol(v string) *UnbindNlbRequest {
	s.Protocol = &v
	return s
}

func (s *UnbindNlbRequest) Validate() error {
	return dara.Validate(s)
}
