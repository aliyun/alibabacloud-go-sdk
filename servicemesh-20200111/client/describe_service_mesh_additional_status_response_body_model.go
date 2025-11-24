// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshAdditionalStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) *DescribeServiceMeshAdditionalStatusResponseBody
	GetClusterStatus() *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus
	SetRequestId(v string) *DescribeServiceMeshAdditionalStatusResponseBody
	GetRequestId() *string
}

type DescribeServiceMeshAdditionalStatusResponseBody struct {
	// The cluster status.
	ClusterStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 11fd0027-c27e-41bb-a565-75583054****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBody) GetClusterStatus() *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	return s.ClusterStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshAdditionalStatusResponseBody) SetClusterStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) *DescribeServiceMeshAdditionalStatusResponseBody {
	s.ClusterStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBody) SetRequestId(v string) *DescribeServiceMeshAdditionalStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBody) Validate() error {
	if s.ClusterStatus != nil {
		if err := s.ClusterStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus struct {
	// Indicates whether access logs exist. Valid values:
	//
	// 	- `exist`: Access logs exist.
	//
	// 	- `not_exist`: Access logs do not exist.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check times out.
	//
	// example:
	//
	// exist
	AccessLogProjectStatus *string `json:"AccessLogProjectStatus,omitempty" xml:"AccessLogProjectStatus,omitempty"`
	// The check result of the EIP associated with the API server. Valid values:
	//
	// 	- `exist`: The EIP exists.
	//
	// 	- `not_exist`: The EIP does not exist.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check times out.
	//
	// 	- `not_in_use`: The EIP is not associated with the API Server.
	//
	// 	- `locked`: The EIP is locked.
	//
	// example:
	//
	// exist
	ApiServerEIPStatus *string `json:"ApiServerEIPStatus,omitempty" xml:"ApiServerEIPStatus,omitempty"`
	// The check results of the CLB instance created for exposing the API server.
	ApiServerLoadBalancerStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus `json:"ApiServerLoadBalancerStatus,omitempty" xml:"ApiServerLoadBalancerStatus,omitempty" type:"Struct"`
	// Indicates whether audit logs exist. Valid values:
	//
	// 	- `exist`
	//
	// 	- `not exist`
	//
	// example:
	//
	// exist
	AuditProjectStatus   *string `json:"AuditProjectStatus,omitempty" xml:"AuditProjectStatus,omitempty"`
	CanaryPilotEIPStatus *string `json:"CanaryPilotEIPStatus,omitempty" xml:"CanaryPilotEIPStatus,omitempty"`
	// The check results of the CLB instance that is created for exposing Istio Pilot and used during canary release.
	CanaryPilotLoadBalancerStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus `json:"CanaryPilotLoadBalancerStatus,omitempty" xml:"CanaryPilotLoadBalancerStatus,omitempty" type:"Struct"`
	// Indicates whether control plane logs exist. Valid values:
	//
	// 	- `exist`: Control plane logs exist.
	//
	// 	- `not_exist`: Control plane logs do not exist.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check times out.
	//
	// example:
	//
	// exist
	ControlPlaneProjectStatus *string `json:"ControlPlaneProjectStatus,omitempty" xml:"ControlPlaneProjectStatus,omitempty"`
	// Indicates whether Logtail is installed in clusters on the data plane.
	LogtailStatusRecord map[string]interface{} `json:"LogtailStatusRecord,omitempty" xml:"LogtailStatusRecord,omitempty"`
	// The check result of whether the EIP is associated with the API server. Valid values:
	//
	// 	- `exist`: The EIP exists.
	//
	// 	- `not_exist`: The EIP does not exist.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check is timed out.
	//
	// 	- `not_in_use`: The EIP is not associated with the API server.
	//
	// 	- `locked`: The EIP is locked.
	//
	// example:
	//
	// exist
	PilotEIPStatus *string `json:"PilotEIPStatus,omitempty" xml:"PilotEIPStatus,omitempty"`
	// The check results of the CLB instance created for exposing Istio Pilot.
	PilotLoadBalancerStatus *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus `json:"PilotLoadBalancerStatus,omitempty" xml:"PilotLoadBalancerStatus,omitempty" type:"Struct"`
	// The status of the RAM OAuth application that is integrated with Mesh Topology. Valid values:
	//
	// 	- `exist`: The RAM OAuth application exists.
	//
	// 	- `reused`: The RAM OAuth application is reused.
	//
	// 	- `not_exist`: The RAM OAuth application does not exist.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check times out.
	//
	// example:
	//
	// reused
	RAMApplicationStatus *string `json:"RAMApplicationStatus,omitempty" xml:"RAMApplicationStatus,omitempty"`
	// Indicates whether the security group is reused. Valid values:
	//
	// 	- `reused`: The security group is reused.
	//
	// 	- `not_reused`: The security group is not reused.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check times out.
	//
	// example:
	//
	// reused
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetAccessLogProjectStatus() *string {
	return s.AccessLogProjectStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetApiServerEIPStatus() *string {
	return s.ApiServerEIPStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetApiServerLoadBalancerStatus() *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	return s.ApiServerLoadBalancerStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetAuditProjectStatus() *string {
	return s.AuditProjectStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetCanaryPilotEIPStatus() *string {
	return s.CanaryPilotEIPStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetCanaryPilotLoadBalancerStatus() *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	return s.CanaryPilotLoadBalancerStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetControlPlaneProjectStatus() *string {
	return s.ControlPlaneProjectStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetLogtailStatusRecord() map[string]interface{} {
	return s.LogtailStatusRecord
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetPilotEIPStatus() *string {
	return s.PilotEIPStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetPilotLoadBalancerStatus() *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	return s.PilotLoadBalancerStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetRAMApplicationStatus() *string {
	return s.RAMApplicationStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) GetSgStatus() *string {
	return s.SgStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetAccessLogProjectStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.AccessLogProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetApiServerEIPStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.ApiServerEIPStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetApiServerLoadBalancerStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.ApiServerLoadBalancerStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetAuditProjectStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.AuditProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetCanaryPilotEIPStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.CanaryPilotEIPStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetCanaryPilotLoadBalancerStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.CanaryPilotLoadBalancerStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetControlPlaneProjectStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.ControlPlaneProjectStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetLogtailStatusRecord(v map[string]interface{}) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.LogtailStatusRecord = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetPilotEIPStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.PilotEIPStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetPilotLoadBalancerStatus(v *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.PilotLoadBalancerStatus = v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetRAMApplicationStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.RAMApplicationStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) SetSgStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus {
	s.SgStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatus) Validate() error {
	if s.ApiServerLoadBalancerStatus != nil {
		if err := s.ApiServerLoadBalancerStatus.Validate(); err != nil {
			return err
		}
	}
	if s.CanaryPilotLoadBalancerStatus != nil {
		if err := s.CanaryPilotLoadBalancerStatus.Validate(); err != nil {
			return err
		}
	}
	if s.PilotLoadBalancerStatus != nil {
		if err := s.PilotLoadBalancerStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus struct {
	// Indicates whether the CLB instance is locked. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The billing method of the CLB instance. Valid values:
	//
	// 	- `PrePay`: subscription
	//
	// 	- `PayOnDemand`: pay-as-you-go
	//
	// example:
	//
	// PrePay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates whether the CLB instance is reused. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Reused *bool `json:"Reused,omitempty" xml:"Reused,omitempty"`
	// The check results of the number of backend servers of the CLB instance created for exposing Istio Pilot. Valid values:
	//
	// 	- `too_much`: An excessive number of backend servers are created.
	//
	// 	- `num_exact`: A proper number of backend servers are created.
	//
	// 	- `too_little`: The number of backend servers falls short.
	//
	// example:
	//
	// num_exact
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	// The check results of the CLB instance. Valid values:
	//
	// 	- `exist`: The CLB instance exists.
	//
	// 	- `not_exist`: The CLB instance does not exist.
	//
	// 	- `conflict`: Conflicts are detected.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check times out.
	//
	// example:
	//
	// exist
	SLBExistStatus *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GetLocked() *bool {
	return s.Locked
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GetPayType() *string {
	return s.PayType
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GetReused() *bool {
	return s.Reused
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GetSLBBackEndServerNumStatus() *string {
	return s.SLBBackEndServerNumStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) GetSLBExistStatus() *string {
	return s.SLBExistStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetLocked(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.Locked = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetPayType(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.PayType = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetReused(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.Reused = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetSLBBackEndServerNumStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.SLBBackEndServerNumStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) SetSLBExistStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus {
	s.SLBExistStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusApiServerLoadBalancerStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus struct {
	// Indicates whether the CLB instance is locked due to overdue payments. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The billing method of the CLB instance. Valid values:
	//
	// 	- `PrePay`: subscription
	//
	// 	- `PayOnDemand` (default): pay-as-you-go
	//
	// example:
	//
	// PayOnDemand
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates whether the CLB instance is reused. Valid values:
	//
	// 	- `true`: The CLB instance is reused. Non-ASM listener configuration is detected in the listener configurations of the CLB instance.
	//
	// 	- `false`: The CLB instance is not reused.
	//
	// example:
	//
	// false
	Reused *bool `json:"Reused,omitempty" xml:"Reused,omitempty"`
	// The check result of the number of backend servers of the CLB instance created for exposing Istio Pilot. Valid values:
	//
	// 	- `num_exact`: A proper number of backend servers are created.
	//
	// 	- `too_much`: An excessive number of backend servers are created.
	//
	// 	- `too_little`: The number of backend servers falls short.
	//
	// example:
	//
	// num_exact
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	// The check result of the CLB instance. Valid values:
	//
	// 	- `exist`: The CLB instance exists.
	//
	// 	- `not_exist`: The CLB instance does not exist.
	//
	// 	- `time_out`: The check times out.
	//
	// 	- `failed`: The CLB instance has expired.
	//
	// example:
	//
	// exist
	SLBExistStatus *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) GetLocked() *bool {
	return s.Locked
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) GetPayType() *string {
	return s.PayType
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) GetReused() *bool {
	return s.Reused
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) GetSLBBackEndServerNumStatus() *string {
	return s.SLBBackEndServerNumStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) GetSLBExistStatus() *string {
	return s.SLBExistStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetLocked(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.Locked = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetPayType(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.PayType = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetReused(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.Reused = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetSLBBackEndServerNumStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.SLBBackEndServerNumStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) SetSLBExistStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus {
	s.SLBExistStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusCanaryPilotLoadBalancerStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus struct {
	// Indicates whether the CLB instance is locked. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	Locked *bool `json:"Locked,omitempty" xml:"Locked,omitempty"`
	// The billing method of the CLB instance. Valid values:
	//
	// 	- `PrePay`: subscription
	//
	// 	- `PayOnDemand`: pay-as-you-go
	//
	// example:
	//
	// PayOnDemand
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Indicates whether the CLB instance is reused. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Reused *bool `json:"Reused,omitempty" xml:"Reused,omitempty"`
	// The check results of the number of backend servers of the CLB instance created for exposing Istio Pilot. Valid values:
	//
	// 	- `too_much`: An excessive number of backend servers are created.
	//
	// 	- `num_exact`: A proper number of backend servers are created.
	//
	// 	- `too_little`: The number of backend servers falls short.
	//
	// example:
	//
	// num_exact
	SLBBackEndServerNumStatus *string `json:"SLBBackEndServerNumStatus,omitempty" xml:"SLBBackEndServerNumStatus,omitempty"`
	// The check results of the CLB instance. Valid values:
	//
	// 	- `exist`: The CLB instance exists.
	//
	// 	- `not_exist`: The CLB instance does not exist.
	//
	// 	- `conflict`: Conflicts are detected.
	//
	// 	- `failed`: The check fails.
	//
	// 	- `time_out`: The check times out.
	//
	// example:
	//
	// exist
	SLBExistStatus *string `json:"SLBExistStatus,omitempty" xml:"SLBExistStatus,omitempty"`
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GetLocked() *bool {
	return s.Locked
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GetPayType() *string {
	return s.PayType
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GetReused() *bool {
	return s.Reused
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GetSLBBackEndServerNumStatus() *string {
	return s.SLBBackEndServerNumStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) GetSLBExistStatus() *string {
	return s.SLBExistStatus
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetLocked(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.Locked = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetPayType(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.PayType = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetReused(v bool) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.Reused = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetSLBBackEndServerNumStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.SLBBackEndServerNumStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) SetSLBExistStatus(v string) *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus {
	s.SLBExistStatus = &v
	return s
}

func (s *DescribeServiceMeshAdditionalStatusResponseBodyClusterStatusPilotLoadBalancerStatus) Validate() error {
	return dara.Validate(s)
}
