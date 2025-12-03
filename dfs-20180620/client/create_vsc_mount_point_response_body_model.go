// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountPointId(v string) *CreateVscMountPointResponseBody
	GetMountPointId() *string
	SetRequestId(v string) *CreateVscMountPointResponseBody
	GetRequestId() *string
}

type CreateVscMountPointResponseBody struct {
	// example:
	//
	// e389e5c7-bcb4-4558-846a-e5afc444****
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVscMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVscMountPointResponseBody) GetMountPointId() *string {
	return s.MountPointId
}

func (s *CreateVscMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVscMountPointResponseBody) SetMountPointId(v string) *CreateVscMountPointResponseBody {
	s.MountPointId = &v
	return s
}

func (s *CreateVscMountPointResponseBody) SetRequestId(v string) *CreateVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVscMountPointResponseBody) Validate() error {
	return dara.Validate(s)
}
