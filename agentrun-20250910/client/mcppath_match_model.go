// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMCPPathMatch interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *MCPPathMatch
	GetType() *string
	SetValue(v string) *MCPPathMatch
	GetValue() *string
}

type MCPPathMatch struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s MCPPathMatch) String() string {
	return dara.Prettify(s)
}

func (s MCPPathMatch) GoString() string {
	return s.String()
}

func (s *MCPPathMatch) GetType() *string {
	return s.Type
}

func (s *MCPPathMatch) GetValue() *string {
	return s.Value
}

func (s *MCPPathMatch) SetType(v string) *MCPPathMatch {
	s.Type = &v
	return s
}

func (s *MCPPathMatch) SetValue(v string) *MCPPathMatch {
	s.Value = &v
	return s
}

func (s *MCPPathMatch) Validate() error {
	return dara.Validate(s)
}
