// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlackListStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackListStrategyShrink(v string) *ModifyBlackListStrategyShrinkRequest
	GetBlackListStrategyShrink() *string
	SetRegionId(v string) *ModifyBlackListStrategyShrinkRequest
	GetRegionId() *string
}

type ModifyBlackListStrategyShrinkRequest struct {
	// Blacklist rule.
	BlackListStrategyShrink *string `json:"BlackListStrategy,omitempty" xml:"BlackListStrategy,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyBlackListStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackListStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBlackListStrategyShrinkRequest) GetBlackListStrategyShrink() *string {
	return s.BlackListStrategyShrink
}

func (s *ModifyBlackListStrategyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBlackListStrategyShrinkRequest) SetBlackListStrategyShrink(v string) *ModifyBlackListStrategyShrinkRequest {
	s.BlackListStrategyShrink = &v
	return s
}

func (s *ModifyBlackListStrategyShrinkRequest) SetRegionId(v string) *ModifyBlackListStrategyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBlackListStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
