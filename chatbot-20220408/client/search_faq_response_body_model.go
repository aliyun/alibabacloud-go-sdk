// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFaqResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFaqHits(v []*SearchFaqResponseBodyFaqHits) *SearchFaqResponseBody
	GetFaqHits() []*SearchFaqResponseBodyFaqHits
	SetPageNumber(v int32) *SearchFaqResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchFaqResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchFaqResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *SearchFaqResponseBody
	GetTotalCount() *int32
}

type SearchFaqResponseBody struct {
	FaqHits []*SearchFaqResponseBodyFaqHits `json:"FaqHits,omitempty" xml:"FaqHits,omitempty" type:"Repeated"`
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
	// E45491D5-7E0A-42C6-9B21-91D1066B1475
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1075
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchFaqResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchFaqResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFaqResponseBody) GetFaqHits() []*SearchFaqResponseBodyFaqHits {
	return s.FaqHits
}

func (s *SearchFaqResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchFaqResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchFaqResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchFaqResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchFaqResponseBody) SetFaqHits(v []*SearchFaqResponseBodyFaqHits) *SearchFaqResponseBody {
	s.FaqHits = v
	return s
}

func (s *SearchFaqResponseBody) SetPageNumber(v int32) *SearchFaqResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFaqResponseBody) SetPageSize(v int32) *SearchFaqResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchFaqResponseBody) SetRequestId(v string) *SearchFaqResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFaqResponseBody) SetTotalCount(v int32) *SearchFaqResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFaqResponseBody) Validate() error {
	if s.FaqHits != nil {
		for _, item := range s.FaqHits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchFaqResponseBodyFaqHits struct {
	// example:
	//
	// 30000055639
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 18453
	CreateUserId *int64 `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// test01
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 20
	EffectStatus     *int32    `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	HitSimilarTitles []*string `json:"HitSimilarTitles,omitempty" xml:"HitSimilarTitles,omitempty" type:"Repeated"`
	HitSolutions     []*string `json:"HitSolutions,omitempty" xml:"HitSolutions,omitempty" type:"Repeated"`
	// example:
	//
	// 30002145804
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 18453
	ModifyUserId *int64 `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// test01
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 3
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchFaqResponseBodyFaqHits) String() string {
	return dara.Prettify(s)
}

func (s SearchFaqResponseBodyFaqHits) GoString() string {
	return s.String()
}

func (s *SearchFaqResponseBodyFaqHits) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *SearchFaqResponseBodyFaqHits) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchFaqResponseBodyFaqHits) GetCreateUserId() *int64 {
	return s.CreateUserId
}

func (s *SearchFaqResponseBodyFaqHits) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *SearchFaqResponseBodyFaqHits) GetEffectStatus() *int32 {
	return s.EffectStatus
}

func (s *SearchFaqResponseBodyFaqHits) GetHitSimilarTitles() []*string {
	return s.HitSimilarTitles
}

func (s *SearchFaqResponseBodyFaqHits) GetHitSolutions() []*string {
	return s.HitSolutions
}

func (s *SearchFaqResponseBodyFaqHits) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *SearchFaqResponseBodyFaqHits) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *SearchFaqResponseBodyFaqHits) GetModifyUserId() *int64 {
	return s.ModifyUserId
}

func (s *SearchFaqResponseBodyFaqHits) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *SearchFaqResponseBodyFaqHits) GetStatus() *int32 {
	return s.Status
}

func (s *SearchFaqResponseBodyFaqHits) GetTitle() *string {
	return s.Title
}

func (s *SearchFaqResponseBodyFaqHits) SetCategoryId(v int64) *SearchFaqResponseBodyFaqHits {
	s.CategoryId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetCreateTime(v string) *SearchFaqResponseBodyFaqHits {
	s.CreateTime = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetCreateUserId(v int64) *SearchFaqResponseBodyFaqHits {
	s.CreateUserId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetCreateUserName(v string) *SearchFaqResponseBodyFaqHits {
	s.CreateUserName = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetEffectStatus(v int32) *SearchFaqResponseBodyFaqHits {
	s.EffectStatus = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetHitSimilarTitles(v []*string) *SearchFaqResponseBodyFaqHits {
	s.HitSimilarTitles = v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetHitSolutions(v []*string) *SearchFaqResponseBodyFaqHits {
	s.HitSolutions = v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetKnowledgeId(v int64) *SearchFaqResponseBodyFaqHits {
	s.KnowledgeId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetModifyTime(v string) *SearchFaqResponseBodyFaqHits {
	s.ModifyTime = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetModifyUserId(v int64) *SearchFaqResponseBodyFaqHits {
	s.ModifyUserId = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetModifyUserName(v string) *SearchFaqResponseBodyFaqHits {
	s.ModifyUserName = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetStatus(v int32) *SearchFaqResponseBodyFaqHits {
	s.Status = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) SetTitle(v string) *SearchFaqResponseBodyFaqHits {
	s.Title = &v
	return s
}

func (s *SearchFaqResponseBodyFaqHits) Validate() error {
	return dara.Validate(s)
}
