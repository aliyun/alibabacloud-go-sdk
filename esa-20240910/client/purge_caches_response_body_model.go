// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurgeCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PurgeCachesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *PurgeCachesResponseBody
	GetTaskId() *string
}

type PurgeCachesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PurgeCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurgeCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PurgeCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurgeCachesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PurgeCachesResponseBody) SetRequestId(v string) *PurgeCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurgeCachesResponseBody) SetTaskId(v string) *PurgeCachesResponseBody {
	s.TaskId = &v
	return s
}

func (s *PurgeCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
