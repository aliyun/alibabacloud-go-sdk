// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDISyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDISyncTaskResponseBodyData) *CreateDISyncTaskResponseBody
	GetData() *CreateDISyncTaskResponseBodyData
	SetRequestId(v string) *CreateDISyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDISyncTaskResponseBody
	GetSuccess() *bool
}

type CreateDISyncTaskResponseBody struct {
	// The information that indicates whether the data synchronization task is created.
	Data *CreateDISyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDISyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDISyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDISyncTaskResponseBody) GetData() *CreateDISyncTaskResponseBodyData {
	return s.Data
}

func (s *CreateDISyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDISyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDISyncTaskResponseBody) SetData(v *CreateDISyncTaskResponseBodyData) *CreateDISyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateDISyncTaskResponseBody) SetRequestId(v string) *CreateDISyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDISyncTaskResponseBody) SetSuccess(v bool) *CreateDISyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDISyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateDISyncTaskResponseBodyData struct {
	// The ID of the data synchronization task that is created.
	//
	// example:
	//
	// 1000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The error message that is returned if the data synchronization task fails to be created. If the data synchronization task is successfully created, this parameter is not returned. If the data synchronization task fails to be created, an error message in the "Invalid path: Workflow/xxxx/Data Integration" format is returned.
	//
	// example:
	//
	// Invalid path: Business Flow/xxxx/Data Integration
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The creation status of the data synchronization task. Valid values:
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

func (s CreateDISyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDISyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDISyncTaskResponseBodyData) GetFileId() *int64 {
	return s.FileId
}

func (s *CreateDISyncTaskResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CreateDISyncTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateDISyncTaskResponseBodyData) SetFileId(v int64) *CreateDISyncTaskResponseBodyData {
	s.FileId = &v
	return s
}

func (s *CreateDISyncTaskResponseBodyData) SetMessage(v string) *CreateDISyncTaskResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateDISyncTaskResponseBodyData) SetStatus(v string) *CreateDISyncTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateDISyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
