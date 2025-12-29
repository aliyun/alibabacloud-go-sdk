// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForRealResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PhoneNumberStatusForRealResponseBody
	GetCode() *string
	SetData(v *PhoneNumberStatusForRealResponseBodyData) *PhoneNumberStatusForRealResponseBody
	GetData() *PhoneNumberStatusForRealResponseBodyData
	SetMessage(v string) *PhoneNumberStatusForRealResponseBody
	GetMessage() *string
	SetRequestId(v string) *PhoneNumberStatusForRealResponseBody
	GetRequestId() *string
}

type PhoneNumberStatusForRealResponseBody struct {
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **OperatorLimit**: The carrier prohibits the query of the phone number.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *PhoneNumberStatusForRealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. It is a common parameter and can be used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForRealResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForRealResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponseBody) GetCode() *string {
	return s.Code
}

func (s *PhoneNumberStatusForRealResponseBody) GetData() *PhoneNumberStatusForRealResponseBodyData {
	return s.Data
}

func (s *PhoneNumberStatusForRealResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PhoneNumberStatusForRealResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PhoneNumberStatusForRealResponseBody) SetCode(v string) *PhoneNumberStatusForRealResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetData(v *PhoneNumberStatusForRealResponseBodyData) *PhoneNumberStatusForRealResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetMessage(v string) *PhoneNumberStatusForRealResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) SetRequestId(v string) *PhoneNumberStatusForRealResponseBody {
	s.RequestId = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumberStatusForRealResponseBodyData struct {
	// The basic carrier who assigns the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned. Valid values:
	//
	// 	- **CMCC**: China Mobile
	//
	// 	- **CUCC**: China Unicom
	//
	// 	- **CTCC**: China Telecom
	//
	// >  You are not allowed to query the phone numbers assigned by China Broadnet.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The returned status for the queried phone number. Valid values:
	//
	// 	- **NORMAL**: The queried phone number can be reached.
	//
	// 	- **SHUTDOWN**: The queried phone number is suspended.
	//
	// 	- **POWER_OFF**: The phone is powered off.
	//
	// 	- **NOT_EXIST**: The queried phone number is a nonexistent number.
	//
	// 	- **BUSY**: The queried phone number is busy.
	//
	// 	- **SUSPECTED_POWER_OFF**: The phone is suspected to be powered off.
	//
	// 	- **DEFECT**: The queried phone number is invalid.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// >  Due to system adjustment of the carrier, the BUSY and POWER_OFF states cannot be returned for the numbers assigned by China Telecom. [For more information, see the official announcements](https://help.aliyun.com/document_detail/2489709.html).
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForRealResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForRealResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *PhoneNumberStatusForRealResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *PhoneNumberStatusForRealResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForRealResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBodyData) SetStatus(v string) *PhoneNumberStatusForRealResponseBodyData {
	s.Status = &v
	return s
}

func (s *PhoneNumberStatusForRealResponseBodyData) Validate() error {
	return dara.Validate(s)
}
