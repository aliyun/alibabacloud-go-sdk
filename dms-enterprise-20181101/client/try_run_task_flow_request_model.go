// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTryRunTaskFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *TryRunTaskFlowRequest
	GetDagId() *int64
}

type TryRunTaskFlowRequest struct {
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
}

func (s TryRunTaskFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s TryRunTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *TryRunTaskFlowRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *TryRunTaskFlowRequest) SetDagId(v int64) *TryRunTaskFlowRequest {
	s.DagId = &v
	return s
}

func (s *TryRunTaskFlowRequest) Validate() error {
	return dara.Validate(s)
}
