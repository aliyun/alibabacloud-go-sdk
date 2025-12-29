// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOnlineTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePhoneNumberOnlineTimeResponseBody
	GetCode() *string
	SetData(v *DescribePhoneNumberOnlineTimeResponseBodyData) *DescribePhoneNumberOnlineTimeResponseBody
	GetData() *DescribePhoneNumberOnlineTimeResponseBodyData
	SetMessage(v string) *DescribePhoneNumberOnlineTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneNumberOnlineTimeResponseBody
	GetRequestId() *string
}

type DescribePhoneNumberOnlineTimeResponseBody struct {
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **PortabilityNumberNotSupported**: The phone number that is involved in mobile number portability is not supported.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// >  You are charged if the value of Code is OK and the value of VerifyResult is not -1. For more information, see [Pricing](https://help.aliyun.com/document_detail/154751.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribePhoneNumberOnlineTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) GetData() *DescribePhoneNumberOnlineTimeResponseBodyData {
	return s.Data
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetCode(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetData(v *DescribePhoneNumberOnlineTimeResponseBodyData) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetMessage(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) SetRequestId(v string) *DescribePhoneNumberOnlineTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneNumberOnlineTimeResponseBodyData struct {
	// The carrier code. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// 	- **CBN**: China Broadnet
	//
	// example:
	//
	// CMCC
	CarrierCode *string `json:"CarrierCode,omitempty" xml:"CarrierCode,omitempty"`
	// The enumerated value of the usage period of a phone number. Valid values:
	//
	// 	- **-1**: No usage period information is available for the phone number.
	//
	// 	- **0**: The phone number status is abnormal. For example, the phone number is a nonexistent number.
	//
	// 	- **1*	- :[0-3) months.
	//
	// 	- **2*	- :[3-6] months.
	//
	// 	- **3*	- :(6-12] months.
	//
	// 	- **4*	- :(12-24] months.
	//
	// 	- **5*	- :(24,+) months.
	//
	// example:
	//
	// 1
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneNumberOnlineTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) GetCarrierCode() *string {
	return s.CarrierCode
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) SetCarrierCode(v string) *DescribePhoneNumberOnlineTimeResponseBodyData {
	s.CarrierCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) SetVerifyResult(v string) *DescribePhoneNumberOnlineTimeResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
