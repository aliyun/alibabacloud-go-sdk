// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDictInfo interface {
	dara.Model
	String() string
	GoString() string
	SetFileSize(v int64) *DictInfo
	GetFileSize() *int64
	SetName(v string) *DictInfo
	GetName() *string
	SetSourceType(v string) *DictInfo
	GetSourceType() *string
	SetType(v string) *DictInfo
	GetType() *string
}

type DictInfo struct {
	FileSize   *int64  `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DictInfo) String() string {
	return dara.Prettify(s)
}

func (s DictInfo) GoString() string {
	return s.String()
}

func (s *DictInfo) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DictInfo) GetName() *string {
	return s.Name
}

func (s *DictInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *DictInfo) GetType() *string {
	return s.Type
}

func (s *DictInfo) SetFileSize(v int64) *DictInfo {
	s.FileSize = &v
	return s
}

func (s *DictInfo) SetName(v string) *DictInfo {
	s.Name = &v
	return s
}

func (s *DictInfo) SetSourceType(v string) *DictInfo {
	s.SourceType = &v
	return s
}

func (s *DictInfo) SetType(v string) *DictInfo {
	s.Type = &v
	return s
}

func (s *DictInfo) Validate() error {
	return dara.Validate(s)
}
