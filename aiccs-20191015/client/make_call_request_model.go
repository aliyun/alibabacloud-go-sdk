// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *MakeCallRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *MakeCallRequest
	GetCallingNumber() *string
	SetCommandCode(v string) *MakeCallRequest
	GetCommandCode() *string
	SetExtInfo(v string) *MakeCallRequest
	GetExtInfo() *string
	SetOuterAccountId(v string) *MakeCallRequest
	GetOuterAccountId() *string
	SetOuterAccountType(v string) *MakeCallRequest
	GetOuterAccountType() *string
}

type MakeCallRequest struct {
	// Called number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130****0000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0571456****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Operation code.
	//
	// This parameter is required.
	//
	// example:
	//
	// outBound_Call
	CommandCode *string `json:"CommandCode,omitempty" xml:"CommandCode,omitempty"`
	// Pass-through business information.
	//
	// example:
	//
	// {"bizId": 23323}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// External account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	OuterAccountId *string `json:"OuterAccountId,omitempty" xml:"OuterAccountId,omitempty"`
	// External account type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIPAY
	OuterAccountType *string `json:"OuterAccountType,omitempty" xml:"OuterAccountType,omitempty"`
}

func (s MakeCallRequest) String() string {
	return dara.Prettify(s)
}

func (s MakeCallRequest) GoString() string {
	return s.String()
}

func (s *MakeCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *MakeCallRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *MakeCallRequest) GetCommandCode() *string {
	return s.CommandCode
}

func (s *MakeCallRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *MakeCallRequest) GetOuterAccountId() *string {
	return s.OuterAccountId
}

func (s *MakeCallRequest) GetOuterAccountType() *string {
	return s.OuterAccountType
}

func (s *MakeCallRequest) SetCalledNumber(v string) *MakeCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *MakeCallRequest) SetCallingNumber(v string) *MakeCallRequest {
	s.CallingNumber = &v
	return s
}

func (s *MakeCallRequest) SetCommandCode(v string) *MakeCallRequest {
	s.CommandCode = &v
	return s
}

func (s *MakeCallRequest) SetExtInfo(v string) *MakeCallRequest {
	s.ExtInfo = &v
	return s
}

func (s *MakeCallRequest) SetOuterAccountId(v string) *MakeCallRequest {
	s.OuterAccountId = &v
	return s
}

func (s *MakeCallRequest) SetOuterAccountType(v string) *MakeCallRequest {
	s.OuterAccountType = &v
	return s
}

func (s *MakeCallRequest) Validate() error {
	return dara.Validate(s)
}
