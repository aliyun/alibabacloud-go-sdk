// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineSeatInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOnlineSeatInformationResponseBody
	GetCode() *string
	SetData(v *GetOnlineSeatInformationResponseBodyData) *GetOnlineSeatInformationResponseBody
	GetData() *GetOnlineSeatInformationResponseBodyData
	SetMessage(v string) *GetOnlineSeatInformationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOnlineSeatInformationResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetOnlineSeatInformationResponseBody
	GetSuccess() *string
}

type GetOnlineSeatInformationResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetOnlineSeatInformationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Status code description.
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOnlineSeatInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineSeatInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOnlineSeatInformationResponseBody) GetData() *GetOnlineSeatInformationResponseBodyData {
	return s.Data
}

func (s *GetOnlineSeatInformationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOnlineSeatInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOnlineSeatInformationResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetOnlineSeatInformationResponseBody) SetCode(v string) *GetOnlineSeatInformationResponseBody {
	s.Code = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetData(v *GetOnlineSeatInformationResponseBodyData) *GetOnlineSeatInformationResponseBody {
	s.Data = v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetMessage(v string) *GetOnlineSeatInformationResponseBody {
	s.Message = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetRequestId(v string) *GetOnlineSeatInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) SetSuccess(v string) *GetOnlineSeatInformationResponseBody {
	s.Success = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOnlineSeatInformationResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size.
	//
	// example:
	//
	// 2000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information is a JSON string of the List<Map> type.
	//
	// example:
	//
	// {"tenant_id":"905","max_service_num":"3","servicer_id":"36****","tenant_name":"非单元测试化BU","record_gmt_modified":"2021-04-01 23:20:03","pk_id":"3690012021****","record_gmt_create":"2021-04-01 23:19:55","department_id":"10****"}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOnlineSeatInformationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineSeatInformationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetOnlineSeatInformationResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOnlineSeatInformationResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetOnlineSeatInformationResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetOnlineSeatInformationResponseBodyData) SetPageNum(v int32) *GetOnlineSeatInformationResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBodyData) SetPageSize(v int32) *GetOnlineSeatInformationResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBodyData) SetRows(v string) *GetOnlineSeatInformationResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBodyData) SetTotalNum(v int32) *GetOnlineSeatInformationResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetOnlineSeatInformationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
