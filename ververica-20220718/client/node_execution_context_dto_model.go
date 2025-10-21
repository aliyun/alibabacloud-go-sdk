// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeExecutionContextDTO interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v map[string]*string) *NodeExecutionContextDTO
	GetContext() map[string]*string
}

type NodeExecutionContextDTO struct {
	Context map[string]*string `json:"context,omitempty" xml:"context,omitempty"`
}

func (s NodeExecutionContextDTO) String() string {
	return dara.Prettify(s)
}

func (s NodeExecutionContextDTO) GoString() string {
	return s.String()
}

func (s *NodeExecutionContextDTO) GetContext() map[string]*string {
	return s.Context
}

func (s *NodeExecutionContextDTO) SetContext(v map[string]*string) *NodeExecutionContextDTO {
	s.Context = v
	return s
}

func (s *NodeExecutionContextDTO) Validate() error {
	return dara.Validate(s)
}
