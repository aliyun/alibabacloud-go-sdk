// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSkillGroupResponseBody
	GetCode() *string
	SetData(v []*ListSkillGroupResponseBodyData) *ListSkillGroupResponseBody
	GetData() []*ListSkillGroupResponseBodyData
	SetMessage(v string) *ListSkillGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSkillGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSkillGroupResponseBody
	GetSuccess() *bool
}

type ListSkillGroupResponseBody struct {
	// Status code. A value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Skill group information.
	Data []*ListSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
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

func (s ListSkillGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSkillGroupResponseBody) GetData() []*ListSkillGroupResponseBodyData {
	return s.Data
}

func (s *ListSkillGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSkillGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSkillGroupResponseBody) SetCode(v string) *ListSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetData(v []*ListSkillGroupResponseBodyData) *ListSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupResponseBody) SetMessage(v string) *ListSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetRequestId(v string) *ListSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetSuccess(v bool) *ListSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListSkillGroupResponseBody) Validate() error {
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

type ListSkillGroupResponseBodyData struct {
	// Channel type of the skill group.
	//
	// example:
	//
	// 2
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Skill group description.
	//
	// example:
	//
	// 自动化技能组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Display name of the skill group.
	//
	// example:
	//
	// 自动化技能组
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Name of the skill group.
	//
	// example:
	//
	// 自动化技能组
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Skill group ID.
	//
	// example:
	//
	// 123456
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListSkillGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBodyData) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *ListSkillGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListSkillGroupResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListSkillGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSkillGroupResponseBodyData) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *ListSkillGroupResponseBodyData) SetChannelType(v int32) *ListSkillGroupResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetDescription(v string) *ListSkillGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetDisplayName(v string) *ListSkillGroupResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetName(v string) *ListSkillGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetSkillGroupId(v int64) *ListSkillGroupResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
