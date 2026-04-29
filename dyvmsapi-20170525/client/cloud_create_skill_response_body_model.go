// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateSkillResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateSkillResponseBody
	GetCode() *string
	SetData(v *CloudCreateSkillResponseBodyData) *CloudCreateSkillResponseBody
	GetData() *CloudCreateSkillResponseBodyData
	SetMessage(v string) *CloudCreateSkillResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateSkillResponseBody
	GetRequestId() *string
}

type CloudCreateSkillResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateSkillResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateSkillResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateSkillResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateSkillResponseBody) GetData() *CloudCreateSkillResponseBodyData {
	return s.Data
}

func (s *CloudCreateSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateSkillResponseBody) SetAccessDeniedDetail(v string) *CloudCreateSkillResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateSkillResponseBody) SetCode(v string) *CloudCreateSkillResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateSkillResponseBody) SetData(v *CloudCreateSkillResponseBodyData) *CloudCreateSkillResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateSkillResponseBody) SetMessage(v string) *CloudCreateSkillResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateSkillResponseBody) SetRequestId(v string) *CloudCreateSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateSkillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateSkillResponseBodyData struct {
	// 描述
	//
	// example:
	//
	// Comment3
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 创建时间,精确到秒
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
	// 技能id
	//
	// example:
	//
	// 49
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 技能名称
	//
	// example:
	//
	// skillName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CloudCreateSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateSkillResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *CloudCreateSkillResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateSkillResponseBodyData) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateSkillResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *CloudCreateSkillResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CloudCreateSkillResponseBodyData) SetComment(v string) *CloudCreateSkillResponseBodyData {
	s.Comment = &v
	return s
}

func (s *CloudCreateSkillResponseBodyData) SetCreateTime(v string) *CloudCreateSkillResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateSkillResponseBodyData) SetEnterpriseId(v int64) *CloudCreateSkillResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateSkillResponseBodyData) SetId(v int64) *CloudCreateSkillResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudCreateSkillResponseBodyData) SetName(v string) *CloudCreateSkillResponseBodyData {
	s.Name = &v
	return s
}

func (s *CloudCreateSkillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
