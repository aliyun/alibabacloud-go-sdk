// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDISyncTaskConfigForUpdatingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GenerateDISyncTaskConfigForUpdatingResponseBodyData) *GenerateDISyncTaskConfigForUpdatingResponseBody
	GetData() *GenerateDISyncTaskConfigForUpdatingResponseBodyData
	SetRequestId(v string) *GenerateDISyncTaskConfigForUpdatingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateDISyncTaskConfigForUpdatingResponseBody
	GetSuccess() *bool
}

type GenerateDISyncTaskConfigForUpdatingResponseBody struct {
	// The information returned for the ID of the asynchronous thread.
	Data *GenerateDISyncTaskConfigForUpdatingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GenerateDISyncTaskConfigForUpdatingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDISyncTaskConfigForUpdatingResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBody) GetData() *GenerateDISyncTaskConfigForUpdatingResponseBodyData {
	return s.Data
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBody) SetData(v *GenerateDISyncTaskConfigForUpdatingResponseBodyData) *GenerateDISyncTaskConfigForUpdatingResponseBody {
	s.Data = v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBody) SetRequestId(v string) *GenerateDISyncTaskConfigForUpdatingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBody) SetSuccess(v bool) *GenerateDISyncTaskConfigForUpdatingResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateDISyncTaskConfigForUpdatingResponseBodyData struct {
	// The reason why the ID of the asynchronous thread fails to be generated. If the ID is successfully generated, no value is returned for this parameter.
	//
	// example:
	//
	// XXX is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the asynchronous thread. You can call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. The parameters are used to update a real-time synchronization task in Data Integration.
	//
	// example:
	//
	// 10
	ProcessId *int64 `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Indicates whether the ID of the asynchronous thread is generated. Valid values:
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

func (s GenerateDISyncTaskConfigForUpdatingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateDISyncTaskConfigForUpdatingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBodyData) GetProcessId() *int64 {
	return s.ProcessId
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBodyData) SetMessage(v string) *GenerateDISyncTaskConfigForUpdatingResponseBodyData {
	s.Message = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBodyData) SetProcessId(v int64) *GenerateDISyncTaskConfigForUpdatingResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBodyData) SetStatus(v string) *GenerateDISyncTaskConfigForUpdatingResponseBodyData {
	s.Status = &v
	return s
}

func (s *GenerateDISyncTaskConfigForUpdatingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
