// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstanceVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *ListAgentInstanceVersionsResponseBody
	GetData() []*string
	SetRequestId(v string) *ListAgentInstanceVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAgentInstanceVersionsResponseBody
	GetTotalCount() *int32
}

type ListAgentInstanceVersionsResponseBody struct {
	// The list of agent instance version numbers, which are deduplicated and sorted in ascending order.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4C5A2E8F-9B3D-4E7A-A1C6-2F8D0E5B7A91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total count.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentInstanceVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceVersionsResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListAgentInstanceVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentInstanceVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentInstanceVersionsResponseBody) SetData(v []*string) *ListAgentInstanceVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentInstanceVersionsResponseBody) SetRequestId(v string) *ListAgentInstanceVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentInstanceVersionsResponseBody) SetTotalCount(v int32) *ListAgentInstanceVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentInstanceVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}
