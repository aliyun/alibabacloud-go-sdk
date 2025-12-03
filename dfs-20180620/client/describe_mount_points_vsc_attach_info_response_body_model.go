// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountPointsVscAttachInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachInfos(v []*DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) *DescribeMountPointsVscAttachInfoResponseBody
	GetAttachInfos() []*DescribeMountPointsVscAttachInfoResponseBodyAttachInfos
	SetMaxResults(v int32) *DescribeMountPointsVscAttachInfoResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeMountPointsVscAttachInfoResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeMountPointsVscAttachInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeMountPointsVscAttachInfoResponseBody
	GetTotalCount() *string
}

type DescribeMountPointsVscAttachInfoResponseBody struct {
	AttachInfos []*DescribeMountPointsVscAttachInfoResponseBodyAttachInfos `json:"AttachInfos,omitempty" xml:"AttachInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// asd0daj****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountPointsVscAttachInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountPointsVscAttachInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) GetAttachInfos() []*DescribeMountPointsVscAttachInfoResponseBodyAttachInfos {
	return s.AttachInfos
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) SetAttachInfos(v []*DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) *DescribeMountPointsVscAttachInfoResponseBody {
	s.AttachInfos = v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) SetMaxResults(v int32) *DescribeMountPointsVscAttachInfoResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) SetNextToken(v string) *DescribeMountPointsVscAttachInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) SetRequestId(v string) *DescribeMountPointsVscAttachInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) SetTotalCount(v string) *DescribeMountPointsVscAttachInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBody) Validate() error {
	if s.AttachInfos != nil {
		for _, item := range s.AttachInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMountPointsVscAttachInfoResponseBodyAttachInfos struct {
	// example:
	//
	// vsc-agent-f5a9bb4b041541f595a0c239c9e0f971-cn-shanghai
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// f-9dd3c6bajmy110.cn-zhangjiakou.dfs.aliyuncs.com
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// vsc-bp18fhqie89loyqixyieal
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// example:
	//
	// cxc
	VscName   *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscStatus *string `json:"VscStatus,omitempty" xml:"VscStatus,omitempty"`
	// example:
	//
	// Primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) GoString() string {
	return s.String()
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) GetMountPointId() *string {
	return s.MountPointId
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) GetVscId() *string {
	return s.VscId
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) GetVscName() *string {
	return s.VscName
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) GetVscStatus() *string {
	return s.VscStatus
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) GetVscType() *string {
	return s.VscType
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) SetInstanceId(v string) *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) SetMountPointId(v string) *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos {
	s.MountPointId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) SetVscId(v string) *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos {
	s.VscId = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) SetVscName(v string) *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos {
	s.VscName = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) SetVscStatus(v string) *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos {
	s.VscStatus = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) SetVscType(v string) *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos {
	s.VscType = &v
	return s
}

func (s *DescribeMountPointsVscAttachInfoResponseBodyAttachInfos) Validate() error {
	return dara.Validate(s)
}
