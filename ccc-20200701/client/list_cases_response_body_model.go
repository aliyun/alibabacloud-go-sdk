// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCasesResponseBody
	GetCode() *string
	SetData(v *ListCasesResponseBodyData) *ListCasesResponseBody
	GetData() *ListCasesResponseBodyData
	SetHttpStatusCode(v int64) *ListCasesResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListCasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCasesResponseBody
	GetRequestId() *string
}

type ListCasesResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data.
	Data *ListCasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 7CC6523B-0E51-1B62-8DA5-6A9831CAE316
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCasesResponseBody) GetData() *ListCasesResponseBodyData {
	return s.Data
}

func (s *ListCasesResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListCasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCasesResponseBody) SetCode(v string) *ListCasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCasesResponseBody) SetData(v *ListCasesResponseBodyData) *ListCasesResponseBody {
	s.Data = v
	return s
}

func (s *ListCasesResponseBody) SetHttpStatusCode(v int64) *ListCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCasesResponseBody) SetMessage(v string) *ListCasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCasesResponseBody) SetRequestId(v string) *ListCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCasesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCasesResponseBodyData struct {
	// List of contact list execution details.
	List []*ListCasesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number, ranging from 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, ranging from 1 to 100.
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBodyData) GetList() []*ListCasesResponseBodyDataList {
	return s.List
}

func (s *ListCasesResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCasesResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCasesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCasesResponseBodyData) SetList(v []*ListCasesResponseBodyDataList) *ListCasesResponseBodyData {
	s.List = v
	return s
}

func (s *ListCasesResponseBodyData) SetPageNumber(v int64) *ListCasesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCasesResponseBodyData) SetPageSize(v int64) *ListCasesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCasesResponseBodyData) SetTotalCount(v int64) *ListCasesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCasesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCasesResponseBodyDataList struct {
	// Phase at which the call was abandoned.
	//
	// - IVR (IVR interaction phase)
	//
	// - Queuing (queuing phase)
	//
	// - Ringing (agent ringing phase)
	//
	// example:
	//
	// Ringing
	AbandonPhase *string `json:"AbandonPhase,omitempty" xml:"AbandonPhase,omitempty"`
	// The type of call abandonment, that is, the reason why the call was not successfully connected.
	//
	// example:
	//
	// NA
	AbandonType *string `json:"AbandonType,omitempty" xml:"AbandonType,omitempty"`
	// Number of attempts, which is the total number of calls made to this contact.
	//
	// example:
	//
	// 1
	AttemptCount *int64 `json:"AttemptCount,omitempty" xml:"AttemptCount,omitempty"`
	// System-generated contact ID. Customers do not need to concern themselves with this.
	//
	// example:
	//
	// 60ecb1a2-4480-4d01-bede-c5b7655bfadf
	CaseId *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// Custom variables defined by the customer, formatted as a JSON object. The object can contain up to 10 properties, each with a name and value defined by the customer. These can be configured when creating a predictive outbound dialing activity.
	//
	// example:
	//
	// {"name":"yy","客戶标签":"tag-yy"}
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	// Reason for outbound call failure.
	//
	// example:
	//
	// 无
	FailureReason *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	// Phone number.
	//
	// example:
	//
	// 1888888****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Status.
	//
	// example:
	//
	// Connected
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListCasesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListCasesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBodyDataList) GetAbandonPhase() *string {
	return s.AbandonPhase
}

func (s *ListCasesResponseBodyDataList) GetAbandonType() *string {
	return s.AbandonType
}

func (s *ListCasesResponseBodyDataList) GetAttemptCount() *int64 {
	return s.AttemptCount
}

func (s *ListCasesResponseBodyDataList) GetCaseId() *string {
	return s.CaseId
}

func (s *ListCasesResponseBodyDataList) GetCustomVariables() *string {
	return s.CustomVariables
}

func (s *ListCasesResponseBodyDataList) GetFailureReason() *string {
	return s.FailureReason
}

func (s *ListCasesResponseBodyDataList) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListCasesResponseBodyDataList) GetState() *string {
	return s.State
}

func (s *ListCasesResponseBodyDataList) SetAbandonPhase(v string) *ListCasesResponseBodyDataList {
	s.AbandonPhase = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetAbandonType(v string) *ListCasesResponseBodyDataList {
	s.AbandonType = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetAttemptCount(v int64) *ListCasesResponseBodyDataList {
	s.AttemptCount = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCaseId(v string) *ListCasesResponseBodyDataList {
	s.CaseId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCustomVariables(v string) *ListCasesResponseBodyDataList {
	s.CustomVariables = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetFailureReason(v string) *ListCasesResponseBodyDataList {
	s.FailureReason = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetPhoneNumber(v string) *ListCasesResponseBodyDataList {
	s.PhoneNumber = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetState(v string) *ListCasesResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListCasesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
