// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForSmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PhoneNumberStatusForSmsResponseBody
	GetCode() *string
	SetData(v *PhoneNumberStatusForSmsResponseBodyData) *PhoneNumberStatusForSmsResponseBody
	GetData() *PhoneNumberStatusForSmsResponseBodyData
	SetMessage(v string) *PhoneNumberStatusForSmsResponseBody
	GetMessage() *string
	SetRequestId(v string) *PhoneNumberStatusForSmsResponseBody
	GetRequestId() *string
}

type PhoneNumberStatusForSmsResponseBody struct {
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
	Data *PhoneNumberStatusForSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 68A40250-50CD-034C-B728-0BD135850177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForSmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponseBody) GetCode() *string {
	return s.Code
}

func (s *PhoneNumberStatusForSmsResponseBody) GetData() *PhoneNumberStatusForSmsResponseBodyData {
	return s.Data
}

func (s *PhoneNumberStatusForSmsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PhoneNumberStatusForSmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PhoneNumberStatusForSmsResponseBody) SetCode(v string) *PhoneNumberStatusForSmsResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetData(v *PhoneNumberStatusForSmsResponseBodyData) *PhoneNumberStatusForSmsResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetMessage(v string) *PhoneNumberStatusForSmsResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) SetRequestId(v string) *PhoneNumberStatusForSmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumberStatusForSmsResponseBodyData struct {
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
	// 	- **DEFECT**: The queried phone number is invalid.
	//
	// 	- **UNKNOWN**: The queried phone number is unknown.
	//
	// >  Due to system adjustment of the carrier, the BUSY, SUSPECTED_POWER_OFF, and POWER_OFF states cannot be returned for the numbers assigned by China Telecom. [For more information, see the official announcements](https://help.aliyun.com/document_detail/2489709.html).
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForSmsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForSmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *PhoneNumberStatusForSmsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *PhoneNumberStatusForSmsResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForSmsResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBodyData) SetStatus(v string) *PhoneNumberStatusForSmsResponseBodyData {
	s.Status = &v
	return s
}

func (s *PhoneNumberStatusForSmsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
