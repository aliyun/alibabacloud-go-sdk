// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDatasetVersionRequest
	GetDescription() *string
	SetVersion(v string) *CreateDatasetVersionRequest
	GetVersion() *string
	SetClientToken(v string) *CreateDatasetVersionRequest
	GetClientToken() *string
}

type CreateDatasetVersionRequest struct {
	// The description of the version. This parameter is optional. If this parameter is not specified, an empty string is used.
	//
	// example:
	//
	// Dataset version used for the first evaluation
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the version to be created, which is used to identify the version within the current dataset. The name cannot be empty after leading and trailing whitespaces are removed. The length of the UTF-8 encoded name cannot exceed 128 bytes, and the name cannot contain forward slashes (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The optional client token that is generated and passed by the client. Do not use this token to determine whether the version is created. You can call ListDatasetVersions to query the version.
	//
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetVersionRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateDatasetVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDatasetVersionRequest) SetDescription(v string) *CreateDatasetVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetVersion(v string) *CreateDatasetVersionRequest {
	s.Version = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetClientToken(v string) *CreateDatasetVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
