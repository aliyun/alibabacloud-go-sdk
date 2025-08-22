// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCalendarNamesResponseBody
	GetCode() *int32
	SetData(v []*string) *ListCalendarNamesResponseBody
	GetData() []*string
	SetMessage(v string) *ListCalendarNamesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCalendarNamesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCalendarNamesResponseBody
	GetSuccess() *bool
}

type ListCalendarNamesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AA3538A0-FBE6-5E31-AD88-A02C6FF0DACC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCalendarNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalendarNamesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCalendarNamesResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListCalendarNamesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCalendarNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCalendarNamesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCalendarNamesResponseBody) SetCode(v int32) *ListCalendarNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCalendarNamesResponseBody) SetData(v []*string) *ListCalendarNamesResponseBody {
	s.Data = v
	return s
}

func (s *ListCalendarNamesResponseBody) SetMessage(v string) *ListCalendarNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCalendarNamesResponseBody) SetRequestId(v string) *ListCalendarNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalendarNamesResponseBody) SetSuccess(v bool) *ListCalendarNamesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCalendarNamesResponseBody) Validate() error {
	return dara.Validate(s)
}
