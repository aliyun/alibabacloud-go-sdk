// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductDataRedundancyTypeStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeProductDataRedundancyTypeStatResponseBodyData) *DescribeProductDataRedundancyTypeStatResponseBody
	GetData() *DescribeProductDataRedundancyTypeStatResponseBodyData
	SetRequestId(v string) *DescribeProductDataRedundancyTypeStatResponseBody
	GetRequestId() *string
}

type DescribeProductDataRedundancyTypeStatResponseBody struct {
	// The response data.
	Data *DescribeProductDataRedundancyTypeStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique request ID.
	//
	// example:
	//
	// 5748C531-80B1-5C31-8421-63A1830B9E48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProductDataRedundancyTypeStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductDataRedundancyTypeStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductDataRedundancyTypeStatResponseBody) GetData() *DescribeProductDataRedundancyTypeStatResponseBodyData {
	return s.Data
}

func (s *DescribeProductDataRedundancyTypeStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductDataRedundancyTypeStatResponseBody) SetData(v *DescribeProductDataRedundancyTypeStatResponseBodyData) *DescribeProductDataRedundancyTypeStatResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponseBody) SetRequestId(v string) *DescribeProductDataRedundancyTypeStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProductDataRedundancyTypeStatResponseBodyData struct {
	// The list of records returned by the request.
	Content []*DescribeProductDataRedundancyTypeStatResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
}

func (s DescribeProductDataRedundancyTypeStatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductDataRedundancyTypeStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyData) GetContent() []*DescribeProductDataRedundancyTypeStatResponseBodyDataContent {
	return s.Content
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyData) SetContent(v []*DescribeProductDataRedundancyTypeStatResponseBodyDataContent) *DescribeProductDataRedundancyTypeStatResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProductDataRedundancyTypeStatResponseBodyDataContent struct {
	// The data redundancy type.
	//
	// example:
	//
	// LRS
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	// The resource count.
	//
	// example:
	//
	// 1
	ResourceCount *int64 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The storage class.
	//
	// example:
	//
	// STANDARD
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s DescribeProductDataRedundancyTypeStatResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductDataRedundancyTypeStatResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyDataContent) GetDataRedundancyType() *string {
	return s.DataRedundancyType
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyDataContent) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyDataContent) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyDataContent) SetDataRedundancyType(v string) *DescribeProductDataRedundancyTypeStatResponseBodyDataContent {
	s.DataRedundancyType = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyDataContent) SetResourceCount(v int64) *DescribeProductDataRedundancyTypeStatResponseBodyDataContent {
	s.ResourceCount = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyDataContent) SetStorageClass(v string) *DescribeProductDataRedundancyTypeStatResponseBodyDataContent {
	s.StorageClass = &v
	return s
}

func (s *DescribeProductDataRedundancyTypeStatResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
