// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeTagsResponseBodyItems) *DescribeTagsResponseBody
	GetItems() *DescribeTagsResponseBodyItems
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
}

type DescribeTagsResponseBody struct {
	Items *DescribeTagsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetItems() *DescribeTagsResponseBodyItems {
	return s.Items
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) SetItems(v *DescribeTagsResponseBodyItems) *DescribeTagsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsResponseBodyItems struct {
	TagInfos []*DescribeTagsResponseBodyItemsTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyItems) GetTagInfos() []*DescribeTagsResponseBodyItemsTagInfos {
	return s.TagInfos
}

func (s *DescribeTagsResponseBodyItems) SetTagInfos(v []*DescribeTagsResponseBodyItemsTagInfos) *DescribeTagsResponseBodyItems {
	s.TagInfos = v
	return s
}

func (s *DescribeTagsResponseBodyItems) Validate() error {
	if s.TagInfos != nil {
		for _, item := range s.TagInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagsResponseBodyItemsTagInfos struct {
	DBInstanceIds *DescribeTagsResponseBodyItemsTagInfosDBInstanceIds `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Struct"`
	TagKey        *string                                             `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue      *string                                             `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyItemsTagInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyItemsTagInfos) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyItemsTagInfos) GetDBInstanceIds() *DescribeTagsResponseBodyItemsTagInfosDBInstanceIds {
	return s.DBInstanceIds
}

func (s *DescribeTagsResponseBodyItemsTagInfos) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyItemsTagInfos) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeTagsResponseBodyItemsTagInfos) SetDBInstanceIds(v *DescribeTagsResponseBodyItemsTagInfosDBInstanceIds) *DescribeTagsResponseBodyItemsTagInfos {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeTagsResponseBodyItemsTagInfos) SetTagKey(v string) *DescribeTagsResponseBodyItemsTagInfos {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyItemsTagInfos) SetTagValue(v string) *DescribeTagsResponseBodyItemsTagInfos {
	s.TagValue = &v
	return s
}

func (s *DescribeTagsResponseBodyItemsTagInfos) Validate() error {
	if s.DBInstanceIds != nil {
		if err := s.DBInstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsResponseBodyItemsTagInfosDBInstanceIds struct {
	DBInstanceIds []*string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyItemsTagInfosDBInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyItemsTagInfosDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyItemsTagInfosDBInstanceIds) GetDBInstanceIds() []*string {
	return s.DBInstanceIds
}

func (s *DescribeTagsResponseBodyItemsTagInfosDBInstanceIds) SetDBInstanceIds(v []*string) *DescribeTagsResponseBodyItemsTagInfosDBInstanceIds {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeTagsResponseBodyItemsTagInfosDBInstanceIds) Validate() error {
	return dara.Validate(s)
}
