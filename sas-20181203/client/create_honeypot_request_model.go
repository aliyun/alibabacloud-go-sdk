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
	// The honeypot image ID.
	//
	// > You can obtain this value from the **HoneypotImageId*	- field returned by the [ListAvailableHoneypot](~~ListAvailableHoneypot~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// sha256:ebc4c102ac407d53733c2373e8888a733ddce86f163ccbe7492ae1cbf26****
	HoneypotImageId *string `json:"HoneypotImageId,omitempty" xml:"HoneypotImageId,omitempty"`
	// The honeypot image name.
	//
	// > You can obtain this value from the **HoneypotImageName*	- field returned by the [ListAvailableHoneypot](~~ListAvailableHoneypot~~) operation.
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
	// The custom configuration of the honeypot in JSON format. The following fields are included:
	//
	// - **trojan_git**: The Git counter-intelligence method. Valid values:
	//
	//     -   **zip**: Git source code package.
	//
	//     -  **web**: .git folder leak.
	//
	//     -  **close**: Shutdown.
	//
	// - **trojan_git_addr**: The Git counter-intelligence endpoint.
	//
	// - **trojan_git.zip**: The Git counter-intelligence trojan package.
	//
	// - **burp**: The Burp counter-intelligence method. Valid values:
	//
	//      - **open**: Enabled.
	//
	//     - **close**: Shutdown.
	//
	// - **portrait_option**: The tracing configuration. Valid values:
	//
	//     - **false**: Shutdown.
	//
	//     - **true**: Enabled.
	//
	// example:
	//
	// {\\"trojan_git\\":\\"close\\",\\"burp\\":\\"close\\",\\"portrait_option\\":\\"false\\"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The ID of the honeypot management node.
	//
	// > Call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to obtain this value.
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
