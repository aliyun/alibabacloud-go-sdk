// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNameCheckResult interface {
	dara.Model
	String() string
	GoString() string
	SetExistFileId(v string) *NameCheckResult
	GetExistFileId() *string
	SetExistFileType(v string) *NameCheckResult
	GetExistFileType() *string
}

type NameCheckResult struct {
	ExistFileId   *string `json:"exist_file_id,omitempty" xml:"exist_file_id,omitempty"`
	ExistFileType *string `json:"exist_file_type,omitempty" xml:"exist_file_type,omitempty"`
}

func (s NameCheckResult) String() string {
	return dara.Prettify(s)
}

func (s NameCheckResult) GoString() string {
	return s.String()
}

func (s *NameCheckResult) GetExistFileId() *string {
	return s.ExistFileId
}

func (s *NameCheckResult) GetExistFileType() *string {
	return s.ExistFileType
}

func (s *NameCheckResult) SetExistFileId(v string) *NameCheckResult {
	s.ExistFileId = &v
	return s
}

func (s *NameCheckResult) SetExistFileType(v string) *NameCheckResult {
	s.ExistFileType = &v
	return s
}

func (s *NameCheckResult) Validate() error {
	return dara.Validate(s)
}
