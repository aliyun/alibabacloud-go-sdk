// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotPresetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotImageName(v string) *CreateHoneypotPresetRequest
	GetHoneypotImageName() *string
	SetLang(v string) *CreateHoneypotPresetRequest
	GetLang() *string
	SetMeta(v string) *CreateHoneypotPresetRequest
	GetMeta() *string
	SetNodeId(v string) *CreateHoneypotPresetRequest
	GetNodeId() *string
	SetPresetName(v string) *CreateHoneypotPresetRequest
	GetPresetName() *string
}

type CreateHoneypotPresetRequest struct {
	// The name of the honeypot image.
	//
	// This parameter is required.
	//
	// example:
	//
	// webmin
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The custom configurations of the honeypot template. The value is a JSON string that contains the following fields:
	//
	// 	- **portrait_option**: Social Source Tracing
	//
	// 	- **burp**: Burp-specific Defense
	//
	// 	- **trojan_git**: Git-specific Defense
	//
	// This parameter is required.
	//
	// example:
	//
	// {"burp":"close","trojan_git":"close","portrait_option":"true"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The ID of the management node to which you want to deploy honeypots.
	//
	// > You can call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to query the IDs of management nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9373fe59-74d5-4505-bb24-c85352fb****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The custom name of the honeypot template.
	//
	// This parameter is required.
	//
	// example:
	//
	// WebMin-online
	PresetName *string `json:"PresetName,omitempty" xml:"PresetName,omitempty"`
}

func (s CreateHoneypotPresetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotPresetRequest) GoString() string {
	return s.String()
}

func (s *CreateHoneypotPresetRequest) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *CreateHoneypotPresetRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateHoneypotPresetRequest) GetMeta() *string {
	return s.Meta
}

func (s *CreateHoneypotPresetRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateHoneypotPresetRequest) GetPresetName() *string {
	return s.PresetName
}

func (s *CreateHoneypotPresetRequest) SetHoneypotImageName(v string) *CreateHoneypotPresetRequest {
	s.HoneypotImageName = &v
	return s
}

func (s *CreateHoneypotPresetRequest) SetLang(v string) *CreateHoneypotPresetRequest {
	s.Lang = &v
	return s
}

func (s *CreateHoneypotPresetRequest) SetMeta(v string) *CreateHoneypotPresetRequest {
	s.Meta = &v
	return s
}

func (s *CreateHoneypotPresetRequest) SetNodeId(v string) *CreateHoneypotPresetRequest {
	s.NodeId = &v
	return s
}

func (s *CreateHoneypotPresetRequest) SetPresetName(v string) *CreateHoneypotPresetRequest {
	s.PresetName = &v
	return s
}

func (s *CreateHoneypotPresetRequest) Validate() error {
	return dara.Validate(s)
}
