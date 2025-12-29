// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredDeviceBrandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListInfraredDeviceBrandsRequest
	GetCategory() *string
	SetServiceProvider(v string) *ListInfraredDeviceBrandsRequest
	GetServiceProvider() *string
}

type ListInfraredDeviceBrandsRequest struct {
	// This parameter is required.
	Category        *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ServiceProvider *string `json:"ServiceProvider,omitempty" xml:"ServiceProvider,omitempty"`
}

func (s ListInfraredDeviceBrandsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredDeviceBrandsRequest) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListInfraredDeviceBrandsRequest) GetServiceProvider() *string {
	return s.ServiceProvider
}

func (s *ListInfraredDeviceBrandsRequest) SetCategory(v string) *ListInfraredDeviceBrandsRequest {
	s.Category = &v
	return s
}

func (s *ListInfraredDeviceBrandsRequest) SetServiceProvider(v string) *ListInfraredDeviceBrandsRequest {
	s.ServiceProvider = &v
	return s
}

func (s *ListInfraredDeviceBrandsRequest) Validate() error {
	return dara.Validate(s)
}
