// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtension(v *GetExtensionResponseBodyExtension) *GetExtensionResponseBody
	GetExtension() *GetExtensionResponseBodyExtension
	SetRequestId(v string) *GetExtensionResponseBody
	GetRequestId() *string
}

type GetExtensionResponseBody struct {
	// The details of the extension.
	Extension *GetExtensionResponseBodyExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *GetExtensionResponseBody) GetExtension() *GetExtensionResponseBodyExtension {
	return s.Extension
}

func (s *GetExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExtensionResponseBody) SetExtension(v *GetExtensionResponseBodyExtension) *GetExtensionResponseBody {
	s.Extension = v
	return s
}

func (s *GetExtensionResponseBody) SetRequestId(v string) *GetExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetExtensionResponseBodyExtension struct {
	// The list of extension points.
	BindEventList []*GetExtensionResponseBodyExtensionBindEventList `json:"BindEventList,omitempty" xml:"BindEventList,omitempty" type:"Repeated"`
	// The URL of the extension details page, on which users can view the details of the process blocked by the extension.
	//
	// example:
	//
	// https://www.aliyun.com/
	DetailUrl *string `json:"DetailUrl,omitempty" xml:"DetailUrl,omitempty"`
	// The list of event types.
	EventCategoryList []*GetExtensionResponseBodyExtensionEventCategoryList `json:"EventCategoryList,omitempty" xml:"EventCategoryList,omitempty" type:"Repeated"`
	// The unique code of the extension.
	//
	// example:
	//
	// ce4*********086da5
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
	// The description of the extension.
	//
	// example:
	//
	// This is a description
	ExtensionDesc *string `json:"ExtensionDesc,omitempty" xml:"ExtensionDesc,omitempty"`
	// The name of the extension.
	//
	// example:
	//
	// Extension name
	ExtensionName *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
	// The URL of the help documentation of the extension.
	//
	// example:
	//
	// https://www.aliyun.com/
	HelpDocUrl *string `json:"HelpDocUrl,omitempty" xml:"HelpDocUrl,omitempty"`
	// The options defined for the extension.
	//
	// example:
	//
	// Option configuration
	OptionSetting *string `json:"OptionSetting,omitempty" xml:"OptionSetting,omitempty"`
	// The parameter settings of the extension. For more information, see [Configure extension parameters](https://help.aliyun.com/document_detail/405354.html).
	//
	// example:
	//
	// extension.project.disabled=projectId1,projectId2,projectId3
	ParameterSetting *string `json:"ParameterSetting,omitempty" xml:"ParameterSetting,omitempty"`
	// The workspace for testing. If the extension is being tested, the extension can be used only in the workspace for testing.
	//
	// example:
	//
	// 13552
	ProjectTesting *int64 `json:"ProjectTesting,omitempty" xml:"ProjectTesting,omitempty"`
	// The state of the extension. 0: Testing, 1: Publishing, 3: Disabled, 4: Processing, 5: Approved, 6: Approve Failed
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExtensionResponseBodyExtension) String() string {
	return dara.Prettify(s)
}

func (s GetExtensionResponseBodyExtension) GoString() string {
	return s.String()
}

func (s *GetExtensionResponseBodyExtension) GetBindEventList() []*GetExtensionResponseBodyExtensionBindEventList {
	return s.BindEventList
}

func (s *GetExtensionResponseBodyExtension) GetDetailUrl() *string {
	return s.DetailUrl
}

func (s *GetExtensionResponseBodyExtension) GetEventCategoryList() []*GetExtensionResponseBodyExtensionEventCategoryList {
	return s.EventCategoryList
}

func (s *GetExtensionResponseBodyExtension) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *GetExtensionResponseBodyExtension) GetExtensionDesc() *string {
	return s.ExtensionDesc
}

func (s *GetExtensionResponseBodyExtension) GetExtensionName() *string {
	return s.ExtensionName
}

func (s *GetExtensionResponseBodyExtension) GetHelpDocUrl() *string {
	return s.HelpDocUrl
}

func (s *GetExtensionResponseBodyExtension) GetOptionSetting() *string {
	return s.OptionSetting
}

func (s *GetExtensionResponseBodyExtension) GetParameterSetting() *string {
	return s.ParameterSetting
}

func (s *GetExtensionResponseBodyExtension) GetProjectTesting() *int64 {
	return s.ProjectTesting
}

func (s *GetExtensionResponseBodyExtension) GetStatus() *int32 {
	return s.Status
}

func (s *GetExtensionResponseBodyExtension) SetBindEventList(v []*GetExtensionResponseBodyExtensionBindEventList) *GetExtensionResponseBodyExtension {
	s.BindEventList = v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetDetailUrl(v string) *GetExtensionResponseBodyExtension {
	s.DetailUrl = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetEventCategoryList(v []*GetExtensionResponseBodyExtensionEventCategoryList) *GetExtensionResponseBodyExtension {
	s.EventCategoryList = v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetExtensionCode(v string) *GetExtensionResponseBodyExtension {
	s.ExtensionCode = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetExtensionDesc(v string) *GetExtensionResponseBodyExtension {
	s.ExtensionDesc = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetExtensionName(v string) *GetExtensionResponseBodyExtension {
	s.ExtensionName = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetHelpDocUrl(v string) *GetExtensionResponseBodyExtension {
	s.HelpDocUrl = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetOptionSetting(v string) *GetExtensionResponseBodyExtension {
	s.OptionSetting = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetParameterSetting(v string) *GetExtensionResponseBodyExtension {
	s.ParameterSetting = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetProjectTesting(v int64) *GetExtensionResponseBodyExtension {
	s.ProjectTesting = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) SetStatus(v int32) *GetExtensionResponseBodyExtension {
	s.Status = &v
	return s
}

func (s *GetExtensionResponseBodyExtension) Validate() error {
	return dara.Validate(s)
}

type GetExtensionResponseBodyExtensionBindEventList struct {
	// The code of the extension point event.
	//
	// example:
	//
	// commit-file
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The name of the extension point event.
	//
	// example:
	//
	// File submission pre-event
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
}

func (s GetExtensionResponseBodyExtensionBindEventList) String() string {
	return dara.Prettify(s)
}

func (s GetExtensionResponseBodyExtensionBindEventList) GoString() string {
	return s.String()
}

func (s *GetExtensionResponseBodyExtensionBindEventList) GetEventCode() *string {
	return s.EventCode
}

func (s *GetExtensionResponseBodyExtensionBindEventList) GetEventName() *string {
	return s.EventName
}

func (s *GetExtensionResponseBodyExtensionBindEventList) SetEventCode(v string) *GetExtensionResponseBodyExtensionBindEventList {
	s.EventCode = &v
	return s
}

func (s *GetExtensionResponseBodyExtensionBindEventList) SetEventName(v string) *GetExtensionResponseBodyExtensionBindEventList {
	s.EventName = &v
	return s
}

func (s *GetExtensionResponseBodyExtensionBindEventList) Validate() error {
	return dara.Validate(s)
}

type GetExtensionResponseBodyExtensionEventCategoryList struct {
	// The code of the event type.
	//
	// example:
	//
	// file-change
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// The name of the event type.
	//
	// example:
	//
	// File change event
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s GetExtensionResponseBodyExtensionEventCategoryList) String() string {
	return dara.Prettify(s)
}

func (s GetExtensionResponseBodyExtensionEventCategoryList) GoString() string {
	return s.String()
}

func (s *GetExtensionResponseBodyExtensionEventCategoryList) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *GetExtensionResponseBodyExtensionEventCategoryList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetExtensionResponseBodyExtensionEventCategoryList) SetCategoryCode(v string) *GetExtensionResponseBodyExtensionEventCategoryList {
	s.CategoryCode = &v
	return s
}

func (s *GetExtensionResponseBodyExtensionEventCategoryList) SetCategoryName(v string) *GetExtensionResponseBodyExtensionEventCategoryList {
	s.CategoryName = &v
	return s
}

func (s *GetExtensionResponseBodyExtensionEventCategoryList) Validate() error {
	return dara.Validate(s)
}
