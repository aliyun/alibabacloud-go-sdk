// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIssueCouponForCustomerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IssueCouponForCustomerResponseBody
	GetCode() *string
	SetMessage(v string) *IssueCouponForCustomerResponseBody
	GetMessage() *string
	SetRequestId(v string) *IssueCouponForCustomerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IssueCouponForCustomerResponseBody
	GetSuccess() *bool
	SetData(v *IssueCouponForCustomerResponseBodyData) *IssueCouponForCustomerResponseBody
	GetData() *IssueCouponForCustomerResponseBodyData
}

type IssueCouponForCustomerResponseBody struct {
	// example:
	//
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data    *IssueCouponForCustomerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s IssueCouponForCustomerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IssueCouponForCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerResponseBody) GetCode() *string {
	return s.Code
}

func (s *IssueCouponForCustomerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IssueCouponForCustomerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IssueCouponForCustomerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IssueCouponForCustomerResponseBody) GetData() *IssueCouponForCustomerResponseBodyData {
	return s.Data
}

func (s *IssueCouponForCustomerResponseBody) SetCode(v string) *IssueCouponForCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetMessage(v string) *IssueCouponForCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetRequestId(v string) *IssueCouponForCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetSuccess(v bool) *IssueCouponForCustomerResponseBody {
	s.Success = &v
	return s
}

func (s *IssueCouponForCustomerResponseBody) SetData(v *IssueCouponForCustomerResponseBodyData) *IssueCouponForCustomerResponseBody {
	s.Data = v
	return s
}

func (s *IssueCouponForCustomerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IssueCouponForCustomerResponseBodyData struct {
	// example:
	//
	// 5075915
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// example:
	//
	// 2024-03-05 18:24:07
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 111,2222
	Uidlist *string `json:"Uidlist,omitempty" xml:"Uidlist,omitempty"`
}

func (s IssueCouponForCustomerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s IssueCouponForCustomerResponseBodyData) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerResponseBodyData) GetCouponTemplateId() *int64 {
	return s.CouponTemplateId
}

func (s *IssueCouponForCustomerResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *IssueCouponForCustomerResponseBodyData) GetUidlist() *string {
	return s.Uidlist
}

func (s *IssueCouponForCustomerResponseBodyData) SetCouponTemplateId(v int64) *IssueCouponForCustomerResponseBodyData {
	s.CouponTemplateId = &v
	return s
}

func (s *IssueCouponForCustomerResponseBodyData) SetCreateTime(v string) *IssueCouponForCustomerResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *IssueCouponForCustomerResponseBodyData) SetUidlist(v string) *IssueCouponForCustomerResponseBodyData {
	s.Uidlist = &v
	return s
}

func (s *IssueCouponForCustomerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
