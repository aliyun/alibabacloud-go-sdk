// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStageModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeStageModelsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStageModelsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeStageModelsResponseBody
	GetRequestId() *string
	SetStageModelInfoList(v []*DescribeStageModelsResponseBodyStageModelInfoList) *DescribeStageModelsResponseBody
	GetStageModelInfoList() []*DescribeStageModelsResponseBodyStageModelInfoList
	SetTotalCount(v int32) *DescribeStageModelsResponseBody
	GetTotalCount() *int32
}

type DescribeStageModelsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZxxx
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StageModelInfoList []*DescribeStageModelsResponseBodyStageModelInfoList `json:"StageModelInfoList,omitempty" xml:"StageModelInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeStageModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStageModelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStageModelsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStageModelsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStageModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStageModelsResponseBody) GetStageModelInfoList() []*DescribeStageModelsResponseBodyStageModelInfoList {
	return s.StageModelInfoList
}

func (s *DescribeStageModelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeStageModelsResponseBody) SetPageNumber(v int32) *DescribeStageModelsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStageModelsResponseBody) SetPageSize(v int32) *DescribeStageModelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStageModelsResponseBody) SetRequestId(v string) *DescribeStageModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStageModelsResponseBody) SetStageModelInfoList(v []*DescribeStageModelsResponseBodyStageModelInfoList) *DescribeStageModelsResponseBody {
	s.StageModelInfoList = v
	return s
}

func (s *DescribeStageModelsResponseBody) SetTotalCount(v int32) *DescribeStageModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeStageModelsResponseBody) Validate() error {
	if s.StageModelInfoList != nil {
		for _, item := range s.StageModelInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStageModelsResponseBodyStageModelInfoList struct {
	// example:
	//
	// 2025-08-13T01:54:03Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// Stage description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2024-12-10T00:01:09Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	StageAlias   *string `json:"StageAlias,omitempty" xml:"StageAlias,omitempty"`
	// example:
	//
	// j3j435j23l4j23l55xxx
	StageModelId *string `json:"StageModelId,omitempty" xml:"StageModelId,omitempty"`
	// example:
	//
	// PRE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeStageModelsResponseBodyStageModelInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStageModelsResponseBodyStageModelInfoList) GoString() string {
	return s.String()
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) GetDescription() *string {
	return s.Description
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) GetStageAlias() *string {
	return s.StageAlias
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) GetStageModelId() *string {
	return s.StageModelId
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) GetStageName() *string {
	return s.StageName
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) GetType() *string {
	return s.Type
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) SetCreatedTime(v string) *DescribeStageModelsResponseBodyStageModelInfoList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) SetDescription(v string) *DescribeStageModelsResponseBodyStageModelInfoList {
	s.Description = &v
	return s
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) SetModifiedTime(v string) *DescribeStageModelsResponseBodyStageModelInfoList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) SetStageAlias(v string) *DescribeStageModelsResponseBodyStageModelInfoList {
	s.StageAlias = &v
	return s
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) SetStageModelId(v string) *DescribeStageModelsResponseBodyStageModelInfoList {
	s.StageModelId = &v
	return s
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) SetStageName(v string) *DescribeStageModelsResponseBodyStageModelInfoList {
	s.StageName = &v
	return s
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) SetType(v string) *DescribeStageModelsResponseBodyStageModelInfoList {
	s.Type = &v
	return s
}

func (s *DescribeStageModelsResponseBodyStageModelInfoList) Validate() error {
	return dara.Validate(s)
}
