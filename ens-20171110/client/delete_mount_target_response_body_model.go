// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMountTargetResponseBody
	GetRequestId() *string
}

type DeleteMountTargetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMountTargetResponseBody) SetRequestId(v string) *DeleteMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMountTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
