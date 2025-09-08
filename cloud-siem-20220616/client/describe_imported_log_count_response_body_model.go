// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportedLogCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeImportedLogCountResponseBodyData) *DescribeImportedLogCountResponseBody
	GetData() *DescribeImportedLogCountResponseBodyData
	SetRequestId(v string) *DescribeImportedLogCountResponseBody
	GetRequestId() *string
}

type DescribeImportedLogCountResponseBody struct {
	// The data returned.
	Data *DescribeImportedLogCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImportedLogCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountResponseBody) GetData() *DescribeImportedLogCountResponseBodyData {
	return s.Data
}

func (s *DescribeImportedLogCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImportedLogCountResponseBody) SetData(v *DescribeImportedLogCountResponseBodyData) *DescribeImportedLogCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImportedLogCountResponseBody) SetRequestId(v string) *DescribeImportedLogCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImportedLogCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImportedLogCountResponseBodyData struct {
	// The number of logs that are added.
	//
	// example:
	//
	// 10
	ImportedLogCount *int32 `json:"ImportedLogCount,omitempty" xml:"ImportedLogCount,omitempty"`
	// The total number of logs.
	//
	// example:
	//
	// 59
	TotalLogCount *int32 `json:"TotalLogCount,omitempty" xml:"TotalLogCount,omitempty"`
	// The number of logs that are not added.
	//
	// example:
	//
	// 49
	UnImportedLogCount *int32 `json:"UnImportedLogCount,omitempty" xml:"UnImportedLogCount,omitempty"`
}

func (s DescribeImportedLogCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedLogCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImportedLogCountResponseBodyData) GetImportedLogCount() *int32 {
	return s.ImportedLogCount
}

func (s *DescribeImportedLogCountResponseBodyData) GetTotalLogCount() *int32 {
	return s.TotalLogCount
}

func (s *DescribeImportedLogCountResponseBodyData) GetUnImportedLogCount() *int32 {
	return s.UnImportedLogCount
}

func (s *DescribeImportedLogCountResponseBodyData) SetImportedLogCount(v int32) *DescribeImportedLogCountResponseBodyData {
	s.ImportedLogCount = &v
	return s
}

func (s *DescribeImportedLogCountResponseBodyData) SetTotalLogCount(v int32) *DescribeImportedLogCountResponseBodyData {
	s.TotalLogCount = &v
	return s
}

func (s *DescribeImportedLogCountResponseBodyData) SetUnImportedLogCount(v int32) *DescribeImportedLogCountResponseBodyData {
	s.UnImportedLogCount = &v
	return s
}

func (s *DescribeImportedLogCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
