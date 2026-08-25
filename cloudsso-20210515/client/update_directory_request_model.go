// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateDirectoryRequest
	GetDirectoryId() *string
	SetNewDirectoryName(v string) *UpdateDirectoryRequest
	GetNewDirectoryName() *string
}

type UpdateDirectoryRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new name of the directory. The name must be globally unique.
	//
	// The name can contain lowercase letters, digits, and hyphens (-). The name cannot start or end with a hyphen (-) and cannot contain two consecutive hyphens (-). If you want the new name of the directory to start with `d-`, you must set this parameter to the ID of the directory.
	//
	// The name must be 2 to 64 characters in length.
	//
	// example:
	//
	// new-example
	NewDirectoryName *string `json:"NewDirectoryName,omitempty" xml:"NewDirectoryName,omitempty"`
}

func (s UpdateDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateDirectoryRequest) GetNewDirectoryName() *string {
	return s.NewDirectoryName
}

func (s *UpdateDirectoryRequest) SetDirectoryId(v string) *UpdateDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDirectoryRequest) SetNewDirectoryName(v string) *UpdateDirectoryRequest {
	s.NewDirectoryName = &v
	return s
}

func (s *UpdateDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
