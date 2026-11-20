// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ClientCreateCmd
	GetAddress() *string
	SetAllowedModels(v string) *ClientCreateCmd
	GetAllowedModels() *string
	SetContact(v string) *ClientCreateCmd
	GetContact() *string
	SetName(v string) *ClientCreateCmd
	GetName() *string
	SetRemark(v string) *ClientCreateCmd
	GetRemark() *string
}

type ClientCreateCmd struct {
	// The company address.
	//
	// example:
	//
	// Hangzhou
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The list of allowed model IDs, separated by commas. An empty value indicates all models.
	//
	// example:
	//
	// 1,2,3
	AllowedModels *string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty"`
	// The contact information.
	//
	// example:
	//
	// 13800138000
	Contact *string `json:"contact,omitempty" xml:"contact,omitempty"`
	// The customer name.
	//
	// example:
	//
	// My customer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Remark
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s ClientCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ClientCreateCmd) GoString() string {
	return s.String()
}

func (s *ClientCreateCmd) GetAddress() *string {
	return s.Address
}

func (s *ClientCreateCmd) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *ClientCreateCmd) GetContact() *string {
	return s.Contact
}

func (s *ClientCreateCmd) GetName() *string {
	return s.Name
}

func (s *ClientCreateCmd) GetRemark() *string {
	return s.Remark
}

func (s *ClientCreateCmd) SetAddress(v string) *ClientCreateCmd {
	s.Address = &v
	return s
}

func (s *ClientCreateCmd) SetAllowedModels(v string) *ClientCreateCmd {
	s.AllowedModels = &v
	return s
}

func (s *ClientCreateCmd) SetContact(v string) *ClientCreateCmd {
	s.Contact = &v
	return s
}

func (s *ClientCreateCmd) SetName(v string) *ClientCreateCmd {
	s.Name = &v
	return s
}

func (s *ClientCreateCmd) SetRemark(v string) *ClientCreateCmd {
	s.Remark = &v
	return s
}

func (s *ClientCreateCmd) Validate() error {
	return dara.Validate(s)
}
