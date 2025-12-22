// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeedbacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFeedbacksRequest
	GetInstanceId() *string
	SetTaskIdList(v string) *ListFeedbacksRequest
	GetTaskIdList() *string
}

type ListFeedbacksRequest struct {
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["20251216-8B9B7B02-16FE-54BE-942A-F59DE0656032"]
	TaskIdList *string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty"`
}

func (s ListFeedbacksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFeedbacksRequest) GoString() string {
	return s.String()
}

func (s *ListFeedbacksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFeedbacksRequest) GetTaskIdList() *string {
	return s.TaskIdList
}

func (s *ListFeedbacksRequest) SetInstanceId(v string) *ListFeedbacksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFeedbacksRequest) SetTaskIdList(v string) *ListFeedbacksRequest {
	s.TaskIdList = &v
	return s
}

func (s *ListFeedbacksRequest) Validate() error {
	return dara.Validate(s)
}
