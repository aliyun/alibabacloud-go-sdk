// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneTwiceTelVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePhoneTwiceTelVerifyResponseBody
	GetCode() *string
	SetData(v *DescribePhoneTwiceTelVerifyResponseBodyData) *DescribePhoneTwiceTelVerifyResponseBody
	GetData() *DescribePhoneTwiceTelVerifyResponseBodyData
	SetMessage(v string) *DescribePhoneTwiceTelVerifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePhoneTwiceTelVerifyResponseBody
	GetRequestId() *string
}

type DescribePhoneTwiceTelVerifyResponseBody struct {
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **PortabilityNumberNotSupported**: The phone number that is involved in mobile number portability is not supported.
	//
	// 	- **RequestNumberNotSupported**: You are not allowed to query phone numbers assigned by China Broadnet (that is, phone numbers start with 192) and phone numbers assigned by virtual network operators (VNOs).
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// >  You are charged for phone number verifications if the value of Code is OK and the value of VerifyResult is not 0. For more information, see [Pricing](https://help.aliyun.com/document_detail/154751.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *DescribePhoneTwiceTelVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) GetData() *DescribePhoneTwiceTelVerifyResponseBodyData {
	return s.Data
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetCode(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetData(v *DescribePhoneTwiceTelVerifyResponseBodyData) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Data = v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetMessage(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) SetRequestId(v string) *DescribePhoneTwiceTelVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePhoneTwiceTelVerifyResponseBodyData struct {
	// The carrier. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  The returned result indicates the carrier who assigns the phone number. If the phone number involves mobile number portability, the carrier after mobile number portability is returned.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The result of the request. Valid values:
	//
	// 	- **0**: It is unable to judge whether the phone number is a reassigned number.
	//
	// 	- **1**: The phone number is a reassigned number.
	//
	// 	- **2**: The phone number is not a reassigned number.
	//
	// 	- **3**: The phone number has been canceled.
	//
	// example:
	//
	// 1
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) SetCarrier(v string) *DescribePhoneTwiceTelVerifyResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) SetVerifyResult(v string) *DescribePhoneTwiceTelVerifyResponseBodyData {
	s.VerifyResult = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
