// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDmAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountCode(v string) *BindDmAccountRequest
	GetAccountCode() *string
	SetCustSpaceId(v string) *BindDmAccountRequest
	GetCustSpaceId() *string
	SetExtendAttr(v *BindDmAccountRequestExtendAttr) *BindDmAccountRequest
	GetExtendAttr() *BindDmAccountRequestExtendAttr
	SetOwnerId(v int64) *BindDmAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *BindDmAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindDmAccountRequest
	GetResourceOwnerId() *int64
}

type BindDmAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	AccountCode *string `json:"AccountCode,omitempty" xml:"AccountCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	ExtendAttr           *BindDmAccountRequestExtendAttr `json:"ExtendAttr,omitempty" xml:"ExtendAttr,omitempty" type:"Struct"`
	OwnerId              *int64                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindDmAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDmAccountRequest) GoString() string {
	return s.String()
}

func (s *BindDmAccountRequest) GetAccountCode() *string {
	return s.AccountCode
}

func (s *BindDmAccountRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *BindDmAccountRequest) GetExtendAttr() *BindDmAccountRequestExtendAttr {
	return s.ExtendAttr
}

func (s *BindDmAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindDmAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindDmAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindDmAccountRequest) SetAccountCode(v string) *BindDmAccountRequest {
	s.AccountCode = &v
	return s
}

func (s *BindDmAccountRequest) SetCustSpaceId(v string) *BindDmAccountRequest {
	s.CustSpaceId = &v
	return s
}

func (s *BindDmAccountRequest) SetExtendAttr(v *BindDmAccountRequestExtendAttr) *BindDmAccountRequest {
	s.ExtendAttr = v
	return s
}

func (s *BindDmAccountRequest) SetOwnerId(v int64) *BindDmAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDmAccountRequest) SetResourceOwnerAccount(v string) *BindDmAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindDmAccountRequest) SetResourceOwnerId(v int64) *BindDmAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindDmAccountRequest) Validate() error {
	if s.ExtendAttr != nil {
		if err := s.ExtendAttr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindDmAccountRequestExtendAttr struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	SendType *string `json:"SendType,omitempty" xml:"SendType,omitempty"`
}

func (s BindDmAccountRequestExtendAttr) String() string {
	return dara.Prettify(s)
}

func (s BindDmAccountRequestExtendAttr) GoString() string {
	return s.String()
}

func (s *BindDmAccountRequestExtendAttr) GetAccountName() *string {
	return s.AccountName
}

func (s *BindDmAccountRequestExtendAttr) GetSendType() *string {
	return s.SendType
}

func (s *BindDmAccountRequestExtendAttr) SetAccountName(v string) *BindDmAccountRequestExtendAttr {
	s.AccountName = &v
	return s
}

func (s *BindDmAccountRequestExtendAttr) SetSendType(v string) *BindDmAccountRequestExtendAttr {
	s.SendType = &v
	return s
}

func (s *BindDmAccountRequestExtendAttr) Validate() error {
	return dara.Validate(s)
}
