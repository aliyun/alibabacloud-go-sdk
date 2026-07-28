// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructDlfTagDto interface {
	dara.Model
	String() string
	GoString() string
	SetSnapshotId(v int64) *OpenStructDlfTagDto
	GetSnapshotId() *int64
	SetTagName(v string) *OpenStructDlfTagDto
	GetTagName() *string
	SetTimeMillis(v int64) *OpenStructDlfTagDto
	GetTimeMillis() *int64
	SetTotalRecordCount(v int64) *OpenStructDlfTagDto
	GetTotalRecordCount() *int64
}

type OpenStructDlfTagDto struct {
	// example:
	//
	// 123456789
	SnapshotId *int64 `json:"snapshotId,omitempty" xml:"snapshotId,omitempty"`
	// example:
	//
	// milvus-auto-20260101
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// example:
	//
	// 1711334400000
	TimeMillis *int64 `json:"timeMillis,omitempty" xml:"timeMillis,omitempty"`
	// example:
	//
	// 100000
	TotalRecordCount *int64 `json:"totalRecordCount,omitempty" xml:"totalRecordCount,omitempty"`
}

func (s OpenStructDlfTagDto) String() string {
	return dara.Prettify(s)
}

func (s OpenStructDlfTagDto) GoString() string {
	return s.String()
}

func (s *OpenStructDlfTagDto) GetSnapshotId() *int64 {
	return s.SnapshotId
}

func (s *OpenStructDlfTagDto) GetTagName() *string {
	return s.TagName
}

func (s *OpenStructDlfTagDto) GetTimeMillis() *int64 {
	return s.TimeMillis
}

func (s *OpenStructDlfTagDto) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *OpenStructDlfTagDto) SetSnapshotId(v int64) *OpenStructDlfTagDto {
	s.SnapshotId = &v
	return s
}

func (s *OpenStructDlfTagDto) SetTagName(v string) *OpenStructDlfTagDto {
	s.TagName = &v
	return s
}

func (s *OpenStructDlfTagDto) SetTimeMillis(v int64) *OpenStructDlfTagDto {
	s.TimeMillis = &v
	return s
}

func (s *OpenStructDlfTagDto) SetTotalRecordCount(v int64) *OpenStructDlfTagDto {
	s.TotalRecordCount = &v
	return s
}

func (s *OpenStructDlfTagDto) Validate() error {
	return dara.Validate(s)
}
