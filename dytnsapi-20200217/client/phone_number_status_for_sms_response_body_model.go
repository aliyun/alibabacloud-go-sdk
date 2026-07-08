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
	// The request status code. Valid values:
	//
	// - **OK**: The request was successful.
	//
	// - **OperatorLimit**: The carrier restricts queries for this phone number.
	//
	// - **RequestFrequencyLimit**: Indicates that requests for a single number are too frequent. Due to carrier restrictions, repeated queries for the same number within a short period are prohibited. If you receive this error code, try again later.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A container for the returned data.
	Data *PhoneNumberStatusForSmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request. Use this ID to troubleshoot issues.
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
	// The carrier of the phone number. If the number has been ported, this parameter returns the current carrier. Valid values:
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
	// - **NORMAL**: Active.
	//
	// - **SHUTDOWN**: Shutdown.
	//
	// - **POWER_OFF**: Powered off.
	//
	// - **NOT_EXIST**: Non-existent number.
	//
	// - **DEFECT**: Invalid number.
	//
	// - **UNKNOWN**: Unknown.
	//
	// > Due to carrier system adjustments, the statuses for busy, suspected to be powered off, and powered off are not returned for China Telecom numbers. For more information, see the [official announcement](https://help.aliyun.com/document_detail/2489709.html).
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
