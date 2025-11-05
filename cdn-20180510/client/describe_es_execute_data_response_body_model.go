// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEsExecuteDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContents(v []*DescribeEsExecuteDataResponseBodyContents) *DescribeEsExecuteDataResponseBody
	GetContents() []*DescribeEsExecuteDataResponseBodyContents
	SetRequestId(v string) *DescribeEsExecuteDataResponseBody
	GetRequestId() *string
}

type DescribeEsExecuteDataResponseBody struct {
	// The content of the script.
	Contents []*DescribeEsExecuteDataResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 99D61AB3-6164-4CF2-A0DE-129C9B07618B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEsExecuteDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExecuteDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataResponseBody) GetContents() []*DescribeEsExecuteDataResponseBodyContents {
	return s.Contents
}

func (s *DescribeEsExecuteDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEsExecuteDataResponseBody) SetContents(v []*DescribeEsExecuteDataResponseBodyContents) *DescribeEsExecuteDataResponseBody {
	s.Contents = v
	return s
}

func (s *DescribeEsExecuteDataResponseBody) SetRequestId(v string) *DescribeEsExecuteDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEsExecuteDataResponseBody) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEsExecuteDataResponseBodyContents struct {
	// The time and column names in the table that shows the status of the script.
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The name of the table that shows the status of the script.
	//
	// example:
	//
	// Exception
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of timestamps and values in the corresponding columns of the table that shows the status of the script.
	Points []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeEsExecuteDataResponseBodyContents) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExecuteDataResponseBodyContents) GoString() string {
	return s.String()
}

func (s *DescribeEsExecuteDataResponseBodyContents) GetColumns() []*string {
	return s.Columns
}

func (s *DescribeEsExecuteDataResponseBodyContents) GetName() *string {
	return s.Name
}

func (s *DescribeEsExecuteDataResponseBodyContents) GetPoints() []*string {
	return s.Points
}

func (s *DescribeEsExecuteDataResponseBodyContents) SetColumns(v []*string) *DescribeEsExecuteDataResponseBodyContents {
	s.Columns = v
	return s
}

func (s *DescribeEsExecuteDataResponseBodyContents) SetName(v string) *DescribeEsExecuteDataResponseBodyContents {
	s.Name = &v
	return s
}

func (s *DescribeEsExecuteDataResponseBodyContents) SetPoints(v []*string) *DescribeEsExecuteDataResponseBodyContents {
	s.Points = v
	return s
}

func (s *DescribeEsExecuteDataResponseBodyContents) Validate() error {
	return dara.Validate(s)
}
