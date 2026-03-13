// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerTCPListenerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendServerPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetBackendServerPort() *int32
	SetBandwidth(v int32) *CreateLoadBalancerTCPListenerRequest
	GetBandwidth() *int32
	SetConnectPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetConnectPort() *int32
	SetConnectTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetConnectTimeout() *int32
	SetEstablishedTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetEstablishedTimeout() *int32
	SetHealthCheck(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheck() *string
	SetHealthCheckDomain(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckDomain() *string
	SetHealthCheckHttpCode(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckHttpCode() *string
	SetHealthCheckType(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckType() *string
	SetHealthCheckURI(v string) *CreateLoadBalancerTCPListenerRequest
	GetHealthCheckURI() *string
	SetHealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest
	GetHealthyThreshold() *int32
	SetInterval(v int32) *CreateLoadBalancerTCPListenerRequest
	GetInterval() *int32
	SetListenerPort(v int32) *CreateLoadBalancerTCPListenerRequest
	GetListenerPort() *int32
	SetListenerStatus(v string) *CreateLoadBalancerTCPListenerRequest
	GetListenerStatus() *string
	SetLoadBalancerId(v string) *CreateLoadBalancerTCPListenerRequest
	GetLoadBalancerId() *string
	SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest
	GetMasterSlaveServerGroupId() *string
	SetMaxConnection(v int32) *CreateLoadBalancerTCPListenerRequest
	GetMaxConnection() *int32
	SetOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest
	GetOwnerId() *int64
	SetPersistenceTimeout(v int32) *CreateLoadBalancerTCPListenerRequest
	GetPersistenceTimeout() *int32
	SetResourceOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest
	GetResourceOwnerId() *int64
	SetScheduler(v string) *CreateLoadBalancerTCPListenerRequest
	GetScheduler() *string
	SetTags(v string) *CreateLoadBalancerTCPListenerRequest
	GetTags() *string
	SetUnhealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest
	GetUnhealthyThreshold() *int32
	SetVServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest
	GetVServerGroupId() *string
	SetAccessKeyId(v string) *CreateLoadBalancerTCPListenerRequest
	GetAccessKeyId() *string
}

type CreateLoadBalancerTCPListenerRequest struct {
	BackendServerPort   *int32  `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	Bandwidth           *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectPort         *int32  `json:"ConnectPort,omitempty" xml:"ConnectPort,omitempty"`
	ConnectTimeout      *int32  `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	EstablishedTimeout  *int32  `json:"EstablishedTimeout,omitempty" xml:"EstablishedTimeout,omitempty"`
	HealthCheck         *string `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty"`
	HealthCheckDomain   *string `json:"HealthCheckDomain,omitempty" xml:"HealthCheckDomain,omitempty"`
	HealthCheckHttpCode *string `json:"HealthCheckHttpCode,omitempty" xml:"HealthCheckHttpCode,omitempty"`
	HealthCheckType     *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	HealthCheckURI      *string `json:"HealthCheckURI,omitempty" xml:"HealthCheckURI,omitempty"`
	HealthyThreshold    *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	Interval            *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	ListenerPort   *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerStatus *string `json:"ListenerStatus,omitempty" xml:"ListenerStatus,omitempty"`
	// This parameter is required.
	LoadBalancerId           *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	MaxConnection            *int32  `json:"MaxConnection,omitempty" xml:"MaxConnection,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PersistenceTimeout       *int32  `json:"PersistenceTimeout,omitempty" xml:"PersistenceTimeout,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Scheduler                *string `json:"Scheduler,omitempty" xml:"Scheduler,omitempty"`
	Tags                     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UnhealthyThreshold       *int32  `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
	VServerGroupId           *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	AccessKeyId              *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s CreateLoadBalancerTCPListenerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerRequest) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateLoadBalancerTCPListenerRequest) GetConnectPort() *int32 {
	return s.ConnectPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetEstablishedTimeout() *int32 {
	return s.EstablishedTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheck() *string {
	return s.HealthCheck
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckDomain() *string {
	return s.HealthCheckDomain
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckHttpCode() *string {
	return s.HealthCheckHttpCode
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthCheckURI() *string {
	return s.HealthCheckURI
}

func (s *CreateLoadBalancerTCPListenerRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *CreateLoadBalancerTCPListenerRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateLoadBalancerTCPListenerRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateLoadBalancerTCPListenerRequest) GetListenerStatus() *string {
	return s.ListenerStatus
}

func (s *CreateLoadBalancerTCPListenerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetMaxConnection() *int32 {
	return s.MaxConnection
}

func (s *CreateLoadBalancerTCPListenerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLoadBalancerTCPListenerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetPersistenceTimeout() *int32 {
	return s.PersistenceTimeout
}

func (s *CreateLoadBalancerTCPListenerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLoadBalancerTCPListenerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetScheduler() *string {
	return s.Scheduler
}

func (s *CreateLoadBalancerTCPListenerRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateLoadBalancerTCPListenerRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *CreateLoadBalancerTCPListenerRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *CreateLoadBalancerTCPListenerRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *CreateLoadBalancerTCPListenerRequest) SetBackendServerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.BackendServerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetBandwidth(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetConnectPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ConnectPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetConnectTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ConnectTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetEstablishedTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.EstablishedTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheck(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheck = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckDomain(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckDomain = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckHttpCode(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckHttpCode = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckType(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthCheckURI(v string) *CreateLoadBalancerTCPListenerRequest {
	s.HealthCheckURI = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetHealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetInterval(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.Interval = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetListenerPort(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetListenerStatus(v string) *CreateLoadBalancerTCPListenerRequest {
	s.ListenerStatus = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetLoadBalancerId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetMasterSlaveServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetMaxConnection(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.MaxConnection = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetPersistenceTimeout(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.PersistenceTimeout = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerAccount(v string) *CreateLoadBalancerTCPListenerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetResourceOwnerId(v int64) *CreateLoadBalancerTCPListenerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetScheduler(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Scheduler = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetTags(v string) *CreateLoadBalancerTCPListenerRequest {
	s.Tags = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetUnhealthyThreshold(v int32) *CreateLoadBalancerTCPListenerRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetVServerGroupId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.VServerGroupId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) SetAccessKeyId(v string) *CreateLoadBalancerTCPListenerRequest {
	s.AccessKeyId = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerRequest) Validate() error {
	return dara.Validate(s)
}
