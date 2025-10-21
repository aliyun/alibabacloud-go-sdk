// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyDefaultOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetPolicyDefaultOptionsRequest
	GetRegionId() *string
}

type GetPolicyDefaultOptionsRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolicyDefaultOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyDefaultOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyDefaultOptionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyDefaultOptionsRequest) SetRegionId(v string) *GetPolicyDefaultOptionsRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyDefaultOptionsRequest) Validate() error {
	return dara.Validate(s)
}
