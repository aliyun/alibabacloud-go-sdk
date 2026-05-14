// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrObResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListCdrObResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListCdrObResponseBody
	GetCode() *string
	SetData(v *ClinkListCdrObResponseBodyData) *ClinkListCdrObResponseBody
	GetData() *ClinkListCdrObResponseBodyData
	SetMessage(v string) *ClinkListCdrObResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListCdrObResponseBody
	GetRequestId() *string
}

type ClinkListCdrObResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListCdrObResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListCdrObResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListCdrObResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListCdrObResponseBody) GetData() *ClinkListCdrObResponseBodyData {
	return s.Data
}

func (s *ClinkListCdrObResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListCdrObResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListCdrObResponseBody) SetAccessDeniedDetail(v string) *ClinkListCdrObResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListCdrObResponseBody) SetCode(v string) *ClinkListCdrObResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListCdrObResponseBody) SetData(v *ClinkListCdrObResponseBodyData) *ClinkListCdrObResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListCdrObResponseBody) SetMessage(v string) *ClinkListCdrObResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListCdrObResponseBody) SetRequestId(v string) *ClinkListCdrObResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListCdrObResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListCdrObResponseBodyData struct {
	// [外呼记录列表] #外呼记录列表
	CdrOb []*ClinkListCdrObResponseBodyDataCdrOb `json:"CdrOb,omitempty" xml:"CdrOb,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 一页展示条数
	//
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 滚动查询ID
	//
	// example:
	//
	// ScrollId
	ScrollId *string `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
	// 总条数
	//
	// example:
	//
	// 31
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListCdrObResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObResponseBodyData) GetCdrOb() []*ClinkListCdrObResponseBodyDataCdrOb {
	return s.CdrOb
}

func (s *ClinkListCdrObResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListCdrObResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListCdrObResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListCdrObResponseBodyData) GetScrollId() *string {
	return s.ScrollId
}

func (s *ClinkListCdrObResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListCdrObResponseBodyData) SetCdrOb(v []*ClinkListCdrObResponseBodyDataCdrOb) *ClinkListCdrObResponseBodyData {
	s.CdrOb = v
	return s
}

func (s *ClinkListCdrObResponseBodyData) SetClinkRequestId(v string) *ClinkListCdrObResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyData) SetPageNumber(v int64) *ClinkListCdrObResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListCdrObResponseBodyData) SetPageSize(v int64) *ClinkListCdrObResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListCdrObResponseBodyData) SetScrollId(v string) *ClinkListCdrObResponseBodyData {
	s.ScrollId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyData) SetTotalCount(v int64) *ClinkListCdrObResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListCdrObResponseBodyData) Validate() error {
	if s.CdrOb != nil {
		for _, item := range s.CdrOb {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListCdrObResponseBodyDataCdrOb struct {
	// 转写记录标识 1：存在，0：不存在
	//
	// example:
	//
	// 0
	AsrRecord *int64 `json:"AsrRecord,omitempty" xml:"AsrRecord,omitempty"`
	// 通话时长
	//
	// example:
	//
	// 76
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 客户接听时间
	//
	// example:
	//
	// 1715326248
	BridgeTime *int64 `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 业务记录ID
	BusinessIds []*int64 `json:"BusinessIds,omitempty" xml:"BusinessIds,omitempty" type:"Repeated"`
	// CallID
	//
	// example:
	//
	// xxx
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 通话类型
	//
	// example:
	//
	// 示例值
	CallType *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 通话类型Key值
	//
	// example:
	//
	// 4
	CallTypeKey *int64 `json:"CallTypeKey,omitempty" xml:"CallTypeKey,omitempty"`
	// 客户响铃时长
	//
	// example:
	//
	// 29
	CalleeRingingDuration *int64 `json:"CalleeRingingDuration,omitempty" xml:"CalleeRingingDuration,omitempty"`
	// 客户响铃时间
	//
	// example:
	//
	// 1715326251
	CalleeRingingTime *int64 `json:"CalleeRingingTime,omitempty" xml:"CalleeRingingTime,omitempty"`
	// 座席名称
	//
	// example:
	//
	// ClientName
	ClientName *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	// 座席电话
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户地区，省份/城市
	//
	// example:
	//
	// 示例值
	CustomerArea *string `json:"CustomerArea,omitempty" xml:"CustomerArea,omitempty"`
	// 客户号码区号
	//
	// example:
	//
	// 010
	CustomerAreaCode *string `json:"CustomerAreaCode,omitempty" xml:"CustomerAreaCode,omitempty"`
	// 客户号码城市
	//
	// example:
	//
	// 示例值示例值示例值
	CustomerCity *string `json:"CustomerCity,omitempty" xml:"CustomerCity,omitempty"`
	// 客户号码
	//
	// example:
	//
	// xxx
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 客户号码加密串
	//
	// example:
	//
	// xxx
	CustomerNumberEncrypt *string `json:"CustomerNumberEncrypt,omitempty" xml:"CustomerNumberEncrypt,omitempty"`
	// 追加客户邮编
	//
	// example:
	//
	// xxx
	CustomerPostCode *string `json:"CustomerPostCode,omitempty" xml:"CustomerPostCode,omitempty"`
	// 客户号码省份
	//
	// example:
	//
	// 示例值
	CustomerProvince *string `json:"CustomerProvince,omitempty" xml:"CustomerProvince,omitempty"`
	// 客户VIP标识
	//
	// example:
	//
	// 1
	CustomerVip *int64 `json:"CustomerVip,omitempty" xml:"CustomerVip,omitempty"`
	// 结束时间
	//
	// example:
	//
	// 1715326248
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 企业号
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否邀评
	//
	// example:
	//
	// 0
	Evaluation *string `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
	// 是否存在业务记录
	//
	// example:
	//
	// 0
	ExistBusiness *int64 `json:"ExistBusiness,omitempty" xml:"ExistBusiness,omitempty"`
	// 是否存在客户资料
	//
	// example:
	//
	// true
	ExistCustomer *bool `json:"ExistCustomer,omitempty" xml:"ExistCustomer,omitempty"`
	// 是否存在工单
	//
	// example:
	//
	// 0
	ExistTicket *int64 `json:"ExistTicket,omitempty" xml:"ExistTicket,omitempty"`
	// 满意度评价
	//
	// example:
	//
	// null
	InvestigationKeys *int64 `json:"InvestigationKeys,omitempty" xml:"InvestigationKeys,omitempty"`
	// 外显号码
	//
	// example:
	//
	// xxx
	LeftClid *string `json:"LeftClid,omitempty" xml:"LeftClid,omitempty"`
	// 外显号码归属地区，省份/城市
	//
	// example:
	//
	// 示例值示例值
	LeftClidArea *string `json:"LeftClidArea,omitempty" xml:"LeftClidArea,omitempty"`
	// 外显号码归属城市
	//
	// example:
	//
	// 示例值示例值示例值
	LeftClidCity *string `json:"LeftClidCity,omitempty" xml:"LeftClidCity,omitempty"`
	// 外显号码归属省份
	//
	// example:
	//
	// 示例值示例值
	LeftClidProvince *string `json:"LeftClidProvince,omitempty" xml:"LeftClidProvince,omitempty"`
	// 通话ID
	//
	// example:
	//
	// medias_1-171532xxxx.53
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 备注
	//
	// example:
	//
	// 2
	MarkData *string `json:"MarkData,omitempty" xml:"MarkData,omitempty"`
	// 呼叫结果
	//
	// example:
	//
	// 示例值
	ObSipCause *string `json:"ObSipCause,omitempty" xml:"ObSipCause,omitempty"`
	// 被叫状态
	//
	// example:
	//
	// 示例值示例值示例值
	ObSipCauseRaw *string `json:"ObSipCauseRaw,omitempty" xml:"ObSipCauseRaw,omitempty"`
	// 外呼等待时长
	//
	// example:
	//
	// 24
	ObWaitDuration *int64 `json:"ObWaitDuration,omitempty" xml:"ObWaitDuration,omitempty"`
	// 挂断方
	//
	// example:
	//
	// 示例值示例值
	OnHookSource *string `json:"OnHookSource,omitempty" xml:"OnHookSource,omitempty"`
	// 响铃前等待时长
	//
	// example:
	//
	// 10
	PreRingWaitDuration *int64 `json:"PreRingWaitDuration,omitempty" xml:"PreRingWaitDuration,omitempty"`
	// 录音文件
	//
	// example:
	//
	// RecordFile
	RecordFile *string `json:"RecordFile,omitempty" xml:"RecordFile,omitempty"`
	// 通话唯一ID
	//
	// example:
	//
	// xxx
	RequestUniqueId *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	// RtcUid
	//
	// example:
	//
	// xxx
	RtcUid *string `json:"RtcUid,omitempty" xml:"RtcUid,omitempty"`
	// 开始时间
	//
	// example:
	//
	// 1715326248
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 接听状态值
	//
	// example:
	//
	// 3
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 接听状态描述
	//
	// example:
	//
	// 示例值示例值
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	// 通话标签
	TagNames []*string `json:"TagNames,omitempty" xml:"TagNames,omitempty" type:"Repeated"`
	// 外呼任务id
	//
	// example:
	//
	// 85
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 外呼任务详情id
	//
	// example:
	//
	// 83
	TaskInventoryId *int64 `json:"TaskInventoryId,omitempty" xml:"TaskInventoryId,omitempty"`
	// 外呼任务名称
	//
	// example:
	//
	// TaskName
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 工单id
	TicketIds []*int64 `json:"TicketIds,omitempty" xml:"TicketIds,omitempty" type:"Repeated"`
	// 总时长
	//
	// example:
	//
	// 27
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 中继组号
	//
	// example:
	//
	// 10088
	TrunkGroupKey *string `json:"TrunkGroupKey,omitempty" xml:"TrunkGroupKey,omitempty"`
	// 座席接起时间
	//
	// example:
	//
	// 1715326251
	UpTime *int64 `json:"UpTime,omitempty" xml:"UpTime,omitempty"`
	// WebRTCCallID
	//
	// example:
	//
	// xxxx
	WebrtcCallId *string `json:"WebrtcCallId,omitempty" xml:"WebrtcCallId,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// xxx
	Xnumber *string `json:"Xnumber,omitempty" xml:"Xnumber,omitempty"`
	// 虚拟号码归属地区，省份/城市
	//
	// example:
	//
	// xxx
	XnumberArea *string `json:"XnumberArea,omitempty" xml:"XnumberArea,omitempty"`
	// 虚拟号码归属城市
	//
	// example:
	//
	// 示例值示例值
	XnumberCity *string `json:"XnumberCity,omitempty" xml:"XnumberCity,omitempty"`
	// 虚拟号码归属省份
	//
	// example:
	//
	// null
	XnumberProvince *string `json:"XnumberProvince,omitempty" xml:"XnumberProvince,omitempty"`
}

func (s ClinkListCdrObResponseBodyDataCdrOb) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObResponseBodyDataCdrOb) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetAsrRecord() *int64 {
	return s.AsrRecord
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetBusinessIds() []*int64 {
	return s.BusinessIds
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCallId() *string {
	return s.CallId
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCallType() *string {
	return s.CallType
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCallTypeKey() *int64 {
	return s.CallTypeKey
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCalleeRingingDuration() *int64 {
	return s.CalleeRingingDuration
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCalleeRingingTime() *int64 {
	return s.CalleeRingingTime
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetClientName() *string {
	return s.ClientName
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCno() *string {
	return s.Cno
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerArea() *string {
	return s.CustomerArea
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerAreaCode() *string {
	return s.CustomerAreaCode
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerCity() *string {
	return s.CustomerCity
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerNumberEncrypt() *string {
	return s.CustomerNumberEncrypt
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerPostCode() *string {
	return s.CustomerPostCode
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerProvince() *string {
	return s.CustomerProvince
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetCustomerVip() *int64 {
	return s.CustomerVip
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetEvaluation() *string {
	return s.Evaluation
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetExistBusiness() *int64 {
	return s.ExistBusiness
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetExistCustomer() *bool {
	return s.ExistCustomer
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetExistTicket() *int64 {
	return s.ExistTicket
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetInvestigationKeys() *int64 {
	return s.InvestigationKeys
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetLeftClid() *string {
	return s.LeftClid
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetLeftClidArea() *string {
	return s.LeftClidArea
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetLeftClidCity() *string {
	return s.LeftClidCity
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetLeftClidProvince() *string {
	return s.LeftClidProvince
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetMarkData() *string {
	return s.MarkData
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetObSipCause() *string {
	return s.ObSipCause
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetObSipCauseRaw() *string {
	return s.ObSipCauseRaw
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetObWaitDuration() *int64 {
	return s.ObWaitDuration
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetOnHookSource() *string {
	return s.OnHookSource
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetPreRingWaitDuration() *int64 {
	return s.PreRingWaitDuration
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetRecordFile() *string {
	return s.RecordFile
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetRtcUid() *string {
	return s.RtcUid
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetStatus() *int64 {
	return s.Status
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetTagNames() []*string {
	return s.TagNames
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetTaskInventoryId() *int64 {
	return s.TaskInventoryId
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetTaskName() *string {
	return s.TaskName
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetTicketIds() []*int64 {
	return s.TicketIds
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetTrunkGroupKey() *string {
	return s.TrunkGroupKey
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetUpTime() *int64 {
	return s.UpTime
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetWebrtcCallId() *string {
	return s.WebrtcCallId
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetXnumber() *string {
	return s.Xnumber
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetXnumberArea() *string {
	return s.XnumberArea
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetXnumberCity() *string {
	return s.XnumberCity
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) GetXnumberProvince() *string {
	return s.XnumberProvince
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetAsrRecord(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.AsrRecord = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetBridgeDuration(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetBridgeTime(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.BridgeTime = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetBusinessIds(v []*int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.BusinessIds = v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCallId(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CallId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCallType(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CallType = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCallTypeKey(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CallTypeKey = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCalleeRingingDuration(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CalleeRingingDuration = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCalleeRingingTime(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CalleeRingingTime = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetClientName(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ClientName = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetClientNumber(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCno(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.Cno = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerArea(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerArea = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerAreaCode(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerAreaCode = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerCity(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerCity = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerNumber(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerNumberEncrypt(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerNumberEncrypt = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerPostCode(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerPostCode = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerProvince(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerProvince = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetCustomerVip(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.CustomerVip = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetEndTime(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetEnterpriseId(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetEvaluation(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.Evaluation = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetExistBusiness(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ExistBusiness = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetExistCustomer(v bool) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ExistCustomer = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetExistTicket(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ExistTicket = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetInvestigationKeys(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.InvestigationKeys = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetLeftClid(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.LeftClid = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetLeftClidArea(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.LeftClidArea = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetLeftClidCity(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.LeftClidCity = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetLeftClidProvince(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.LeftClidProvince = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetMainUniqueId(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetMarkData(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.MarkData = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetObSipCause(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ObSipCause = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetObSipCauseRaw(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ObSipCauseRaw = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetObWaitDuration(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.ObWaitDuration = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetOnHookSource(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.OnHookSource = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetPreRingWaitDuration(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.PreRingWaitDuration = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetRecordFile(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.RecordFile = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetRequestUniqueId(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.RequestUniqueId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetRtcUid(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.RtcUid = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetStartTime(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetStatus(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.Status = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetStatusDesc(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.StatusDesc = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetTagNames(v []*string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.TagNames = v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetTaskId(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.TaskId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetTaskInventoryId(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.TaskInventoryId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetTaskName(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.TaskName = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetTicketIds(v []*int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.TicketIds = v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetTotalDuration(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.TotalDuration = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetTrunkGroupKey(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.TrunkGroupKey = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetUpTime(v int64) *ClinkListCdrObResponseBodyDataCdrOb {
	s.UpTime = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetWebrtcCallId(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.WebrtcCallId = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetXnumber(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.Xnumber = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetXnumberArea(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.XnumberArea = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetXnumberCity(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.XnumberCity = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) SetXnumberProvince(v string) *ClinkListCdrObResponseBodyDataCdrOb {
	s.XnumberProvince = &v
	return s
}

func (s *ClinkListCdrObResponseBodyDataCdrOb) Validate() error {
	return dara.Validate(s)
}
