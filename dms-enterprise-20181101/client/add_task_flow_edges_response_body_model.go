// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskFlowEdgesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdgeIds(v *AddTaskFlowEdgesResponseBodyEdgeIds) *AddTaskFlowEdgesResponseBody
	GetEdgeIds() *AddTaskFlowEdgesResponseBodyEdgeIds
	SetErrorCode(v string) *AddTaskFlowEdgesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddTaskFlowEdgesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddTaskFlowEdgesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTaskFlowEdgesResponseBody
	GetSuccess() *bool
}

type AddTaskFlowEdgesResponseBody struct {
	// The list of task flow edge IDs.
	EdgeIds *AddTaskFlowEdgesResponseBodyEdgeIds `json:"EdgeIds,omitempty" xml:"EdgeIds,omitempty" type:"Struct"`
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
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// E5EE2B9E-2F95-57FA-B284-CB441CEE49D6
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

func (s AddTaskFlowEdgesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTaskFlowEdgesResponseBody) GoString() string {
	return s.String()
}

func (s *AddTaskFlowEdgesResponseBody) GetEdgeIds() *AddTaskFlowEdgesResponseBodyEdgeIds {
	return s.EdgeIds
}

func (s *AddTaskFlowEdgesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddTaskFlowEdgesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddTaskFlowEdgesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTaskFlowEdgesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTaskFlowEdgesResponseBody) SetEdgeIds(v *AddTaskFlowEdgesResponseBodyEdgeIds) *AddTaskFlowEdgesResponseBody {
	s.EdgeIds = v
	return s
}

func (s *AddTaskFlowEdgesResponseBody) SetErrorCode(v string) *AddTaskFlowEdgesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddTaskFlowEdgesResponseBody) SetErrorMessage(v string) *AddTaskFlowEdgesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddTaskFlowEdgesResponseBody) SetRequestId(v string) *AddTaskFlowEdgesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTaskFlowEdgesResponseBody) SetSuccess(v bool) *AddTaskFlowEdgesResponseBody {
	s.Success = &v
	return s
}

func (s *AddTaskFlowEdgesResponseBody) Validate() error {
	if s.EdgeIds != nil {
		if err := s.EdgeIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTaskFlowEdgesResponseBodyEdgeIds struct {
	EdgeId []*int64 `json:"EdgeId,omitempty" xml:"EdgeId,omitempty" type:"Repeated"`
}

func (s AddTaskFlowEdgesResponseBodyEdgeIds) String() string {
	return dara.Prettify(s)
}

func (s AddTaskFlowEdgesResponseBodyEdgeIds) GoString() string {
	return s.String()
}

func (s *AddTaskFlowEdgesResponseBodyEdgeIds) GetEdgeId() []*int64 {
	return s.EdgeId
}

func (s *AddTaskFlowEdgesResponseBodyEdgeIds) SetEdgeId(v []*int64) *AddTaskFlowEdgesResponseBodyEdgeIds {
	s.EdgeId = v
	return s
}

func (s *AddTaskFlowEdgesResponseBodyEdgeIds) Validate() error {
	return dara.Validate(s)
}
