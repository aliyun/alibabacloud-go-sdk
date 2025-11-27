// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallRCCloudAssistantResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InstallRCCloudAssistantResponseBody
	GetRequestId() *string
}

type InstallRCCloudAssistantResponseBody struct {
	// example:
	//
	// 842B73C8-5776-4BD9-9872-69C8C46DD7D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallRCCloudAssistantResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallRCCloudAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *InstallRCCloudAssistantResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallRCCloudAssistantResponseBody) SetRequestId(v string) *InstallRCCloudAssistantResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallRCCloudAssistantResponseBody) Validate() error {
	return dara.Validate(s)
}
