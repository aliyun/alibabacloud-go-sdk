// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryDetailResponseBody
	GetCode() *string
	SetGroups(v []*GetMemoryDetailResponseBodyGroups) *GetMemoryDetailResponseBody
	GetGroups() []*GetMemoryDetailResponseBodyGroups
	SetMessage(v string) *GetMemoryDetailResponseBody
	GetMessage() *string
	SetProfile(v *GetMemoryDetailResponseBodyProfile) *GetMemoryDetailResponseBody
	GetProfile() *GetMemoryDetailResponseBodyProfile
	SetRequestId(v string) *GetMemoryDetailResponseBody
	GetRequestId() *string
}

type GetMemoryDetailResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of groups for automatically collected memories. If no automatic memories exist, an empty array is returned.
	//
	// example:
	//
	// [{"type":"user","items":[{"id":"user/example.md","title":"Sample preference","content":"Prefers concise answers","description":"Preference summary from recent conversations"}],"gmtModified":1720000000}]
	Groups []*GetMemoryDetailResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// The error description.
	//
	// example:
	//
	// The current user is not an active member of the group cg_12345
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The manually added memory from profile.md. It includes the full content, existence status, and modification time.
	//
	// example:
	//
	// {"content":"Sample manual memory","exists":true,"gmtModified":1720000000}
	Profile *GetMemoryDetailResponseBodyProfile `json:"profile,omitempty" xml:"profile,omitempty" type:"Struct"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMemoryDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryDetailResponseBody) GetGroups() []*GetMemoryDetailResponseBodyGroups {
	return s.Groups
}

func (s *GetMemoryDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryDetailResponseBody) GetProfile() *GetMemoryDetailResponseBodyProfile {
	return s.Profile
}

func (s *GetMemoryDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryDetailResponseBody) SetCode(v string) *GetMemoryDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryDetailResponseBody) SetGroups(v []*GetMemoryDetailResponseBodyGroups) *GetMemoryDetailResponseBody {
	s.Groups = v
	return s
}

func (s *GetMemoryDetailResponseBody) SetMessage(v string) *GetMemoryDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryDetailResponseBody) SetProfile(v *GetMemoryDetailResponseBodyProfile) *GetMemoryDetailResponseBody {
	s.Profile = v
	return s
}

func (s *GetMemoryDetailResponseBody) SetRequestId(v string) *GetMemoryDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryDetailResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Profile != nil {
		if err := s.Profile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryDetailResponseBodyGroups struct {
	// The last modification time of the group. The value is a UNIX timestamp in seconds. If no files exist, null is returned.
	//
	// example:
	//
	// 1720000000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The list of automatic memory entries under this type. If no entries exist, an empty array is returned.
	//
	// example:
	//
	// [{"id":"user/example.md","title":"Sample preference","content":"Prefers concise answers","description":"Preference summary from recent conversations"}]
	Items []*GetMemoryDetailResponseBodyGroupsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The type of the automatic memory. Valid values: user, feedback, reference, and project.
	//
	// example:
	//
	// user
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMemoryDetailResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBodyGroups) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetMemoryDetailResponseBodyGroups) GetItems() []*GetMemoryDetailResponseBodyGroupsItems {
	return s.Items
}

func (s *GetMemoryDetailResponseBodyGroups) GetType() *string {
	return s.Type
}

func (s *GetMemoryDetailResponseBodyGroups) SetGmtModified(v int64) *GetMemoryDetailResponseBodyGroups {
	s.GmtModified = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroups) SetItems(v []*GetMemoryDetailResponseBodyGroupsItems) *GetMemoryDetailResponseBodyGroups {
	s.Items = v
	return s
}

func (s *GetMemoryDetailResponseBodyGroups) SetType(v string) *GetMemoryDetailResponseBodyGroups {
	s.Type = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroups) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMemoryDetailResponseBodyGroupsItems struct {
	// The body of the automatic memory, excluding the YAML frontmatter.
	//
	// example:
	//
	// Prefers concise answers
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The additional description of the automatic memory. If no description is available, an empty string is returned.
	//
	// example:
	//
	// Preference summary from recent conversations
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The unique identifier of the automatic memory entry, which is a relative OSS path.
	//
	// example:
	//
	// user/example.md
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The title of the automatic memory, which is derived from the file metadata or file name.
	//
	// example:
	//
	// Sample preference
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetMemoryDetailResponseBodyGroupsItems) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBodyGroupsItems) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetContent() *string {
	return s.Content
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetDescription() *string {
	return s.Description
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetId() *string {
	return s.Id
}

func (s *GetMemoryDetailResponseBodyGroupsItems) GetTitle() *string {
	return s.Title
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetContent(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Content = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetDescription(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Description = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetId(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Id = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) SetTitle(v string) *GetMemoryDetailResponseBodyGroupsItems {
	s.Title = &v
	return s
}

func (s *GetMemoryDetailResponseBodyGroupsItems) Validate() error {
	return dara.Validate(s)
}

type GetMemoryDetailResponseBodyProfile struct {
	// The full body of the manual memory. If the memory has never been saved, an empty string is returned.
	//
	// example:
	//
	// Sample manual memory
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Indicates whether profile.md exists. If the memory has never been saved, false is returned.
	//
	// example:
	//
	// true
	Exists *bool `json:"exists,omitempty" xml:"exists,omitempty"`
	// The last modification time of profile.md. The value is a UNIX timestamp in seconds. If the file does not exist or its metadata cannot be obtained, null is returned.
	//
	// example:
	//
	// 1720000000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
}

func (s GetMemoryDetailResponseBodyProfile) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryDetailResponseBodyProfile) GoString() string {
	return s.String()
}

func (s *GetMemoryDetailResponseBodyProfile) GetContent() *string {
	return s.Content
}

func (s *GetMemoryDetailResponseBodyProfile) GetExists() *bool {
	return s.Exists
}

func (s *GetMemoryDetailResponseBodyProfile) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetMemoryDetailResponseBodyProfile) SetContent(v string) *GetMemoryDetailResponseBodyProfile {
	s.Content = &v
	return s
}

func (s *GetMemoryDetailResponseBodyProfile) SetExists(v bool) *GetMemoryDetailResponseBodyProfile {
	s.Exists = &v
	return s
}

func (s *GetMemoryDetailResponseBodyProfile) SetGmtModified(v int64) *GetMemoryDetailResponseBodyProfile {
	s.GmtModified = &v
	return s
}

func (s *GetMemoryDetailResponseBodyProfile) Validate() error {
	return dara.Validate(s)
}
