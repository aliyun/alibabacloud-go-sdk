// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskItemTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeRiskItemTypeResponseBodyList) *DescribeRiskItemTypeResponseBody
	GetList() []*DescribeRiskItemTypeResponseBodyList
	SetRequestId(v string) *DescribeRiskItemTypeResponseBody
	GetRequestId() *string
}

type DescribeRiskItemTypeResponseBody struct {
	// An array that consists of information about the type of the check item.
	List []*DescribeRiskItemTypeResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3B3F3A90-46A5-4023-A2D8-D68B14262F96
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRiskItemTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskItemTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskItemTypeResponseBody) GetList() []*DescribeRiskItemTypeResponseBodyList {
	return s.List
}

func (s *DescribeRiskItemTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskItemTypeResponseBody) SetList(v []*DescribeRiskItemTypeResponseBodyList) *DescribeRiskItemTypeResponseBody {
	s.List = v
	return s
}

func (s *DescribeRiskItemTypeResponseBody) SetRequestId(v string) *DescribeRiskItemTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskItemTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskItemTypeResponseBodyList struct {
	// The ID of the check item.
	//
	// example:
	//
	// 37625
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the check type.
	//
	// example:
	//
	// Identity authentication and permissions
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeRiskItemTypeResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskItemTypeResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeRiskItemTypeResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *DescribeRiskItemTypeResponseBodyList) GetTitle() *string {
	return s.Title
}

func (s *DescribeRiskItemTypeResponseBodyList) SetId(v int64) *DescribeRiskItemTypeResponseBodyList {
	s.Id = &v
	return s
}

func (s *DescribeRiskItemTypeResponseBodyList) SetTitle(v string) *DescribeRiskItemTypeResponseBodyList {
	s.Title = &v
	return s
}

func (s *DescribeRiskItemTypeResponseBodyList) Validate() error {
	return dara.Validate(s)
}
