// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCitiesByProvinceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProvince(v string) *ListCitiesByProvinceRequest
	GetProvince() *string
}

type ListCitiesByProvinceRequest struct {
	// This parameter is required.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListCitiesByProvinceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCitiesByProvinceRequest) GoString() string {
	return s.String()
}

func (s *ListCitiesByProvinceRequest) GetProvince() *string {
	return s.Province
}

func (s *ListCitiesByProvinceRequest) SetProvince(v string) *ListCitiesByProvinceRequest {
	s.Province = &v
	return s
}

func (s *ListCitiesByProvinceRequest) Validate() error {
	return dara.Validate(s)
}
