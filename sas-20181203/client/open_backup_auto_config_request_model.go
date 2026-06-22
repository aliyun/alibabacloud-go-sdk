// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBackupAutoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxBatchSize(v int32) *OpenBackupAutoConfigRequest
	GetMaxBatchSize() *int32
}

type OpenBackupAutoConfigRequest struct {
	// The number of servers included in a single batch when the anti-ransomware managed service automatically generates policies.
	//
	// > The maximum value is 50. If you specify a value greater than 50, the value is set to 50.
	//
	// example:
	//
	// 20
	MaxBatchSize *int32 `json:"MaxBatchSize,omitempty" xml:"MaxBatchSize,omitempty"`
}

func (s OpenBackupAutoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenBackupAutoConfigRequest) GoString() string {
	return s.String()
}

func (s *OpenBackupAutoConfigRequest) GetMaxBatchSize() *int32 {
	return s.MaxBatchSize
}

func (s *OpenBackupAutoConfigRequest) SetMaxBatchSize(v int32) *OpenBackupAutoConfigRequest {
	s.MaxBatchSize = &v
	return s
}

func (s *OpenBackupAutoConfigRequest) Validate() error {
	return dara.Validate(s)
}
