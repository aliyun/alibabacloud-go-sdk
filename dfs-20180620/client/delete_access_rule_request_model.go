// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *DeleteAccessRuleRequest
	GetAccessGroupId() *string
	SetAccessRuleId(v string) *DeleteAccessRuleRequest
	GetAccessRuleId() *string
	SetInputRegionId(v string) *DeleteAccessRuleRequest
	GetInputRegionId() *string
}

type DeleteAccessRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acg-e3755fb0-358d-4286-9942-8d461048****
	AccessGroupId *string `json:"AccessGroupId,omitempty" xml:"AccessGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acr-c38028f0-f313-4385-9456-3501b1f5****
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
}

func (s DeleteAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *DeleteAccessRuleRequest) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *DeleteAccessRuleRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteAccessRuleRequest) SetAccessGroupId(v string) *DeleteAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetAccessRuleId(v string) *DeleteAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetInputRegionId(v string) *DeleteAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
