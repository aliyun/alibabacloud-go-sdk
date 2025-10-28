// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetK8sClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterPage(v *GetK8sClusterResponseBodyClusterPage) *GetK8sClusterResponseBody
	GetClusterPage() *GetK8sClusterResponseBodyClusterPage
	SetCode(v int32) *GetK8sClusterResponseBody
	GetCode() *int32
	SetMessage(v string) *GetK8sClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetK8sClusterResponseBody
	GetRequestId() *string
}

type GetK8sClusterResponseBody struct {
	// The cluster data that is returned by page.
	ClusterPage *GetK8sClusterResponseBodyClusterPage `json:"ClusterPage,omitempty" xml:"ClusterPage,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C3CE915C-0C83-4AA5-8D66-E8BEED62939E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetK8sClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetK8sClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBody) GetClusterPage() *GetK8sClusterResponseBodyClusterPage {
	return s.ClusterPage
}

func (s *GetK8sClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetK8sClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetK8sClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetK8sClusterResponseBody) SetClusterPage(v *GetK8sClusterResponseBodyClusterPage) *GetK8sClusterResponseBody {
	s.ClusterPage = v
	return s
}

func (s *GetK8sClusterResponseBody) SetCode(v int32) *GetK8sClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetK8sClusterResponseBody) SetMessage(v string) *GetK8sClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetK8sClusterResponseBody) SetRequestId(v string) *GetK8sClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetK8sClusterResponseBody) Validate() error {
	if s.ClusterPage != nil {
		if err := s.ClusterPage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetK8sClusterResponseBodyClusterPage struct {
	// The list of clusters.
	ClusterList *GetK8sClusterResponseBodyClusterPageClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	// The number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: 1000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages that are returned.
	//
	// example:
	//
	// 5
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetK8sClusterResponseBodyClusterPage) String() string {
	return dara.Prettify(s)
}

func (s GetK8sClusterResponseBodyClusterPage) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBodyClusterPage) GetClusterList() *GetK8sClusterResponseBodyClusterPageClusterList {
	return s.ClusterList
}

func (s *GetK8sClusterResponseBodyClusterPage) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetK8sClusterResponseBodyClusterPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetK8sClusterResponseBodyClusterPage) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetK8sClusterResponseBodyClusterPage) SetClusterList(v *GetK8sClusterResponseBodyClusterPageClusterList) *GetK8sClusterResponseBodyClusterPage {
	s.ClusterList = v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPage) SetCurrentPage(v int32) *GetK8sClusterResponseBodyClusterPage {
	s.CurrentPage = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPage) SetPageSize(v int32) *GetK8sClusterResponseBodyClusterPage {
	s.PageSize = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPage) SetTotalSize(v int32) *GetK8sClusterResponseBodyClusterPage {
	s.TotalSize = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPage) Validate() error {
	if s.ClusterList != nil {
		if err := s.ClusterList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetK8sClusterResponseBodyClusterPageClusterList struct {
	Cluster []*GetK8sClusterResponseBodyClusterPageClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s GetK8sClusterResponseBodyClusterPageClusterList) String() string {
	return dara.Prettify(s)
}

func (s GetK8sClusterResponseBodyClusterPageClusterList) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBodyClusterPageClusterList) GetCluster() []*GetK8sClusterResponseBodyClusterPageClusterListCluster {
	return s.Cluster
}

func (s *GetK8sClusterResponseBodyClusterPageClusterList) SetCluster(v []*GetK8sClusterResponseBodyClusterPageClusterListCluster) *GetK8sClusterResponseBodyClusterPageClusterList {
	s.Cluster = v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterList) Validate() error {
	if s.Cluster != nil {
		for _, item := range s.Cluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetK8sClusterResponseBodyClusterPageClusterListCluster struct {
	// The ID of the cluster.
	//
	// example:
	//
	// 81453e4b-4df0-4592-****-b835a2ee****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The import state of the cluster. Valid values:
	//
	// 	- 0: The cluster is not imported.
	//
	// 	- 1: The cluster is imported.
	//
	// 	- 2: The cluster fails to be imported.
	//
	// 	- 3: The cluster is being imported.
	//
	// 	- 4: The cluster is deleted.
	//
	// example:
	//
	// 1
	ClusterImportStatus *int32 `json:"ClusterImportStatus,omitempty" xml:"ClusterImportStatus,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- 1: The cluster runs as expected.
	//
	// 	- 2: The cluster does not run as expected.
	//
	// 	- 3: The cluster is offline.
	//
	// example:
	//
	// 1
	ClusterStatus *int32 `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- 2: Elastic Compute Service (ECS) cluster
	//
	// 	- 5: ACK cluster or Serverless Kubernetes cluster
	//
	// example:
	//
	// 5
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The ID of the ACK cluster.
	//
	// example:
	//
	// 2ce62869f4d4466b920312315f05****
	CsClusterId *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty"`
	// The state of the ACK cluster. Valid values:
	//
	// 	- initial: The cluster is being initialized.
	//
	// 	- failed: The cluster fails to be created.
	//
	// 	- running: The cluster is running.
	//
	// 	- updating: The cluster is being updated.
	//
	// 	- scaling: The cluster is being scaled out.
	//
	// 	- removing: Nodes are being removed from the cluster.
	//
	// 	- upgrading: The cluster is being upgraded.
	//
	// 	- deleting: The cluster is being deleted.
	//
	// 	- delete_failed: The cluster fails to be deleted.
	//
	// 	- deleted: The cluster is deleted. The deleted cluster is invisible to users.
	//
	// example:
	//
	// running
	CsClusterStatus *string `json:"CsClusterStatus,omitempty" xml:"CsClusterStatus,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total size of memory. Unit: MB.
	//
	// example:
	//
	// 2048
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- 1: classic network
	//
	// 	- 2: VPC
	//
	// example:
	//
	// 2
	NetworkMode *int32 `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 4
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// test
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subtype of the cluster. Valid values:
	//
	// 	- Ask: Serverless Kubernetes cluster
	//
	// 	- ManagedKubernetes: ACK cluster
	//
	// example:
	//
	// Ask
	SubClusterType *string `json:"SubClusterType,omitempty" xml:"SubClusterType,omitempty"`
	// The CIDR block of the subnet.
	//
	// example:
	//
	// 172.20.0.0/16
	SubNetCidr *string `json:"SubNetCidr,omitempty" xml:"SubNetCidr,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-**z1mlwpbjx3e9m**
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1uf97****xjxgip****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s GetK8sClusterResponseBodyClusterPageClusterListCluster) String() string {
	return dara.Prettify(s)
}

func (s GetK8sClusterResponseBodyClusterPageClusterListCluster) GoString() string {
	return s.String()
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetClusterImportStatus() *int32 {
	return s.ClusterImportStatus
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetClusterStatus() *int32 {
	return s.ClusterStatus
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetCsClusterStatus() *string {
	return s.CsClusterStatus
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetDescription() *string {
	return s.Description
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetMem() *int32 {
	return s.Mem
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetNetworkMode() *int32 {
	return s.NetworkMode
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetSubClusterType() *string {
	return s.SubClusterType
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetSubNetCidr() *string {
	return s.SubNetCidr
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterImportStatus(v int32) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterImportStatus = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterName(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterStatus(v int32) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterStatus = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetClusterType(v int32) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetCpu(v int32) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.Cpu = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetCsClusterId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.CsClusterId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetCsClusterStatus(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.CsClusterStatus = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetDescription(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.Description = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetMem(v int32) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.Mem = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetNetworkMode(v int32) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.NetworkMode = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetNodeNum(v int32) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.NodeNum = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetRegionId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetSubClusterType(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.SubClusterType = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetSubNetCidr(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.SubNetCidr = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetVpcId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.VpcId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) SetVswitchId(v string) *GetK8sClusterResponseBodyClusterPageClusterListCluster {
	s.VswitchId = &v
	return s
}

func (s *GetK8sClusterResponseBodyClusterPageClusterListCluster) Validate() error {
	return dara.Validate(s)
}
