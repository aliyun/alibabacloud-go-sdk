// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecycleBinConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDeleteEnabled(v bool) *RecycleBinConfig
	GetAutoDeleteEnabled() *bool
	SetAutoDeleteKeepSecond(v int32) *RecycleBinConfig
	GetAutoDeleteKeepSecond() *int32
	SetDeleteTrashNormalFileDisabled(v bool) *RecycleBinConfig
	GetDeleteTrashNormalFileDisabled() *bool
}

type RecycleBinConfig struct {
	AutoDeleteEnabled             *bool  `json:"auto_delete_enabled,omitempty" xml:"auto_delete_enabled,omitempty"`
	AutoDeleteKeepSecond          *int32 `json:"auto_delete_keep_second,omitempty" xml:"auto_delete_keep_second,omitempty"`
	DeleteTrashNormalFileDisabled *bool  `json:"delete_trash_normal_file_disabled,omitempty" xml:"delete_trash_normal_file_disabled,omitempty"`
}

func (s RecycleBinConfig) String() string {
	return dara.Prettify(s)
}

func (s RecycleBinConfig) GoString() string {
	return s.String()
}

func (s *RecycleBinConfig) GetAutoDeleteEnabled() *bool {
	return s.AutoDeleteEnabled
}

func (s *RecycleBinConfig) GetAutoDeleteKeepSecond() *int32 {
	return s.AutoDeleteKeepSecond
}

func (s *RecycleBinConfig) GetDeleteTrashNormalFileDisabled() *bool {
	return s.DeleteTrashNormalFileDisabled
}

func (s *RecycleBinConfig) SetAutoDeleteEnabled(v bool) *RecycleBinConfig {
	s.AutoDeleteEnabled = &v
	return s
}

func (s *RecycleBinConfig) SetAutoDeleteKeepSecond(v int32) *RecycleBinConfig {
	s.AutoDeleteKeepSecond = &v
	return s
}

func (s *RecycleBinConfig) SetDeleteTrashNormalFileDisabled(v bool) *RecycleBinConfig {
	s.DeleteTrashNormalFileDisabled = &v
	return s
}

func (s *RecycleBinConfig) Validate() error {
	return dara.Validate(s)
}
