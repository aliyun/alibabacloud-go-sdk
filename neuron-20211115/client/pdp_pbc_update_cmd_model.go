// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpPbcUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpPbcUpdateCmd
	GetAlias() *string
	SetDescription(v string) *PdpPbcUpdateCmd
	GetDescription() *string
	SetId(v int64) *PdpPbcUpdateCmd
	GetId() *int64
	SetName(v string) *PdpPbcUpdateCmd
	GetName() *string
}

type PdpPbcUpdateCmd struct {
	Alias       *string `json:"alias,omitempty" xml:"alias,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s PdpPbcUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpPbcUpdateCmd) GoString() string {
	return s.String()
}

func (s *PdpPbcUpdateCmd) GetAlias() *string {
	return s.Alias
}

func (s *PdpPbcUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *PdpPbcUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *PdpPbcUpdateCmd) GetName() *string {
	return s.Name
}

func (s *PdpPbcUpdateCmd) SetAlias(v string) *PdpPbcUpdateCmd {
	s.Alias = &v
	return s
}

func (s *PdpPbcUpdateCmd) SetDescription(v string) *PdpPbcUpdateCmd {
	s.Description = &v
	return s
}

func (s *PdpPbcUpdateCmd) SetId(v int64) *PdpPbcUpdateCmd {
	s.Id = &v
	return s
}

func (s *PdpPbcUpdateCmd) SetName(v string) *PdpPbcUpdateCmd {
	s.Name = &v
	return s
}

func (s *PdpPbcUpdateCmd) Validate() error {
	return dara.Validate(s)
}
