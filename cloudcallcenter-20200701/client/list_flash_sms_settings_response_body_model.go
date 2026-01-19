// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlashSmsSettingsResponseBody
	GetCode() *string
	SetData(v *ListFlashSmsSettingsResponseBodyData) *ListFlashSmsSettingsResponseBody
	GetData() *ListFlashSmsSettingsResponseBodyData
	SetHttpStatusCode(v int32) *ListFlashSmsSettingsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFlashSmsSettingsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListFlashSmsSettingsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListFlashSmsSettingsResponseBody
	GetRequestId() *string
}

type ListFlashSmsSettingsResponseBody struct {
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListFlashSmsSettingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// FCEFE806-E67C-577E-865B-4ED398F2F648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlashSmsSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlashSmsSettingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFlashSmsSettingsResponseBody) GetData() *ListFlashSmsSettingsResponseBodyData {
	return s.Data
}

func (s *ListFlashSmsSettingsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFlashSmsSettingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlashSmsSettingsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListFlashSmsSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlashSmsSettingsResponseBody) SetCode(v string) *ListFlashSmsSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBody) SetData(v *ListFlashSmsSettingsResponseBodyData) *ListFlashSmsSettingsResponseBody {
	s.Data = v
	return s
}

func (s *ListFlashSmsSettingsResponseBody) SetHttpStatusCode(v int32) *ListFlashSmsSettingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBody) SetMessage(v string) *ListFlashSmsSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBody) SetParams(v []*string) *ListFlashSmsSettingsResponseBody {
	s.Params = v
	return s
}

func (s *ListFlashSmsSettingsResponseBody) SetRequestId(v string) *ListFlashSmsSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFlashSmsSettingsResponseBodyData struct {
	List []*ListFlashSmsSettingsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlashSmsSettingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsSettingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlashSmsSettingsResponseBodyData) GetList() []*ListFlashSmsSettingsResponseBodyDataList {
	return s.List
}

func (s *ListFlashSmsSettingsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsSettingsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsSettingsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFlashSmsSettingsResponseBodyData) SetList(v []*ListFlashSmsSettingsResponseBodyDataList) *ListFlashSmsSettingsResponseBodyData {
	s.List = v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyData) SetPageNumber(v int32) *ListFlashSmsSettingsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyData) SetPageSize(v int32) *ListFlashSmsSettingsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyData) SetTotalCount(v int32) *ListFlashSmsSettingsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFlashSmsSettingsResponseBodyDataList struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 57fd969d-1936-4879-baaf-4bc57b3caef0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// skg1@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// skg
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListFlashSmsSettingsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsSettingsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFlashSmsSettingsResponseBodyDataList) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListFlashSmsSettingsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsSettingsResponseBodyDataList) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListFlashSmsSettingsResponseBodyDataList) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListFlashSmsSettingsResponseBodyDataList) SetEnabled(v bool) *ListFlashSmsSettingsResponseBodyDataList {
	s.Enabled = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyDataList) SetInstanceId(v string) *ListFlashSmsSettingsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyDataList) SetSkillGroupId(v string) *ListFlashSmsSettingsResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyDataList) SetSkillGroupName(v string) *ListFlashSmsSettingsResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListFlashSmsSettingsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
