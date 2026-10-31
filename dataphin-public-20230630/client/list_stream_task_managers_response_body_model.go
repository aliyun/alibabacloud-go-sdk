// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStreamTaskManagersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListStreamTaskManagersResponseBody
	GetCode() *string
	SetData(v []*ListStreamTaskManagersResponseBodyData) *ListStreamTaskManagersResponseBody
	GetData() []*ListStreamTaskManagersResponseBodyData
	SetHttpStatusCode(v int32) *ListStreamTaskManagersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListStreamTaskManagersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListStreamTaskManagersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListStreamTaskManagersResponseBody
	GetSuccess() *bool
}

type ListStreamTaskManagersResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of TaskManagers for the real-time task.
	Data []*ListStreamTaskManagersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend response exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListStreamTaskManagersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagersResponseBody) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListStreamTaskManagersResponseBody) GetData() []*ListStreamTaskManagersResponseBodyData {
	return s.Data
}

func (s *ListStreamTaskManagersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListStreamTaskManagersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStreamTaskManagersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStreamTaskManagersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStreamTaskManagersResponseBody) SetCode(v string) *ListStreamTaskManagersResponseBody {
	s.Code = &v
	return s
}

func (s *ListStreamTaskManagersResponseBody) SetData(v []*ListStreamTaskManagersResponseBodyData) *ListStreamTaskManagersResponseBody {
	s.Data = v
	return s
}

func (s *ListStreamTaskManagersResponseBody) SetHttpStatusCode(v int32) *ListStreamTaskManagersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListStreamTaskManagersResponseBody) SetMessage(v string) *ListStreamTaskManagersResponseBody {
	s.Message = &v
	return s
}

func (s *ListStreamTaskManagersResponseBody) SetRequestId(v string) *ListStreamTaskManagersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStreamTaskManagersResponseBody) SetSuccess(v bool) *ListStreamTaskManagersResponseBody {
	s.Success = &v
	return s
}

func (s *ListStreamTaskManagersResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStreamTaskManagersResponseBodyData struct {
	// The communication address or machine information.
	//
	// example:
	//
	// akka.tcp://flink@10.0.0.12:6122/user/rpc/taskmanager_0
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// taskmanagerId
	//
	// example:
	//
	// dataphin-session-7db678957dc449aab0086b931c9f8218-taskmanazjkxq
	TaskmanagerId *string `json:"TaskmanagerId,omitempty" xml:"TaskmanagerId,omitempty"`
}

func (s ListStreamTaskManagersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListStreamTaskManagersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStreamTaskManagersResponseBodyData) GetPath() *string {
	return s.Path
}

func (s *ListStreamTaskManagersResponseBodyData) GetTaskmanagerId() *string {
	return s.TaskmanagerId
}

func (s *ListStreamTaskManagersResponseBodyData) SetPath(v string) *ListStreamTaskManagersResponseBodyData {
	s.Path = &v
	return s
}

func (s *ListStreamTaskManagersResponseBodyData) SetTaskmanagerId(v string) *ListStreamTaskManagersResponseBodyData {
	s.TaskmanagerId = &v
	return s
}

func (s *ListStreamTaskManagersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
