// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCloudBenchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunCloudBenchTaskResponseBody
	GetCode() *string
	SetData(v *RunCloudBenchTaskResponseBodyData) *RunCloudBenchTaskResponseBody
	GetData() *RunCloudBenchTaskResponseBodyData
	SetMessage(v string) *RunCloudBenchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunCloudBenchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *RunCloudBenchTaskResponseBody
	GetSuccess() *string
}

type RunCloudBenchTaskResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of returned entries.
	Data *RunCloudBenchTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunCloudBenchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCloudBenchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunCloudBenchTaskResponseBody) GetData() *RunCloudBenchTaskResponseBodyData {
	return s.Data
}

func (s *RunCloudBenchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunCloudBenchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCloudBenchTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *RunCloudBenchTaskResponseBody) SetCode(v string) *RunCloudBenchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetData(v *RunCloudBenchTaskResponseBodyData) *RunCloudBenchTaskResponseBody {
	s.Data = v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetMessage(v string) *RunCloudBenchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetRequestId(v string) *RunCloudBenchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCloudBenchTaskResponseBody) SetSuccess(v string) *RunCloudBenchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *RunCloudBenchTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunCloudBenchTaskResponseBodyData struct {
	PreCheckItem []*RunCloudBenchTaskResponseBodyDataPreCheckItem `json:"PreCheckItem,omitempty" xml:"PreCheckItem,omitempty" type:"Repeated"`
}

func (s RunCloudBenchTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunCloudBenchTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponseBodyData) GetPreCheckItem() []*RunCloudBenchTaskResponseBodyDataPreCheckItem {
	return s.PreCheckItem
}

func (s *RunCloudBenchTaskResponseBodyData) SetPreCheckItem(v []*RunCloudBenchTaskResponseBodyDataPreCheckItem) *RunCloudBenchTaskResponseBodyData {
	s.PreCheckItem = v
	return s
}

func (s *RunCloudBenchTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RunCloudBenchTaskResponseBodyDataPreCheckItem struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information of the check item.
	//
	// example:
	//
	// "Data": { "total": 1, "list":[...] }, "Code": 200, "Success": true }
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the check item. Valid values:
	//
	// 	- **SqlArchiveStatusChecker**: checks whether SQL Explorer is available.
	//
	// 	- **BenchClientEnvChecker**: checks whether the runtime environment for programs on the stress testing client is available.
	//
	// 	- **SpecChecker**: checks whether the destination instance type and the instance type of the stress testing client support this API operation.
	//
	// 	- **SourceInstanceChecker**: checks whether the account of the source instance is available and whether the source instance is connected to the destination instance.
	//
	// 	- **BenchTargetChecker**: checks whether the account of the destination instance is available and whether the source instance is connected to the destination instance.
	//
	// example:
	//
	// BenchTargetChecker
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sequence number of the check item. Valid values: **0*	- to **10**.
	//
	// example:
	//
	// 0
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **IGNORED**: The task is ignored.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **EXCEPTION**: An error occurred.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RunCloudBenchTaskResponseBodyDataPreCheckItem) String() string {
	return dara.Prettify(s)
}

func (s RunCloudBenchTaskResponseBodyDataPreCheckItem) GoString() string {
	return s.String()
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) GetCode() *int32 {
	return s.Code
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) GetDetails() *string {
	return s.Details
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) GetMessage() *string {
	return s.Message
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) GetName() *string {
	return s.Name
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) GetOrder() *int32 {
	return s.Order
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) GetStatus() *string {
	return s.Status
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetCode(v int32) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Code = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetDetails(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Details = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetMessage(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Message = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetName(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Name = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetOrder(v int32) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Order = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) SetStatus(v string) *RunCloudBenchTaskResponseBodyDataPreCheckItem {
	s.Status = &v
	return s
}

func (s *RunCloudBenchTaskResponseBodyDataPreCheckItem) Validate() error {
	return dara.Validate(s)
}
