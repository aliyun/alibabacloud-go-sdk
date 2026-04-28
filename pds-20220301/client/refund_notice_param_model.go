// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundNoticeParam interface {
	dara.Model
	String() string
	GoString() string
	SetAliuid(v int64) *RefundNoticeParam
	GetAliuid() *int64
	SetAliyunProduceCode(v string) *RefundNoticeParam
	GetAliyunProduceCode() *string
	SetCommodityCode(v string) *RefundNoticeParam
	GetCommodityCode() *string
	SetInstanceId(v string) *RefundNoticeParam
	GetInstanceId() *string
	SetNewExpireTime(v interface{}) *RefundNoticeParam
	GetNewExpireTime() interface{}
	SetOrderIds(v []*int64) *RefundNoticeParam
	GetOrderIds() []*int64
	SetRefundParamMap(v map[string]*string) *RefundNoticeParam
	GetRefundParamMap() map[string]*string
	SetRefundType(v string) *RefundNoticeParam
	GetRefundType() *string
}

type RefundNoticeParam struct {
	Aliuid            *int64             `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	AliyunProduceCode *string            `json:"aliyunProduceCode,omitempty" xml:"aliyunProduceCode,omitempty"`
	CommodityCode     *string            `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	InstanceId        *string            `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NewExpireTime     interface{}        `json:"newExpireTime,omitempty" xml:"newExpireTime,omitempty"`
	OrderIds          []*int64           `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	RefundParamMap    map[string]*string `json:"refundParamMap,omitempty" xml:"refundParamMap,omitempty"`
	RefundType        *string            `json:"refundType,omitempty" xml:"refundType,omitempty"`
}

func (s RefundNoticeParam) String() string {
	return dara.Prettify(s)
}

func (s RefundNoticeParam) GoString() string {
	return s.String()
}

func (s *RefundNoticeParam) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *RefundNoticeParam) GetAliyunProduceCode() *string {
	return s.AliyunProduceCode
}

func (s *RefundNoticeParam) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *RefundNoticeParam) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RefundNoticeParam) GetNewExpireTime() interface{} {
	return s.NewExpireTime
}

func (s *RefundNoticeParam) GetOrderIds() []*int64 {
	return s.OrderIds
}

func (s *RefundNoticeParam) GetRefundParamMap() map[string]*string {
	return s.RefundParamMap
}

func (s *RefundNoticeParam) GetRefundType() *string {
	return s.RefundType
}

func (s *RefundNoticeParam) SetAliuid(v int64) *RefundNoticeParam {
	s.Aliuid = &v
	return s
}

func (s *RefundNoticeParam) SetAliyunProduceCode(v string) *RefundNoticeParam {
	s.AliyunProduceCode = &v
	return s
}

func (s *RefundNoticeParam) SetCommodityCode(v string) *RefundNoticeParam {
	s.CommodityCode = &v
	return s
}

func (s *RefundNoticeParam) SetInstanceId(v string) *RefundNoticeParam {
	s.InstanceId = &v
	return s
}

func (s *RefundNoticeParam) SetNewExpireTime(v interface{}) *RefundNoticeParam {
	s.NewExpireTime = v
	return s
}

func (s *RefundNoticeParam) SetOrderIds(v []*int64) *RefundNoticeParam {
	s.OrderIds = v
	return s
}

func (s *RefundNoticeParam) SetRefundParamMap(v map[string]*string) *RefundNoticeParam {
	s.RefundParamMap = v
	return s
}

func (s *RefundNoticeParam) SetRefundType(v string) *RefundNoticeParam {
	s.RefundType = &v
	return s
}

func (s *RefundNoticeParam) Validate() error {
	return dara.Validate(s)
}
