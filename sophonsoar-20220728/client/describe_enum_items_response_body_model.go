// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnumItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeEnumItemsResponseBodyData) *DescribeEnumItemsResponseBody
	GetData() []*DescribeEnumItemsResponseBodyData
	SetRequestId(v string) *DescribeEnumItemsResponseBody
	GetRequestId() *string
}

type DescribeEnumItemsResponseBody struct {
	// The information about the enumeration item.
	Data []*DescribeEnumItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E7698CFB-4E1C-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnumItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnumItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponseBody) GetData() []*DescribeEnumItemsResponseBodyData {
	return s.Data
}

func (s *DescribeEnumItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnumItemsResponseBody) SetData(v []*DescribeEnumItemsResponseBodyData) *DescribeEnumItemsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnumItemsResponseBody) SetRequestId(v string) *DescribeEnumItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnumItemsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnumItemsResponseBodyData struct {
	// The key of the enumeration item.
	//
	// example:
	//
	// system_xxxxx_process_book
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the enumeration item.
	//
	// example:
	//
	// system_xxxxx_process_book
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEnumItemsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnumItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *DescribeEnumItemsResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *DescribeEnumItemsResponseBodyData) SetKey(v string) *DescribeEnumItemsResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeEnumItemsResponseBodyData) SetValue(v string) *DescribeEnumItemsResponseBodyData {
	s.Value = &v
	return s
}

func (s *DescribeEnumItemsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
