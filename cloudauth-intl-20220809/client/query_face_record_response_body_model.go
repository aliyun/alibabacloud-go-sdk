// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceRecordResponseBody
	GetCode() *string
	SetCurrentPage(v int64) *QueryFaceRecordResponseBody
	GetCurrentPage() *int64
	SetItems(v []*QueryFaceRecordResponseBodyItems) *QueryFaceRecordResponseBody
	GetItems() []*QueryFaceRecordResponseBodyItems
	SetMaxResults(v int32) *QueryFaceRecordResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *QueryFaceRecordResponseBody
	GetMessage() *string
	SetNextToken(v string) *QueryFaceRecordResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *QueryFaceRecordResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryFaceRecordResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryFaceRecordResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *QueryFaceRecordResponseBody
	GetTotalPage() *int32
}

type QueryFaceRecordResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2
	CurrentPage *int64                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*QueryFaceRecordResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAARfZmVDe9NvRXloR5+8CK9nwqHyx44CQz3pa71+mmu0e
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7F971622-38C0-5F56-B2EC-315367979B4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryFaceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceRecordResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceRecordResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryFaceRecordResponseBody) GetItems() []*QueryFaceRecordResponseBodyItems {
	return s.Items
}

func (s *QueryFaceRecordResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryFaceRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryFaceRecordResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryFaceRecordResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceRecordResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryFaceRecordResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *QueryFaceRecordResponseBody) SetCode(v string) *QueryFaceRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetCurrentPage(v int64) *QueryFaceRecordResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetItems(v []*QueryFaceRecordResponseBodyItems) *QueryFaceRecordResponseBody {
	s.Items = v
	return s
}

func (s *QueryFaceRecordResponseBody) SetMaxResults(v int32) *QueryFaceRecordResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetMessage(v string) *QueryFaceRecordResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetNextToken(v string) *QueryFaceRecordResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetPageSize(v int32) *QueryFaceRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetRequestId(v string) *QueryFaceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetTotalCount(v int32) *QueryFaceRecordResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryFaceRecordResponseBody) SetTotalPage(v int32) *QueryFaceRecordResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QueryFaceRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryFaceRecordResponseBodyItems struct {
	// example:
	//
	// 230642938
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// example:
	//
	// 2025-01-15T02:20:28Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 16112
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// http://www.xxxx.com/1.jpg
	ImgOssUrl *string `json:"ImgOssUrl,omitempty" xml:"ImgOssUrl,omitempty"`
	// example:
	//
	// CuN6hiD08qr
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// MANUAL
	RegistrationType *string `json:"RegistrationType,omitempty" xml:"RegistrationType,omitempty"`
}

func (s QueryFaceRecordResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceRecordResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryFaceRecordResponseBodyItems) GetFaceId() *string {
	return s.FaceId
}

func (s *QueryFaceRecordResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryFaceRecordResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *QueryFaceRecordResponseBodyItems) GetImgOssUrl() *string {
	return s.ImgOssUrl
}

func (s *QueryFaceRecordResponseBodyItems) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *QueryFaceRecordResponseBodyItems) GetRegistrationType() *string {
	return s.RegistrationType
}

func (s *QueryFaceRecordResponseBodyItems) SetFaceId(v string) *QueryFaceRecordResponseBodyItems {
	s.FaceId = &v
	return s
}

func (s *QueryFaceRecordResponseBodyItems) SetGmtCreate(v string) *QueryFaceRecordResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *QueryFaceRecordResponseBodyItems) SetId(v int64) *QueryFaceRecordResponseBodyItems {
	s.Id = &v
	return s
}

func (s *QueryFaceRecordResponseBodyItems) SetImgOssUrl(v string) *QueryFaceRecordResponseBodyItems {
	s.ImgOssUrl = &v
	return s
}

func (s *QueryFaceRecordResponseBodyItems) SetMerchantUserId(v string) *QueryFaceRecordResponseBodyItems {
	s.MerchantUserId = &v
	return s
}

func (s *QueryFaceRecordResponseBodyItems) SetRegistrationType(v string) *QueryFaceRecordResponseBodyItems {
	s.RegistrationType = &v
	return s
}

func (s *QueryFaceRecordResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
