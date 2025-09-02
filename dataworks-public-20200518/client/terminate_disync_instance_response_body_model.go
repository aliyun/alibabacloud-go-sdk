// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateDISyncInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TerminateDISyncInstanceResponseBodyData) *TerminateDISyncInstanceResponseBody
	GetData() *TerminateDISyncInstanceResponseBodyData
	SetRequestId(v string) *TerminateDISyncInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TerminateDISyncInstanceResponseBody
	GetSuccess() *bool
}

type TerminateDISyncInstanceResponseBody struct {
	// The returned results.
	Data *TerminateDISyncInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s TerminateDISyncInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateDISyncInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateDISyncInstanceResponseBody) GetData() *TerminateDISyncInstanceResponseBodyData {
	return s.Data
}

func (s *TerminateDISyncInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateDISyncInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TerminateDISyncInstanceResponseBody) SetData(v *TerminateDISyncInstanceResponseBodyData) *TerminateDISyncInstanceResponseBody {
	s.Data = v
	return s
}

func (s *TerminateDISyncInstanceResponseBody) SetRequestId(v string) *TerminateDISyncInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateDISyncInstanceResponseBody) SetSuccess(v bool) *TerminateDISyncInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *TerminateDISyncInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type TerminateDISyncInstanceResponseBodyData struct {
	// The reason why the real-time synchronization task fails to be terminated. If the real-time synchronization task is undeployed, the value of this parameter is null.
	//
	// example:
	//
	// fileId:[100] is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the real-time synchronization task is undeployed. Valid values:
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

func (s TerminateDISyncInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TerminateDISyncInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *TerminateDISyncInstanceResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *TerminateDISyncInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *TerminateDISyncInstanceResponseBodyData) SetMessage(v string) *TerminateDISyncInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *TerminateDISyncInstanceResponseBodyData) SetStatus(v string) *TerminateDISyncInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *TerminateDISyncInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
