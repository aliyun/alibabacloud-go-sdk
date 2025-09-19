// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTrialFixCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInfo(v string) *CheckTrialFixCountRequest
	GetInfo() *string
	SetType(v string) *CheckTrialFixCountRequest
	GetType() *string
	SetUuids(v []*string) *CheckTrialFixCountRequest
	GetUuids() []*string
	SetVulNames(v []*string) *CheckTrialFixCountRequest
	GetVulNames() []*string
}

type CheckTrialFixCountRequest struct {
	// The information about the vulnerability. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **name**: the name of the vulnerability.
	//
	// 	- **uuid**: the UUID of the server on which the vulnerability is detected.
	//
	// 	- **tag**: the tag that is added to the vulnerability. Valid values:
	//
	//     	- **oval**: Linux software vulnerability.
	//
	//     	- **system**: Windows system vulnerability.
	//
	//     	- **cms**: Web-CMS vulnerability.
	//
	// >  You must specify a value for Info or values for VulNames and Uuids to identify a vulnerability.
	//
	// example:
	//
	// [{\\"name\\":\\"oval:com.redhat.rhsa:def:20192143\\",\\"uuid\\":\\"80ee3226-1f96-4da0-a3ed-55c104e2****\\",\\"tag\\":\\"oval\\"}]
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The type of the vulnerability that you want to fix. Valid values:
	//
	// 	- **cve**: Linux software vulnerability.
	//
	// 	- **sys**: Windows system vulnerability.
	//
	// 	- **cms**: Web-CMS vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUIDs of the servers.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
	// The names of the vulnerabilities.
	VulNames []*string `json:"VulNames,omitempty" xml:"VulNames,omitempty" type:"Repeated"`
}

func (s CheckTrialFixCountRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckTrialFixCountRequest) GoString() string {
	return s.String()
}

func (s *CheckTrialFixCountRequest) GetInfo() *string {
	return s.Info
}

func (s *CheckTrialFixCountRequest) GetType() *string {
	return s.Type
}

func (s *CheckTrialFixCountRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *CheckTrialFixCountRequest) GetVulNames() []*string {
	return s.VulNames
}

func (s *CheckTrialFixCountRequest) SetInfo(v string) *CheckTrialFixCountRequest {
	s.Info = &v
	return s
}

func (s *CheckTrialFixCountRequest) SetType(v string) *CheckTrialFixCountRequest {
	s.Type = &v
	return s
}

func (s *CheckTrialFixCountRequest) SetUuids(v []*string) *CheckTrialFixCountRequest {
	s.Uuids = v
	return s
}

func (s *CheckTrialFixCountRequest) SetVulNames(v []*string) *CheckTrialFixCountRequest {
	s.VulNames = v
	return s
}

func (s *CheckTrialFixCountRequest) Validate() error {
	return dara.Validate(s)
}
