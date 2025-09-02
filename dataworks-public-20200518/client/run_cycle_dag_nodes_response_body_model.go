// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCycleDagNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *RunCycleDagNodesResponseBody
	GetData() []*int64
	SetErrorCode(v string) *RunCycleDagNodesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RunCycleDagNodesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *RunCycleDagNodesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RunCycleDagNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunCycleDagNodesResponseBody
	GetSuccess() *bool
}

type RunCycleDagNodesResponseBody struct {
	// The IDs of the nodes in the workflow. You can query instances based on the IDs.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// >E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunCycleDagNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCycleDagNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RunCycleDagNodesResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *RunCycleDagNodesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunCycleDagNodesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunCycleDagNodesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RunCycleDagNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCycleDagNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunCycleDagNodesResponseBody) SetData(v []*int64) *RunCycleDagNodesResponseBody {
	s.Data = v
	return s
}

func (s *RunCycleDagNodesResponseBody) SetErrorCode(v string) *RunCycleDagNodesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RunCycleDagNodesResponseBody) SetErrorMessage(v string) *RunCycleDagNodesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RunCycleDagNodesResponseBody) SetHttpStatusCode(v int32) *RunCycleDagNodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunCycleDagNodesResponseBody) SetRequestId(v string) *RunCycleDagNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCycleDagNodesResponseBody) SetSuccess(v bool) *RunCycleDagNodesResponseBody {
	s.Success = &v
	return s
}

func (s *RunCycleDagNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
