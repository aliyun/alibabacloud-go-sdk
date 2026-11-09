// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSandboxRebuildOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSandboxRebuildOperationResponseBody
	GetCode() *string
	SetData(v *GetSandboxRebuildOperationResponseBodyData) *GetSandboxRebuildOperationResponseBody
	GetData() *GetSandboxRebuildOperationResponseBodyData
	SetHttpStatusCode(v int32) *GetSandboxRebuildOperationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSandboxRebuildOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSandboxRebuildOperationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSandboxRebuildOperationResponseBody
	GetSuccess() *bool
}

type GetSandboxRebuildOperationResponseBody struct {
	// Task result code of the API request. A successful request returns SUCCESS. Query task result of the asynchronous task separately.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The overall result of the rebuild task.
	Data *GetSandboxRebuildOperationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code corresponding to the API request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result description of the API request. This does not indicate that the asynchronous rebuild task is completed.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to locate the API call.
	//
	// example:
	//
	// 01A1BE93-46A7-1A2A-BD62-159FCDAA****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API request is successful. A value of true does not indicate that the sandbox rebuild is completed.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSandboxRebuildOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSandboxRebuildOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSandboxRebuildOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSandboxRebuildOperationResponseBody) GetData() *GetSandboxRebuildOperationResponseBodyData {
	return s.Data
}

func (s *GetSandboxRebuildOperationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSandboxRebuildOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSandboxRebuildOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSandboxRebuildOperationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSandboxRebuildOperationResponseBody) SetCode(v string) *GetSandboxRebuildOperationResponseBody {
	s.Code = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBody) SetData(v *GetSandboxRebuildOperationResponseBodyData) *GetSandboxRebuildOperationResponseBody {
	s.Data = v
	return s
}

func (s *GetSandboxRebuildOperationResponseBody) SetHttpStatusCode(v int32) *GetSandboxRebuildOperationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBody) SetMessage(v string) *GetSandboxRebuildOperationResponseBody {
	s.Message = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBody) SetRequestId(v string) *GetSandboxRebuildOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBody) SetSuccess(v bool) *GetSandboxRebuildOperationResponseBody {
	s.Success = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSandboxRebuildOperationResponseBodyData struct {
	// The task error that can be publicly displayed. This is returned upon task failure, cancellation, or override by subsequent operations. This field can be empty in other states.
	Error *GetSandboxRebuildOperationResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The initial rebuild task ID. Use the operationId returned by RebuildSandbox. Always use this ID to query the overall result. You do not need to switch to the ID of a subsequent creation task.
	//
	// example:
	//
	// 01a1be92-7fa4-76b2-a871-b376d6d0****
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The logical sandbox ID corresponding to the rebuild. The logical ID remains unchanged after the rebuild.
	//
	// example:
	//
	// 01a1b56a-b00f-7aa3-bacf-cc9d04b0****
	SandboxId *string `json:"sandboxId,omitempty" xml:"sandboxId,omitempty"`
	// The overall rebuild status. Valid values: PENDING: waiting to be processed. RUNNING: execute. SUCCEEDED: overall success. FAILED: failed. CANCELED: canceled. SUPERSEDED: override by subsequent operations. UNKNOWN: unconfirmed. Only SUCCEEDED indicates completion. Do not treat UNKNOWN, missing status, or query exceptions as success, and do not automatically resubmit the rebuild.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetSandboxRebuildOperationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSandboxRebuildOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSandboxRebuildOperationResponseBodyData) GetError() *GetSandboxRebuildOperationResponseBodyDataError {
	return s.Error
}

func (s *GetSandboxRebuildOperationResponseBodyData) GetOperationId() *string {
	return s.OperationId
}

func (s *GetSandboxRebuildOperationResponseBodyData) GetSandboxId() *string {
	return s.SandboxId
}

func (s *GetSandboxRebuildOperationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSandboxRebuildOperationResponseBodyData) SetError(v *GetSandboxRebuildOperationResponseBodyDataError) *GetSandboxRebuildOperationResponseBodyData {
	s.Error = v
	return s
}

func (s *GetSandboxRebuildOperationResponseBodyData) SetOperationId(v string) *GetSandboxRebuildOperationResponseBodyData {
	s.OperationId = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBodyData) SetSandboxId(v string) *GetSandboxRebuildOperationResponseBodyData {
	s.SandboxId = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBodyData) SetStatus(v string) *GetSandboxRebuildOperationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBodyData) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSandboxRebuildOperationResponseBodyDataError struct {
	// The task error code, which is used to distinguish between rebuild failures, subsequent instance creation failures, cancellations, or supersessions.
	//
	// example:
	//
	// SandboxRebuild.Failed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The task error description that can be publicly displayed. Keep the task ID for troubleshooting.
	//
	// example:
	//
	// The Sandbox rebuild did not complete. Contact support with the operation ID.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetSandboxRebuildOperationResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s GetSandboxRebuildOperationResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *GetSandboxRebuildOperationResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *GetSandboxRebuildOperationResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *GetSandboxRebuildOperationResponseBodyDataError) SetCode(v string) *GetSandboxRebuildOperationResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBodyDataError) SetMessage(v string) *GetSandboxRebuildOperationResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *GetSandboxRebuildOperationResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}
