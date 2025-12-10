// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeInputSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputId(v string) *GetNodeInputSchemaRequest
	GetInputId() *string
	SetInputIndex(v int32) *GetNodeInputSchemaRequest
	GetInputIndex() *int32
}

type GetNodeInputSchemaRequest struct {
	// example:
	//
	// inputTable
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	// example:
	//
	// 0
	InputIndex *int32 `json:"InputIndex,omitempty" xml:"InputIndex,omitempty"`
}

func (s GetNodeInputSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeInputSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetNodeInputSchemaRequest) GetInputId() *string {
	return s.InputId
}

func (s *GetNodeInputSchemaRequest) GetInputIndex() *int32 {
	return s.InputIndex
}

func (s *GetNodeInputSchemaRequest) SetInputId(v string) *GetNodeInputSchemaRequest {
	s.InputId = &v
	return s
}

func (s *GetNodeInputSchemaRequest) SetInputIndex(v int32) *GetNodeInputSchemaRequest {
	s.InputIndex = &v
	return s
}

func (s *GetNodeInputSchemaRequest) Validate() error {
	return dara.Validate(s)
}
