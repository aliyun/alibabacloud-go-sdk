// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillDownOrUpgradeCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *BillDownOrUpgradeCmd
	GetAccountId() *string
	SetActiveDate(v string) *BillDownOrUpgradeCmd
	GetActiveDate() *string
	SetCodeType(v string) *BillDownOrUpgradeCmd
	GetCodeType() *string
	SetOperateTypEnum(v string) *BillDownOrUpgradeCmd
	GetOperateTypEnum() *string
	SetQps(v int32) *BillDownOrUpgradeCmd
	GetQps() *int32
}

type BillDownOrUpgradeCmd struct {
	AccountId      *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ActiveDate     *string `json:"activeDate,omitempty" xml:"activeDate,omitempty"`
	CodeType       *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	OperateTypEnum *string `json:"operateTypEnum,omitempty" xml:"operateTypEnum,omitempty"`
	Qps            *int32  `json:"qps,omitempty" xml:"qps,omitempty"`
}

func (s BillDownOrUpgradeCmd) String() string {
	return dara.Prettify(s)
}

func (s BillDownOrUpgradeCmd) GoString() string {
	return s.String()
}

func (s *BillDownOrUpgradeCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *BillDownOrUpgradeCmd) GetActiveDate() *string {
	return s.ActiveDate
}

func (s *BillDownOrUpgradeCmd) GetCodeType() *string {
	return s.CodeType
}

func (s *BillDownOrUpgradeCmd) GetOperateTypEnum() *string {
	return s.OperateTypEnum
}

func (s *BillDownOrUpgradeCmd) GetQps() *int32 {
	return s.Qps
}

func (s *BillDownOrUpgradeCmd) SetAccountId(v string) *BillDownOrUpgradeCmd {
	s.AccountId = &v
	return s
}

func (s *BillDownOrUpgradeCmd) SetActiveDate(v string) *BillDownOrUpgradeCmd {
	s.ActiveDate = &v
	return s
}

func (s *BillDownOrUpgradeCmd) SetCodeType(v string) *BillDownOrUpgradeCmd {
	s.CodeType = &v
	return s
}

func (s *BillDownOrUpgradeCmd) SetOperateTypEnum(v string) *BillDownOrUpgradeCmd {
	s.OperateTypEnum = &v
	return s
}

func (s *BillDownOrUpgradeCmd) SetQps(v int32) *BillDownOrUpgradeCmd {
	s.Qps = &v
	return s
}

func (s *BillDownOrUpgradeCmd) Validate() error {
	return dara.Validate(s)
}
