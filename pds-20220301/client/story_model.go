// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStory interface {
	dara.Model
	String() string
	GoString() string
	SetCoverFileId(v string) *Story
	GetCoverFileId() *string
	SetCoverFileThumbnailUrl(v string) *Story
	GetCoverFileThumbnailUrl() *string
	SetCreatedAt(v string) *Story
	GetCreatedAt() *string
	SetCustomLabels(v map[string]interface{}) *Story
	GetCustomLabels() map[string]interface{}
	SetFaceGroupIds(v []*string) *Story
	GetFaceGroupIds() []*string
	SetStoryEndTime(v string) *Story
	GetStoryEndTime() *string
	SetStoryFileList(v []*File) *Story
	GetStoryFileList() []*File
	SetStoryId(v string) *Story
	GetStoryId() *string
	SetStoryName(v string) *Story
	GetStoryName() *string
	SetStoryStartTime(v string) *Story
	GetStoryStartTime() *string
	SetStorySubType(v string) *Story
	GetStorySubType() *string
	SetStoryType(v string) *Story
	GetStoryType() *string
	SetUpdatedAt(v string) *Story
	GetUpdatedAt() *string
}

type Story struct {
	// The ID of the story cover file.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	CoverFileId *string `json:"cover_file_id,omitempty" xml:"cover_file_id,omitempty"`
	// The URL of thumbnail of the story cover file.
	//
	// example:
	//
	// https://pds-domain-bucket.oss-cn-hangzhou.aliyuncs.com/Bh1HqdAs%2F1001%2F642a88c26f99cad589904fc8a6b2db6021c6601f%2F642a88c29391c47f221e49e88a21fa9b0a9e5063?security-token=CAIS%2BgF1q6Ft5B2yfSjIr5eGE8nGtYZ59bSnbUODlmsjTttVhpbFpjz2IHpPfHdoBe0btvU%2BlWxX6fwZlq5rR4QAXlDfNXD6XXO2qFHPWZHInuDox55m4cTXNAr%2BIhr%2F29CoEIedZdjBe%2FCrRknZnytou9XTfimjWFrXWv%2Fgy%2BQQDLItUxK%2FcCBNCfpPOwJms7V6D3bKMuu3OROY6Qi5TmgQ41En1DIlt%2FXuk5DCtkqB12eXkLFF%2B97DRbG%2FdNRpMZtFVNO44fd7bKKp0lQLsUMSqv8q0fEcqGaW4o7CWQJLnzyCMvvJ9OVDFyN0aKEnH7J%2Bq%2FzxhTPrMnpkSlacGoABVgD3mIAeUoDYKTJrFTPdmAPIrlUUQ1HZstB0ZVHvYByygr3oeitG9FEQ8lcfaZYJOtWRchdgH8tsSwHj2jI1vMeAiLHgbW98SPA3yfN8OR1kLW73uq3LAvekc9Qldz33YTM4JA5MNJzR8x5O6Oe4rY5OMSSVU5F%2Ffr8bJZD%2FXMM%3D&x-oss-access-key-id=STS.NT3XsrXYXBvLoe2rkvBTyiSnD&x-oss-expires=1685966251&x-oss-process=image%2Fresize%2Cw_500&x-oss-signature=XimW6haBxYktKg4cstMfpd%2F2LSoMibj6A%2B7ZqBJcSbQ%3D&x-oss-signature-version=OSS2
	CoverFileThumbnailUrl *string `json:"cover_file_thumbnail_url,omitempty" xml:"cover_file_thumbnail_url,omitempty"`
	// The time when the story was created. The time follows the RFC3339 standard.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The custom tags. You can specify key-value pairs based on your business requirements to search for stories by calling the FindStories operation.
	CustomLabels map[string]interface{} `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// The information about face-based groups. This parameter is valid only if story_type is set to PeopleMemory. This parameter is invalid for stories of other types or custom stories.
	FaceGroupIds []*string `json:"face_group_ids,omitempty" xml:"face_group_ids,omitempty" type:"Repeated"`
	// The time when the story ends. The time follows the RFC3339 standard.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	StoryEndTime *string `json:"story_end_time,omitempty" xml:"story_end_time,omitempty"`
	// The story files.
	StoryFileList []*File `json:"story_file_list,omitempty" xml:"story_file_list,omitempty" type:"Repeated"`
	// The story ID.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	// The name of the story.
	//
	// example:
	//
	// test_name
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
	// The time when the story starts. The time follows the RFC3339 standard.
	//
	// example:
	//
	// 2021-01-14T10:17:18.102700407+08:00
	StoryStartTime *string `json:"story_start_time,omitempty" xml:"story_start_time,omitempty"`
	// The subtype of the story. It is specified when the story is created.
	//
	// example:
	//
	// Solo
	StorySubType *string `json:"story_sub_type,omitempty" xml:"story_sub_type,omitempty"`
	// The type of the story. It is specified when the story is created.
	//
	// example:
	//
	// PeopleMemory
	StoryType *string `json:"story_type,omitempty" xml:"story_type,omitempty"`
	// The time when the story was updated. The time follows the RFC3339 standard.
	//
	// example:
	//
	// 2022-01-14T10:17:18.102700407+08:00
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s Story) String() string {
	return dara.Prettify(s)
}

func (s Story) GoString() string {
	return s.String()
}

func (s *Story) GetCoverFileId() *string {
	return s.CoverFileId
}

func (s *Story) GetCoverFileThumbnailUrl() *string {
	return s.CoverFileThumbnailUrl
}

func (s *Story) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Story) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *Story) GetFaceGroupIds() []*string {
	return s.FaceGroupIds
}

func (s *Story) GetStoryEndTime() *string {
	return s.StoryEndTime
}

func (s *Story) GetStoryFileList() []*File {
	return s.StoryFileList
}

func (s *Story) GetStoryId() *string {
	return s.StoryId
}

func (s *Story) GetStoryName() *string {
	return s.StoryName
}

func (s *Story) GetStoryStartTime() *string {
	return s.StoryStartTime
}

func (s *Story) GetStorySubType() *string {
	return s.StorySubType
}

func (s *Story) GetStoryType() *string {
	return s.StoryType
}

func (s *Story) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *Story) SetCoverFileId(v string) *Story {
	s.CoverFileId = &v
	return s
}

func (s *Story) SetCoverFileThumbnailUrl(v string) *Story {
	s.CoverFileThumbnailUrl = &v
	return s
}

func (s *Story) SetCreatedAt(v string) *Story {
	s.CreatedAt = &v
	return s
}

func (s *Story) SetCustomLabels(v map[string]interface{}) *Story {
	s.CustomLabels = v
	return s
}

func (s *Story) SetFaceGroupIds(v []*string) *Story {
	s.FaceGroupIds = v
	return s
}

func (s *Story) SetStoryEndTime(v string) *Story {
	s.StoryEndTime = &v
	return s
}

func (s *Story) SetStoryFileList(v []*File) *Story {
	s.StoryFileList = v
	return s
}

func (s *Story) SetStoryId(v string) *Story {
	s.StoryId = &v
	return s
}

func (s *Story) SetStoryName(v string) *Story {
	s.StoryName = &v
	return s
}

func (s *Story) SetStoryStartTime(v string) *Story {
	s.StoryStartTime = &v
	return s
}

func (s *Story) SetStorySubType(v string) *Story {
	s.StorySubType = &v
	return s
}

func (s *Story) SetStoryType(v string) *Story {
	s.StoryType = &v
	return s
}

func (s *Story) SetUpdatedAt(v string) *Story {
	s.UpdatedAt = &v
	return s
}

func (s *Story) Validate() error {
	if s.StoryFileList != nil {
		for _, item := range s.StoryFileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
