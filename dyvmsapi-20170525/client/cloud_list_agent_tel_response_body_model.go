// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAgentTelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListAgentTelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListAgentTelResponseBody
	GetCode() *string
	SetData(v *CloudListAgentTelResponseBodyData) *CloudListAgentTelResponseBody
	GetData() *CloudListAgentTelResponseBodyData
	SetMessage(v string) *CloudListAgentTelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListAgentTelResponseBody
	GetRequestId() *string
}

type CloudListAgentTelResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListAgentTelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListAgentTelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentTelResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListAgentTelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListAgentTelResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListAgentTelResponseBody) GetData() *CloudListAgentTelResponseBodyData {
	return s.Data
}

func (s *CloudListAgentTelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListAgentTelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListAgentTelResponseBody) SetAccessDeniedDetail(v string) *CloudListAgentTelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListAgentTelResponseBody) SetCode(v string) *CloudListAgentTelResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListAgentTelResponseBody) SetData(v *CloudListAgentTelResponseBodyData) *CloudListAgentTelResponseBody {
	s.Data = v
	return s
}

func (s *CloudListAgentTelResponseBody) SetMessage(v string) *CloudListAgentTelResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListAgentTelResponseBody) SetRequestId(v string) *CloudListAgentTelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListAgentTelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListAgentTelResponseBodyData struct {
	// 座席电话数组
	List []*CloudListAgentTelResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudListAgentTelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentTelResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListAgentTelResponseBodyData) GetList() []*CloudListAgentTelResponseBodyDataList {
	return s.List
}

func (s *CloudListAgentTelResponseBodyData) SetList(v []*CloudListAgentTelResponseBodyDataList) *CloudListAgentTelResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListAgentTelResponseBodyData) Validate() error {
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

type CloudListAgentTelResponseBodyDataList struct {
	// 座席id
	//
	// example:
	//
	// 19
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 7098
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7122600
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 座席电话id
	//
	// example:
	//
	// 9
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否绑定，0:未绑定 1:绑定
	//
	// example:
	//
	// 1
	IsBind *int64 `json:"IsBind,omitempty" xml:"IsBind,omitempty"`
	// 是否验证，0:未验证 1:已验证
	//
	// example:
	//
	// 1
	IsValidity *int64 `json:"IsValidity,omitempty" xml:"IsValidity,omitempty"`
	// 电话号码
	//
	// example:
	//
	// 41008502
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
	// 电话类型，1:固话 2:手机 3:分机 4:软电话
	//
	// example:
	//
	// 1
	TelType *int64 `json:"TelType,omitempty" xml:"TelType,omitempty"`
}

func (s CloudListAgentTelResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentTelResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListAgentTelResponseBodyDataList) GetAgentId() *int64 {
	return s.AgentId
}

func (s *CloudListAgentTelResponseBodyDataList) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudListAgentTelResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudListAgentTelResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListAgentTelResponseBodyDataList) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListAgentTelResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *CloudListAgentTelResponseBodyDataList) GetIsBind() *int64 {
	return s.IsBind
}

func (s *CloudListAgentTelResponseBodyDataList) GetIsValidity() *int64 {
	return s.IsValidity
}

func (s *CloudListAgentTelResponseBodyDataList) GetTel() *string {
	return s.Tel
}

func (s *CloudListAgentTelResponseBodyDataList) GetTelType() *int64 {
	return s.TelType
}

func (s *CloudListAgentTelResponseBodyDataList) SetAgentId(v int64) *CloudListAgentTelResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetAreaCode(v string) *CloudListAgentTelResponseBodyDataList {
	s.AreaCode = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetCno(v string) *CloudListAgentTelResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetCreateTime(v string) *CloudListAgentTelResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetEnterpriseId(v int64) *CloudListAgentTelResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetId(v int64) *CloudListAgentTelResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetIsBind(v int64) *CloudListAgentTelResponseBodyDataList {
	s.IsBind = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetIsValidity(v int64) *CloudListAgentTelResponseBodyDataList {
	s.IsValidity = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetTel(v string) *CloudListAgentTelResponseBodyDataList {
	s.Tel = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) SetTelType(v int64) *CloudListAgentTelResponseBodyDataList {
	s.TelType = &v
	return s
}

func (s *CloudListAgentTelResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
