// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnList(v *DescribeTableResponseBodyColumnList) *DescribeTableResponseBody
	GetColumnList() *DescribeTableResponseBodyColumnList
	SetMessage(v string) *DescribeTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTableResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeTableResponseBody
	GetStatus() *string
}

type DescribeTableResponseBody struct {
	// The columns of the table.
	ColumnList *DescribeTableResponseBodyColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBody) GetColumnList() *DescribeTableResponseBodyColumnList {
	return s.ColumnList
}

func (s *DescribeTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeTableResponseBody) SetColumnList(v *DescribeTableResponseBodyColumnList) *DescribeTableResponseBody {
	s.ColumnList = v
	return s
}

func (s *DescribeTableResponseBody) SetMessage(v string) *DescribeTableResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTableResponseBody) SetRequestId(v string) *DescribeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableResponseBody) SetStatus(v string) *DescribeTableResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeTableResponseBody) Validate() error {
	if s.ColumnList != nil {
		if err := s.ColumnList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTableResponseBodyColumnList struct {
	ColumnList []*ColumnMetadata `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
}

func (s DescribeTableResponseBodyColumnList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableResponseBodyColumnList) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyColumnList) GetColumnList() []*ColumnMetadata {
	return s.ColumnList
}

func (s *DescribeTableResponseBodyColumnList) SetColumnList(v []*ColumnMetadata) *DescribeTableResponseBodyColumnList {
	s.ColumnList = v
	return s
}

func (s *DescribeTableResponseBodyColumnList) Validate() error {
	if s.ColumnList != nil {
		for _, item := range s.ColumnList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
