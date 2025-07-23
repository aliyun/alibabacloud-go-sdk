// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v *GetClusterResponseBodyCluster) *GetClusterResponseBody
	GetCluster() *GetClusterResponseBodyCluster
	SetRequestId(v string) *GetClusterResponseBody
	GetRequestId() *string
}

type GetClusterResponseBody struct {
	// The cluster details.
	Cluster *GetClusterResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) GetCluster() *GetClusterResponseBodyCluster {
	return s.Cluster
}

func (s *GetClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterResponseBody) SetCluster(v *GetClusterResponseBodyCluster) *GetClusterResponseBody {
	s.Cluster = v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyCluster struct {
	CertManaged *bool `json:"CertManaged,omitempty" xml:"CertManaged,omitempty"`
	// The cluster certificate.
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
	ClusterCertificate *string `json:"ClusterCertificate,omitempty" xml:"ClusterCertificate,omitempty"`
	// The certificate signing request (CSR) file of the cluster.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----\\nMIIC5TCCAc0CAQAwgZ8xWTAJBgNVBAYTAlVTMAkGA1UECAwCQ0EwDQYDVQQKDAZD\\nYXZpdW0wDQYDVQQLDAZOM0ZJUFMwDgYDVQQHDAdTYW5Kb3NlMBMGA1UdEQwMMTk****
	//
	// -----END CERTIFICATE REQUEST-----
	ClusterCsr *string `json:"ClusterCsr,omitempty" xml:"ClusterCsr,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// cluster-p94y1dud9ts****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster mode.
	//
	// 2: automatically synchronizes the cluster.
	//
	// example:
	//
	// 2
	ClusterMode *int32 `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// cluster_polar_****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The self-signed certificate of the cluster.
	//
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
	// The time when the cluster was created. Unit: milliseconds. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1641275680000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the device.
	//
	// example:
	//
	// jnta
	DeviceType           *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	EntityCertExpireTime *string `json:"EntityCertExpireTime,omitempty" xml:"EntityCertExpireTime,omitempty"`
	// The HSMs in the cluster.
	Instances []*GetClusterResponseBodyClusterInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the region in which the cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of hardware security modules (HSMs) in the cluster.
	//
	// example:
	//
	// 2
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- NEW: The cluster is not initialized.
	//
	// 	- INITIALIZED: The cluster is initialized.
	//
	// 	- DELETED: The cluster is deleted.
	//
	// 	- SYNCHRONIZING: The cluster is being synchronized.
	//
	// 	- TO_DELETE: The cluster is pending deletion.
	//
	// example:
	//
	// NEW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the cluster belongs.
	//
	// example:
	//
	// vpc-8vbt0fjdm29hofvbo****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IP address whitelist of the cluster.
	//
	// example:
	//
	// 130.176.XX.XX
	Whitelist *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
	// The information about the zones in which the cluster is deployed.
	Zones []*GetClusterResponseBodyClusterZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s GetClusterResponseBodyCluster) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyCluster) GetCertManaged() *bool {
	return s.CertManaged
}

func (s *GetClusterResponseBodyCluster) GetClusterCertificate() *string {
	return s.ClusterCertificate
}

func (s *GetClusterResponseBodyCluster) GetClusterCsr() *string {
	return s.ClusterCsr
}

func (s *GetClusterResponseBodyCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterResponseBodyCluster) GetClusterMode() *int32 {
	return s.ClusterMode
}

func (s *GetClusterResponseBodyCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetClusterResponseBodyCluster) GetClusterOwnerCertificate() *string {
	return s.ClusterOwnerCertificate
}

func (s *GetClusterResponseBodyCluster) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetClusterResponseBodyCluster) GetDeviceType() *string {
	return s.DeviceType
}

func (s *GetClusterResponseBodyCluster) GetEntityCertExpireTime() *string {
	return s.EntityCertExpireTime
}

func (s *GetClusterResponseBodyCluster) GetInstances() []*GetClusterResponseBodyClusterInstances {
	return s.Instances
}

func (s *GetClusterResponseBodyCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *GetClusterResponseBodyCluster) GetSize() *int32 {
	return s.Size
}

func (s *GetClusterResponseBodyCluster) GetStatus() *string {
	return s.Status
}

func (s *GetClusterResponseBodyCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *GetClusterResponseBodyCluster) GetWhitelist() *string {
	return s.Whitelist
}

func (s *GetClusterResponseBodyCluster) GetZones() []*GetClusterResponseBodyClusterZones {
	return s.Zones
}

func (s *GetClusterResponseBodyCluster) SetCertManaged(v bool) *GetClusterResponseBodyCluster {
	s.CertManaged = &v
	return s
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

func (s *GetClusterResponseBodyCluster) SetClusterMode(v int32) *GetClusterResponseBodyCluster {
	s.ClusterMode = &v
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

func (s *GetClusterResponseBodyCluster) SetEntityCertExpireTime(v string) *GetClusterResponseBodyCluster {
	s.EntityCertExpireTime = &v
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

func (s *GetClusterResponseBodyCluster) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyClusterInstances struct {
	// The ID of the HSM.
	//
	// example:
	//
	// hsm-cn-g6z3v0uf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the HSM is a master HSM. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Master *bool `json:"Master,omitempty" xml:"Master,omitempty"`
	// The ID of the HSM in the cluster.
	//
	// example:
	//
	// 1
	NodeId *int32 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s GetClusterResponseBodyClusterInstances) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyClusterInstances) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyClusterInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetClusterResponseBodyClusterInstances) GetMaster() *bool {
	return s.Master
}

func (s *GetClusterResponseBodyClusterInstances) GetNodeId() *int32 {
	return s.NodeId
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

func (s *GetClusterResponseBodyClusterInstances) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyClusterZones struct {
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf61s651p69bdgmki****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetClusterResponseBodyClusterZones) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyClusterZones) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyClusterZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetClusterResponseBodyClusterZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetClusterResponseBodyClusterZones) SetVSwitchId(v string) *GetClusterResponseBodyClusterZones {
	s.VSwitchId = &v
	return s
}

func (s *GetClusterResponseBodyClusterZones) SetZoneId(v string) *GetClusterResponseBodyClusterZones {
	s.ZoneId = &v
	return s
}

func (s *GetClusterResponseBodyClusterZones) Validate() error {
	return dara.Validate(s)
}
