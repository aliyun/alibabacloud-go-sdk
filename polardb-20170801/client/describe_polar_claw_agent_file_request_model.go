// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DescribePolarClawAgentFileRequest
	GetAgentId() *string
	SetApplicationId(v string) *DescribePolarClawAgentFileRequest
	GetApplicationId() *string
	SetFileName(v string) *DescribePolarClawAgentFileRequest
	GetFileName() *string
}

type DescribePolarClawAgentFileRequest struct {
	// Agent ID
	//
	// This parameter is required.
	//
	// example:
	//
	// main
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// SOUL.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s DescribePolarClawAgentFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentFileRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentFileRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribePolarClawAgentFileRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawAgentFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribePolarClawAgentFileRequest) SetAgentId(v string) *DescribePolarClawAgentFileRequest {
	s.AgentId = &v
	return s
}

func (s *DescribePolarClawAgentFileRequest) SetApplicationId(v string) *DescribePolarClawAgentFileRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawAgentFileRequest) SetFileName(v string) *DescribePolarClawAgentFileRequest {
	s.FileName = &v
	return s
}

func (s *DescribePolarClawAgentFileRequest) Validate() error {
	return dara.Validate(s)
}
