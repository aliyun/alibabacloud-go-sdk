// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListAgentGroupResponseBody
	GetCode() *string
	SetData(v *CloudListAgentGroupResponseBodyData) *CloudListAgentGroupResponseBody
	GetData() *CloudListAgentGroupResponseBodyData
	SetMessage(v string) *CloudListAgentGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListAgentGroupResponseBody
	GetRequestId() *string
}

type CloudListAgentGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListAgentGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListAgentGroupResponseBody) GetData() *CloudListAgentGroupResponseBodyData {
	return s.Data
}

func (s *CloudListAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListAgentGroupResponseBody) SetAccessDeniedDetail(v string) *CloudListAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListAgentGroupResponseBody) SetCode(v string) *CloudListAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListAgentGroupResponseBody) SetData(v *CloudListAgentGroupResponseBodyData) *CloudListAgentGroupResponseBody {
	s.Data = v
	return s
}

func (s *CloudListAgentGroupResponseBody) SetMessage(v string) *CloudListAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListAgentGroupResponseBody) SetRequestId(v string) *CloudListAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListAgentGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListAgentGroupResponseBodyData struct {
	List []*CloudListAgentGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s CloudListAgentGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListAgentGroupResponseBodyData) GetList() []*CloudListAgentGroupResponseBodyDataList {
	return s.List
}

func (s *CloudListAgentGroupResponseBodyData) SetList(v []*CloudListAgentGroupResponseBodyDataList) *CloudListAgentGroupResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListAgentGroupResponseBodyData) Validate() error {
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

type CloudListAgentGroupResponseBodyDataList struct {
	// example:
	//
	// ""
	ChildGnos         *string                                                   `json:"ChildGnos,omitempty" xml:"ChildGnos,omitempty"`
	CtiLinkAgentGroup *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup `json:"CtiLinkAgentGroup,omitempty" xml:"CtiLinkAgentGroup,omitempty" type:"Struct"`
}

func (s CloudListAgentGroupResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListAgentGroupResponseBodyDataList) GetChildGnos() *string {
	return s.ChildGnos
}

func (s *CloudListAgentGroupResponseBodyDataList) GetCtiLinkAgentGroup() *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup {
	return s.CtiLinkAgentGroup
}

func (s *CloudListAgentGroupResponseBodyDataList) SetChildGnos(v string) *CloudListAgentGroupResponseBodyDataList {
	s.ChildGnos = &v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataList) SetCtiLinkAgentGroup(v *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) *CloudListAgentGroupResponseBodyDataList {
	s.CtiLinkAgentGroup = v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataList) Validate() error {
	if s.CtiLinkAgentGroup != nil {
		if err := s.CtiLinkAgentGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup struct {
	// 描述信息
	//
	// example:
	//
	// comment1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
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
	// 外呼组编号
	//
	// example:
	//
	// ZO874912
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 外呼组名称
	//
	// example:
	//
	// gpname1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 外呼组类型 1：外呼总组 2：外呼组
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) String() string {
	return dara.Prettify(s)
}

func (s CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) GoString() string {
	return s.String()
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) GetComment() *string {
	return s.Comment
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) GetGno() *string {
	return s.Gno
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) GetType() *int64 {
	return s.Type
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) SetComment(v string) *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup {
	s.Comment = &v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) SetCreateTime(v string) *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup {
	s.CreateTime = &v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) SetEnterpriseId(v int64) *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) SetGno(v string) *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup {
	s.Gno = &v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) SetGroupName(v string) *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup {
	s.GroupName = &v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) SetType(v int64) *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup {
	s.Type = &v
	return s
}

func (s *CloudListAgentGroupResponseBodyDataListCtiLinkAgentGroup) Validate() error {
	return dara.Validate(s)
}
