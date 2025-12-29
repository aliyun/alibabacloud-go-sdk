// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSTBServiceProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *ListSTBServiceProvidersRequest
	GetCity() *string
	SetProvince(v string) *ListSTBServiceProvidersRequest
	GetProvince() *string
}

type ListSTBServiceProvidersRequest struct {
	// This parameter is required.
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// This parameter is required.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListSTBServiceProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSTBServiceProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersRequest) GetCity() *string {
	return s.City
}

func (s *ListSTBServiceProvidersRequest) GetProvince() *string {
	return s.Province
}

func (s *ListSTBServiceProvidersRequest) SetCity(v string) *ListSTBServiceProvidersRequest {
	s.City = &v
	return s
}

func (s *ListSTBServiceProvidersRequest) SetProvince(v string) *ListSTBServiceProvidersRequest {
	s.Province = &v
	return s
}

func (s *ListSTBServiceProvidersRequest) Validate() error {
	return dara.Validate(s)
}
