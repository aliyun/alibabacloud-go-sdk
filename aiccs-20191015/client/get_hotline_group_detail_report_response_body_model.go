// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineGroupDetailReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineGroupDetailReportResponseBody
	GetCode() *string
	SetData(v *GetHotlineGroupDetailReportResponseBodyData) *GetHotlineGroupDetailReportResponseBody
	GetData() *GetHotlineGroupDetailReportResponseBodyData
	SetMessage(v string) *GetHotlineGroupDetailReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineGroupDetailReportResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetHotlineGroupDetailReportResponseBody
	GetSuccess() *string
}

type GetHotlineGroupDetailReportResponseBody struct {
	// Status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Query result data.
	Data *GetHotlineGroupDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API was invoked successfully. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// Public
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineGroupDetailReportResponseBody) GetData() *GetHotlineGroupDetailReportResponseBodyData {
	return s.Data
}

func (s *GetHotlineGroupDetailReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineGroupDetailReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineGroupDetailReportResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetHotlineGroupDetailReportResponseBody) SetCode(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetData(v *GetHotlineGroupDetailReportResponseBodyData) *GetHotlineGroupDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetMessage(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetRequestId(v string) *GetHotlineGroupDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetSuccess(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotlineGroupDetailReportResponseBodyData struct {
	// Description of returned columns.
	Columns []*GetHotlineGroupDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// Current page.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Returned data results.
	Rows []map[string]interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// Total number of records.
	//
	// example:
	//
	// 9
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyData) GetColumns() []*GetHotlineGroupDetailReportResponseBodyDataColumns {
	return s.Columns
}

func (s *GetHotlineGroupDetailReportResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *GetHotlineGroupDetailReportResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetHotlineGroupDetailReportResponseBodyData) GetRows() []map[string]interface{} {
	return s.Rows
}

func (s *GetHotlineGroupDetailReportResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetColumns(v []*GetHotlineGroupDetailReportResponseBodyDataColumns) *GetHotlineGroupDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPage(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineGroupDetailReportResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) Validate() error {
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

type GetHotlineGroupDetailReportResponseBodyDataColumns struct {
	// Metric.
	//
	// example:
	//
	// skillGroupName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Metric description.
	//
	// example:
	//
	// 技能组名称
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) GetKey() *string {
	return s.Key
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) GetTitle() *string {
	return s.Title
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) Validate() error {
	return dara.Validate(s)
}
