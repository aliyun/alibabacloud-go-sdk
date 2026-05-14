// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrObsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedId(v int64) *ClinkListCdrObsRequest
	GetAssociatedId() *int64
	SetCity(v string) *ClinkListCdrObsRequest
	GetCity() *string
	SetClientNumber(v string) *ClinkListCdrObsRequest
	GetClientNumber() *string
	SetCno(v string) *ClinkListCdrObsRequest
	GetCno() *string
	SetCustomerNumber(v string) *ClinkListCdrObsRequest
	GetCustomerNumber() *string
	SetEndTime(v int64) *ClinkListCdrObsRequest
	GetEndTime() *int64
	SetEnterpriseId(v int64) *ClinkListCdrObsRequest
	GetEnterpriseId() *int64
	SetEvaluation(v int64) *ClinkListCdrObsRequest
	GetEvaluation() *int64
	SetHiddenType(v int64) *ClinkListCdrObsRequest
	GetHiddenType() *int64
	SetHotline(v string) *ClinkListCdrObsRequest
	GetHotline() *string
	SetIdType(v int64) *ClinkListCdrObsRequest
	GetIdType() *int64
	SetLimit(v int64) *ClinkListCdrObsRequest
	GetLimit() *int64
	SetMainUniqueId(v string) *ClinkListCdrObsRequest
	GetMainUniqueId() *string
	SetMark(v int64) *ClinkListCdrObsRequest
	GetMark() *int64
	SetOffset(v int64) *ClinkListCdrObsRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListCdrObsRequest
	GetOwnerId() *int64
	SetProvince(v string) *ClinkListCdrObsRequest
	GetProvince() *string
	SetQueueAnswerInTime(v int64) *ClinkListCdrObsRequest
	GetQueueAnswerInTime() *int64
	SetRequestUniqueId(v string) *ClinkListCdrObsRequest
	GetRequestUniqueId() *string
	SetResourceOwnerAccount(v string) *ClinkListCdrObsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListCdrObsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v int64) *ClinkListCdrObsRequest
	GetStartTime() *int64
	SetStatus(v int64) *ClinkListCdrObsRequest
	GetStatus() *int64
}

type ClinkListCdrObsRequest struct {
	// 业务ID
	//
	// example:
	//
	// AssociatedId
	AssociatedId *int64 `json:"AssociatedId,omitempty" xml:"AssociatedId,omitempty"`
	// 客户城市
	//
	// example:
	//
	// 示例值示例值
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// 座席电话
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席号，要求只能是 4-11 位数字
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户号码
	//
	// example:
	//
	// 177xxxx
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 结束时间，时间戳格式精确到秒，开始时间和结束时间跨度不能超过一个月。默认值取当前时间
	//
	// example:
	//
	// 1775024775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否邀评: 0: 否 1: 是 2: -
	//
	// example:
	//
	// 0
	Evaluation *int64 `json:"Evaluation,omitempty" xml:"Evaluation,omitempty"`
	// 是否隐藏号码。 0: 不隐藏，1: 中间四位，2: 最后八位 3: 全部号码，4: 最后四位。默认值为 0
	//
	// example:
	//
	// 0
	HiddenType *int64 `json:"HiddenType,omitempty" xml:"HiddenType,omitempty"`
	// 热线号码
	//
	// example:
	//
	// Hotline
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 业务ID类型，1：工单；2：业务记录
	//
	// example:
	//
	// 44
	IdType *int64 `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// 查询条数，范围 10-100。默认值为 10。注：limit + offset 不允许超过100000
	//
	// example:
	//
	// 100
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// MainUniqueId
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 标记。取值范围如下： 为空表示全部， 1:留言 2:转移 3:咨询 4:三方 5:传真接收 6:会议 7:交互 8:IVR中放弃 9:已进入IVR 10:未进入IVR 11:队列中放弃 12:队列中溢出 注： 其中mark值2,3,4,7仅在status=1(座席接听)时有效， mark值1,5,6,8,9,10,11,12在status=3(系统接听)时有效
	//
	// example:
	//
	// 3
	Mark *int64 `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// 偏移量，范围 0-99990。默认值为 0。注：limit + offset 不允许超过100000
	//
	// example:
	//
	// 0
	Offset  *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 客户省份
	//
	// example:
	//
	// 示例值示例值
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// 队列及时应答: 0: 否 1: 是 2: -
	//
	// example:
	//
	// 0
	QueueAnswerInTime *int64 `json:"QueueAnswerInTime,omitempty" xml:"QueueAnswerInTime,omitempty"`
	// 请求唯一标识
	//
	// example:
	//
	// RequestUniqueId
	RequestUniqueId      *string `json:"RequestUniqueId,omitempty" xml:"RequestUniqueId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 开始时间，时间戳格式精确到秒。默认值取当前月份第一天
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 接听状态。取值范围如下： 0: 全部 1: 客户未接听 2: 座席未接听 3: 双方接听 默认值为0
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ClinkListCdrObsRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObsRequest) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObsRequest) GetAssociatedId() *int64 {
	return s.AssociatedId
}

func (s *ClinkListCdrObsRequest) GetCity() *string {
	return s.City
}

func (s *ClinkListCdrObsRequest) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrObsRequest) GetCno() *string {
	return s.Cno
}

func (s *ClinkListCdrObsRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrObsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrObsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListCdrObsRequest) GetEvaluation() *int64 {
	return s.Evaluation
}

func (s *ClinkListCdrObsRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkListCdrObsRequest) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkListCdrObsRequest) GetIdType() *int64 {
	return s.IdType
}

func (s *ClinkListCdrObsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListCdrObsRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrObsRequest) GetMark() *int64 {
	return s.Mark
}

func (s *ClinkListCdrObsRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListCdrObsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListCdrObsRequest) GetProvince() *string {
	return s.Province
}

func (s *ClinkListCdrObsRequest) GetQueueAnswerInTime() *int64 {
	return s.QueueAnswerInTime
}

func (s *ClinkListCdrObsRequest) GetRequestUniqueId() *string {
	return s.RequestUniqueId
}

func (s *ClinkListCdrObsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListCdrObsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListCdrObsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrObsRequest) GetStatus() *int64 {
	return s.Status
}

func (s *ClinkListCdrObsRequest) SetAssociatedId(v int64) *ClinkListCdrObsRequest {
	s.AssociatedId = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetCity(v string) *ClinkListCdrObsRequest {
	s.City = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetClientNumber(v string) *ClinkListCdrObsRequest {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetCno(v string) *ClinkListCdrObsRequest {
	s.Cno = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetCustomerNumber(v string) *ClinkListCdrObsRequest {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetEndTime(v int64) *ClinkListCdrObsRequest {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetEnterpriseId(v int64) *ClinkListCdrObsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetEvaluation(v int64) *ClinkListCdrObsRequest {
	s.Evaluation = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetHiddenType(v int64) *ClinkListCdrObsRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetHotline(v string) *ClinkListCdrObsRequest {
	s.Hotline = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetIdType(v int64) *ClinkListCdrObsRequest {
	s.IdType = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetLimit(v int64) *ClinkListCdrObsRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetMainUniqueId(v string) *ClinkListCdrObsRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetMark(v int64) *ClinkListCdrObsRequest {
	s.Mark = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetOffset(v int64) *ClinkListCdrObsRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetOwnerId(v int64) *ClinkListCdrObsRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetProvince(v string) *ClinkListCdrObsRequest {
	s.Province = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetQueueAnswerInTime(v int64) *ClinkListCdrObsRequest {
	s.QueueAnswerInTime = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetRequestUniqueId(v string) *ClinkListCdrObsRequest {
	s.RequestUniqueId = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetResourceOwnerAccount(v string) *ClinkListCdrObsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetResourceOwnerId(v int64) *ClinkListCdrObsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetStartTime(v int64) *ClinkListCdrObsRequest {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrObsRequest) SetStatus(v int64) *ClinkListCdrObsRequest {
	s.Status = &v
	return s
}

func (s *ClinkListCdrObsRequest) Validate() error {
	return dara.Validate(s)
}
