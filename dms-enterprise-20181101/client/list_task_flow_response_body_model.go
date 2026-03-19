// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTaskFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTaskFlowResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTaskFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskFlowResponseBody
	GetSuccess() *bool
	SetTaskFlowList(v *ListTaskFlowResponseBodyTaskFlowList) *ListTaskFlowResponseBody
	GetTaskFlowList() *ListTaskFlowResponseBodyTaskFlowList
}

type ListTaskFlowResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4116147E-C628-5816-8779-8EEAF8E973F4
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
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskFlowList *ListTaskFlowResponseBodyTaskFlowList `json:"TaskFlowList,omitempty" xml:"TaskFlowList,omitempty" type:"Struct"`
}

func (s ListTaskFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTaskFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTaskFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskFlowResponseBody) GetTaskFlowList() *ListTaskFlowResponseBodyTaskFlowList {
	return s.TaskFlowList
}

func (s *ListTaskFlowResponseBody) SetErrorCode(v string) *ListTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTaskFlowResponseBody) SetErrorMessage(v string) *ListTaskFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTaskFlowResponseBody) SetRequestId(v string) *ListTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskFlowResponseBody) SetSuccess(v bool) *ListTaskFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskFlowResponseBody) SetTaskFlowList(v *ListTaskFlowResponseBodyTaskFlowList) *ListTaskFlowResponseBody {
	s.TaskFlowList = v
	return s
}

func (s *ListTaskFlowResponseBody) Validate() error {
	if s.TaskFlowList != nil {
		if err := s.TaskFlowList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskFlowResponseBodyTaskFlowList struct {
	TaskFlow []*ListTaskFlowResponseBodyTaskFlowListTaskFlow `json:"TaskFlow,omitempty" xml:"TaskFlow,omitempty" type:"Repeated"`
}

func (s ListTaskFlowResponseBodyTaskFlowList) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowResponseBodyTaskFlowList) GoString() string {
	return s.String()
}

func (s *ListTaskFlowResponseBodyTaskFlowList) GetTaskFlow() []*ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	return s.TaskFlow
}

func (s *ListTaskFlowResponseBodyTaskFlowList) SetTaskFlow(v []*ListTaskFlowResponseBodyTaskFlowListTaskFlow) *ListTaskFlowResponseBodyTaskFlowList {
	s.TaskFlow = v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowList) Validate() error {
	if s.TaskFlow != nil {
		for _, item := range s.TaskFlow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskFlowResponseBodyTaskFlowListTaskFlow struct {
	CreatorId            *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorNickName      *string `json:"CreatorNickName,omitempty" xml:"CreatorNickName,omitempty"`
	DagOwnerNickName     *string `json:"DagOwnerNickName,omitempty" xml:"DagOwnerNickName,omitempty"`
	DeployId             *int64  `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	LatestInstanceStatus *int32  `json:"LatestInstanceStatus,omitempty" xml:"LatestInstanceStatus,omitempty"`
	LatestInstanceTime   *string `json:"LatestInstanceTime,omitempty" xml:"LatestInstanceTime,omitempty"`
	Status               *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTaskFlowResponseBodyTaskFlowListTaskFlow) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowResponseBodyTaskFlowListTaskFlow) GoString() string {
	return s.String()
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetCreatorNickName() *string {
	return s.CreatorNickName
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetDagOwnerNickName() *string {
	return s.DagOwnerNickName
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetDeployId() *int64 {
	return s.DeployId
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetId() *int64 {
	return s.Id
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetLatestInstanceStatus() *int32 {
	return s.LatestInstanceStatus
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetLatestInstanceTime() *string {
	return s.LatestInstanceTime
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) GetStatus() *int32 {
	return s.Status
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetCreatorId(v string) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.CreatorId = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetCreatorNickName(v string) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.CreatorNickName = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetDagOwnerNickName(v string) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.DagOwnerNickName = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetDeployId(v int64) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.DeployId = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetId(v int64) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.Id = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetLatestInstanceStatus(v int32) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.LatestInstanceStatus = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetLatestInstanceTime(v string) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.LatestInstanceTime = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) SetStatus(v int32) *ListTaskFlowResponseBodyTaskFlowListTaskFlow {
	s.Status = &v
	return s
}

func (s *ListTaskFlowResponseBodyTaskFlowListTaskFlow) Validate() error {
	return dara.Validate(s)
}
