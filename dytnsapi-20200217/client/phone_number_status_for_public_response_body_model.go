// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForPublicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PhoneNumberStatusForPublicResponseBody
	GetCode() *string
	SetData(v *PhoneNumberStatusForPublicResponseBodyData) *PhoneNumberStatusForPublicResponseBody
	GetData() *PhoneNumberStatusForPublicResponseBodyData
	SetMessage(v string) *PhoneNumberStatusForPublicResponseBody
	GetMessage() *string
	SetRequestId(v string) *PhoneNumberStatusForPublicResponseBody
	GetRequestId() *string
}

type PhoneNumberStatusForPublicResponseBody struct {
	// The response code. Valid values:
	//
	// 	- **OK**: The request is successful.
	//
	// 	- **OperatorLimit**: The carrier prohibits the query of the phone number.
	//
	// 	- **RequestFrequencyLimit**: Repeated queries for the same phone number at a high frequency within a short period of time are prohibited due to restrictions that are set by carriers. If this error code is returned, please try again later.
	//
	// >  For a list of error codes, see [Service error codes](https://next.api.aliyun.com/document/Dytnsapi/2020-02-17/errorCode).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *PhoneNumberStatusForPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// CC3BB6D2-****-****-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForPublicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForPublicResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicResponseBody) GetCode() *string {
	return s.Code
}

func (s *PhoneNumberStatusForPublicResponseBody) GetData() *PhoneNumberStatusForPublicResponseBodyData {
	return s.Data
}

func (s *PhoneNumberStatusForPublicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PhoneNumberStatusForPublicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PhoneNumberStatusForPublicResponseBody) SetCode(v string) *PhoneNumberStatusForPublicResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBody) SetData(v *PhoneNumberStatusForPublicResponseBodyData) *PhoneNumberStatusForPublicResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBody) SetMessage(v string) *PhoneNumberStatusForPublicResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBody) SetRequestId(v string) *PhoneNumberStatusForPublicResponseBody {
	s.RequestId = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumberStatusForPublicResponseBodyData struct {
	// The basic carrier who assigns the phone number. If the queried phone number involves mobile number portability, the carrier after mobile number portability is returned.
	//
	// Valid values:
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
	// 	- **SUSPECTED_POWER_OFF**: The phone is suspected to be powered off.
	//
	// 	- **BUSY**: The queried phone number is busy.
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

func (s PhoneNumberStatusForPublicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *PhoneNumberStatusForPublicResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *PhoneNumberStatusForPublicResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForPublicResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBodyData) SetStatus(v string) *PhoneNumberStatusForPublicResponseBodyData {
	s.Status = &v
	return s
}

func (s *PhoneNumberStatusForPublicResponseBodyData) Validate() error {
	return dara.Validate(s)
}
