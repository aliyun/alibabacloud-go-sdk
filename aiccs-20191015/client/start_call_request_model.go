// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *StartCallRequest
	GetAccountName() *string
	SetCallee(v string) *StartCallRequest
	GetCallee() *string
	SetCaller(v string) *StartCallRequest
	GetCaller() *string
	SetClientToken(v string) *StartCallRequest
	GetClientToken() *string
	SetInstanceId(v string) *StartCallRequest
	GetInstanceId() *string
}

type StartCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1360987****
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 906****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartCallRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCallRequest) GoString() string {
	return s.String()
}

func (s *StartCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *StartCallRequest) GetCallee() *string {
	return s.Callee
}

func (s *StartCallRequest) GetCaller() *string {
	return s.Caller
}

func (s *StartCallRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartCallRequest) SetAccountName(v string) *StartCallRequest {
	s.AccountName = &v
	return s
}

func (s *StartCallRequest) SetCallee(v string) *StartCallRequest {
	s.Callee = &v
	return s
}

func (s *StartCallRequest) SetCaller(v string) *StartCallRequest {
	s.Caller = &v
	return s
}

func (s *StartCallRequest) SetClientToken(v string) *StartCallRequest {
	s.ClientToken = &v
	return s
}

func (s *StartCallRequest) SetInstanceId(v string) *StartCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartCallRequest) Validate() error {
	return dara.Validate(s)
}
