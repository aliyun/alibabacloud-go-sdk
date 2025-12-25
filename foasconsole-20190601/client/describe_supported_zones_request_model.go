// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportedZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *DescribeSupportedZonesRequest
	GetArchitectureType() *string
	SetRegion(v string) *DescribeSupportedZonesRequest
	GetRegion() *string
}

type DescribeSupportedZonesRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeSupportedZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportedZonesRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeSupportedZonesRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeSupportedZonesRequest) SetArchitectureType(v string) *DescribeSupportedZonesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeSupportedZonesRequest) SetRegion(v string) *DescribeSupportedZonesRequest {
	s.Region = &v
	return s
}

func (s *DescribeSupportedZonesRequest) Validate() error {
	return dara.Validate(s)
}
