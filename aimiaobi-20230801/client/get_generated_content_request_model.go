// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGeneratedContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetGeneratedContentRequest
	GetAgentKey() *string
	SetId(v int64) *GetGeneratedContentRequest
	GetId() *int64
}

type GetGeneratedContentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetGeneratedContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetGeneratedContentRequest) GetId() *int64 {
	return s.Id
}

func (s *GetGeneratedContentRequest) SetAgentKey(v string) *GetGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *GetGeneratedContentRequest) SetId(v int64) *GetGeneratedContentRequest {
	s.Id = &v
	return s
}

func (s *GetGeneratedContentRequest) Validate() error {
	return dara.Validate(s)
}
