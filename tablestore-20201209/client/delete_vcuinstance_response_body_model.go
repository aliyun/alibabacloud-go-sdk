// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVCUInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVCUInstanceResponseBody
	GetRequestId() *string
}

type DeleteVCUInstanceResponseBody struct {
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVCUInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVCUInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVCUInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVCUInstanceResponseBody) SetRequestId(v string) *DeleteVCUInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVCUInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
