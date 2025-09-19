// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstallCodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstallCodes(v []*DescribeInstallCodesResponseBodyInstallCodes) *DescribeInstallCodesResponseBody
	GetInstallCodes() []*DescribeInstallCodesResponseBodyInstallCodes
	SetRequestId(v string) *DescribeInstallCodesResponseBody
	GetRequestId() *string
}

type DescribeInstallCodesResponseBody struct {
	// The information about the installation commands.
	InstallCodes []*DescribeInstallCodesResponseBodyInstallCodes `json:"InstallCodes,omitempty" xml:"InstallCodes,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C0D6119F-92EE-1276-B8B6-C81A7F9D57F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstallCodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstallCodesResponseBody) GetInstallCodes() []*DescribeInstallCodesResponseBodyInstallCodes {
	return s.InstallCodes
}

func (s *DescribeInstallCodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstallCodesResponseBody) SetInstallCodes(v []*DescribeInstallCodesResponseBodyInstallCodes) *DescribeInstallCodesResponseBody {
	s.InstallCodes = v
	return s
}

func (s *DescribeInstallCodesResponseBody) SetRequestId(v string) *DescribeInstallCodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstallCodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstallCodesResponseBodyInstallCodes struct {
	// The verification code for you to manually install the Security Center agent.
	//
	// example:
	//
	// 15v02r
	CaptchaCode *string `json:"CaptchaCode,omitempty" xml:"CaptchaCode,omitempty"`
	// The timestamp generated when the commands used to install the Security Center agent expire. Unit: milliseconds.
	//
	// example:
	//
	// 1637810007000
	ExpiredDate *int64 `json:"ExpiredDate,omitempty" xml:"ExpiredDate,omitempty"`
	// The ID of the server group to which the server belongs.
	//
	// example:
	//
	// 9165712
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the server group to which the server belongs.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an image is used to install the Security Center agent. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	OnlyImage *bool `json:"OnlyImage,omitempty" xml:"OnlyImage,omitempty"`
	// The operating system of the server. Valid values:
	//
	// 	- **linux**
	//
	// 	- **windows**
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The ID of the PrivateLink endpoint.
	//
	// example:
	//
	// 123
	PrivateLinkEndpointId *int64 `json:"PrivateLinkEndpointId,omitempty" xml:"PrivateLinkEndpointId,omitempty"`
	// The name of the proxy cluster.
	//
	// example:
	//
	// proxy_test
	ProxyCluster *string `json:"ProxyCluster,omitempty" xml:"ProxyCluster,omitempty"`
	// The name of the server provider.
	//
	// example:
	//
	// ALIYUN
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s DescribeInstallCodesResponseBodyInstallCodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCodesResponseBodyInstallCodes) GoString() string {
	return s.String()
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetCaptchaCode() *string {
	return s.CaptchaCode
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetExpiredDate() *int64 {
	return s.ExpiredDate
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetOnlyImage() *bool {
	return s.OnlyImage
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetOs() *string {
	return s.Os
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetPrivateLinkEndpointId() *int64 {
	return s.PrivateLinkEndpointId
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetProxyCluster() *string {
	return s.ProxyCluster
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) GetVendorName() *string {
	return s.VendorName
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetCaptchaCode(v string) *DescribeInstallCodesResponseBodyInstallCodes {
	s.CaptchaCode = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetExpiredDate(v int64) *DescribeInstallCodesResponseBodyInstallCodes {
	s.ExpiredDate = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetGroupId(v int64) *DescribeInstallCodesResponseBodyInstallCodes {
	s.GroupId = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetGroupName(v string) *DescribeInstallCodesResponseBodyInstallCodes {
	s.GroupName = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetOnlyImage(v bool) *DescribeInstallCodesResponseBodyInstallCodes {
	s.OnlyImage = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetOs(v string) *DescribeInstallCodesResponseBodyInstallCodes {
	s.Os = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetPrivateLinkEndpointId(v int64) *DescribeInstallCodesResponseBodyInstallCodes {
	s.PrivateLinkEndpointId = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetProxyCluster(v string) *DescribeInstallCodesResponseBodyInstallCodes {
	s.ProxyCluster = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) SetVendorName(v string) *DescribeInstallCodesResponseBodyInstallCodes {
	s.VendorName = &v
	return s
}

func (s *DescribeInstallCodesResponseBodyInstallCodes) Validate() error {
	return dara.Validate(s)
}
