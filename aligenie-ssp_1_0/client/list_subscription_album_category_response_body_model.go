// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionAlbumCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSubscriptionAlbumCategoryResponseBody
	GetCode() *int32
	SetMessage(v string) *ListSubscriptionAlbumCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSubscriptionAlbumCategoryResponseBody
	GetRequestId() *string
	SetResult(v []*ListSubscriptionAlbumCategoryResponseBodyResult) *ListSubscriptionAlbumCategoryResponseBody
	GetResult() []*ListSubscriptionAlbumCategoryResponseBodyResult
}

type ListSubscriptionAlbumCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 60E7A523-9766-1D07-87A2-6E587420C59B
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListSubscriptionAlbumCategoryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSubscriptionAlbumCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSubscriptionAlbumCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubscriptionAlbumCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubscriptionAlbumCategoryResponseBody) GetResult() []*ListSubscriptionAlbumCategoryResponseBodyResult {
	return s.Result
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetCode(v int32) *ListSubscriptionAlbumCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetMessage(v string) *ListSubscriptionAlbumCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetRequestId(v string) *ListSubscriptionAlbumCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBody) SetResult(v []*ListSubscriptionAlbumCategoryResponseBodyResult) *ListSubscriptionAlbumCategoryResponseBody {
	s.Result = v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSubscriptionAlbumCategoryResponseBodyResult struct {
	// example:
	//
	// 80011
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 儿童
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s ListSubscriptionAlbumCategoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryResponseBodyResult) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListSubscriptionAlbumCategoryResponseBodyResult) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListSubscriptionAlbumCategoryResponseBodyResult) SetCategoryId(v string) *ListSubscriptionAlbumCategoryResponseBodyResult {
	s.CategoryId = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBodyResult) SetCategoryName(v string) *ListSubscriptionAlbumCategoryResponseBodyResult {
	s.CategoryName = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
