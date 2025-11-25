// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpenCount(v int64) *DescribeSensitiveSwitchResponseBody
	GetOpenCount() *int64
	SetRequestId(v string) *DescribeSensitiveSwitchResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSensitiveSwitchResponseBody
	GetTotalCount() *int64
	SetUserSensitiveDataSwitchList(v []*DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) *DescribeSensitiveSwitchResponseBody
	GetUserSensitiveDataSwitchList() []*DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList
}

type DescribeSensitiveSwitchResponseBody struct {
	// example:
	//
	// 80
	OpenCount *int64 `json:"OpenCount,omitempty" xml:"OpenCount,omitempty"`
	// example:
	//
	// A7F3ED45-5556-5AF3-ADE3-EE48FFF0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount                  *int64                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserSensitiveDataSwitchList []*DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList `json:"UserSensitiveDataSwitchList,omitempty" xml:"UserSensitiveDataSwitchList,omitempty" type:"Repeated"`
}

func (s DescribeSensitiveSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveSwitchResponseBody) GetOpenCount() *int64 {
	return s.OpenCount
}

func (s *DescribeSensitiveSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSensitiveSwitchResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSensitiveSwitchResponseBody) GetUserSensitiveDataSwitchList() []*DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList {
	return s.UserSensitiveDataSwitchList
}

func (s *DescribeSensitiveSwitchResponseBody) SetOpenCount(v int64) *DescribeSensitiveSwitchResponseBody {
	s.OpenCount = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBody) SetRequestId(v string) *DescribeSensitiveSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBody) SetTotalCount(v int64) *DescribeSensitiveSwitchResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBody) SetUserSensitiveDataSwitchList(v []*DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) *DescribeSensitiveSwitchResponseBody {
	s.UserSensitiveDataSwitchList = v
	return s
}

func (s *DescribeSensitiveSwitchResponseBody) Validate() error {
	if s.UserSensitiveDataSwitchList != nil {
		for _, item := range s.UserSensitiveDataSwitchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// id_card
	SensitiveCategory *string `json:"SensitiveCategory,omitempty" xml:"SensitiveCategory,omitempty"`
	// example:
	//
	// S3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// example:
	//
	// 1
	SwitchStatus *int32 `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) GetDescription() *string {
	return s.Description
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) GetSensitiveCategory() *string {
	return s.SensitiveCategory
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) GetSwitchStatus() *int32 {
	return s.SwitchStatus
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) SetCategoryName(v string) *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList {
	s.CategoryName = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) SetDescription(v string) *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList {
	s.Description = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) SetSensitiveCategory(v string) *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList {
	s.SensitiveCategory = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) SetSensitiveLevel(v string) *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) SetSwitchStatus(v int32) *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList {
	s.SwitchStatus = &v
	return s
}

func (s *DescribeSensitiveSwitchResponseBodyUserSensitiveDataSwitchList) Validate() error {
	return dara.Validate(s)
}
