// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWaIdPermissions interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *WaIdPermissions
	GetCode() *string
	SetIsBasicChild(v bool) *WaIdPermissions
	GetIsBasicChild() *bool
	SetName(v string) *WaIdPermissions
	GetName() *string
	SetSubPermissions(v []*WaIdPermissions) *WaIdPermissions
	GetSubPermissions() []*WaIdPermissions
	SetType(v string) *WaIdPermissions
	GetType() *string
}

type WaIdPermissions struct {
	Code           *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	IsBasicChild   *bool              `json:"IsBasicChild,omitempty" xml:"IsBasicChild,omitempty"`
	Name           *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	SubPermissions []*WaIdPermissions `json:"SubPermissions,omitempty" xml:"SubPermissions,omitempty" type:"Repeated"`
	Type           *string            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s WaIdPermissions) String() string {
	return dara.Prettify(s)
}

func (s WaIdPermissions) GoString() string {
	return s.String()
}

func (s *WaIdPermissions) GetCode() *string {
	return s.Code
}

func (s *WaIdPermissions) GetIsBasicChild() *bool {
	return s.IsBasicChild
}

func (s *WaIdPermissions) GetName() *string {
	return s.Name
}

func (s *WaIdPermissions) GetSubPermissions() []*WaIdPermissions {
	return s.SubPermissions
}

func (s *WaIdPermissions) GetType() *string {
	return s.Type
}

func (s *WaIdPermissions) SetCode(v string) *WaIdPermissions {
	s.Code = &v
	return s
}

func (s *WaIdPermissions) SetIsBasicChild(v bool) *WaIdPermissions {
	s.IsBasicChild = &v
	return s
}

func (s *WaIdPermissions) SetName(v string) *WaIdPermissions {
	s.Name = &v
	return s
}

func (s *WaIdPermissions) SetSubPermissions(v []*WaIdPermissions) *WaIdPermissions {
	s.SubPermissions = v
	return s
}

func (s *WaIdPermissions) SetType(v string) *WaIdPermissions {
	s.Type = &v
	return s
}

func (s *WaIdPermissions) Validate() error {
	return dara.Validate(s)
}
