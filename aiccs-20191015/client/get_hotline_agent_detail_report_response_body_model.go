// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentDetailReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineAgentDetailReportResponseBody
	GetCode() *string
	SetData(v *GetHotlineAgentDetailReportResponseBodyData) *GetHotlineAgentDetailReportResponseBody
	GetData() *GetHotlineAgentDetailReportResponseBodyData
	SetHttpStatusCode(v int64) *GetHotlineAgentDetailReportResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetHotlineAgentDetailReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineAgentDetailReportResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetHotlineAgentDetailReportResponseBody
	GetSuccess() *string
}

type GetHotlineAgentDetailReportResponseBody struct {
	// example:
	//
	// Success
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetHotlineAgentDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineAgentDetailReportResponseBody) GetData() *GetHotlineAgentDetailReportResponseBodyData {
	return s.Data
}

func (s *GetHotlineAgentDetailReportResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetHotlineAgentDetailReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineAgentDetailReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineAgentDetailReportResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetHotlineAgentDetailReportResponseBody) SetCode(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetData(v *GetHotlineAgentDetailReportResponseBodyData) *GetHotlineAgentDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentDetailReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetMessage(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetRequestId(v string) *GetHotlineAgentDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetSuccess(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotlineAgentDetailReportResponseBodyData struct {
	Columns []*GetHotlineAgentDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rows     []map[string]interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// example:
	//
	// 7
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyData) GetColumns() []*GetHotlineAgentDetailReportResponseBodyDataColumns {
	return s.Columns
}

func (s *GetHotlineAgentDetailReportResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *GetHotlineAgentDetailReportResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHotlineAgentDetailReportResponseBodyData) GetRows() []map[string]interface{} {
	return s.Rows
}

func (s *GetHotlineAgentDetailReportResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetColumns(v []*GetHotlineAgentDetailReportResponseBodyDataColumns) *GetHotlineAgentDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPage(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineAgentDetailReportResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotlineAgentDetailReportResponseBodyDataColumns struct {
	// example:
	//
	// realName
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) GetKey() *string {
	return s.Key
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) GetTitle() *string {
	return s.Title
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) Validate() error {
	return dara.Validate(s)
}
