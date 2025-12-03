// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVscMountPointResponseBody
	GetRequestId() *string
}

type DeleteVscMountPointResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVscMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVscMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVscMountPointResponseBody) SetRequestId(v string) *DeleteVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVscMountPointResponseBody) Validate() error {
	return dara.Validate(s)
}
