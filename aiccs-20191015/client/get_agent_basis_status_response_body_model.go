// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentBasisStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentBasisStatusResponseBody
	GetCode() *string
	SetData(v *GetAgentBasisStatusResponseBodyData) *GetAgentBasisStatusResponseBody
	GetData() *GetAgentBasisStatusResponseBodyData
	SetMessage(v string) *GetAgentBasisStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentBasisStatusResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAgentBasisStatusResponseBody
	GetSuccess() *string
}

type GetAgentBasisStatusResponseBody struct {
	// The status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data list.
	Data *GetAgentBasisStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Status code description.
	//
	// example:
	//
	// OK
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
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentBasisStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentBasisStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentBasisStatusResponseBody) GetData() *GetAgentBasisStatusResponseBodyData {
	return s.Data
}

func (s *GetAgentBasisStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentBasisStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentBasisStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAgentBasisStatusResponseBody) SetCode(v string) *GetAgentBasisStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetData(v *GetAgentBasisStatusResponseBodyData) *GetAgentBasisStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetMessage(v string) *GetAgentBasisStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetRequestId(v string) *GetAgentBasisStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) SetSuccess(v string) *GetAgentBasisStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentBasisStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentBasisStatusResponseBodyData struct {
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
	// {"tenant_id":"905","servicer_id":"73****","tenant_name":"测试","recordgmtmodified":"2021-04-01 11:36:50","pk_id":"7320372021****","statusstarttime":"2021-04-01 11:36:50","recordgmtcreate":"2021-04-01 10:06:24","department_id":"94****","department_name":"测试技能组","lstlogintime":"2021-04-01 10:06:24","date_id":"20210401","triggerreason":"3","servicer_status":"D","assignstatus":"1","servicerreal_name":"xx","servicerstatusname":"签出","fstlogintime":"2021-04-01 10:06:24","servicer_name":"xx"}
	Rows *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 4
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetAgentBasisStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentBasisStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentBasisStatusResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetAgentBasisStatusResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAgentBasisStatusResponseBodyData) GetRows() *string {
	return s.Rows
}

func (s *GetAgentBasisStatusResponseBodyData) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetAgentBasisStatusResponseBodyData) SetPageNum(v int32) *GetAgentBasisStatusResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetAgentBasisStatusResponseBodyData) SetPageSize(v int32) *GetAgentBasisStatusResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAgentBasisStatusResponseBodyData) SetRows(v string) *GetAgentBasisStatusResponseBodyData {
	s.Rows = &v
	return s
}

func (s *GetAgentBasisStatusResponseBodyData) SetTotalNum(v int32) *GetAgentBasisStatusResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetAgentBasisStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
