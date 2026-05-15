// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ListOutboundPhoneNumberRequest
	GetAccountName() *string
	SetClientToken(v string) *ListOutboundPhoneNumberRequest
	GetClientToken() *string
	SetInstanceId(v string) *ListOutboundPhoneNumberRequest
	GetInstanceId() *string
}

type ListOutboundPhoneNumberRequest struct {
	// example:
	//
	// 123@****.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListOutboundPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ListOutboundPhoneNumberRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListOutboundPhoneNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOutboundPhoneNumberRequest) SetAccountName(v string) *ListOutboundPhoneNumberRequest {
	s.AccountName = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetClientToken(v string) *ListOutboundPhoneNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetInstanceId(v string) *ListOutboundPhoneNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
