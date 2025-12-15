// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodesShrinkRequest
	GetClusterId() *string
	SetHostnamesShrink(v string) *ListNodesShrinkRequest
	GetHostnamesShrink() *string
	SetPageNumber(v int32) *ListNodesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListNodesShrinkRequest
	GetPageSize() *int32
	SetPrivateIpAddressShrink(v string) *ListNodesShrinkRequest
	GetPrivateIpAddressShrink() *string
	SetQueueNamesShrink(v string) *ListNodesShrinkRequest
	GetQueueNamesShrink() *string
	SetSequence(v string) *ListNodesShrinkRequest
	GetSequence() *string
	SetSortBy(v string) *ListNodesShrinkRequest
	GetSortBy() *string
	SetStatusShrink(v string) *ListNodesShrinkRequest
	GetStatusShrink() *string
}

type ListNodesShrinkRequest struct {
	// The cluster ID. You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The hostnames of the compute nodes that you want to query.
	HostnamesShrink *string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty"`
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
	PrivateIpAddressShrink *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The queues to which the nodes belong.
	QueueNamesShrink *string `json:"QueueNames,omitempty" xml:"QueueNames,omitempty"`
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
	StatusShrink *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodesShrinkRequest) GetHostnamesShrink() *string {
	return s.HostnamesShrink
}

func (s *ListNodesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNodesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNodesShrinkRequest) GetPrivateIpAddressShrink() *string {
	return s.PrivateIpAddressShrink
}

func (s *ListNodesShrinkRequest) GetQueueNamesShrink() *string {
	return s.QueueNamesShrink
}

func (s *ListNodesShrinkRequest) GetSequence() *string {
	return s.Sequence
}

func (s *ListNodesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListNodesShrinkRequest) GetStatusShrink() *string {
	return s.StatusShrink
}

func (s *ListNodesShrinkRequest) SetClusterId(v string) *ListNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesShrinkRequest) SetHostnamesShrink(v string) *ListNodesShrinkRequest {
	s.HostnamesShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPageNumber(v int32) *ListNodesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPageSize(v int32) *ListNodesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesShrinkRequest) SetPrivateIpAddressShrink(v string) *ListNodesShrinkRequest {
	s.PrivateIpAddressShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetQueueNamesShrink(v string) *ListNodesShrinkRequest {
	s.QueueNamesShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) SetSequence(v string) *ListNodesShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *ListNodesShrinkRequest) SetSortBy(v string) *ListNodesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesShrinkRequest) SetStatusShrink(v string) *ListNodesShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *ListNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
