// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSpareIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSpareIpsResponseBody
	GetRequestId() *string
}

type DeleteSpareIpsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSpareIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSpareIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSpareIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSpareIpsResponseBody) SetRequestId(v string) *DeleteSpareIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSpareIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
