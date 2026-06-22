// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulTargetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyVulTargetConfigRequest
	GetConfig() *string
	SetSourceIp(v string) *ModifyVulTargetConfigRequest
	GetSourceIp() *string
	SetType(v string) *ModifyVulTargetConfigRequest
	GetType() *string
	SetUuid(v string) *ModifyVulTargetConfigRequest
	GetUuid() *string
}

type ModifyVulTargetConfigRequest struct {
	// Specifies whether to enable or disable vulnerability detection. Valid values:
	//
	// - **on**: Enable vulnerability detection.
	//
	// - **off**: Disable vulnerability detection.
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of vulnerability detection to configure. Valid values:
	//
	// - **cve**: Linux software vulnerability
	//
	// - **sys**: Windows system vulnerability
	//
	// - **cms**: Web-CMS vulnerability
	//
	// - **emg**: urgent vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the server to configure.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-7c676676-06fa-442e-90fb-b802e5d6****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyVulTargetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulTargetConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyVulTargetConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyVulTargetConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyVulTargetConfigRequest) GetType() *string {
	return s.Type
}

func (s *ModifyVulTargetConfigRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyVulTargetConfigRequest) SetConfig(v string) *ModifyVulTargetConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyVulTargetConfigRequest) SetSourceIp(v string) *ModifyVulTargetConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVulTargetConfigRequest) SetType(v string) *ModifyVulTargetConfigRequest {
	s.Type = &v
	return s
}

func (s *ModifyVulTargetConfigRequest) SetUuid(v string) *ModifyVulTargetConfigRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyVulTargetConfigRequest) Validate() error {
	return dara.Validate(s)
}
