// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallNodePoolComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InstallNodePoolComponentsResponseBody
	GetClusterId() *string
	SetRequestId(v string) *InstallNodePoolComponentsResponseBody
	GetRequestId() *string
	SetTaskId(v string) *InstallNodePoolComponentsResponseBody
	GetTaskId() *string
}

type InstallNodePoolComponentsResponseBody struct {
	// example:
	//
	// c8155823d057948c69a****
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// T-67d7ec016ce37c0106000***
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s InstallNodePoolComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallNodePoolComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallNodePoolComponentsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallNodePoolComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallNodePoolComponentsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallNodePoolComponentsResponseBody) SetClusterId(v string) *InstallNodePoolComponentsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *InstallNodePoolComponentsResponseBody) SetRequestId(v string) *InstallNodePoolComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallNodePoolComponentsResponseBody) SetTaskId(v string) *InstallNodePoolComponentsResponseBody {
	s.TaskId = &v
	return s
}

func (s *InstallNodePoolComponentsResponseBody) Validate() error {
	return dara.Validate(s)
}
