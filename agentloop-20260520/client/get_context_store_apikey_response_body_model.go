// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreAPIKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetContextStoreAPIKeyResponseBody
	GetAgentSpace() *string
	SetApiKey(v string) *GetContextStoreAPIKeyResponseBody
	GetApiKey() *string
	SetContextStoreName(v string) *GetContextStoreAPIKeyResponseBody
	GetContextStoreName() *string
	SetCreateTime(v string) *GetContextStoreAPIKeyResponseBody
	GetCreateTime() *string
	SetName(v string) *GetContextStoreAPIKeyResponseBody
	GetName() *string
	SetRegionId(v string) *GetContextStoreAPIKeyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetContextStoreAPIKeyResponseBody
	GetRequestId() *string
}

type GetContextStoreAPIKeyResponseBody struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// sk-abcd****
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// my-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// my-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetContextStoreAPIKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreAPIKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextStoreAPIKeyResponseBody) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetContextStoreAPIKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetContextStoreAPIKeyResponseBody) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *GetContextStoreAPIKeyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetContextStoreAPIKeyResponseBody) GetName() *string {
	return s.Name
}

func (s *GetContextStoreAPIKeyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetContextStoreAPIKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextStoreAPIKeyResponseBody) SetAgentSpace(v string) *GetContextStoreAPIKeyResponseBody {
	s.AgentSpace = &v
	return s
}

func (s *GetContextStoreAPIKeyResponseBody) SetApiKey(v string) *GetContextStoreAPIKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *GetContextStoreAPIKeyResponseBody) SetContextStoreName(v string) *GetContextStoreAPIKeyResponseBody {
	s.ContextStoreName = &v
	return s
}

func (s *GetContextStoreAPIKeyResponseBody) SetCreateTime(v string) *GetContextStoreAPIKeyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetContextStoreAPIKeyResponseBody) SetName(v string) *GetContextStoreAPIKeyResponseBody {
	s.Name = &v
	return s
}

func (s *GetContextStoreAPIKeyResponseBody) SetRegionId(v string) *GetContextStoreAPIKeyResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetContextStoreAPIKeyResponseBody) SetRequestId(v string) *GetContextStoreAPIKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextStoreAPIKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
