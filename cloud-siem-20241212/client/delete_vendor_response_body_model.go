// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVendorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVendorResponseBody
	GetRequestId() *string
}

type DeleteVendorResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVendorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVendorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVendorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVendorResponseBody) SetRequestId(v string) *DeleteVendorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVendorResponseBody) Validate() error {
	return dara.Validate(s)
}
