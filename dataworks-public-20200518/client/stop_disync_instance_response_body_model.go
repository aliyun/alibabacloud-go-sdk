// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDISyncInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StopDISyncInstanceResponseBodyData) *StopDISyncInstanceResponseBody
	GetData() *StopDISyncInstanceResponseBodyData
	SetRequestId(v string) *StopDISyncInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDISyncInstanceResponseBody
	GetSuccess() *bool
}

type StopDISyncInstanceResponseBody struct {
	// The information returned for the synchronization task.
	Data *StopDISyncInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0bc1411515937635973****
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

func (s StopDISyncInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDISyncInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDISyncInstanceResponseBody) GetData() *StopDISyncInstanceResponseBodyData {
	return s.Data
}

func (s *StopDISyncInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDISyncInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDISyncInstanceResponseBody) SetData(v *StopDISyncInstanceResponseBodyData) *StopDISyncInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StopDISyncInstanceResponseBody) SetRequestId(v string) *StopDISyncInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDISyncInstanceResponseBody) SetSuccess(v bool) *StopDISyncInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StopDISyncInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type StopDISyncInstanceResponseBodyData struct {
	// The reason why the synchronization task fails to be stopped.
	//
	// If the synchronization task is stopped, the value null is returned.
	//
	// example:
	//
	// fileId:[100] is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the synchronization task is stopped. Valid values:
	//
	// 	- success
	//
	// 	- fail
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StopDISyncInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StopDISyncInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopDISyncInstanceResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *StopDISyncInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *StopDISyncInstanceResponseBodyData) SetMessage(v string) *StopDISyncInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *StopDISyncInstanceResponseBodyData) SetStatus(v string) *StopDISyncInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *StopDISyncInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
