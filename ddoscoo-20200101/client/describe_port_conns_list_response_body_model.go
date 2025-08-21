// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortConnsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnsList(v []*DescribePortConnsListResponseBodyConnsList) *DescribePortConnsListResponseBody
	GetConnsList() []*DescribePortConnsListResponseBodyConnsList
	SetRequestId(v string) *DescribePortConnsListResponseBody
	GetRequestId() *string
}

type DescribePortConnsListResponseBody struct {
	// Details about the connections established over the port.
	ConnsList []*DescribePortConnsListResponseBodyConnsList `json:"ConnsList,omitempty" xml:"ConnsList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6D48AED0-41DB-5D9B-B484-3B6AAD312AD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortConnsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortConnsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListResponseBody) GetConnsList() []*DescribePortConnsListResponseBodyConnsList {
	return s.ConnsList
}

func (s *DescribePortConnsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortConnsListResponseBody) SetConnsList(v []*DescribePortConnsListResponseBodyConnsList) *DescribePortConnsListResponseBody {
	s.ConnsList = v
	return s
}

func (s *DescribePortConnsListResponseBody) SetRequestId(v string) *DescribePortConnsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortConnsListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePortConnsListResponseBodyConnsList struct {
	// The number of active connections.
	//
	// example:
	//
	// 3
	ActConns *int64 `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	// >  This parameter is in internal preview. Do not use this parameter.
	//
	// example:
	//
	// 8
	Conns *int64 `json:"Conns,omitempty" xml:"Conns,omitempty"`
	// The number of new connections.
	//
	// example:
	//
	// 0
	Cps *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The number of inactive connections.
	//
	// example:
	//
	// 2
	InActConns *int64 `json:"InActConns,omitempty" xml:"InActConns,omitempty"`
	// The index number of the returned data.
	//
	// example:
	//
	// 16506
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s DescribePortConnsListResponseBodyConnsList) String() string {
	return dara.Prettify(s)
}

func (s DescribePortConnsListResponseBodyConnsList) GoString() string {
	return s.String()
}

func (s *DescribePortConnsListResponseBodyConnsList) GetActConns() *int64 {
	return s.ActConns
}

func (s *DescribePortConnsListResponseBodyConnsList) GetConns() *int64 {
	return s.Conns
}

func (s *DescribePortConnsListResponseBodyConnsList) GetCps() *int64 {
	return s.Cps
}

func (s *DescribePortConnsListResponseBodyConnsList) GetInActConns() *int64 {
	return s.InActConns
}

func (s *DescribePortConnsListResponseBodyConnsList) GetIndex() *int64 {
	return s.Index
}

func (s *DescribePortConnsListResponseBodyConnsList) SetActConns(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.ActConns = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetConns(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.Conns = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetCps(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.Cps = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetInActConns(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.InActConns = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) SetIndex(v int64) *DescribePortConnsListResponseBodyConnsList {
	s.Index = &v
	return s
}

func (s *DescribePortConnsListResponseBodyConnsList) Validate() error {
	return dara.Validate(s)
}
