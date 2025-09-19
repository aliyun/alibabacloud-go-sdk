// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotImageId(v string) *CreateHoneypotRequest
	GetHoneypotImageId() *string
	SetHoneypotImageName(v string) *CreateHoneypotRequest
	GetHoneypotImageName() *string
	SetHoneypotName(v string) *CreateHoneypotRequest
	GetHoneypotName() *string
	SetMeta(v string) *CreateHoneypotRequest
	GetMeta() *string
	SetNodeId(v string) *CreateHoneypotRequest
	GetNodeId() *string
}

type CreateHoneypotRequest struct {
	// The ID of the honeypot image.
	//
	// > You can call the [ListAvailableHoneypot](~~ListAvailableHoneypot~~) operation to query the IDs of images from the **HoneypotImageId*	- response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// sha256:ebc4c102ac407d53733c2373e8888a733ddce86f163ccbe7492ae1cbf26****
	HoneypotImageId *string `json:"HoneypotImageId,omitempty" xml:"HoneypotImageId,omitempty"`
	// The name of the honeypot image.
	//
	// > You can call the [ListAvailableHoneypot](~~ListAvailableHoneypot~~) operation to query the names of images from the **HoneypotImageName*	- response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// HoneyPotImageName
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The custom name of the honeypot.
	//
	// This parameter is required.
	//
	// example:
	//
	// ruoyi
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The custom configuration of the honeypot in the JSON format. The value contains the following fields:
	//
	// 	- **trojan_git**: Git-specific Defense. Valid values:
	//
	//     	- **zip**: Git Source Code Package
	//
	//     	- **web**: Git Directory Leak
	//
	//     	- **close**: Disabled
	//
	// 	- **trojan_git_addr**: Git Trojan Address.
	//
	// 	- **trojan_git.zip**: Git Trojan.
	//
	// 	- **burp**: Burp-specific Defense. Valid values:
	//
	//     	- **open**: Enable
	//
	//     	- **close**: Disable
	//
	// 	- **portrait_option**: Source Tracing Configuration. Valid values:
	//
	//     	- **false**: Disable
	//
	//     	- **true**: Enable
	//
	// example:
	//
	// {\\"trojan_git\\":\\"close\\",\\"burp\\":\\"close\\",\\"portrait_option\\":\\"false\\"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The ID of the management node.
	//
	// > You can call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to query the IDs of management nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateHoneypotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotRequest) GoString() string {
	return s.String()
}

func (s *CreateHoneypotRequest) GetHoneypotImageId() *string {
	return s.HoneypotImageId
}

func (s *CreateHoneypotRequest) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *CreateHoneypotRequest) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *CreateHoneypotRequest) GetMeta() *string {
	return s.Meta
}

func (s *CreateHoneypotRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateHoneypotRequest) SetHoneypotImageId(v string) *CreateHoneypotRequest {
	s.HoneypotImageId = &v
	return s
}

func (s *CreateHoneypotRequest) SetHoneypotImageName(v string) *CreateHoneypotRequest {
	s.HoneypotImageName = &v
	return s
}

func (s *CreateHoneypotRequest) SetHoneypotName(v string) *CreateHoneypotRequest {
	s.HoneypotName = &v
	return s
}

func (s *CreateHoneypotRequest) SetMeta(v string) *CreateHoneypotRequest {
	s.Meta = &v
	return s
}

func (s *CreateHoneypotRequest) SetNodeId(v string) *CreateHoneypotRequest {
	s.NodeId = &v
	return s
}

func (s *CreateHoneypotRequest) Validate() error {
	return dara.Validate(s)
}
