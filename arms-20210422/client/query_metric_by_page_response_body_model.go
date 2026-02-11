// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMetricByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMetricByPageResponseBody
	GetCode() *string
	SetData(v *QueryMetricByPageResponseBodyData) *QueryMetricByPageResponseBody
	GetData() *QueryMetricByPageResponseBodyData
	SetMessage(v string) *QueryMetricByPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMetricByPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMetricByPageResponseBody
	GetSuccess() *bool
}

type QueryMetricByPageResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryMetricByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMetricByPageResponseBody) GetData() *QueryMetricByPageResponseBodyData {
	return s.Data
}

func (s *QueryMetricByPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMetricByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMetricByPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMetricByPageResponseBody) SetCode(v string) *QueryMetricByPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetData(v *QueryMetricByPageResponseBodyData) *QueryMetricByPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryMetricByPageResponseBody) SetMessage(v string) *QueryMetricByPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetRequestId(v string) *QueryMetricByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetSuccess(v bool) *QueryMetricByPageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMetricByPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMetricByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMetricByPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *QueryMetricByPageResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *QueryMetricByPageResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMetricByPageResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryMetricByPageResponseBodyData) SetItems(v []map[string]interface{}) *QueryMetricByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPage(v int32) *QueryMetricByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPageSize(v int32) *QueryMetricByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetTotal(v int32) *QueryMetricByPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
