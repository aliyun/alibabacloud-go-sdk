// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHealthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeHealthStatusResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DescribeHealthStatusResponseBody
	GetRequestId() *string
	SetStatus(v *DescribeHealthStatusResponseBodyStatus) *DescribeHealthStatusResponseBody
	GetStatus() *DescribeHealthStatusResponseBodyStatus
}

type DescribeHealthStatusResponseBody struct {
	// The ID of instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D97B4191-104D-10CE-8BC5-53**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried performance metrics. Each performance metric consists of the parameter name, status, and metric value. The metric information is returned only for the performance parameters specified by **Key**. For example, if you set **Key*	- to **adbpg_status**, only the metric information of **adbpg_status*	- is returned.
	//
	// For more information about performance parameters, see [Performance parameters](https://help.aliyun.com/document_detail/86943.html).
	//
	// example:
	//
	// {"node_master_connection_status":{"Status":"healthy","Value":1.6}}
	Status *DescribeHealthStatusResponseBodyStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s DescribeHealthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHealthStatusResponseBody) GetStatus() *DescribeHealthStatusResponseBodyStatus {
	return s.Status
}

func (s *DescribeHealthStatusResponseBody) SetDBClusterId(v string) *DescribeHealthStatusResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetRequestId(v string) *DescribeHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHealthStatusResponseBody) SetStatus(v *DescribeHealthStatusResponseBodyStatus) *DescribeHealthStatusResponseBody {
	s.Status = v
	return s
}

func (s *DescribeHealthStatusResponseBody) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHealthStatusResponseBodyStatus struct {
	// The information of maximum compute node storage usage.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbgpSegmentDiskUsagePercentMax *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax `json:"adbgp_segment_disk_usage_percent_max,omitempty" xml:"adbgp_segment_disk_usage_percent_max,omitempty" type:"Struct"`
	// The information of instance connection health status.
	AdbpgConnectionStatus *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus `json:"adbpg_connection_status,omitempty" xml:"adbpg_connection_status,omitempty" type:"Struct"`
	// The information of instance storage status.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbpgDiskStatus *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus `json:"adbpg_disk_status,omitempty" xml:"adbpg_disk_status,omitempty" type:"Struct"`
	// The information of instance storage usage.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbpgDiskUsagePercent *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent `json:"adbpg_disk_usage_percent,omitempty" xml:"adbpg_disk_usage_percent,omitempty" type:"Struct"`
	// The total amount of cold data storage.
	AdbpgInstanceColdDataGb *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb `json:"adbpg_instance_cold_data_gb,omitempty" xml:"adbpg_instance_cold_data_gb,omitempty" type:"Struct"`
	// The total amount of hot data storage.
	AdbpgInstanceHotDataGb *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb `json:"adbpg_instance_hot_data_gb,omitempty" xml:"adbpg_instance_hot_data_gb,omitempty" type:"Struct"`
	// The total amount of data storage of the instance.
	AdbpgInstanceTotalDataGb *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb `json:"adbpg_instance_total_data_gb,omitempty" xml:"adbpg_instance_total_data_gb,omitempty" type:"Struct"`
	// The information of maximum coordinator node storage usage.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	AdbpgMasterDiskUsagePercentMax *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax `json:"adbpg_master_disk_usage_percent_max,omitempty" xml:"adbpg_master_disk_usage_percent_max,omitempty" type:"Struct"`
	// The information of coordinator node availability status.
	AdbpgMasterStatus *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus `json:"adbpg_master_status,omitempty" xml:"adbpg_master_status,omitempty" type:"Struct"`
	// The information of compute node availability status.
	AdbpgSegmentStatus *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus `json:"adbpg_segment_status,omitempty" xml:"adbpg_segment_status,omitempty" type:"Struct"`
	// The information of instance health status.
	AdbpgStatus *DescribeHealthStatusResponseBodyStatusAdbpgStatus `json:"adbpg_status,omitempty" xml:"adbpg_status,omitempty" type:"Struct"`
	// The information of coordinator node connection health status.
	NodeMasterConnectionStatus *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus `json:"node_master_connection_status,omitempty" xml:"node_master_connection_status,omitempty" type:"Struct"`
	// The information of coordinator node health status.
	NodeMasterStatus *DescribeHealthStatusResponseBodyStatusNodeMasterStatus `json:"node_master_status,omitempty" xml:"node_master_status,omitempty" type:"Struct"`
	// The information of compute node connection health status.
	NodeSegmentConnectionStatus *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus `json:"node_segment_connection_status,omitempty" xml:"node_segment_connection_status,omitempty" type:"Struct"`
	// The information of compute node storage status.
	//
	// >  This parameter value is returned only for instances in elastic storage mode.
	NodeSegmentDiskStatus *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus `json:"node_segment_disk_status,omitempty" xml:"node_segment_disk_status,omitempty" type:"Struct"`
}

func (s DescribeHealthStatusResponseBodyStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbgpSegmentDiskUsagePercentMax() *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax {
	return s.AdbgpSegmentDiskUsagePercentMax
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgConnectionStatus() *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus {
	return s.AdbpgConnectionStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgDiskStatus() *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus {
	return s.AdbpgDiskStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgDiskUsagePercent() *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent {
	return s.AdbpgDiskUsagePercent
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgInstanceColdDataGb() *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb {
	return s.AdbpgInstanceColdDataGb
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgInstanceHotDataGb() *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb {
	return s.AdbpgInstanceHotDataGb
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgInstanceTotalDataGb() *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb {
	return s.AdbpgInstanceTotalDataGb
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgMasterDiskUsagePercentMax() *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax {
	return s.AdbpgMasterDiskUsagePercentMax
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgMasterStatus() *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus {
	return s.AdbpgMasterStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgSegmentStatus() *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus {
	return s.AdbpgSegmentStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetAdbpgStatus() *DescribeHealthStatusResponseBodyStatusAdbpgStatus {
	return s.AdbpgStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetNodeMasterConnectionStatus() *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus {
	return s.NodeMasterConnectionStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetNodeMasterStatus() *DescribeHealthStatusResponseBodyStatusNodeMasterStatus {
	return s.NodeMasterStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetNodeSegmentConnectionStatus() *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus {
	return s.NodeSegmentConnectionStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) GetNodeSegmentDiskStatus() *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus {
	return s.NodeSegmentDiskStatus
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbgpSegmentDiskUsagePercentMax(v *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) *DescribeHealthStatusResponseBodyStatus {
	s.AdbgpSegmentDiskUsagePercentMax = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgConnectionStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgDiskStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgDiskStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgDiskUsagePercent(v *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgDiskUsagePercent = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgInstanceColdDataGb(v *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgInstanceColdDataGb = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgInstanceHotDataGb(v *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgInstanceHotDataGb = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgInstanceTotalDataGb(v *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgInstanceTotalDataGb = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgMasterDiskUsagePercentMax(v *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgMasterDiskUsagePercentMax = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgMasterStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgMasterStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgSegmentStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgSegmentStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetAdbpgStatus(v *DescribeHealthStatusResponseBodyStatusAdbpgStatus) *DescribeHealthStatusResponseBodyStatus {
	s.AdbpgStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeMasterConnectionStatus(v *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeMasterConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeMasterStatus(v *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeMasterStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeSegmentConnectionStatus(v *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeSegmentConnectionStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) SetNodeSegmentDiskStatus(v *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) *DescribeHealthStatusResponseBodyStatus {
	s.NodeSegmentDiskStatus = v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatus) Validate() error {
	if s.AdbgpSegmentDiskUsagePercentMax != nil {
		if err := s.AdbgpSegmentDiskUsagePercentMax.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgConnectionStatus != nil {
		if err := s.AdbpgConnectionStatus.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgDiskStatus != nil {
		if err := s.AdbpgDiskStatus.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgDiskUsagePercent != nil {
		if err := s.AdbpgDiskUsagePercent.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgInstanceColdDataGb != nil {
		if err := s.AdbpgInstanceColdDataGb.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgInstanceHotDataGb != nil {
		if err := s.AdbpgInstanceHotDataGb.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgInstanceTotalDataGb != nil {
		if err := s.AdbpgInstanceTotalDataGb.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgMasterDiskUsagePercentMax != nil {
		if err := s.AdbpgMasterDiskUsagePercentMax.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgMasterStatus != nil {
		if err := s.AdbpgMasterStatus.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgSegmentStatus != nil {
		if err := s.AdbpgSegmentStatus.Validate(); err != nil {
			return err
		}
	}
	if s.AdbpgStatus != nil {
		if err := s.AdbpgStatus.Validate(); err != nil {
			return err
		}
	}
	if s.NodeMasterConnectionStatus != nil {
		if err := s.NodeMasterConnectionStatus.Validate(); err != nil {
			return err
		}
	}
	if s.NodeMasterStatus != nil {
		if err := s.NodeMasterStatus.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSegmentConnectionStatus != nil {
		if err := s.NodeSegmentConnectionStatus.Validate(); err != nil {
			return err
		}
	}
	if s.NodeSegmentDiskStatus != nil {
		if err := s.NodeSegmentDiskStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax struct {
	// The status corresponding to the maximum storage usage among all compute nodes. Valid values:
	//
	// 	- **critical**: The compute node storage usage is greater than or equal to 90%. In this case, the instance is locked.
	//
	// 	- **warning**: The compute node storage usage is greater than or equal to 80% and less than 90%.
	//
	// 	- **healthy**: The compute node storage usage is less than 80%.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum compute node storage usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 1.52
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbgpSegmentDiskUsagePercentMax) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus struct {
	// The connection health status of the instance. Valid values:
	//
	// 	- **critical**: The instance connection usage is greater than 95%. In this case, this metric is marked in red in the console.
	//
	// 	- **warning**: The instance connection usage is greater than 90% and less than or equal to 95%. In this case, this metric is marked in yellow in the console.
	//
	// 	- **healthy**: The instance connection usage is less than or equal to 90%. In this case, this metric is marked in green in the console.
	//
	// >  The instance connection usage is the maximum connection usage among all the coordinator and compute nodes.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance connection usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 1.71
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgConnectionStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus struct {
	// The storage status of the instance. Valid values:
	//
	// 	- **critical**: The instance storage usage is greater than or equal to 90%. In this case, this metric is marked in red in the console and the instance is locked.
	//
	// 	- **warning**: The instance storage usage is greater than or equal to 70% and less than 90%. In this case, this metric is marked in yellow in the console.
	//
	// 	- **healthy**: The instance storage usage is less than 70%. In this case, this metric is marked in green in the console.
	//
	// >  The instance storage usage is the average storage usage of all compute nodes.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance storage usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 1.52
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent struct {
	// The status corresponding to the storage usage of the instance. Valid values:
	//
	// 	- **critical**: The instance storage usage is greater than or equal to 90%. In this case, the instance is locked.
	//
	// 	- **warning**: The instance storage usage is greater than or equal to 70% and less than 90%.
	//
	// 	- **healthy**: The instance storage usage is less than 70%.
	//
	// >  The instance storage usage is the average storage usage of all compute nodes.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance storage usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 1.52
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgDiskUsagePercent) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb struct {
	// The total amount of cold data storage. Unit: GB.
	//
	// example:
	//
	// 0
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceColdDataGb) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb struct {
	// The total amount of hot data storage. Unit: GB.
	//
	// example:
	//
	// 4.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceHotDataGb) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb struct {
	// The total amount of data storage of the instance. Unit: GB.
	//
	// example:
	//
	// 4.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgInstanceTotalDataGb) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax struct {
	// The status corresponding to the maximum storage usage of the coordinator node. Valid values:
	//
	// 	- **critical**: The coordinator node storage usage is greater than or equal to 90%. In this case, the instance is locked.
	//
	// 	- **warning**: The coordinator node storage usage is greater than or equal to 70% and less than 90%.
	//
	// 	- **healthy**: The coordinator node storage usage is less than 70%.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum coordinator node storage usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 1.34
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterDiskUsagePercentMax) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus struct {
	// The availability status of the coordinator node. Valid values:
	//
	// 	- **critical**: Both the primary and standby coordinator nodes are unavailable. In this case, this metric is marked in red in the console.
	//
	// 	- **warning**: The primary or standby coordinator node is unavailable. In this case, this metric is marked in yellow in the console.
	//
	// 	- **healthy**: Both the primary and standby coordinator nodes are available. In this case, this metric is marked in green in the console.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of coordinator node availability status. Valid values:
	//
	// example:
	//
	// 1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgMasterStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus struct {
	// The availability status of compute nodes. Valid values:
	//
	// 	- **critical**: All the primary and secondary compute nodes are unavailable. In this case, this metric is marked in red in the console.
	//
	// 	- **warning**: Fifty percent or more than fifty percent of compute nodes are unavailable. In this case, this metric is marked in yellow in the console.
	//
	// 	- **healthy**: All compute nodes are available. In this case, this metric is marked in green in the console.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of compute node availability status.
	//
	// example:
	//
	// 1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgSegmentStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusAdbpgStatus struct {
	// The health status of the instance. Valid values:
	//
	// 	- **critical**: The coordinator node or a compute node is unavailable. In this case, this metric is marked in red in the console.
	//
	// 	- **healthy**: All nodes are available. In this case, this metric is marked in green in the console.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of instance health status. Valid values:
	//
	// 	- **1**: healthy
	//
	// 	- **0**: critical
	//
	// example:
	//
	// 1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusAdbpgStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusAdbpgStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusAdbpgStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusAdbpgStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus struct {
	// The connection health status of the coordinator node. Valid values:
	//
	// 	- **critical**: The coordinator node connection usage is greater than 95%. In this case, this metric is marked in red in the console.
	//
	// 	- **warning**: The coordinator node connection usage is greater than or equal to 90% and less than 95%. In this case, this metric is marked in yellow in the console.
	//
	// 	- **healthy**: The coordinator node connection usage is less than 90%. In this case, this metric is marked in green in the console.
	//
	// >  The coordinator node connection usage is the maximum connection usage of the coordinator node.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of coordinator node connection usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 1.71
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterConnectionStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusNodeMasterStatus struct {
	// The health status of the coordinator node. Valid values:
	//
	// 	- **critical**: The primary or standby coordinator node is unavailable. In this case, this metric is marked in red in the console.
	//
	// 	- **healthy**: Both the primary and standby coordinator nodes are available. In this case, this metric is marked in green in the console.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of coordinator node health status. Valid values:
	//
	// 	- **1**: healthy
	//
	// 	- **0**: critical
	//
	// example:
	//
	// 1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeMasterStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeMasterStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeMasterStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeMasterStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus struct {
	// The connection health status of compute nodes. Valid values:
	//
	// 	- **critical**: The compute node connection usage is greater than or equal to 95%. In this case, this metric is marked in red in the console.
	//
	// 	- **warning**: The compute node connection usage is greater than or equal to 90% and less than 95%. In this case, this metric is marked in yellow in the console.
	//
	// 	- **healthy**: The compute node connection usage is less than 90%. In this case, this metric is marked in green in the console.
	//
	// >  The compute node connection usage is the maximum connection usage among all compute nodes.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum compute node connection usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 0.48
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentConnectionStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus struct {
	// The storage status of compute nodes. Valid values:
	//
	// 	- **critical**: The compute node storage usage is greater than or equal to 90%. In this case, this metric is marked in red in the console and the instance is locked.
	//
	// 	- **warning**: The compute node storage usage is greater than or equal to 80% and less than 90%. In this case, this metric is marked in yellow in the console.
	//
	// 	- **healthy**: The compute node storage usage is less than 80%. In this case, this metric is marked in green in the console.
	//
	// >  The compute node storage usage is the maximum storage usage among all compute nodes.
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metric value of maximum compute node storage usage.
	//
	// Unit: %.
	//
	// example:
	//
	// 1.52
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) GoString() string {
	return s.String()
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) GetValue() *float32 {
	return s.Value
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) SetStatus(v string) *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus {
	s.Status = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) SetValue(v float32) *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus {
	s.Value = &v
	return s
}

func (s *DescribeHealthStatusResponseBodyStatusNodeSegmentDiskStatus) Validate() error {
	return dara.Validate(s)
}
