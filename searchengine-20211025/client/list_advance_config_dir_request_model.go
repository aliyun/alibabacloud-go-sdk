// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdvanceConfigDirRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirName(v string) *ListAdvanceConfigDirRequest
	GetDirName() *string
}

type ListAdvanceConfigDirRequest struct {
	// The name of the directory
	//
	// This parameter is required.
	//
	// example:
	//
	// /zones/general
	DirName *string `json:"dirName,omitempty" xml:"dirName,omitempty"`
}

func (s ListAdvanceConfigDirRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigDirRequest) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigDirRequest) GetDirName() *string {
	return s.DirName
}

func (s *ListAdvanceConfigDirRequest) SetDirName(v string) *ListAdvanceConfigDirRequest {
	s.DirName = &v
	return s
}

func (s *ListAdvanceConfigDirRequest) Validate() error {
	return dara.Validate(s)
}
