// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCallV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *StartCallV2Request
	GetAccountName() *string
	SetCallee(v string) *StartCallV2Request
	GetCallee() *string
	SetCaller(v string) *StartCallV2Request
	GetCaller() *string
	SetCallerType(v int32) *StartCallV2Request
	GetCallerType() *int32
	SetClientToken(v string) *StartCallV2Request
	GetClientToken() *string
	SetInstanceId(v string) *StartCallV2Request
	GetInstanceId() *string
}

type StartCallV2Request struct {
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
	// 136****1111
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9065****
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CallerType *int32 `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
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

func (s StartCallV2Request) String() string {
	return dara.Prettify(s)
}

func (s StartCallV2Request) GoString() string {
	return s.String()
}

func (s *StartCallV2Request) GetAccountName() *string {
	return s.AccountName
}

func (s *StartCallV2Request) GetCallee() *string {
	return s.Callee
}

func (s *StartCallV2Request) GetCaller() *string {
	return s.Caller
}

func (s *StartCallV2Request) GetCallerType() *int32 {
	return s.CallerType
}

func (s *StartCallV2Request) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartCallV2Request) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartCallV2Request) SetAccountName(v string) *StartCallV2Request {
	s.AccountName = &v
	return s
}

func (s *StartCallV2Request) SetCallee(v string) *StartCallV2Request {
	s.Callee = &v
	return s
}

func (s *StartCallV2Request) SetCaller(v string) *StartCallV2Request {
	s.Caller = &v
	return s
}

func (s *StartCallV2Request) SetCallerType(v int32) *StartCallV2Request {
	s.CallerType = &v
	return s
}

func (s *StartCallV2Request) SetClientToken(v string) *StartCallV2Request {
	s.ClientToken = &v
	return s
}

func (s *StartCallV2Request) SetInstanceId(v string) *StartCallV2Request {
	s.InstanceId = &v
	return s
}

func (s *StartCallV2Request) Validate() error {
	return dara.Validate(s)
}
