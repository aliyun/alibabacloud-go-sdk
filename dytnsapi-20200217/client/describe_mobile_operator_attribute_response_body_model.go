// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileOperatorAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeMobileOperatorAttributeResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DescribeMobileOperatorAttributeResponseBody
	GetCode() *string
	SetData(v *DescribeMobileOperatorAttributeResponseBodyData) *DescribeMobileOperatorAttributeResponseBody
	GetData() *DescribeMobileOperatorAttributeResponseBodyData
	SetMessage(v string) *DescribeMobileOperatorAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMobileOperatorAttributeResponseBody
	GetRequestId() *string
}

type DescribeMobileOperatorAttributeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeMobileOperatorAttributeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMobileOperatorAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileOperatorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeMobileOperatorAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMobileOperatorAttributeResponseBody) GetData() *DescribeMobileOperatorAttributeResponseBodyData {
	return s.Data
}

func (s *DescribeMobileOperatorAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMobileOperatorAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetAccessDeniedDetail(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetCode(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetData(v *DescribeMobileOperatorAttributeResponseBodyData) *DescribeMobileOperatorAttributeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetMessage(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) SetRequestId(v string) *DescribeMobileOperatorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMobileOperatorAttributeResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// example:
	//
	// 示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// true
	IsNumberPortability *bool `json:"IsNumberPortability,omitempty" xml:"IsNumberPortability,omitempty"`
	// example:
	//
	// 示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// 示例值示例值
	RealNumber *string `json:"RealNumber,omitempty" xml:"RealNumber,omitempty"`
	// example:
	//
	// 示例值示例值
	SegmentCarrier *string `json:"SegmentCarrier,omitempty" xml:"SegmentCarrier,omitempty"`
}

func (s DescribeMobileOperatorAttributeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileOperatorAttributeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) GetCity() *string {
	return s.City
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) GetIsNumberPortability() *bool {
	return s.IsNumberPortability
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) GetProvince() *string {
	return s.Province
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) GetRealNumber() *string {
	return s.RealNumber
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) GetSegmentCarrier() *string {
	return s.SegmentCarrier
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetBasicCarrier(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetCarrier(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetCity(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.City = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetIsNumberPortability(v bool) *DescribeMobileOperatorAttributeResponseBodyData {
	s.IsNumberPortability = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetProvince(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.Province = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetRealNumber(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.RealNumber = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) SetSegmentCarrier(v string) *DescribeMobileOperatorAttributeResponseBodyData {
	s.SegmentCarrier = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
