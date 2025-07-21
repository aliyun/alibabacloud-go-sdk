// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListVirtualHostsResponseBodyData) *ListVirtualHostsResponseBody
	GetData() *ListVirtualHostsResponseBodyData
	SetRequestId(v string) *ListVirtualHostsResponseBody
	GetRequestId() *string
}

type ListVirtualHostsResponseBody struct {
	// The returned data.
	Data *ListVirtualHostsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EF4DB019-DA4A-4CE3-B220-223BBC93F***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirtualHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBody) GetData() *ListVirtualHostsResponseBodyData {
	return s.Data
}

func (s *ListVirtualHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirtualHostsResponseBody) SetData(v *ListVirtualHostsResponseBodyData) *ListVirtualHostsResponseBody {
	s.Data = v
	return s
}

func (s *ListVirtualHostsResponseBody) SetRequestId(v string) *ListVirtualHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualHostsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVirtualHostsResponseBodyData struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The vhosts.
	VirtualHosts []*ListVirtualHostsResponseBodyDataVirtualHosts `json:"VirtualHosts,omitempty" xml:"VirtualHosts,omitempty" type:"Repeated"`
}

func (s ListVirtualHostsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualHostsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVirtualHostsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVirtualHostsResponseBodyData) GetVirtualHosts() []*ListVirtualHostsResponseBodyDataVirtualHosts {
	return s.VirtualHosts
}

func (s *ListVirtualHostsResponseBodyData) SetMaxResults(v int32) *ListVirtualHostsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualHostsResponseBodyData) SetNextToken(v string) *ListVirtualHostsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListVirtualHostsResponseBodyData) SetVirtualHosts(v []*ListVirtualHostsResponseBodyDataVirtualHosts) *ListVirtualHostsResponseBodyData {
	s.VirtualHosts = v
	return s
}

func (s *ListVirtualHostsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListVirtualHostsResponseBodyDataVirtualHosts struct {
	// The vhost name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListVirtualHostsResponseBodyDataVirtualHosts) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualHostsResponseBodyDataVirtualHosts) GoString() string {
	return s.String()
}

func (s *ListVirtualHostsResponseBodyDataVirtualHosts) GetName() *string {
	return s.Name
}

func (s *ListVirtualHostsResponseBodyDataVirtualHosts) SetName(v string) *ListVirtualHostsResponseBodyDataVirtualHosts {
	s.Name = &v
	return s
}

func (s *ListVirtualHostsResponseBodyDataVirtualHosts) Validate() error {
	return dara.Validate(s)
}
