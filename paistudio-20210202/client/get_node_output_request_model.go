// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputIndex(v string) *GetNodeOutputRequest
	GetOutputIndex() *string
}

type GetNodeOutputRequest struct {
	// example:
	//
	// 0
	OutputIndex *string `json:"OutputIndex,omitempty" xml:"OutputIndex,omitempty"`
}

func (s GetNodeOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeOutputRequest) GoString() string {
	return s.String()
}

func (s *GetNodeOutputRequest) GetOutputIndex() *string {
	return s.OutputIndex
}

func (s *GetNodeOutputRequest) SetOutputIndex(v string) *GetNodeOutputRequest {
	s.OutputIndex = &v
	return s
}

func (s *GetNodeOutputRequest) Validate() error {
	return dara.Validate(s)
}
