// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDISyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDISyncTaskResponseBodyData) *DeleteDISyncTaskResponseBody
	GetData() *DeleteDISyncTaskResponseBodyData
	SetRequestId(v string) *DeleteDISyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDISyncTaskResponseBody
	GetSuccess() *bool
}

type DeleteDISyncTaskResponseBody struct {
	// The deletion result.
	Data *DeleteDISyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DeleteDISyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDISyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDISyncTaskResponseBody) GetData() *DeleteDISyncTaskResponseBodyData {
	return s.Data
}

func (s *DeleteDISyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDISyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDISyncTaskResponseBody) SetData(v *DeleteDISyncTaskResponseBodyData) *DeleteDISyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDISyncTaskResponseBody) SetRequestId(v string) *DeleteDISyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDISyncTaskResponseBody) SetSuccess(v bool) *DeleteDISyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDISyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteDISyncTaskResponseBodyData struct {
	// The reason why the synchronization task fails to be deleted. If the synchronization task is deleted, the value null is returned for this parameter.
	//
	// example:
	//
	// fileId:[100] is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the synchronization task is deleted. Valid values:
	//
	// 	- success: The synchronization task is deleted.
	//
	// 	- fail: The synchronization task fails to be deleted. You can troubleshoot the issue based on the failure reason.
	//
	// example:
	//
	// fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteDISyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDISyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDISyncTaskResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteDISyncTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteDISyncTaskResponseBodyData) SetMessage(v string) *DeleteDISyncTaskResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteDISyncTaskResponseBodyData) SetStatus(v string) *DeleteDISyncTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteDISyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
