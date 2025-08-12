// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySiteMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySiteMonitorResponseBody
	GetCode() *string
	SetData(v *ModifySiteMonitorResponseBodyData) *ModifySiteMonitorResponseBody
	GetData() *ModifySiteMonitorResponseBodyData
	SetMessage(v string) *ModifySiteMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySiteMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifySiteMonitorResponseBody
	GetSuccess() *string
}

type ModifySiteMonitorResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of modifying the task.
	Data *ModifySiteMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68192f5d-0d45-4b98-9724-892813f86c71
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySiteMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySiteMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySiteMonitorResponseBody) GetData() *ModifySiteMonitorResponseBodyData {
	return s.Data
}

func (s *ModifySiteMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySiteMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySiteMonitorResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifySiteMonitorResponseBody) SetCode(v string) *ModifySiteMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetData(v *ModifySiteMonitorResponseBodyData) *ModifySiteMonitorResponseBody {
	s.Data = v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetMessage(v string) *ModifySiteMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetRequestId(v string) *ModifySiteMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySiteMonitorResponseBody) SetSuccess(v string) *ModifySiteMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySiteMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifySiteMonitorResponseBodyData struct {
	// The number of site monitoring tasks.
	//
	// example:
	//
	// 1
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s ModifySiteMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySiteMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *ModifySiteMonitorResponseBodyData) SetCount(v int32) *ModifySiteMonitorResponseBodyData {
	s.Count = &v
	return s
}

func (s *ModifySiteMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
