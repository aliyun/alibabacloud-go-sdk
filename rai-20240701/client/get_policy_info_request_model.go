// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *GetPolicyInfoRequest
	GetPolicyId() *int64
	SetRegionId(v string) *GetPolicyInfoRequest
	GetRegionId() *string
}

type GetPolicyInfoRequest struct {
	// Detection policy ID
	//
	// example:
	//
	// 16
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPolicyInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyInfoRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *GetPolicyInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPolicyInfoRequest) SetPolicyId(v int64) *GetPolicyInfoRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyInfoRequest) SetRegionId(v string) *GetPolicyInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyInfoRequest) Validate() error {
	return dara.Validate(s)
}
