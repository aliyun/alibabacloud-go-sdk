// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMobileAgentPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPackageIds(v []*string) *DeleteMobileAgentPackageRequest
	GetPackageIds() []*string
}

type DeleteMobileAgentPackageRequest struct {
	// The list of packages.
	PackageIds []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
}

func (s DeleteMobileAgentPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMobileAgentPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteMobileAgentPackageRequest) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *DeleteMobileAgentPackageRequest) SetPackageIds(v []*string) *DeleteMobileAgentPackageRequest {
	s.PackageIds = v
	return s
}

func (s *DeleteMobileAgentPackageRequest) Validate() error {
	return dara.Validate(s)
}
