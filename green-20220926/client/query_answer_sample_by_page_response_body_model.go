// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAnswerSampleByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryAnswerSampleByPageResponseBody
	GetCurrentPage() *int32
	SetItems(v []*QueryAnswerSampleByPageResponseBodyItems) *QueryAnswerSampleByPageResponseBody
	GetItems() []*QueryAnswerSampleByPageResponseBodyItems
	SetPageSize(v int32) *QueryAnswerSampleByPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryAnswerSampleByPageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryAnswerSampleByPageResponseBody
	GetTotalCount() *int64
}

type QueryAnswerSampleByPageResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                      `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*QueryAnswerSampleByPageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 58
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAnswerSampleByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAnswerSampleByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryAnswerSampleByPageResponseBody) GetItems() []*QueryAnswerSampleByPageResponseBodyItems {
	return s.Items
}

func (s *QueryAnswerSampleByPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAnswerSampleByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAnswerSampleByPageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryAnswerSampleByPageResponseBody) SetCurrentPage(v int32) *QueryAnswerSampleByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetItems(v []*QueryAnswerSampleByPageResponseBodyItems) *QueryAnswerSampleByPageResponseBody {
	s.Items = v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetPageSize(v int32) *QueryAnswerSampleByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetRequestId(v string) *QueryAnswerSampleByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) SetTotalCount(v int64) *QueryAnswerSampleByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAnswerSampleByPageResponseBodyItems struct {
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 2023-07-31 06:16:06
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1666
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// custom_xxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// UID。
	//
	// example:
	//
	// 104813*****2399
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryAnswerSampleByPageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s QueryAnswerSampleByPageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryAnswerSampleByPageResponseBodyItems) GetAnswer() *string {
	return s.Answer
}

func (s *QueryAnswerSampleByPageResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryAnswerSampleByPageResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *QueryAnswerSampleByPageResponseBodyItems) GetLibId() *string {
	return s.LibId
}

func (s *QueryAnswerSampleByPageResponseBodyItems) GetUid() *string {
	return s.Uid
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetAnswer(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.Answer = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetGmtCreate(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetId(v int64) *QueryAnswerSampleByPageResponseBodyItems {
	s.Id = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetLibId(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.LibId = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) SetUid(v string) *QueryAnswerSampleByPageResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *QueryAnswerSampleByPageResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
