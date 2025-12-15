// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodesRequest
	GetClusterId() *string
	SetHostnames(v []*string) *ListNodesRequest
	GetHostnames() []*string
	SetPageNumber(v int32) *ListNodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodesRequest
	GetPageSize() *int32
	SetPrivateIpAddress(v []*string) *ListNodesRequest
	GetPrivateIpAddress() []*string
	SetQueueNames(v []*string) *ListNodesRequest
	GetQueueNames() []*string
	SetSequence(v string) *ListNodesRequest
	GetSequence() *string
	SetSortBy(v string) *ListNodesRequest
	GetSortBy() *string
	SetStatus(v []*string) *ListNodesRequest
	GetStatus() []*string
}

type ListNodesRequest struct {
	// The cluster ID. You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The hostnames of the compute nodes that you want to query.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IP addresses of the compute nodes that you want to query.
	PrivateIpAddress []*string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty" type:"Repeated"`
	// The queues to which the nodes belong.
	QueueNames []*string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty" type:"Repeated"`
	// Specifies whether the results are sorted in ascending or descending order. Valid values:
	//
	// 	- Forward: ascending
	//
	// 	- Backward: descending
	//
	// Default value: Forward.
	//
	// example:
	//
	// Forward
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The sorting method of the node list. Valid values:
	//
	// 	- AddedTime: sorts the nodes by the time that they were added.
	//
	// 	- HostName: sorts the nodes by their hostnames.
	//
	// Default value: addedtime.
	//
	// example:
	//
	// AddedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The states of the compute nodes to be queried.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodesRequest) GetHostnames() []*string {
	return s.Hostnames
}

func (s *ListNodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesRequest) GetPrivateIpAddress() []*string {
	return s.PrivateIpAddress
}

func (s *ListNodesRequest) GetQueueNames() []*string {
	return s.QueueNames
}

func (s *ListNodesRequest) GetSequence() *string {
	return s.Sequence
}

func (s *ListNodesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListNodesRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListNodesRequest) SetClusterId(v string) *ListNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesRequest) SetHostnames(v []*string) *ListNodesRequest {
	s.Hostnames = v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetPrivateIpAddress(v []*string) *ListNodesRequest {
	s.PrivateIpAddress = v
	return s
}

func (s *ListNodesRequest) SetQueueNames(v []*string) *ListNodesRequest {
	s.QueueNames = v
	return s
}

func (s *ListNodesRequest) SetSequence(v string) *ListNodesRequest {
	s.Sequence = &v
	return s
}

func (s *ListNodesRequest) SetSortBy(v string) *ListNodesRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesRequest) SetStatus(v []*string) *ListNodesRequest {
	s.Status = v
	return s
}

func (s *ListNodesRequest) Validate() error {
	return dara.Validate(s)
}
