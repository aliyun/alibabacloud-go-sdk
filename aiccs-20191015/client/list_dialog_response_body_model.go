// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDialogResponseBody
	GetCode() *string
	SetData(v []*ListDialogResponseBodyData) *ListDialogResponseBody
	GetData() []*ListDialogResponseBodyData
	SetMessage(v string) *ListDialogResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDialogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDialogResponseBody
	GetSuccess() *bool
}

type ListDialogResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListDialogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDialogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDialogResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialogResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDialogResponseBody) GetData() []*ListDialogResponseBodyData {
	return s.Data
}

func (s *ListDialogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDialogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDialogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDialogResponseBody) SetCode(v string) *ListDialogResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialogResponseBody) SetData(v []*ListDialogResponseBodyData) *ListDialogResponseBody {
	s.Data = v
	return s
}

func (s *ListDialogResponseBody) SetMessage(v string) *ListDialogResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialogResponseBody) SetRequestId(v string) *ListDialogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialogResponseBody) SetSuccess(v bool) *ListDialogResponseBody {
	s.Success = &v
	return s
}

func (s *ListDialogResponseBody) Validate() error {
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

type ListDialogResponseBodyData struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// normal
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// robot
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Tag  *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 1619763900718
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListDialogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDialogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDialogResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *ListDialogResponseBodyData) GetNodeType() *string {
	return s.NodeType
}

func (s *ListDialogResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *ListDialogResponseBodyData) GetTag() *string {
	return s.Tag
}

func (s *ListDialogResponseBodyData) GetTime() *string {
	return s.Time
}

func (s *ListDialogResponseBodyData) SetContent(v string) *ListDialogResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListDialogResponseBodyData) SetNodeType(v string) *ListDialogResponseBodyData {
	s.NodeType = &v
	return s
}

func (s *ListDialogResponseBodyData) SetRole(v string) *ListDialogResponseBodyData {
	s.Role = &v
	return s
}

func (s *ListDialogResponseBodyData) SetTag(v string) *ListDialogResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListDialogResponseBodyData) SetTime(v string) *ListDialogResponseBodyData {
	s.Time = &v
	return s
}

func (s *ListDialogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
