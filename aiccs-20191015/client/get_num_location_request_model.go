// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetNumLocationRequest
	GetClientToken() *string
	SetInstanceId(v string) *GetNumLocationRequest
	GetInstanceId() *string
	SetPhoneNum(v string) *GetNumLocationRequest
	GetPhoneNum() *string
}

type GetNumLocationRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1360987****
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s GetNumLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNumLocationRequest) GoString() string {
	return s.String()
}

func (s *GetNumLocationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetNumLocationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNumLocationRequest) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *GetNumLocationRequest) SetClientToken(v string) *GetNumLocationRequest {
	s.ClientToken = &v
	return s
}

func (s *GetNumLocationRequest) SetInstanceId(v string) *GetNumLocationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNumLocationRequest) SetPhoneNum(v string) *GetNumLocationRequest {
	s.PhoneNum = &v
	return s
}

func (s *GetNumLocationRequest) Validate() error {
	return dara.Validate(s)
}
