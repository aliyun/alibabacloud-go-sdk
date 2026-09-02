// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchChangeTableOwnerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *GetBatchChangeTableOwnerStatusRequest
	GetBatchId() *string
}

type GetBatchChangeTableOwnerStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 524257_openapi-req-abc123
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s GetBatchChangeTableOwnerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchChangeTableOwnerStatusRequest) GoString() string {
	return s.String()
}

func (s *GetBatchChangeTableOwnerStatusRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *GetBatchChangeTableOwnerStatusRequest) SetBatchId(v string) *GetBatchChangeTableOwnerStatusRequest {
	s.BatchId = &v
	return s
}

func (s *GetBatchChangeTableOwnerStatusRequest) Validate() error {
	return dara.Validate(s)
}
