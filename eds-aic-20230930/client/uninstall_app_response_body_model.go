// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChildTaskInfo(v []*UninstallAppResponseBodyChildTaskInfo) *UninstallAppResponseBody
	GetChildTaskInfo() []*UninstallAppResponseBodyChildTaskInfo
	SetRequestId(v string) *UninstallAppResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UninstallAppResponseBody
	GetTaskId() *string
}

type UninstallAppResponseBody struct {
	ChildTaskInfo []*UninstallAppResponseBodyChildTaskInfo `json:"ChildTaskInfo,omitempty" xml:"ChildTaskInfo,omitempty" type:"Repeated"`
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
	// t-1ljew7on6ay0j****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAppResponseBody) GetChildTaskInfo() []*UninstallAppResponseBodyChildTaskInfo {
	return s.ChildTaskInfo
}

func (s *UninstallAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallAppResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallAppResponseBody) SetChildTaskInfo(v []*UninstallAppResponseBodyChildTaskInfo) *UninstallAppResponseBody {
	s.ChildTaskInfo = v
	return s
}

func (s *UninstallAppResponseBody) SetRequestId(v string) *UninstallAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAppResponseBody) SetTaskId(v string) *UninstallAppResponseBody {
	s.TaskId = &v
	return s
}

func (s *UninstallAppResponseBody) Validate() error {
	if s.ChildTaskInfo != nil {
		for _, item := range s.ChildTaskInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UninstallAppResponseBodyChildTaskInfo struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// t-ybde48cevxxxx
	ChildTaskId *string `json:"ChildTaskId,omitempty" xml:"ChildTaskId,omitempty"`
	// example:
	//
	// acp-ty3bnd7b9xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UninstallAppResponseBodyChildTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppResponseBodyChildTaskInfo) GoString() string {
	return s.String()
}

func (s *UninstallAppResponseBodyChildTaskInfo) GetAppId() *string {
	return s.AppId
}

func (s *UninstallAppResponseBodyChildTaskInfo) GetChildTaskId() *string {
	return s.ChildTaskId
}

func (s *UninstallAppResponseBodyChildTaskInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UninstallAppResponseBodyChildTaskInfo) SetAppId(v string) *UninstallAppResponseBodyChildTaskInfo {
	s.AppId = &v
	return s
}

func (s *UninstallAppResponseBodyChildTaskInfo) SetChildTaskId(v string) *UninstallAppResponseBodyChildTaskInfo {
	s.ChildTaskId = &v
	return s
}

func (s *UninstallAppResponseBodyChildTaskInfo) SetInstanceId(v string) *UninstallAppResponseBodyChildTaskInfo {
	s.InstanceId = &v
	return s
}

func (s *UninstallAppResponseBodyChildTaskInfo) Validate() error {
	return dara.Validate(s)
}
