// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RenameFolderRequest
	GetName() *string
}

type RenameFolderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RenameFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameFolderRequest) GoString() string {
	return s.String()
}

func (s *RenameFolderRequest) GetName() *string {
	return s.Name
}

func (s *RenameFolderRequest) SetName(v string) *RenameFolderRequest {
	s.Name = &v
	return s
}

func (s *RenameFolderRequest) Validate() error {
	return dara.Validate(s)
}
