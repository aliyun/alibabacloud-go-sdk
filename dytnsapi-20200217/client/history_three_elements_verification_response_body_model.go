// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHistoryThreeElementsVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *HistoryThreeElementsVerificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *HistoryThreeElementsVerificationResponseBody
	GetCode() *string
	SetData(v *HistoryThreeElementsVerificationResponseBodyData) *HistoryThreeElementsVerificationResponseBody
	GetData() *HistoryThreeElementsVerificationResponseBodyData
	SetMessage(v string) *HistoryThreeElementsVerificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *HistoryThreeElementsVerificationResponseBody
	GetRequestId() *string
}

type HistoryThreeElementsVerificationResponseBody struct {
	// Details about why access is denied.
	//
	// > This parameter is returned only when RAM authentication fails.
	//
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request.
	//
	// - A value of `OK` indicates the request was successful.
	//
	// - For other values, see the Error Codes section.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query results.
	Data *HistoryThreeElementsVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the status code.
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

func (s HistoryThreeElementsVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HistoryThreeElementsVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *HistoryThreeElementsVerificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *HistoryThreeElementsVerificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *HistoryThreeElementsVerificationResponseBody) GetData() *HistoryThreeElementsVerificationResponseBodyData {
	return s.Data
}

func (s *HistoryThreeElementsVerificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HistoryThreeElementsVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HistoryThreeElementsVerificationResponseBody) SetAccessDeniedDetail(v string) *HistoryThreeElementsVerificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *HistoryThreeElementsVerificationResponseBody) SetCode(v string) *HistoryThreeElementsVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *HistoryThreeElementsVerificationResponseBody) SetData(v *HistoryThreeElementsVerificationResponseBodyData) *HistoryThreeElementsVerificationResponseBody {
	s.Data = v
	return s
}

func (s *HistoryThreeElementsVerificationResponseBody) SetMessage(v string) *HistoryThreeElementsVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *HistoryThreeElementsVerificationResponseBody) SetRequestId(v string) *HistoryThreeElementsVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *HistoryThreeElementsVerificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HistoryThreeElementsVerificationResponseBodyData struct {
	// The consistency of the verification result. Valid values:
	//
	// - `0`: No record found.
	//
	// - `1`: The phone number, ID number, and name match the carrier\\"s records.
	//
	// - `2`: The phone number and ID number match the carrier\\"s records, but the name does not.
	//
	// - `3`: The phone number and name match the carrier\\"s records, but the ID number does not.
	//
	// - `4`: The phone number matches the carrier\\"s records, but the name and ID number do not.
	//
	// example:
	//
	// 72
	IsConsistent *int64 `json:"IsConsistent,omitempty" xml:"IsConsistent,omitempty"`
	// The carrier to which the request was routed.
	//
	// example:
	//
	// CMCC
	RequestCarrier *string `json:"RequestCarrier,omitempty" xml:"RequestCarrier,omitempty"`
}

func (s HistoryThreeElementsVerificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s HistoryThreeElementsVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *HistoryThreeElementsVerificationResponseBodyData) GetIsConsistent() *int64 {
	return s.IsConsistent
}

func (s *HistoryThreeElementsVerificationResponseBodyData) GetRequestCarrier() *string {
	return s.RequestCarrier
}

func (s *HistoryThreeElementsVerificationResponseBodyData) SetIsConsistent(v int64) *HistoryThreeElementsVerificationResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *HistoryThreeElementsVerificationResponseBodyData) SetRequestCarrier(v string) *HistoryThreeElementsVerificationResponseBodyData {
	s.RequestCarrier = &v
	return s
}

func (s *HistoryThreeElementsVerificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
