// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListResourceTagsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListResourceTagsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListResourceTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListResourceTagsResponseBody
	GetRequestId() *string
	SetResourceTags(v *ListResourceTagsResponseBodyResourceTags) *ListResourceTagsResponseBody
	GetResourceTags() *ListResourceTagsResponseBodyResourceTags
	SetSuccess(v bool) *ListResourceTagsResponseBody
	GetSuccess() *bool
}

type ListResourceTagsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {}
	ResourceTags *ListResourceTagsResponseBodyResourceTags `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourceTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListResourceTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListResourceTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListResourceTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceTagsResponseBody) GetResourceTags() *ListResourceTagsResponseBodyResourceTags {
	return s.ResourceTags
}

func (s *ListResourceTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourceTagsResponseBody) SetCode(v string) *ListResourceTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListResourceTagsResponseBody) SetHttpStatusCode(v int32) *ListResourceTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListResourceTagsResponseBody) SetMessage(v string) *ListResourceTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListResourceTagsResponseBody) SetRequestId(v string) *ListResourceTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTagsResponseBody) SetResourceTags(v *ListResourceTagsResponseBodyResourceTags) *ListResourceTagsResponseBody {
	s.ResourceTags = v
	return s
}

func (s *ListResourceTagsResponseBody) SetSuccess(v bool) *ListResourceTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourceTagsResponseBody) Validate() error {
	if s.ResourceTags != nil {
		if err := s.ResourceTags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceTagsResponseBodyResourceTags struct {
	// example:
	//
	// []
	List []*ListResourceTagsResponseBodyResourceTagsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceTagsResponseBodyResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsResponseBodyResourceTags) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBodyResourceTags) GetList() []*ListResourceTagsResponseBodyResourceTagsList {
	return s.List
}

func (s *ListResourceTagsResponseBodyResourceTags) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceTagsResponseBodyResourceTags) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceTagsResponseBodyResourceTags) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceTagsResponseBodyResourceTags) SetList(v []*ListResourceTagsResponseBodyResourceTagsList) *ListResourceTagsResponseBodyResourceTags {
	s.List = v
	return s
}

func (s *ListResourceTagsResponseBodyResourceTags) SetPageNumber(v int32) *ListResourceTagsResponseBodyResourceTags {
	s.PageNumber = &v
	return s
}

func (s *ListResourceTagsResponseBodyResourceTags) SetPageSize(v int32) *ListResourceTagsResponseBodyResourceTags {
	s.PageSize = &v
	return s
}

func (s *ListResourceTagsResponseBodyResourceTags) SetTotalCount(v int32) *ListResourceTagsResponseBodyResourceTags {
	s.TotalCount = &v
	return s
}

func (s *ListResourceTagsResponseBodyResourceTags) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceTagsResponseBodyResourceTagsList struct {
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// xxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceTagsResponseBodyResourceTagsList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTagsResponseBodyResourceTagsList) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseBodyResourceTagsList) GetKey() *string {
	return s.Key
}

func (s *ListResourceTagsResponseBodyResourceTagsList) GetValue() *string {
	return s.Value
}

func (s *ListResourceTagsResponseBodyResourceTagsList) SetKey(v string) *ListResourceTagsResponseBodyResourceTagsList {
	s.Key = &v
	return s
}

func (s *ListResourceTagsResponseBodyResourceTagsList) SetValue(v string) *ListResourceTagsResponseBodyResourceTagsList {
	s.Value = &v
	return s
}

func (s *ListResourceTagsResponseBodyResourceTagsList) Validate() error {
	return dara.Validate(s)
}
