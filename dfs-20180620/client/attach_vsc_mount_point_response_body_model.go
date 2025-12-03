// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachVscMountPointResponseBody
	GetRequestId() *string
}

type AttachVscMountPointResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachVscMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachVscMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *AttachVscMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachVscMountPointResponseBody) SetRequestId(v string) *AttachVscMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachVscMountPointResponseBody) Validate() error {
	return dara.Validate(s)
}
