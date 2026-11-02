// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryName(v string) *CreateDirectoryRequest
	GetDirectoryName() *string
}

type CreateDirectoryRequest struct {
	// The name of the directory. The name must be globally unique.
	//
	// Format: The name can contain lowercase letters, digits, and hyphens (-). It cannot start or end with a hyphen (-) and cannot contain two consecutive hyphens (-). It cannot start with `d-`.
	//
	// Length: 2 to 64 characters.
	//
	// > If you do not specify this parameter, the system automatically generates a name.
	//
	// example:
	//
	// example
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
}

func (s CreateDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequest) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *CreateDirectoryRequest) SetDirectoryName(v string) *CreateDirectoryRequest {
	s.DirectoryName = &v
	return s
}

func (s *CreateDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
