// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishMigrationStageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *FinishMigrationStageRequest
	GetInstanceId() *string
}

type FinishMigrationStageRequest struct {
	// example:
	//
	// rmq-cn-pe334f1nx04
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s FinishMigrationStageRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishMigrationStageRequest) GoString() string {
	return s.String()
}

func (s *FinishMigrationStageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FinishMigrationStageRequest) SetInstanceId(v string) *FinishMigrationStageRequest {
	s.InstanceId = &v
	return s
}

func (s *FinishMigrationStageRequest) Validate() error {
	return dara.Validate(s)
}
