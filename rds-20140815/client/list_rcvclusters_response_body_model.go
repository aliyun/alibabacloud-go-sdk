// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRCVClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRCVClustersResponseBody
	GetRequestId() *string
	SetVClusters(v []*ListRCVClustersResponseBodyVClusters) *ListRCVClustersResponseBody
	GetVClusters() []*ListRCVClustersResponseBodyVClusters
}

type ListRCVClustersResponseBody struct {
	// example:
	//
	// 07F6177E-6DE4-408A-BB4F-0723301340F3
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VClusters []*ListRCVClustersResponseBodyVClusters `json:"VClusters,omitempty" xml:"VClusters,omitempty" type:"Repeated"`
}

func (s ListRCVClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRCVClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRCVClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRCVClustersResponseBody) GetVClusters() []*ListRCVClustersResponseBodyVClusters {
	return s.VClusters
}

func (s *ListRCVClustersResponseBody) SetRequestId(v string) *ListRCVClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRCVClustersResponseBody) SetVClusters(v []*ListRCVClustersResponseBodyVClusters) *ListRCVClustersResponseBody {
	s.VClusters = v
	return s
}

func (s *ListRCVClustersResponseBody) Validate() error {
	if s.VClusters != nil {
		for _, item := range s.VClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRCVClustersResponseBodyVClusters struct {
	// example:
	//
	// cd21387ea640145bab79a78276c1a****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	InstanceCount *int64 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId                    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SupportDiskPerformanceLevel []*string `json:"SupportDiskPerformanceLevel,omitempty" xml:"SupportDiskPerformanceLevel,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-2zeqj40j2ce0s5yhg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListRCVClustersResponseBodyVClusters) String() string {
	return dara.Prettify(s)
}

func (s ListRCVClustersResponseBodyVClusters) GoString() string {
	return s.String()
}

func (s *ListRCVClustersResponseBodyVClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListRCVClustersResponseBodyVClusters) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListRCVClustersResponseBodyVClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRCVClustersResponseBodyVClusters) GetSupportDiskPerformanceLevel() []*string {
	return s.SupportDiskPerformanceLevel
}

func (s *ListRCVClustersResponseBodyVClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListRCVClustersResponseBodyVClusters) SetClusterId(v string) *ListRCVClustersResponseBodyVClusters {
	s.ClusterId = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetInstanceCount(v int64) *ListRCVClustersResponseBodyVClusters {
	s.InstanceCount = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetRegionId(v string) *ListRCVClustersResponseBodyVClusters {
	s.RegionId = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetSupportDiskPerformanceLevel(v []*string) *ListRCVClustersResponseBodyVClusters {
	s.SupportDiskPerformanceLevel = v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetVpcId(v string) *ListRCVClustersResponseBodyVClusters {
	s.VpcId = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) Validate() error {
	return dara.Validate(s)
}
