// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchConnectionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) *DescribeOpenSearchConnectionInfoResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchConnectionInfoResponseBodyData) *DescribeOpenSearchConnectionInfoResponseBody
	GetData() *DescribeOpenSearchConnectionInfoResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchConnectionInfoResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchConnectionInfoResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DescribeOpenSearchConnectionInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchConnectionInfoResponseBody) GetData() *DescribeOpenSearchConnectionInfoResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchConnectionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchConnectionInfoResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) *DescribeOpenSearchConnectionInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBody) SetData(v *DescribeOpenSearchConnectionInfoResponseBodyData) *DescribeOpenSearchConnectionInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBody) SetRequestId(v string) *DescribeOpenSearchConnectionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The description is the same as above.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The diagnostic information.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// NoPermissionType
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchConnectionInfoResponseBodyData struct {
	// The internal endpoint of the OpenSearch Dashboard.
	DashboardEndpoint *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint `json:"DashboardEndpoint,omitempty" xml:"DashboardEndpoint,omitempty" type:"Struct"`
	// The public network access endpoint of the OpenSearch Dashboard.
	DashboardPublicEndpoint *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint `json:"DashboardPublicEndpoint,omitempty" xml:"DashboardPublicEndpoint,omitempty" type:"Struct"`
	// The default account name of OpenSearch.
	//
	// example:
	//
	// elastic
	DefaultUsername *string `json:"DefaultUsername,omitempty" xml:"DefaultUsername,omitempty"`
	// The VPC endpoint of the instance.
	PrivateEndpoint *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty" type:"Struct"`
	// The protocol of the monitoring task. Valid values:
	//
	// - **ICMP**.
	//
	// - **TCP**.
	//
	// - **HTTP**.
	//
	// > Private network monitoring supports only the ICMP and TCP protocols.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The public endpoint of the instance.
	PublicEndpoint *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty" type:"Struct"`
}

func (s DescribeOpenSearchConnectionInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) GetDashboardEndpoint() *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint {
	return s.DashboardEndpoint
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) GetDashboardPublicEndpoint() *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint {
	return s.DashboardPublicEndpoint
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) GetDefaultUsername() *string {
	return s.DefaultUsername
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) GetPrivateEndpoint() *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint {
	return s.PrivateEndpoint
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) GetPublicEndpoint() *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint {
	return s.PublicEndpoint
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) SetDashboardEndpoint(v *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) *DescribeOpenSearchConnectionInfoResponseBodyData {
	s.DashboardEndpoint = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) SetDashboardPublicEndpoint(v *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) *DescribeOpenSearchConnectionInfoResponseBodyData {
	s.DashboardPublicEndpoint = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) SetDefaultUsername(v string) *DescribeOpenSearchConnectionInfoResponseBodyData {
	s.DefaultUsername = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) SetPrivateEndpoint(v *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) *DescribeOpenSearchConnectionInfoResponseBodyData {
	s.PrivateEndpoint = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) SetProtocol(v string) *DescribeOpenSearchConnectionInfoResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) SetPublicEndpoint(v *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) *DescribeOpenSearchConnectionInfoResponseBodyData {
	s.PublicEndpoint = v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyData) Validate() error {
	if s.DashboardEndpoint != nil {
		if err := s.DashboardEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.DashboardPublicEndpoint != nil {
		if err := s.DashboardPublicEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateEndpoint != nil {
		if err := s.PrivateEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.PublicEndpoint != nil {
		if err := s.PublicEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint struct {
	// Specifies whether static frame check is enabled. Default value: false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The host address.
	//
	// example:
	//
	// 100.118.102.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The URL.
	//
	// example:
	//
	// https://static.yipigai.cn/timuocr/tmp_c29e30497575a40193a24a7a83654e30e21b951cc6856cdb.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) GetHost() *string {
	return s.Host
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) GetPort() *int32 {
	return s.Port
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) GetUrl() *string {
	return s.Url
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) SetEnabled(v bool) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint {
	s.Enabled = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) SetHost(v string) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint {
	s.Host = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) SetPort(v int32) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) SetUrl(v string) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint {
	s.Url = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint struct {
	// The service activation status. Valid values:
	//
	// - **on**: Activated.
	//
	// - **off**: Not activated.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The hostname. Retrieves data under the specified host.
	//
	// example:
	//
	// https://secnet-defense-vastip.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The URL.
	//
	// example:
	//
	// https://static.yipigai.cn/timuocr/tmp_c29e30497575a40193a24a7a83654e30e21b951cc6856cdb.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) GetHost() *string {
	return s.Host
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) GetPort() *int32 {
	return s.Port
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) GetUrl() *string {
	return s.Url
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) SetEnabled(v bool) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint {
	s.Enabled = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) SetHost(v string) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint {
	s.Host = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) SetPort(v int32) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) SetUrl(v string) *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint {
	s.Url = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataDashboardPublicEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint struct {
	// Specifies whether to enable the echo feature. This parameter is required. Valid values: true/false.
	//
	// example:
	//
	// True
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The OSS domain name.
	//
	// example:
	//
	// 100.118.214.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) GetHost() *string {
	return s.Host
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) GetPort() *int32 {
	return s.Port
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) SetEnabled(v bool) *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint {
	s.Enabled = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) SetHost(v string) *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint {
	s.Host = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) SetPort(v int32) *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPrivateEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint struct {
	// Specifies whether to enable dead-letter message delivery.
	//
	// example:
	//
	// True
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The machine.
	//
	// example:
	//
	// 100.98.83.0/24
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) GetHost() *string {
	return s.Host
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) GetPort() *int32 {
	return s.Port
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) SetEnabled(v bool) *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint {
	s.Enabled = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) SetHost(v string) *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint {
	s.Host = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) SetPort(v int32) *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoResponseBodyDataPublicEndpoint) Validate() error {
	return dara.Validate(s)
}
