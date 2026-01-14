// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinAgentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *JoinAgentGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *JoinAgentGroupResponseBody
	GetCode() *string
	SetMessage(v string) *JoinAgentGroupResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *JoinAgentGroupResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *JoinAgentGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *JoinAgentGroupResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *JoinAgentGroupResponseBody
	GetTimestamp() *int64
}

type JoinAgentGroupResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
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
	// 30
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s JoinAgentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinAgentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *JoinAgentGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *JoinAgentGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *JoinAgentGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *JoinAgentGroupResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *JoinAgentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinAgentGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *JoinAgentGroupResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *JoinAgentGroupResponseBody) SetAccessDeniedDetail(v string) *JoinAgentGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *JoinAgentGroupResponseBody) SetCode(v string) *JoinAgentGroupResponseBody {
	s.Code = &v
	return s
}

func (s *JoinAgentGroupResponseBody) SetMessage(v string) *JoinAgentGroupResponseBody {
	s.Message = &v
	return s
}

func (s *JoinAgentGroupResponseBody) SetModel(v map[string]interface{}) *JoinAgentGroupResponseBody {
	s.Model = v
	return s
}

func (s *JoinAgentGroupResponseBody) SetRequestId(v string) *JoinAgentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinAgentGroupResponseBody) SetSuccess(v bool) *JoinAgentGroupResponseBody {
	s.Success = &v
	return s
}

func (s *JoinAgentGroupResponseBody) SetTimestamp(v int64) *JoinAgentGroupResponseBody {
	s.Timestamp = &v
	return s
}

func (s *JoinAgentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
