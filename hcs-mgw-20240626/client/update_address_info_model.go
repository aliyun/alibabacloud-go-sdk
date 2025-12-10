// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAddressInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentList(v string) *UpdateAddressInfo
	GetAgentList() *string
}

type UpdateAddressInfo struct {
	AgentList *string `json:"AgentList,omitempty" xml:"AgentList,omitempty"`
}

func (s UpdateAddressInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateAddressInfo) GoString() string {
	return s.String()
}

func (s *UpdateAddressInfo) GetAgentList() *string {
	return s.AgentList
}

func (s *UpdateAddressInfo) SetAgentList(v string) *UpdateAddressInfo {
	s.AgentList = &v
	return s
}

func (s *UpdateAddressInfo) Validate() error {
	return dara.Validate(s)
}
