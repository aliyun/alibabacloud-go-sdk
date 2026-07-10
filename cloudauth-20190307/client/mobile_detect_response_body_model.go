// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobileDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MobileDetectResponseBody
	GetCode() *string
	SetMessage(v string) *MobileDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *MobileDetectResponseBody
	GetRequestId() *string
	SetResultObject(v *MobileDetectResponseBodyResultObject) *MobileDetectResponseBody
	GetResultObject() *MobileDetectResponseBodyResultObject
}

type MobileDetectResponseBody struct {
	// The return code. A value of 200 indicates success. Other values indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 969434DF-926B-4997-9881-4DE94E39F805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information.
	ResultObject *MobileDetectResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
}

func (s MobileDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MobileDetectResponseBody) GoString() string {
	return s.String()
}

func (s *MobileDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *MobileDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MobileDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MobileDetectResponseBody) GetResultObject() *MobileDetectResponseBodyResultObject {
	return s.ResultObject
}

func (s *MobileDetectResponseBody) SetCode(v string) *MobileDetectResponseBody {
	s.Code = &v
	return s
}

func (s *MobileDetectResponseBody) SetMessage(v string) *MobileDetectResponseBody {
	s.Message = &v
	return s
}

func (s *MobileDetectResponseBody) SetRequestId(v string) *MobileDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *MobileDetectResponseBody) SetResultObject(v *MobileDetectResponseBodyResultObject) *MobileDetectResponseBody {
	s.ResultObject = v
	return s
}

func (s *MobileDetectResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MobileDetectResponseBodyResultObject struct {
	// The total number of billing counts in a single request.
	//
	// example:
	//
	// 2
	ChargeCount *string `json:"ChargeCount,omitempty" xml:"ChargeCount,omitempty"`
	// The verification result set.
	Items []*MobileDetectResponseBodyResultObjectItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s MobileDetectResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s MobileDetectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *MobileDetectResponseBodyResultObject) GetChargeCount() *string {
	return s.ChargeCount
}

func (s *MobileDetectResponseBodyResultObject) GetItems() []*MobileDetectResponseBodyResultObjectItems {
	return s.Items
}

func (s *MobileDetectResponseBodyResultObject) SetChargeCount(v string) *MobileDetectResponseBodyResultObject {
	s.ChargeCount = &v
	return s
}

func (s *MobileDetectResponseBodyResultObject) SetItems(v []*MobileDetectResponseBodyResultObjectItems) *MobileDetectResponseBodyResultObject {
	s.Items = v
	return s
}

func (s *MobileDetectResponseBodyResultObject) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MobileDetectResponseBodyResultObjectItems struct {
	// The location to which the phone number belongs. This field is available only for plaintext phone numbers.
	//
	// example:
	//
	// 安徽-马**
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The verification result. Valid values:
	//
	// - 1: Active and available.
	//
	// - 2: Not in active and available status.
	//
	// - 3: No query results.
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// The carrier name. Valid values:
	//
	// - CMCC: China Mobile
	//
	// - CUCC: China Unicom
	//
	// - CTCC: China Telecom.
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 131********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The verification details. Valid values:
	//
	// - 101: Available number.
	//
	// - 102: Empty number.
	//
	// - 103: Suspended.
	//
	// - 104: Silent number (inactive secondary number in the past six months, new number, or non-smartphone user).
	//
	// - 105: Risky number (user with prolonged shutdown, voice service not activated, or prone to complaints).
	//
	// - 301: No record found.
	//
	// example:
	//
	// 101
	SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s MobileDetectResponseBodyResultObjectItems) String() string {
	return dara.Prettify(s)
}

func (s MobileDetectResponseBodyResultObjectItems) GoString() string {
	return s.String()
}

func (s *MobileDetectResponseBodyResultObjectItems) GetArea() *string {
	return s.Area
}

func (s *MobileDetectResponseBodyResultObjectItems) GetBizCode() *string {
	return s.BizCode
}

func (s *MobileDetectResponseBodyResultObjectItems) GetIspName() *string {
	return s.IspName
}

func (s *MobileDetectResponseBodyResultObjectItems) GetMobile() *string {
	return s.Mobile
}

func (s *MobileDetectResponseBodyResultObjectItems) GetSubCode() *string {
	return s.SubCode
}

func (s *MobileDetectResponseBodyResultObjectItems) SetArea(v string) *MobileDetectResponseBodyResultObjectItems {
	s.Area = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetBizCode(v string) *MobileDetectResponseBodyResultObjectItems {
	s.BizCode = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetIspName(v string) *MobileDetectResponseBodyResultObjectItems {
	s.IspName = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetMobile(v string) *MobileDetectResponseBodyResultObjectItems {
	s.Mobile = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) SetSubCode(v string) *MobileDetectResponseBodyResultObjectItems {
	s.SubCode = &v
	return s
}

func (s *MobileDetectResponseBodyResultObjectItems) Validate() error {
	return dara.Validate(s)
}
