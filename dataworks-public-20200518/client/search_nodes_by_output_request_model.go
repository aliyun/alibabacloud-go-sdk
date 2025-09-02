// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchNodesByOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputs(v string) *SearchNodesByOutputRequest
	GetOutputs() *string
	SetProjectEnv(v string) *SearchNodesByOutputRequest
	GetProjectEnv() *string
}

type SearchNodesByOutputRequest struct {
	// The output names of the node. If you specify multiple output names, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// aaaaa.1231412_out,bbbb.12313123_out
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The environment of Operation Center. Valid values: PROD and DEV.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s SearchNodesByOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchNodesByOutputRequest) GoString() string {
	return s.String()
}

func (s *SearchNodesByOutputRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *SearchNodesByOutputRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *SearchNodesByOutputRequest) SetOutputs(v string) *SearchNodesByOutputRequest {
	s.Outputs = &v
	return s
}

func (s *SearchNodesByOutputRequest) SetProjectEnv(v string) *SearchNodesByOutputRequest {
	s.ProjectEnv = &v
	return s
}

func (s *SearchNodesByOutputRequest) Validate() error {
	return dara.Validate(s)
}
