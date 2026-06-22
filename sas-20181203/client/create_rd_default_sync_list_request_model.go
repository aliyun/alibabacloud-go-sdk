// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRdDefaultSyncListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderIds(v string) *CreateRdDefaultSyncListRequest
	GetFolderIds() *string
}

type CreateRdDefaultSyncListRequest struct {
	// The list of folder IDs in the resource directory.
	//
	// > You can call [GetRdTree](~~GetRdTree~~) to obtain folder IDs. Separate multiple folder IDs with commas (,). If you do not specify this parameter, the existing configuration list is cleared.
	//
	// example:
	//
	// fd-BwoXuf****,fd-CFamY7****
	FolderIds *string `json:"FolderIds,omitempty" xml:"FolderIds,omitempty"`
}

func (s CreateRdDefaultSyncListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRdDefaultSyncListRequest) GoString() string {
	return s.String()
}

func (s *CreateRdDefaultSyncListRequest) GetFolderIds() *string {
	return s.FolderIds
}

func (s *CreateRdDefaultSyncListRequest) SetFolderIds(v string) *CreateRdDefaultSyncListRequest {
	s.FolderIds = &v
	return s
}

func (s *CreateRdDefaultSyncListRequest) Validate() error {
	return dara.Validate(s)
}
