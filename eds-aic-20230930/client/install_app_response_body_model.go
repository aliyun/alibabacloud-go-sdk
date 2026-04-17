// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChildTaskInfo(v []*InstallAppResponseBodyChildTaskInfo) *InstallAppResponseBody
	GetChildTaskInfo() []*InstallAppResponseBodyChildTaskInfo
	SetRequestId(v string) *InstallAppResponseBody
	GetRequestId() *string
	SetTaskId(v string) *InstallAppResponseBody
	GetTaskId() *string
}

type InstallAppResponseBody struct {
	ChildTaskInfo []*InstallAppResponseBodyChildTaskInfo `json:"ChildTaskInfo,omitempty" xml:"ChildTaskInfo,omitempty" type:"Repeated"`
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

func (s *InstallAppResponseBody) GetChildTaskInfo() []*InstallAppResponseBodyChildTaskInfo {
	return s.ChildTaskInfo
}

func (s *InstallAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAppResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallAppResponseBody) SetChildTaskInfo(v []*InstallAppResponseBodyChildTaskInfo) *InstallAppResponseBody {
	s.ChildTaskInfo = v
	return s
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

type InstallAppResponseBodyChildTaskInfo struct {
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

func (s InstallAppResponseBodyChildTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s InstallAppResponseBodyChildTaskInfo) GoString() string {
	return s.String()
}

func (s *InstallAppResponseBodyChildTaskInfo) GetAppId() *string {
	return s.AppId
}

func (s *InstallAppResponseBodyChildTaskInfo) GetChildTaskId() *string {
	return s.ChildTaskId
}

func (s *InstallAppResponseBodyChildTaskInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstallAppResponseBodyChildTaskInfo) SetAppId(v string) *InstallAppResponseBodyChildTaskInfo {
	s.AppId = &v
	return s
}

func (s *InstallAppResponseBodyChildTaskInfo) SetChildTaskId(v string) *InstallAppResponseBodyChildTaskInfo {
	s.ChildTaskId = &v
	return s
}

func (s *InstallAppResponseBodyChildTaskInfo) SetInstanceId(v string) *InstallAppResponseBodyChildTaskInfo {
	s.InstanceId = &v
	return s
}

func (s *InstallAppResponseBodyChildTaskInfo) Validate() error {
	return dara.Validate(s)
}
