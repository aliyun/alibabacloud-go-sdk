// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToMountPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachVscToMountPointsResponseBody
	GetRequestId() *string
}

type AttachVscToMountPointsResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachVscToMountPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToMountPointsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVscToMountPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachVscToMountPointsResponseBody) SetRequestId(v string) *AttachVscToMountPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachVscToMountPointsResponseBody) Validate() error {
	return dara.Validate(s)
}
