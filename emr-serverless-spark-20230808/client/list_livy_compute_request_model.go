// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *ListLivyComputeRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListLivyComputeRequest
	GetRegionId() *string
}

type ListLivyComputeRequest struct {
	// example:
	//
	// ev-cq31c7tlhtgm9nrrlj4g
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListLivyComputeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeRequest) GoString() string {
	return s.String()
}

func (s *ListLivyComputeRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListLivyComputeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLivyComputeRequest) SetEnvironmentId(v string) *ListLivyComputeRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListLivyComputeRequest) SetRegionId(v string) *ListLivyComputeRequest {
	s.RegionId = &v
	return s
}

func (s *ListLivyComputeRequest) Validate() error {
	return dara.Validate(s)
}
