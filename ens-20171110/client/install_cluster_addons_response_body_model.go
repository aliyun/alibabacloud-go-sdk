// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallClusterAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InstallClusterAddonsResponseBody
	GetClusterId() *string
	SetRequestId(v string) *InstallClusterAddonsResponseBody
	GetRequestId() *string
}

type InstallClusterAddonsResponseBody struct {
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallClusterAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallClusterAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallClusterAddonsResponseBody) SetClusterId(v string) *InstallClusterAddonsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *InstallClusterAddonsResponseBody) SetRequestId(v string) *InstallClusterAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallClusterAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}
