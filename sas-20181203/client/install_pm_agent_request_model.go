// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPmAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *InstallPmAgentRequest
	GetLang() *string
	SetSourceIp(v string) *InstallPmAgentRequest
	GetSourceIp() *string
	SetType(v string) *InstallPmAgentRequest
	GetType() *string
	SetUuids(v string) *InstallPmAgentRequest
	GetUuids() *string
}

type InstallPmAgentRequest struct {
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
	// The source IP address of the request.
	//
	// example:
	//
	// 58.35.xx.xx
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the client.
	//
	// 	- **aliyun_assist**: Cloud Assistant client
	//
	// 	- **aliyun_monitor**: CloudMonitor client
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun_monitor
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server. If you specify multiple UUIDs, separate the UUIDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-a6444920-d303-4ccf-ab87-a1d3cd49****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s InstallPmAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallPmAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallPmAgentRequest) GetLang() *string {
	return s.Lang
}

func (s *InstallPmAgentRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *InstallPmAgentRequest) GetType() *string {
	return s.Type
}

func (s *InstallPmAgentRequest) GetUuids() *string {
	return s.Uuids
}

func (s *InstallPmAgentRequest) SetLang(v string) *InstallPmAgentRequest {
	s.Lang = &v
	return s
}

func (s *InstallPmAgentRequest) SetSourceIp(v string) *InstallPmAgentRequest {
	s.SourceIp = &v
	return s
}

func (s *InstallPmAgentRequest) SetType(v string) *InstallPmAgentRequest {
	s.Type = &v
	return s
}

func (s *InstallPmAgentRequest) SetUuids(v string) *InstallPmAgentRequest {
	s.Uuids = &v
	return s
}

func (s *InstallPmAgentRequest) Validate() error {
	return dara.Validate(s)
}
