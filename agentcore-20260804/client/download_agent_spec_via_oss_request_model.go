// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadAgentSpecViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpecVersion(v string) *DownloadAgentSpecViaOssRequest
	GetAgentSpecVersion() *string
}

type DownloadAgentSpecViaOssRequest struct {
	// The version number. If not specified, the version corresponding to the latest label is downloaded.
	//
	// example:
	//
	// 1.0.0
	AgentSpecVersion *string `json:"agentSpecVersion,omitempty" xml:"agentSpecVersion,omitempty"`
}

func (s DownloadAgentSpecViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadAgentSpecViaOssRequest) GoString() string {
	return s.String()
}

func (s *DownloadAgentSpecViaOssRequest) GetAgentSpecVersion() *string {
	return s.AgentSpecVersion
}

func (s *DownloadAgentSpecViaOssRequest) SetAgentSpecVersion(v string) *DownloadAgentSpecViaOssRequest {
	s.AgentSpecVersion = &v
	return s
}

func (s *DownloadAgentSpecViaOssRequest) Validate() error {
	return dara.Validate(s)
}
