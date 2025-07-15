// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcClassicLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableVpcClassicLinkResponseBody
	GetRequestId() *string
}

type DisableVpcClassicLinkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVpcClassicLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcClassicLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVpcClassicLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableVpcClassicLinkResponseBody) SetRequestId(v string) *DisableVpcClassicLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableVpcClassicLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
