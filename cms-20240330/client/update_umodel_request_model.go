// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonSchemaRef(v []*UpdateUmodelRequestCommonSchemaRef) *UpdateUmodelRequest
	GetCommonSchemaRef() []*UpdateUmodelRequestCommonSchemaRef
	SetDescription(v string) *UpdateUmodelRequest
	GetDescription() *string
}

type UpdateUmodelRequest struct {
	CommonSchemaRef []*UpdateUmodelRequestCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUmodelRequest) GoString() string {
	return s.String()
}

func (s *UpdateUmodelRequest) GetCommonSchemaRef() []*UpdateUmodelRequestCommonSchemaRef {
	return s.CommonSchemaRef
}

func (s *UpdateUmodelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUmodelRequest) SetCommonSchemaRef(v []*UpdateUmodelRequestCommonSchemaRef) *UpdateUmodelRequest {
	s.CommonSchemaRef = v
	return s
}

func (s *UpdateUmodelRequest) SetDescription(v string) *UpdateUmodelRequest {
	s.Description = &v
	return s
}

func (s *UpdateUmodelRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateUmodelRequestCommonSchemaRef struct {
	// example:
	//
	// test-bmp-123123
	Group *string   `json:"group,omitempty" xml:"group,omitempty"`
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2.5.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateUmodelRequestCommonSchemaRef) String() string {
	return dara.Prettify(s)
}

func (s UpdateUmodelRequestCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *UpdateUmodelRequestCommonSchemaRef) GetGroup() *string {
	return s.Group
}

func (s *UpdateUmodelRequestCommonSchemaRef) GetItems() []*string {
	return s.Items
}

func (s *UpdateUmodelRequestCommonSchemaRef) GetVersion() *string {
	return s.Version
}

func (s *UpdateUmodelRequestCommonSchemaRef) SetGroup(v string) *UpdateUmodelRequestCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *UpdateUmodelRequestCommonSchemaRef) SetItems(v []*string) *UpdateUmodelRequestCommonSchemaRef {
	s.Items = v
	return s
}

func (s *UpdateUmodelRequestCommonSchemaRef) SetVersion(v string) *UpdateUmodelRequestCommonSchemaRef {
	s.Version = &v
	return s
}

func (s *UpdateUmodelRequestCommonSchemaRef) Validate() error {
	return dara.Validate(s)
}
