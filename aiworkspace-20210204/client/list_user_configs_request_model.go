// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryNames(v string) *ListUserConfigsRequest
	GetCategoryNames() *string
	SetConfigKeys(v string) *ListUserConfigsRequest
	GetConfigKeys() *string
}

type ListUserConfigsRequest struct {
	// The category. Currently, only DataPrivacyConfig is supported.
	//
	// example:
	//
	// DataPrivacyConfig
	CategoryNames *string `json:"CategoryNames,omitempty" xml:"CategoryNames,omitempty"`
	// The configuration item keys. Currently, only customizePAIAssumedRole is supported.
	//
	// example:
	//
	// customizePAIAssumedRole
	ConfigKeys *string `json:"ConfigKeys,omitempty" xml:"ConfigKeys,omitempty"`
}

func (s ListUserConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListUserConfigsRequest) GetCategoryNames() *string {
	return s.CategoryNames
}

func (s *ListUserConfigsRequest) GetConfigKeys() *string {
	return s.ConfigKeys
}

func (s *ListUserConfigsRequest) SetCategoryNames(v string) *ListUserConfigsRequest {
	s.CategoryNames = &v
	return s
}

func (s *ListUserConfigsRequest) SetConfigKeys(v string) *ListUserConfigsRequest {
	s.ConfigKeys = &v
	return s
}

func (s *ListUserConfigsRequest) Validate() error {
	return dara.Validate(s)
}
