// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsSwitchRuleCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *SdnsSwitchRuleCategoryRequest
	GetEnable() *bool
	SetId(v string) *SdnsSwitchRuleCategoryRequest
	GetId() *string
}

type SdnsSwitchRuleCategoryRequest struct {
	// This parameter is required.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// 123456789012345678
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SdnsSwitchRuleCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s SdnsSwitchRuleCategoryRequest) GoString() string {
	return s.String()
}

func (s *SdnsSwitchRuleCategoryRequest) GetEnable() *bool {
	return s.Enable
}

func (s *SdnsSwitchRuleCategoryRequest) GetId() *string {
	return s.Id
}

func (s *SdnsSwitchRuleCategoryRequest) SetEnable(v bool) *SdnsSwitchRuleCategoryRequest {
	s.Enable = &v
	return s
}

func (s *SdnsSwitchRuleCategoryRequest) SetId(v string) *SdnsSwitchRuleCategoryRequest {
	s.Id = &v
	return s
}

func (s *SdnsSwitchRuleCategoryRequest) Validate() error {
	return dara.Validate(s)
}
