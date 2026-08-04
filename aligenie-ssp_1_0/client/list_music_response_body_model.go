// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMusicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListMusicResponseBody
	GetCode() *int32
	SetMessage(v string) *ListMusicResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMusicResponseBody
	GetRequestId() *string
	SetResult(v *ListMusicResponseBodyResult) *ListMusicResponseBody
	GetResult() *ListMusicResponseBodyResult
}

type ListMusicResponseBody struct {
	// Status code returned by the alarm service
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// error message
	//
	// example:
	//
	// 音乐类型名称为空
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 43***28C-A810-5***-8747-EC226A086881
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// collection of ringtone query results
	Result *ListMusicResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListMusicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMusicResponseBody) GoString() string {
	return s.String()
}

func (s *ListMusicResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListMusicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMusicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMusicResponseBody) GetResult() *ListMusicResponseBodyResult {
	return s.Result
}

func (s *ListMusicResponseBody) SetCode(v int32) *ListMusicResponseBody {
	s.Code = &v
	return s
}

func (s *ListMusicResponseBody) SetMessage(v string) *ListMusicResponseBody {
	s.Message = &v
	return s
}

func (s *ListMusicResponseBody) SetRequestId(v string) *ListMusicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMusicResponseBody) SetResult(v *ListMusicResponseBodyResult) *ListMusicResponseBody {
	s.Result = v
	return s
}

func (s *ListMusicResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMusicResponseBodyResult struct {
	// Current page
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of ringtones
	Model []*ListMusicResponseBodyResultModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// Total number of pages
	//
	// example:
	//
	// 10
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// Number of entries per page: maximum value is 100; values exceeding 100 are treated as 100
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of entries
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMusicResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListMusicResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListMusicResponseBodyResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMusicResponseBodyResult) GetModel() []*ListMusicResponseBodyResultModel {
	return s.Model
}

func (s *ListMusicResponseBodyResult) GetPageCount() *int32 {
	return s.PageCount
}

func (s *ListMusicResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMusicResponseBodyResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMusicResponseBodyResult) SetCurrentPage(v int32) *ListMusicResponseBodyResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMusicResponseBodyResult) SetModel(v []*ListMusicResponseBodyResultModel) *ListMusicResponseBodyResult {
	s.Model = v
	return s
}

func (s *ListMusicResponseBodyResult) SetPageCount(v int32) *ListMusicResponseBodyResult {
	s.PageCount = &v
	return s
}

func (s *ListMusicResponseBodyResult) SetPageSize(v int32) *ListMusicResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *ListMusicResponseBodyResult) SetTotalCount(v int32) *ListMusicResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListMusicResponseBodyResult) Validate() error {
	if s.Model != nil {
		for _, item := range s.Model {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMusicResponseBodyResultModel struct {
	// Ringtone ID
	//
	// example:
	//
	// 1
	MusicId *int64 `json:"MusicId,omitempty" xml:"MusicId,omitempty"`
	// Ringtone name
	//
	// example:
	//
	// xx
	MusicName *string `json:"MusicName,omitempty" xml:"MusicName,omitempty"`
	// Ringtone category ID
	//
	// example:
	//
	// 1
	MusicType *int64 `json:"MusicType,omitempty" xml:"MusicType,omitempty"`
	// Ringtone category name
	//
	// example:
	//
	// xx
	MusicTypeName *string `json:"MusicTypeName,omitempty" xml:"MusicTypeName,omitempty"`
	// Ringtone URL
	//
	// example:
	//
	// http://xx
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
}

func (s ListMusicResponseBodyResultModel) String() string {
	return dara.Prettify(s)
}

func (s ListMusicResponseBodyResultModel) GoString() string {
	return s.String()
}

func (s *ListMusicResponseBodyResultModel) GetMusicId() *int64 {
	return s.MusicId
}

func (s *ListMusicResponseBodyResultModel) GetMusicName() *string {
	return s.MusicName
}

func (s *ListMusicResponseBodyResultModel) GetMusicType() *int64 {
	return s.MusicType
}

func (s *ListMusicResponseBodyResultModel) GetMusicTypeName() *string {
	return s.MusicTypeName
}

func (s *ListMusicResponseBodyResultModel) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *ListMusicResponseBodyResultModel) SetMusicId(v int64) *ListMusicResponseBodyResultModel {
	s.MusicId = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicName(v string) *ListMusicResponseBodyResultModel {
	s.MusicName = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicType(v int64) *ListMusicResponseBodyResultModel {
	s.MusicType = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicTypeName(v string) *ListMusicResponseBodyResultModel {
	s.MusicTypeName = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) SetMusicUrl(v string) *ListMusicResponseBodyResultModel {
	s.MusicUrl = &v
	return s
}

func (s *ListMusicResponseBodyResultModel) Validate() error {
	return dara.Validate(s)
}
