// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineServiceVolumeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOnlineServiceVolumeResponseBody
	GetCode() *string
	SetData(v *GetOnlineServiceVolumeResponseBodyData) *GetOnlineServiceVolumeResponseBody
	GetData() *GetOnlineServiceVolumeResponseBodyData
	SetMessage(v string) *GetOnlineServiceVolumeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOnlineServiceVolumeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetOnlineServiceVolumeResponseBody
	GetSuccess() *string
}

type GetOnlineServiceVolumeResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetOnlineServiceVolumeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetOnlineServiceVolumeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineServiceVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOnlineServiceVolumeResponseBody) GetData() *GetOnlineServiceVolumeResponseBodyData {
	return s.Data
}

func (s *GetOnlineServiceVolumeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOnlineServiceVolumeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOnlineServiceVolumeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetOnlineServiceVolumeResponseBody) SetCode(v string) *GetOnlineServiceVolumeResponseBody {
	s.Code = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetData(v *GetOnlineServiceVolumeResponseBodyData) *GetOnlineServiceVolumeResponseBody {
	s.Data = v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetMessage(v string) *GetOnlineServiceVolumeResponseBody {
	s.Message = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetRequestId(v string) *GetOnlineServiceVolumeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) SetSuccess(v string) *GetOnlineServiceVolumeResponseBody {
	s.Success = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOnlineServiceVolumeResponseBodyData struct {
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
	// A JSON string of type List<Map>.
	//
	// example:
	//
	// {"tenant_id":"905","online_40s_transfer_ready_cnt":109,"minute_id":"-1","wait_time_len":1215,"pickup_rate":"63.09%","thirty_seconds_to_pickUp":"2560","date_id":"-1","online_over_out_cnt":0,"online_20s_transfer_ready_cnt":109,"thirty_seconds_response_rate":"63.09%","abandonment_rate":"63.09%","service_time_len":68378,"service_pickup":"2560","hour_id":"-1","online_10s_transfer_ready_cnt":109}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetOnlineServiceVolumeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineServiceVolumeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOnlineServiceVolumeResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetOnlineServiceVolumeResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOnlineServiceVolumeResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetOnlineServiceVolumeResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetPageNum(v int32) *GetOnlineServiceVolumeResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetPageSize(v int32) *GetOnlineServiceVolumeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetRows(v string) *GetOnlineServiceVolumeResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBodyData) SetTotalNum(v int32) *GetOnlineServiceVolumeResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetOnlineServiceVolumeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
