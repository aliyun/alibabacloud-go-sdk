// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceTypeFamiliesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypeFamilies(v *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) *DescribeRCInstanceTypeFamiliesResponseBody
	GetInstanceTypeFamilies() *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies
	SetRequestId(v string) *DescribeRCInstanceTypeFamiliesResponseBody
	GetRequestId() *string
}

type DescribeRCInstanceTypeFamiliesResponseBody struct {
	// The instance families.
	InstanceTypeFamilies *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F2911788-25E8-42E5-A3A3-1B38D263F01E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCInstanceTypeFamiliesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypeFamiliesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypeFamiliesResponseBody) GetInstanceTypeFamilies() *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies {
	return s.InstanceTypeFamilies
}

func (s *DescribeRCInstanceTypeFamiliesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInstanceTypeFamiliesResponseBody) SetInstanceTypeFamilies(v *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) *DescribeRCInstanceTypeFamiliesResponseBody {
	s.InstanceTypeFamilies = v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponseBody) SetRequestId(v string) *DescribeRCInstanceTypeFamiliesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponseBody) Validate() error {
	if s.InstanceTypeFamilies != nil {
		if err := s.InstanceTypeFamilies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies struct {
	// The instance family.
	InstanceTypeFamily []*DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) GetInstanceTypeFamily() []*DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily {
	return s.InstanceTypeFamily
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) SetInstanceTypeFamily(v []*DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies {
	s.InstanceTypeFamily = v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) Validate() error {
	if s.InstanceTypeFamily != nil {
		for _, item := range s.InstanceTypeFamily {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily struct {
	// The description of the instance family.
	InstanceTypeFamilyDesc *string `json:"InstanceTypeFamilyDesc,omitempty" xml:"InstanceTypeFamilyDesc,omitempty"`
	// The ID of the instance family.
	//
	// example:
	//
	// x.6cm
	InstanceTypeFamilyId *string `json:"InstanceTypeFamilyId,omitempty" xml:"InstanceTypeFamilyId,omitempty"`
}

func (s DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) GetInstanceTypeFamilyDesc() *string {
	return s.InstanceTypeFamilyDesc
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) GetInstanceTypeFamilyId() *string {
	return s.InstanceTypeFamilyId
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) SetInstanceTypeFamilyDesc(v string) *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily {
	s.InstanceTypeFamilyDesc = &v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) SetInstanceTypeFamilyId(v string) *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily {
	s.InstanceTypeFamilyId = &v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) Validate() error {
	return dara.Validate(s)
}
