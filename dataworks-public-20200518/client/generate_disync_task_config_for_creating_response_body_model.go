// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDISyncTaskConfigForCreatingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateDISyncTaskConfigForCreatingResponseBodyData) *GenerateDISyncTaskConfigForCreatingResponseBody
	GetData() *GenerateDISyncTaskConfigForCreatingResponseBodyData
	SetRequestId(v string) *GenerateDISyncTaskConfigForCreatingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateDISyncTaskConfigForCreatingResponseBody
	GetSuccess() *bool
}

type GenerateDISyncTaskConfigForCreatingResponseBody struct {
	// The information returned for the ID of the asynchronous thread.
	Data *GenerateDISyncTaskConfigForCreatingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
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

func (s GenerateDISyncTaskConfigForCreatingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDISyncTaskConfigForCreatingResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBody) GetData() *GenerateDISyncTaskConfigForCreatingResponseBodyData {
	return s.Data
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBody) SetData(v *GenerateDISyncTaskConfigForCreatingResponseBodyData) *GenerateDISyncTaskConfigForCreatingResponseBody {
	s.Data = v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBody) SetRequestId(v string) *GenerateDISyncTaskConfigForCreatingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBody) SetSuccess(v bool) *GenerateDISyncTaskConfigForCreatingResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateDISyncTaskConfigForCreatingResponseBodyData struct {
	// The reason why the ID of the asynchronous thread fails to be generated. If the ID is successfully generated, no value is returned for this parameter.
	//
	// example:
	//
	// XXX is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the asynchronous thread. You can call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. The parameters are used to create a synchronization task in Data Integration.
	//
	// example:
	//
	// 10
	ProcessId *int64 `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Indicates whether the ID of the asynchronous thread is generated. Valid values: Valid values:
	//
	// 	- success: indicates that the ID of the asynchronous thread is generated.
	//
	// 	- fail: indicates that the ID of the asynchronous thread fails to be generated. You can view the reason for the failure and troubleshoot the issue based on the reason.
	//
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GenerateDISyncTaskConfigForCreatingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateDISyncTaskConfigForCreatingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBodyData) GetProcessId() *int64 {
	return s.ProcessId
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBodyData) SetMessage(v string) *GenerateDISyncTaskConfigForCreatingResponseBodyData {
	s.Message = &v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBodyData) SetProcessId(v int64) *GenerateDISyncTaskConfigForCreatingResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBodyData) SetStatus(v string) *GenerateDISyncTaskConfigForCreatingResponseBodyData {
	s.Status = &v
	return s
}

func (s *GenerateDISyncTaskConfigForCreatingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
