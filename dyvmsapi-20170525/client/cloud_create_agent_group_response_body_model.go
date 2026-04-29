// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateAgentGroupResponseBody
	GetCode() *string
	SetData(v *CloudCreateAgentGroupResponseBodyData) *CloudCreateAgentGroupResponseBody
	GetData() *CloudCreateAgentGroupResponseBodyData
	SetMessage(v string) *CloudCreateAgentGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateAgentGroupResponseBody
	GetRequestId() *string
}

type CloudCreateAgentGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateAgentGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateAgentGroupResponseBody) GetData() *CloudCreateAgentGroupResponseBodyData {
	return s.Data
}

func (s *CloudCreateAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateAgentGroupResponseBody) SetAccessDeniedDetail(v string) *CloudCreateAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBody) SetCode(v string) *CloudCreateAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBody) SetData(v *CloudCreateAgentGroupResponseBodyData) *CloudCreateAgentGroupResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateAgentGroupResponseBody) SetMessage(v string) *CloudCreateAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBody) SetRequestId(v string) *CloudCreateAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateAgentGroupResponseBodyData struct {
	// 描述信息
	//
	// example:
	//
	// comment3
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
	// WH12
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

func (s CloudCreateAgentGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAgentGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateAgentGroupResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *CloudCreateAgentGroupResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateAgentGroupResponseBodyData) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateAgentGroupResponseBodyData) GetGno() *string {
	return s.Gno
}

func (s *CloudCreateAgentGroupResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *CloudCreateAgentGroupResponseBodyData) GetType() *int64 {
	return s.Type
}

func (s *CloudCreateAgentGroupResponseBodyData) SetComment(v string) *CloudCreateAgentGroupResponseBodyData {
	s.Comment = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBodyData) SetCreateTime(v string) *CloudCreateAgentGroupResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBodyData) SetEnterpriseId(v int64) *CloudCreateAgentGroupResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBodyData) SetGno(v string) *CloudCreateAgentGroupResponseBodyData {
	s.Gno = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBodyData) SetGroupName(v string) *CloudCreateAgentGroupResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBodyData) SetType(v int64) *CloudCreateAgentGroupResponseBodyData {
	s.Type = &v
	return s
}

func (s *CloudCreateAgentGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
