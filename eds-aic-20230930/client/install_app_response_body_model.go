// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallAppResponseBody
	GetRequestId() *string
	SetTaskId(v string) *InstallAppResponseBody
	GetTaskId() *string
}

type InstallAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-14xwibw7yp73q****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InstallAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAppResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallAppResponseBody) SetRequestId(v string) *InstallAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAppResponseBody) SetTaskId(v string) *InstallAppResponseBody {
	s.TaskId = &v
	return s
}

func (s *InstallAppResponseBody) Validate() error {
	return dara.Validate(s)
}
