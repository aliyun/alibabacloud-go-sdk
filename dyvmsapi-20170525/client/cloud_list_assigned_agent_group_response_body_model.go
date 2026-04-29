// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAssignedAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListAssignedAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListAssignedAgentGroupResponseBody
	GetCode() *string
	SetData(v *CloudListAssignedAgentGroupResponseBodyData) *CloudListAssignedAgentGroupResponseBody
	GetData() *CloudListAssignedAgentGroupResponseBodyData
	SetMessage(v string) *CloudListAssignedAgentGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListAssignedAgentGroupResponseBody
	GetRequestId() *string
}

type CloudListAssignedAgentGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListAssignedAgentGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListAssignedAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListAssignedAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListAssignedAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListAssignedAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListAssignedAgentGroupResponseBody) GetData() *CloudListAssignedAgentGroupResponseBodyData {
	return s.Data
}

func (s *CloudListAssignedAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListAssignedAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListAssignedAgentGroupResponseBody) SetAccessDeniedDetail(v string) *CloudListAssignedAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBody) SetCode(v string) *CloudListAssignedAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBody) SetData(v *CloudListAssignedAgentGroupResponseBodyData) *CloudListAssignedAgentGroupResponseBody {
	s.Data = v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBody) SetMessage(v string) *CloudListAssignedAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBody) SetRequestId(v string) *CloudListAssignedAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListAssignedAgentGroupResponseBodyData struct {
	// 返回数据
	List []*CloudListAssignedAgentGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudListAssignedAgentGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListAssignedAgentGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListAssignedAgentGroupResponseBodyData) GetList() []*CloudListAssignedAgentGroupResponseBodyDataList {
	return s.List
}

func (s *CloudListAssignedAgentGroupResponseBodyData) SetList(v []*CloudListAssignedAgentGroupResponseBodyDataList) *CloudListAssignedAgentGroupResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBodyData) Validate() error {
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

type CloudListAssignedAgentGroupResponseBodyDataList struct {
	// 座席名称
	//
	// example:
	//
	// name1
	Cname *int64 `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// 座席编号
	//
	// example:
	//
	// 59517582
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2026-03-30 06:12:30
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 呼叫中心编号
	//
	// example:
	//
	// 7122600
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 外呼组编号
	//
	// example:
	//
	// ZO874912
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
}

func (s CloudListAssignedAgentGroupResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListAssignedAgentGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) GetCname() *int64 {
	return s.Cname
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) GetGno() *string {
	return s.Gno
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) SetCname(v int64) *CloudListAssignedAgentGroupResponseBodyDataList {
	s.Cname = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) SetCno(v string) *CloudListAssignedAgentGroupResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) SetCreateTime(v string) *CloudListAssignedAgentGroupResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) SetEnterpriseId(v string) *CloudListAssignedAgentGroupResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) SetGno(v string) *CloudListAssignedAgentGroupResponseBodyDataList {
	s.Gno = &v
	return s
}

func (s *CloudListAssignedAgentGroupResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
