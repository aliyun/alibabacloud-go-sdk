// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchGdnMemberRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SwitchGdnMemberRoleResponseBodyData) *SwitchGdnMemberRoleResponseBody
	GetData() *SwitchGdnMemberRoleResponseBodyData
	SetMessage(v string) *SwitchGdnMemberRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *SwitchGdnMemberRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SwitchGdnMemberRoleResponseBody
	GetSuccess() *bool
}

type SwitchGdnMemberRoleResponseBody struct {
	Data *SwitchGdnMemberRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SwitchGdnMemberRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchGdnMemberRoleResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleResponseBody) GetData() *SwitchGdnMemberRoleResponseBodyData {
	return s.Data
}

func (s *SwitchGdnMemberRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SwitchGdnMemberRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchGdnMemberRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SwitchGdnMemberRoleResponseBody) SetData(v *SwitchGdnMemberRoleResponseBodyData) *SwitchGdnMemberRoleResponseBody {
	s.Data = v
	return s
}

func (s *SwitchGdnMemberRoleResponseBody) SetMessage(v string) *SwitchGdnMemberRoleResponseBody {
	s.Message = &v
	return s
}

func (s *SwitchGdnMemberRoleResponseBody) SetRequestId(v string) *SwitchGdnMemberRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchGdnMemberRoleResponseBody) SetSuccess(v bool) *SwitchGdnMemberRoleResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchGdnMemberRoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type SwitchGdnMemberRoleResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SwitchGdnMemberRoleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SwitchGdnMemberRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *SwitchGdnMemberRoleResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *SwitchGdnMemberRoleResponseBodyData) SetTaskId(v int32) *SwitchGdnMemberRoleResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SwitchGdnMemberRoleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
