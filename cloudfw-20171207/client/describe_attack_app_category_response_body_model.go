// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAppCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppCategories(v []*DescribeAttackAppCategoryResponseBodyAppCategories) *DescribeAttackAppCategoryResponseBody
	GetAppCategories() []*DescribeAttackAppCategoryResponseBodyAppCategories
	SetRequestId(v string) *DescribeAttackAppCategoryResponseBody
	GetRequestId() *string
}

type DescribeAttackAppCategoryResponseBody struct {
	AppCategories []*DescribeAttackAppCategoryResponseBodyAppCategories `json:"AppCategories,omitempty" xml:"AppCategories,omitempty" type:"Repeated"`
	// example:
	//
	// B14757D0-4640-4B44-AC67-7F558FE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAttackAppCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAppCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAttackAppCategoryResponseBody) GetAppCategories() []*DescribeAttackAppCategoryResponseBodyAppCategories {
	return s.AppCategories
}

func (s *DescribeAttackAppCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAttackAppCategoryResponseBody) SetAppCategories(v []*DescribeAttackAppCategoryResponseBodyAppCategories) *DescribeAttackAppCategoryResponseBody {
	s.AppCategories = v
	return s
}

func (s *DescribeAttackAppCategoryResponseBody) SetRequestId(v string) *DescribeAttackAppCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAttackAppCategoryResponseBody) Validate() error {
	if s.AppCategories != nil {
		for _, item := range s.AppCategories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAttackAppCategoryResponseBodyAppCategories struct {
	AttackApps []*string `json:"AttackApps,omitempty" xml:"AttackApps,omitempty" type:"Repeated"`
	// example:
	//
	// test
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s DescribeAttackAppCategoryResponseBodyAppCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAppCategoryResponseBodyAppCategories) GoString() string {
	return s.String()
}

func (s *DescribeAttackAppCategoryResponseBodyAppCategories) GetAttackApps() []*string {
	return s.AttackApps
}

func (s *DescribeAttackAppCategoryResponseBodyAppCategories) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeAttackAppCategoryResponseBodyAppCategories) SetAttackApps(v []*string) *DescribeAttackAppCategoryResponseBodyAppCategories {
	s.AttackApps = v
	return s
}

func (s *DescribeAttackAppCategoryResponseBodyAppCategories) SetCategoryName(v string) *DescribeAttackAppCategoryResponseBodyAppCategories {
	s.CategoryName = &v
	return s
}

func (s *DescribeAttackAppCategoryResponseBodyAppCategories) Validate() error {
	return dara.Validate(s)
}
