// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountFactoryBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v string) *DeleteAccountFactoryBaselineRequest
	GetBaselineId() *string
	SetRegionId(v string) *DeleteAccountFactoryBaselineRequest
	GetRegionId() *string
}

type DeleteAccountFactoryBaselineRequest struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAccountFactoryBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountFactoryBaselineRequest) GetBaselineId() *string {
	return s.BaselineId
}

func (s *DeleteAccountFactoryBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAccountFactoryBaselineRequest) SetBaselineId(v string) *DeleteAccountFactoryBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *DeleteAccountFactoryBaselineRequest) SetRegionId(v string) *DeleteAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccountFactoryBaselineRequest) Validate() error {
	return dara.Validate(s)
}
