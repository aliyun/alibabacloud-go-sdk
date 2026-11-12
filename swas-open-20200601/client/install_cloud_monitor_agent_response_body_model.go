// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudMonitorAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallCloudMonitorAgentResponseBody
	GetRequestId() *string
}

type InstallCloudMonitorAgentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallCloudMonitorAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudMonitorAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallCloudMonitorAgentResponseBody) SetRequestId(v string) *InstallCloudMonitorAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallCloudMonitorAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
