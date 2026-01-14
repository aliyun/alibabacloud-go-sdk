// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddAgentResponseBody
	GetCode() *string
	SetMessage(v string) *AddAgentResponseBody
	GetMessage() *string
	SetModel(v *AddAgentResponseBodyModel) *AddAgentResponseBody
	GetModel() *AddAgentResponseBodyModel
	SetRequestId(v string) *AddAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAgentResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *AddAgentResponseBody
	GetTimestamp() *int64
}

type AddAgentResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 坐席ID
	Model *AddAgentResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 82
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AddAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAgentResponseBody) GoString() string {
	return s.String()
}

func (s *AddAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAgentResponseBody) GetModel() *AddAgentResponseBodyModel {
	return s.Model
}

func (s *AddAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAgentResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AddAgentResponseBody) SetAccessDeniedDetail(v string) *AddAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddAgentResponseBody) SetCode(v string) *AddAgentResponseBody {
	s.Code = &v
	return s
}

func (s *AddAgentResponseBody) SetMessage(v string) *AddAgentResponseBody {
	s.Message = &v
	return s
}

func (s *AddAgentResponseBody) SetModel(v *AddAgentResponseBodyModel) *AddAgentResponseBody {
	s.Model = v
	return s
}

func (s *AddAgentResponseBody) SetRequestId(v string) *AddAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAgentResponseBody) SetSuccess(v bool) *AddAgentResponseBody {
	s.Success = &v
	return s
}

func (s *AddAgentResponseBody) SetTimestamp(v int64) *AddAgentResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AddAgentResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddAgentResponseBodyModel struct {
	// 坐席ID
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
}

func (s AddAgentResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AddAgentResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AddAgentResponseBodyModel) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AddAgentResponseBodyModel) SetAgentId(v int64) *AddAgentResponseBodyModel {
	s.AgentId = &v
	return s
}

func (s *AddAgentResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
