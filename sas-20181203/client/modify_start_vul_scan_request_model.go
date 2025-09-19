// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStartVulScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTypes(v string) *ModifyStartVulScanRequest
	GetTypes() *string
	SetUuids(v string) *ModifyStartVulScanRequest
	GetUuids() *string
}

type ModifyStartVulScanRequest struct {
	// The types of vulnerabilities that can be detected. Valid values:
	//
	// 	- **cve**: Linux software vulnerabilities
	//
	// 	- **sys**: Windows system vulnerabilities
	//
	// 	- **cms**: Web-CMS vulnerabilities
	//
	// 	- **app**: application vulnerabilities
	//
	// 	- **emg**: urgent vulnerabilities
	//
	// 	- **image**: container image vulnerabilities
	//
	// 	- **sca**: vulnerabilities that are detected based on software component analysis
	//
	// > If you leave this parameter empty, all types of vulnerabilities can be detected.
	//
	// example:
	//
	// "cve,sys,cms,app,emg"
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
	// The UUIDs of servers.
	//
	// example:
	//
	// {"i-sdada-xxxxx","i-ifaedada-sfsasdxxx"}
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ModifyStartVulScanRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStartVulScanRequest) GoString() string {
	return s.String()
}

func (s *ModifyStartVulScanRequest) GetTypes() *string {
	return s.Types
}

func (s *ModifyStartVulScanRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ModifyStartVulScanRequest) SetTypes(v string) *ModifyStartVulScanRequest {
	s.Types = &v
	return s
}

func (s *ModifyStartVulScanRequest) SetUuids(v string) *ModifyStartVulScanRequest {
	s.Uuids = &v
	return s
}

func (s *ModifyStartVulScanRequest) Validate() error {
	return dara.Validate(s)
}
