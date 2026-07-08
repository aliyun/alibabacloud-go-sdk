// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PhoneNumberStatusForVoiceResponseBody
	GetCode() *string
	SetData(v *PhoneNumberStatusForVoiceResponseBodyData) *PhoneNumberStatusForVoiceResponseBody
	GetData() *PhoneNumberStatusForVoiceResponseBodyData
	SetMessage(v string) *PhoneNumberStatusForVoiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *PhoneNumberStatusForVoiceResponseBody
	GetRequestId() *string
}

type PhoneNumberStatusForVoiceResponseBody struct {
	// The status code of the request. Valid values:
	//
	// - **OK**: The request was successful.
	//
	// - **OperatorLimit**: The carrier restricts queries for this phone number.
	//
	// - **RequestFrequencyLimit**: Carrier restrictions limit how frequently you can query the same number. If you receive this error, try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *PhoneNumberStatusForVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. This is a common parameter. Each request has a unique ID that you can use to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PhoneNumberStatusForVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *PhoneNumberStatusForVoiceResponseBody) GetData() *PhoneNumberStatusForVoiceResponseBodyData {
	return s.Data
}

func (s *PhoneNumberStatusForVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PhoneNumberStatusForVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetCode(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetData(v *PhoneNumberStatusForVoiceResponseBodyData) *PhoneNumberStatusForVoiceResponseBody {
	s.Data = v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetMessage(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) SetRequestId(v string) *PhoneNumberStatusForVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PhoneNumberStatusForVoiceResponseBodyData struct {
	// The current carrier for the number. If the number has been ported, this field returns the new carrier. Valid values:
	//
	// - **CMCC**: China Mobile
	//
	// - **CUCC**: China Unicom
	//
	// - **CTCC**: China Telecom
	//
	// > Queries for China Broadnet numbers are not supported.
	//
	// example:
	//
	// CTCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The status of the phone number. Valid values:
	//
	// - **NORMAL**: The number is active.
	//
	// - **SHUTDOWN**: The service for the number is suspended.
	//
	// - **POWER_OFF**: The phone is powered off.
	//
	// - **NOT_EXIST**: The number does not exist.
	//
	// - **SUSPECTED_POWER_OFF**: The phone is likely powered off.
	//
	// - **DEFECT**: The number is invalid.
	//
	// - **UNKNOWN**: The status is unknown.
	//
	// > Due to carrier system adjustments, the `SUSPECTED_POWER_OFF` and `POWER_OFF` statuses are not returned for China Telecom numbers. [For more information, see the official announcement.](https://help.aliyun.com/document_detail/2489709.html)
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PhoneNumberStatusForVoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForVoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) GetCarrier() *string {
	return s.Carrier
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) SetCarrier(v string) *PhoneNumberStatusForVoiceResponseBodyData {
	s.Carrier = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) SetStatus(v string) *PhoneNumberStatusForVoiceResponseBodyData {
	s.Status = &v
	return s
}

func (s *PhoneNumberStatusForVoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
