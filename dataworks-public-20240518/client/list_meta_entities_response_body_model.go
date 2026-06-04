// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListMetaEntitiesResponseBodyPagingInfo) *ListMetaEntitiesResponseBody
	GetPagingInfo() *ListMetaEntitiesResponseBodyPagingInfo
	SetRequestId(v string) *ListMetaEntitiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaEntitiesResponseBody
	GetSuccess() *bool
}

type ListMetaEntitiesResponseBody struct {
	PagingInfo *ListMetaEntitiesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// ADFASDFASDFA-ADFASDF-ASDFADSDF-AFFADS
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMetaEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaEntitiesResponseBody) GetPagingInfo() *ListMetaEntitiesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListMetaEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaEntitiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaEntitiesResponseBody) SetPagingInfo(v *ListMetaEntitiesResponseBodyPagingInfo) *ListMetaEntitiesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListMetaEntitiesResponseBody) SetRequestId(v string) *ListMetaEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaEntitiesResponseBody) SetSuccess(v bool) *ListMetaEntitiesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaEntitiesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMetaEntitiesResponseBodyPagingInfo struct {
	// example:
	//
	// 10
	MaxResults   *int32        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	MetaEntities []*MetaEntity `json:"MetaEntities,omitempty" xml:"MetaEntities,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAAdEdsXbwG2ZlbWCzN4wTTg6NmTbhyvglcoMCJsiEdngaTov15YaMyduvjIHYeTOIcEeXqCevM1qffZkwCkUTUYc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetaEntitiesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntitiesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) GetMetaEntities() []*MetaEntity {
	return s.MetaEntities
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) SetMaxResults(v int32) *ListMetaEntitiesResponseBodyPagingInfo {
	s.MaxResults = &v
	return s
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) SetMetaEntities(v []*MetaEntity) *ListMetaEntitiesResponseBodyPagingInfo {
	s.MetaEntities = v
	return s
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) SetNextToken(v string) *ListMetaEntitiesResponseBodyPagingInfo {
	s.NextToken = &v
	return s
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) SetTotalCount(v int64) *ListMetaEntitiesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListMetaEntitiesResponseBodyPagingInfo) Validate() error {
	if s.MetaEntities != nil {
		for _, item := range s.MetaEntities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
