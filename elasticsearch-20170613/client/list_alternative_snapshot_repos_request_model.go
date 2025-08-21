// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlternativeSnapshotReposRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlreadySetItems(v bool) *ListAlternativeSnapshotReposRequest
	GetAlreadySetItems() *bool
}

type ListAlternativeSnapshotReposRequest struct {
	// Indicates whether to return the OSS reference repository added. The return value. Valid values: true and false.
	//
	// example:
	//
	// true
	AlreadySetItems *bool `json:"alreadySetItems,omitempty" xml:"alreadySetItems,omitempty"`
}

func (s ListAlternativeSnapshotReposRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlternativeSnapshotReposRequest) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposRequest) GetAlreadySetItems() *bool {
	return s.AlreadySetItems
}

func (s *ListAlternativeSnapshotReposRequest) SetAlreadySetItems(v bool) *ListAlternativeSnapshotReposRequest {
	s.AlreadySetItems = &v
	return s
}

func (s *ListAlternativeSnapshotReposRequest) Validate() error {
	return dara.Validate(s)
}
