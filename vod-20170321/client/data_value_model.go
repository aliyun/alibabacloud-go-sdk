// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataValue interface {
	dara.Model
	String() string
	GoString() string
	SetItemId(v string) *DataValue
	GetItemId() *string
	SetBusinessType(v string) *DataValue
	GetBusinessType() *string
	SetAppName(v string) *DataValue
	GetAppName() *string
	SetStatus(v string) *DataValue
	GetStatus() *string
	SetPurchaseStatus(v string) *DataValue
	GetPurchaseStatus() *string
	SetExpiredOn(v string) *DataValue
	GetExpiredOn() *string
}

type DataValue struct {
	ItemId         *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	BusinessType   *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PurchaseStatus *string `json:"PurchaseStatus,omitempty" xml:"PurchaseStatus,omitempty"`
	ExpiredOn      *string `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
}

func (s DataValue) String() string {
	return dara.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) GetItemId() *string {
	return s.ItemId
}

func (s *DataValue) GetBusinessType() *string {
	return s.BusinessType
}

func (s *DataValue) GetAppName() *string {
	return s.AppName
}

func (s *DataValue) GetStatus() *string {
	return s.Status
}

func (s *DataValue) GetPurchaseStatus() *string {
	return s.PurchaseStatus
}

func (s *DataValue) GetExpiredOn() *string {
	return s.ExpiredOn
}

func (s *DataValue) SetItemId(v string) *DataValue {
	s.ItemId = &v
	return s
}

func (s *DataValue) SetBusinessType(v string) *DataValue {
	s.BusinessType = &v
	return s
}

func (s *DataValue) SetAppName(v string) *DataValue {
	s.AppName = &v
	return s
}

func (s *DataValue) SetStatus(v string) *DataValue {
	s.Status = &v
	return s
}

func (s *DataValue) SetPurchaseStatus(v string) *DataValue {
	s.PurchaseStatus = &v
	return s
}

func (s *DataValue) SetExpiredOn(v string) *DataValue {
	s.ExpiredOn = &v
	return s
}

func (s *DataValue) Validate() error {
	return dara.Validate(s)
}
