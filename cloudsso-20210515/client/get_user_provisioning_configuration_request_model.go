// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetUserProvisioningConfigurationRequest
	GetDirectoryId() *string
}

type GetUserProvisioningConfigurationRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
}

func (s GetUserProvisioningConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningConfigurationRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningConfigurationRequest) SetDirectoryId(v string) *GetUserProvisioningConfigurationRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
