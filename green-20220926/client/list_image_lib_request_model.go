// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListImageLibRequest
	GetRegionId() *string
}

type ListImageLibRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImageLibRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImageLibRequest) GoString() string {
	return s.String()
}

func (s *ListImageLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImageLibRequest) SetRegionId(v string) *ListImageLibRequest {
	s.RegionId = &v
	return s
}

func (s *ListImageLibRequest) Validate() error {
	return dara.Validate(s)
}
