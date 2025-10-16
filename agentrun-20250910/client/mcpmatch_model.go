// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPMatch interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v *MCPPathMatch) *MCPMatch
	GetPath() *MCPPathMatch
}

type MCPMatch struct {
	Path *MCPPathMatch `json:"path,omitempty" xml:"path,omitempty"`
}

func (s MCPMatch) String() string {
	return dara.Prettify(s)
}

func (s MCPMatch) GoString() string {
	return s.String()
}

func (s *MCPMatch) GetPath() *MCPPathMatch {
	return s.Path
}

func (s *MCPMatch) SetPath(v *MCPPathMatch) *MCPMatch {
	s.Path = v
	return s
}

func (s *MCPMatch) Validate() error {
	if s.Path != nil {
		if err := s.Path.Validate(); err != nil {
			return err
		}
	}
	return nil
}
