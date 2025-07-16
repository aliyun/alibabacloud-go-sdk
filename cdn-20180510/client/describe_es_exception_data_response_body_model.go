// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEsExceptionDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContents(v []*DescribeEsExceptionDataResponseBodyContents) *DescribeEsExceptionDataResponseBody
	GetContents() []*DescribeEsExceptionDataResponseBodyContents
	SetRequestId(v string) *DescribeEsExceptionDataResponseBody
	GetRequestId() *string
}

type DescribeEsExceptionDataResponseBody struct {
	// The content of the script for which an error was reported.
	Contents []*DescribeEsExceptionDataResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 99D61AB3-6164-4CF2-A0DE-129C9B07618B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEsExceptionDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExceptionDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataResponseBody) GetContents() []*DescribeEsExceptionDataResponseBodyContents {
	return s.Contents
}

func (s *DescribeEsExceptionDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEsExceptionDataResponseBody) SetContents(v []*DescribeEsExceptionDataResponseBodyContents) *DescribeEsExceptionDataResponseBody {
	s.Contents = v
	return s
}

func (s *DescribeEsExceptionDataResponseBody) SetRequestId(v string) *DescribeEsExceptionDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEsExceptionDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEsExceptionDataResponseBodyContents struct {
	// Information about the time column and the error column name.
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The name of the table that shows the errors of the script.
	//
	// example:
	//
	// 401
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time columns and error column names.
	Points []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeEsExceptionDataResponseBodyContents) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExceptionDataResponseBodyContents) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataResponseBodyContents) GetColumns() []*string {
	return s.Columns
}

func (s *DescribeEsExceptionDataResponseBodyContents) GetName() *string {
	return s.Name
}

func (s *DescribeEsExceptionDataResponseBodyContents) GetPoints() []*string {
	return s.Points
}

func (s *DescribeEsExceptionDataResponseBodyContents) SetColumns(v []*string) *DescribeEsExceptionDataResponseBodyContents {
	s.Columns = v
	return s
}

func (s *DescribeEsExceptionDataResponseBodyContents) SetName(v string) *DescribeEsExceptionDataResponseBodyContents {
	s.Name = &v
	return s
}

func (s *DescribeEsExceptionDataResponseBodyContents) SetPoints(v []*string) *DescribeEsExceptionDataResponseBodyContents {
	s.Points = v
	return s
}

func (s *DescribeEsExceptionDataResponseBodyContents) Validate() error {
	return dara.Validate(s)
}
