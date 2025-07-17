// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBackFillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDagInstanceId(v int64) *BackFillResponseBody
	GetDagInstanceId() *int64
	SetErrorCode(v string) *BackFillResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BackFillResponseBody
	GetErrorMessage() *string
	SetNodeId(v int64) *BackFillResponseBody
	GetNodeId() *int64
	SetRequestId(v string) *BackFillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BackFillResponseBody
	GetSuccess() *bool
}

type BackFillResponseBody struct {
	// The ID of the execution record of the task flow.
	//
	// example:
	//
	// 47****
	DagInstanceId *int64 `json:"DagInstanceId,omitempty" xml:"DagInstanceId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 43****
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BackFillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BackFillResponseBody) GoString() string {
	return s.String()
}

func (s *BackFillResponseBody) GetDagInstanceId() *int64 {
	return s.DagInstanceId
}

func (s *BackFillResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BackFillResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BackFillResponseBody) GetNodeId() *int64 {
	return s.NodeId
}

func (s *BackFillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BackFillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BackFillResponseBody) SetDagInstanceId(v int64) *BackFillResponseBody {
	s.DagInstanceId = &v
	return s
}

func (s *BackFillResponseBody) SetErrorCode(v string) *BackFillResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BackFillResponseBody) SetErrorMessage(v string) *BackFillResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BackFillResponseBody) SetNodeId(v int64) *BackFillResponseBody {
	s.NodeId = &v
	return s
}

func (s *BackFillResponseBody) SetRequestId(v string) *BackFillResponseBody {
	s.RequestId = &v
	return s
}

func (s *BackFillResponseBody) SetSuccess(v bool) *BackFillResponseBody {
	s.Success = &v
	return s
}

func (s *BackFillResponseBody) Validate() error {
	return dara.Validate(s)
}
