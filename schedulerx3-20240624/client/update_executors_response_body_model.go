// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateExecutorsResponseBody
	GetCode() *int32
	SetData(v *UpdateExecutorsResponseBodyData) *UpdateExecutorsResponseBody
	GetData() *UpdateExecutorsResponseBodyData
	SetMessage(v string) *UpdateExecutorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateExecutorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateExecutorsResponseBody
	GetSuccess() *bool
}

type UpdateExecutorsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *UpdateExecutorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateExecutorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExecutorsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateExecutorsResponseBody) GetData() *UpdateExecutorsResponseBodyData {
	return s.Data
}

func (s *UpdateExecutorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateExecutorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExecutorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateExecutorsResponseBody) SetCode(v int32) *UpdateExecutorsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExecutorsResponseBody) SetData(v *UpdateExecutorsResponseBodyData) *UpdateExecutorsResponseBody {
	s.Data = v
	return s
}

func (s *UpdateExecutorsResponseBody) SetMessage(v string) *UpdateExecutorsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExecutorsResponseBody) SetRequestId(v string) *UpdateExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExecutorsResponseBody) SetSuccess(v bool) *UpdateExecutorsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateExecutorsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateExecutorsResponseBodyData struct {
	// example:
	//
	// 1
	AppGroupId *int64 `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	FailedService *string `json:"FailedService,omitempty" xml:"FailedService,omitempty"`
	// example:
	//
	// 1
	WorkId *int32 `json:"WorkId,omitempty" xml:"WorkId,omitempty"`
	// example:
	//
	// k8s_service
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s UpdateExecutorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateExecutorsResponseBodyData) GetAppGroupId() *int64 {
	return s.AppGroupId
}

func (s *UpdateExecutorsResponseBodyData) GetAppType() *int32 {
	return s.AppType
}

func (s *UpdateExecutorsResponseBodyData) GetFailedService() *string {
	return s.FailedService
}

func (s *UpdateExecutorsResponseBodyData) GetWorkId() *int32 {
	return s.WorkId
}

func (s *UpdateExecutorsResponseBodyData) GetWorkerType() *string {
	return s.WorkerType
}

func (s *UpdateExecutorsResponseBodyData) GetWorkers() *string {
	return s.Workers
}

func (s *UpdateExecutorsResponseBodyData) SetAppGroupId(v int64) *UpdateExecutorsResponseBodyData {
	s.AppGroupId = &v
	return s
}

func (s *UpdateExecutorsResponseBodyData) SetAppType(v int32) *UpdateExecutorsResponseBodyData {
	s.AppType = &v
	return s
}

func (s *UpdateExecutorsResponseBodyData) SetFailedService(v string) *UpdateExecutorsResponseBodyData {
	s.FailedService = &v
	return s
}

func (s *UpdateExecutorsResponseBodyData) SetWorkId(v int32) *UpdateExecutorsResponseBodyData {
	s.WorkId = &v
	return s
}

func (s *UpdateExecutorsResponseBodyData) SetWorkerType(v string) *UpdateExecutorsResponseBodyData {
	s.WorkerType = &v
	return s
}

func (s *UpdateExecutorsResponseBodyData) SetWorkers(v string) *UpdateExecutorsResponseBodyData {
	s.Workers = &v
	return s
}

func (s *UpdateExecutorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
