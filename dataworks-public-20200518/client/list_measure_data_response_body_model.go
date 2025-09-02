// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMeasureDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListMeasureDataResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMeasureDataResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListMeasureDataResponseBody
	GetHttpStatusCode() *int32
	SetMeasureDatas(v []*ListMeasureDataResponseBodyMeasureDatas) *ListMeasureDataResponseBody
	GetMeasureDatas() []*ListMeasureDataResponseBodyMeasureDatas
	SetRequestId(v string) *ListMeasureDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMeasureDataResponseBody
	GetSuccess() *bool
}

type ListMeasureDataResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 100001001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The user is not in tenant.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The measurement results.
	MeasureDatas []*ListMeasureDataResponseBodyMeasureDatas `json:"MeasureDatas,omitempty" xml:"MeasureDatas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMeasureDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMeasureDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListMeasureDataResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMeasureDataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMeasureDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMeasureDataResponseBody) GetMeasureDatas() []*ListMeasureDataResponseBodyMeasureDatas {
	return s.MeasureDatas
}

func (s *ListMeasureDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMeasureDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMeasureDataResponseBody) SetErrorCode(v string) *ListMeasureDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMeasureDataResponseBody) SetErrorMessage(v string) *ListMeasureDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMeasureDataResponseBody) SetHttpStatusCode(v int32) *ListMeasureDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMeasureDataResponseBody) SetMeasureDatas(v []*ListMeasureDataResponseBodyMeasureDatas) *ListMeasureDataResponseBody {
	s.MeasureDatas = v
	return s
}

func (s *ListMeasureDataResponseBody) SetRequestId(v string) *ListMeasureDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMeasureDataResponseBody) SetSuccess(v bool) *ListMeasureDataResponseBody {
	s.Success = &v
	return s
}

func (s *ListMeasureDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMeasureDataResponseBodyMeasureDatas struct {
	// The measurement component.
	//
	// example:
	//
	// Count
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The item that is measured.
	//
	// example:
	//
	// DideAlarmPhone
	DomainCode *string `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	// The end timestamp of the measurement period, in milliseconds.
	//
	// example:
	//
	// 1717430400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start timestamp of the measurement period, in milliseconds.
	//
	// example:
	//
	// 1717344000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total quantity used within the measurement period.
	//
	// example:
	//
	// 10
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListMeasureDataResponseBodyMeasureDatas) String() string {
	return dara.Prettify(s)
}

func (s ListMeasureDataResponseBodyMeasureDatas) GoString() string {
	return s.String()
}

func (s *ListMeasureDataResponseBodyMeasureDatas) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *ListMeasureDataResponseBodyMeasureDatas) GetDomainCode() *string {
	return s.DomainCode
}

func (s *ListMeasureDataResponseBodyMeasureDatas) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMeasureDataResponseBodyMeasureDatas) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMeasureDataResponseBodyMeasureDatas) GetUsage() *int64 {
	return s.Usage
}

func (s *ListMeasureDataResponseBodyMeasureDatas) SetComponentCode(v string) *ListMeasureDataResponseBodyMeasureDatas {
	s.ComponentCode = &v
	return s
}

func (s *ListMeasureDataResponseBodyMeasureDatas) SetDomainCode(v string) *ListMeasureDataResponseBodyMeasureDatas {
	s.DomainCode = &v
	return s
}

func (s *ListMeasureDataResponseBodyMeasureDatas) SetEndTime(v int64) *ListMeasureDataResponseBodyMeasureDatas {
	s.EndTime = &v
	return s
}

func (s *ListMeasureDataResponseBodyMeasureDatas) SetStartTime(v int64) *ListMeasureDataResponseBodyMeasureDatas {
	s.StartTime = &v
	return s
}

func (s *ListMeasureDataResponseBodyMeasureDatas) SetUsage(v int64) *ListMeasureDataResponseBodyMeasureDatas {
	s.Usage = &v
	return s
}

func (s *ListMeasureDataResponseBodyMeasureDatas) Validate() error {
	return dara.Validate(s)
}
