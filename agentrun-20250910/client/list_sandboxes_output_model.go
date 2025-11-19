// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesOutput interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*Sandbox) *ListSandboxesOutput
	GetItems() []*Sandbox
	SetNextToken(v string) *ListSandboxesOutput
	GetNextToken() *string
}

type ListSandboxesOutput struct {
	// This parameter is required.
	Items     []*Sandbox `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextToken *string    `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListSandboxesOutput) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesOutput) GoString() string {
	return s.String()
}

func (s *ListSandboxesOutput) GetItems() []*Sandbox {
	return s.Items
}

func (s *ListSandboxesOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxesOutput) SetItems(v []*Sandbox) *ListSandboxesOutput {
	s.Items = v
	return s
}

func (s *ListSandboxesOutput) SetNextToken(v string) *ListSandboxesOutput {
	s.NextToken = &v
	return s
}

func (s *ListSandboxesOutput) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
