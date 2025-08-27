// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightSegmentAvailableCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *IntlFlightSegmentAvailableCertResponseBodyModule) *IntlFlightSegmentAvailableCertResponseBody
	GetModule() *IntlFlightSegmentAvailableCertResponseBodyModule
	SetRequestId(v string) *IntlFlightSegmentAvailableCertResponseBody
	GetRequestId() *string
	SetResultCode(v string) *IntlFlightSegmentAvailableCertResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *IntlFlightSegmentAvailableCertResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *IntlFlightSegmentAvailableCertResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightSegmentAvailableCertResponseBody
	GetTraceId() *string
}

type IntlFlightSegmentAvailableCertResponseBody struct {
	Module *IntlFlightSegmentAvailableCertResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 210bc81a17090871660176894d008c
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	ResultMsg  *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 213e1ea516895592036143147e5864
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightSegmentAvailableCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightSegmentAvailableCertResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightSegmentAvailableCertResponseBody) GetModule() *IntlFlightSegmentAvailableCertResponseBodyModule {
	return s.Module
}

func (s *IntlFlightSegmentAvailableCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightSegmentAvailableCertResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *IntlFlightSegmentAvailableCertResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *IntlFlightSegmentAvailableCertResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightSegmentAvailableCertResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightSegmentAvailableCertResponseBody) SetModule(v *IntlFlightSegmentAvailableCertResponseBodyModule) *IntlFlightSegmentAvailableCertResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBody) SetRequestId(v string) *IntlFlightSegmentAvailableCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBody) SetResultCode(v string) *IntlFlightSegmentAvailableCertResponseBody {
	s.ResultCode = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBody) SetResultMsg(v string) *IntlFlightSegmentAvailableCertResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBody) SetSuccess(v bool) *IntlFlightSegmentAvailableCertResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBody) SetTraceId(v string) *IntlFlightSegmentAvailableCertResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightSegmentAvailableCertResponseBodyModule struct {
	SegmentAvailableCertList []*IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList `json:"segment_available_cert_list,omitempty" xml:"segment_available_cert_list,omitempty" type:"Repeated"`
}

func (s IntlFlightSegmentAvailableCertResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightSegmentAvailableCertResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModule) GetSegmentAvailableCertList() []*IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList {
	return s.SegmentAvailableCertList
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModule) SetSegmentAvailableCertList(v []*IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) *IntlFlightSegmentAvailableCertResponseBodyModule {
	s.SegmentAvailableCertList = v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList struct {
	CertTypes       []*int32                                                                                 `json:"cert_types,omitempty" xml:"cert_types,omitempty" type:"Repeated"`
	SegmentPosition *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
}

func (s IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) GoString() string {
	return s.String()
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) GetCertTypes() []*int32 {
	return s.CertTypes
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) GetSegmentPosition() *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition {
	return s.SegmentPosition
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) SetCertTypes(v []*int32) *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList {
	s.CertTypes = v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) SetSegmentPosition(v *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList {
	s.SegmentPosition = v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) GoString() string {
	return s.String()
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) SetJourneyIndex(v int32) *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) SetSegmentIndex(v int32) *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponseBodyModuleSegmentAvailableCertListSegmentPosition) Validate() error {
	return dara.Validate(s)
}
