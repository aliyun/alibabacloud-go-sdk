// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeApplicationResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpgradeApplicationResponseBody
	GetTaskId() *string
}

type UpgradeApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID. You can use the task ID to query the upgrade progress or status.
	//
	// example:
	//
	// 6f24a774-6bd5-4026-bb7d-deffb1dad875
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpgradeApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeApplicationResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeApplicationResponseBody) SetRequestId(v string) *UpgradeApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeApplicationResponseBody) SetTaskId(v string) *UpgradeApplicationResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
