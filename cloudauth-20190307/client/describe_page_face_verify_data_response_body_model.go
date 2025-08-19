// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePageFaceVerifyDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePageFaceVerifyDataResponseBody
	GetCode() *string
	SetCurrentPage(v string) *DescribePageFaceVerifyDataResponseBody
	GetCurrentPage() *string
	SetItems(v []*DescribePageFaceVerifyDataResponseBodyItems) *DescribePageFaceVerifyDataResponseBody
	GetItems() []*DescribePageFaceVerifyDataResponseBodyItems
	SetMessage(v string) *DescribePageFaceVerifyDataResponseBody
	GetMessage() *string
	SetPageSize(v string) *DescribePageFaceVerifyDataResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribePageFaceVerifyDataResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribePageFaceVerifyDataResponseBody
	GetSuccess() *string
	SetTotalCount(v string) *DescribePageFaceVerifyDataResponseBody
	GetTotalCount() *string
	SetTotalPage(v string) *DescribePageFaceVerifyDataResponseBody
	GetTotalPage() *string
}

type DescribePageFaceVerifyDataResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *string                                        `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribePageFaceVerifyDataResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 473469C7-A***B-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 5
	TotalPage *string `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribePageFaceVerifyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePageFaceVerifyDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePageFaceVerifyDataResponseBody) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribePageFaceVerifyDataResponseBody) GetItems() []*DescribePageFaceVerifyDataResponseBodyItems {
	return s.Items
}

func (s *DescribePageFaceVerifyDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePageFaceVerifyDataResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribePageFaceVerifyDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePageFaceVerifyDataResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribePageFaceVerifyDataResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribePageFaceVerifyDataResponseBody) GetTotalPage() *string {
	return s.TotalPage
}

func (s *DescribePageFaceVerifyDataResponseBody) SetCode(v string) *DescribePageFaceVerifyDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetCurrentPage(v string) *DescribePageFaceVerifyDataResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetItems(v []*DescribePageFaceVerifyDataResponseBodyItems) *DescribePageFaceVerifyDataResponseBody {
	s.Items = v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetMessage(v string) *DescribePageFaceVerifyDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetPageSize(v string) *DescribePageFaceVerifyDataResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetRequestId(v string) *DescribePageFaceVerifyDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetSuccess(v string) *DescribePageFaceVerifyDataResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetTotalCount(v string) *DescribePageFaceVerifyDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) SetTotalPage(v string) *DescribePageFaceVerifyDataResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePageFaceVerifyDataResponseBodyItems struct {
	// example:
	//
	// 2024-03-24T00:00:00.000Z
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// ID_PLUS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 20**40
	SceneId   *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	SuccessCount *string `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 19
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePageFaceVerifyDataResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePageFaceVerifyDataResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) GetDate() *string {
	return s.Date
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) GetSceneName() *string {
	return s.SceneName
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) GetSuccessCount() *string {
	return s.SuccessCount
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetDate(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.Date = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetProductCode(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetSceneId(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.SceneId = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetSceneName(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.SceneName = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetSuccessCount(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.SuccessCount = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) SetTotalCount(v string) *DescribePageFaceVerifyDataResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribePageFaceVerifyDataResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
