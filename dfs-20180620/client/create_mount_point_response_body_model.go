// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountPointId(v string) *CreateMountPointResponseBody
	GetMountPointId() *string
	SetRequestId(v string) *CreateMountPointResponseBody
	GetRequestId() *string
}

type CreateMountPointResponseBody struct {
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMountPointResponseBody) GetMountPointId() *string {
	return s.MountPointId
}

func (s *CreateMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMountPointResponseBody) SetMountPointId(v string) *CreateMountPointResponseBody {
	s.MountPointId = &v
	return s
}

func (s *CreateMountPointResponseBody) SetRequestId(v string) *CreateMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMountPointResponseBody) Validate() error {
	return dara.Validate(s)
}
