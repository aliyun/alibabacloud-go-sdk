// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyThreatIntelligenceSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryList(v []*ModifyThreatIntelligenceSwitchRequestCategoryList) *ModifyThreatIntelligenceSwitchRequest
	GetCategoryList() []*ModifyThreatIntelligenceSwitchRequestCategoryList
}

type ModifyThreatIntelligenceSwitchRequest struct {
	CategoryList []*ModifyThreatIntelligenceSwitchRequestCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Repeated"`
}

func (s ModifyThreatIntelligenceSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyThreatIntelligenceSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyThreatIntelligenceSwitchRequest) GetCategoryList() []*ModifyThreatIntelligenceSwitchRequestCategoryList {
	return s.CategoryList
}

func (s *ModifyThreatIntelligenceSwitchRequest) SetCategoryList(v []*ModifyThreatIntelligenceSwitchRequestCategoryList) *ModifyThreatIntelligenceSwitchRequest {
	s.CategoryList = v
	return s
}

func (s *ModifyThreatIntelligenceSwitchRequest) Validate() error {
	if s.CategoryList != nil {
		for _, item := range s.CategoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyThreatIntelligenceSwitchRequestCategoryList struct {
	// example:
	//
	// alert
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// IpOutThreatTorExit
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 1
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
}

func (s ModifyThreatIntelligenceSwitchRequestCategoryList) String() string {
	return dara.Prettify(s)
}

func (s ModifyThreatIntelligenceSwitchRequestCategoryList) GoString() string {
	return s.String()
}

func (s *ModifyThreatIntelligenceSwitchRequestCategoryList) GetAction() *string {
	return s.Action
}

func (s *ModifyThreatIntelligenceSwitchRequestCategoryList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ModifyThreatIntelligenceSwitchRequestCategoryList) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ModifyThreatIntelligenceSwitchRequestCategoryList) SetAction(v string) *ModifyThreatIntelligenceSwitchRequestCategoryList {
	s.Action = &v
	return s
}

func (s *ModifyThreatIntelligenceSwitchRequestCategoryList) SetCategoryId(v string) *ModifyThreatIntelligenceSwitchRequestCategoryList {
	s.CategoryId = &v
	return s
}

func (s *ModifyThreatIntelligenceSwitchRequestCategoryList) SetEnableStatus(v string) *ModifyThreatIntelligenceSwitchRequestCategoryList {
	s.EnableStatus = &v
	return s
}

func (s *ModifyThreatIntelligenceSwitchRequestCategoryList) Validate() error {
	return dara.Validate(s)
}
