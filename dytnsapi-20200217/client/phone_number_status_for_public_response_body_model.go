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
	// The status code of the request. Valid values:
	//
	// - **OK**: The request was successful.
	//
	// - **OperatorLimit**: The query for the phone number is prohibited by the carrier.
	//
	// - **RequestFrequencyLimit**: Carrier restrictions prohibit frequent queries for the same number in a short period. If this error code is returned, try again later.
	//
	// > For a list of other error codes, see [API Error Center](https://next.api.aliyun.com/document/Dytnsapi/2020-02-17/errorCode).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *PhoneNumberStatusForPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
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
	// The basic carrier of the number. If the number has been ported, this parameter returns the current carrier.
	//
	// Valid values:
	//
	// - **CMCC**: China Mobile
	//
	// - **CUCC**: China Unicom
	//
	// - **CTCC**: China Telecom
	//
	// - **CBN**: China Broadnet
	//
	// example:
	//
	// CMCC
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The status of the queried phone number. Valid values:
	//
	// - **NORMAL**: The number is in service.
	//
	// - **SHUTDOWN**: The service for the number is suspended.
	//
	// - **POWER_OFF**: The phone is powered off.
	//
	// - **NOT_EXIST**: The number is non-existent.
	//
	// - **SUSPECTED_POWER_OFF**: The phone is suspected to be powered off.
	//
	// - **BUSY**: The line is busy.
	//
	// - **UNKNOWN**: The status is unknown.
	//
	// > Due to carrier system adjustments, the `BUSY`, `SUSPECTED_POWER_OFF`, and `POWER_OFF` statuses are not returned for China Telecom numbers. For more information, see the [official announcement](https://help.aliyun.com/document_detail/2489709.html).
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
