// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeDoubleCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *MakeDoubleCallRequest
	GetAccountName() *string
	SetBizData(v string) *MakeDoubleCallRequest
	GetBizData() *string
	SetInstanceId(v string) *MakeDoubleCallRequest
	GetInstanceId() *string
	SetMemberPhone(v string) *MakeDoubleCallRequest
	GetMemberPhone() *string
	SetOutboundCallNumber(v string) *MakeDoubleCallRequest
	GetOutboundCallNumber() *string
	SetServicerPhone(v string) *MakeDoubleCallRequest
	GetServicerPhone() *string
}

type MakeDoubleCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// {"bizId": 123456}
	BizData *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1502123****
	MemberPhone *string `json:"MemberPhone,omitempty" xml:"MemberPhone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0571000****
	OutboundCallNumber *string `json:"OutboundCallNumber,omitempty" xml:"OutboundCallNumber,omitempty"`
	// example:
	//
	// 150****1234
	ServicerPhone *string `json:"ServicerPhone,omitempty" xml:"ServicerPhone,omitempty"`
}

func (s MakeDoubleCallRequest) String() string {
	return dara.Prettify(s)
}

func (s MakeDoubleCallRequest) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *MakeDoubleCallRequest) GetBizData() *string {
	return s.BizData
}

func (s *MakeDoubleCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *MakeDoubleCallRequest) GetMemberPhone() *string {
	return s.MemberPhone
}

func (s *MakeDoubleCallRequest) GetOutboundCallNumber() *string {
	return s.OutboundCallNumber
}

func (s *MakeDoubleCallRequest) GetServicerPhone() *string {
	return s.ServicerPhone
}

func (s *MakeDoubleCallRequest) SetAccountName(v string) *MakeDoubleCallRequest {
	s.AccountName = &v
	return s
}

func (s *MakeDoubleCallRequest) SetBizData(v string) *MakeDoubleCallRequest {
	s.BizData = &v
	return s
}

func (s *MakeDoubleCallRequest) SetInstanceId(v string) *MakeDoubleCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MakeDoubleCallRequest) SetMemberPhone(v string) *MakeDoubleCallRequest {
	s.MemberPhone = &v
	return s
}

func (s *MakeDoubleCallRequest) SetOutboundCallNumber(v string) *MakeDoubleCallRequest {
	s.OutboundCallNumber = &v
	return s
}

func (s *MakeDoubleCallRequest) SetServicerPhone(v string) *MakeDoubleCallRequest {
	s.ServicerPhone = &v
	return s
}

func (s *MakeDoubleCallRequest) Validate() error {
	return dara.Validate(s)
}
