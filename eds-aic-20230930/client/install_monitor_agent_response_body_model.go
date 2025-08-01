// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitorAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallMonitorAgentResponseBody
	GetRequestId() *string
}

type InstallMonitorAgentResponseBody struct {
	// example:
	//
	// DB070C80-45AC-52CA-8101-937C25DA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallMonitorAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallMonitorAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallMonitorAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallMonitorAgentResponseBody) SetRequestId(v string) *InstallMonitorAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallMonitorAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
