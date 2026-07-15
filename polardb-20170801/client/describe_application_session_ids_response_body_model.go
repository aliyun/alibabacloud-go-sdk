// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSessionIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationSessionIdsResponseBody
	GetApplicationId() *string
	SetItems(v *DescribeApplicationSessionIdsResponseBodyItems) *DescribeApplicationSessionIdsResponseBody
	GetItems() *DescribeApplicationSessionIdsResponseBodyItems
	SetPageNumber(v int32) *DescribeApplicationSessionIdsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeApplicationSessionIdsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeApplicationSessionIdsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeApplicationSessionIdsResponseBody
	GetTotalRecordCount() *string
}

type DescribeApplicationSessionIdsResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-********************
	ApplicationId *string                                         `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Items         *DescribeApplicationSessionIdsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7F2007D3-7E74-4ECB-89A8-BF130D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeApplicationSessionIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionIdsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionIdsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationSessionIdsResponseBody) GetItems() *DescribeApplicationSessionIdsResponseBodyItems {
	return s.Items
}

func (s *DescribeApplicationSessionIdsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationSessionIdsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeApplicationSessionIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationSessionIdsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeApplicationSessionIdsResponseBody) SetApplicationId(v string) *DescribeApplicationSessionIdsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBody) SetItems(v *DescribeApplicationSessionIdsResponseBodyItems) *DescribeApplicationSessionIdsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBody) SetPageNumber(v int32) *DescribeApplicationSessionIdsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBody) SetPageRecordCount(v int32) *DescribeApplicationSessionIdsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBody) SetRequestId(v string) *DescribeApplicationSessionIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBody) SetTotalRecordCount(v string) *DescribeApplicationSessionIdsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationSessionIdsResponseBodyItems struct {
	Items []*DescribeApplicationSessionIdsResponseBodyItemsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s DescribeApplicationSessionIdsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionIdsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionIdsResponseBodyItems) GetItems() []*DescribeApplicationSessionIdsResponseBodyItemsItems {
	return s.Items
}

func (s *DescribeApplicationSessionIdsResponseBodyItems) SetItems(v []*DescribeApplicationSessionIdsResponseBodyItemsItems) *DescribeApplicationSessionIdsResponseBodyItems {
	s.Items = v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBodyItems) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationSessionIdsResponseBodyItemsItems struct {
	Agent     *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Time      *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeApplicationSessionIdsResponseBodyItemsItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionIdsResponseBodyItemsItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionIdsResponseBodyItemsItems) GetAgent() *string {
	return s.Agent
}

func (s *DescribeApplicationSessionIdsResponseBodyItemsItems) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeApplicationSessionIdsResponseBodyItemsItems) GetTime() *string {
	return s.Time
}

func (s *DescribeApplicationSessionIdsResponseBodyItemsItems) SetAgent(v string) *DescribeApplicationSessionIdsResponseBodyItemsItems {
	s.Agent = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBodyItemsItems) SetSessionId(v string) *DescribeApplicationSessionIdsResponseBodyItemsItems {
	s.SessionId = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBodyItemsItems) SetTime(v string) *DescribeApplicationSessionIdsResponseBodyItemsItems {
	s.Time = &v
	return s
}

func (s *DescribeApplicationSessionIdsResponseBodyItemsItems) Validate() error {
	return dara.Validate(s)
}
