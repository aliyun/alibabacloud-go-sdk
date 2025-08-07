// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateBackupRequest
	GetDescription() *string
	SetServiceInstanceId(v string) *CreateBackupRequest
	GetServiceInstanceId() *string
}

type CreateBackupRequest struct {
	// Backup description
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the service instance to be transferred to official version.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-56cfb7ec3a634448a96c
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateBackupRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateBackupRequest) SetDescription(v string) *CreateBackupRequest {
	s.Description = &v
	return s
}

func (s *CreateBackupRequest) SetServiceInstanceId(v string) *CreateBackupRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateBackupRequest) Validate() error {
	return dara.Validate(s)
}
