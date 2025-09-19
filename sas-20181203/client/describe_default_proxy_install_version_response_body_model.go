// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultProxyInstallVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstallVersion(v string) *DescribeDefaultProxyInstallVersionResponseBody
	GetInstallVersion() *string
	SetRequestId(v string) *DescribeDefaultProxyInstallVersionResponseBody
	GetRequestId() *string
}

type DescribeDefaultProxyInstallVersionResponseBody struct {
	// The default installation version.
	//
	// example:
	//
	// proxy_01_03
	InstallVersion *string `json:"InstallVersion,omitempty" xml:"InstallVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F9FCB51A-5078-5D31-9C4D-3B25BEF068C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefaultProxyInstallVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultProxyInstallVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefaultProxyInstallVersionResponseBody) GetInstallVersion() *string {
	return s.InstallVersion
}

func (s *DescribeDefaultProxyInstallVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefaultProxyInstallVersionResponseBody) SetInstallVersion(v string) *DescribeDefaultProxyInstallVersionResponseBody {
	s.InstallVersion = &v
	return s
}

func (s *DescribeDefaultProxyInstallVersionResponseBody) SetRequestId(v string) *DescribeDefaultProxyInstallVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefaultProxyInstallVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
