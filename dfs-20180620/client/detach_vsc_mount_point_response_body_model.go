// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachVscMountPointResponseBody
	GetRequestId() *string
}

type DetachVscMountPointResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachVscMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *DetachVscMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachVscMountPointResponseBody) SetRequestId(v string) *DetachVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachVscMountPointResponseBody) Validate() error {
	return dara.Validate(s)
}
