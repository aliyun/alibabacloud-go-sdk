// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMPUTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMPUTaskStatusResponseBody
	GetRequestId() *string
	SetStatus(v int32) *GetMPUTaskStatusResponseBody
	GetStatus() *int32
}

type GetMPUTaskStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status ID of the bypass transcoding task. Valid values:
	//
	//  - **0**: Waiting for the channel to start.
	//
	// > If the task is not started within 30 days, the task is automatically stopped.
	//
	//  - **1**: The task is running.
	//
	//  - **2**: The task is stopped.
	//
	//  - **3**: The task is stopped by the user.
	//
	//  - **4**: The channel is stopped.
	//
	//  - **5**: Live streaming is stopped due to CDN network issues.
	//
	//  - **6**: Live streaming is stopped due to live streaming URL issues.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMPUTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMPUTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMPUTaskStatusResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetMPUTaskStatusResponseBody) SetRequestId(v string) *GetMPUTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMPUTaskStatusResponseBody) SetStatus(v int32) *GetMPUTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetMPUTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
