// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDISyncInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartDISyncInstanceResponseBodyData) *StartDISyncInstanceResponseBody
	GetData() *StartDISyncInstanceResponseBodyData
	SetRequestId(v string) *StartDISyncInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartDISyncInstanceResponseBody
	GetSuccess() *bool
}

type StartDISyncInstanceResponseBody struct {
	// The result returned for the start.
	Data *StartDISyncInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s StartDISyncInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDISyncInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDISyncInstanceResponseBody) GetData() *StartDISyncInstanceResponseBodyData {
	return s.Data
}

func (s *StartDISyncInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDISyncInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartDISyncInstanceResponseBody) SetData(v *StartDISyncInstanceResponseBodyData) *StartDISyncInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StartDISyncInstanceResponseBody) SetRequestId(v string) *StartDISyncInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDISyncInstanceResponseBody) SetSuccess(v bool) *StartDISyncInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartDISyncInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartDISyncInstanceResponseBodyData struct {
	// The reason why the real-time synchronization task or the data synchronization solution fails to be started.
	//
	// If the real-time synchronization task or the data synchronization solution is started, the value null is returned.
	//
	// example:
	//
	// fileId:[100] is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the real-time synchronization task or the data synchronization solution is started. Valid values:
	//
	// 	- success: The real-time synchronization task or the data synchronization solution is started.
	//
	// 	- fail: The real-time synchronization task or the data synchronization solution fails to be started. You can troubleshoot the issue based on the provided cause.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StartDISyncInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartDISyncInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartDISyncInstanceResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *StartDISyncInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *StartDISyncInstanceResponseBodyData) SetMessage(v string) *StartDISyncInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *StartDISyncInstanceResponseBodyData) SetStatus(v string) *StartDISyncInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *StartDISyncInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
