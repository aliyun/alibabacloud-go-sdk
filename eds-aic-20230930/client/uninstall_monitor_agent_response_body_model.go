// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallMonitorAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallMonitorAgentResponseBody
	GetRequestId() *string
}

type UninstallMonitorAgentResponseBody struct {
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallMonitorAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallMonitorAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallMonitorAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallMonitorAgentResponseBody) SetRequestId(v string) *UninstallMonitorAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallMonitorAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
