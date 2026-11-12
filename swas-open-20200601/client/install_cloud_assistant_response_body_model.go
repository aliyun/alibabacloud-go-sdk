// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAssistantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallCloudAssistantResponseBody
	GetRequestId() *string
}

type InstallCloudAssistantResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallCloudAssistantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallCloudAssistantResponseBody) SetRequestId(v string) *InstallCloudAssistantResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallCloudAssistantResponseBody) Validate() error {
	return dara.Validate(s)
}
