// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRobotCallDialogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRobotCallDialogResponseBody
	GetCode() *string
	SetData(v []*ListRobotCallDialogResponseBodyData) *ListRobotCallDialogResponseBody
	GetData() []*ListRobotCallDialogResponseBodyData
	SetMessage(v string) *ListRobotCallDialogResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRobotCallDialogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRobotCallDialogResponseBody
	GetSuccess() *bool
}

type ListRobotCallDialogResponseBody struct {
	// Request status code. A return value of OK indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Conversation records.
	Data []*ListRobotCallDialogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRobotCallDialogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRobotCallDialogResponseBody) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRobotCallDialogResponseBody) GetData() []*ListRobotCallDialogResponseBodyData {
	return s.Data
}

func (s *ListRobotCallDialogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRobotCallDialogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRobotCallDialogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRobotCallDialogResponseBody) SetCode(v string) *ListRobotCallDialogResponseBody {
	s.Code = &v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetData(v []*ListRobotCallDialogResponseBodyData) *ListRobotCallDialogResponseBody {
	s.Data = v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetMessage(v string) *ListRobotCallDialogResponseBody {
	s.Message = &v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetRequestId(v string) *ListRobotCallDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRobotCallDialogResponseBody) SetSuccess(v bool) *ListRobotCallDialogResponseBody {
	s.Success = &v
	return s
}

func (s *ListRobotCallDialogResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRobotCallDialogResponseBodyData struct {
	// Script content.
	//
	// example:
	//
	// 我是某某的客服，看您之前在我们家找过工作，做个回访，现在工作怎么样？
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// File Type.
	//
	// example:
	//
	// 开场白
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// Role in the conversation content.
	//
	// example:
	//
	// robot
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Intent label.
	//
	// example:
	//
	// 拒绝
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// Start Time. UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1621483557000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListRobotCallDialogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRobotCallDialogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRobotCallDialogResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListRobotCallDialogResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *ListRobotCallDialogResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *ListRobotCallDialogResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ListRobotCallDialogResponseBodyData) GetTime() *string {
	return s.Time
}

func (s *ListRobotCallDialogResponseBodyData) SetContent(v string) *ListRobotCallDialogResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetNodeType(v string) *ListRobotCallDialogResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetRole(v string) *ListRobotCallDialogResponseBodyData {
	s.Role = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetTag(v string) *ListRobotCallDialogResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) SetTime(v string) *ListRobotCallDialogResponseBodyData {
	s.Time = &v
	return s
}

func (s *ListRobotCallDialogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
