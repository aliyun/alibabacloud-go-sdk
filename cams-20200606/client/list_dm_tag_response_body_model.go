// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDmTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListDmTagResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListDmTagResponseBody
	GetCode() *string
	SetData(v []*ListDmTagResponseBodyData) *ListDmTagResponseBody
	GetData() []*ListDmTagResponseBodyData
	SetMessage(v string) *ListDmTagResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDmTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDmTagResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListDmTagResponseBody
	GetTotal() *int64
}

type ListDmTagResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListDmTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// hgfh77-gfh55***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 42
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDmTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDmTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListDmTagResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListDmTagResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDmTagResponseBody) GetData() []*ListDmTagResponseBodyData {
	return s.Data
}

func (s *ListDmTagResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDmTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDmTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDmTagResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListDmTagResponseBody) SetAccessDeniedDetail(v string) *ListDmTagResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListDmTagResponseBody) SetCode(v string) *ListDmTagResponseBody {
	s.Code = &v
	return s
}

func (s *ListDmTagResponseBody) SetData(v []*ListDmTagResponseBodyData) *ListDmTagResponseBody {
	s.Data = v
	return s
}

func (s *ListDmTagResponseBody) SetMessage(v string) *ListDmTagResponseBody {
	s.Message = &v
	return s
}

func (s *ListDmTagResponseBody) SetRequestId(v string) *ListDmTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDmTagResponseBody) SetSuccess(v bool) *ListDmTagResponseBody {
	s.Success = &v
	return s
}

func (s *ListDmTagResponseBody) SetTotal(v int64) *ListDmTagResponseBody {
	s.Total = &v
	return s
}

func (s *ListDmTagResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDmTagResponseBodyData struct {
	// example:
	//
	// xx
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// example:
	//
	// xx
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// xx
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListDmTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDmTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDmTagResponseBodyData) GetTagDescription() *string {
	return s.TagDescription
}

func (s *ListDmTagResponseBodyData) GetTagId() *string {
	return s.TagId
}

func (s *ListDmTagResponseBodyData) GetTagName() *string {
	return s.TagName
}

func (s *ListDmTagResponseBodyData) SetTagDescription(v string) *ListDmTagResponseBodyData {
	s.TagDescription = &v
	return s
}

func (s *ListDmTagResponseBodyData) SetTagId(v string) *ListDmTagResponseBodyData {
	s.TagId = &v
	return s
}

func (s *ListDmTagResponseBodyData) SetTagName(v string) *ListDmTagResponseBodyData {
	s.TagName = &v
	return s
}

func (s *ListDmTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
