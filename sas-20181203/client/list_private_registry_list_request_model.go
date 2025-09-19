// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateRegistryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryType(v string) *ListPrivateRegistryListRequest
	GetRegistryType() *string
}

type ListPrivateRegistryListRequest struct {
	// The type of the image repository. Valid values:
	//
	// 	- **acr**: Container Registry
	//
	// 	- **harbor**: Harbor
	//
	// 	- **quay**: Quay
	//
	// 	- **CI/CD**: Jenkins
	//
	// example:
	//
	// harbor
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
}

func (s ListPrivateRegistryListRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateRegistryListRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateRegistryListRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *ListPrivateRegistryListRequest) SetRegistryType(v string) *ListPrivateRegistryListRequest {
	s.RegistryType = &v
	return s
}

func (s *ListPrivateRegistryListRequest) Validate() error {
	return dara.Validate(s)
}
