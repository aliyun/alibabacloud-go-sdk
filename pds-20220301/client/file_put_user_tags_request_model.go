// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilePutUserTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *FilePutUserTagsRequest
	GetDriveId() *string
	SetFileId(v string) *FilePutUserTagsRequest
	GetFileId() *string
	SetUserTags(v []*FilePutUserTagsRequestUserTags) *FilePutUserTagsRequest
	GetUserTags() []*FilePutUserTagsRequestUserTags
}

type FilePutUserTagsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The tags to be added to the file. You cannot leave this parameter empty. You can specify up to 1,000 tags. You cannot specify tags that have the same name.
	//
	// This parameter is required.
	UserTags []*FilePutUserTagsRequestUserTags `json:"user_tags,omitempty" xml:"user_tags,omitempty" type:"Repeated"`
}

func (s FilePutUserTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s FilePutUserTagsRequest) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *FilePutUserTagsRequest) GetFileId() *string {
	return s.FileId
}

func (s *FilePutUserTagsRequest) GetUserTags() []*FilePutUserTagsRequestUserTags {
	return s.UserTags
}

func (s *FilePutUserTagsRequest) SetDriveId(v string) *FilePutUserTagsRequest {
	s.DriveId = &v
	return s
}

func (s *FilePutUserTagsRequest) SetFileId(v string) *FilePutUserTagsRequest {
	s.FileId = &v
	return s
}

func (s *FilePutUserTagsRequest) SetUserTags(v []*FilePutUserTagsRequestUserTags) *FilePutUserTagsRequest {
	s.UserTags = v
	return s
}

func (s *FilePutUserTagsRequest) Validate() error {
	if s.UserTags != nil {
		for _, item := range s.UserTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FilePutUserTagsRequestUserTags struct {
	// The name of the tag. The tag name cannot be empty and cannot contain number signs (#).
	//
	// This parameter is required.
	//
	// example:
	//
	// tag
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag. The tag value cannot contain number signs (#).
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FilePutUserTagsRequestUserTags) String() string {
	return dara.Prettify(s)
}

func (s FilePutUserTagsRequestUserTags) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsRequestUserTags) GetKey() *string {
	return s.Key
}

func (s *FilePutUserTagsRequestUserTags) GetValue() *string {
	return s.Value
}

func (s *FilePutUserTagsRequestUserTags) SetKey(v string) *FilePutUserTagsRequestUserTags {
	s.Key = &v
	return s
}

func (s *FilePutUserTagsRequestUserTags) SetValue(v string) *FilePutUserTagsRequestUserTags {
	s.Value = &v
	return s
}

func (s *FilePutUserTagsRequestUserTags) Validate() error {
	return dara.Validate(s)
}
