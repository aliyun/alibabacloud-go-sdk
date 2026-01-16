// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *DescribeAddonRequest
	GetAddonName() *string
	SetAddonVersion(v string) *DescribeAddonRequest
	GetAddonVersion() *string
}

type DescribeAddonRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// edge-csi-lite
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// example:
	//
	// v1
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
}

func (s DescribeAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *DescribeAddonRequest) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *DescribeAddonRequest) SetAddonName(v string) *DescribeAddonRequest {
	s.AddonName = &v
	return s
}

func (s *DescribeAddonRequest) SetAddonVersion(v string) *DescribeAddonRequest {
	s.AddonVersion = &v
	return s
}

func (s *DescribeAddonRequest) Validate() error {
	return dara.Validate(s)
}
