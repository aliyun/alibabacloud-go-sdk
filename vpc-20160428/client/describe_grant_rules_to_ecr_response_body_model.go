// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGrantRulesToEcrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeGrantRulesToEcrResponseBody
	GetCount() *int32
	SetEcrGrantRules(v []*DescribeGrantRulesToEcrResponseBodyEcrGrantRules) *DescribeGrantRulesToEcrResponseBody
	GetEcrGrantRules() []*DescribeGrantRulesToEcrResponseBodyEcrGrantRules
	SetPageNumber(v int32) *DescribeGrantRulesToEcrResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGrantRulesToEcrResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGrantRulesToEcrResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGrantRulesToEcrResponseBody
	GetTotalCount() *int32
}

type DescribeGrantRulesToEcrResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The cross-account authorization list of the ECR
	EcrGrantRules []*DescribeGrantRulesToEcrResponseBodyEcrGrantRules `json:"EcrGrantRules,omitempty" xml:"EcrGrantRules,omitempty" type:"Repeated"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C1CCAB91-6AE6-50E3-AAA3-D0E5A2BC6ADE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGrantRulesToEcrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToEcrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToEcrResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGrantRulesToEcrResponseBody) GetEcrGrantRules() []*DescribeGrantRulesToEcrResponseBodyEcrGrantRules {
	return s.EcrGrantRules
}

func (s *DescribeGrantRulesToEcrResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGrantRulesToEcrResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGrantRulesToEcrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGrantRulesToEcrResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGrantRulesToEcrResponseBody) SetCount(v int32) *DescribeGrantRulesToEcrResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBody) SetEcrGrantRules(v []*DescribeGrantRulesToEcrResponseBodyEcrGrantRules) *DescribeGrantRulesToEcrResponseBody {
	s.EcrGrantRules = v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBody) SetPageNumber(v int32) *DescribeGrantRulesToEcrResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBody) SetPageSize(v int32) *DescribeGrantRulesToEcrResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBody) SetRequestId(v string) *DescribeGrantRulesToEcrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBody) SetTotalCount(v int32) *DescribeGrantRulesToEcrResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBody) Validate() error {
	if s.EcrGrantRules != nil {
		for _, item := range s.EcrGrantRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGrantRulesToEcrResponseBodyEcrGrantRules struct {
	// The authorization time. The time follows the ISO8601 standard and uses UTC time. The format is YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2025-09-15T14:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ECR account ID.
	//
	// example:
	//
	// ecr-xxxxxx
	EcrInstanceId *string `json:"EcrInstanceId,omitempty" xml:"EcrInstanceId,omitempty"`
	// The ECR account ID.
	//
	// example:
	//
	// 11111111111
	EcrUid *int64 `json:"EcrUid,omitempty" xml:"EcrUid,omitempty"`
}

func (s DescribeGrantRulesToEcrResponseBodyEcrGrantRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeGrantRulesToEcrResponseBodyEcrGrantRules) GoString() string {
	return s.String()
}

func (s *DescribeGrantRulesToEcrResponseBodyEcrGrantRules) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGrantRulesToEcrResponseBodyEcrGrantRules) GetEcrInstanceId() *string {
	return s.EcrInstanceId
}

func (s *DescribeGrantRulesToEcrResponseBodyEcrGrantRules) GetEcrUid() *int64 {
	return s.EcrUid
}

func (s *DescribeGrantRulesToEcrResponseBodyEcrGrantRules) SetCreateTime(v string) *DescribeGrantRulesToEcrResponseBodyEcrGrantRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBodyEcrGrantRules) SetEcrInstanceId(v string) *DescribeGrantRulesToEcrResponseBodyEcrGrantRules {
	s.EcrInstanceId = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBodyEcrGrantRules) SetEcrUid(v int64) *DescribeGrantRulesToEcrResponseBodyEcrGrantRules {
	s.EcrUid = &v
	return s
}

func (s *DescribeGrantRulesToEcrResponseBodyEcrGrantRules) Validate() error {
	return dara.Validate(s)
}
