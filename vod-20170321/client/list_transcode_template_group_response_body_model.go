// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeTemplateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTranscodeTemplateGroupResponseBody
	GetRequestId() *string
	SetTranscodeTemplateGroupList(v []*ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) *ListTranscodeTemplateGroupResponseBody
	GetTranscodeTemplateGroupList() []*ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList
}

type ListTranscodeTemplateGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The transcoding template groups.
	TranscodeTemplateGroupList []*ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList `json:"TranscodeTemplateGroupList,omitempty" xml:"TranscodeTemplateGroupList,omitempty" type:"Repeated"`
}

func (s ListTranscodeTemplateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTranscodeTemplateGroupResponseBody) GetTranscodeTemplateGroupList() []*ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	return s.TranscodeTemplateGroupList
}

func (s *ListTranscodeTemplateGroupResponseBody) SetRequestId(v string) *ListTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroupList(v []*ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) *ListTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroupList = v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList struct {
	// The ID of the application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the template group was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-05T10:20:09Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the template group is the default one. Valid values:
	//
	// 	- **Default**: The template group is the default one.
	//
	// 	- **NotDefault**: The template group is not the default one.
	//
	// example:
	//
	// Default
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The lock status of the transcoding template group. Valid values:
	//
	// 	- **Disabled**: The template group is not locked.
	//
	// 	- **Enabled**: The template group is locked.
	//
	// example:
	//
	// Disabled
	Locked *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The time when the template group was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-05T10:22:09Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the template group.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the transcoding template group.
	//
	// example:
	//
	// 17a9889fc66852*****d791c886700932
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GetAppId() *string {
	return s.AppId
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GetIsDefault() *string {
	return s.IsDefault
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GetLocked() *string {
	return s.Locked
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GetName() *string {
	return s.Name
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetAppId(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.AppId = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetCreationTime(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.CreationTime = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetIsDefault(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.IsDefault = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetLocked(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.Locked = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetModifyTime(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.ModifyTime = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetName(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.Name = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetTranscodeTemplateGroupId(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) Validate() error {
	return dara.Validate(s)
}
