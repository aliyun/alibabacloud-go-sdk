// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppApiByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAppApiByPageResponseBody
	GetCode() *int32
	SetData(v *GetAppApiByPageResponseBodyData) *GetAppApiByPageResponseBody
	GetData() *GetAppApiByPageResponseBodyData
	SetMessage(v string) *GetAppApiByPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppApiByPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppApiByPageResponseBody
	GetSuccess() *bool
}

type GetAppApiByPageResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAppApiByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppApiByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppApiByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAppApiByPageResponseBody) GetData() *GetAppApiByPageResponseBodyData {
	return s.Data
}

func (s *GetAppApiByPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppApiByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppApiByPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppApiByPageResponseBody) SetCode(v int32) *GetAppApiByPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetData(v *GetAppApiByPageResponseBodyData) *GetAppApiByPageResponseBody {
	s.Data = v
	return s
}

func (s *GetAppApiByPageResponseBody) SetMessage(v string) *GetAppApiByPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetRequestId(v string) *GetAppApiByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetSuccess(v bool) *GetAppApiByPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppApiByPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppApiByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *string                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAppApiByPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAppApiByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *GetAppApiByPageResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *GetAppApiByPageResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAppApiByPageResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *GetAppApiByPageResponseBodyData) SetItems(v []map[string]interface{}) *GetAppApiByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPage(v int32) *GetAppApiByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPageSize(v int32) *GetAppApiByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetTotal(v string) *GetAppApiByPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
