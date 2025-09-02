// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesByOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputs(v string) *ListNodesByOutputRequest
	GetOutputs() *string
	SetProjectEnv(v string) *ListNodesByOutputRequest
	GetProjectEnv() *string
}

type ListNodesByOutputRequest struct {
	// The output name of the node. You can specify multiple output names. Separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// aaaaa.1231412_out,bbbb.12313123_out
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The environment of Operation Center. Valid values: PROD and DEV. The value PROD indicates the production environment, and the value DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s ListNodesByOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesByOutputRequest) GoString() string {
	return s.String()
}

func (s *ListNodesByOutputRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *ListNodesByOutputRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListNodesByOutputRequest) SetOutputs(v string) *ListNodesByOutputRequest {
	s.Outputs = &v
	return s
}

func (s *ListNodesByOutputRequest) SetProjectEnv(v string) *ListNodesByOutputRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListNodesByOutputRequest) Validate() error {
	return dara.Validate(s)
}
