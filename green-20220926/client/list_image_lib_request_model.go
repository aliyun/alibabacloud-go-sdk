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
	SetServiceCode(v string) *ListImageLibRequest
	GetServiceCode() *string
}

type ListImageLibRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
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

func (s *ListImageLibRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListImageLibRequest) SetRegionId(v string) *ListImageLibRequest {
	s.RegionId = &v
	return s
}

func (s *ListImageLibRequest) SetServiceCode(v string) *ListImageLibRequest {
	s.ServiceCode = &v
	return s
}

func (s *ListImageLibRequest) Validate() error {
	return dara.Validate(s)
}
