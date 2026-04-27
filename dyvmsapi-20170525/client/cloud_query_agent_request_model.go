// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v int64) *CloudQueryAgentRequest
	GetActive() *int64
	SetCnos(v string) *CloudQueryAgentRequest
	GetCnos() *string
	SetComment(v string) *CloudQueryAgentRequest
	GetComment() *string
	SetEndTime(v string) *CloudQueryAgentRequest
	GetEndTime() *string
	SetEnterpriseId(v int64) *CloudQueryAgentRequest
	GetEnterpriseId() *int64
	SetLimit(v int64) *CloudQueryAgentRequest
	GetLimit() *int64
	SetName(v string) *CloudQueryAgentRequest
	GetName() *string
	SetOrder(v int64) *CloudQueryAgentRequest
	GetOrder() *int64
	SetOwnerId(v int64) *CloudQueryAgentRequest
	GetOwnerId() *int64
	SetQno(v string) *CloudQueryAgentRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *CloudQueryAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudQueryAgentRequest
	GetResourceOwnerId() *int64
	SetReturnQueue(v int64) *CloudQueryAgentRequest
	GetReturnQueue() *int64
	SetStart(v int64) *CloudQueryAgentRequest
	GetStart() *int64
	SetStartTime(v string) *CloudQueryAgentRequest
	GetStartTime() *string
	SetStatus(v int64) *CloudQueryAgentRequest
	GetStatus() *int64
	SetWebrtcUrlType(v int64) *CloudQueryAgentRequest
	GetWebrtcUrlType() *int64
}

type CloudQueryAgentRequest struct {
	// 是否启用；正整数，只能是0或者1，0是停用，1是启用
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 多个座席号以英文逗号分隔 最多支持500个座席
	//
	// example:
	//
	// 1111,2222
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 通过座席备注信息检索；取值说明：前缀模糊匹配
	//
	// example:
	//
	// comment1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 创建时间终止时间点；取值说明："2019-10-11 23:59:59"
	//
	// example:
	//
	// 2026-04-20 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 查询条数；正整数，大于0，最大不能超过1000，默认10
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 通过座席名称检索；取值说明：前后缀模糊匹配
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 排序方式,按照创建时间排序；0正序 1倒序
	//
	// example:
	//
	// 0
	Order   *int64 `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列号；当qno不为空时，查询指定队列的所有座席记录
	//
	// example:
	//
	// 333
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 是否返回绑定座席的队列信息；0:不返回 1:需要返回 默认值:1
	//
	// example:
	//
	// 10
	ReturnQueue *int64 `json:"ReturnQueue,omitempty" xml:"ReturnQueue,omitempty"`
	// 从第几条开始；正整数，大于等于0，默认0
	//
	// example:
	//
	// 0
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// 创建时间起始时间点；取值说明："2019-10-11 00:00:00"
	//
	// example:
	//
	// 2026-04-20 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 是否在线；正整数，只能是0或者1，0是离线，1是在线
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// webrtc软电话返回地址；取值说明：0：企业默认 1：公网域名 2：专线域名 3：公网IP 4：专线IP
	//
	// example:
	//
	// 1
	WebrtcUrlType *int64 `json:"WebrtcUrlType,omitempty" xml:"WebrtcUrlType,omitempty"`
}

func (s CloudQueryAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentRequest) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentRequest) GetActive() *int64 {
	return s.Active
}

func (s *CloudQueryAgentRequest) GetCnos() *string {
	return s.Cnos
}

func (s *CloudQueryAgentRequest) GetComment() *string {
	return s.Comment
}

func (s *CloudQueryAgentRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CloudQueryAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryAgentRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *CloudQueryAgentRequest) GetName() *string {
	return s.Name
}

func (s *CloudQueryAgentRequest) GetOrder() *int64 {
	return s.Order
}

func (s *CloudQueryAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudQueryAgentRequest) GetQno() *string {
	return s.Qno
}

func (s *CloudQueryAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudQueryAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudQueryAgentRequest) GetReturnQueue() *int64 {
	return s.ReturnQueue
}

func (s *CloudQueryAgentRequest) GetStart() *int64 {
	return s.Start
}

func (s *CloudQueryAgentRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudQueryAgentRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudQueryAgentRequest) GetWebrtcUrlType() *int64 {
	return s.WebrtcUrlType
}

func (s *CloudQueryAgentRequest) SetActive(v int64) *CloudQueryAgentRequest {
	s.Active = &v
	return s
}

func (s *CloudQueryAgentRequest) SetCnos(v string) *CloudQueryAgentRequest {
	s.Cnos = &v
	return s
}

func (s *CloudQueryAgentRequest) SetComment(v string) *CloudQueryAgentRequest {
	s.Comment = &v
	return s
}

func (s *CloudQueryAgentRequest) SetEndTime(v string) *CloudQueryAgentRequest {
	s.EndTime = &v
	return s
}

func (s *CloudQueryAgentRequest) SetEnterpriseId(v int64) *CloudQueryAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentRequest) SetLimit(v int64) *CloudQueryAgentRequest {
	s.Limit = &v
	return s
}

func (s *CloudQueryAgentRequest) SetName(v string) *CloudQueryAgentRequest {
	s.Name = &v
	return s
}

func (s *CloudQueryAgentRequest) SetOrder(v int64) *CloudQueryAgentRequest {
	s.Order = &v
	return s
}

func (s *CloudQueryAgentRequest) SetOwnerId(v int64) *CloudQueryAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudQueryAgentRequest) SetQno(v string) *CloudQueryAgentRequest {
	s.Qno = &v
	return s
}

func (s *CloudQueryAgentRequest) SetResourceOwnerAccount(v string) *CloudQueryAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudQueryAgentRequest) SetResourceOwnerId(v int64) *CloudQueryAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudQueryAgentRequest) SetReturnQueue(v int64) *CloudQueryAgentRequest {
	s.ReturnQueue = &v
	return s
}

func (s *CloudQueryAgentRequest) SetStart(v int64) *CloudQueryAgentRequest {
	s.Start = &v
	return s
}

func (s *CloudQueryAgentRequest) SetStartTime(v string) *CloudQueryAgentRequest {
	s.StartTime = &v
	return s
}

func (s *CloudQueryAgentRequest) SetStatus(v int64) *CloudQueryAgentRequest {
	s.Status = &v
	return s
}

func (s *CloudQueryAgentRequest) SetWebrtcUrlType(v int64) *CloudQueryAgentRequest {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudQueryAgentRequest) Validate() error {
	return dara.Validate(s)
}
