// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAegisForLingjunStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuidsShrink(v string) *ListAegisForLingjunStatusShrinkRequest
	GetUuidsShrink() *string
}

type ListAegisForLingjunStatusShrinkRequest struct {
	// List of unique UUIDs for Lingjun bare metal.
	UuidsShrink *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ListAegisForLingjunStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAegisForLingjunStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAegisForLingjunStatusShrinkRequest) GetUuidsShrink() *string {
	return s.UuidsShrink
}

func (s *ListAegisForLingjunStatusShrinkRequest) SetUuidsShrink(v string) *ListAegisForLingjunStatusShrinkRequest {
	s.UuidsShrink = &v
	return s
}

func (s *ListAegisForLingjunStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
