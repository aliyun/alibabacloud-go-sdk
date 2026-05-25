// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogDeviceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *GetTlogDeviceListResponseBody
	GetErrorCode() *int64
	SetMessage(v string) *GetTlogDeviceListResponseBody
	GetMessage() *string
	SetModel(v *GetTlogDeviceListResponseBodyModel) *GetTlogDeviceListResponseBody
	GetModel() *GetTlogDeviceListResponseBodyModel
	SetRequestId(v string) *GetTlogDeviceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTlogDeviceListResponseBody
	GetSuccess() *bool
}

type GetTlogDeviceListResponseBody struct {
	// example:
	//
	// Success
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// internal error
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetTlogDeviceListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// AB8AB5EC-9636-421D-AE7C-BB227DFC95B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTlogDeviceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceListResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *GetTlogDeviceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTlogDeviceListResponseBody) GetModel() *GetTlogDeviceListResponseBodyModel {
	return s.Model
}

func (s *GetTlogDeviceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTlogDeviceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTlogDeviceListResponseBody) SetErrorCode(v int64) *GetTlogDeviceListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTlogDeviceListResponseBody) SetMessage(v string) *GetTlogDeviceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTlogDeviceListResponseBody) SetModel(v *GetTlogDeviceListResponseBodyModel) *GetTlogDeviceListResponseBody {
	s.Model = v
	return s
}

func (s *GetTlogDeviceListResponseBody) SetRequestId(v string) *GetTlogDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTlogDeviceListResponseBody) SetSuccess(v bool) *GetTlogDeviceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTlogDeviceListResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTlogDeviceListResponseBodyModel struct {
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
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTlogDeviceListResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceListResponseBodyModel) GetItems() []interface{} {
	return s.Items
}

func (s *GetTlogDeviceListResponseBodyModel) GetPageNum() *int64 {
	return s.PageNum
}

func (s *GetTlogDeviceListResponseBodyModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetTlogDeviceListResponseBodyModel) GetPages() *int64 {
	return s.Pages
}

func (s *GetTlogDeviceListResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *GetTlogDeviceListResponseBodyModel) SetItems(v []interface{}) *GetTlogDeviceListResponseBodyModel {
	s.Items = v
	return s
}

func (s *GetTlogDeviceListResponseBodyModel) SetPageNum(v int64) *GetTlogDeviceListResponseBodyModel {
	s.PageNum = &v
	return s
}

func (s *GetTlogDeviceListResponseBodyModel) SetPageSize(v int64) *GetTlogDeviceListResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *GetTlogDeviceListResponseBodyModel) SetPages(v int64) *GetTlogDeviceListResponseBodyModel {
	s.Pages = &v
	return s
}

func (s *GetTlogDeviceListResponseBodyModel) SetTotal(v int64) *GetTlogDeviceListResponseBodyModel {
	s.Total = &v
	return s
}

func (s *GetTlogDeviceListResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
