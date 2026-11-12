// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListPlansRequest
	GetRegionId() *string
}

type ListPlansRequest struct {
	// The region ID. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the available regions for simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPlansRequest) GoString() string {
	return s.String()
}

func (s *ListPlansRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPlansRequest) SetRegionId(v string) *ListPlansRequest {
	s.RegionId = &v
	return s
}

func (s *ListPlansRequest) Validate() error {
	return dara.Validate(s)
}
