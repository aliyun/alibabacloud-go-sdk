// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeAIServiceResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeAIServiceResponseBody
	GetDescription() *string
	SetPrivateApiDevUrl(v string) *DescribeAIServiceResponseBody
	GetPrivateApiDevUrl() *string
	SetPrivateWorkbenchUrl(v string) *DescribeAIServiceResponseBody
	GetPrivateWorkbenchUrl() *string
	SetPublicApiDevUrl(v string) *DescribeAIServiceResponseBody
	GetPublicApiDevUrl() *string
	SetPublicWorkbenchUrl(v string) *DescribeAIServiceResponseBody
	GetPublicWorkbenchUrl() *string
	SetRequestId(v string) *DescribeAIServiceResponseBody
	GetRequestId() *string
	SetSecurityIpList(v string) *DescribeAIServiceResponseBody
	GetSecurityIpList() *string
	SetServiceAccount(v string) *DescribeAIServiceResponseBody
	GetServiceAccount() *string
	SetServiceId(v string) *DescribeAIServiceResponseBody
	GetServiceId() *string
	SetStatus(v string) *DescribeAIServiceResponseBody
	GetStatus() *string
}

type DescribeAIServiceResponseBody struct {
	// The time when the service was created.
	//
	// example:
	//
	// 2026-03-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// dramatest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The internal endpoint for API debugging.
	//
	// example:
	//
	// 192.168.0.1/api-docs
	PrivateApiDevUrl *string `json:"PrivateApiDevUrl,omitempty" xml:"PrivateApiDevUrl,omitempty"`
	// The internal endpoint of the Workbench.
	//
	// example:
	//
	// 192.168.0.1
	PrivateWorkbenchUrl *string `json:"PrivateWorkbenchUrl,omitempty" xml:"PrivateWorkbenchUrl,omitempty"`
	// The public endpoint for API debugging.
	//
	// example:
	//
	// 8.8.8.8/api-docs
	PublicApiDevUrl *string `json:"PublicApiDevUrl,omitempty" xml:"PublicApiDevUrl,omitempty"`
	// The public endpoint of the Workbench.
	//
	// example:
	//
	// 8.8.8.8
	PublicWorkbenchUrl *string `json:"PublicWorkbenchUrl,omitempty" xml:"PublicWorkbenchUrl,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A comma-separated list of IP addresses in the IP address whitelist group.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	// The service account.
	//
	// example:
	//
	// dramauser
	ServiceAccount *string `json:"ServiceAccount,omitempty" xml:"ServiceAccount,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// drama-123456
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The status of the service. Valid values:
	//
	// - deploying: The service is being deployed.
	//
	// - active: The service is running.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAIServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIServiceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAIServiceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeAIServiceResponseBody) GetPrivateApiDevUrl() *string {
	return s.PrivateApiDevUrl
}

func (s *DescribeAIServiceResponseBody) GetPrivateWorkbenchUrl() *string {
	return s.PrivateWorkbenchUrl
}

func (s *DescribeAIServiceResponseBody) GetPublicApiDevUrl() *string {
	return s.PublicApiDevUrl
}

func (s *DescribeAIServiceResponseBody) GetPublicWorkbenchUrl() *string {
	return s.PublicWorkbenchUrl
}

func (s *DescribeAIServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIServiceResponseBody) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *DescribeAIServiceResponseBody) GetServiceAccount() *string {
	return s.ServiceAccount
}

func (s *DescribeAIServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeAIServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAIServiceResponseBody) SetCreateTime(v string) *DescribeAIServiceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetDescription(v string) *DescribeAIServiceResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetPrivateApiDevUrl(v string) *DescribeAIServiceResponseBody {
	s.PrivateApiDevUrl = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetPrivateWorkbenchUrl(v string) *DescribeAIServiceResponseBody {
	s.PrivateWorkbenchUrl = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetPublicApiDevUrl(v string) *DescribeAIServiceResponseBody {
	s.PublicApiDevUrl = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetPublicWorkbenchUrl(v string) *DescribeAIServiceResponseBody {
	s.PublicWorkbenchUrl = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetRequestId(v string) *DescribeAIServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetSecurityIpList(v string) *DescribeAIServiceResponseBody {
	s.SecurityIpList = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetServiceAccount(v string) *DescribeAIServiceResponseBody {
	s.ServiceAccount = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetServiceId(v string) *DescribeAIServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeAIServiceResponseBody) SetStatus(v string) *DescribeAIServiceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAIServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
