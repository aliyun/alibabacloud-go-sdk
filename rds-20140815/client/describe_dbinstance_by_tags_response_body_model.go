// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceByTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDBInstanceByTagsResponseBodyItems) *DescribeDBInstanceByTagsResponseBody
	GetItems() *DescribeDBInstanceByTagsResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstanceByTagsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDBInstanceByTagsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDBInstanceByTagsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstanceByTagsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstanceByTagsResponseBody struct {
	Items *DescribeDBInstanceByTagsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstanceByTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceByTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceByTagsResponseBody) GetItems() *DescribeDBInstanceByTagsResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceByTagsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceByTagsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDBInstanceByTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceByTagsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstanceByTagsResponseBody) SetItems(v *DescribeDBInstanceByTagsResponseBodyItems) *DescribeDBInstanceByTagsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBody) SetPageNumber(v int32) *DescribeDBInstanceByTagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBody) SetPageRecordCount(v int32) *DescribeDBInstanceByTagsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBody) SetRequestId(v string) *DescribeDBInstanceByTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstanceByTagsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceByTagsResponseBodyItems struct {
	DBInstanceTag []*DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag `json:"DBInstanceTag,omitempty" xml:"DBInstanceTag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceByTagsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceByTagsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceByTagsResponseBodyItems) GetDBInstanceTag() []*DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag {
	return s.DBInstanceTag
}

func (s *DescribeDBInstanceByTagsResponseBodyItems) SetDBInstanceTag(v []*DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) *DescribeDBInstanceByTagsResponseBodyItems {
	s.DBInstanceTag = v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBodyItems) Validate() error {
	if s.DBInstanceTag != nil {
		for _, item := range s.DBInstanceTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag struct {
	DBInstanceId *string                                                     `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Tags         *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) GetTags() *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags {
	return s.Tags
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) SetDBInstanceId(v string) *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) SetTags(v *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags) *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag {
	s.Tags = v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTag) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags struct {
	Tag []*DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags) GetTag() []*DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag {
	return s.Tag
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags) SetTag(v []*DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags {
	s.Tag = v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) SetTagKey(v string) *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) SetTagValue(v string) *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponseBodyItemsDBInstanceTagTagsTag) Validate() error {
	return dara.Validate(s)
}
