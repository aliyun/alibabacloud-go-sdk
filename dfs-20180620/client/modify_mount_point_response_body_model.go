// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMountPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMountPointResponseBody
	GetRequestId() *string
}

type ModifyMountPointResponseBody struct {
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMountPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMountPointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMountPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMountPointResponseBody) SetRequestId(v string) *ModifyMountPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMountPointResponseBody) Validate() error {
	return dara.Validate(s)
}
