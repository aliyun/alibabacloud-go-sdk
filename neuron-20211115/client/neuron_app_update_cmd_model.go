// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNeuronAppUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *NeuronAppUpdateCmd
	GetDesc() *string
	SetId(v int64) *NeuronAppUpdateCmd
	GetId() *int64
	SetStatus(v string) *NeuronAppUpdateCmd
	GetStatus() *string
}

type NeuronAppUpdateCmd struct {
	// example:
	//
	// 升级订单功能
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// NEW
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s NeuronAppUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s NeuronAppUpdateCmd) GoString() string {
	return s.String()
}

func (s *NeuronAppUpdateCmd) GetDesc() *string {
	return s.Desc
}

func (s *NeuronAppUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *NeuronAppUpdateCmd) GetStatus() *string {
	return s.Status
}

func (s *NeuronAppUpdateCmd) SetDesc(v string) *NeuronAppUpdateCmd {
	s.Desc = &v
	return s
}

func (s *NeuronAppUpdateCmd) SetId(v int64) *NeuronAppUpdateCmd {
	s.Id = &v
	return s
}

func (s *NeuronAppUpdateCmd) SetStatus(v string) *NeuronAppUpdateCmd {
	s.Status = &v
	return s
}

func (s *NeuronAppUpdateCmd) Validate() error {
	return dara.Validate(s)
}
