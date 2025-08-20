// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsHiveProgress(v []*DescribeApsProgressResponseBodyApsHiveProgress) *DescribeApsProgressResponseBody
	GetApsHiveProgress() []*DescribeApsProgressResponseBodyApsHiveProgress
	SetRequestId(v string) *DescribeApsProgressResponseBody
	GetRequestId() *string
	SetSuccessPercentage(v int32) *DescribeApsProgressResponseBody
	GetSuccessPercentage() *int32
	SetSuccessTableCount(v int32) *DescribeApsProgressResponseBody
	GetSuccessTableCount() *int32
	SetTotalTableCount(v int32) *DescribeApsProgressResponseBody
	GetTotalTableCount() *int32
}

type DescribeApsProgressResponseBody struct {
	// The migration progress.
	//
	// example:
	//
	// -
	ApsHiveProgress []*DescribeApsProgressResponseBodyApsHiveProgress `json:"ApsHiveProgress,omitempty" xml:"ApsHiveProgress,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The success rate.
	//
	// example:
	//
	// 100
	SuccessPercentage *int32 `json:"SuccessPercentage,omitempty" xml:"SuccessPercentage,omitempty"`
	// The total number of migrated tables returned.
	//
	// example:
	//
	// 10
	SuccessTableCount *int32 `json:"SuccessTableCount,omitempty" xml:"SuccessTableCount,omitempty"`
	// The total number of tables to be migrated.
	//
	// example:
	//
	// 10
	TotalTableCount *int32 `json:"TotalTableCount,omitempty" xml:"TotalTableCount,omitempty"`
}

func (s DescribeApsProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsProgressResponseBody) GetApsHiveProgress() []*DescribeApsProgressResponseBodyApsHiveProgress {
	return s.ApsHiveProgress
}

func (s *DescribeApsProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsProgressResponseBody) GetSuccessPercentage() *int32 {
	return s.SuccessPercentage
}

func (s *DescribeApsProgressResponseBody) GetSuccessTableCount() *int32 {
	return s.SuccessTableCount
}

func (s *DescribeApsProgressResponseBody) GetTotalTableCount() *int32 {
	return s.TotalTableCount
}

func (s *DescribeApsProgressResponseBody) SetApsHiveProgress(v []*DescribeApsProgressResponseBodyApsHiveProgress) *DescribeApsProgressResponseBody {
	s.ApsHiveProgress = v
	return s
}

func (s *DescribeApsProgressResponseBody) SetRequestId(v string) *DescribeApsProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsProgressResponseBody) SetSuccessPercentage(v int32) *DescribeApsProgressResponseBody {
	s.SuccessPercentage = &v
	return s
}

func (s *DescribeApsProgressResponseBody) SetSuccessTableCount(v int32) *DescribeApsProgressResponseBody {
	s.SuccessTableCount = &v
	return s
}

func (s *DescribeApsProgressResponseBody) SetTotalTableCount(v int32) *DescribeApsProgressResponseBody {
	s.TotalTableCount = &v
	return s
}

func (s *DescribeApsProgressResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApsProgressResponseBodyApsHiveProgress struct {
	// The name of the database.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The migration progress.
	//
	// example:
	//
	// 95
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The migration speed.
	//
	// example:
	//
	// 2
	Speed *string `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TbName *string `json:"TbName,omitempty" xml:"TbName,omitempty"`
}

func (s DescribeApsProgressResponseBodyApsHiveProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsProgressResponseBodyApsHiveProgress) GoString() string {
	return s.String()
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) GetDbName() *string {
	return s.DbName
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) GetProgress() *string {
	return s.Progress
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) GetSpeed() *string {
	return s.Speed
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) GetTbName() *string {
	return s.TbName
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) SetDbName(v string) *DescribeApsProgressResponseBodyApsHiveProgress {
	s.DbName = &v
	return s
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) SetProgress(v string) *DescribeApsProgressResponseBodyApsHiveProgress {
	s.Progress = &v
	return s
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) SetSpeed(v string) *DescribeApsProgressResponseBodyApsHiveProgress {
	s.Speed = &v
	return s
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) SetTbName(v string) *DescribeApsProgressResponseBodyApsHiveProgress {
	s.TbName = &v
	return s
}

func (s *DescribeApsProgressResponseBodyApsHiveProgress) Validate() error {
	return dara.Validate(s)
}
