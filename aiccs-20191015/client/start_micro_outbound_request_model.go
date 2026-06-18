// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMicroOutboundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *StartMicroOutboundRequest
	GetAccountId() *string
	SetAccountType(v string) *StartMicroOutboundRequest
	GetAccountType() *string
	SetAppName(v string) *StartMicroOutboundRequest
	GetAppName() *string
	SetCalledNumber(v string) *StartMicroOutboundRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *StartMicroOutboundRequest
	GetCallingNumber() *string
	SetCommandCode(v string) *StartMicroOutboundRequest
	GetCommandCode() *string
	SetExtInfo(v string) *StartMicroOutboundRequest
	GetExtInfo() *string
	SetOwnerId(v int64) *StartMicroOutboundRequest
	GetOwnerId() *int64
	SetProdCode(v string) *StartMicroOutboundRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *StartMicroOutboundRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartMicroOutboundRequest
	GetResourceOwnerId() *int64
}

type StartMicroOutboundRequest struct {
	// Account ID.
	//
	// example:
	//
	// 223457****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Account type.
	//
	// example:
	//
	// BUC_TYPE
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Custom AppName for the business.
	//
	// example:
	//
	// aliyun
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Called number.
	//
	// example:
	//
	// 0571456****
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// Calling number.
	//
	// example:
	//
	// 1367123****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Instruction code.
	//
	// example:
	//
	// outBound_Call
	CommandCode *string `json:"CommandCode,omitempty" xml:"CommandCode,omitempty"`
	// Business information.
	//
	// example:
	//
	// {"caseId":23232****}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Product name. Default value: **aiccs**.
	//
	// example:
	//
	// aiccs
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartMicroOutboundRequest) String() string {
	return dara.Prettify(s)
}

func (s StartMicroOutboundRequest) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *StartMicroOutboundRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *StartMicroOutboundRequest) GetAppName() *string {
	return s.AppName
}

func (s *StartMicroOutboundRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *StartMicroOutboundRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *StartMicroOutboundRequest) GetCommandCode() *string {
	return s.CommandCode
}

func (s *StartMicroOutboundRequest) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *StartMicroOutboundRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartMicroOutboundRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *StartMicroOutboundRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartMicroOutboundRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartMicroOutboundRequest) SetAccountId(v string) *StartMicroOutboundRequest {
	s.AccountId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAccountType(v string) *StartMicroOutboundRequest {
	s.AccountType = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAppName(v string) *StartMicroOutboundRequest {
	s.AppName = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCalledNumber(v string) *StartMicroOutboundRequest {
	s.CalledNumber = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCallingNumber(v string) *StartMicroOutboundRequest {
	s.CallingNumber = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCommandCode(v string) *StartMicroOutboundRequest {
	s.CommandCode = &v
	return s
}

func (s *StartMicroOutboundRequest) SetExtInfo(v string) *StartMicroOutboundRequest {
	s.ExtInfo = &v
	return s
}

func (s *StartMicroOutboundRequest) SetOwnerId(v int64) *StartMicroOutboundRequest {
	s.OwnerId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetProdCode(v string) *StartMicroOutboundRequest {
	s.ProdCode = &v
	return s
}

func (s *StartMicroOutboundRequest) SetResourceOwnerAccount(v string) *StartMicroOutboundRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartMicroOutboundRequest) SetResourceOwnerId(v int64) *StartMicroOutboundRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartMicroOutboundRequest) Validate() error {
	return dara.Validate(s)
}
