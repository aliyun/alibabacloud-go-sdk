// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddAgentGroupResponseBody
	GetCode() *string
	SetMessage(v string) *AddAgentGroupResponseBody
	GetMessage() *string
	SetModel(v *AddAgentGroupResponseBodyModel) *AddAgentGroupResponseBody
	GetModel() *AddAgentGroupResponseBodyModel
	SetRequestId(v string) *AddAgentGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAgentGroupResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *AddAgentGroupResponseBody
	GetTimestamp() *int64
}

type AddAgentGroupResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AddAgentGroupResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 40
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AddAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAgentGroupResponseBody) GetModel() *AddAgentGroupResponseBodyModel {
	return s.Model
}

func (s *AddAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAgentGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAgentGroupResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AddAgentGroupResponseBody) SetAccessDeniedDetail(v string) *AddAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddAgentGroupResponseBody) SetCode(v string) *AddAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddAgentGroupResponseBody) SetMessage(v string) *AddAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddAgentGroupResponseBody) SetModel(v *AddAgentGroupResponseBodyModel) *AddAgentGroupResponseBody {
	s.Model = v
	return s
}

func (s *AddAgentGroupResponseBody) SetRequestId(v string) *AddAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAgentGroupResponseBody) SetSuccess(v bool) *AddAgentGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AddAgentGroupResponseBody) SetTimestamp(v int64) *AddAgentGroupResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AddAgentGroupResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAgentGroupResponseBodyModel struct {
	// 坐席组ID
	//
	// example:
	//
	// 48
	AgentGroupId *int64 `json:"AgentGroupId,omitempty" xml:"AgentGroupId,omitempty"`
	// 坐席组名称
	//
	// example:
	//
	// 示例值示例值示例值
	AgentGroupName *string `json:"AgentGroupName,omitempty" xml:"AgentGroupName,omitempty"`
	// 创建坐席组的时间
	//
	// example:
	//
	// 示例值
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s AddAgentGroupResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AddAgentGroupResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AddAgentGroupResponseBodyModel) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *AddAgentGroupResponseBodyModel) GetAgentGroupName() *string {
	return s.AgentGroupName
}

func (s *AddAgentGroupResponseBodyModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddAgentGroupResponseBodyModel) SetAgentGroupId(v int64) *AddAgentGroupResponseBodyModel {
	s.AgentGroupId = &v
	return s
}

func (s *AddAgentGroupResponseBodyModel) SetAgentGroupName(v string) *AddAgentGroupResponseBodyModel {
	s.AgentGroupName = &v
	return s
}

func (s *AddAgentGroupResponseBodyModel) SetCreateTime(v string) *AddAgentGroupResponseBodyModel {
	s.CreateTime = &v
	return s
}

func (s *AddAgentGroupResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
