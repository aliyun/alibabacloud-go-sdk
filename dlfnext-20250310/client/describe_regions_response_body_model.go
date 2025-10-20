// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() []*DescribeRegionsResponseBodyRegions
}

type DescribeRegionsResponseBody struct {
	Regions []*DescribeRegionsResponseBodyRegions `json:"regions,omitempty" xml:"regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegions() []*DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegions struct {
	// The region description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The region name
	//
	// example:
	//
	// cn-hangzhou
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region show name
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
	// The region type
	//
	// example:
	//
	// region
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetDescription() *string {
	return s.Description
}

func (s *DescribeRegionsResponseBodyRegions) GetName() *string {
	return s.Name
}

func (s *DescribeRegionsResponseBodyRegions) GetShowName() *string {
	return s.ShowName
}

func (s *DescribeRegionsResponseBodyRegions) GetType() *string {
	return s.Type
}

func (s *DescribeRegionsResponseBodyRegions) SetDescription(v string) *DescribeRegionsResponseBodyRegions {
	s.Description = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetName(v string) *DescribeRegionsResponseBodyRegions {
	s.Name = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetShowName(v string) *DescribeRegionsResponseBodyRegions {
	s.ShowName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetType(v string) *DescribeRegionsResponseBodyRegions {
	s.Type = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
