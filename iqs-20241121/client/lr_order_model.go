// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLrOrder interface {
	dara.Model
	String() string
	GoString() string
	SetAccountCredentials(v []*AccountCredentialDTO) *LrOrder
	GetAccountCredentials() []*AccountCredentialDTO
	SetAliUid(v string) *LrOrder
	GetAliUid() *string
	SetCommodityCode(v string) *LrOrder
	GetCommodityCode() *string
	SetCommodityId(v string) *LrOrder
	GetCommodityId() *string
	SetCommoditySpec(v string) *LrOrder
	GetCommoditySpec() *string
	SetCustomerName(v string) *LrOrder
	GetCustomerName() *string
	SetEventTime(v string) *LrOrder
	GetEventTime() *string
	SetExpirationTime(v string) *LrOrder
	GetExpirationTime() *string
	SetGmtCreate(v string) *LrOrder
	GetGmtCreate() *string
	SetGmtModified(v string) *LrOrder
	GetGmtModified() *string
	SetInstanceId(v string) *LrOrder
	GetInstanceId() *string
	SetOrderCategory(v string) *LrOrder
	GetOrderCategory() *string
	SetOrderId(v int64) *LrOrder
	GetOrderId() *int64
	SetQps(v int32) *LrOrder
	GetQps() *int32
	SetStatus(v string) *LrOrder
	GetStatus() *string
}

type LrOrder struct {
	AccountCredentials []*AccountCredentialDTO `json:"accountCredentials,omitempty" xml:"accountCredentials,omitempty" type:"Repeated"`
	AliUid             *string                 `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	CommodityCode      *string                 `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityId        *string                 `json:"commodityId,omitempty" xml:"commodityId,omitempty"`
	CommoditySpec      *string                 `json:"commoditySpec,omitempty" xml:"commoditySpec,omitempty"`
	CustomerName       *string                 `json:"customerName,omitempty" xml:"customerName,omitempty"`
	EventTime          *string                 `json:"eventTime,omitempty" xml:"eventTime,omitempty"`
	ExpirationTime     *string                 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	GmtCreate          *string                 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified        *string                 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	InstanceId         *string                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrderCategory      *string                 `json:"orderCategory,omitempty" xml:"orderCategory,omitempty"`
	OrderId            *int64                  `json:"orderId,omitempty" xml:"orderId,omitempty"`
	Qps                *int32                  `json:"qps,omitempty" xml:"qps,omitempty"`
	Status             *string                 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s LrOrder) String() string {
	return dara.Prettify(s)
}

func (s LrOrder) GoString() string {
	return s.String()
}

func (s *LrOrder) GetAccountCredentials() []*AccountCredentialDTO {
	return s.AccountCredentials
}

func (s *LrOrder) GetAliUid() *string {
	return s.AliUid
}

func (s *LrOrder) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *LrOrder) GetCommodityId() *string {
	return s.CommodityId
}

func (s *LrOrder) GetCommoditySpec() *string {
	return s.CommoditySpec
}

func (s *LrOrder) GetCustomerName() *string {
	return s.CustomerName
}

func (s *LrOrder) GetEventTime() *string {
	return s.EventTime
}

func (s *LrOrder) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *LrOrder) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *LrOrder) GetGmtModified() *string {
	return s.GmtModified
}

func (s *LrOrder) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LrOrder) GetOrderCategory() *string {
	return s.OrderCategory
}

func (s *LrOrder) GetOrderId() *int64 {
	return s.OrderId
}

func (s *LrOrder) GetQps() *int32 {
	return s.Qps
}

func (s *LrOrder) GetStatus() *string {
	return s.Status
}

func (s *LrOrder) SetAccountCredentials(v []*AccountCredentialDTO) *LrOrder {
	s.AccountCredentials = v
	return s
}

func (s *LrOrder) SetAliUid(v string) *LrOrder {
	s.AliUid = &v
	return s
}

func (s *LrOrder) SetCommodityCode(v string) *LrOrder {
	s.CommodityCode = &v
	return s
}

func (s *LrOrder) SetCommodityId(v string) *LrOrder {
	s.CommodityId = &v
	return s
}

func (s *LrOrder) SetCommoditySpec(v string) *LrOrder {
	s.CommoditySpec = &v
	return s
}

func (s *LrOrder) SetCustomerName(v string) *LrOrder {
	s.CustomerName = &v
	return s
}

func (s *LrOrder) SetEventTime(v string) *LrOrder {
	s.EventTime = &v
	return s
}

func (s *LrOrder) SetExpirationTime(v string) *LrOrder {
	s.ExpirationTime = &v
	return s
}

func (s *LrOrder) SetGmtCreate(v string) *LrOrder {
	s.GmtCreate = &v
	return s
}

func (s *LrOrder) SetGmtModified(v string) *LrOrder {
	s.GmtModified = &v
	return s
}

func (s *LrOrder) SetInstanceId(v string) *LrOrder {
	s.InstanceId = &v
	return s
}

func (s *LrOrder) SetOrderCategory(v string) *LrOrder {
	s.OrderCategory = &v
	return s
}

func (s *LrOrder) SetOrderId(v int64) *LrOrder {
	s.OrderId = &v
	return s
}

func (s *LrOrder) SetQps(v int32) *LrOrder {
	s.Qps = &v
	return s
}

func (s *LrOrder) SetStatus(v string) *LrOrder {
	s.Status = &v
	return s
}

func (s *LrOrder) Validate() error {
	if s.AccountCredentials != nil {
		for _, item := range s.AccountCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
