// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMountPointResponseBody
	GetRequestId() *string
}

type DeleteMountPointResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMountPointResponseBody) SetRequestId(v string) *DeleteMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMountPointResponseBody) Validate() error {
	return dara.Validate(s)
}
