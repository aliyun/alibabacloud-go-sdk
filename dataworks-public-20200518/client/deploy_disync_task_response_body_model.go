// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDISyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeployDISyncTaskResponseBodyData) *DeployDISyncTaskResponseBody
	GetData() *DeployDISyncTaskResponseBodyData
	SetRequestId(v string) *DeployDISyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeployDISyncTaskResponseBody
	GetSuccess() *bool
}

type DeployDISyncTaskResponseBody struct {
	// Indicates whether the real-time synchronization node or data synchronization solution is deployed. Valid values:
	//
	// 	- success: The real-time synchronization node or data synchronization solution is deployed.
	//
	// 	- fail: The real-time synchronization node or data synchronization solution fails to be deployed.
	Data *DeployDISyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the request. You can query logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeployDISyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployDISyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeployDISyncTaskResponseBody) GetData() *DeployDISyncTaskResponseBodyData {
	return s.Data
}

func (s *DeployDISyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployDISyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeployDISyncTaskResponseBody) SetData(v *DeployDISyncTaskResponseBodyData) *DeployDISyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *DeployDISyncTaskResponseBody) SetRequestId(v string) *DeployDISyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployDISyncTaskResponseBody) SetSuccess(v bool) *DeployDISyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeployDISyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeployDISyncTaskResponseBodyData struct {
	// The reason for the failure to publish a data integration synchronization task.
	//
	// If the data integration synchronization task is published successfully, the return value of this parameter is empty.
	//
	// example:
	//
	// submit and deploy fail.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The cause of the failure to deploy the real-time synchronization node or data synchronization solution.
	//
	// If the real-time synchronization node or data synchronization solution is deployed, the value null is returned.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeployDISyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeployDISyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeployDISyncTaskResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeployDISyncTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeployDISyncTaskResponseBodyData) SetMessage(v string) *DeployDISyncTaskResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeployDISyncTaskResponseBodyData) SetStatus(v string) *DeployDISyncTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeployDISyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
