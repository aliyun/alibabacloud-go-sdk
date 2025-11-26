// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceFilterSetting interface {
	dara.Model
	String() string
	GoString() string
	SetTagSelector(v *FilterSetting) *WorkspaceFilterSetting
	GetTagSelector() *FilterSetting
	SetWorkspaceUuids(v []*string) *WorkspaceFilterSetting
	GetWorkspaceUuids() []*string
}

type WorkspaceFilterSetting struct {
	TagSelector    *FilterSetting `json:"tagSelector,omitempty" xml:"tagSelector,omitempty"`
	WorkspaceUuids []*string      `json:"workspaceUuids,omitempty" xml:"workspaceUuids,omitempty" type:"Repeated"`
}

func (s WorkspaceFilterSetting) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceFilterSetting) GoString() string {
	return s.String()
}

func (s *WorkspaceFilterSetting) GetTagSelector() *FilterSetting {
	return s.TagSelector
}

func (s *WorkspaceFilterSetting) GetWorkspaceUuids() []*string {
	return s.WorkspaceUuids
}

func (s *WorkspaceFilterSetting) SetTagSelector(v *FilterSetting) *WorkspaceFilterSetting {
	s.TagSelector = v
	return s
}

func (s *WorkspaceFilterSetting) SetWorkspaceUuids(v []*string) *WorkspaceFilterSetting {
	s.WorkspaceUuids = v
	return s
}

func (s *WorkspaceFilterSetting) Validate() error {
	if s.TagSelector != nil {
		if err := s.TagSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
