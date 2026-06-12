// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillSpacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSkillSpacesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillSpacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSkillSpacesResponseBody
	GetRequestId() *string
	SetSkillSpaces(v []*ListSkillSpacesResponseBodySkillSpaces) *ListSkillSpacesResponseBody
	GetSkillSpaces() []*ListSkillSpacesResponseBodySkillSpaces
	SetTotalCount(v int32) *ListSkillSpacesResponseBody
	GetTotalCount() *int32
}

type ListSkillSpacesResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. If this parameter is not returned, no more results are available.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of SkillSpaces.
	SkillSpaces []*ListSkillSpacesResponseBodySkillSpaces `json:"SkillSpaces,omitempty" xml:"SkillSpaces,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillSpacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillSpacesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillSpacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillSpacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillSpacesResponseBody) GetSkillSpaces() []*ListSkillSpacesResponseBodySkillSpaces {
	return s.SkillSpaces
}

func (s *ListSkillSpacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillSpacesResponseBody) SetMaxResults(v int32) *ListSkillSpacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSkillSpacesResponseBody) SetNextToken(v string) *ListSkillSpacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSkillSpacesResponseBody) SetRequestId(v string) *ListSkillSpacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillSpacesResponseBody) SetSkillSpaces(v []*ListSkillSpacesResponseBodySkillSpaces) *ListSkillSpacesResponseBody {
	s.SkillSpaces = v
	return s
}

func (s *ListSkillSpacesResponseBody) SetTotalCount(v int32) *ListSkillSpacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSkillSpacesResponseBody) Validate() error {
	if s.SkillSpaces != nil {
		for _, item := range s.SkillSpaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillSpacesResponseBodySkillSpaces struct {
	// The time when the SkillSpace was created. The time is in UTC.
	//
	// example:
	//
	// 2025-11-03T22:58:52Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The SkillSpace description.
	//
	// example:
	//
	// 1111
	SkillSpaceDescription *string `json:"SkillSpaceDescription,omitempty" xml:"SkillSpaceDescription,omitempty"`
	// The SkillSpace ID.
	//
	// example:
	//
	// ss-111111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// The SkillSpace name.
	//
	// example:
	//
	// 1111
	SkillSpaceName *string `json:"SkillSpaceName,omitempty" xml:"SkillSpaceName,omitempty"`
	// The time when the SkillSpace was last updated. The time is in UTC.
	//
	// example:
	//
	// 2025-11-03T22:57:29Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSkillSpacesResponseBodySkillSpaces) String() string {
	return dara.Prettify(s)
}

func (s ListSkillSpacesResponseBodySkillSpaces) GoString() string {
	return s.String()
}

func (s *ListSkillSpacesResponseBodySkillSpaces) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSkillSpacesResponseBodySkillSpaces) GetSkillSpaceDescription() *string {
	return s.SkillSpaceDescription
}

func (s *ListSkillSpacesResponseBodySkillSpaces) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *ListSkillSpacesResponseBodySkillSpaces) GetSkillSpaceName() *string {
	return s.SkillSpaceName
}

func (s *ListSkillSpacesResponseBodySkillSpaces) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSkillSpacesResponseBodySkillSpaces) SetCreateTime(v string) *ListSkillSpacesResponseBodySkillSpaces {
	s.CreateTime = &v
	return s
}

func (s *ListSkillSpacesResponseBodySkillSpaces) SetSkillSpaceDescription(v string) *ListSkillSpacesResponseBodySkillSpaces {
	s.SkillSpaceDescription = &v
	return s
}

func (s *ListSkillSpacesResponseBodySkillSpaces) SetSkillSpaceId(v string) *ListSkillSpacesResponseBodySkillSpaces {
	s.SkillSpaceId = &v
	return s
}

func (s *ListSkillSpacesResponseBodySkillSpaces) SetSkillSpaceName(v string) *ListSkillSpacesResponseBodySkillSpaces {
	s.SkillSpaceName = &v
	return s
}

func (s *ListSkillSpacesResponseBodySkillSpaces) SetUpdateTime(v string) *ListSkillSpacesResponseBodySkillSpaces {
	s.UpdateTime = &v
	return s
}

func (s *ListSkillSpacesResponseBodySkillSpaces) Validate() error {
	return dara.Validate(s)
}
