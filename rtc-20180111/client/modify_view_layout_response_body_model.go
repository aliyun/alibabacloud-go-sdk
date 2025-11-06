// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyViewLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyViewLayoutResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyViewLayoutResponseBody
	GetTaskId() *string
}

type ModifyViewLayoutResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyViewLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyViewLayoutResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyViewLayoutResponseBody) SetRequestId(v string) *ModifyViewLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyViewLayoutResponseBody) SetTaskId(v string) *ModifyViewLayoutResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyViewLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
