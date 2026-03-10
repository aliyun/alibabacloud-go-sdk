// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetMemoriesRequest
	GetAgentId() *string
	SetAppId(v string) *GetMemoriesRequest
	GetAppId() *string
	SetPage(v int32) *GetMemoriesRequest
	GetPage() *int32
	SetPageSize(v int32) *GetMemoriesRequest
	GetPageSize() *int32
	SetRunId(v string) *GetMemoriesRequest
	GetRunId() *string
	SetUserId(v string) *GetMemoriesRequest
	GetUserId() *string
}

type GetMemoriesRequest struct {
	// example:
	//
	// 952730652285943809
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 150130323
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 1000
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// jr-80ded1d6953c64ea
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// test_user_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoriesRequest) GoString() string {
	return s.String()
}

func (s *GetMemoriesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *GetMemoriesRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMemoriesRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetMemoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMemoriesRequest) GetRunId() *string {
	return s.RunId
}

func (s *GetMemoriesRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetMemoriesRequest) SetAgentId(v string) *GetMemoriesRequest {
	s.AgentId = &v
	return s
}

func (s *GetMemoriesRequest) SetAppId(v string) *GetMemoriesRequest {
	s.AppId = &v
	return s
}

func (s *GetMemoriesRequest) SetPage(v int32) *GetMemoriesRequest {
	s.Page = &v
	return s
}

func (s *GetMemoriesRequest) SetPageSize(v int32) *GetMemoriesRequest {
	s.PageSize = &v
	return s
}

func (s *GetMemoriesRequest) SetRunId(v string) *GetMemoriesRequest {
	s.RunId = &v
	return s
}

func (s *GetMemoriesRequest) SetUserId(v string) *GetMemoriesRequest {
	s.UserId = &v
	return s
}

func (s *GetMemoriesRequest) Validate() error {
	return dara.Validate(s)
}
