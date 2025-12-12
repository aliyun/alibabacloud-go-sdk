// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventMetaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeEventMetaInfoResponseBodyItems) *DescribeEventMetaInfoResponseBody
	GetItems() []*DescribeEventMetaInfoResponseBodyItems
	SetRequestId(v string) *DescribeEventMetaInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEventMetaInfoResponseBody
	GetTotalCount() *int32
}

type DescribeEventMetaInfoResponseBody struct {
	Items []*DescribeEventMetaInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 94F92113-FF63-5E57-8401-6FE123AD11DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventMetaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventMetaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventMetaInfoResponseBody) GetItems() []*DescribeEventMetaInfoResponseBodyItems {
	return s.Items
}

func (s *DescribeEventMetaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventMetaInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEventMetaInfoResponseBody) SetItems(v []*DescribeEventMetaInfoResponseBodyItems) *DescribeEventMetaInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeEventMetaInfoResponseBody) SetRequestId(v string) *DescribeEventMetaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventMetaInfoResponseBody) SetTotalCount(v int32) *DescribeEventMetaInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventMetaInfoResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventMetaInfoResponseBodyItems struct {
	// example:
	//
	// Running
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Task.TaskStatus
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
}

func (s DescribeEventMetaInfoResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventMetaInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeEventMetaInfoResponseBodyItems) GetCode() *string {
	return s.Code
}

func (s *DescribeEventMetaInfoResponseBodyItems) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeEventMetaInfoResponseBodyItems) SetCode(v string) *DescribeEventMetaInfoResponseBodyItems {
	s.Code = &v
	return s
}

func (s *DescribeEventMetaInfoResponseBodyItems) SetSourceCode(v string) *DescribeEventMetaInfoResponseBodyItems {
	s.SourceCode = &v
	return s
}

func (s *DescribeEventMetaInfoResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
