// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountExistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIds(v []*CheckAccountExistRequestEcIdAccountIds) *CheckAccountExistRequest
	GetEcIdAccountIds() []*CheckAccountExistRequestEcIdAccountIds
	SetNbid(v string) *CheckAccountExistRequest
	GetNbid() *string
	SetToUserType(v int32) *CheckAccountExistRequest
	GetToUserType() *int32
	SetTransferAccount(v string) *CheckAccountExistRequest
	GetTransferAccount() *string
}

type CheckAccountExistRequest struct {
	EcIdAccountIds []*CheckAccountExistRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 0
	ToUserType *int32 `json:"ToUserType,omitempty" xml:"ToUserType,omitempty"`
	// example:
	//
	// 12323
	TransferAccount *string `json:"TransferAccount,omitempty" xml:"TransferAccount,omitempty"`
}

func (s CheckAccountExistRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountExistRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountExistRequest) GetEcIdAccountIds() []*CheckAccountExistRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *CheckAccountExistRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CheckAccountExistRequest) GetToUserType() *int32 {
	return s.ToUserType
}

func (s *CheckAccountExistRequest) GetTransferAccount() *string {
	return s.TransferAccount
}

func (s *CheckAccountExistRequest) SetEcIdAccountIds(v []*CheckAccountExistRequestEcIdAccountIds) *CheckAccountExistRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *CheckAccountExistRequest) SetNbid(v string) *CheckAccountExistRequest {
	s.Nbid = &v
	return s
}

func (s *CheckAccountExistRequest) SetToUserType(v int32) *CheckAccountExistRequest {
	s.ToUserType = &v
	return s
}

func (s *CheckAccountExistRequest) SetTransferAccount(v string) *CheckAccountExistRequest {
	s.TransferAccount = &v
	return s
}

func (s *CheckAccountExistRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckAccountExistRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s CheckAccountExistRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountExistRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *CheckAccountExistRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *CheckAccountExistRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *CheckAccountExistRequestEcIdAccountIds) SetAccountIds(v []*int64) *CheckAccountExistRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *CheckAccountExistRequestEcIdAccountIds) SetEcId(v string) *CheckAccountExistRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *CheckAccountExistRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
