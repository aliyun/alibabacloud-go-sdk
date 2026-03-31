// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudSdkServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeHybridCloudSdkServersResponseBody
	GetRequestId() *string
	SetSdkServers(v []*DescribeHybridCloudSdkServersResponseBodySdkServers) *DescribeHybridCloudSdkServersResponseBody
	GetSdkServers() []*DescribeHybridCloudSdkServersResponseBodySdkServers
	SetTotalCount(v int32) *DescribeHybridCloudSdkServersResponseBody
	GetTotalCount() *int32
}

type DescribeHybridCloudSdkServersResponseBody struct {
	// example:
	//
	// 3600F008-2E76-5D0B-BC76-EFBD****6D
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkServers []*DescribeHybridCloudSdkServersResponseBodySdkServers `json:"SdkServers,omitempty" xml:"SdkServers,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudSdkServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudSdkServersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudSdkServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudSdkServersResponseBody) GetSdkServers() []*DescribeHybridCloudSdkServersResponseBodySdkServers {
	return s.SdkServers
}

func (s *DescribeHybridCloudSdkServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHybridCloudSdkServersResponseBody) SetRequestId(v string) *DescribeHybridCloudSdkServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBody) SetSdkServers(v []*DescribeHybridCloudSdkServersResponseBodySdkServers) *DescribeHybridCloudSdkServersResponseBody {
	s.SdkServers = v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBody) SetTotalCount(v int32) *DescribeHybridCloudSdkServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBody) Validate() error {
	if s.SdkServers != nil {
		for _, item := range s.SdkServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHybridCloudSdkServersResponseBodySdkServers struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// 1621428205000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// qsh5-sec-8-fedd**005
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// example:
	//
	// 127.0.0.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// SDKID。
	//
	// example:
	//
	// b11327c21790846374051d5d**83c
	Mid                    *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	ProtectionGroupAddress *string `json:"ProtectionGroupAddress,omitempty" xml:"ProtectionGroupAddress,omitempty"`
	// example:
	//
	// on
	PullinStatus *string `json:"PullinStatus,omitempty" xml:"PullinStatus,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1719489906000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeHybridCloudSdkServersResponseBodySdkServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudSdkServersResponseBodySdkServers) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetHostName() *string {
	return s.HostName
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetIp() *string {
	return s.Ip
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetMid() *string {
	return s.Mid
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetProtectionGroupAddress() *string {
	return s.ProtectionGroupAddress
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetPullinStatus() *string {
	return s.PullinStatus
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetStatus() *string {
	return s.Status
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetClusterName(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetCreateTime(v int64) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.CreateTime = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetHostName(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.HostName = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetIp(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.Ip = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetMid(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.Mid = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetProtectionGroupAddress(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.ProtectionGroupAddress = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetPullinStatus(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.PullinStatus = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetResourceId(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.ResourceId = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetStatus(v string) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.Status = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) SetUpdateTime(v int64) *DescribeHybridCloudSdkServersResponseBodySdkServers {
	s.UpdateTime = &v
	return s
}

func (s *DescribeHybridCloudSdkServersResponseBodySdkServers) Validate() error {
	return dara.Validate(s)
}
