// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDetectionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetectionRuleId(v string) *DeleteDetectionRuleRequest
	GetDetectionRuleId() *string
	SetLang(v string) *DeleteDetectionRuleRequest
	GetLang() *string
	SetRegionId(v string) *DeleteDetectionRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteDetectionRuleRequest
	GetRoleFor() *int64
}

type DeleteDetectionRuleRequest struct {
	// This parameter is required.
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleFor         *int64  `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteDetectionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDetectionRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDetectionRuleRequest) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *DeleteDetectionRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDetectionRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDetectionRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteDetectionRuleRequest) SetDetectionRuleId(v string) *DeleteDetectionRuleRequest {
	s.DetectionRuleId = &v
	return s
}

func (s *DeleteDetectionRuleRequest) SetLang(v string) *DeleteDetectionRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDetectionRuleRequest) SetRegionId(v string) *DeleteDetectionRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDetectionRuleRequest) SetRoleFor(v int64) *DeleteDetectionRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteDetectionRuleRequest) Validate() error {
	return dara.Validate(s)
}
