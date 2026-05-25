// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogCollectListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *GetTlogCollectListResponseBody
	GetErrorCode() *int64
	SetMessage(v string) *GetTlogCollectListResponseBody
	GetMessage() *string
	SetModel(v *GetTlogCollectListResponseBodyModel) *GetTlogCollectListResponseBody
	GetModel() *GetTlogCollectListResponseBodyModel
	SetRequestId(v string) *GetTlogCollectListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTlogCollectListResponseBody
	GetSuccess() *bool
}

type GetTlogCollectListResponseBody struct {
	// example:
	//
	// 500
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetTlogCollectListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// AB8AB5EC-9636-421D-AE7C-BB227DFC95B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTlogCollectListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTlogCollectListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTlogCollectListResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *GetTlogCollectListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTlogCollectListResponseBody) GetModel() *GetTlogCollectListResponseBodyModel {
	return s.Model
}

func (s *GetTlogCollectListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTlogCollectListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTlogCollectListResponseBody) SetErrorCode(v int64) *GetTlogCollectListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTlogCollectListResponseBody) SetMessage(v string) *GetTlogCollectListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTlogCollectListResponseBody) SetModel(v *GetTlogCollectListResponseBodyModel) *GetTlogCollectListResponseBody {
	s.Model = v
	return s
}

func (s *GetTlogCollectListResponseBody) SetRequestId(v string) *GetTlogCollectListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTlogCollectListResponseBody) SetSuccess(v bool) *GetTlogCollectListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTlogCollectListResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTlogCollectListResponseBodyModel struct {
	Items []interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTlogCollectListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetTlogCollectListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetTlogCollectListResponseBodyModel) GetItems() []interface{} {
	return s.Items
}

func (s *GetTlogCollectListResponseBodyModel) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetTlogCollectListResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTlogCollectListResponseBodyModel) GetPages() *int64 {
	return s.Pages
}

func (s *GetTlogCollectListResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *GetTlogCollectListResponseBodyModel) SetItems(v []interface{}) *GetTlogCollectListResponseBodyModel {
	s.Items = v
	return s
}

func (s *GetTlogCollectListResponseBodyModel) SetPageNum(v int64) *GetTlogCollectListResponseBodyModel {
	s.PageNum = &v
	return s
}

func (s *GetTlogCollectListResponseBodyModel) SetPageSize(v int64) *GetTlogCollectListResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *GetTlogCollectListResponseBodyModel) SetPages(v int64) *GetTlogCollectListResponseBodyModel {
	s.Pages = &v
	return s
}

func (s *GetTlogCollectListResponseBodyModel) SetTotal(v int64) *GetTlogCollectListResponseBodyModel {
	s.Total = &v
	return s
}

func (s *GetTlogCollectListResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
