// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentIndexRealTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentIndexRealTimeResponseBody
	GetCode() *string
	SetData(v *GetAgentIndexRealTimeResponseBodyData) *GetAgentIndexRealTimeResponseBody
	GetData() *GetAgentIndexRealTimeResponseBodyData
	SetMessage(v string) *GetAgentIndexRealTimeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentIndexRealTimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentIndexRealTimeResponseBody
	GetSuccess() *bool
}

type GetAgentIndexRealTimeResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetAgentIndexRealTimeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentIndexRealTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIndexRealTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentIndexRealTimeResponseBody) GetData() *GetAgentIndexRealTimeResponseBodyData {
	return s.Data
}

func (s *GetAgentIndexRealTimeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentIndexRealTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentIndexRealTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentIndexRealTimeResponseBody) SetCode(v string) *GetAgentIndexRealTimeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetData(v *GetAgentIndexRealTimeResponseBodyData) *GetAgentIndexRealTimeResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetMessage(v string) *GetAgentIndexRealTimeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetRequestId(v string) *GetAgentIndexRealTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) SetSuccess(v bool) *GetAgentIndexRealTimeResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentIndexRealTimeResponseBodyData struct {
	// Description of returned columns.
	Columns []*GetAgentIndexRealTimeResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Returned data results.
	Rows []map[string]interface{} `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// Total number of records.
	//
	// example:
	//
	// 4
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAgentIndexRealTimeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIndexRealTimeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponseBodyData) GetColumns() []*GetAgentIndexRealTimeResponseBodyDataColumns {
	return s.Columns
}

func (s *GetAgentIndexRealTimeResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *GetAgentIndexRealTimeResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentIndexRealTimeResponseBodyData) GetRows() []map[string]interface{} {
	return s.Rows
}

func (s *GetAgentIndexRealTimeResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetColumns(v []*GetAgentIndexRealTimeResponseBodyDataColumns) *GetAgentIndexRealTimeResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetPage(v int32) *GetAgentIndexRealTimeResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetPageSize(v int32) *GetAgentIndexRealTimeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetRows(v []map[string]interface{}) *GetAgentIndexRealTimeResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) SetTotal(v int32) *GetAgentIndexRealTimeResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyData) Validate() error {
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

type GetAgentIndexRealTimeResponseBodyDataColumns struct {
	// Metric.
	//
	// example:
	//
	// 客服ID
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Metric description.
	//
	// example:
	//
	// servicerId
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetAgentIndexRealTimeResponseBodyDataColumns) String() string {
	return dara.Prettify(s)
}

func (s GetAgentIndexRealTimeResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetAgentIndexRealTimeResponseBodyDataColumns) GetKey() *string {
	return s.Key
}

func (s *GetAgentIndexRealTimeResponseBodyDataColumns) GetTitle() *string {
	return s.Title
}

func (s *GetAgentIndexRealTimeResponseBodyDataColumns) SetKey(v string) *GetAgentIndexRealTimeResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyDataColumns) SetTitle(v string) *GetAgentIndexRealTimeResponseBodyDataColumns {
	s.Title = &v
	return s
}

func (s *GetAgentIndexRealTimeResponseBodyDataColumns) Validate() error {
	return dara.Validate(s)
}
