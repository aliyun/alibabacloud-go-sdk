// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentImportNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AgentImportNumberResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AgentImportNumberResponseBody
	GetCode() *string
	SetMessage(v string) *AgentImportNumberResponseBody
	GetMessage() *string
	SetModel(v *AgentImportNumberResponseBodyModel) *AgentImportNumberResponseBody
	GetModel() *AgentImportNumberResponseBodyModel
	SetRequestId(v string) *AgentImportNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AgentImportNumberResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *AgentImportNumberResponseBody
	GetTimestamp() *int64
}

type AgentImportNumberResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AgentImportNumberResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 93
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AgentImportNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AgentImportNumberResponseBody) GoString() string {
	return s.String()
}

func (s *AgentImportNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AgentImportNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AgentImportNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AgentImportNumberResponseBody) GetModel() *AgentImportNumberResponseBodyModel {
	return s.Model
}

func (s *AgentImportNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentImportNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AgentImportNumberResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AgentImportNumberResponseBody) SetAccessDeniedDetail(v string) *AgentImportNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AgentImportNumberResponseBody) SetCode(v string) *AgentImportNumberResponseBody {
	s.Code = &v
	return s
}

func (s *AgentImportNumberResponseBody) SetMessage(v string) *AgentImportNumberResponseBody {
	s.Message = &v
	return s
}

func (s *AgentImportNumberResponseBody) SetModel(v *AgentImportNumberResponseBodyModel) *AgentImportNumberResponseBody {
	s.Model = v
	return s
}

func (s *AgentImportNumberResponseBody) SetRequestId(v string) *AgentImportNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgentImportNumberResponseBody) SetSuccess(v bool) *AgentImportNumberResponseBody {
	s.Success = &v
	return s
}

func (s *AgentImportNumberResponseBody) SetTimestamp(v int64) *AgentImportNumberResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AgentImportNumberResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentImportNumberResponseBodyModel struct {
	// 批次ID
	//
	// example:
	//
	// 示例值示例值示例值
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 外呼编号ID
	//
	// example:
	//
	// 80
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AgentImportNumberResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AgentImportNumberResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AgentImportNumberResponseBodyModel) GetBatchId() *string {
	return s.BatchId
}

func (s *AgentImportNumberResponseBodyModel) GetId() *int64 {
	return s.Id
}

func (s *AgentImportNumberResponseBodyModel) SetBatchId(v string) *AgentImportNumberResponseBodyModel {
	s.BatchId = &v
	return s
}

func (s *AgentImportNumberResponseBodyModel) SetId(v int64) *AgentImportNumberResponseBodyModel {
	s.Id = &v
	return s
}

func (s *AgentImportNumberResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
