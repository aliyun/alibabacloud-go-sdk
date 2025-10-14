// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpgradeItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListUpgradeItemsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUpgradeItemsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUpgradeItemsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListUpgradeItemsResponseBody
	GetTotalCount() *int32
	SetUpgradeItems(v []*ListUpgradeItemsResponseBodyUpgradeItems) *ListUpgradeItemsResponseBody
	GetUpgradeItems() []*ListUpgradeItemsResponseBodyUpgradeItems
}

type ListUpgradeItemsResponseBody struct {
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UpgradeItems []*ListUpgradeItemsResponseBodyUpgradeItems `json:"UpgradeItems,omitempty" xml:"UpgradeItems,omitempty" type:"Repeated"`
}

func (s ListUpgradeItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUpgradeItemsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUpgradeItemsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUpgradeItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUpgradeItemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUpgradeItemsResponseBody) GetUpgradeItems() []*ListUpgradeItemsResponseBodyUpgradeItems {
	return s.UpgradeItems
}

func (s *ListUpgradeItemsResponseBody) SetMaxResults(v int32) *ListUpgradeItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUpgradeItemsResponseBody) SetNextToken(v string) *ListUpgradeItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUpgradeItemsResponseBody) SetRequestId(v string) *ListUpgradeItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUpgradeItemsResponseBody) SetTotalCount(v int32) *ListUpgradeItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUpgradeItemsResponseBody) SetUpgradeItems(v []*ListUpgradeItemsResponseBodyUpgradeItems) *ListUpgradeItemsResponseBody {
	s.UpgradeItems = v
	return s
}

func (s *ListUpgradeItemsResponseBody) Validate() error {
	if s.UpgradeItems != nil {
		for _, item := range s.UpgradeItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUpgradeItemsResponseBodyUpgradeItems struct {
	// example:
	//
	// data_storage_2_upgrade
	UpgradeItemId *string `json:"UpgradeItemId,omitempty" xml:"UpgradeItemId,omitempty"`
}

func (s ListUpgradeItemsResponseBodyUpgradeItems) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeItemsResponseBodyUpgradeItems) GoString() string {
	return s.String()
}

func (s *ListUpgradeItemsResponseBodyUpgradeItems) GetUpgradeItemId() *string {
	return s.UpgradeItemId
}

func (s *ListUpgradeItemsResponseBodyUpgradeItems) SetUpgradeItemId(v string) *ListUpgradeItemsResponseBodyUpgradeItems {
	s.UpgradeItemId = &v
	return s
}

func (s *ListUpgradeItemsResponseBodyUpgradeItems) Validate() error {
	return dara.Validate(s)
}
