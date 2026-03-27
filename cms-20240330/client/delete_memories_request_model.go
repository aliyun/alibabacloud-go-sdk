// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeleteMemoriesRequest
	GetAgentId() *string
	SetAppId(v string) *DeleteMemoriesRequest
	GetAppId() *string
	SetRunId(v string) *DeleteMemoriesRequest
	GetRunId() *string
	SetUserId(v string) *DeleteMemoriesRequest
	GetUserId() *string
}

type DeleteMemoriesRequest struct {
	// example:
	//
	// 952730733889060865
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// 98ea19fe-128b-4841-b318-0359bec3c65d
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// jr-dd7c645fd6fe50d4
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// test_user_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemoriesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteMemoriesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMemoriesRequest) GetRunId() *string {
	return s.RunId
}

func (s *DeleteMemoriesRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteMemoriesRequest) SetAgentId(v string) *DeleteMemoriesRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteMemoriesRequest) SetAppId(v string) *DeleteMemoriesRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMemoriesRequest) SetRunId(v string) *DeleteMemoriesRequest {
	s.RunId = &v
	return s
}

func (s *DeleteMemoriesRequest) SetUserId(v string) *DeleteMemoriesRequest {
	s.UserId = &v
	return s
}

func (s *DeleteMemoriesRequest) Validate() error {
	return dara.Validate(s)
}
