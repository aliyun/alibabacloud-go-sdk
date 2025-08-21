// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumDetailByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAlbumDetailByIdResponseBody
	GetCode() *int32
	SetMessage(v string) *GetAlbumDetailByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlbumDetailByIdResponseBody
	GetRequestId() *string
	SetResult(v *GetAlbumDetailByIdResponseBodyResult) *GetAlbumDetailByIdResponseBody
	GetResult() *GetAlbumDetailByIdResponseBodyResult
}

type GetAlbumDetailByIdResponseBody struct {
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
	// A0B7CACD-485B-14E2-854F-39EACB09E45B
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetAlbumDetailByIdResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAlbumDetailByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAlbumDetailByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlbumDetailByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlbumDetailByIdResponseBody) GetResult() *GetAlbumDetailByIdResponseBodyResult {
	return s.Result
}

func (s *GetAlbumDetailByIdResponseBody) SetCode(v int32) *GetAlbumDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBody) SetMessage(v string) *GetAlbumDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBody) SetRequestId(v string) *GetAlbumDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBody) SetResult(v *GetAlbumDetailByIdResponseBodyResult) *GetAlbumDetailByIdResponseBody {
	s.Result = v
	return s
}

func (s *GetAlbumDetailByIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAlbumDetailByIdResponseBodyResult struct {
	AlbumContentList []*GetAlbumDetailByIdResponseBodyResultAlbumContentList `json:"AlbumContentList,omitempty" xml:"AlbumContentList,omitempty" type:"Repeated"`
	// example:
	//
	// https://ailabs.alibabausercontent.com/images/8838/1600839452498.jpg
	AlbumCoverUrl *string `json:"AlbumCoverUrl,omitempty" xml:"AlbumCoverUrl,omitempty"`
	// example:
	//
	// 每次一个百科知识或者故事\n丰富孩子的视野，拓展眼界和知识面，培养和孩子的探究能力和好奇心\n\n
	AlbumDescription *string `json:"AlbumDescription,omitempty" xml:"AlbumDescription,omitempty"`
	// example:
	//
	// 51999575
	AlbumId *string `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	// example:
	//
	// 小科学家探索
	AlbumTitle *string `json:"AlbumTitle,omitempty" xml:"AlbumTitle,omitempty"`
}

func (s GetAlbumDetailByIdResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumDetailByIdResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponseBodyResult) GetAlbumContentList() []*GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	return s.AlbumContentList
}

func (s *GetAlbumDetailByIdResponseBodyResult) GetAlbumCoverUrl() *string {
	return s.AlbumCoverUrl
}

func (s *GetAlbumDetailByIdResponseBodyResult) GetAlbumDescription() *string {
	return s.AlbumDescription
}

func (s *GetAlbumDetailByIdResponseBodyResult) GetAlbumId() *string {
	return s.AlbumId
}

func (s *GetAlbumDetailByIdResponseBodyResult) GetAlbumTitle() *string {
	return s.AlbumTitle
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumContentList(v []*GetAlbumDetailByIdResponseBodyResultAlbumContentList) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumContentList = v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumCoverUrl(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumCoverUrl = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumDescription(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumDescription = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumId(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumId = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) SetAlbumTitle(v string) *GetAlbumDetailByIdResponseBodyResult {
	s.AlbumTitle = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetAlbumDetailByIdResponseBodyResultAlbumContentList struct {
	// example:
	//
	// 3分24秒
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 468009044
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1000
	OrderIndex *string `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// example:
	//
	// 001为什么肚子饿时会咕咕叫
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetAlbumDetailByIdResponseBodyResultAlbumContentList) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumDetailByIdResponseBodyResultAlbumContentList) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) GetDuration() *string {
	return s.Duration
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) GetId() *string {
	return s.Id
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) GetOrderIndex() *string {
	return s.OrderIndex
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) GetTitle() *string {
	return s.Title
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetDuration(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.Duration = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetId(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.Id = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetOrderIndex(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.OrderIndex = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) SetTitle(v string) *GetAlbumDetailByIdResponseBodyResultAlbumContentList {
	s.Title = &v
	return s
}

func (s *GetAlbumDetailByIdResponseBodyResultAlbumContentList) Validate() error {
	return dara.Validate(s)
}
