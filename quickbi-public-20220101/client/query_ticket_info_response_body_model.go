// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTicketInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryTicketInfoResponseBody
	GetRequestId() *string
	SetResult(v *QueryTicketInfoResponseBodyResult) *QueryTicketInfoResponseBody
	GetResult() *QueryTicketInfoResponseBodyResult
	SetSuccess(v bool) *QueryTicketInfoResponseBody
	GetSuccess() *bool
}

type QueryTicketInfoResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D787E1A3-A93C-424A-B626-C2B05DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the generated ticket.
	Result *QueryTicketInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTicketInfoResponseBody) GetResult() *QueryTicketInfoResponseBodyResult {
	return s.Result
}

func (s *QueryTicketInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTicketInfoResponseBody) SetRequestId(v string) *QueryTicketInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketInfoResponseBody) SetResult(v *QueryTicketInfoResponseBodyResult) *QueryTicketInfoResponseBody {
	s.Result = v
	return s
}

func (s *QueryTicketInfoResponseBody) SetSuccess(v bool) *QueryTicketInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTicketInfoResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTicketInfoResponseBodyResult struct {
	// Notes.
	//
	// example:
	//
	// a27a9aec-****-****-bd40-1a21ea41d7c5
	AccessTicket *string `json:"AccessTicket,omitempty" xml:"AccessTicket,omitempty"`
	// The ID of the component in the report.
	//
	// example:
	//
	// sfdgsds-****-****-a608-mghdgd
	CmptId *string `json:"CmptId,omitempty" xml:"CmptId,omitempty"`
	// Global parameters.
	//
	// example:
	//
	// [&{quot;paramKey\\&quot;:\\&quot;price\\&quot;,\\&quot;joinType\\&quot;and\\&quot;,\\&quot;conditionList\\&quot;:[{\\&quot; operation\\&quot;\\&quot;\\&quot;\\&quot;\\&quot;\\&quot;\\&quot;value ;& quot;\\&quot;\\&quot;\\&quot;\\&quot;\\&quot;\\&quot;\\&quot;\\&quot;\\&quot product_type\\&quot;,\\&quot;joinType\\&quot;:\\&quot;and ";,& quot;conditionList\\&quot;, the conditions must be:[{\\&quot;operate" ;:& quot;in\\&quot;,\\&quot;value\\&quot;, the conditions must be:[\\&quot; office supplies\\&quot;,\\&quot; furniture products\\&quot;]}]}]\\n
	GlobalParam *string `json:"GlobalParam,omitempty" xml:"GlobalParam,omitempty"`
	// Expiration time of the note.
	//
	// example:
	//
	// 2022-01-30 03:03:49
	InvalidTime *string `json:"InvalidTime,omitempty" xml:"InvalidTime,omitempty"`
	// The maximum number of supported bills.
	//
	// example:
	//
	// 9999
	MaxTicketNum *int32 `json:"MaxTicketNum,omitempty" xml:"MaxTicketNum,omitempty"`
	// The ID of the organization.
	//
	// example:
	//
	// 2fe4fbd8-****-****-b3e1-e92c7af083ea
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The registration time of the ticket.
	//
	// example:
	//
	// 2022-01-09 22:23:49
	RegisterTime *string `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// The number of bills that have been consumed.
	//
	// example:
	//
	// 47
	UsedTicketNum *int32 `json:"UsedTicketNum,omitempty" xml:"UsedTicketNum,omitempty"`
	// The user ID of the Quick BI.
	//
	// example:
	//
	// 974e50**********9033f46
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Set the watermarking parameters.
	//
	// example:
	//
	// Tripartite embedding of Ticket
	WatermarkParam *string `json:"WatermarkParam,omitempty" xml:"WatermarkParam,omitempty"`
	// The ID of the operations report.
	//
	// example:
	//
	// ccd3428c-****-****-a608-26bae29dffee
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s QueryTicketInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoResponseBodyResult) GetAccessTicket() *string {
	return s.AccessTicket
}

func (s *QueryTicketInfoResponseBodyResult) GetCmptId() *string {
	return s.CmptId
}

func (s *QueryTicketInfoResponseBodyResult) GetGlobalParam() *string {
	return s.GlobalParam
}

func (s *QueryTicketInfoResponseBodyResult) GetInvalidTime() *string {
	return s.InvalidTime
}

func (s *QueryTicketInfoResponseBodyResult) GetMaxTicketNum() *int32 {
	return s.MaxTicketNum
}

func (s *QueryTicketInfoResponseBodyResult) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *QueryTicketInfoResponseBodyResult) GetRegisterTime() *string {
	return s.RegisterTime
}

func (s *QueryTicketInfoResponseBodyResult) GetUsedTicketNum() *int32 {
	return s.UsedTicketNum
}

func (s *QueryTicketInfoResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *QueryTicketInfoResponseBodyResult) GetWatermarkParam() *string {
	return s.WatermarkParam
}

func (s *QueryTicketInfoResponseBodyResult) GetWorksId() *string {
	return s.WorksId
}

func (s *QueryTicketInfoResponseBodyResult) SetAccessTicket(v string) *QueryTicketInfoResponseBodyResult {
	s.AccessTicket = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetCmptId(v string) *QueryTicketInfoResponseBodyResult {
	s.CmptId = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetGlobalParam(v string) *QueryTicketInfoResponseBodyResult {
	s.GlobalParam = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetInvalidTime(v string) *QueryTicketInfoResponseBodyResult {
	s.InvalidTime = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetMaxTicketNum(v int32) *QueryTicketInfoResponseBodyResult {
	s.MaxTicketNum = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetOrganizationId(v string) *QueryTicketInfoResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetRegisterTime(v string) *QueryTicketInfoResponseBodyResult {
	s.RegisterTime = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetUsedTicketNum(v int32) *QueryTicketInfoResponseBodyResult {
	s.UsedTicketNum = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetUserId(v string) *QueryTicketInfoResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetWatermarkParam(v string) *QueryTicketInfoResponseBodyResult {
	s.WatermarkParam = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) SetWorksId(v string) *QueryTicketInfoResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *QueryTicketInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
