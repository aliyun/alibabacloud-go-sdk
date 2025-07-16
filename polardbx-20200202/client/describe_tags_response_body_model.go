// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetTagInfos(v []*DescribeTagsResponseBodyTagInfos) *DescribeTagsResponseBody
	GetTagInfos() []*DescribeTagsResponseBodyTagInfos
}

type DescribeTagsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// requestid
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagInfos  []*DescribeTagsResponseBodyTagInfos `json:"TagInfos,omitempty" xml:"TagInfos,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetTagInfos() []*DescribeTagsResponseBodyTagInfos {
	return s.TagInfos
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTagInfos(v []*DescribeTagsResponseBodyTagInfos) *DescribeTagsResponseBody {
	s.TagInfos = v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsResponseBodyTagInfos struct {
	DBInstanceIds []*string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTagInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagInfos) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagInfos) GetDBInstanceIds() []*string {
	return s.DBInstanceIds
}

func (s *DescribeTagsResponseBodyTagInfos) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyTagInfos) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeTagsResponseBodyTagInfos) SetDBInstanceIds(v []*string) *DescribeTagsResponseBodyTagInfos {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeTagsResponseBodyTagInfos) SetTagKey(v string) *DescribeTagsResponseBodyTagInfos {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTagInfos) SetTagValue(v string) *DescribeTagsResponseBodyTagInfos {
	s.TagValue = &v
	return s
}

func (s *DescribeTagsResponseBodyTagInfos) Validate() error {
	return dara.Validate(s)
}
