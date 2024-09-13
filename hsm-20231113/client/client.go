// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConfigAuditLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// enable
	AuditAction *string `json:"AuditAction,omitempty" xml:"AuditAction,omitempty"`
	// example:
	//
	// hsm-log
	AuditOssBucket *string `json:"AuditOssBucket,omitempty" xml:"AuditOssBucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigAuditLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigAuditLogRequest) GoString() string {
	return s.String()
}

func (s *ConfigAuditLogRequest) SetAuditAction(v string) *ConfigAuditLogRequest {
	s.AuditAction = &v
	return s
}

func (s *ConfigAuditLogRequest) SetAuditOssBucket(v string) *ConfigAuditLogRequest {
	s.AuditOssBucket = &v
	return s
}

func (s *ConfigAuditLogRequest) SetRegionId(v string) *ConfigAuditLogRequest {
	s.RegionId = &v
	return s
}

type ConfigAuditLogResponseBody struct {
	// example:
	//
	// 42B118FB-16A6-56FB-B877-D58637EEC6AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigAuditLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAuditLogResponseBody) SetRequestId(v string) *ConfigAuditLogResponseBody {
	s.RequestId = &v
	return s
}

type ConfigAuditLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigAuditLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigAuditLogResponse) GoString() string {
	return s.String()
}

func (s *ConfigAuditLogResponse) SetHeaders(v map[string]*string) *ConfigAuditLogResponse {
	s.Headers = v
	return s
}

func (s *ConfigAuditLogResponse) SetStatusCode(v int32) *ConfigAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigAuditLogResponse) SetBody(v *ConfigAuditLogResponseBody) *ConfigAuditLogResponse {
	s.Body = v
	return s
}

type ConfigBackupRemarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// backup-fdb897sdfg****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// backup-test
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ConfigBackupRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigBackupRemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigBackupRemarkRequest) SetBackupId(v string) *ConfigBackupRemarkRequest {
	s.BackupId = &v
	return s
}

func (s *ConfigBackupRemarkRequest) SetName(v string) *ConfigBackupRemarkRequest {
	s.Name = &v
	return s
}

func (s *ConfigBackupRemarkRequest) SetRemark(v string) *ConfigBackupRemarkRequest {
	s.Remark = &v
	return s
}

type ConfigBackupRemarkResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigBackupRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigBackupRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigBackupRemarkResponseBody) SetRequestId(v string) *ConfigBackupRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ConfigBackupRemarkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigBackupRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigBackupRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigBackupRemarkResponse) GoString() string {
	return s.String()
}

func (s *ConfigBackupRemarkResponse) SetHeaders(v map[string]*string) *ConfigBackupRemarkResponse {
	s.Headers = v
	return s
}

func (s *ConfigBackupRemarkResponse) SetStatusCode(v int32) *ConfigBackupRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigBackupRemarkResponse) SetBody(v *ConfigBackupRemarkResponseBody) *ConfigBackupRemarkResponse {
	s.Body = v
	return s
}

type ConfigBackupTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12
	BackupHourInDay *int64 `json:"BackupHourInDay,omitempty" xml:"BackupHourInDay,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// backup-173620705****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	BackupPeriod        *int64    `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	Manual2PeriodicList []*string `json:"Manual2PeriodicList,omitempty" xml:"Manual2PeriodicList,omitempty" type:"Repeated"`
	Periodic2ManualList []*string `json:"Periodic2ManualList,omitempty" xml:"Periodic2ManualList,omitempty" type:"Repeated"`
}

func (s ConfigBackupTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigBackupTaskRequest) GoString() string {
	return s.String()
}

func (s *ConfigBackupTaskRequest) SetBackupHourInDay(v int64) *ConfigBackupTaskRequest {
	s.BackupHourInDay = &v
	return s
}

func (s *ConfigBackupTaskRequest) SetBackupId(v string) *ConfigBackupTaskRequest {
	s.BackupId = &v
	return s
}

func (s *ConfigBackupTaskRequest) SetBackupPeriod(v int64) *ConfigBackupTaskRequest {
	s.BackupPeriod = &v
	return s
}

func (s *ConfigBackupTaskRequest) SetManual2PeriodicList(v []*string) *ConfigBackupTaskRequest {
	s.Manual2PeriodicList = v
	return s
}

func (s *ConfigBackupTaskRequest) SetPeriodic2ManualList(v []*string) *ConfigBackupTaskRequest {
	s.Periodic2ManualList = v
	return s
}

type ConfigBackupTaskResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigBackupTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigBackupTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigBackupTaskResponseBody) SetRequestId(v string) *ConfigBackupTaskResponseBody {
	s.RequestId = &v
	return s
}

type ConfigBackupTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigBackupTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigBackupTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigBackupTaskResponse) GoString() string {
	return s.String()
}

func (s *ConfigBackupTaskResponse) SetHeaders(v map[string]*string) *ConfigBackupTaskResponse {
	s.Headers = v
	return s
}

func (s *ConfigBackupTaskResponse) SetStatusCode(v int32) *ConfigBackupTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigBackupTaskResponse) SetBody(v *ConfigBackupTaskResponseBody) *ConfigBackupTaskResponse {
	s.Body = v
	return s
}

type ConfigClusterCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDaTCCAlECAQEwDQYJKoZIhvcNAQELBQAwVTELMAkGA1UEBhMCY24xCzAJBgNV
	//
	// BAgMAnpqMQswCQYDVQQHDAJoejEWMBQGA1UECgwNQWxpYmFiYSBDbG91ZDEUMBIG
	//
	// A1UECwwLU2VjQ2xvdWRIc20wHhcNMjQwNzAzM****-----END CERTIFICATE-----
	ClusterCertificate *string `json:"ClusterCertificate,omitempty" xml:"ClusterCertificate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsg****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDfTCCAmWgAwIBAgIJAMRqQMr5if66MA0GCSqGSIb3DQEBCwUAMFUxCzAJBgNV
	//
	// BAYTAmNuMQswCQYDVQQIDAJ6ajELMAkGA1UEBwwCaHoxFjAUBgNVBAoMDUFsaWJh
	//
	// YmEgQ2xvdWQxFDA****
	//
	// -----END CERTIFICATE-----
	IssuerCertificate *string `json:"IssuerCertificate,omitempty" xml:"IssuerCertificate,omitempty"`
}

func (s ConfigClusterCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterCertificateRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterCertificateRequest) SetClusterCertificate(v string) *ConfigClusterCertificateRequest {
	s.ClusterCertificate = &v
	return s
}

func (s *ConfigClusterCertificateRequest) SetClusterId(v string) *ConfigClusterCertificateRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterCertificateRequest) SetIssuerCertificate(v string) *ConfigClusterCertificateRequest {
	s.IssuerCertificate = &v
	return s
}

type ConfigClusterCertificateResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterCertificateResponseBody) SetRequestId(v string) *ConfigClusterCertificateResponseBody {
	s.RequestId = &v
	return s
}

type ConfigClusterCertificateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterCertificateResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterCertificateResponse) SetHeaders(v map[string]*string) *ConfigClusterCertificateResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterCertificateResponse) SetStatusCode(v int32) *ConfigClusterCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterCertificateResponse) SetBody(v *ConfigClusterCertificateResponseBody) *ConfigClusterCertificateResponse {
	s.Body = v
	return s
}

type ConfigClusterNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsgfaisdf****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cluster_on****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s ConfigClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterNameRequest) SetClusterId(v string) *ConfigClusterNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterNameRequest) SetClusterName(v string) *ConfigClusterNameRequest {
	s.ClusterName = &v
	return s
}

type ConfigClusterNameResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterNameResponseBody) SetRequestId(v string) *ConfigClusterNameResponseBody {
	s.RequestId = &v
	return s
}

type ConfigClusterNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterNameResponse) SetHeaders(v map[string]*string) *ConfigClusterNameResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterNameResponse) SetStatusCode(v int32) *ConfigClusterNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterNameResponse) SetBody(v *ConfigClusterNameResponseBody) *ConfigClusterNameResponse {
	s.Body = v
	return s
}

type ConfigClusterSubnetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsg****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigClusterSubnetRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetRequest) SetClusterId(v string) *ConfigClusterSubnetRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterSubnetRequest) SetRegionId(v string) *ConfigClusterSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigClusterSubnetRequest) SetVSwitchIds(v []*string) *ConfigClusterSubnetRequest {
	s.VSwitchIds = v
	return s
}

func (s *ConfigClusterSubnetRequest) SetVpcId(v string) *ConfigClusterSubnetRequest {
	s.VpcId = &v
	return s
}

type ConfigClusterSubnetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsg****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigClusterSubnetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetShrinkRequest) SetClusterId(v string) *ConfigClusterSubnetShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetRegionId(v string) *ConfigClusterSubnetShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetVSwitchIdsShrink(v string) *ConfigClusterSubnetShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetVpcId(v string) *ConfigClusterSubnetShrinkRequest {
	s.VpcId = &v
	return s
}

type ConfigClusterSubnetResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetResponseBody) SetRequestId(v string) *ConfigClusterSubnetResponseBody {
	s.RequestId = &v
	return s
}

type ConfigClusterSubnetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterSubnetResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetResponse) SetHeaders(v map[string]*string) *ConfigClusterSubnetResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterSubnetResponse) SetStatusCode(v int32) *ConfigClusterSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterSubnetResponse) SetBody(v *ConfigClusterSubnetResponseBody) *ConfigClusterSubnetResponse {
	s.Body = v
	return s
}

type ConfigClusterWhitelistRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-729dm40FG****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18.68.XX.XX
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s ConfigClusterWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterWhitelistRequest) SetClusterId(v string) *ConfigClusterWhitelistRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterWhitelistRequest) SetWhitelist(v string) *ConfigClusterWhitelistRequest {
	s.Whitelist = &v
	return s
}

type ConfigClusterWhitelistResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterWhitelistResponseBody) SetRequestId(v string) *ConfigClusterWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ConfigClusterWhitelistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterWhitelistResponse) SetHeaders(v map[string]*string) *ConfigClusterWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterWhitelistResponse) SetStatusCode(v int32) *ConfigClusterWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterWhitelistResponse) SetBody(v *ConfigClusterWhitelistResponseBody) *ConfigClusterWhitelistResponse {
	s.Body = v
	return s
}

type ConfigImageRemarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// image-d79x4k11pmg19****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-****
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ConfigImageRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigImageRemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigImageRemarkRequest) SetImageId(v string) *ConfigImageRemarkRequest {
	s.ImageId = &v
	return s
}

func (s *ConfigImageRemarkRequest) SetRemark(v string) *ConfigImageRemarkRequest {
	s.Remark = &v
	return s
}

type ConfigImageRemarkResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigImageRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigImageRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigImageRemarkResponseBody) SetRequestId(v string) *ConfigImageRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ConfigImageRemarkResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigImageRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigImageRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigImageRemarkResponse) GoString() string {
	return s.String()
}

func (s *ConfigImageRemarkResponse) SetHeaders(v map[string]*string) *ConfigImageRemarkResponse {
	s.Headers = v
	return s
}

func (s *ConfigImageRemarkResponse) SetStatusCode(v int32) *ConfigImageRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigImageRemarkResponse) SetBody(v *ConfigImageRemarkResponseBody) *ConfigImageRemarkResponse {
	s.Body = v
	return s
}

type ConfigInstanceIpAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-u7gb0qahu****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-lmkmivmo6****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigInstanceIpAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceIpAddressRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceIpAddressRequest) SetInstanceId(v string) *ConfigInstanceIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetIp(v string) *ConfigInstanceIpAddressRequest {
	s.Ip = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetRegionId(v string) *ConfigInstanceIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetVSwitchId(v string) *ConfigInstanceIpAddressRequest {
	s.VSwitchId = &v
	return s
}

func (s *ConfigInstanceIpAddressRequest) SetVpcId(v string) *ConfigInstanceIpAddressRequest {
	s.VpcId = &v
	return s
}

type ConfigInstanceIpAddressResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceIpAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceIpAddressResponseBody) SetRequestId(v string) *ConfigInstanceIpAddressResponseBody {
	s.RequestId = &v
	return s
}

type ConfigInstanceIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceIpAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceIpAddressResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceIpAddressResponse) SetHeaders(v map[string]*string) *ConfigInstanceIpAddressResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceIpAddressResponse) SetStatusCode(v int32) *ConfigInstanceIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceIpAddressResponse) SetBody(v *ConfigInstanceIpAddressResponseBody) *ConfigInstanceIpAddressResponse {
	s.Body = v
	return s
}

type ConfigInstanceRemarkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsmOnline
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ConfigInstanceRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceRemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceRemarkRequest) SetInstanceId(v string) *ConfigInstanceRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceRemarkRequest) SetRemark(v string) *ConfigInstanceRemarkRequest {
	s.Remark = &v
	return s
}

type ConfigInstanceRemarkResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceRemarkResponseBody) SetRequestId(v string) *ConfigInstanceRemarkResponseBody {
	s.RequestId = &v
	return s
}

type ConfigInstanceRemarkResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceRemarkResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceRemarkResponse) SetHeaders(v map[string]*string) *ConfigInstanceRemarkResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceRemarkResponse) SetStatusCode(v int32) *ConfigInstanceRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceRemarkResponse) SetBody(v *ConfigInstanceRemarkResponseBody) *ConfigInstanceRemarkResponse {
	s.Body = v
	return s
}

type ConfigInstanceWhitelistRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18.68.XX.XX,18.68.XX.XX
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s ConfigInstanceWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhitelistRequest) SetInstanceId(v string) *ConfigInstanceWhitelistRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceWhitelistRequest) SetWhitelist(v string) *ConfigInstanceWhitelistRequest {
	s.Whitelist = &v
	return s
}

type ConfigInstanceWhitelistResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhitelistResponseBody) SetRequestId(v string) *ConfigInstanceWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ConfigInstanceWhitelistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhitelistResponse) SetHeaders(v map[string]*string) *ConfigInstanceWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceWhitelistResponse) SetStatusCode(v int32) *ConfigInstanceWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceWhitelistResponse) SetBody(v *ConfigInstanceWhitelistResponseBody) *ConfigInstanceWhitelistResponse {
	s.Body = v
	return s
}

type CopyImageRequest struct {
	// example:
	//
	// image-hafiudfahdd****
	ImageUid *string `json:"ImageUid,omitempty" xml:"ImageUid,omitempty"`
	// example:
	//
	// cn-beijing
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s CopyImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyImageRequest) GoString() string {
	return s.String()
}

func (s *CopyImageRequest) SetImageUid(v string) *CopyImageRequest {
	s.ImageUid = &v
	return s
}

func (s *CopyImageRequest) SetTargetRegionId(v string) *CopyImageRequest {
	s.TargetRegionId = &v
	return s
}

type CopyImageResponseBody struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1724379766191
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-202401250936hze747fd7e0007005
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CopyImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CopyImageResponseBody) SetCompleted(v bool) *CopyImageResponseBody {
	s.Completed = &v
	return s
}

func (s *CopyImageResponseBody) SetCreateTime(v string) *CopyImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CopyImageResponseBody) SetError(v string) *CopyImageResponseBody {
	s.Error = &v
	return s
}

func (s *CopyImageResponseBody) SetJobId(v string) *CopyImageResponseBody {
	s.JobId = &v
	return s
}

func (s *CopyImageResponseBody) SetProgress(v int32) *CopyImageResponseBody {
	s.Progress = &v
	return s
}

func (s *CopyImageResponseBody) SetRequestId(v string) *CopyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyImageResponseBody) SetResponse(v string) *CopyImageResponseBody {
	s.Response = &v
	return s
}

func (s *CopyImageResponseBody) SetStatus(v string) *CopyImageResponseBody {
	s.Status = &v
	return s
}

func (s *CopyImageResponseBody) SetType(v string) *CopyImageResponseBody {
	s.Type = &v
	return s
}

type CopyImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyImageResponse) GoString() string {
	return s.String()
}

func (s *CopyImageResponse) SetHeaders(v map[string]*string) *CopyImageResponse {
	s.Headers = v
	return s
}

func (s *CopyImageResponse) SetStatusCode(v int32) *CopyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyImageResponse) SetBody(v *CopyImageResponseBody) *CopyImageResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster_on****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm_intl-sg-uz63ixak****
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceId(v string) *CreateClusterRequest {
	s.MasterInstanceId = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

type CreateClusterResponseBody struct {
	// example:
	//
	// cluster-729dm40FG****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 5F58413E-8F57-585B-BE48-64CC1E31133C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-NZB9Oj5Yfd8Y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteClusterResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones    []*DescribeRegionsResponseBodyRegionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*DescribeRegionsResponseBodyRegionsZones) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsZones struct {
	// example:
	//
	// yes
	Cluster   *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetCluster(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.Cluster = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type EnableBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// backup-1736207****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupRequest) GoString() string {
	return s.String()
}

func (s *EnableBackupRequest) SetBackupId(v string) *EnableBackupRequest {
	s.BackupId = &v
	return s
}

func (s *EnableBackupRequest) SetInstanceId(v string) *EnableBackupRequest {
	s.InstanceId = &v
	return s
}

type EnableBackupResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupResponseBody) GoString() string {
	return s.String()
}

func (s *EnableBackupResponseBody) SetRequestId(v string) *EnableBackupResponseBody {
	s.RequestId = &v
	return s
}

type EnableBackupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableBackupResponse) GoString() string {
	return s.String()
}

func (s *EnableBackupResponse) SetHeaders(v map[string]*string) *EnableBackupResponse {
	s.Headers = v
	return s
}

func (s *EnableBackupResponse) SetStatusCode(v int32) *EnableBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableBackupResponse) SetBody(v *EnableBackupResponseBody) *EnableBackupResponse {
	s.Body = v
	return s
}

type ExportImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// image-8vbdd5uc6v10ecn5****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ExportImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportImageRequest) GoString() string {
	return s.String()
}

func (s *ExportImageRequest) SetImageId(v string) *ExportImageRequest {
	s.ImageId = &v
	return s
}

func (s *ExportImageRequest) SetInstanceId(v string) *ExportImageRequest {
	s.InstanceId = &v
	return s
}

type ExportImageResponseBody struct {
	Job *ExportImageResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExportImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportImageResponseBody) GoString() string {
	return s.String()
}

func (s *ExportImageResponseBody) SetJob(v *ExportImageResponseBodyJob) *ExportImageResponseBody {
	s.Job = v
	return s
}

func (s *ExportImageResponseBody) SetRequestId(v string) *ExportImageResponseBody {
	s.RequestId = &v
	return s
}

type ExportImageResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// Job.Canceled
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// b1748ca6-6b55-49f4-a6d4-2d694a9f3693
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 100
	Process *int32 `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ExportImageResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s ExportImageResponseBodyJob) GoString() string {
	return s.String()
}

func (s *ExportImageResponseBodyJob) SetCompleted(v bool) *ExportImageResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *ExportImageResponseBodyJob) SetError(v string) *ExportImageResponseBodyJob {
	s.Error = &v
	return s
}

func (s *ExportImageResponseBodyJob) SetJobId(v string) *ExportImageResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *ExportImageResponseBodyJob) SetProcess(v int32) *ExportImageResponseBodyJob {
	s.Process = &v
	return s
}

func (s *ExportImageResponseBodyJob) SetResponse(v string) *ExportImageResponseBodyJob {
	s.Response = &v
	return s
}

func (s *ExportImageResponseBodyJob) SetStatus(v string) *ExportImageResponseBodyJob {
	s.Status = &v
	return s
}

func (s *ExportImageResponseBodyJob) SetType(v string) *ExportImageResponseBodyJob {
	s.Type = &v
	return s
}

type ExportImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportImageResponse) GoString() string {
	return s.String()
}

func (s *ExportImageResponse) SetHeaders(v map[string]*string) *ExportImageResponse {
	s.Headers = v
	return s
}

func (s *ExportImageResponse) SetStatusCode(v int32) *ExportImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportImageResponse) SetBody(v *ExportImageResponseBody) *ExportImageResponse {
	s.Body = v
	return s
}

type GetAuditLogStatusRequest struct {
	// example:
	//
	// true
	GetOssBucket *bool `json:"GetOssBucket,omitempty" xml:"GetOssBucket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAuditLogStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuditLogStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAuditLogStatusRequest) SetGetOssBucket(v bool) *GetAuditLogStatusRequest {
	s.GetOssBucket = &v
	return s
}

func (s *GetAuditLogStatusRequest) SetRegionId(v string) *GetAuditLogStatusRequest {
	s.RegionId = &v
	return s
}

type GetAuditLogStatusResponseBody struct {
	// example:
	//
	// enable
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	// example:
	//
	// bucket-test
	AuditOssBucket *string `json:"AuditOssBucket,omitempty" xml:"AuditOssBucket,omitempty"`
	// example:
	//
	// true
	GrantedServiceAccess *bool     `json:"GrantedServiceAccess,omitempty" xml:"GrantedServiceAccess,omitempty"`
	OssBuckets           []*string `json:"OssBuckets,omitempty" xml:"OssBuckets,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuditLogStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuditLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditLogStatusResponseBody) SetAuditLogStatus(v string) *GetAuditLogStatusResponseBody {
	s.AuditLogStatus = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetAuditOssBucket(v string) *GetAuditLogStatusResponseBody {
	s.AuditOssBucket = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetGrantedServiceAccess(v bool) *GetAuditLogStatusResponseBody {
	s.GrantedServiceAccess = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetOssBuckets(v []*string) *GetAuditLogStatusResponseBody {
	s.OssBuckets = v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetRegionId(v string) *GetAuditLogStatusResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetRequestId(v string) *GetAuditLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditLogStatusResponseBody) SetSuccess(v bool) *GetAuditLogStatusResponseBody {
	s.Success = &v
	return s
}

type GetAuditLogStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuditLogStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuditLogStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuditLogStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAuditLogStatusResponse) SetHeaders(v map[string]*string) *GetAuditLogStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAuditLogStatusResponse) SetStatusCode(v int32) *GetAuditLogStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuditLogStatusResponse) SetBody(v *GetAuditLogStatusResponseBody) *GetAuditLogStatusResponse {
	s.Body = v
	return s
}

type GetBackupRequest struct {
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s GetBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBackupRequest) GoString() string {
	return s.String()
}

func (s *GetBackupRequest) SetBackupId(v string) *GetBackupRequest {
	s.BackupId = &v
	return s
}

type GetBackupResponseBody struct {
	Backup *GetBackupResponseBodyBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBackupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupResponseBody) SetBackup(v *GetBackupResponseBodyBackup) *GetBackupResponseBody {
	s.Backup = v
	return s
}

func (s *GetBackupResponseBody) SetRequestId(v string) *GetBackupResponseBody {
	s.RequestId = &v
	return s
}

type GetBackupResponseBodyBackup struct {
	// example:
	//
	// 1
	AutoImageCount *int64 `json:"AutoImageCount,omitempty" xml:"AutoImageCount,omitempty"`
	// example:
	//
	// 10
	BackupHourInDay *string `json:"BackupHourInDay,omitempty" xml:"BackupHourInDay,omitempty"`
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// 3
	BackupPeriod *int64 `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// example:
	//
	// 1682417553781
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1682417553781
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// hsm-cn-5yd35431****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 3
	MaxImageCount *string `json:"MaxImageCount,omitempty" xml:"MaxImageCount,omitempty"`
	// example:
	//
	// backup-te****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1682417553781
	NextImageCreateTime *int64 `json:"NextImageCreateTime,omitempty" xml:"NextImageCreateTime,omitempty"`
	// example:
	//
	// hsm-cn-huoahd****
	OwnerInstanceId *string `json:"OwnerInstanceId,omitempty" xml:"OwnerInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1641275680000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// backup-fdb897sdfg53****
	SpInstanceId *string `json:"SpInstanceId,omitempty" xml:"SpInstanceId,omitempty"`
	// example:
	//
	// NEW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DEFAULT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBackupResponseBodyBackup) String() string {
	return tea.Prettify(s)
}

func (s GetBackupResponseBodyBackup) GoString() string {
	return s.String()
}

func (s *GetBackupResponseBodyBackup) SetAutoImageCount(v int64) *GetBackupResponseBodyBackup {
	s.AutoImageCount = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetBackupHourInDay(v string) *GetBackupResponseBodyBackup {
	s.BackupHourInDay = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetBackupId(v string) *GetBackupResponseBodyBackup {
	s.BackupId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetBackupPeriod(v int64) *GetBackupResponseBodyBackup {
	s.BackupPeriod = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetCreateTime(v int64) *GetBackupResponseBodyBackup {
	s.CreateTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetExpireTime(v int64) *GetBackupResponseBodyBackup {
	s.ExpireTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetInstanceId(v string) *GetBackupResponseBodyBackup {
	s.InstanceId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetMaxImageCount(v string) *GetBackupResponseBodyBackup {
	s.MaxImageCount = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetName(v string) *GetBackupResponseBodyBackup {
	s.Name = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetNextImageCreateTime(v int64) *GetBackupResponseBodyBackup {
	s.NextImageCreateTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetOwnerInstanceId(v string) *GetBackupResponseBodyBackup {
	s.OwnerInstanceId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetRegionId(v string) *GetBackupResponseBodyBackup {
	s.RegionId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetReleaseTime(v int64) *GetBackupResponseBodyBackup {
	s.ReleaseTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetRemark(v string) *GetBackupResponseBodyBackup {
	s.Remark = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetSpInstanceId(v string) *GetBackupResponseBodyBackup {
	s.SpInstanceId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetStatus(v string) *GetBackupResponseBodyBackup {
	s.Status = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetType(v string) *GetBackupResponseBodyBackup {
	s.Type = &v
	return s
}

type GetBackupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBackupResponse) GoString() string {
	return s.String()
}

func (s *GetBackupResponse) SetHeaders(v map[string]*string) *GetBackupResponse {
	s.Headers = v
	return s
}

func (s *GetBackupResponse) SetStatusCode(v int32) *GetBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupResponse) SetBody(v *GetBackupResponseBody) *GetBackupResponse {
	s.Body = v
	return s
}

type GetClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-p94y1dud9ts****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetClusterId(v string) *GetClusterRequest {
	s.ClusterId = &v
	return s
}

type GetClusterResponseBody struct {
	Cluster *GetClusterResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetCluster(v *GetClusterResponseBodyCluster) *GetClusterResponseBody {
	s.Cluster = v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

type GetClusterResponseBodyCluster struct {
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIDfTCCAmWgAwIBAgIJAMRqQMr5if66MA0GCSqGSIb3DQEBCwUAMFUxCzAJBgNV
	//
	// BAYTAmNuMQswCQYDVQQIDAJ6ajELMAkGA1UEBwwCaHoxFjAUBgNVBAoMDUFsaWJh
	//
	// YmEgQ2xvdWQxFDA****
	//
	// -----END CERTIFICATE-----
	ClusterCertificate *string `json:"ClusterCertificate,omitempty" xml:"ClusterCertificate,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----\\nMIIC5TCCAc0CAQAwgZ8xWTAJBgNVBAYTAlVTMAkGA1UECAwCQ0EwDQYDVQQKDAZD\\nYXZpdW0wDQYDVQQLDAZOM0ZJUFMwDgYDVQQHDAdTYW5Kb3NlMBMGA1UdEQwMMTk****
	//
	// -----END CERTIFICATE REQUEST-----
	ClusterCsr *string `json:"ClusterCsr,omitempty" xml:"ClusterCsr,omitempty"`
	// example:
	//
	// cluster-p94y1dud9ts****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// cluster_polar_****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// ----BEGIN CERTIFICATE-----
	//
	// MIIDaTCCAlECAQEwDQYJKoZIhvcNAQELBQAwVTELMAkGA1UEBhMCY24xCzAJBgNV
	//
	// BAgMAnpqMQswCQYDVQQHDAJoejEWMBQGA1UECgwNQWxpYmFiYSBDbG91ZDEUMBIG
	//
	// A1UECwwLU2VjQ2xvdWRIc20wHhcNMjQwNzAzM****
	//
	// -----END CERTIFICATE-----
	ClusterOwnerCertificate *string `json:"ClusterOwnerCertificate,omitempty" xml:"ClusterOwnerCertificate,omitempty"`
	// example:
	//
	// 1641275680000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// jnta
	DeviceType *string                                   `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Instances  []*GetClusterResponseBodyClusterInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// NEW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vpc-8vbt0fjdm29hofvbo****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 130.176.XX.XX
	Whitelist *string                               `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	Zones     []*GetClusterResponseBodyClusterZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s GetClusterResponseBodyCluster) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyCluster) SetClusterCertificate(v string) *GetClusterResponseBodyCluster {
	s.ClusterCertificate = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterCsr(v string) *GetClusterResponseBodyCluster {
	s.ClusterCsr = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterId(v string) *GetClusterResponseBodyCluster {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterName(v string) *GetClusterResponseBodyCluster {
	s.ClusterName = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterOwnerCertificate(v string) *GetClusterResponseBodyCluster {
	s.ClusterOwnerCertificate = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCreateTime(v int64) *GetClusterResponseBodyCluster {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetDeviceType(v string) *GetClusterResponseBodyCluster {
	s.DeviceType = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetInstances(v []*GetClusterResponseBodyClusterInstances) *GetClusterResponseBodyCluster {
	s.Instances = v
	return s
}

func (s *GetClusterResponseBodyCluster) SetRegionId(v string) *GetClusterResponseBodyCluster {
	s.RegionId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetSize(v int32) *GetClusterResponseBodyCluster {
	s.Size = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetStatus(v string) *GetClusterResponseBodyCluster {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetVpcId(v string) *GetClusterResponseBodyCluster {
	s.VpcId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetWhitelist(v string) *GetClusterResponseBodyCluster {
	s.Whitelist = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetZones(v []*GetClusterResponseBodyClusterZones) *GetClusterResponseBodyCluster {
	s.Zones = v
	return s
}

type GetClusterResponseBodyClusterInstances struct {
	// example:
	//
	// hsm-cn-g6z3v0uf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Master *bool `json:"Master,omitempty" xml:"Master,omitempty"`
	// example:
	//
	// 1
	NodeId *int32 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s GetClusterResponseBodyClusterInstances) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyClusterInstances) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyClusterInstances) SetInstanceId(v string) *GetClusterResponseBodyClusterInstances {
	s.InstanceId = &v
	return s
}

func (s *GetClusterResponseBodyClusterInstances) SetMaster(v bool) *GetClusterResponseBodyClusterInstances {
	s.Master = &v
	return s
}

func (s *GetClusterResponseBodyClusterInstances) SetNodeId(v int32) *GetClusterResponseBodyClusterInstances {
	s.NodeId = &v
	return s
}

type GetClusterResponseBodyClusterZones struct {
	// example:
	//
	// vsw-uf61s651p69bdgmki****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetClusterResponseBodyClusterZones) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyClusterZones) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyClusterZones) SetVSwitchId(v string) *GetClusterResponseBodyClusterZones {
	s.VSwitchId = &v
	return s
}

func (s *GetClusterResponseBodyClusterZones) SetZoneId(v string) *GetClusterResponseBodyClusterZones {
	s.ZoneId = &v
	return s
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// image-wz9c5ths5dfuwx47****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetImageId(v string) *GetImageRequest {
	s.ImageId = &v
	return s
}

type GetImageResponseBody struct {
	Image *GetImageResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody {
	s.Image = v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

type GetImageResponseBodyImage struct {
	// example:
	//
	// backup-1618017313
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// 1641275680000
	CopyTime *int64 `json:"CopyTime,omitempty" xml:"CopyTime,omitempty"`
	// example:
	//
	// 1786776567788
	ExportTime *int64 `json:"ExportTime,omitempty" xml:"ExportTime,omitempty"`
	// example:
	//
	// image-wz9c5ths5dfuwx47****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// hsm-cn-9lb32vll****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// hsm-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// backup-gfuiasdfa****
	SourceBackupUid *string `json:"SourceBackupUid,omitempty" xml:"SourceBackupUid,omitempty"`
	// example:
	//
	// image-kklhhhh****
	SourceImageUid *string `json:"SourceImageUid,omitempty" xml:"SourceImageUid,omitempty"`
	// example:
	//
	// hsm-wz9fnmvx190shfbk****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// example:
	//
	// NEW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3kGeHnmQzXwSsfF0Jk9eJYhe2gP6An0/HlYIiZh1****
	VsmDigest *string `json:"VsmDigest,omitempty" xml:"VsmDigest,omitempty"`
}

func (s GetImageResponseBodyImage) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImage) SetBackupId(v string) *GetImageResponseBodyImage {
	s.BackupId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetCopyTime(v int64) *GetImageResponseBodyImage {
	s.CopyTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetExportTime(v int64) *GetImageResponseBodyImage {
	s.ExportTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetImageId(v string) *GetImageResponseBodyImage {
	s.ImageId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetInstanceId(v string) *GetImageResponseBodyImage {
	s.InstanceId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetMode(v string) *GetImageResponseBodyImage {
	s.Mode = &v
	return s
}

func (s *GetImageResponseBodyImage) SetRegionId(v string) *GetImageResponseBodyImage {
	s.RegionId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetRemark(v string) *GetImageResponseBodyImage {
	s.Remark = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceBackupUid(v string) *GetImageResponseBodyImage {
	s.SourceBackupUid = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceImageUid(v string) *GetImageResponseBodyImage {
	s.SourceImageUid = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceInstanceId(v string) *GetImageResponseBodyImage {
	s.SourceInstanceId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSourceRegionId(v string) *GetImageResponseBodyImage {
	s.SourceRegionId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetStatus(v string) *GetImageResponseBodyImage {
	s.Status = &v
	return s
}

func (s *GetImageResponseBodyImage) SetVsmDigest(v string) *GetImageResponseBodyImage {
	s.VsmDigest = &v
	return s
}

type GetImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

func (s *GetImageResponse) SetStatusCode(v int32) *GetImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	// example:
	//
	// hsm-cn-vj30bil****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetInstanceId(v string) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

type GetInstanceResponseBody struct {
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceResponseBodyInstance struct {
	// example:
	//
	// cluster-w3G9vOJI2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// cluster_online
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 1699515963000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// jnta.SJJ1528-G
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1699496389720
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// hsm-cn-g4t3jwsc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10.192.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// false
	IsTrial *bool `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	// example:
	//
	// true
	Master *bool `json:"Master,omitempty" xml:"Master,omitempty"`
	// example:
	//
	// 23576634952****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// hsmOnline
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// EXPIRED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-bp1mvfs31ltt0wyhf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// jnta
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// example:
	//
	// vpc-uf69i66j9kmoko52p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// 18.68.XX.XX
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) SetClusterId(v string) *GetInstanceResponseBodyInstance {
	s.ClusterId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetClusterName(v string) *GetInstanceResponseBodyInstance {
	s.ClusterName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCreateTime(v int64) *GetInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDeviceType(v string) *GetInstanceResponseBodyInstance {
	s.DeviceType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetExpireTime(v int64) *GetInstanceResponseBodyInstance {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetIp(v string) *GetInstanceResponseBodyInstance {
	s.Ip = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetIsTrial(v bool) *GetInstanceResponseBodyInstance {
	s.IsTrial = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetMaster(v bool) *GetInstanceResponseBodyInstance {
	s.Master = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetOrderId(v string) *GetInstanceResponseBodyInstance {
	s.OrderId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetRegionId(v string) *GetInstanceResponseBodyInstance {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetRemark(v string) *GetInstanceResponseBodyInstance {
	s.Remark = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetStatus(v string) *GetInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVSwitchId(v string) *GetInstanceResponseBodyInstance {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVendor(v string) *GetInstanceResponseBodyInstance {
	s.Vendor = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVpcId(v string) *GetInstanceResponseBodyInstance {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetWhitelist(v string) *GetInstanceResponseBodyInstance {
	s.Whitelist = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetZoneId(v string) *GetInstanceResponseBodyInstance {
	s.ZoneId = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-202401250936hze747fd7e0007005
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetJobId(v string) *GetJobRequest {
	s.JobId = &v
	return s
}

type GetJobResponseBody struct {
	Job *GetJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetJob(v *GetJobResponseBodyJob) *GetJobResponseBody {
	s.Job = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

type GetJobResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-202401250936hze747fd7e0007005
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 95
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJob) SetCompleted(v bool) *GetJobResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *GetJobResponseBodyJob) SetError(v string) *GetJobResponseBodyJob {
	s.Error = &v
	return s
}

func (s *GetJobResponseBodyJob) SetJobId(v string) *GetJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJob) SetProgress(v int64) *GetJobResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *GetJobResponseBodyJob) SetResponse(v string) *GetJobResponseBodyJob {
	s.Response = &v
	return s
}

func (s *GetJobResponseBodyJob) SetStatus(v string) *GetJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJob) SetType(v string) *GetJobResponseBodyJob {
	s.Type = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type InitializeAuditLogResponseBody struct {
	// example:
	//
	// 4FE969D9-E1C7-5274-BE7D-8C3534587605
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeAuditLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeAuditLogResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeAuditLogResponseBody) SetRequestId(v string) *InitializeAuditLogResponseBody {
	s.RequestId = &v
	return s
}

type InitializeAuditLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeAuditLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeAuditLogResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeAuditLogResponse) GoString() string {
	return s.String()
}

func (s *InitializeAuditLogResponse) SetHeaders(v map[string]*string) *InitializeAuditLogResponse {
	s.Headers = v
	return s
}

func (s *InitializeAuditLogResponse) SetStatusCode(v int32) *InitializeAuditLogResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeAuditLogResponse) SetBody(v *InitializeAuditLogResponseBody) *InitializeAuditLogResponse {
	s.Body = v
	return s
}

type InitializeClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-p94y1dud9****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InitializeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeClusterRequest) GoString() string {
	return s.String()
}

func (s *InitializeClusterRequest) SetClusterId(v string) *InitializeClusterRequest {
	s.ClusterId = &v
	return s
}

type InitializeClusterResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeClusterResponseBody) SetRequestId(v string) *InitializeClusterResponseBody {
	s.RequestId = &v
	return s
}

type InitializeClusterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeClusterResponse) GoString() string {
	return s.String()
}

func (s *InitializeClusterResponse) SetHeaders(v map[string]*string) *InitializeClusterResponse {
	s.Headers = v
	return s
}

func (s *InitializeClusterResponse) SetStatusCode(v int32) *InitializeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeClusterResponse) SetBody(v *InitializeClusterResponseBody) *InitializeClusterResponse {
	s.Body = v
	return s
}

type JoinClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-NZB9Oj5Yfd8Y****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s JoinClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinClusterRequest) GoString() string {
	return s.String()
}

func (s *JoinClusterRequest) SetClusterId(v string) *JoinClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *JoinClusterRequest) SetInstanceId(v string) *JoinClusterRequest {
	s.InstanceId = &v
	return s
}

type JoinClusterResponseBody struct {
	Job *JoinClusterResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinClusterResponseBody) GoString() string {
	return s.String()
}

func (s *JoinClusterResponseBody) SetJob(v *JoinClusterResponseBodyJob) *JoinClusterResponseBody {
	s.Job = v
	return s
}

func (s *JoinClusterResponseBody) SetRequestId(v string) *JoinClusterResponseBody {
	s.RequestId = &v
	return s
}

type JoinClusterResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1711764127000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-202401250936hze747fd7e0007005
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 86
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s JoinClusterResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s JoinClusterResponseBodyJob) GoString() string {
	return s.String()
}

func (s *JoinClusterResponseBodyJob) SetCompleted(v bool) *JoinClusterResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetCreateTime(v string) *JoinClusterResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetError(v string) *JoinClusterResponseBodyJob {
	s.Error = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetJobId(v string) *JoinClusterResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetProgress(v int32) *JoinClusterResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetResponse(v string) *JoinClusterResponseBodyJob {
	s.Response = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetStatus(v string) *JoinClusterResponseBodyJob {
	s.Status = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetType(v string) *JoinClusterResponseBodyJob {
	s.Type = &v
	return s
}

type JoinClusterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinClusterResponse) GoString() string {
	return s.String()
}

func (s *JoinClusterResponse) SetHeaders(v map[string]*string) *JoinClusterResponse {
	s.Headers = v
	return s
}

func (s *JoinClusterResponse) SetStatusCode(v int32) *JoinClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinClusterResponse) SetBody(v *JoinClusterResponseBody) *JoinClusterResponse {
	s.Body = v
	return s
}

type LeaveClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-729dm40FG****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s LeaveClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s LeaveClusterRequest) GoString() string {
	return s.String()
}

func (s *LeaveClusterRequest) SetClusterId(v string) *LeaveClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *LeaveClusterRequest) SetInstanceId(v string) *LeaveClusterRequest {
	s.InstanceId = &v
	return s
}

type LeaveClusterResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LeaveClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LeaveClusterResponseBody) GoString() string {
	return s.String()
}

func (s *LeaveClusterResponseBody) SetRequestId(v string) *LeaveClusterResponseBody {
	s.RequestId = &v
	return s
}

type LeaveClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LeaveClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LeaveClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s LeaveClusterResponse) GoString() string {
	return s.String()
}

func (s *LeaveClusterResponse) SetHeaders(v map[string]*string) *LeaveClusterResponse {
	s.Headers = v
	return s
}

func (s *LeaveClusterResponse) SetStatusCode(v int32) *LeaveClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *LeaveClusterResponse) SetBody(v *LeaveClusterResponseBody) *LeaveClusterResponse {
	s.Body = v
	return s
}

type ListBackupsRequest struct {
	// example:
	//
	// backup-1648438****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// hsm-te****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBackupsRequest) GoString() string {
	return s.String()
}

func (s *ListBackupsRequest) SetBackupId(v string) *ListBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *ListBackupsRequest) SetCurrentPage(v int64) *ListBackupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListBackupsRequest) SetInstanceId(v string) *ListBackupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBackupsRequest) SetName(v string) *ListBackupsRequest {
	s.Name = &v
	return s
}

func (s *ListBackupsRequest) SetPageSize(v int64) *ListBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBackupsRequest) SetRegionId(v string) *ListBackupsRequest {
	s.RegionId = &v
	return s
}

type ListBackupsResponseBody struct {
	Backups []*ListBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackupsResponseBody) SetBackups(v []*ListBackupsResponseBodyBackups) *ListBackupsResponseBody {
	s.Backups = v
	return s
}

func (s *ListBackupsResponseBody) SetCurrentPage(v int32) *ListBackupsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListBackupsResponseBody) SetPageSize(v int32) *ListBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBackupsResponseBody) SetRequestId(v string) *ListBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackupsResponseBody) SetTotalCount(v int32) *ListBackupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListBackupsResponseBodyBackups struct {
	// example:
	//
	// 1
	AutoImageCount *int64 `json:"AutoImageCount,omitempty" xml:"AutoImageCount,omitempty"`
	// example:
	//
	// 13
	BackupHourInDay *string `json:"BackupHourInDay,omitempty" xml:"BackupHourInDay,omitempty"`
	// example:
	//
	// backup-1648438****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// 3
	BackupPeriod *int64 `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// example:
	//
	// 1637229596000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1682417553781
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 3
	MaxImageCount *string `json:"MaxImageCount,omitempty" xml:"MaxImageCount,omitempty"`
	// example:
	//
	// backup-te****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1682417553781
	NextImageCreateTime *int64 `json:"NextImageCreateTime,omitempty" xml:"NextImageCreateTime,omitempty"`
	// example:
	//
	// hsm-cn-vj30bil8****
	OwnerInstanceId *string `json:"OwnerInstanceId,omitempty" xml:"OwnerInstanceId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1641275680000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// normal backup
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// backup-fdb897sdfg534-****
	SpInstanceId *string `json:"SpInstanceId,omitempty" xml:"SpInstanceId,omitempty"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBackupsResponseBodyBackups) String() string {
	return tea.Prettify(s)
}

func (s ListBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *ListBackupsResponseBodyBackups) SetAutoImageCount(v int64) *ListBackupsResponseBodyBackups {
	s.AutoImageCount = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetBackupHourInDay(v string) *ListBackupsResponseBodyBackups {
	s.BackupHourInDay = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetBackupId(v string) *ListBackupsResponseBodyBackups {
	s.BackupId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetBackupPeriod(v int64) *ListBackupsResponseBodyBackups {
	s.BackupPeriod = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetCreateTime(v int64) *ListBackupsResponseBodyBackups {
	s.CreateTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetExpireTime(v int64) *ListBackupsResponseBodyBackups {
	s.ExpireTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetInstanceId(v string) *ListBackupsResponseBodyBackups {
	s.InstanceId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetMaxImageCount(v string) *ListBackupsResponseBodyBackups {
	s.MaxImageCount = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetName(v string) *ListBackupsResponseBodyBackups {
	s.Name = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetNextImageCreateTime(v int64) *ListBackupsResponseBodyBackups {
	s.NextImageCreateTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetOwnerInstanceId(v string) *ListBackupsResponseBodyBackups {
	s.OwnerInstanceId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetRegionId(v string) *ListBackupsResponseBodyBackups {
	s.RegionId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetReleaseTime(v int64) *ListBackupsResponseBodyBackups {
	s.ReleaseTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetRemark(v string) *ListBackupsResponseBodyBackups {
	s.Remark = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetSpInstanceId(v string) *ListBackupsResponseBodyBackups {
	s.SpInstanceId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetStatus(v string) *ListBackupsResponseBodyBackups {
	s.Status = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetType(v string) *ListBackupsResponseBodyBackups {
	s.Type = &v
	return s
}

type ListBackupsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBackupsResponse) GoString() string {
	return s.String()
}

func (s *ListBackupsResponse) SetHeaders(v map[string]*string) *ListBackupsResponse {
	s.Headers = v
	return s
}

func (s *ListBackupsResponse) SetStatusCode(v int32) *ListBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBackupsResponse) SetBody(v *ListBackupsResponseBody) *ListBackupsResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetCurrentPage(v int32) *ListClustersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

type ListClustersResponseBody struct {
	Clusters []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 114
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetCurrentPage(v int32) *ListClustersResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotal(v int32) *ListClustersResponseBody {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	// example:
	//
	// cluster-w3G9vOJI2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// INITIALIZED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetStatus(v string) *ListClustersResponseBodyClusters {
	s.Status = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetBackupId(v string) *ListImagesRequest {
	s.BackupId = &v
	return s
}

func (s *ListImagesRequest) SetCurrentPage(v int32) *ListImagesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesRequest) SetMode(v string) *ListImagesRequest {
	s.Mode = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int32) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) SetRegionId(v string) *ListImagesRequest {
	s.RegionId = &v
	return s
}

type ListImagesResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                          `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Images      []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetCurrentPage(v int32) *ListImagesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetPageSize(v int32) *ListImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int32) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListImagesResponseBodyImages struct {
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// 1641275680000
	CopyTime *string `json:"CopyTime,omitempty" xml:"CopyTime,omitempty"`
	// example:
	//
	// 1782849566738
	ExportTime *int64 `json:"ExportTime,omitempty" xml:"ExportTime,omitempty"`
	// example:
	//
	// image-d79x4k11pmg19****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// hsm-cn-6ja1xknf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// MANUAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// hsm-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// backup-hodfhaol****
	SourceBackupUid *string `json:"SourceBackupUid,omitempty" xml:"SourceBackupUid,omitempty"`
	// example:
	//
	// image-ooopjygsn****
	SourceImageUid *string `json:"SourceImageUid,omitempty" xml:"SourceImageUid,omitempty"`
	// example:
	//
	// hsm-cn-wz9i2dmefudfxtmb****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// example:
	//
	// cn-shanghai
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3kGeHnmQzXwSsfF0Jk9eJYhe2gP6An0/HlYIiZh1****
	VsmDigest *string `json:"VsmDigest,omitempty" xml:"VsmDigest,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetBackupId(v string) *ListImagesResponseBodyImages {
	s.BackupId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetCopyTime(v string) *ListImagesResponseBodyImages {
	s.CopyTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetExportTime(v int64) *ListImagesResponseBodyImages {
	s.ExportTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetInstanceId(v string) *ListImagesResponseBodyImages {
	s.InstanceId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetMode(v string) *ListImagesResponseBodyImages {
	s.Mode = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRegionId(v string) *ListImagesResponseBodyImages {
	s.RegionId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetRemark(v string) *ListImagesResponseBodyImages {
	s.Remark = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceBackupUid(v string) *ListImagesResponseBodyImages {
	s.SourceBackupUid = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceImageUid(v string) *ListImagesResponseBodyImages {
	s.SourceImageUid = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceInstanceId(v string) *ListImagesResponseBodyImages {
	s.SourceInstanceId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSourceRegionId(v string) *ListImagesResponseBodyImages {
	s.SourceRegionId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetStatus(v string) *ListImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetVsmDigest(v string) *ListImagesResponseBodyImages {
	s.VsmDigest = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetCurrentPage(v int32) *ListInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

type ListInstancesResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Instances   []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 80
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCurrentPage(v int32) *ListInstancesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int32) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotal(v int32) *ListInstancesResponseBody {
	s.Total = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rg-aek2tsvbnfe****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-2ze0qae64mjuc0ni****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type PauseInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PauseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceRequest) GoString() string {
	return s.String()
}

func (s *PauseInstanceRequest) SetInstanceId(v string) *PauseInstanceRequest {
	s.InstanceId = &v
	return s
}

type PauseInstanceResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponseBody) SetRequestId(v string) *PauseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type PauseInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseInstanceResponse) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponse) SetHeaders(v map[string]*string) *PauseInstanceResponse {
	s.Headers = v
	return s
}

func (s *PauseInstanceResponse) SetStatusCode(v int32) *PauseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseInstanceResponse) SetBody(v *PauseInstanceResponseBody) *PauseInstanceResponse {
	s.Body = v
	return s
}

type QuickInitInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QuickInitInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QuickInitInstanceRequest) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceRequest) SetInstanceId(v string) *QuickInitInstanceRequest {
	s.InstanceId = &v
	return s
}

type QuickInitInstanceResponseBody struct {
	Job *QuickInitInstanceResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuickInitInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuickInitInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceResponseBody) SetJob(v *QuickInitInstanceResponseBodyJob) *QuickInitInstanceResponseBody {
	s.Job = v
	return s
}

func (s *QuickInitInstanceResponseBody) SetRequestId(v string) *QuickInitInstanceResponseBody {
	s.RequestId = &v
	return s
}

type QuickInitInstanceResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1699515963000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-000fi9k1v2hclo321sal
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuickInitInstanceResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s QuickInitInstanceResponseBodyJob) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceResponseBodyJob) SetCompleted(v bool) *QuickInitInstanceResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetCreateTime(v string) *QuickInitInstanceResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetError(v string) *QuickInitInstanceResponseBodyJob {
	s.Error = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetJobId(v string) *QuickInitInstanceResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetProgress(v int32) *QuickInitInstanceResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetResponse(v string) *QuickInitInstanceResponseBodyJob {
	s.Response = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetStatus(v string) *QuickInitInstanceResponseBodyJob {
	s.Status = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetType(v string) *QuickInitInstanceResponseBodyJob {
	s.Type = &v
	return s
}

type QuickInitInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuickInitInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuickInitInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QuickInitInstanceResponse) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceResponse) SetHeaders(v map[string]*string) *QuickInitInstanceResponse {
	s.Headers = v
	return s
}

func (s *QuickInitInstanceResponse) SetStatusCode(v int32) *QuickInitInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QuickInitInstanceResponse) SetBody(v *QuickInitInstanceResponseBody) *QuickInitInstanceResponse {
	s.Body = v
	return s
}

type ResetBackupRequest struct {
	// example:
	//
	// backup-fdb897sdfg5****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s ResetBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetBackupRequest) GoString() string {
	return s.String()
}

func (s *ResetBackupRequest) SetBackupId(v string) *ResetBackupRequest {
	s.BackupId = &v
	return s
}

type ResetBackupResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetBackupResponseBody) GoString() string {
	return s.String()
}

func (s *ResetBackupResponseBody) SetRequestId(v string) *ResetBackupResponseBody {
	s.RequestId = &v
	return s
}

type ResetBackupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetBackupResponse) GoString() string {
	return s.String()
}

func (s *ResetBackupResponse) SetHeaders(v map[string]*string) *ResetBackupResponse {
	s.Headers = v
	return s
}

func (s *ResetBackupResponse) SetStatusCode(v int32) *ResetBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetBackupResponse) SetBody(v *ResetBackupResponseBody) *ResetBackupResponse {
	s.Body = v
	return s
}

type ResetInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ResetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResetInstanceRequest) SetInstanceId(v string) *ResetInstanceRequest {
	s.InstanceId = &v
	return s
}

type ResetInstanceResponseBody struct {
	Job *ResetInstanceResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponseBody) SetJob(v *ResetInstanceResponseBodyJob) *ResetInstanceResponseBody {
	s.Job = v
	return s
}

func (s *ResetInstanceResponseBody) SetRequestId(v string) *ResetInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ResetInstanceResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1653274407000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-0007bl8oev0u3jqyfu6a
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 80
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ResetInstanceResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s ResetInstanceResponseBodyJob) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponseBodyJob) SetCompleted(v bool) *ResetInstanceResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetCreateTime(v string) *ResetInstanceResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetError(v string) *ResetInstanceResponseBodyJob {
	s.Error = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetJobId(v string) *ResetInstanceResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetProgress(v int32) *ResetInstanceResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetResponse(v string) *ResetInstanceResponseBodyJob {
	s.Response = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetStatus(v string) *ResetInstanceResponseBodyJob {
	s.Status = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetType(v string) *ResetInstanceResponseBodyJob {
	s.Type = &v
	return s
}

type ResetInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponse) SetHeaders(v map[string]*string) *ResetInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResetInstanceResponse) SetStatusCode(v int32) *ResetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetInstanceResponse) SetBody(v *ResetInstanceResponseBody) *ResetInstanceResponse {
	s.Body = v
	return s
}

type RestoreInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// image-eaOGHkRDQgh4****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-mp90fxef****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestoreInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestoreInstanceRequest) SetImageId(v string) *RestoreInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *RestoreInstanceRequest) SetInstanceId(v string) *RestoreInstanceRequest {
	s.InstanceId = &v
	return s
}

type RestoreInstanceResponseBody struct {
	Job *RestoreInstanceResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBody) SetJob(v *RestoreInstanceResponseBodyJob) *RestoreInstanceResponseBody {
	s.Job = v
	return s
}

func (s *RestoreInstanceResponseBody) SetRequestId(v string) *RestoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestoreInstanceResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1711764127000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-540356379023708160
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 50
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RestoreInstanceResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s RestoreInstanceResponseBodyJob) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBodyJob) SetCompleted(v bool) *RestoreInstanceResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetCreateTime(v string) *RestoreInstanceResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetError(v string) *RestoreInstanceResponseBodyJob {
	s.Error = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetJobId(v string) *RestoreInstanceResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetProgress(v int32) *RestoreInstanceResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetResponse(v string) *RestoreInstanceResponseBodyJob {
	s.Response = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetStatus(v string) *RestoreInstanceResponseBodyJob {
	s.Status = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetType(v string) *RestoreInstanceResponseBodyJob {
	s.Type = &v
	return s
}

type RestoreInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponse) SetHeaders(v map[string]*string) *RestoreInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestoreInstanceResponse) SetStatusCode(v int32) *RestoreInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreInstanceResponse) SetBody(v *RestoreInstanceResponseBody) *RestoreInstanceResponse {
	s.Body = v
	return s
}

type ResumeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ResumeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeInstanceRequest) SetInstanceId(v string) *ResumeInstanceRequest {
	s.InstanceId = &v
	return s
}

type ResumeInstanceResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ResumeInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponse) SetHeaders(v map[string]*string) *ResumeInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeInstanceResponse) SetStatusCode(v int32) *ResumeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeInstanceResponse) SetBody(v *ResumeInstanceResponseBody) *ResumeInstanceResponse {
	s.Body = v
	return s
}

type SwitchClusterMasterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-w3G9vOJI2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SwitchClusterMasterRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchClusterMasterRequest) GoString() string {
	return s.String()
}

func (s *SwitchClusterMasterRequest) SetClusterId(v string) *SwitchClusterMasterRequest {
	s.ClusterId = &v
	return s
}

func (s *SwitchClusterMasterRequest) SetInstanceId(v string) *SwitchClusterMasterRequest {
	s.InstanceId = &v
	return s
}

type SwitchClusterMasterResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchClusterMasterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchClusterMasterResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchClusterMasterResponseBody) SetRequestId(v string) *SwitchClusterMasterResponseBody {
	s.RequestId = &v
	return s
}

type SwitchClusterMasterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchClusterMasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchClusterMasterResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchClusterMasterResponse) GoString() string {
	return s.String()
}

func (s *SwitchClusterMasterResponse) SetHeaders(v map[string]*string) *SwitchClusterMasterResponse {
	s.Headers = v
	return s
}

func (s *SwitchClusterMasterResponse) SetStatusCode(v int32) *SwitchClusterMasterResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchClusterMasterResponse) SetBody(v *SwitchClusterMasterResponseBody) *SwitchClusterMasterResponse {
	s.Body = v
	return s
}

type SyncClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsgytet****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s SyncClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncClusterRequest) GoString() string {
	return s.String()
}

func (s *SyncClusterRequest) SetClusterId(v string) *SyncClusterRequest {
	s.ClusterId = &v
	return s
}

type SyncClusterResponseBody struct {
	Job *SyncClusterResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncClusterResponseBody) GoString() string {
	return s.String()
}

func (s *SyncClusterResponseBody) SetJob(v *SyncClusterResponseBodyJob) *SyncClusterResponseBody {
	s.Job = v
	return s
}

func (s *SyncClusterResponseBody) SetRequestId(v string) *SyncClusterResponseBody {
	s.RequestId = &v
	return s
}

type SyncClusterResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1711764127000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-000bu7m5vjmyz9s7qz85
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 90
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SyncClusterResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s SyncClusterResponseBodyJob) GoString() string {
	return s.String()
}

func (s *SyncClusterResponseBodyJob) SetCompleted(v bool) *SyncClusterResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetCreateTime(v string) *SyncClusterResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetError(v string) *SyncClusterResponseBodyJob {
	s.Error = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetJobId(v string) *SyncClusterResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetProgress(v int32) *SyncClusterResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetResponse(v string) *SyncClusterResponseBodyJob {
	s.Response = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetStatus(v string) *SyncClusterResponseBodyJob {
	s.Status = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetType(v string) *SyncClusterResponseBodyJob {
	s.Type = &v
	return s
}

type SyncClusterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncClusterResponse) GoString() string {
	return s.String()
}

func (s *SyncClusterResponse) SetHeaders(v map[string]*string) *SyncClusterResponse {
	s.Headers = v
	return s
}

func (s *SyncClusterResponse) SetStatusCode(v int32) *SyncClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncClusterResponse) SetBody(v *SyncClusterResponseBody) *SyncClusterResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("hsm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户地域下审计日志功能开通
//
// @param request - ConfigAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigAuditLogResponse
func (client *Client) ConfigAuditLogWithOptions(request *ConfigAuditLogRequest, runtime *util.RuntimeOptions) (_result *ConfigAuditLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuditAction)) {
		query["AuditAction"] = request.AuditAction
	}

	if !tea.BoolValue(util.IsUnset(request.AuditOssBucket)) {
		query["AuditOssBucket"] = request.AuditOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigAuditLog"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户地域下审计日志功能开通
//
// @param request - ConfigAuditLogRequest
//
// @return ConfigAuditLogResponse
func (client *Client) ConfigAuditLog(request *ConfigAuditLogRequest) (_result *ConfigAuditLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigAuditLogResponse{}
	_body, _err := client.ConfigAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置备份名与备注
//
// @param request - ConfigBackupRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigBackupRemarkResponse
func (client *Client) ConfigBackupRemarkWithOptions(request *ConfigBackupRemarkRequest, runtime *util.RuntimeOptions) (_result *ConfigBackupRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigBackupRemark"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigBackupRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置备份名与备注
//
// @param request - ConfigBackupRemarkRequest
//
// @return ConfigBackupRemarkResponse
func (client *Client) ConfigBackupRemark(request *ConfigBackupRemarkRequest) (_result *ConfigBackupRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigBackupRemarkResponse{}
	_body, _err := client.ConfigBackupRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置备份自动轮转任务
//
// @param request - ConfigBackupTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigBackupTaskResponse
func (client *Client) ConfigBackupTaskWithOptions(request *ConfigBackupTaskRequest, runtime *util.RuntimeOptions) (_result *ConfigBackupTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupHourInDay)) {
		query["BackupHourInDay"] = request.BackupHourInDay
	}

	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.BackupPeriod)) {
		query["BackupPeriod"] = request.BackupPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.Manual2PeriodicList)) {
		query["Manual2PeriodicList"] = request.Manual2PeriodicList
	}

	if !tea.BoolValue(util.IsUnset(request.Periodic2ManualList)) {
		query["Periodic2ManualList"] = request.Periodic2ManualList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigBackupTask"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigBackupTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置备份自动轮转任务
//
// @param request - ConfigBackupTaskRequest
//
// @return ConfigBackupTaskResponse
func (client *Client) ConfigBackupTask(request *ConfigBackupTaskRequest) (_result *ConfigBackupTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigBackupTaskResponse{}
	_body, _err := client.ConfigBackupTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 国际站配置HSM集群证书
//
// @param request - ConfigClusterCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterCertificateResponse
func (client *Client) ConfigClusterCertificateWithOptions(request *ConfigClusterCertificateRequest, runtime *util.RuntimeOptions) (_result *ConfigClusterCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterCertificate)) {
		body["ClusterCertificate"] = request.ClusterCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IssuerCertificate)) {
		body["IssuerCertificate"] = request.IssuerCertificate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigClusterCertificate"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigClusterCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 国际站配置HSM集群证书
//
// @param request - ConfigClusterCertificateRequest
//
// @return ConfigClusterCertificateResponse
func (client *Client) ConfigClusterCertificate(request *ConfigClusterCertificateRequest) (_result *ConfigClusterCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigClusterCertificateResponse{}
	_body, _err := client.ConfigClusterCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置集群名称
//
// @param request - ConfigClusterNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterNameResponse
func (client *Client) ConfigClusterNameWithOptions(request *ConfigClusterNameRequest, runtime *util.RuntimeOptions) (_result *ConfigClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigClusterName"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigClusterNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置集群名称
//
// @param request - ConfigClusterNameRequest
//
// @return ConfigClusterNameResponse
func (client *Client) ConfigClusterName(request *ConfigClusterNameRequest) (_result *ConfigClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigClusterNameResponse{}
	_body, _err := client.ConfigClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置集群子网
//
// @param tmpReq - ConfigClusterSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterSubnetResponse
func (client *Client) ConfigClusterSubnetWithOptions(tmpReq *ConfigClusterSubnetRequest, runtime *util.RuntimeOptions) (_result *ConfigClusterSubnetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ConfigClusterSubnetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.VSwitchIds)) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, tea.String("VSwitchIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIdsShrink)) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigClusterSubnet"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigClusterSubnetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置集群子网
//
// @param request - ConfigClusterSubnetRequest
//
// @return ConfigClusterSubnetResponse
func (client *Client) ConfigClusterSubnet(request *ConfigClusterSubnetRequest) (_result *ConfigClusterSubnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigClusterSubnetResponse{}
	_body, _err := client.ConfigClusterSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置集群名称
//
// @param request - ConfigClusterWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterWhitelistResponse
func (client *Client) ConfigClusterWhitelistWithOptions(request *ConfigClusterWhitelistRequest, runtime *util.RuntimeOptions) (_result *ConfigClusterWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Whitelist)) {
		body["Whitelist"] = request.Whitelist
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigClusterWhitelist"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigClusterWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置集群名称
//
// @param request - ConfigClusterWhitelistRequest
//
// @return ConfigClusterWhitelistResponse
func (client *Client) ConfigClusterWhitelist(request *ConfigClusterWhitelistRequest) (_result *ConfigClusterWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigClusterWhitelistResponse{}
	_body, _err := client.ConfigClusterWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置镜像备注
//
// @param request - ConfigImageRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigImageRemarkResponse
func (client *Client) ConfigImageRemarkWithOptions(request *ConfigImageRemarkRequest, runtime *util.RuntimeOptions) (_result *ConfigImageRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigImageRemark"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigImageRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置镜像备注
//
// @param request - ConfigImageRemarkRequest
//
// @return ConfigImageRemarkResponse
func (client *Client) ConfigImageRemark(request *ConfigImageRemarkRequest) (_result *ConfigImageRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigImageRemarkResponse{}
	_body, _err := client.ConfigImageRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ConfigInstanceIpAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceIpAddressResponse
func (client *Client) ConfigInstanceIpAddressWithOptions(request *ConfigInstanceIpAddressRequest, runtime *util.RuntimeOptions) (_result *ConfigInstanceIpAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigInstanceIpAddress"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigInstanceIpAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConfigInstanceIpAddressRequest
//
// @return ConfigInstanceIpAddressResponse
func (client *Client) ConfigInstanceIpAddress(request *ConfigInstanceIpAddressRequest) (_result *ConfigInstanceIpAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigInstanceIpAddressResponse{}
	_body, _err := client.ConfigInstanceIpAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ConfigInstanceRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceRemarkResponse
func (client *Client) ConfigInstanceRemarkWithOptions(request *ConfigInstanceRemarkRequest, runtime *util.RuntimeOptions) (_result *ConfigInstanceRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigInstanceRemark"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigInstanceRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ConfigInstanceRemarkRequest
//
// @return ConfigInstanceRemarkResponse
func (client *Client) ConfigInstanceRemark(request *ConfigInstanceRemarkRequest) (_result *ConfigInstanceRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigInstanceRemarkResponse{}
	_body, _err := client.ConfigInstanceRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// TODO 不允许控制台直接修改集群内实例的白名单实现重构
//
// @param request - ConfigInstanceWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigInstanceWhitelistResponse
func (client *Client) ConfigInstanceWhitelistWithOptions(request *ConfigInstanceWhitelistRequest, runtime *util.RuntimeOptions) (_result *ConfigInstanceWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Whitelist)) {
		body["Whitelist"] = request.Whitelist
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigInstanceWhitelist"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigInstanceWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// TODO 不允许控制台直接修改集群内实例的白名单实现重构
//
// @param request - ConfigInstanceWhitelistRequest
//
// @return ConfigInstanceWhitelistResponse
func (client *Client) ConfigInstanceWhitelist(request *ConfigInstanceWhitelistRequest) (_result *ConfigInstanceWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigInstanceWhitelistResponse{}
	_body, _err := client.ConfigInstanceWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 跨地域复制镜像
//
// @param request - CopyImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyImageResponse
func (client *Client) CopyImageWithOptions(request *CopyImageRequest, runtime *util.RuntimeOptions) (_result *CopyImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageUid)) {
		body["ImageUid"] = request.ImageUid
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRegionId)) {
		body["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyImage"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 跨地域复制镜像
//
// @param request - CopyImageRequest
//
// @return CopyImageResponse
func (client *Client) CopyImage(request *CopyImageRequest) (_result *CopyImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyImageResponse{}
	_body, _err := client.CopyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.MasterInstanceId)) {
		body["MasterInstanceId"] = request.MasterInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用备份
//
// @param request - EnableBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableBackupResponse
func (client *Client) EnableBackupWithOptions(request *EnableBackupRequest, runtime *util.RuntimeOptions) (_result *EnableBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableBackup"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用备份
//
// @param request - EnableBackupRequest
//
// @return EnableBackupResponse
func (client *Client) EnableBackup(request *EnableBackupRequest) (_result *EnableBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableBackupResponse{}
	_body, _err := client.EnableBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 手动导出镜像
//
// @param request - ExportImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportImageResponse
func (client *Client) ExportImageWithOptions(request *ExportImageRequest, runtime *util.RuntimeOptions) (_result *ExportImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportImage"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 手动导出镜像
//
// @param request - ExportImageRequest
//
// @return ExportImageResponse
func (client *Client) ExportImage(request *ExportImageRequest) (_result *ExportImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportImageResponse{}
	_body, _err := client.ExportImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询审计日志功能开通状态
//
// @param request - GetAuditLogStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuditLogStatusResponse
func (client *Client) GetAuditLogStatusWithOptions(request *GetAuditLogStatusRequest, runtime *util.RuntimeOptions) (_result *GetAuditLogStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GetOssBucket)) {
		query["GetOssBucket"] = request.GetOssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuditLogStatus"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuditLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询审计日志功能开通状态
//
// @param request - GetAuditLogStatusRequest
//
// @return GetAuditLogStatusResponse
func (client *Client) GetAuditLogStatus(request *GetAuditLogStatusRequest) (_result *GetAuditLogStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuditLogStatusResponse{}
	_body, _err := client.GetAuditLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示用户备份
//
// @param request - GetBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBackupResponse
func (client *Client) GetBackupWithOptions(request *GetBackupRequest, runtime *util.RuntimeOptions) (_result *GetBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBackup"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示用户备份
//
// @param request - GetBackupRequest
//
// @return GetBackupResponse
func (client *Client) GetBackup(request *GetBackupRequest) (_result *GetBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBackupResponse{}
	_body, _err := client.GetBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCluster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示备份下的用户镜像
//
// @param request - GetImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
func (client *Client) GetImageWithOptions(request *GetImageRequest, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImage"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示备份下的用户镜像
//
// @param request - GetImageRequest
//
// @return GetImageResponse
func (client *Client) GetImage(request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务执行信息
//
// @param request - GetJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(request *GetJobRequest, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务执行信息
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 为用户创建审计日志的服务关联角色
//
// @param request - InitializeAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeAuditLogResponse
func (client *Client) InitializeAuditLogWithOptions(runtime *util.RuntimeOptions) (_result *InitializeAuditLogResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("InitializeAuditLog"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为用户创建审计日志的服务关联角色
//
// @return InitializeAuditLogResponse
func (client *Client) InitializeAuditLog() (_result *InitializeAuditLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeAuditLogResponse{}
	_body, _err := client.InitializeAuditLogWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 初始化集群（原activeCluster）
//
// @param request - InitializeClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitializeClusterResponse
func (client *Client) InitializeClusterWithOptions(request *InitializeClusterRequest, runtime *util.RuntimeOptions) (_result *InitializeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitializeCluster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化集群（原activeCluster）
//
// @param request - InitializeClusterRequest
//
// @return InitializeClusterResponse
func (client *Client) InitializeCluster(request *InitializeClusterRequest) (_result *InitializeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeClusterResponse{}
	_body, _err := client.InitializeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - JoinClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinClusterResponse
func (client *Client) JoinClusterWithOptions(request *JoinClusterRequest, runtime *util.RuntimeOptions) (_result *JoinClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinCluster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - JoinClusterRequest
//
// @return JoinClusterResponse
func (client *Client) JoinCluster(request *JoinClusterRequest) (_result *JoinClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinClusterResponse{}
	_body, _err := client.JoinClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - LeaveClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LeaveClusterResponse
func (client *Client) LeaveClusterWithOptions(request *LeaveClusterRequest, runtime *util.RuntimeOptions) (_result *LeaveClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LeaveCluster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LeaveClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - LeaveClusterRequest
//
// @return LeaveClusterResponse
func (client *Client) LeaveCluster(request *LeaveClusterRequest) (_result *LeaveClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LeaveClusterResponse{}
	_body, _err := client.LeaveClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询备份列表
//
// @param request - ListBackupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackupsResponse
func (client *Client) ListBackupsWithOptions(request *ListBackupsRequest, runtime *util.RuntimeOptions) (_result *ListBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBackups"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询备份列表
//
// @param request - ListBackupsRequest
//
// @return ListBackupsResponse
func (client *Client) ListBackups(request *ListBackupsRequest) (_result *ListBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBackupsResponse{}
	_body, _err := client.ListBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户镜像列表
//
// @param request - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户镜像列表
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资源组移动资源
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源组移动资源
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PauseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseInstanceResponse
func (client *Client) PauseInstanceWithOptions(request *PauseInstanceRequest, runtime *util.RuntimeOptions) (_result *PauseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PauseInstance"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PauseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PauseInstanceRequest
//
// @return PauseInstanceResponse
func (client *Client) PauseInstance(request *PauseInstanceRequest) (_result *PauseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseInstanceResponse{}
	_body, _err := client.PauseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuickInitInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuickInitInstanceResponse
func (client *Client) QuickInitInstanceWithOptions(request *QuickInitInstanceRequest, runtime *util.RuntimeOptions) (_result *QuickInitInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QuickInitInstance"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuickInitInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuickInitInstanceRequest
//
// @return QuickInitInstanceResponse
func (client *Client) QuickInitInstance(request *QuickInitInstanceRequest) (_result *QuickInitInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuickInitInstanceResponse{}
	_body, _err := client.QuickInitInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置备份
//
// @param request - ResetBackupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetBackupResponse
func (client *Client) ResetBackupWithOptions(request *ResetBackupRequest, runtime *util.RuntimeOptions) (_result *ResetBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupId)) {
		query["BackupId"] = request.BackupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetBackup"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置备份
//
// @param request - ResetBackupRequest
//
// @return ResetBackupResponse
func (client *Client) ResetBackup(request *ResetBackupRequest) (_result *ResetBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetBackupResponse{}
	_body, _err := client.ResetBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetInstanceResponse
func (client *Client) ResetInstanceWithOptions(request *ResetInstanceRequest, runtime *util.RuntimeOptions) (_result *ResetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetInstance"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResetInstanceRequest
//
// @return ResetInstanceResponse
func (client *Client) ResetInstance(request *ResetInstanceRequest) (_result *ResetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetInstanceResponse{}
	_body, _err := client.ResetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RestoreInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstanceWithOptions(request *RestoreInstanceRequest, runtime *util.RuntimeOptions) (_result *RestoreInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreInstance"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RestoreInstanceRequest
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstance(request *RestoreInstanceRequest) (_result *RestoreInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.RestoreInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResumeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstanceWithOptions(request *ResumeInstanceRequest, runtime *util.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeInstance"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResumeInstanceRequest
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstance(request *ResumeInstanceRequest) (_result *ResumeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SwitchClusterMasterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchClusterMasterResponse
func (client *Client) SwitchClusterMasterWithOptions(request *SwitchClusterMasterRequest, runtime *util.RuntimeOptions) (_result *SwitchClusterMasterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchClusterMaster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchClusterMasterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SwitchClusterMasterRequest
//
// @return SwitchClusterMasterResponse
func (client *Client) SwitchClusterMaster(request *SwitchClusterMasterRequest) (_result *SwitchClusterMasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchClusterMasterResponse{}
	_body, _err := client.SwitchClusterMasterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SyncClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncClusterResponse
func (client *Client) SyncClusterWithOptions(request *SyncClusterRequest, runtime *util.RuntimeOptions) (_result *SyncClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncCluster"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - SyncClusterRequest
//
// @return SyncClusterResponse
func (client *Client) SyncCluster(request *SyncClusterRequest) (_result *SyncClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncClusterResponse{}
	_body, _err := client.SyncClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
