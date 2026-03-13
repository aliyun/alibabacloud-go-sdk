// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkspaceSpecs interface {
	dara.Model
	String() string
	GoString() string
	SetProduct(v string) *WorkspaceSpecs
	GetProduct() *string
	SetSpecs(v []*WorkspaceSpec) *WorkspaceSpecs
	GetSpecs() []*WorkspaceSpec
	SetWorkspaceId(v string) *WorkspaceSpecs
	GetWorkspaceId() *string
}

type WorkspaceSpecs struct {
	Product     *string          `json:"Product,omitempty" xml:"Product,omitempty"`
	Specs       []*WorkspaceSpec `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
	WorkspaceId *string          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s WorkspaceSpecs) String() string {
	return dara.Prettify(s)
}

func (s WorkspaceSpecs) GoString() string {
	return s.String()
}

func (s *WorkspaceSpecs) GetProduct() *string {
	return s.Product
}

func (s *WorkspaceSpecs) GetSpecs() []*WorkspaceSpec {
	return s.Specs
}

func (s *WorkspaceSpecs) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *WorkspaceSpecs) SetProduct(v string) *WorkspaceSpecs {
	s.Product = &v
	return s
}

func (s *WorkspaceSpecs) SetSpecs(v []*WorkspaceSpec) *WorkspaceSpecs {
	s.Specs = v
	return s
}

func (s *WorkspaceSpecs) SetWorkspaceId(v string) *WorkspaceSpecs {
	s.WorkspaceId = &v
	return s
}

func (s *WorkspaceSpecs) Validate() error {
	if s.Specs != nil {
		for _, item := range s.Specs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
