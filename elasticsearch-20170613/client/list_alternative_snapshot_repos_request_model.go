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
	// Specifies whether to return the OSS reference repositories that have already been added. Valid values:
	//
	// - true (default): Returns the already added repositories.
	//
	// - false: Does not return the already added repositories.
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
