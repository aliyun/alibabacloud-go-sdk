// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountFactoryBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineId(v string) *GetAccountFactoryBaselineRequest
	GetBaselineId() *string
	SetRegionId(v string) *GetAccountFactoryBaselineRequest
	GetRegionId() *string
}

type GetAccountFactoryBaselineRequest struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1nf0enuzb89az*****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAccountFactoryBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineRequest) GetBaselineId() *string {
	return s.BaselineId
}

func (s *GetAccountFactoryBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAccountFactoryBaselineRequest) SetBaselineId(v string) *GetAccountFactoryBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *GetAccountFactoryBaselineRequest) SetRegionId(v string) *GetAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *GetAccountFactoryBaselineRequest) Validate() error {
	return dara.Validate(s)
}
