// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuitAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QuitAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QuitAgentGroupResponseBody
	GetCode() *string
	SetMessage(v string) *QuitAgentGroupResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *QuitAgentGroupResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *QuitAgentGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuitAgentGroupResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *QuitAgentGroupResponseBody
	GetTimestamp() *int64
}

type QuitAgentGroupResponseBody struct {
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
	// 示例值示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 15
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s QuitAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuitAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QuitAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QuitAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuitAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuitAgentGroupResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *QuitAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuitAgentGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuitAgentGroupResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *QuitAgentGroupResponseBody) SetAccessDeniedDetail(v string) *QuitAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuitAgentGroupResponseBody) SetCode(v string) *QuitAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QuitAgentGroupResponseBody) SetMessage(v string) *QuitAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *QuitAgentGroupResponseBody) SetModel(v map[string]interface{}) *QuitAgentGroupResponseBody {
	s.Model = v
	return s
}

func (s *QuitAgentGroupResponseBody) SetRequestId(v string) *QuitAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuitAgentGroupResponseBody) SetSuccess(v bool) *QuitAgentGroupResponseBody {
	s.Success = &v
	return s
}

func (s *QuitAgentGroupResponseBody) SetTimestamp(v int64) *QuitAgentGroupResponseBody {
	s.Timestamp = &v
	return s
}

func (s *QuitAgentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
