// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpServiceGroupUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpServiceGroupUpdateCmd
	GetAlias() *string
	SetDescription(v string) *PdpServiceGroupUpdateCmd
	GetDescription() *string
	SetId(v int64) *PdpServiceGroupUpdateCmd
	GetId() *int64
}

type PdpServiceGroupUpdateCmd struct {
	// example:
	//
	// 管理
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 管理
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s PdpServiceGroupUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpServiceGroupUpdateCmd) GoString() string {
	return s.String()
}

func (s *PdpServiceGroupUpdateCmd) GetAlias() *string {
	return s.Alias
}

func (s *PdpServiceGroupUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *PdpServiceGroupUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *PdpServiceGroupUpdateCmd) SetAlias(v string) *PdpServiceGroupUpdateCmd {
	s.Alias = &v
	return s
}

func (s *PdpServiceGroupUpdateCmd) SetDescription(v string) *PdpServiceGroupUpdateCmd {
	s.Description = &v
	return s
}

func (s *PdpServiceGroupUpdateCmd) SetId(v int64) *PdpServiceGroupUpdateCmd {
	s.Id = &v
	return s
}

func (s *PdpServiceGroupUpdateCmd) Validate() error {
	return dara.Validate(s)
}
