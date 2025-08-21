// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortConnsCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActConns(v int64) *DescribePortConnsCountResponseBody
	GetActConns() *int64
	SetConns(v int64) *DescribePortConnsCountResponseBody
	GetConns() *int64
	SetCps(v int64) *DescribePortConnsCountResponseBody
	GetCps() *int64
	SetInActConns(v int64) *DescribePortConnsCountResponseBody
	GetInActConns() *int64
	SetRequestId(v string) *DescribePortConnsCountResponseBody
	GetRequestId() *string
}

type DescribePortConnsCountResponseBody struct {
	// The number of active connections.
	//
	// example:
	//
	// 159
	ActConns *int64 `json:"ActConns,omitempty" xml:"ActConns,omitempty"`
	// The number of concurrent connections.
	//
	// example:
	//
	// 46340
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
	// 121
	InActConns *int64 `json:"InActConns,omitempty" xml:"InActConns,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 48859E14-A9FB-4100-99FF-AAB75CA46776
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortConnsCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortConnsCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortConnsCountResponseBody) GetActConns() *int64 {
	return s.ActConns
}

func (s *DescribePortConnsCountResponseBody) GetConns() *int64 {
	return s.Conns
}

func (s *DescribePortConnsCountResponseBody) GetCps() *int64 {
	return s.Cps
}

func (s *DescribePortConnsCountResponseBody) GetInActConns() *int64 {
	return s.InActConns
}

func (s *DescribePortConnsCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortConnsCountResponseBody) SetActConns(v int64) *DescribePortConnsCountResponseBody {
	s.ActConns = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetConns(v int64) *DescribePortConnsCountResponseBody {
	s.Conns = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetCps(v int64) *DescribePortConnsCountResponseBody {
	s.Cps = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetInActConns(v int64) *DescribePortConnsCountResponseBody {
	s.InActConns = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) SetRequestId(v string) *DescribePortConnsCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortConnsCountResponseBody) Validate() error {
	return dara.Validate(s)
}
