// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateAclRequest
	GetAcceptLanguage() *string
	SetAclEntryList(v string) *UpdateAclRequest
	GetAclEntryList() *string
	SetInstanceId(v string) *UpdateAclRequest
	GetInstanceId() *string
	SetNetworkType(v string) *UpdateAclRequest
	GetNetworkType() *string
}

type UpdateAclRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The IP addresses or CIDR blocks in the IP address whitelist.
	//
	// example:
	//
	// 192.168.0.0/XX,192.168.0.0/XX
	AclEntryList *string `json:"AclEntryList,omitempty" xml:"AclEntryList,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-78v1l83****
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s UpdateAclRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateAclRequest) GetAclEntryList() *string {
	return s.AclEntryList
}

func (s *UpdateAclRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAclRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateAclRequest) SetAcceptLanguage(v string) *UpdateAclRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateAclRequest) SetAclEntryList(v string) *UpdateAclRequest {
	s.AclEntryList = &v
	return s
}

func (s *UpdateAclRequest) SetInstanceId(v string) *UpdateAclRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAclRequest) SetNetworkType(v string) *UpdateAclRequest {
	s.NetworkType = &v
	return s
}

func (s *UpdateAclRequest) Validate() error {
	return dara.Validate(s)
}
