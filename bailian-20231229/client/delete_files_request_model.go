// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileIds(v []*string) *DeleteFilesRequest
	GetFileIds() []*string
}

type DeleteFilesRequest struct {
	// This parameter is required.
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
}

func (s DeleteFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilesRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *DeleteFilesRequest) SetFileIds(v []*string) *DeleteFilesRequest {
	s.FileIds = v
	return s
}

func (s *DeleteFilesRequest) Validate() error {
	return dara.Validate(s)
}
