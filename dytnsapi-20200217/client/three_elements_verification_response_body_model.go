// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThreeElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ThreeElementsVerificationResponseBody
	GetCode() *string
	SetData(v *ThreeElementsVerificationResponseBodyData) *ThreeElementsVerificationResponseBody
	GetData() *ThreeElementsVerificationResponseBodyData
	SetMessage(v string) *ThreeElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ThreeElementsVerificationResponseBody
	GetRequestId() *string
}

type ThreeElementsVerificationResponseBody struct {
	// The response code.
	//
	// 	- **OK**: The request is successful.
	//
	// 	- For more information, see Error codes in this documentation.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *ThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ThreeElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ThreeElementsVerificationResponseBody) GetData() *ThreeElementsVerificationResponseBodyData {
	return s.Data
}

func (s *ThreeElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ThreeElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ThreeElementsVerificationResponseBody) SetCode(v string) *ThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetData(v *ThreeElementsVerificationResponseBodyData) *ThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetMessage(v string) *ThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) SetRequestId(v string) *ThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ThreeElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ThreeElementsVerificationResponseBodyData struct {
	// The basic carrier. Valid values:
	//
	// 	- **China Mobile**
	//
	// 	- **China Unicom**
	//
	// 	- **China Telecom**
	//
	// example:
	//
	// China Mobile
	BasicCarrier *string `json:"BasicCarrier,omitempty" xml:"BasicCarrier,omitempty"`
	// Indicates whether the specified name, phone number, and ID card number belong to the same user. Valid values:
	//
	// 	- **1**: The specified name, phone number, and ID card number belong to the same user.
	//
	// 	- **0**: The specified name, phone number, and ID card number do not belong to the same user.
	//
	// 	- **2**: The specified name, phone number, and ID card number cannot be found.
	//
	// **Note*	- The phone number registration data of a user is usually updated one or three days after registration. The registration data can be queried only after the update. The following table shows the verification results under different phone number states.
	//
	// |Carrier/Phone number state|Out-of-service|Nonexistent|Canceled|
	//
	// |---|---|---|---|
	//
	// |China Mobile|Verifications can be carried out normally.|The specified name, phone number, and ID card number cannot be found.|The specified name, phone number, and ID card number cannot be found.|
	//
	// |China Unicom|Verifications can be carried out normally.|The specified name, phone number, and ID card number do not belong to the same user.|The specified name, phone number, and ID card number do not belong to the same user.|
	//
	// |China Telecom|Verifications can be carried out normally.|The specified name, phone number, and ID card number cannot be found.|The specified name, phone number, and ID card number cannot be found.|
	//
	// example:
	//
	// 1
	IsConsistent *int32 `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
}

func (s ThreeElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationResponseBodyData) GetBasicCarrier() *string {
	return s.BasicCarrier
}

func (s *ThreeElementsVerificationResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *ThreeElementsVerificationResponseBodyData) SetBasicCarrier(v string) *ThreeElementsVerificationResponseBodyData {
	s.BasicCarrier = &v
	return s
}

func (s *ThreeElementsVerificationResponseBodyData) SetIsConsistent(v int32) *ThreeElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *ThreeElementsVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
