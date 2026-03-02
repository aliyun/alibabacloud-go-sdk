// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronProductScope interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *NeuronProductScope
	GetDescription() *string
	SetGmtCreate(v string) *NeuronProductScope
	GetGmtCreate() *string
	SetGmtModified(v string) *NeuronProductScope
	GetGmtModified() *string
	SetName(v string) *NeuronProductScope
	GetName() *string
	SetProductId(v int64) *NeuronProductScope
	GetProductId() *int64
	SetScope(v string) *NeuronProductScope
	GetScope() *string
}

type NeuronProductScope struct {
	// example:
	//
	// order
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 读取订单数据
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// readOrder
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s NeuronProductScope) String() string {
	return dara.Prettify(s)
}

func (s NeuronProductScope) GoString() string {
	return s.String()
}

func (s *NeuronProductScope) GetDescription() *string {
	return s.Description
}

func (s *NeuronProductScope) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *NeuronProductScope) GetGmtModified() *string {
	return s.GmtModified
}

func (s *NeuronProductScope) GetName() *string {
	return s.Name
}

func (s *NeuronProductScope) GetProductId() *int64 {
	return s.ProductId
}

func (s *NeuronProductScope) GetScope() *string {
	return s.Scope
}

func (s *NeuronProductScope) SetDescription(v string) *NeuronProductScope {
	s.Description = &v
	return s
}

func (s *NeuronProductScope) SetGmtCreate(v string) *NeuronProductScope {
	s.GmtCreate = &v
	return s
}

func (s *NeuronProductScope) SetGmtModified(v string) *NeuronProductScope {
	s.GmtModified = &v
	return s
}

func (s *NeuronProductScope) SetName(v string) *NeuronProductScope {
	s.Name = &v
	return s
}

func (s *NeuronProductScope) SetProductId(v int64) *NeuronProductScope {
	s.ProductId = &v
	return s
}

func (s *NeuronProductScope) SetScope(v string) *NeuronProductScope {
	s.Scope = &v
	return s
}

func (s *NeuronProductScope) Validate() error {
	return dara.Validate(s)
}
