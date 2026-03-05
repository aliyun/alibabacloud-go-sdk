// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCheckResourceMeasureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBelongId(v string) *BatchCheckResourceMeasureRequest
	GetBelongId() *string
	SetBelongIdType(v string) *BatchCheckResourceMeasureRequest
	GetBelongIdType() *string
	SetBizType(v string) *BatchCheckResourceMeasureRequest
	GetBizType() *string
	SetEspBizId(v string) *BatchCheckResourceMeasureRequest
	GetEspBizId() *string
	SetOrderComponentParams(v string) *BatchCheckResourceMeasureRequest
	GetOrderComponentParams() *string
	SetResourceCheckItems(v string) *BatchCheckResourceMeasureRequest
	GetResourceCheckItems() *string
}

type BatchCheckResourceMeasureRequest struct {
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// example:
	//
	// USER
	BelongIdType *string `json:"BelongIdType,omitempty" xml:"BelongIdType,omitempty"`
	// example:
	//
	// 4
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// p20202933455
	EspBizId *string `json:"EspBizId,omitempty" xml:"EspBizId,omitempty"`
	// example:
	//
	// {"siteversion":"test"}
	OrderComponentParams *string `json:"OrderComponentParams,omitempty" xml:"OrderComponentParams,omitempty"`
	// example:
	//
	// [
	//
	//                   {
	//
	//                     "resourceCode": "InspirationTokens",
	//
	//                     "resourceValue": "0"
	//
	//                   }
	//
	//                 ]
	ResourceCheckItems *string `json:"ResourceCheckItems,omitempty" xml:"ResourceCheckItems,omitempty"`
}

func (s BatchCheckResourceMeasureRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCheckResourceMeasureRequest) GoString() string {
	return s.String()
}

func (s *BatchCheckResourceMeasureRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *BatchCheckResourceMeasureRequest) GetBelongIdType() *string {
	return s.BelongIdType
}

func (s *BatchCheckResourceMeasureRequest) GetBizType() *string {
	return s.BizType
}

func (s *BatchCheckResourceMeasureRequest) GetEspBizId() *string {
	return s.EspBizId
}

func (s *BatchCheckResourceMeasureRequest) GetOrderComponentParams() *string {
	return s.OrderComponentParams
}

func (s *BatchCheckResourceMeasureRequest) GetResourceCheckItems() *string {
	return s.ResourceCheckItems
}

func (s *BatchCheckResourceMeasureRequest) SetBelongId(v string) *BatchCheckResourceMeasureRequest {
	s.BelongId = &v
	return s
}

func (s *BatchCheckResourceMeasureRequest) SetBelongIdType(v string) *BatchCheckResourceMeasureRequest {
	s.BelongIdType = &v
	return s
}

func (s *BatchCheckResourceMeasureRequest) SetBizType(v string) *BatchCheckResourceMeasureRequest {
	s.BizType = &v
	return s
}

func (s *BatchCheckResourceMeasureRequest) SetEspBizId(v string) *BatchCheckResourceMeasureRequest {
	s.EspBizId = &v
	return s
}

func (s *BatchCheckResourceMeasureRequest) SetOrderComponentParams(v string) *BatchCheckResourceMeasureRequest {
	s.OrderComponentParams = &v
	return s
}

func (s *BatchCheckResourceMeasureRequest) SetResourceCheckItems(v string) *BatchCheckResourceMeasureRequest {
	s.ResourceCheckItems = &v
	return s
}

func (s *BatchCheckResourceMeasureRequest) Validate() error {
	return dara.Validate(s)
}
