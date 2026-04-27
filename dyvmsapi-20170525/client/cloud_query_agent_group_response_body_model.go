// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryAgentGroupResponseBody
	GetCode() *string
	SetData(v *CloudQueryAgentGroupResponseBodyData) *CloudQueryAgentGroupResponseBody
	GetData() *CloudQueryAgentGroupResponseBodyData
	SetMessage(v string) *CloudQueryAgentGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryAgentGroupResponseBody
	GetRequestId() *string
}

type CloudQueryAgentGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryAgentGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryAgentGroupResponseBody) GetData() *CloudQueryAgentGroupResponseBodyData {
	return s.Data
}

func (s *CloudQueryAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryAgentGroupResponseBody) SetAccessDeniedDetail(v string) *CloudQueryAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBody) SetCode(v string) *CloudQueryAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBody) SetData(v *CloudQueryAgentGroupResponseBodyData) *CloudQueryAgentGroupResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryAgentGroupResponseBody) SetMessage(v string) *CloudQueryAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBody) SetRequestId(v string) *CloudQueryAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryAgentGroupResponseBodyData struct {
	// 描述信息
	//
	// example:
	//
	// description
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
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 外呼组编号
	//
	// example:
	//
	// 233
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 外呼组名称
	//
	// example:
	//
	// name1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 外呼组id
	//
	// example:
	//
	// 33
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 外呼组类型 1：外呼总组 2：外呼组
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CloudQueryAgentGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentGroupResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *CloudQueryAgentGroupResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudQueryAgentGroupResponseBodyData) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryAgentGroupResponseBodyData) GetGno() *string {
	return s.Gno
}

func (s *CloudQueryAgentGroupResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *CloudQueryAgentGroupResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CloudQueryAgentGroupResponseBodyData) GetType() *int64 {
	return s.Type
}

func (s *CloudQueryAgentGroupResponseBodyData) SetComment(v string) *CloudQueryAgentGroupResponseBodyData {
	s.Comment = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBodyData) SetCreateTime(v string) *CloudQueryAgentGroupResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBodyData) SetEnterpriseId(v int64) *CloudQueryAgentGroupResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBodyData) SetGno(v string) *CloudQueryAgentGroupResponseBodyData {
	s.Gno = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBodyData) SetGroupName(v string) *CloudQueryAgentGroupResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBodyData) SetId(v int64) *CloudQueryAgentGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBodyData) SetType(v int64) *CloudQueryAgentGroupResponseBodyData {
	s.Type = &v
	return s
}

func (s *CloudQueryAgentGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
