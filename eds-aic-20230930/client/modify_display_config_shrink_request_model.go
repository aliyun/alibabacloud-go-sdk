// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDisplayConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *ModifyDisplayConfigShrinkRequest
	GetAndroidInstanceIds() []*string
	SetDisplayConfigShrink(v string) *ModifyDisplayConfigShrinkRequest
	GetDisplayConfigShrink() *string
}

type ModifyDisplayConfigShrinkRequest struct {
	AndroidInstanceIds  []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	DisplayConfigShrink *string   `json:"DisplayConfig,omitempty" xml:"DisplayConfig,omitempty"`
}

func (s ModifyDisplayConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDisplayConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDisplayConfigShrinkRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *ModifyDisplayConfigShrinkRequest) GetDisplayConfigShrink() *string {
	return s.DisplayConfigShrink
}

func (s *ModifyDisplayConfigShrinkRequest) SetAndroidInstanceIds(v []*string) *ModifyDisplayConfigShrinkRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *ModifyDisplayConfigShrinkRequest) SetDisplayConfigShrink(v string) *ModifyDisplayConfigShrinkRequest {
	s.DisplayConfigShrink = &v
	return s
}

func (s *ModifyDisplayConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
