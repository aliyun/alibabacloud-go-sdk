// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSpareIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSpareIpsResponseBody
	GetRequestId() *string
}

type CreateSpareIpsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSpareIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSpareIpsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSpareIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSpareIpsResponseBody) SetRequestId(v string) *CreateSpareIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSpareIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
