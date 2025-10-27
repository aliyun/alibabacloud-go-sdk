// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterVipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterVipResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDBClusterVipResponseBody
	GetTaskId() *string
}

type ModifyDBClusterVipResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-****-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1564657730
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDBClusterVipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterVipResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterVipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterVipResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDBClusterVipResponseBody) SetRequestId(v string) *ModifyDBClusterVipResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterVipResponseBody) SetTaskId(v string) *ModifyDBClusterVipResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDBClusterVipResponseBody) Validate() error {
	return dara.Validate(s)
}
