// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportCalendarResponseBody
	GetCode() *int32
	SetData(v []*string) *ImportCalendarResponseBody
	GetData() []*string
	SetMessage(v string) *ImportCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportCalendarResponseBody
	GetSuccess() *bool
}

type ImportCalendarResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2ECA6FC9-7557-5576-AF5F-FC3E7BCC9C21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportCalendarResponseBody) GetData() []*string {
	return s.Data
}

func (s *ImportCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportCalendarResponseBody) SetCode(v int32) *ImportCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *ImportCalendarResponseBody) SetData(v []*string) *ImportCalendarResponseBody {
	s.Data = v
	return s
}

func (s *ImportCalendarResponseBody) SetMessage(v string) *ImportCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *ImportCalendarResponseBody) SetRequestId(v string) *ImportCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCalendarResponseBody) SetSuccess(v bool) *ImportCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *ImportCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
