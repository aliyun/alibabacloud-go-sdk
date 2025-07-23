// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonSchemaRef(v []*CreateUmodelRequestCommonSchemaRef) *CreateUmodelRequest
	GetCommonSchemaRef() []*CreateUmodelRequestCommonSchemaRef
	SetDescription(v string) *CreateUmodelRequest
	GetDescription() *string
}

type CreateUmodelRequest struct {
	CommonSchemaRef []*CreateUmodelRequestCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUmodelRequest) GoString() string {
	return s.String()
}

func (s *CreateUmodelRequest) GetCommonSchemaRef() []*CreateUmodelRequestCommonSchemaRef {
	return s.CommonSchemaRef
}

func (s *CreateUmodelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUmodelRequest) SetCommonSchemaRef(v []*CreateUmodelRequestCommonSchemaRef) *CreateUmodelRequest {
	s.CommonSchemaRef = v
	return s
}

func (s *CreateUmodelRequest) SetDescription(v string) *CreateUmodelRequest {
	s.Description = &v
	return s
}

func (s *CreateUmodelRequest) Validate() error {
	return dara.Validate(s)
}

type CreateUmodelRequestCommonSchemaRef struct {
	// example:
	//
	// test-job-123
	Group *string   `json:"group,omitempty" xml:"group,omitempty"`
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s CreateUmodelRequestCommonSchemaRef) String() string {
	return dara.Prettify(s)
}

func (s CreateUmodelRequestCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *CreateUmodelRequestCommonSchemaRef) GetGroup() *string {
	return s.Group
}

func (s *CreateUmodelRequestCommonSchemaRef) GetItems() []*string {
	return s.Items
}

func (s *CreateUmodelRequestCommonSchemaRef) GetVersion() *string {
	return s.Version
}

func (s *CreateUmodelRequestCommonSchemaRef) SetGroup(v string) *CreateUmodelRequestCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *CreateUmodelRequestCommonSchemaRef) SetItems(v []*string) *CreateUmodelRequestCommonSchemaRef {
	s.Items = v
	return s
}

func (s *CreateUmodelRequestCommonSchemaRef) SetVersion(v string) *CreateUmodelRequestCommonSchemaRef {
	s.Version = &v
	return s
}

func (s *CreateUmodelRequestCommonSchemaRef) Validate() error {
	return dara.Validate(s)
}
