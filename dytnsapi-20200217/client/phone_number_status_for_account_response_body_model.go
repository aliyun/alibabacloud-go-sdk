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
	// - **OK**: The request was successful.
	//
	// - **OperatorLimit**: The query is prohibited by the carrier.
	//
	// - **RequestFrequencyLimit**: Carriers restrict frequent queries for the same number within a short period. If you receive this error code, try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response object.
	Data *PhoneNumberStatusForAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. This ID is unique to each request and can be used for troubleshooting.
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
	// The number\\"s current carrier. If the number has been ported to a new carrier through mobile number portability, the new carrier is returned. Valid values:
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
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The status of the phone number. Valid values:
	//
	// - **NORMAL**: The number is active.
	//
	// - **SHUTDOWN**: The number is suspended or temporarily out of service.
	//
	// - **POWER_OFF**: The phone is powered off.
	//
	// - **NOT_EXIST**: The number is non-existent.
	//
	// - **DEFECT**: The number is invalid.
	//
	// - **UNKNOWN**: The status is unknown.
	//
	// > Due to adjustments in the carrier\\"s system, China Telecom numbers do not return the `busy` and `powered off` statuses. For more information, [see the official announcement](https://help.aliyun.com/document_detail/2489709.html).
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
