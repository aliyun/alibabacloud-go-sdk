// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoupondeductProductCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCoupondeductProductCodeResponseBody
	GetCode() *string
	SetData(v []*GetCoupondeductProductCodeResponseBodyData) *GetCoupondeductProductCodeResponseBody
	GetData() []*GetCoupondeductProductCodeResponseBodyData
	SetMessage(v string) *GetCoupondeductProductCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCoupondeductProductCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCoupondeductProductCodeResponseBody
	GetSuccess() *bool
}

type GetCoupondeductProductCodeResponseBody struct {
	// example:
	//
	// code
	Code    *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*GetCoupondeductProductCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 210e876f16704666020714468dab35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCoupondeductProductCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCoupondeductProductCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCoupondeductProductCodeResponseBody) GetData() []*GetCoupondeductProductCodeResponseBodyData {
	return s.Data
}

func (s *GetCoupondeductProductCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCoupondeductProductCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCoupondeductProductCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCoupondeductProductCodeResponseBody) SetCode(v string) *GetCoupondeductProductCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetData(v []*GetCoupondeductProductCodeResponseBodyData) *GetCoupondeductProductCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetMessage(v string) *GetCoupondeductProductCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetRequestId(v string) *GetCoupondeductProductCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) SetSuccess(v bool) *GetCoupondeductProductCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetCoupondeductProductCodeResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCoupondeductProductCodeResponseBodyData struct {
	// example:
	//
	// code1
	ProductType interface{} `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetCoupondeductProductCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCoupondeductProductCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeResponseBodyData) GetProductType() interface{} {
	return s.ProductType
}

func (s *GetCoupondeductProductCodeResponseBodyData) SetProductType(v interface{}) *GetCoupondeductProductCodeResponseBodyData {
	s.ProductType = v
	return s
}

func (s *GetCoupondeductProductCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
