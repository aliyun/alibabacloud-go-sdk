// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTwoElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TwoElementsVerificationResponseBody
	GetCode() *string
	SetData(v *TwoElementsVerificationResponseBodyData) *TwoElementsVerificationResponseBody
	GetData() *TwoElementsVerificationResponseBodyData
	SetMessage(v string) *TwoElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *TwoElementsVerificationResponseBody
	GetRequestId() *string
}

type TwoElementsVerificationResponseBody struct {
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- For more information, see Error codes in this documentation.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number or name at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *TwoElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TwoElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TwoElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *TwoElementsVerificationResponseBody) GetData() *TwoElementsVerificationResponseBodyData {
	return s.Data
}

func (s *TwoElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TwoElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TwoElementsVerificationResponseBody) SetCode(v string) *TwoElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetData(v *TwoElementsVerificationResponseBodyData) *TwoElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetMessage(v string) *TwoElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) SetRequestId(v string) *TwoElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *TwoElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TwoElementsVerificationResponseBodyData struct {
	// The basic carriers. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	//
	// >  You are not allowed to verify numbers assigned by China Broadnet.
	//
	// example:
	//
	// China Mobile
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// Indicates whether the specified name and phone number belong to the same user. Valid values:
	//
	// 	- **1**: The specified name and phone number belong to the same user.
	//
	// 	- **0**: The specified name and phone number do not belong to the same user.
	//
	// 	- **2**: The specified name and phone number cannot be found.
	//
	// The phone number registration data of a user is usually updated one or three days after registration. The registration data can be queried only after the update. The following table shows the verification results under different phone number states.
	//
	// |Carrier/Phone number state|Out-of-service|Nonexistent|Canceled|
	//
	// |---|---|---|---|
	//
	// |China Mobile|Verifications can be carried out normally.|The specified name and phone number cannot be found.|The specified name and phone number cannot be found.|
	//
	// |China Unicom|Verifications can be carried out normally.|The specified name and phone number do not belong to the same user.|The specified name and phone number do not belong to the same user.|
	//
	// |China Telecom|Verifications can be carried out normally.|The specified name and phone number cannot be found.|The specified name and phone number cannot be found.|
	//
	// example:
	//
	// 1
	IsConsistent *int32 `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s TwoElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TwoElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *TwoElementsVerificationResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *TwoElementsVerificationResponseBodyData) SetBasicCarrier(v string) *TwoElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *TwoElementsVerificationResponseBodyData) SetIsConsistent(v int32) *TwoElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *TwoElementsVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
