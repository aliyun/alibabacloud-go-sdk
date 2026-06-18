// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDepGroupTreeDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDepGroupTreeDataResponseBody
	GetCode() *string
	SetData(v []*GetDepGroupTreeDataResponseBodyData) *GetDepGroupTreeDataResponseBody
	GetData() []*GetDepGroupTreeDataResponseBodyData
	SetMessage(v string) *GetDepGroupTreeDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDepGroupTreeDataResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDepGroupTreeDataResponseBody
	GetSuccess() *string
}

type GetDepGroupTreeDataResponseBody struct {
	// The status code. A value of Success indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Department information.
	Data []*GetDepGroupTreeDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDepGroupTreeDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDepGroupTreeDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDepGroupTreeDataResponseBody) GetData() []*GetDepGroupTreeDataResponseBodyData {
	return s.Data
}

func (s *GetDepGroupTreeDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDepGroupTreeDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDepGroupTreeDataResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDepGroupTreeDataResponseBody) SetCode(v string) *GetDepGroupTreeDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetData(v []*GetDepGroupTreeDataResponseBodyData) *GetDepGroupTreeDataResponseBody {
	s.Data = v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetMessage(v string) *GetDepGroupTreeDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetRequestId(v string) *GetDepGroupTreeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) SetSuccess(v string) *GetDepGroupTreeDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBody) Validate() error {
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

type GetDepGroupTreeDataResponseBodyData struct {
	// The department ID.
	//
	// example:
	//
	// 10****
	DepGroupId *string `json:"DepGroupId,omitempty" xml:"DepGroupId,omitempty"`
	// The department name.
	//
	// example:
	//
	// 部门A
	DepGroupName *string `json:"DepGroupName,omitempty" xml:"DepGroupName,omitempty"`
	// Skill group data.
	GroupDTOS []*GetDepGroupTreeDataResponseBodyDataGroupDTOS `json:"GroupDTOS,omitempty" xml:"GroupDTOS,omitempty" type:"Repeated"`
}

func (s GetDepGroupTreeDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDepGroupTreeDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponseBodyData) GetDepGroupId() *string {
	return s.DepGroupId
}

func (s *GetDepGroupTreeDataResponseBodyData) GetDepGroupName() *string {
	return s.DepGroupName
}

func (s *GetDepGroupTreeDataResponseBodyData) GetGroupDTOS() []*GetDepGroupTreeDataResponseBodyDataGroupDTOS {
	return s.GroupDTOS
}

func (s *GetDepGroupTreeDataResponseBodyData) SetDepGroupId(v string) *GetDepGroupTreeDataResponseBodyData {
	s.DepGroupId = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyData) SetDepGroupName(v string) *GetDepGroupTreeDataResponseBodyData {
	s.DepGroupName = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyData) SetGroupDTOS(v []*GetDepGroupTreeDataResponseBodyDataGroupDTOS) *GetDepGroupTreeDataResponseBodyData {
	s.GroupDTOS = v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyData) Validate() error {
	if s.GroupDTOS != nil {
		for _, item := range s.GroupDTOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDepGroupTreeDataResponseBodyDataGroupDTOS struct {
	// The name of the skill group.
	//
	// example:
	//
	// 自动化技能组
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The skill group ID.
	//
	// example:
	//
	// 55****
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s GetDepGroupTreeDataResponseBodyDataGroupDTOS) String() string {
	return dara.Prettify(s)
}

func (s GetDepGroupTreeDataResponseBodyDataGroupDTOS) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataResponseBodyDataGroupDTOS) GetName() *string {
	return s.Name
}

func (s *GetDepGroupTreeDataResponseBodyDataGroupDTOS) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *GetDepGroupTreeDataResponseBodyDataGroupDTOS) SetName(v string) *GetDepGroupTreeDataResponseBodyDataGroupDTOS {
	s.Name = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyDataGroupDTOS) SetSkillGroupId(v int64) *GetDepGroupTreeDataResponseBodyDataGroupDTOS {
	s.SkillGroupId = &v
	return s
}

func (s *GetDepGroupTreeDataResponseBodyDataGroupDTOS) Validate() error {
	return dara.Validate(s)
}
