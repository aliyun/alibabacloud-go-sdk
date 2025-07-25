// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpamResponseBody
	GetRequestId() *string
}

type DeleteIpamResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30A20EE2-6223-5D0F-BF49-D7C78F206063
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpamResponseBody) SetRequestId(v string) *DeleteIpamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpamResponseBody) Validate() error {
	return dara.Validate(s)
}
