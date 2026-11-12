// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildSandboxResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RebuildSandboxResponseBody
	GetCode() *string
	SetData(v *RebuildSandboxResponseBodyData) *RebuildSandboxResponseBody
	GetData() *RebuildSandboxResponseBodyData
	SetHttpStatusCode(v int32) *RebuildSandboxResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RebuildSandboxResponseBody
	GetMessage() *string
	SetRequestId(v string) *RebuildSandboxResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RebuildSandboxResponseBody
	GetSuccess() *bool
}

type RebuildSandboxResponseBody struct {
	// Task result code of this API request. A successful request returns SUCCESS. Query task result of the asynchronous task separately.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The acceptance result of the rebuild request.
	Data *RebuildSandboxResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code corresponding to this API request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result description of this API request. This does not indicate that the asynchronous rebuild task is complete.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID used to trace this API call.
	//
	// example:
	//
	// 01A1BE93-46A7-1A2A-BD62-159FCDAA****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether this API request is successful. A value of true does not indicate that the sandbox rebuild is complete.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RebuildSandboxResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebuildSandboxResponseBody) GoString() string {
	return s.String()
}

func (s *RebuildSandboxResponseBody) GetCode() *string {
	return s.Code
}

func (s *RebuildSandboxResponseBody) GetData() *RebuildSandboxResponseBodyData {
	return s.Data
}

func (s *RebuildSandboxResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RebuildSandboxResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RebuildSandboxResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebuildSandboxResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RebuildSandboxResponseBody) SetCode(v string) *RebuildSandboxResponseBody {
	s.Code = &v
	return s
}

func (s *RebuildSandboxResponseBody) SetData(v *RebuildSandboxResponseBodyData) *RebuildSandboxResponseBody {
	s.Data = v
	return s
}

func (s *RebuildSandboxResponseBody) SetHttpStatusCode(v int32) *RebuildSandboxResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RebuildSandboxResponseBody) SetMessage(v string) *RebuildSandboxResponseBody {
	s.Message = &v
	return s
}

func (s *RebuildSandboxResponseBody) SetRequestId(v string) *RebuildSandboxResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebuildSandboxResponseBody) SetSuccess(v bool) *RebuildSandboxResponseBody {
	s.Success = &v
	return s
}

func (s *RebuildSandboxResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RebuildSandboxResponseBodyData struct {
	// The target instance incarnation for this rebuild. The rebuild increases the incarnation. The returned value only indicates that the request is accepted, not that the new instance is ready.
	//
	// example:
	//
	// 2
	Incarnation *int64 `json:"incarnation,omitempty" xml:"incarnation,omitempty"`
	// The initial rebuild task ID. Use the operationId returned by RebuildSandbox. Always use this ID to query the overall result. You do not need to switch to the IDs of subsequently created tasks.
	//
	// example:
	//
	// 01a1be92-7fa4-76b2-a871-b376d6d0****
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The logical sandbox ID corresponding to this rebuild. The logical ID remains unchanged after the rebuild.
	//
	// example:
	//
	// 01a1b56a-b00f-7aa3-bacf-cc9d04b0****
	SandboxId *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
}

func (s RebuildSandboxResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RebuildSandboxResponseBodyData) GoString() string {
	return s.String()
}

func (s *RebuildSandboxResponseBodyData) GetIncarnation() *int64 {
	return s.Incarnation
}

func (s *RebuildSandboxResponseBodyData) GetOperationId() *string {
	return s.OperationId
}

func (s *RebuildSandboxResponseBodyData) GetSandboxId() *string {
	return s.SandboxId
}

func (s *RebuildSandboxResponseBodyData) SetIncarnation(v int64) *RebuildSandboxResponseBodyData {
	s.Incarnation = &v
	return s
}

func (s *RebuildSandboxResponseBodyData) SetOperationId(v string) *RebuildSandboxResponseBodyData {
	s.OperationId = &v
	return s
}

func (s *RebuildSandboxResponseBodyData) SetSandboxId(v string) *RebuildSandboxResponseBodyData {
	s.SandboxId = &v
	return s
}

func (s *RebuildSandboxResponseBodyData) Validate() error {
	return dara.Validate(s)
}
