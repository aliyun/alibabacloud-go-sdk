// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlgorithmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmDescription(v string) *CreateAlgorithmRequest
	GetAlgorithmDescription() *string
	SetAlgorithmName(v string) *CreateAlgorithmRequest
	GetAlgorithmName() *string
	SetDisplayName(v string) *CreateAlgorithmRequest
	GetDisplayName() *string
	SetWorkspaceId(v string) *CreateAlgorithmRequest
	GetWorkspaceId() *string
}

type CreateAlgorithmRequest struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// llm_training
	AlgorithmName *string `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateAlgorithmRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmRequest) GetAlgorithmDescription() *string {
	return s.AlgorithmDescription
}

func (s *CreateAlgorithmRequest) GetAlgorithmName() *string {
	return s.AlgorithmName
}

func (s *CreateAlgorithmRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAlgorithmRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateAlgorithmRequest) SetAlgorithmDescription(v string) *CreateAlgorithmRequest {
	s.AlgorithmDescription = &v
	return s
}

func (s *CreateAlgorithmRequest) SetAlgorithmName(v string) *CreateAlgorithmRequest {
	s.AlgorithmName = &v
	return s
}

func (s *CreateAlgorithmRequest) SetDisplayName(v string) *CreateAlgorithmRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAlgorithmRequest) SetWorkspaceId(v string) *CreateAlgorithmRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateAlgorithmRequest) Validate() error {
	return dara.Validate(s)
}
