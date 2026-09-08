// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadNumGroupByClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadNumGroupByClassResponseBody
	GetCode() *string
	SetData(v []*ReadNumGroupByClassResponseBodyData) *ReadNumGroupByClassResponseBody
	GetData() []*ReadNumGroupByClassResponseBodyData
	SetMessage(v string) *ReadNumGroupByClassResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadNumGroupByClassResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadNumGroupByClassResponseBody
	GetSuccess() *bool
}

type ReadNumGroupByClassResponseBody struct {
	// The error code returned when the call fails. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The execution result.
	Data []*ReadNumGroupByClassResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message returned when the call fails.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadNumGroupByClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupByClassResponseBody) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadNumGroupByClassResponseBody) GetData() []*ReadNumGroupByClassResponseBodyData {
	return s.Data
}

func (s *ReadNumGroupByClassResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadNumGroupByClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadNumGroupByClassResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadNumGroupByClassResponseBody) SetCode(v string) *ReadNumGroupByClassResponseBody {
	s.Code = &v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetData(v []*ReadNumGroupByClassResponseBodyData) *ReadNumGroupByClassResponseBody {
	s.Data = v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetMessage(v string) *ReadNumGroupByClassResponseBody {
	s.Message = &v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetRequestId(v string) *ReadNumGroupByClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadNumGroupByClassResponseBody) SetSuccess(v bool) *ReadNumGroupByClassResponseBody {
	s.Success = &v
	return s
}

func (s *ReadNumGroupByClassResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadNumGroupByClassResponseBodyData struct {
	// The message category ID.
	//
	// example:
	//
	// 1
	ClassId *int64 `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The number of unread messages in the category.
	//
	// example:
	//
	// 1
	MsgCount *int64 `json:"MsgCount,omitempty" xml:"MsgCount,omitempty"`
}

func (s ReadNumGroupByClassResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupByClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassResponseBodyData) GetClassId() *int64 {
	return s.ClassId
}

func (s *ReadNumGroupByClassResponseBodyData) GetMsgCount() *int64 {
	return s.MsgCount
}

func (s *ReadNumGroupByClassResponseBodyData) SetClassId(v int64) *ReadNumGroupByClassResponseBodyData {
	s.ClassId = &v
	return s
}

func (s *ReadNumGroupByClassResponseBodyData) SetMsgCount(v int64) *ReadNumGroupByClassResponseBodyData {
	s.MsgCount = &v
	return s
}

func (s *ReadNumGroupByClassResponseBodyData) Validate() error {
	return dara.Validate(s)
}
