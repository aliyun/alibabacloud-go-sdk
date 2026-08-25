// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupClientResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIdsShrink(v string) *DeleteBackupClientResourceShrinkRequest
	GetClientIdsShrink() *string
}

type DeleteBackupClientResourceShrinkRequest struct {
	// A list of client IDs. The list can contain up to 100 client IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["c-*********************"]
	ClientIdsShrink *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
}

func (s DeleteBackupClientResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupClientResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceShrinkRequest) GetClientIdsShrink() *string {
	return s.ClientIdsShrink
}

func (s *DeleteBackupClientResourceShrinkRequest) SetClientIdsShrink(v string) *DeleteBackupClientResourceShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *DeleteBackupClientResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
