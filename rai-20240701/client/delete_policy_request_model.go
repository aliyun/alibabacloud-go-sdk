// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIdList(v []*int64) *DeletePolicyRequest
	GetPolicyIdList() []*int64
	SetRegionId(v string) *DeletePolicyRequest
	GetRegionId() *string
}

type DeletePolicyRequest struct {
	// List of detection policy IDs
	PolicyIdList []*int64 `json:"PolicyIdList,omitempty" xml:"PolicyIdList,omitempty" type:"Repeated"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetPolicyIdList() []*int64 {
	return s.PolicyIdList
}

func (s *DeletePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolicyRequest) SetPolicyIdList(v []*int64) *DeletePolicyRequest {
	s.PolicyIdList = v
	return s
}

func (s *DeletePolicyRequest) SetRegionId(v string) *DeletePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
