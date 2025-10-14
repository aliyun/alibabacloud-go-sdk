// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRuleVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNormalizationRuleVersionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRuleVersionsResponseBody
	GetNextToken() *string
	SetNormalizationRuleVersions(v []*ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) *ListNormalizationRuleVersionsResponseBody
	GetNormalizationRuleVersions() []*ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions
	SetRequestId(v string) *ListNormalizationRuleVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNormalizationRuleVersionsResponseBody
	GetTotalCount() *int32
}

type ListNormalizationRuleVersionsResponseBody struct {
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken                 *string                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NormalizationRuleVersions []*ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions `json:"NormalizationRuleVersions,omitempty" xml:"NormalizationRuleVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNormalizationRuleVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRuleVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRuleVersionsResponseBody) GetNormalizationRuleVersions() []*ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions {
	return s.NormalizationRuleVersions
}

func (s *ListNormalizationRuleVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNormalizationRuleVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNormalizationRuleVersionsResponseBody) SetMaxResults(v int32) *ListNormalizationRuleVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBody) SetNextToken(v string) *ListNormalizationRuleVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBody) SetNormalizationRuleVersions(v []*ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) *ListNormalizationRuleVersionsResponseBody {
	s.NormalizationRuleVersions = v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBody) SetRequestId(v string) *ListNormalizationRuleVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBody) SetTotalCount(v int32) *ListNormalizationRuleVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBody) Validate() error {
	if s.NormalizationRuleVersions != nil {
		for _, item := range s.NormalizationRuleVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions struct {
	// example:
	//
	// 1733269771123。
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 	- | pack-fields -include=\\"[\\s\\S]+\\" as extend_content。
	NormalizationRuleExpression *string `json:"NormalizationRuleExpression,omitempty" xml:"NormalizationRuleExpression,omitempty"`
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp。
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// example:
	//
	// V1。
	NormalizationRuleVersion *int32 `json:"NormalizationRuleVersion,omitempty" xml:"NormalizationRuleVersion,omitempty"`
	// example:
	//
	// 1733269771123。
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) GetNormalizationRuleExpression() *string {
	return s.NormalizationRuleExpression
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) GetNormalizationRuleVersion() *int32 {
	return s.NormalizationRuleVersion
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) SetCreateTime(v int64) *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions {
	s.CreateTime = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) SetNormalizationRuleExpression(v string) *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions {
	s.NormalizationRuleExpression = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) SetNormalizationRuleId(v string) *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) SetNormalizationRuleVersion(v int32) *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions {
	s.NormalizationRuleVersion = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) SetUpdateTime(v int64) *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListNormalizationRuleVersionsResponseBodyNormalizationRuleVersions) Validate() error {
	return dara.Validate(s)
}
