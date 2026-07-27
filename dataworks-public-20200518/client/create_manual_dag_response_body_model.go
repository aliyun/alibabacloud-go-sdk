// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManualDagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *CreateManualDagResponseBody
	GetDagId() *int64
	SetRequestId(v string) *CreateManualDagResponseBody
	GetRequestId() *string
}

type CreateManualDagResponseBody struct {
	// The instance ID of the DAG generated when the manual workflow runs. You can use this DagId together with the relevant API to query the details and status of the internal node instances of this manual workflow run.
	//
	// example:
	//
	// 700000123123141
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// AASFDFSDFG-DFSDF-DFSDFD-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateManualDagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateManualDagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateManualDagResponseBody) GetDagId() *int64 {
	return s.DagId
}

func (s *CreateManualDagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateManualDagResponseBody) SetDagId(v int64) *CreateManualDagResponseBody {
	s.DagId = &v
	return s
}

func (s *CreateManualDagResponseBody) SetRequestId(v string) *CreateManualDagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateManualDagResponseBody) Validate() error {
	return dara.Validate(s)
}
