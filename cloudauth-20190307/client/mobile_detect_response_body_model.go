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
	// Return code: 200 for success, others for failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 969434DF-926B-4997-9881-4DE94E39F805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
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
	return dara.Validate(s)
}

type MobileDetectResponseBodyResultObject struct {
	// Billing count, the total billing count in one request
	//
	// example:
	//
	// 2
	ChargeCount *string `json:"ChargeCount,omitempty" xml:"ChargeCount,omitempty"`
	// Verification results set
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
	return dara.Validate(s)
}

type MobileDetectResponseBodyResultObjectItems struct {
	// Phone number\\"s area (only for plaintext phone numbers)
	//
	// example:
	//
	// 安徽-马**
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// Verification result
	//
	// - 1: Available online
	//
	// - 2: Not available online
	//
	// - 3: No query result
	//
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// Operator name
	//
	// - CMCC: China Mobile
	//
	// - CUCC: China Unicom
	//
	// - CTCC: China Telecom
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// Phone number
	//
	// example:
	//
	// 131********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Verification details
	//
	// - 101: Available number
	//
	// - 102: Empty number
	//
	// - 103: Suspended
	//
	// - 104: Silent number (inactive small number, new number, non-smartphone user within the last six months)
	//
	// - 105: Risky number (long-term shutdown or no voice service activated and prone to complaints)
	//
	// - 301: No record found
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
