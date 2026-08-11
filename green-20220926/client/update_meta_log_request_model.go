// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *UpdateMetaLogRequest
	GetCommodityCode() *string
	SetDeliveryRegion(v string) *UpdateMetaLogRequest
	GetDeliveryRegion() *string
	SetStorage(v int64) *UpdateMetaLogRequest
	GetStorage() *int64
	SetTtl(v int32) *UpdateMetaLogRequest
	GetTtl() *int32
}

type UpdateMetaLogRequest struct {
	// The commodity code.
	//
	// example:
	//
	// lvwang_guardrail_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The delivery region.
	//
	// example:
	//
	// cn-beijing
	DeliveryRegion *string `json:"DeliveryRegion,omitempty" xml:"DeliveryRegion,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 1
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The time to live.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateMetaLogRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaLogRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaLogRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *UpdateMetaLogRequest) GetDeliveryRegion() *string {
	return s.DeliveryRegion
}

func (s *UpdateMetaLogRequest) GetStorage() *int64 {
	return s.Storage
}

func (s *UpdateMetaLogRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateMetaLogRequest) SetCommodityCode(v string) *UpdateMetaLogRequest {
	s.CommodityCode = &v
	return s
}

func (s *UpdateMetaLogRequest) SetDeliveryRegion(v string) *UpdateMetaLogRequest {
	s.DeliveryRegion = &v
	return s
}

func (s *UpdateMetaLogRequest) SetStorage(v int64) *UpdateMetaLogRequest {
	s.Storage = &v
	return s
}

func (s *UpdateMetaLogRequest) SetTtl(v int32) *UpdateMetaLogRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateMetaLogRequest) Validate() error {
	return dara.Validate(s)
}
