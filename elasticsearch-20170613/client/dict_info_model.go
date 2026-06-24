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
	// The size of the dictionary file. Unit: bytes.
	//
	// example:
	//
	// 2782602
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The name of the dictionary file. Requirements:
	//
	// - Main dictionary or stopword list: one word per line, saved as a UTF-8 encoded DIC file. The file name can contain uppercase and lowercase letters, digits, and underscores, and cannot exceed 30 characters in length. Files with duplicate names are not allowed. The main dictionary file and the stopword file cannot share the same name.
	//
	// - Synonym dictionary: one synonym expression per line, saved as a UTF-8 encoded TXT file.
	//
	// - Alibaba dictionary: the file name must be aliws_ext_dict.txt. The file must be in UTF-8 format. Each line contains one word with no leading or trailing whitespace. Use UNIX or Linux line endings, where each line ends with
	//
	// . If the file is generated on a Windows system, use the dos2unix tool on a Linux machine to process the dictionary file before uploading it.
	//
	// example:
	//
	// aliws_ext_dict.txt
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source type of the dictionary file. Valid values:
	//
	// - OSS: Object Storage Service (OSS). Ensure that the OSS bucket has public-read permission.
	//
	// - ORIGIN: open-source Elasticsearch
	//
	// - UPLOAD: uploaded file.
	//
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The type of the dictionary file. Valid values:
	//
	// - STOP: stopword list
	//
	// - MAIN: main dictionary
	//
	// - SYNONYMS: synonym dictionary
	//
	// - ALI_WS: Alibaba dictionary.
	//
	// example:
	//
	// ALI_WS
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
