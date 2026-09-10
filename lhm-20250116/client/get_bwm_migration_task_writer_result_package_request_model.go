// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationTaskWriterResultPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetBwmMigrationTaskWriterResultPackageRequest
	GetInstanceId() *string
}

type GetBwmMigrationTaskWriterResultPackageRequest struct {
	// The submit instance ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s GetBwmMigrationTaskWriterResultPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationTaskWriterResultPackageRequest) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationTaskWriterResultPackageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBwmMigrationTaskWriterResultPackageRequest) SetInstanceId(v string) *GetBwmMigrationTaskWriterResultPackageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetBwmMigrationTaskWriterResultPackageRequest) Validate() error {
	return dara.Validate(s)
}
