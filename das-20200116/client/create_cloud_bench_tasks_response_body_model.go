// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudBenchTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCloudBenchTasksResponseBody
	GetCode() *string
	SetData(v *CreateCloudBenchTasksResponseBodyData) *CreateCloudBenchTasksResponseBody
	GetData() *CreateCloudBenchTasksResponseBodyData
	SetMessage(v string) *CreateCloudBenchTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCloudBenchTasksResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateCloudBenchTasksResponseBody
	GetSuccess() *string
}

type CreateCloudBenchTasksResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *CreateCloudBenchTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCloudBenchTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudBenchTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCloudBenchTasksResponseBody) GetData() *CreateCloudBenchTasksResponseBodyData {
	return s.Data
}

func (s *CreateCloudBenchTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCloudBenchTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudBenchTasksResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateCloudBenchTasksResponseBody) SetCode(v string) *CreateCloudBenchTasksResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetData(v *CreateCloudBenchTasksResponseBodyData) *CreateCloudBenchTasksResponseBody {
	s.Data = v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetMessage(v string) *CreateCloudBenchTasksResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetRequestId(v string) *CreateCloudBenchTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) SetSuccess(v string) *CreateCloudBenchTasksResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCloudBenchTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateCloudBenchTasksResponseBodyData struct {
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s CreateCloudBenchTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudBenchTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCloudBenchTasksResponseBodyData) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *CreateCloudBenchTasksResponseBodyData) SetTaskIds(v []*string) *CreateCloudBenchTasksResponseBodyData {
	s.TaskIds = v
	return s
}

func (s *CreateCloudBenchTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
