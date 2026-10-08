// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenGovernanceServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *OpenGovernanceServiceRequest
	GetRegionId() *string
}

type OpenGovernanceServiceRequest struct {
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenGovernanceServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenGovernanceServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenGovernanceServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenGovernanceServiceRequest) SetRegionId(v string) *OpenGovernanceServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenGovernanceServiceRequest) Validate() error {
	return dara.Validate(s)
}
