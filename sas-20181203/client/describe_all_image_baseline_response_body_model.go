// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllImageBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageBaselines(v *DescribeAllImageBaselineResponseBodyImageBaselines) *DescribeAllImageBaselineResponseBody
	GetImageBaselines() *DescribeAllImageBaselineResponseBodyImageBaselines
	SetRequestId(v string) *DescribeAllImageBaselineResponseBody
	GetRequestId() *string
}

type DescribeAllImageBaselineResponseBody struct {
	// The baselines that are used in image baseline checks.
	ImageBaselines *DescribeAllImageBaselineResponseBodyImageBaselines `json:"ImageBaselines,omitempty" xml:"ImageBaselines,omitempty" type:"Struct"`
	// The ID of the request, which is used to query logs and troubleshoot issues.
	//
	// example:
	//
	// 1A975D03-5F49-5354-B2CB-3918D5DA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAllImageBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllImageBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllImageBaselineResponseBody) GetImageBaselines() *DescribeAllImageBaselineResponseBodyImageBaselines {
	return s.ImageBaselines
}

func (s *DescribeAllImageBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllImageBaselineResponseBody) SetImageBaselines(v *DescribeAllImageBaselineResponseBodyImageBaselines) *DescribeAllImageBaselineResponseBody {
	s.ImageBaselines = v
	return s
}

func (s *DescribeAllImageBaselineResponseBody) SetRequestId(v string) *DescribeAllImageBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBody) Validate() error {
	if s.ImageBaselines != nil {
		if err := s.ImageBaselines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAllImageBaselineResponseBodyImageBaselines struct {
	// An array that consists of baseline types.
	BaselineClassList []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList `json:"BaselineClassList,omitempty" xml:"BaselineClassList,omitempty" type:"Repeated"`
}

func (s DescribeAllImageBaselineResponseBodyImageBaselines) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllImageBaselineResponseBodyImageBaselines) GoString() string {
	return s.String()
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselines) GetBaselineClassList() []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList {
	return s.BaselineClassList
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselines) SetBaselineClassList(v []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) *DescribeAllImageBaselineResponseBodyImageBaselines {
	s.BaselineClassList = v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselines) Validate() error {
	if s.BaselineClassList != nil {
		for _, item := range s.BaselineClassList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList struct {
	// The alias of the baseline type.
	//
	// example:
	//
	// Identity authentication
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The information about the baseline.
	BaselineNameList []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList `json:"BaselineNameList,omitempty" xml:"BaselineNameList,omitempty" type:"Repeated"`
	// The key of the baseline type.
	//
	// example:
	//
	// identification
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
}

func (s DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) GoString() string {
	return s.String()
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) GetAlias() *string {
	return s.Alias
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) GetBaselineNameList() []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList {
	return s.BaselineNameList
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) GetClassKey() *string {
	return s.ClassKey
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) SetAlias(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList {
	s.Alias = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) SetBaselineNameList(v []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList {
	s.BaselineNameList = v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) SetClassKey(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList {
	s.ClassKey = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassList) Validate() error {
	if s.BaselineNameList != nil {
		for _, item := range s.BaselineNameList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList struct {
	// The alias of the baseline.
	//
	// example:
	//
	// Identity authentication
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The information about the baseline check item.
	BaselineItemList []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList `json:"BaselineItemList,omitempty" xml:"BaselineItemList,omitempty" type:"Repeated"`
	// The key of the type for the baseline.
	//
	// example:
	//
	// identification
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The key of the name for the baseline.
	//
	// example:
	//
	// identification
	NameKey *string `json:"NameKey,omitempty" xml:"NameKey,omitempty"`
}

func (s DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) GoString() string {
	return s.String()
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) GetAlias() *string {
	return s.Alias
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) GetBaselineItemList() []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList {
	return s.BaselineItemList
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) GetClassKey() *string {
	return s.ClassKey
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) GetNameKey() *string {
	return s.NameKey
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) SetAlias(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList {
	s.Alias = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) SetBaselineItemList(v []*DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList {
	s.BaselineItemList = v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) SetClassKey(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList {
	s.ClassKey = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) SetNameKey(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList {
	s.NameKey = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameList) Validate() error {
	if s.BaselineItemList != nil {
		for _, item := range s.BaselineItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList struct {
	// The alias of the baseline check item.
	//
	// example:
	//
	// Ensure that no accounts use the same Hash password.
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The key of the type for the baseline.
	//
	// example:
	//
	// identification
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The key of the name for the baseline check item.
	//
	// example:
	//
	// duplicate_pwd_hash
	ItemKey *string `json:"ItemKey,omitempty" xml:"ItemKey,omitempty"`
	// The key of the name for the baseline.
	//
	// example:
	//
	// identification
	NameKey *string `json:"NameKey,omitempty" xml:"NameKey,omitempty"`
}

func (s DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) GoString() string {
	return s.String()
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) GetAlias() *string {
	return s.Alias
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) GetClassKey() *string {
	return s.ClassKey
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) GetItemKey() *string {
	return s.ItemKey
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) GetNameKey() *string {
	return s.NameKey
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) SetAlias(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList {
	s.Alias = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) SetClassKey(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList {
	s.ClassKey = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) SetItemKey(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList {
	s.ItemKey = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) SetNameKey(v string) *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList {
	s.NameKey = &v
	return s
}

func (s *DescribeAllImageBaselineResponseBodyImageBaselinesBaselineClassListBaselineNameListBaselineItemList) Validate() error {
	return dara.Validate(s)
}
