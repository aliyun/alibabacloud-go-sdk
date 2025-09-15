// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonLogFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogFieldList(v []*DescribeCommonLogFieldsResponseBodyLogFieldList) *DescribeCommonLogFieldsResponseBody
	GetLogFieldList() []*DescribeCommonLogFieldsResponseBodyLogFieldList
	SetRequestId(v string) *DescribeCommonLogFieldsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeCommonLogFieldsResponseBody
	GetTotalCount() *int64
}

type DescribeCommonLogFieldsResponseBody struct {
	LogFieldList []*DescribeCommonLogFieldsResponseBodyLogFieldList `json:"LogFieldList,omitempty" xml:"LogFieldList,omitempty" type:"Repeated"`
	// example:
	//
	// AB0775EB-2594-598A-97E1-51B1*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommonLogFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonLogFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommonLogFieldsResponseBody) GetLogFieldList() []*DescribeCommonLogFieldsResponseBodyLogFieldList {
	return s.LogFieldList
}

func (s *DescribeCommonLogFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommonLogFieldsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCommonLogFieldsResponseBody) SetLogFieldList(v []*DescribeCommonLogFieldsResponseBodyLogFieldList) *DescribeCommonLogFieldsResponseBody {
	s.LogFieldList = v
	return s
}

func (s *DescribeCommonLogFieldsResponseBody) SetRequestId(v string) *DescribeCommonLogFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommonLogFieldsResponseBody) SetTotalCount(v int64) *DescribeCommonLogFieldsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCommonLogFieldsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCommonLogFieldsResponseBodyLogFieldList struct {
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// true
	IsRequired *bool `json:"IsRequired,omitempty" xml:"IsRequired,omitempty"`
	// example:
	//
	// final_action
	LogKey *string `json:"LogKey,omitempty" xml:"LogKey,omitempty"`
	// example:
	//
	// 1
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCommonLogFieldsResponseBodyLogFieldList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonLogFieldsResponseBodyLogFieldList) GoString() string {
	return s.String()
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) GetLogKey() *string {
	return s.LogKey
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) GetStatus() *bool {
	return s.Status
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) SetIsDefault(v bool) *DescribeCommonLogFieldsResponseBodyLogFieldList {
	s.IsDefault = &v
	return s
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) SetIsRequired(v bool) *DescribeCommonLogFieldsResponseBodyLogFieldList {
	s.IsRequired = &v
	return s
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) SetLogKey(v string) *DescribeCommonLogFieldsResponseBodyLogFieldList {
	s.LogKey = &v
	return s
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) SetStatus(v bool) *DescribeCommonLogFieldsResponseBodyLogFieldList {
	s.Status = &v
	return s
}

func (s *DescribeCommonLogFieldsResponseBodyLogFieldList) Validate() error {
	return dara.Validate(s)
}
