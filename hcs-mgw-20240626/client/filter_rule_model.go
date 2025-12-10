// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFilterRule interface {
	dara.Model
	String() string
	GoString() string
	SetFileTypeFilters(v *FileTypeFilters) *FilterRule
	GetFileTypeFilters() *FileTypeFilters
	SetKeyFilters(v *KeyFilters) *FilterRule
	GetKeyFilters() *KeyFilters
	SetLastModifiedFilters(v *LastModifiedFilters) *FilterRule
	GetLastModifiedFilters() *LastModifiedFilters
}

type FilterRule struct {
	FileTypeFilters     *FileTypeFilters     `json:"FileTypeFilters,omitempty" xml:"FileTypeFilters,omitempty"`
	KeyFilters          *KeyFilters          `json:"KeyFilters,omitempty" xml:"KeyFilters,omitempty"`
	LastModifiedFilters *LastModifiedFilters `json:"LastModifiedFilters,omitempty" xml:"LastModifiedFilters,omitempty"`
}

func (s FilterRule) String() string {
	return dara.Prettify(s)
}

func (s FilterRule) GoString() string {
	return s.String()
}

func (s *FilterRule) GetFileTypeFilters() *FileTypeFilters {
	return s.FileTypeFilters
}

func (s *FilterRule) GetKeyFilters() *KeyFilters {
	return s.KeyFilters
}

func (s *FilterRule) GetLastModifiedFilters() *LastModifiedFilters {
	return s.LastModifiedFilters
}

func (s *FilterRule) SetFileTypeFilters(v *FileTypeFilters) *FilterRule {
	s.FileTypeFilters = v
	return s
}

func (s *FilterRule) SetKeyFilters(v *KeyFilters) *FilterRule {
	s.KeyFilters = v
	return s
}

func (s *FilterRule) SetLastModifiedFilters(v *LastModifiedFilters) *FilterRule {
	s.LastModifiedFilters = v
	return s
}

func (s *FilterRule) Validate() error {
	if s.FileTypeFilters != nil {
		if err := s.FileTypeFilters.Validate(); err != nil {
			return err
		}
	}
	if s.KeyFilters != nil {
		if err := s.KeyFilters.Validate(); err != nil {
			return err
		}
	}
	if s.LastModifiedFilters != nil {
		if err := s.LastModifiedFilters.Validate(); err != nil {
			return err
		}
	}
	return nil
}
