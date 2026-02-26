// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralBill interface {
	dara.Model
	String() string
	GoString() string
	SetBillId(v string) *GeneralBill
	GetBillId() *string
	SetBillPeriod(v string) *GeneralBill
	GetBillPeriod() *string
	SetDownloadUrl(v []*string) *GeneralBill
	GetDownloadUrl() []*string
	SetEndTime(v string) *GeneralBill
	GetEndTime() *string
	SetGmtCreate(v string) *GeneralBill
	GetGmtCreate() *string
	SetGmtModified(v string) *GeneralBill
	GetGmtModified() *string
	SetShopId(v string) *GeneralBill
	GetShopId() *string
	SetShopName(v string) *GeneralBill
	GetShopName() *string
	SetStartTime(v string) *GeneralBill
	GetStartTime() *string
	SetTotalAmount(v *Money) *GeneralBill
	GetTotalAmount() *Money
}

type GeneralBill struct {
	BillId      *string   `json:"billId,omitempty" xml:"billId,omitempty"`
	BillPeriod  *string   `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	DownloadUrl []*string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty" type:"Repeated"`
	EndTime     *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	GmtCreate   *string   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	ShopId      *string   `json:"shopId,omitempty" xml:"shopId,omitempty"`
	ShopName    *string   `json:"shopName,omitempty" xml:"shopName,omitempty"`
	StartTime   *string   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TotalAmount *Money    `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s GeneralBill) String() string {
	return dara.Prettify(s)
}

func (s GeneralBill) GoString() string {
	return s.String()
}

func (s *GeneralBill) GetBillId() *string {
	return s.BillId
}

func (s *GeneralBill) GetBillPeriod() *string {
	return s.BillPeriod
}

func (s *GeneralBill) GetDownloadUrl() []*string {
	return s.DownloadUrl
}

func (s *GeneralBill) GetEndTime() *string {
	return s.EndTime
}

func (s *GeneralBill) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GeneralBill) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GeneralBill) GetShopId() *string {
	return s.ShopId
}

func (s *GeneralBill) GetShopName() *string {
	return s.ShopName
}

func (s *GeneralBill) GetStartTime() *string {
	return s.StartTime
}

func (s *GeneralBill) GetTotalAmount() *Money {
	return s.TotalAmount
}

func (s *GeneralBill) SetBillId(v string) *GeneralBill {
	s.BillId = &v
	return s
}

func (s *GeneralBill) SetBillPeriod(v string) *GeneralBill {
	s.BillPeriod = &v
	return s
}

func (s *GeneralBill) SetDownloadUrl(v []*string) *GeneralBill {
	s.DownloadUrl = v
	return s
}

func (s *GeneralBill) SetEndTime(v string) *GeneralBill {
	s.EndTime = &v
	return s
}

func (s *GeneralBill) SetGmtCreate(v string) *GeneralBill {
	s.GmtCreate = &v
	return s
}

func (s *GeneralBill) SetGmtModified(v string) *GeneralBill {
	s.GmtModified = &v
	return s
}

func (s *GeneralBill) SetShopId(v string) *GeneralBill {
	s.ShopId = &v
	return s
}

func (s *GeneralBill) SetShopName(v string) *GeneralBill {
	s.ShopName = &v
	return s
}

func (s *GeneralBill) SetStartTime(v string) *GeneralBill {
	s.StartTime = &v
	return s
}

func (s *GeneralBill) SetTotalAmount(v *Money) *GeneralBill {
	s.TotalAmount = v
	return s
}

func (s *GeneralBill) Validate() error {
	if s.TotalAmount != nil {
		if err := s.TotalAmount.Validate(); err != nil {
			return err
		}
	}
	return nil
}
