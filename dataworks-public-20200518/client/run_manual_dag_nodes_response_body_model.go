// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunManualDagNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *RunManualDagNodesResponseBody
	GetDagId() *int64
	SetRequestId(v string) *RunManualDagNodesResponseBody
	GetRequestId() *string
}

type RunManualDagNodesResponseBody struct {
	// The ID of the DAG for the manually triggered workflow. You can call an operation with this parameter as a request parameter to query the details and states of the nodes in the manually triggered workflow.
	//
	// example:
	//
	// 700000123123141
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AASFDFSDFG-DFSDF-DFSDFD-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunManualDagNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunManualDagNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RunManualDagNodesResponseBody) GetDagId() *int64 {
	return s.DagId
}

func (s *RunManualDagNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunManualDagNodesResponseBody) SetDagId(v int64) *RunManualDagNodesResponseBody {
	s.DagId = &v
	return s
}

func (s *RunManualDagNodesResponseBody) SetRequestId(v string) *RunManualDagNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunManualDagNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
