// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaWorkloadsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListQuotaWorkloadsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListQuotaWorkloadsResponseBody
	GetTotalCount() *int64
	SetWorkloads(v []*QueueInfo) *ListQuotaWorkloadsResponseBody
	GetWorkloads() []*QueueInfo
}

type ListQuotaWorkloadsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 42F23B58-3684-5443-848A-8DA81FF99712
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalCount *int64       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Workloads  []*QueueInfo `json:"Workloads,omitempty" xml:"Workloads,omitempty" type:"Repeated"`
}

func (s ListQuotaWorkloadsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaWorkloadsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaWorkloadsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotaWorkloadsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQuotaWorkloadsResponseBody) GetWorkloads() []*QueueInfo {
	return s.Workloads
}

func (s *ListQuotaWorkloadsResponseBody) SetRequestId(v string) *ListQuotaWorkloadsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaWorkloadsResponseBody) SetTotalCount(v int64) *ListQuotaWorkloadsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotaWorkloadsResponseBody) SetWorkloads(v []*QueueInfo) *ListQuotaWorkloadsResponseBody {
	s.Workloads = v
	return s
}

func (s *ListQuotaWorkloadsResponseBody) Validate() error {
	return dara.Validate(s)
}
