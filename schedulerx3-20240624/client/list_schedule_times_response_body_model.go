// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleTimesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListScheduleTimesResponseBody
	GetCode() *int32
	SetData(v []*string) *ListScheduleTimesResponseBody
	GetData() []*string
	SetMessage(v string) *ListScheduleTimesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScheduleTimesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScheduleTimesResponseBody
	GetSuccess() *bool
}

type ListScheduleTimesResponseBody struct {
	// example:
	//
	// 200
	Code *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9A48E22F-F30A-5CE5-AC7A-E0FED1B6942E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScheduleTimesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTimesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleTimesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListScheduleTimesResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListScheduleTimesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScheduleTimesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduleTimesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScheduleTimesResponseBody) SetCode(v int32) *ListScheduleTimesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScheduleTimesResponseBody) SetData(v []*string) *ListScheduleTimesResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduleTimesResponseBody) SetMessage(v string) *ListScheduleTimesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduleTimesResponseBody) SetRequestId(v string) *ListScheduleTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleTimesResponseBody) SetSuccess(v bool) *ListScheduleTimesResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduleTimesResponseBody) Validate() error {
	return dara.Validate(s)
}
