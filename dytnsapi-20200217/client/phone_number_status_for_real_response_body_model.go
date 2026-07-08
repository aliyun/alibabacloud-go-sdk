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
	// The request status code. Valid values:
	//
	// - **OK**: The request was successful.
	//
	// - **OperatorLimit**: The query for the phone number is restricted by the carrier.
	//
	// - **RequestFrequencyLimit**: Carriers prohibit high-frequency queries for the same number within a short period. If this error code is returned, try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned for the request.
	Data *PhoneNumberStatusForRealResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A unique identifier for the request. You can use this ID to troubleshoot issues.
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
	// The carrier that provides service for the phone number. If the number has been ported through mobile number portability (MNP), this field returns the new carrier. Valid values:
	//
	// - **CMCC**: China Mobile.
	//
	// - **CUCC**: China Unicom.
	//
	// - **CTCC**: China Telecom.
	//
	// > Queries for China Broadnet numbers are not supported.
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The status of the phone number. Valid values:
	//
	// - **NORMAL**: The number is in service.
	//
	// - **SHUTDOWN**: The service for the number is suspended.
	//
	// - **POWER_OFF**: The phone is powered off.
	//
	// - **NOT_EXIST**: The number is not in service.
	//
	// - **BUSY**: The line is busy.
	//
	// - **SUSPECTED_POWER_OFF**: The phone is suspected to be powered off.
	//
	// - **DEFECT**: The number is invalid.
	//
	// - **UNKNOWN**: The status is unknown.
	//
	// > Due to carrier system adjustments, China Telecom numbers no longer return the `BUSY`, `SUSPECTED_POWER_OFF`, and `POWER_OFF` statuses. For more information, see the [official announcement](https://help.aliyun.com/document_detail/2489709.html).
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
