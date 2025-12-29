// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PhoneNumberStatusForAccountResponseBody
	GetCode() *string
	SetData(v *PhoneNumberStatusForAccountResponseBodyData) *PhoneNumberStatusForAccountResponseBody
	GetData() *PhoneNumberStatusForAccountResponseBodyData
	SetMessage(v string) *PhoneNumberStatusForAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *PhoneNumberStatusForAccountResponseBody
	GetRequestId() *string
}

type PhoneNumberStatusForAccountResponseBody struct {
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
	Data *PhoneNumberStatusForAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s PhoneNumberStatusForAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *PhoneNumberStatusForAccountResponseBody) GetData() *PhoneNumberStatusForAccountResponseBodyData {
	return s.Data
}

func (s *PhoneNumberStatusForAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PhoneNumberStatusForAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PhoneNumberStatusForAccountResponseBody) SetCode(v string) *PhoneNumberStatusForAccountResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetData(v *PhoneNumberStatusForAccountResponseBodyData) *PhoneNumberStatusForAccountResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetMessage(v string) *PhoneNumberStatusForAccountResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) SetRequestId(v string) *PhoneNumberStatusForAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumberStatusForAccountResponseBodyData struct {
	// The basic carrier who assings the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned. Valid values:
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
	// 	- **NORMAL**: The queried phone number is valid.
	//
	// 	- **SHUTDOWN**: The queried phone number is suspended.
	//
	// 	- **POWER_OFF**: The queried phone number cannot be connected.
	//
	// 	- **NOT_EXIST**: The queried phone number is a nonexistent number.
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

func (s PhoneNumberStatusForAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *PhoneNumberStatusForAccountResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *PhoneNumberStatusForAccountResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForAccountResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBodyData) SetStatus(v string) *PhoneNumberStatusForAccountResponseBodyData {
	s.Status = &v
	return s
}

func (s *PhoneNumberStatusForAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
