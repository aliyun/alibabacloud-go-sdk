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
	// The request status code. Valid values:
	//
	// - **OK**: The request was successful.
	//
	// - **PortabilityNumberNotSupported**: Queries for this ported number are not supported.
	//
	// - **RequestNumberNotSupported**: Queries are not supported for numbers from China Broadnet (starting with 192), mobile virtual network operators, and other unsupported carriers.
	//
	// - **RequestFrequencyLimit**: Carriers limit frequent queries for the same number. If you receive this error code, try again later.
	//
	// > A charge applies when the value of `Code` is `OK` and the value of `VerifyResult` is not `0`. For more information, see [Phone Number Service pricing](https://help.aliyun.com/document_detail/154751.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A data structure that contains the verification results.
	Data *DescribePhoneTwiceTelVerifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// A description of the returned status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. This common parameter is returned with each request. Use this ID to troubleshoot issues.
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
	// The carrier that provides service for the number. Valid values:
	//
	// - **CMCC**: China Mobile.
	//
	// - **CUCC**: China Unicom.
	//
	// - **CTCC**: China Telecom.
	//
	// > The carrier that currently provides service for the number. For a ported number, this is the destination carrier.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The verification result. Valid values:
	//
	// - **0**: Cannot be determined.
	//
	// - **1**: The number is a recycled number.
	//
	// - **2**: The number is not a recycled number.
	//
	// - **3**: The number has been deactivated.
	//
	// - **4**: Unknown: The number was transferred to a new owner.
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
