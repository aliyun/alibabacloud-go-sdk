// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientBalanceDTO interface {
	dara.Model
	String() string
	GoString() string
	SetBalance(v float64) *ClientBalanceDTO
	GetBalance() *float64
	SetBalanceType(v string) *ClientBalanceDTO
	GetBalanceType() *string
	SetClientId(v int64) *ClientBalanceDTO
	GetClientId() *int64
	SetEnableBalance(v bool) *ClientBalanceDTO
	GetEnableBalance() *bool
	SetGmtCreate(v string) *ClientBalanceDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *ClientBalanceDTO
	GetGmtModified() *string
	SetId(v int64) *ClientBalanceDTO
	GetId() *int64
}

type ClientBalanceDTO struct {
	// example:
	//
	// 100.00
	Balance *float64 `json:"balance,omitempty" xml:"balance,omitempty"`
	// example:
	//
	// amount
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// true
	EnableBalance *bool `json:"enableBalance,omitempty" xml:"enableBalance,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ClientBalanceDTO) String() string {
	return dara.Prettify(s)
}

func (s ClientBalanceDTO) GoString() string {
	return s.String()
}

func (s *ClientBalanceDTO) GetBalance() *float64 {
	return s.Balance
}

func (s *ClientBalanceDTO) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ClientBalanceDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *ClientBalanceDTO) GetEnableBalance() *bool {
	return s.EnableBalance
}

func (s *ClientBalanceDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ClientBalanceDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ClientBalanceDTO) GetId() *int64 {
	return s.Id
}

func (s *ClientBalanceDTO) SetBalance(v float64) *ClientBalanceDTO {
	s.Balance = &v
	return s
}

func (s *ClientBalanceDTO) SetBalanceType(v string) *ClientBalanceDTO {
	s.BalanceType = &v
	return s
}

func (s *ClientBalanceDTO) SetClientId(v int64) *ClientBalanceDTO {
	s.ClientId = &v
	return s
}

func (s *ClientBalanceDTO) SetEnableBalance(v bool) *ClientBalanceDTO {
	s.EnableBalance = &v
	return s
}

func (s *ClientBalanceDTO) SetGmtCreate(v string) *ClientBalanceDTO {
	s.GmtCreate = &v
	return s
}

func (s *ClientBalanceDTO) SetGmtModified(v string) *ClientBalanceDTO {
	s.GmtModified = &v
	return s
}

func (s *ClientBalanceDTO) SetId(v int64) *ClientBalanceDTO {
	s.Id = &v
	return s
}

func (s *ClientBalanceDTO) Validate() error {
	return dara.Validate(s)
}
