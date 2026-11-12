// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecutorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateExecutorGroupResponseBody
	GetCode() *int32
	SetData(v *CreateExecutorGroupResponseBodyData) *CreateExecutorGroupResponseBody
	GetData() *CreateExecutorGroupResponseBodyData
	SetMessage(v string) *CreateExecutorGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateExecutorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateExecutorGroupResponseBody
	GetSuccess() *bool
}

type CreateExecutorGroupResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateExecutorGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Parameter format error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
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

func (s CreateExecutorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExecutorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateExecutorGroupResponseBody) GetData() *CreateExecutorGroupResponseBodyData {
	return s.Data
}

func (s *CreateExecutorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateExecutorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExecutorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateExecutorGroupResponseBody) SetCode(v int32) *CreateExecutorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExecutorGroupResponseBody) SetData(v *CreateExecutorGroupResponseBodyData) *CreateExecutorGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateExecutorGroupResponseBody) SetMessage(v string) *CreateExecutorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExecutorGroupResponseBody) SetRequestId(v string) *CreateExecutorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExecutorGroupResponseBody) SetSuccess(v bool) *CreateExecutorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExecutorGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateExecutorGroupResponseBodyData struct {
	// The authentication method used by devices under the product to connect to the IoT platform. Valid values:
	//
	// - **secret**: Uses the device secret for device identity authentication.
	//
	// - **id2**: Uses IoT Internet Device ID (ID²).
	//
	// - **x509**: Uses the device X.509 certificate for device identity authentication.
	//
	// example:
	//
	// md5_spell
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
	// The domain name rule.
	//
	// example:
	//
	// NORMAL
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The protocol types supported by the domain name. Valid values: HTTP (supports only the HTTP protocol), HTTPS (supports only the HTTPS protocol), and HTTP,HTTPS (supports both HTTP and HTTPS protocols).
	//
	// example:
	//
	// All
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

func (s CreateExecutorGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateExecutorGroupResponseBodyData) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateExecutorGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateExecutorGroupResponseBodyData) GetFailedService() *string {
	return s.FailedService
}

func (s *CreateExecutorGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateExecutorGroupResponseBodyData) GetNetwork() *string {
	return s.Network
}

func (s *CreateExecutorGroupResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateExecutorGroupResponseBodyData) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *CreateExecutorGroupResponseBodyData) GetWorkerType() *string {
	return s.WorkerType
}

func (s *CreateExecutorGroupResponseBodyData) GetWorkers() *string {
	return s.Workers
}

func (s *CreateExecutorGroupResponseBodyData) SetAuthType(v string) *CreateExecutorGroupResponseBodyData {
	s.AuthType = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetDescription(v string) *CreateExecutorGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetFailedService(v string) *CreateExecutorGroupResponseBodyData {
	s.FailedService = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetName(v string) *CreateExecutorGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetNetwork(v string) *CreateExecutorGroupResponseBodyData {
	s.Network = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetProtocol(v string) *CreateExecutorGroupResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetWorkerId(v int64) *CreateExecutorGroupResponseBodyData {
	s.WorkerId = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetWorkerType(v string) *CreateExecutorGroupResponseBodyData {
	s.WorkerType = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) SetWorkers(v string) *CreateExecutorGroupResponseBodyData {
	s.Workers = &v
	return s
}

func (s *CreateExecutorGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
