// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryForOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryForOperationResponseBody
	GetCode() *string
	SetData(v *QueryForOperationResponseBodyData) *QueryForOperationResponseBody
	GetData() *QueryForOperationResponseBodyData
	SetMessage(v string) *QueryForOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryForOperationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryForOperationResponseBody
	GetSuccess() *bool
}

type QueryForOperationResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query results.
	Data *QueryForOperationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEE90F8C-EDC2-5394-953B-D07A121612B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryForOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryForOperationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryForOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryForOperationResponseBody) GetData() *QueryForOperationResponseBodyData {
	return s.Data
}

func (s *QueryForOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryForOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryForOperationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryForOperationResponseBody) SetCode(v string) *QueryForOperationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryForOperationResponseBody) SetData(v *QueryForOperationResponseBodyData) *QueryForOperationResponseBody {
	s.Data = v
	return s
}

func (s *QueryForOperationResponseBody) SetMessage(v string) *QueryForOperationResponseBody {
	s.Message = &v
	return s
}

func (s *QueryForOperationResponseBody) SetRequestId(v string) *QueryForOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryForOperationResponseBody) SetSuccess(v bool) *QueryForOperationResponseBody {
	s.Success = &v
	return s
}

func (s *QueryForOperationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryForOperationResponseBodyData struct {
	// The event data.
	List []*QueryForOperationResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *QueryForOperationResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s QueryForOperationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryForOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryForOperationResponseBodyData) GetList() []*QueryForOperationResponseBodyDataList {
	return s.List
}

func (s *QueryForOperationResponseBodyData) GetPageInfo() *QueryForOperationResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *QueryForOperationResponseBodyData) SetList(v []*QueryForOperationResponseBodyDataList) *QueryForOperationResponseBodyData {
	s.List = v
	return s
}

func (s *QueryForOperationResponseBodyData) SetPageInfo(v *QueryForOperationResponseBodyDataPageInfo) *QueryForOperationResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *QueryForOperationResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryForOperationResponseBodyDataList struct {
	// The application time.
	//
	// example:
	//
	// 1788873796000
	ApplyTime *int64 `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The asset type. Valid values:
	//
	// - **EIP**: VPC public IP address.
	//
	// - **SLB**: load balancing.
	//
	// - **ECS**: Public IP address in the classic network.
	//
	// example:
	//
	// EIP
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The review remarks.
	//
	// example:
	//
	// agree
	AuditMemo *string `json:"AuditMemo,omitempty" xml:"AuditMemo,omitempty"`
	// The review time.
	//
	// example:
	//
	// 1791937132000
	AuditTime *int64 `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	// The reviewer.
	//
	// example:
	//
	// 149754
	Auditor *string `json:"Auditor,omitempty" xml:"Auditor,omitempty"`
	// The customer UID.
	//
	// example:
	//
	// 1068723216148614
	CustomerUid *string `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// The type of the sent email.
	//
	// example:
	//
	// xxx
	EmailType *string `json:"EmailType,omitempty" xml:"EmailType,omitempty"`
	// Specifies whether the email contains an unsubscribe option. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	HasUnsubscribe *bool `json:"HasUnsubscribe,omitempty" xml:"HasUnsubscribe,omitempty"`
	// The unique ID.
	//
	// example:
	//
	// 710864062
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 10.0.158.58
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The description of the number of sent emails.
	//
	// example:
	//
	// 10
	SendCount *string `json:"SendCount,omitempty" xml:"SendCount,omitempty"`
	// The task status. Valid values:
	//
	// - **auditing**: Under review.
	//
	// - **approved**: Approved.
	//
	// - **rejected**: Rejected.
	//
	// - **effective**: Effective.
	//
	// - **ineffective**: Ineffective.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description.
	//
	// example:
	//
	// xxx
	UseDesc *string `json:"UseDesc,omitempty" xml:"UseDesc,omitempty"`
}

func (s QueryForOperationResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryForOperationResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryForOperationResponseBodyDataList) GetApplyTime() *int64 {
	return s.ApplyTime
}

func (s *QueryForOperationResponseBodyDataList) GetAssetType() *string {
	return s.AssetType
}

func (s *QueryForOperationResponseBodyDataList) GetAuditMemo() *string {
	return s.AuditMemo
}

func (s *QueryForOperationResponseBodyDataList) GetAuditTime() *int64 {
	return s.AuditTime
}

func (s *QueryForOperationResponseBodyDataList) GetAuditor() *string {
	return s.Auditor
}

func (s *QueryForOperationResponseBodyDataList) GetCustomerUid() *string {
	return s.CustomerUid
}

func (s *QueryForOperationResponseBodyDataList) GetEmailType() *string {
	return s.EmailType
}

func (s *QueryForOperationResponseBodyDataList) GetHasUnsubscribe() *bool {
	return s.HasUnsubscribe
}

func (s *QueryForOperationResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *QueryForOperationResponseBodyDataList) GetIp() *string {
	return s.Ip
}

func (s *QueryForOperationResponseBodyDataList) GetSendCount() *string {
	return s.SendCount
}

func (s *QueryForOperationResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *QueryForOperationResponseBodyDataList) GetUseDesc() *string {
	return s.UseDesc
}

func (s *QueryForOperationResponseBodyDataList) SetApplyTime(v int64) *QueryForOperationResponseBodyDataList {
	s.ApplyTime = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetAssetType(v string) *QueryForOperationResponseBodyDataList {
	s.AssetType = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetAuditMemo(v string) *QueryForOperationResponseBodyDataList {
	s.AuditMemo = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetAuditTime(v int64) *QueryForOperationResponseBodyDataList {
	s.AuditTime = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetAuditor(v string) *QueryForOperationResponseBodyDataList {
	s.Auditor = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetCustomerUid(v string) *QueryForOperationResponseBodyDataList {
	s.CustomerUid = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetEmailType(v string) *QueryForOperationResponseBodyDataList {
	s.EmailType = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetHasUnsubscribe(v bool) *QueryForOperationResponseBodyDataList {
	s.HasUnsubscribe = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetId(v int64) *QueryForOperationResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetIp(v string) *QueryForOperationResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetSendCount(v string) *QueryForOperationResponseBodyDataList {
	s.SendCount = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetStatus(v string) *QueryForOperationResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) SetUseDesc(v string) *QueryForOperationResponseBodyDataList {
	s.UseDesc = &v
	return s
}

func (s *QueryForOperationResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type QueryForOperationResponseBodyDataPageInfo struct {
	// The page number from which the query results are displayed.
	//
	// > Default value: 1, which indicates that the query results are displayed from the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of events.
	//
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryForOperationResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryForOperationResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *QueryForOperationResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryForOperationResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryForOperationResponseBodyDataPageInfo) GetTotal() *int32 {
	return s.Total
}

func (s *QueryForOperationResponseBodyDataPageInfo) SetCurrentPage(v int32) *QueryForOperationResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *QueryForOperationResponseBodyDataPageInfo) SetPageSize(v int32) *QueryForOperationResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *QueryForOperationResponseBodyDataPageInfo) SetTotal(v int32) *QueryForOperationResponseBodyDataPageInfo {
	s.Total = &v
	return s
}

func (s *QueryForOperationResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
