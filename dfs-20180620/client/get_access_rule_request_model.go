// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupId(v string) *GetAccessRuleRequest
	GetAccessGroupId() *string
	SetAccessRuleId(v string) *GetAccessRuleRequest
	GetAccessRuleId() *string
	SetInputRegionId(v string) *GetAccessRuleRequest
	GetInputRegionId() *string
}

type GetAccessRuleRequest struct {
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

func (s GetAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAccessRuleRequest) GetAccessGroupId() *string {
	return s.AccessGroupId
}

func (s *GetAccessRuleRequest) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *GetAccessRuleRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *GetAccessRuleRequest) SetAccessGroupId(v string) *GetAccessRuleRequest {
	s.AccessGroupId = &v
	return s
}

func (s *GetAccessRuleRequest) SetAccessRuleId(v string) *GetAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *GetAccessRuleRequest) SetInputRegionId(v string) *GetAccessRuleRequest {
	s.InputRegionId = &v
	return s
}

func (s *GetAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
