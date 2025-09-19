// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPmAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallPmAgentResponseBody
	GetRequestId() *string
}

type InstallPmAgentResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D49B5134-9511-5736-B447-BEE0AC66****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallPmAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallPmAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallPmAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallPmAgentResponseBody) SetRequestId(v string) *InstallPmAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallPmAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
