// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredRemoteControllersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrand(v string) *ListInfraredRemoteControllersRequest
	GetBrand() *string
	SetCategory(v string) *ListInfraredRemoteControllersRequest
	GetCategory() *string
	SetCity(v string) *ListInfraredRemoteControllersRequest
	GetCity() *string
	SetHotelId(v string) *ListInfraredRemoteControllersRequest
	GetHotelId() *string
	SetProvince(v string) *ListInfraredRemoteControllersRequest
	GetProvince() *string
	SetServiceProvider(v string) *ListInfraredRemoteControllersRequest
	GetServiceProvider() *string
}

type ListInfraredRemoteControllersRequest struct {
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// This parameter is required.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af7***536
	HotelId         *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	Province        *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s ListInfraredRemoteControllersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredRemoteControllersRequest) GoString() string {
	return s.String()
}

func (s *ListInfraredRemoteControllersRequest) GetBrand() *string {
	return s.Brand
}

func (s *ListInfraredRemoteControllersRequest) GetCategory() *string {
	return s.Category
}

func (s *ListInfraredRemoteControllersRequest) GetCity() *string {
	return s.City
}

func (s *ListInfraredRemoteControllersRequest) GetHotelId() *string {
	return s.HotelId
}

func (s *ListInfraredRemoteControllersRequest) GetProvince() *string {
	return s.Province
}

func (s *ListInfraredRemoteControllersRequest) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *ListInfraredRemoteControllersRequest) SetBrand(v string) *ListInfraredRemoteControllersRequest {
	s.Brand = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetCategory(v string) *ListInfraredRemoteControllersRequest {
	s.Category = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetCity(v string) *ListInfraredRemoteControllersRequest {
	s.City = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetHotelId(v string) *ListInfraredRemoteControllersRequest {
	s.HotelId = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetProvince(v string) *ListInfraredRemoteControllersRequest {
	s.Province = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) SetServiceProvider(v string) *ListInfraredRemoteControllersRequest {
	s.ServiceProvider = &v
	return s
}

func (s *ListInfraredRemoteControllersRequest) Validate() error {
	return dara.Validate(s)
}
