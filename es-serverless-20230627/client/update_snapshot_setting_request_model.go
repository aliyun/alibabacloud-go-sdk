// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSnapshotSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *UpdateSnapshotSettingRequest
	GetEnable() *bool
	SetQuartzRegex(v string) *UpdateSnapshotSettingRequest
	GetQuartzRegex() *string
}

type UpdateSnapshotSettingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0 0 01 ? 	- 	- *
	QuartzRegex *string `json:"quartzRegex,omitempty" xml:"quartzRegex,omitempty"`
}

func (s UpdateSnapshotSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSnapshotSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotSettingRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateSnapshotSettingRequest) GetQuartzRegex() *string {
	return s.QuartzRegex
}

func (s *UpdateSnapshotSettingRequest) SetEnable(v bool) *UpdateSnapshotSettingRequest {
	s.Enable = &v
	return s
}

func (s *UpdateSnapshotSettingRequest) SetQuartzRegex(v string) *UpdateSnapshotSettingRequest {
	s.QuartzRegex = &v
	return s
}

func (s *UpdateSnapshotSettingRequest) Validate() error {
	return dara.Validate(s)
}
