// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDownloadUrlV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *GetAgentDownloadUrlV2Request
	GetAgentType() *string
	SetArchType(v string) *GetAgentDownloadUrlV2Request
	GetArchType() *string
	SetOsType(v string) *GetAgentDownloadUrlV2Request
	GetOsType() *string
}

type GetAgentDownloadUrlV2Request struct {
	// The agent type.\\
	//
	// **Valid values:**
	//
	// 	- **JavaAgent**
	//
	// 	- **Instgo**
	//
	// This parameter is required.
	//
	// example:
	//
	// JavaAgent
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The architecture type of the environment where the agent is installed.\\
	//
	// This parameter is required and valid only when **AgentType*	- is set to **Instgo**.\\
	//
	// **Valid values:**
	//
	// 	- **amd64**
	//
	// 	- **arm64**
	//
	// example:
	//
	// amd64
	ArchType *string `json:"ArchType,omitempty" xml:"ArchType,omitempty"`
	// The operating system of the environment where the agent is installed.\\
	//
	// This parameter is required and valid only when **AgentType*	- is set to **Instgo**.\\
	//
	// **Valid values:**
	//
	// 	- **linux**
	//
	// 	- **darwin**
	//
	// 	- **windows**
	//
	// example:
	//
	// linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s GetAgentDownloadUrlV2Request) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDownloadUrlV2Request) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlV2Request) GetAgentType() *string {
	return s.AgentType
}

func (s *GetAgentDownloadUrlV2Request) GetArchType() *string {
	return s.ArchType
}

func (s *GetAgentDownloadUrlV2Request) GetOsType() *string {
	return s.OsType
}

func (s *GetAgentDownloadUrlV2Request) SetAgentType(v string) *GetAgentDownloadUrlV2Request {
	s.AgentType = &v
	return s
}

func (s *GetAgentDownloadUrlV2Request) SetArchType(v string) *GetAgentDownloadUrlV2Request {
	s.ArchType = &v
	return s
}

func (s *GetAgentDownloadUrlV2Request) SetOsType(v string) *GetAgentDownloadUrlV2Request {
	s.OsType = &v
	return s
}

func (s *GetAgentDownloadUrlV2Request) Validate() error {
	return dara.Validate(s)
}
