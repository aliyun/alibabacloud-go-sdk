// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListCustomAttributesResponseBodyPagingInfo) *ListCustomAttributesResponseBody
	GetPagingInfo() *ListCustomAttributesResponseBodyPagingInfo
	SetRequestId(v string) *ListCustomAttributesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomAttributesResponseBody
	GetSuccess() *bool
}

type ListCustomAttributesResponseBody struct {
	PagingInfo *ListCustomAttributesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// 54594ACA-7976-5273-958B-02E15E9B867C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCustomAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAttributesResponseBody) GetPagingInfo() *ListCustomAttributesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListCustomAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomAttributesResponseBody) SetPagingInfo(v *ListCustomAttributesResponseBodyPagingInfo) *ListCustomAttributesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListCustomAttributesResponseBody) SetRequestId(v string) *ListCustomAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAttributesResponseBody) SetSuccess(v bool) *ListCustomAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomAttributesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomAttributesResponseBodyPagingInfo struct {
	CustomAttributes []*CustomAttribute `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty" type:"Repeated"`
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

func (s ListCustomAttributesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAttributesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListCustomAttributesResponseBodyPagingInfo) GetCustomAttributes() []*CustomAttribute {
	return s.CustomAttributes
}

func (s *ListCustomAttributesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomAttributesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomAttributesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomAttributesResponseBodyPagingInfo) SetCustomAttributes(v []*CustomAttribute) *ListCustomAttributesResponseBodyPagingInfo {
	s.CustomAttributes = v
	return s
}

func (s *ListCustomAttributesResponseBodyPagingInfo) SetPageNumber(v int32) *ListCustomAttributesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAttributesResponseBodyPagingInfo) SetPageSize(v int32) *ListCustomAttributesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListCustomAttributesResponseBodyPagingInfo) SetTotalCount(v int64) *ListCustomAttributesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCustomAttributesResponseBodyPagingInfo) Validate() error {
	if s.CustomAttributes != nil {
		for _, item := range s.CustomAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
