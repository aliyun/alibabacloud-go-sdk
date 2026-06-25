// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecutorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateExecutorsResponseBody
	GetCode() *int32
	SetData(v *CreateExecutorsResponseBodyData) *CreateExecutorsResponseBody
	GetData() *CreateExecutorsResponseBodyData
	SetMessage(v string) *CreateExecutorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateExecutorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExecutorsResponseBody
	GetSuccess() *bool
}

type CreateExecutorsResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateExecutorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// F131C3E0-3FAA-5FA4-A6F3-E974D69EF3C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExecutorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExecutorsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateExecutorsResponseBody) GetData() *CreateExecutorsResponseBodyData {
	return s.Data
}

func (s *CreateExecutorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExecutorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExecutorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExecutorsResponseBody) SetCode(v int32) *CreateExecutorsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExecutorsResponseBody) SetData(v *CreateExecutorsResponseBodyData) *CreateExecutorsResponseBody {
	s.Data = v
	return s
}

func (s *CreateExecutorsResponseBody) SetMessage(v string) *CreateExecutorsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExecutorsResponseBody) SetRequestId(v string) *CreateExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExecutorsResponseBody) SetSuccess(v bool) *CreateExecutorsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExecutorsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExecutorsResponseBodyData struct {
	// The App Group ID.
	//
	// example:
	//
	// test-app
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The application type.
	//
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// A list of Kubernetes Services that failed to import.
	//
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	FailedService *string `json:"FailedService,omitempty" xml:"FailedService,omitempty"`
	// The ID of the worker. You can obtain this ID by calling the [ListWorkerResource](https://help.aliyun.com/document_detail/2712224.html) operation.
	//
	// example:
	//
	// ins-95mfvqrtg6fkijt1uko000
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// The worker type.
	//
	// example:
	//
	// k8s_service
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// A JSON string that represents the list of imported workers.
	//
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s CreateExecutorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateExecutorsResponseBodyData) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *CreateExecutorsResponseBodyData) GetAppType() *int32 {
	return s.AppType
}

func (s *CreateExecutorsResponseBodyData) GetFailedService() *string {
	return s.FailedService
}

func (s *CreateExecutorsResponseBodyData) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *CreateExecutorsResponseBodyData) GetWorkerType() *string {
	return s.WorkerType
}

func (s *CreateExecutorsResponseBodyData) GetWorkers() *string {
	return s.Workers
}

func (s *CreateExecutorsResponseBodyData) SetAppGroupId(v int64) *CreateExecutorsResponseBodyData {
	s.AppGroupId = &v
	return s
}

func (s *CreateExecutorsResponseBodyData) SetAppType(v int32) *CreateExecutorsResponseBodyData {
	s.AppType = &v
	return s
}

func (s *CreateExecutorsResponseBodyData) SetFailedService(v string) *CreateExecutorsResponseBodyData {
	s.FailedService = &v
	return s
}

func (s *CreateExecutorsResponseBodyData) SetWorkerId(v int64) *CreateExecutorsResponseBodyData {
	s.WorkerId = &v
	return s
}

func (s *CreateExecutorsResponseBodyData) SetWorkerType(v string) *CreateExecutorsResponseBodyData {
	s.WorkerType = &v
	return s
}

func (s *CreateExecutorsResponseBodyData) SetWorkers(v string) *CreateExecutorsResponseBodyData {
	s.Workers = &v
	return s
}

func (s *CreateExecutorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
