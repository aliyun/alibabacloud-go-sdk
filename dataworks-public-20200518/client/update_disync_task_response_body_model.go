// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDISyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateDISyncTaskResponseBodyData) *UpdateDISyncTaskResponseBody
	GetData() *UpdateDISyncTaskResponseBodyData
	SetRequestId(v string) *UpdateDISyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDISyncTaskResponseBody
	GetSuccess() *bool
}

type UpdateDISyncTaskResponseBody struct {
	// The information that indicates whether the data synchronization task is updated.
	Data *UpdateDISyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdateDISyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDISyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDISyncTaskResponseBody) GetData() *UpdateDISyncTaskResponseBodyData {
	return s.Data
}

func (s *UpdateDISyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDISyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDISyncTaskResponseBody) SetData(v *UpdateDISyncTaskResponseBodyData) *UpdateDISyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDISyncTaskResponseBody) SetRequestId(v string) *UpdateDISyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDISyncTaskResponseBody) SetSuccess(v bool) *UpdateDISyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDISyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateDISyncTaskResponseBodyData struct {
	// The error message returned if the data synchronization task fails to be updated. If the data synchronization task is successfully updated, the value null is returned for this parameter.
	//
	// example:
	//
	// ResourceGroup:[S_res_group_XXX] is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The update status of the data synchronization task. Valid values:
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

func (s UpdateDISyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDISyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDISyncTaskResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *UpdateDISyncTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateDISyncTaskResponseBodyData) SetMessage(v string) *UpdateDISyncTaskResponseBodyData {
	s.Message = &v
	return s
}

func (s *UpdateDISyncTaskResponseBodyData) SetStatus(v string) *UpdateDISyncTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateDISyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
