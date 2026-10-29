// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstances(v []*ListRenderingInstancesResponseBodyRenderingInstances) *ListRenderingInstancesResponseBody
	GetRenderingInstances() []*ListRenderingInstancesResponseBodyRenderingInstances
	SetRequestId(v string) *ListRenderingInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRenderingInstancesResponseBody
	GetTotalCount() *int64
}

type ListRenderingInstancesResponseBody struct {
	// The list of cloud application service instances.
	RenderingInstances []*ListRenderingInstancesResponseBodyRenderingInstances `json:"RenderingInstances,omitempty" xml:"RenderingInstances,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of cloud application service instances.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingInstancesResponseBody) GetRenderingInstances() []*ListRenderingInstancesResponseBodyRenderingInstances {
	return s.RenderingInstances
}

func (s *ListRenderingInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRenderingInstancesResponseBody) SetRenderingInstances(v []*ListRenderingInstancesResponseBodyRenderingInstances) *ListRenderingInstancesResponseBody {
	s.RenderingInstances = v
	return s
}

func (s *ListRenderingInstancesResponseBody) SetRequestId(v string) *ListRenderingInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingInstancesResponseBody) SetTotalCount(v int64) *ListRenderingInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingInstancesResponseBody) Validate() error {
	if s.RenderingInstances != nil {
		for _, item := range s.RenderingInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRenderingInstancesResponseBodyRenderingInstances struct {
	// The billing method.
	//
	// example:
	//
	// Postpaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the cloud application service instance was created. The time is in UTC and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2023-11-17T02:18:04Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2023-11-17T02:18:04Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// TestInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum outbound bandwidth limit. Unit: Mbit/s. A value of 0 indicates that no bandwidth limit is configured.
	//
	// example:
	//
	// 20
	MaxEgressBandwidth *string `json:"MaxEgressBandwidth,omitempty" xml:"MaxEgressBandwidth,omitempty"`
	// The maximum inbound bandwidth limit. Unit: Mbit/s. A value of 0 indicates that no bandwidth limit is configured.
	//
	// example:
	//
	// 10
	MaxIngressBandwidth *string `json:"MaxIngressBandwidth,omitempty" xml:"MaxIngressBandwidth,omitempty"`
	// The instance ID of the cloud application service.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The instance type of the cloud application service instance.
	//
	// example:
	//
	// crs.cp.l1
	RenderingSpec *string `json:"RenderingSpec,omitempty" xml:"RenderingSpec,omitempty"`
	// The status of the service host.
	//
	// example:
	//
	// running
	ServerStatus *string `json:"ServerStatus,omitempty" xml:"ServerStatus,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cloud storage capacity. Unit: GB.
	//
	// example:
	//
	// 100
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s ListRenderingInstancesResponseBodyRenderingInstances) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingInstancesResponseBodyRenderingInstances) GoString() string {
	return s.String()
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetMaxEgressBandwidth() *string {
	return s.MaxEgressBandwidth
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetMaxIngressBandwidth() *string {
	return s.MaxIngressBandwidth
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetRenderingSpec() *string {
	return s.RenderingSpec
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetServerStatus() *string {
	return s.ServerStatus
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetStatus() *string {
	return s.Status
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetChargeType(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.ChargeType = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetCreationTime(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.CreationTime = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetExpireTime(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.ExpireTime = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetInstanceName(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.InstanceName = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetMaxEgressBandwidth(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.MaxEgressBandwidth = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetMaxIngressBandwidth(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.MaxIngressBandwidth = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetRenderingInstanceId(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetRenderingSpec(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.RenderingSpec = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetServerStatus(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.ServerStatus = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetStatus(v string) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.Status = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) SetStorageSize(v int32) *ListRenderingInstancesResponseBodyRenderingInstances {
	s.StorageSize = &v
	return s
}

func (s *ListRenderingInstancesResponseBodyRenderingInstances) Validate() error {
	return dara.Validate(s)
}
