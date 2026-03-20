// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetMemoryResponseBody
	GetAgentId() *string
	SetAppId(v string) *GetMemoryResponseBody
	GetAppId() *string
	SetCreatedAt(v string) *GetMemoryResponseBody
	GetCreatedAt() *string
	SetExpirationDate(v string) *GetMemoryResponseBody
	GetExpirationDate() *string
	SetId(v string) *GetMemoryResponseBody
	GetId() *string
	SetImmutable(v string) *GetMemoryResponseBody
	GetImmutable() *string
	SetMemory(v string) *GetMemoryResponseBody
	GetMemory() *string
	SetMetadata(v string) *GetMemoryResponseBody
	GetMetadata() *string
	SetOrganization(v string) *GetMemoryResponseBody
	GetOrganization() *string
	SetOwner(v string) *GetMemoryResponseBody
	GetOwner() *string
	SetRequestId(v string) *GetMemoryResponseBody
	GetRequestId() *string
	SetRunId(v string) *GetMemoryResponseBody
	GetRunId() *string
	SetUpdatedAt(v string) *GetMemoryResponseBody
	GetUpdatedAt() *string
	SetUserId(v string) *GetMemoryResponseBody
	GetUserId() *string
}

type GetMemoryResponseBody struct {
	// example:
	//
	// test_user_001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// test_user_001
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 1751595283143
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 1751595283143
	ExpirationDate *string `json:"expirationDate,omitempty" xml:"expirationDate,omitempty"`
	// example:
	//
	// 019ca1e5-7307-7d50-b943-5e628326a8ed
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test
	Immutable *string `json:"immutable,omitempty" xml:"immutable,omitempty"`
	// example:
	//
	// My name is Zhang San and I live in Hangzhou.
	Memory *string `json:"memory,omitempty" xml:"memory,omitempty"`
	// example:
	//
	// {"sessionId":"test_session_001"}
	Metadata *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// test
	Organization *string `json:"organization,omitempty" xml:"organization,omitempty"`
	// example:
	//
	// test
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// test_user_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// 1744428159434
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// test_user_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *GetMemoryResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetMemoryResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMemoryResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *GetMemoryResponseBody) GetId() *string {
	return s.Id
}

func (s *GetMemoryResponseBody) GetImmutable() *string {
	return s.Immutable
}

func (s *GetMemoryResponseBody) GetMemory() *string {
	return s.Memory
}

func (s *GetMemoryResponseBody) GetMetadata() *string {
	return s.Metadata
}

func (s *GetMemoryResponseBody) GetOrganization() *string {
	return s.Organization
}

func (s *GetMemoryResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *GetMemoryResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMemoryResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetMemoryResponseBody) SetAgentId(v string) *GetMemoryResponseBody {
	s.AgentId = &v
	return s
}

func (s *GetMemoryResponseBody) SetAppId(v string) *GetMemoryResponseBody {
	s.AppId = &v
	return s
}

func (s *GetMemoryResponseBody) SetCreatedAt(v string) *GetMemoryResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetMemoryResponseBody) SetExpirationDate(v string) *GetMemoryResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *GetMemoryResponseBody) SetId(v string) *GetMemoryResponseBody {
	s.Id = &v
	return s
}

func (s *GetMemoryResponseBody) SetImmutable(v string) *GetMemoryResponseBody {
	s.Immutable = &v
	return s
}

func (s *GetMemoryResponseBody) SetMemory(v string) *GetMemoryResponseBody {
	s.Memory = &v
	return s
}

func (s *GetMemoryResponseBody) SetMetadata(v string) *GetMemoryResponseBody {
	s.Metadata = &v
	return s
}

func (s *GetMemoryResponseBody) SetOrganization(v string) *GetMemoryResponseBody {
	s.Organization = &v
	return s
}

func (s *GetMemoryResponseBody) SetOwner(v string) *GetMemoryResponseBody {
	s.Owner = &v
	return s
}

func (s *GetMemoryResponseBody) SetRequestId(v string) *GetMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryResponseBody) SetRunId(v string) *GetMemoryResponseBody {
	s.RunId = &v
	return s
}

func (s *GetMemoryResponseBody) SetUpdatedAt(v string) *GetMemoryResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetMemoryResponseBody) SetUserId(v string) *GetMemoryResponseBody {
	s.UserId = &v
	return s
}

func (s *GetMemoryResponseBody) Validate() error {
	return dara.Validate(s)
}
