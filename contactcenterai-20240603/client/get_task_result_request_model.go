// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequiredFieldList(v []*string) *GetTaskResultRequest
	GetRequiredFieldList() []*string
	SetTaskId(v string) *GetTaskResultRequest
	GetTaskId() *string
}

type GetTaskResultRequest struct {
	RequiredFieldList []*string `json:"requiredFieldList,omitempty" xml:"requiredFieldList,omitempty" type:"Repeated"`
	// example:
	//
	// 20240905-********-93E9-5D45-B4EF-045743A34071
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetTaskResultRequest) GetRequiredFieldList() []*string {
	return s.RequiredFieldList
}

func (s *GetTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResultRequest) SetRequiredFieldList(v []*string) *GetTaskResultRequest {
	s.RequiredFieldList = v
	return s
}

func (s *GetTaskResultRequest) SetTaskId(v string) *GetTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
