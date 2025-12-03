// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateBackupPlanRequest
	GetClusterId() *string
}

type CreateBackupPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-wz94lbcqc****4x93
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateBackupPlanRequest) SetClusterId(v string) *CreateBackupPlanRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
