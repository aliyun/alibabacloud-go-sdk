// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaEntityDefsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListMetaEntityDefsResponseBodyPagingInfo) *ListMetaEntityDefsResponseBody
	GetPagingInfo() *ListMetaEntityDefsResponseBodyPagingInfo
	SetRequestId(v string) *ListMetaEntityDefsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaEntityDefsResponseBody
	GetSuccess() *bool
}

type ListMetaEntityDefsResponseBody struct {
	PagingInfo *ListMetaEntityDefsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
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

func (s ListMetaEntityDefsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntityDefsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaEntityDefsResponseBody) GetPagingInfo() *ListMetaEntityDefsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListMetaEntityDefsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaEntityDefsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaEntityDefsResponseBody) SetPagingInfo(v *ListMetaEntityDefsResponseBodyPagingInfo) *ListMetaEntityDefsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListMetaEntityDefsResponseBody) SetRequestId(v string) *ListMetaEntityDefsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaEntityDefsResponseBody) SetSuccess(v bool) *ListMetaEntityDefsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaEntityDefsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMetaEntityDefsResponseBodyPagingInfo struct {
	MetaEntityDefs []*MetaEntityDef `json:"MetaEntityDefs,omitempty" xml:"MetaEntityDefs,omitempty" type:"Repeated"`
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
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetaEntityDefsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMetaEntityDefsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) GetMetaEntityDefs() []*MetaEntityDef {
	return s.MetaEntityDefs
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) SetMetaEntityDefs(v []*MetaEntityDef) *ListMetaEntityDefsResponseBodyPagingInfo {
	s.MetaEntityDefs = v
	return s
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) SetPageNumber(v int32) *ListMetaEntityDefsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) SetPageSize(v int32) *ListMetaEntityDefsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) SetTotalCount(v int64) *ListMetaEntityDefsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListMetaEntityDefsResponseBodyPagingInfo) Validate() error {
	if s.MetaEntityDefs != nil {
		for _, item := range s.MetaEntityDefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
