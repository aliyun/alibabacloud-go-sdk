// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVodPackagingGroupResponseBody
	GetRequestId() *string
}

type DeleteVodPackagingGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 13cbb83e-043c-4728-ac35-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodPackagingGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVodPackagingGroupResponseBody) SetRequestId(v string) *DeleteVodPackagingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodPackagingGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
