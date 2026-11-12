// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateExecutorGroupResponseBody
	GetCode() *int32
	SetData(v *UpdateExecutorGroupResponseBodyData) *UpdateExecutorGroupResponseBody
	GetData() *UpdateExecutorGroupResponseBodyData
	SetMessage(v string) *UpdateExecutorGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateExecutorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateExecutorGroupResponseBody
	GetSuccess() *bool
}

type UpdateExecutorGroupResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The updated executor group information.
	Data *UpdateExecutorGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C78E2AD2-5985-515B-BAD2-31A248AFC263
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateExecutorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExecutorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateExecutorGroupResponseBody) GetData() *UpdateExecutorGroupResponseBodyData {
	return s.Data
}

func (s *UpdateExecutorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateExecutorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExecutorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateExecutorGroupResponseBody) SetCode(v int32) *UpdateExecutorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExecutorGroupResponseBody) SetData(v *UpdateExecutorGroupResponseBodyData) *UpdateExecutorGroupResponseBody {
	s.Data = v
	return s
}

func (s *UpdateExecutorGroupResponseBody) SetMessage(v string) *UpdateExecutorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExecutorGroupResponseBody) SetRequestId(v string) *UpdateExecutorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExecutorGroupResponseBody) SetSuccess(v bool) *UpdateExecutorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateExecutorGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateExecutorGroupResponseBodyData struct {
	// The authentication type. Valid values:
	//
	// - **Anonymous**: No authentication is required.
	//
	// - **Application**: Authentication is required.
	//
	// > The default value is **Anonymous**.
	//
	// example:
	//
	// EDIT
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The workflow description.
	//
	// example:
	//
	// my first workflow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Kubernetes services that failed to be imported.
	//
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	FailedService *string `json:"FailedService,omitempty" xml:"FailedService,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// myWorkflow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network data.
	//
	// example:
	//
	// {\\"OfficeSiteId\\":\\"cn-hangzhou+dir-8801187585\\",\\"VSwitchIds\\":[null]}
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The protocol types supported by the domain name. HTTP: Only the HTTP protocol is supported. HTTPS: Only the HTTPS protocol is supported. HTTP,HTTPS: Both HTTP and HTTPS protocols are supported.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The worker ID of the deployment task. You can obtain this ID by calling the [ListWorkerResource](https://help.aliyun.com/document_detail/2712224.html) operation.
	//
	// example:
	//
	// 12
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
	// The executor type.
	//
	// example:
	//
	// k8s_service
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// The list of specified machines in JSON format.
	//
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s UpdateExecutorGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateExecutorGroupResponseBodyData) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateExecutorGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateExecutorGroupResponseBodyData) GetFailedService() *string {
	return s.FailedService
}

func (s *UpdateExecutorGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateExecutorGroupResponseBodyData) GetNetwork() *string {
	return s.Network
}

func (s *UpdateExecutorGroupResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateExecutorGroupResponseBodyData) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *UpdateExecutorGroupResponseBodyData) GetWorkerType() *string {
	return s.WorkerType
}

func (s *UpdateExecutorGroupResponseBodyData) GetWorkers() *string {
	return s.Workers
}

func (s *UpdateExecutorGroupResponseBodyData) SetAuthType(v string) *UpdateExecutorGroupResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetDescription(v string) *UpdateExecutorGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetFailedService(v string) *UpdateExecutorGroupResponseBodyData {
	s.FailedService = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetName(v string) *UpdateExecutorGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetNetwork(v string) *UpdateExecutorGroupResponseBodyData {
	s.Network = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetProtocol(v string) *UpdateExecutorGroupResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetWorkerId(v int64) *UpdateExecutorGroupResponseBodyData {
	s.WorkerId = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetWorkerType(v string) *UpdateExecutorGroupResponseBodyData {
	s.WorkerType = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) SetWorkers(v string) *UpdateExecutorGroupResponseBodyData {
	s.Workers = &v
	return s
}

func (s *UpdateExecutorGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
