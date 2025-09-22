// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResourceCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTemplateResourceCountResponseBody
	GetRequestId() *string
	SetResourceCount(v []*DescribeTemplateResourceCountResponseBodyResourceCount) *DescribeTemplateResourceCountResponseBody
	GetResourceCount() []*DescribeTemplateResourceCountResponseBodyResourceCount
}

type DescribeTemplateResourceCountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B8064433-9781-5E86-806E-C1DD****1D95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of protected objects or protected object groups for which the protection template takes effect.
	ResourceCount []*DescribeTemplateResourceCountResponseBodyResourceCount `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty" type:"Repeated"`
}

func (s DescribeTemplateResourceCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResourceCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplateResourceCountResponseBody) GetResourceCount() []*DescribeTemplateResourceCountResponseBodyResourceCount {
	return s.ResourceCount
}

func (s *DescribeTemplateResourceCountResponseBody) SetRequestId(v string) *DescribeTemplateResourceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBody) SetResourceCount(v []*DescribeTemplateResourceCountResponseBodyResourceCount) *DescribeTemplateResourceCountResponseBody {
	s.ResourceCount = v
	return s
}

func (s *DescribeTemplateResourceCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTemplateResourceCountResponseBodyResourceCount struct {
	// example:
	//
	// 10
	AssetCount *int32 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The number of protected object groups.
	//
	// example:
	//
	// 30
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The number of protected objects.
	//
	// example:
	//
	// 30
	SingleCount *int32 `json:"SingleCount,omitempty" xml:"SingleCount,omitempty"`
	// The ID of the protection template.
	//
	// example:
	//
	// 12345
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeTemplateResourceCountResponseBodyResourceCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResourceCountResponseBodyResourceCount) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) GetAssetCount() *int32 {
	return s.AssetCount
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) GetSingleCount() *int32 {
	return s.SingleCount
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) SetAssetCount(v int32) *DescribeTemplateResourceCountResponseBodyResourceCount {
	s.AssetCount = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) SetGroupCount(v int32) *DescribeTemplateResourceCountResponseBodyResourceCount {
	s.GroupCount = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) SetSingleCount(v int32) *DescribeTemplateResourceCountResponseBodyResourceCount {
	s.SingleCount = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) SetTemplateId(v int64) *DescribeTemplateResourceCountResponseBodyResourceCount {
	s.TemplateId = &v
	return s
}

func (s *DescribeTemplateResourceCountResponseBodyResourceCount) Validate() error {
	return dara.Validate(s)
}
