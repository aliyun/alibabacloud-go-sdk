// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPushAllTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceIp(v string) *ModifyPushAllTaskRequest
	GetSourceIp() *string
	SetTasks(v string) *ModifyPushAllTaskRequest
	GetTasks() *string
	SetUuids(v string) *ModifyPushAllTaskRequest
	GetUuids() *string
}

type ModifyPushAllTaskRequest struct {
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The check items. Separate multiple check items with commas (,). Valid values:
	//
	// 	- **OVAL_ENTITY**: Common Vulnerabilities and Exposures (CVE) vulnerabilities.
	//
	// 	- **CMS**: Web-CMS vulnerabilities.
	//
	// 	- **SYSVUL**: Windows system vulnerabilities.
	//
	// 	- **SCA**: application vulnerabilities.
	//
	// 	- **HEALTH_CHECK**: baselines.
	//
	// 	- **WEBSHELL**: webshells.
	//
	// 	- **PROC_SNAPSHOT**: processes.
	//
	// 	- **PORT_SNAPSHOT**: ports.
	//
	// 	- **ACCOUNT_SNAPSHOT**: accounts.
	//
	// 	- **SOFTWARE_SNAPSHOT**: software assets.
	//
	// 	- **SCA_SNAPSHOT**: middleware, databases, and web services.
	//
	// 	- **CROND_SNAPSHOT**: scheduled tasks.
	//
	// 	- **AUTORUN_SNAPSHOT**: startup items.
	//
	// 	- **LKM_SNAPSHOT**: kernel modules.
	//
	// 	- **SCA_PROXY_SNAPSHOT**: websites.
	//
	// This parameter is required.
	//
	// example:
	//
	// HEALTH_CHECK,OVAL_ENTITY
	Tasks *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	// The UUIDs of servers on which you want to perform security check tasks. Separate multiple UUIDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-923b4538-0e88-409d-80ba-cb2e7487****,dc1691eb-656f-472f-b2aa-04f621f4****,70452f92-9fc1-45c5-ab35-e7bf8552****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ModifyPushAllTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPushAllTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyPushAllTaskRequest) GetTasks() *string {
	return s.Tasks
}

func (s *ModifyPushAllTaskRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ModifyPushAllTaskRequest) SetSourceIp(v string) *ModifyPushAllTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyPushAllTaskRequest) SetTasks(v string) *ModifyPushAllTaskRequest {
	s.Tasks = &v
	return s
}

func (s *ModifyPushAllTaskRequest) SetUuids(v string) *ModifyPushAllTaskRequest {
	s.Uuids = &v
	return s
}

func (s *ModifyPushAllTaskRequest) Validate() error {
	return dara.Validate(s)
}
