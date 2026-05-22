// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BlockObjectResponseBody
	GetRequestId() *string
	SetTaskId(v string) *BlockObjectResponseBody
	GetTaskId() *string
}

type BlockObjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BlockObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BlockObjectResponseBody) GoString() string {
	return s.String()
}

func (s *BlockObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BlockObjectResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *BlockObjectResponseBody) SetRequestId(v string) *BlockObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlockObjectResponseBody) SetTaskId(v string) *BlockObjectResponseBody {
	s.TaskId = &v
	return s
}

func (s *BlockObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
