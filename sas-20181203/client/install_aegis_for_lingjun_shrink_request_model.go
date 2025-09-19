// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAegisForLingjunShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuidsShrink(v string) *InstallAegisForLingjunShrinkRequest
	GetUuidsShrink() *string
}

type InstallAegisForLingjunShrinkRequest struct {
	// List of unique UUIDs for Lingjun bare metal.
	UuidsShrink *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s InstallAegisForLingjunShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAegisForLingjunShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallAegisForLingjunShrinkRequest) GetUuidsShrink() *string {
	return s.UuidsShrink
}

func (s *InstallAegisForLingjunShrinkRequest) SetUuidsShrink(v string) *InstallAegisForLingjunShrinkRequest {
	s.UuidsShrink = &v
	return s
}

func (s *InstallAegisForLingjunShrinkRequest) Validate() error {
	return dara.Validate(s)
}
