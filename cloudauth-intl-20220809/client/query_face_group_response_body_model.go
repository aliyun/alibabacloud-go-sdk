// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFaceGroupResponseBody
	GetCode() *string
	SetCurrentPage(v int64) *QueryFaceGroupResponseBody
	GetCurrentPage() *int64
	SetItems(v []*QueryFaceGroupResponseBodyItems) *QueryFaceGroupResponseBody
	GetItems() []*QueryFaceGroupResponseBodyItems
	SetMaxResults(v int32) *QueryFaceGroupResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *QueryFaceGroupResponseBody
	GetMessage() *string
	SetNextToken(v string) *QueryFaceGroupResponseBody
	GetNextToken() *string
	SetPageSize(v int32) *QueryFaceGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryFaceGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryFaceGroupResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *QueryFaceGroupResponseBody
	GetTotalPage() *int32
}

type QueryFaceGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int64                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*QueryFaceGroupResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// N/zmQeG/x9TDWmaB/pbfBQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5E63B760-0ECB-5C07-8503-A65C27876968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryFaceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFaceGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFaceGroupResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryFaceGroupResponseBody) GetItems() []*QueryFaceGroupResponseBodyItems {
	return s.Items
}

func (s *QueryFaceGroupResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryFaceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryFaceGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryFaceGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryFaceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFaceGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryFaceGroupResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *QueryFaceGroupResponseBody) SetCode(v string) *QueryFaceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetCurrentPage(v int64) *QueryFaceGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetItems(v []*QueryFaceGroupResponseBodyItems) *QueryFaceGroupResponseBody {
	s.Items = v
	return s
}

func (s *QueryFaceGroupResponseBody) SetMaxResults(v int32) *QueryFaceGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetMessage(v string) *QueryFaceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetNextToken(v string) *QueryFaceGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetPageSize(v int32) *QueryFaceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetRequestId(v string) *QueryFaceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetTotalCount(v int32) *QueryFaceGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryFaceGroupResponseBody) SetTotalPage(v int32) *QueryFaceGroupResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QueryFaceGroupResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryFaceGroupResponseBodyItems struct {
	// example:
	//
	// faceGroup001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 162261
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test-888
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryFaceGroupResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceGroupResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryFaceGroupResponseBodyItems) GetCode() *string {
	return s.Code
}

func (s *QueryFaceGroupResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *QueryFaceGroupResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *QueryFaceGroupResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *QueryFaceGroupResponseBodyItems) SetCode(v string) *QueryFaceGroupResponseBodyItems {
	s.Code = &v
	return s
}

func (s *QueryFaceGroupResponseBodyItems) SetDescription(v string) *QueryFaceGroupResponseBodyItems {
	s.Description = &v
	return s
}

func (s *QueryFaceGroupResponseBodyItems) SetId(v int64) *QueryFaceGroupResponseBodyItems {
	s.Id = &v
	return s
}

func (s *QueryFaceGroupResponseBodyItems) SetName(v string) *QueryFaceGroupResponseBodyItems {
	s.Name = &v
	return s
}

func (s *QueryFaceGroupResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
