// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindVscMountPointAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountPointAlias(v string) *BindVscMountPointAliasResponseBody
	GetMountPointAlias() *string
	SetRequestId(v string) *BindVscMountPointAliasResponseBody
	GetRequestId() *string
}

type BindVscMountPointAliasResponseBody struct {
	// example:
	//
	// sdfe
	MountPointAlias *string `json:"MountPointAlias,omitempty" xml:"MountPointAlias,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindVscMountPointAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindVscMountPointAliasResponseBody) GoString() string {
	return s.String()
}

func (s *BindVscMountPointAliasResponseBody) GetMountPointAlias() *string {
	return s.MountPointAlias
}

func (s *BindVscMountPointAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindVscMountPointAliasResponseBody) SetMountPointAlias(v string) *BindVscMountPointAliasResponseBody {
	s.MountPointAlias = &v
	return s
}

func (s *BindVscMountPointAliasResponseBody) SetRequestId(v string) *BindVscMountPointAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindVscMountPointAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
