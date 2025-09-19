// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetsScaProcessNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeAssetsScaProcessNumResponseBodyData) *DescribeAssetsScaProcessNumResponseBody
	GetData() []*DescribeAssetsScaProcessNumResponseBodyData
	SetRequestId(v string) *DescribeAssetsScaProcessNumResponseBody
	GetRequestId() *string
}

type DescribeAssetsScaProcessNumResponseBody struct {
	// The statistical results.
	Data []*DescribeAssetsScaProcessNumResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12BREF20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssetsScaProcessNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsScaProcessNumResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetsScaProcessNumResponseBody) GetData() []*DescribeAssetsScaProcessNumResponseBodyData {
	return s.Data
}

func (s *DescribeAssetsScaProcessNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetsScaProcessNumResponseBody) SetData(v []*DescribeAssetsScaProcessNumResponseBodyData) *DescribeAssetsScaProcessNumResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAssetsScaProcessNumResponseBody) SetRequestId(v string) *DescribeAssetsScaProcessNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetsScaProcessNumResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetsScaProcessNumResponseBodyData struct {
	// The number of Java processes.
	//
	// >  If no processes exist on the asset, no statistical result is returned.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The UUID of the asset.
	//
	// >  If no processes exist on the asset, no statistical result is returned.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAssetsScaProcessNumResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsScaProcessNumResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAssetsScaProcessNumResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAssetsScaProcessNumResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAssetsScaProcessNumResponseBodyData) SetCount(v int32) *DescribeAssetsScaProcessNumResponseBodyData {
	s.Count = &v
	return s
}

func (s *DescribeAssetsScaProcessNumResponseBodyData) SetUuid(v string) *DescribeAssetsScaProcessNumResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *DescribeAssetsScaProcessNumResponseBodyData) Validate() error {
	return dara.Validate(s)
}
